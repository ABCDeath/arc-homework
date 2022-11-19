import itertools
import re
import string
from typing import Iterable

from generator.generator import CodeGeneratorBase


class AdapterGenerator(CodeGeneratorBase):
    _TEMPLATE_PATH = 'adapter'
    _STRUCT_NAME = 'Adapter'
    _INTERFACE_TYPE_RE = re.compile(r'(?<=type\s)\w+(?=\sinterface)')

    def generate(self) -> str:
        interface_def, *body_and_the_rest = self._content[self._line:]

        interface_name = self._INTERFACE_TYPE_RE.findall(interface_def)[0]
        struct_name = self._build_adapter_name(interface_name)

        interface_body = tuple(
            itertools.takewhile(lambda s: '}' not in s, body_and_the_rest)
        )
        methods = self._build_methods(
            struct_name=struct_name,
            dependency_module=self._dependency_module,
            body=interface_body,
        )
        imports = self._build_imports(interface_body)

        return self._build_file(
            imports=imports,
            struct_name=struct_name,
            interface_name=interface_name,
            implementation='\n'.join(methods).strip(),
        )

    def _build_adapter_name(self, interface_name: str) -> str:
        name = f'{interface_name}{self._STRUCT_NAME}'

        return f'{name[0].swapcase()}{name[1:]}'

    def _build_methods(
        self,
        struct_name: str,
        dependency_module: str,
        body: Iterable[str],
    ) -> tuple[str, ...]:
        return tuple(
            self._build_method(
                receiver_name=struct_name,
                signature=signature,
                dependency_module=dependency_module,
            )
            for signature in body
        )

    def _build_method(
        self,
        receiver_name: str,
        signature: str,
        dependency_module: str,
    ) -> str:
        template = string.Template(self._templates['method'])
        signature = signature.strip()
        dependency = self._get_dependency(signature)
        method_args, method_return = self._parse_signature(signature)
        resolve_args = ''

        if method_args:
            resolve_args = ', {}'.format(', '.join(method_args.keys()))

        return template.substitute(
            receiver_name=receiver_name,
            method_signature=signature,
            module=dependency_module,
            dependency=dependency,
            resolve_args=resolve_args,
            return_type=method_return[0],
        )

    @staticmethod
    def _get_dependency(signature: str) -> str:
        method_name = signature.split('(', maxsplit=1)[0]

        if 'Get' in method_name:
            return '{}.Get'.format(method_name.split('Get')[1])

        if 'Set' in method_name:
            return '{}.Set'.format(method_name.split('Set')[1])

        return method_name

    def _parse_signature(self, signature: str) -> tuple[dict[str, str], tuple[str, ...]]:
        return (
            self._parse_signature_for_args(signature),
            self._parse_signature_for_return(signature),
        )

    @staticmethod
    def _parse_signature_for_args(signature: str) -> dict[str, str]:
        args_substring_slice = slice(signature.find('(') + 1, signature.find(')'))
        args = signature[args_substring_slice].split(', ')
        if not args or args[0] == '':
            return {}

        return dict(s.split(' ') for s in args)

    @staticmethod
    def _parse_signature_for_return(signature: str) -> tuple[str, ...]:
        signature = signature[signature.find(')') + 1:]
        ret_closing = signature.rfind(')')
        if ret_closing < 0:
            ret_closing = len(signature)

        ret_substring_slice = slice(signature.find('(') + 1, ret_closing + 1)
        ret_substring = signature[ret_substring_slice].strip().replace('(', '').replace(')', '')

        return tuple(ret_substring.split(', '))

    def _build_imports(self, interface_body: Iterable[str]) -> str:
        current_imports = self._get_current_imports()
        required_modules = set(self._get_required_modules(interface_body))

        imports = tuple(i for m, i in current_imports.items() if m in required_modules)
        imports = imports + ('"arc-homework/space-game/ioc"', '"arc-homework/space-game/moving/object"')

        imports_string = '\n'.join(f'\t{i}' for i in sorted(imports))

        return (
            'import (\n'
            '\t"context"\n'
            '\n'
            f'{imports_string}\n'
            ')'
        )

    def _get_current_imports(self) -> dict[str, str]:
        import_position = -1
        for i, s in enumerate(self._content[:self._line]):
            if 'import' in s:
                import_position = i

                break

        if import_position < 0:
            return {}

        import_closing = import_position + 1
        for i, s in enumerate(self._content[import_position:self._line]):
            if ')' in s:
                import_closing = import_position + i

                break

        return {
            i.rsplit('/', maxsplit=1)[-1].strip().replace('"', ''): i.strip()
            for i in self._content[import_position + 1:import_closing]
        }

    def _get_required_modules(self, interface_body: Iterable[str]) -> tuple[str, ...]:
        return_types = itertools.chain(
            *tuple(self._parse_signature_for_return(s) for s in interface_body)
        )
        filtered_current_package_and_built_in = filter(
            lambda t: '.' in t, return_types
        )

        return tuple(
            t.split('.', maxsplit=1)[0] for t in filtered_current_package_and_built_in
        )

    def _build_file(
        self,
        imports: str,
        struct_name: str,
        interface_name: str,
        implementation: str,
    ) -> str:
        template = string.Template(self._templates['module'])

        return template.substitute(
            package=self._go_package,
            imports=imports,
            struct_name=struct_name,
            implementation=implementation,
            interface_name=interface_name,
        )
