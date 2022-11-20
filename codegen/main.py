import os
from typing import List

from args import get_codegen_params
from generator.factory import generator_factory


_FILE_EXTENSION = 'go'


def _read_source(filename: str) -> List[str]:
    with open(filename) as f:
        return f.readlines()


def _save_code(filename: str, code: str) -> None:
    with open(filename, 'w') as f:
        f.write(code)


if __name__ == '__main__':
    codegen_root = file_path = os.path.realpath(__file__).rsplit('/', maxsplit=1)[0]
    codegen_params = get_codegen_params()

    generator_class = generator_factory(codegen_params.object_type)
    generator = generator_class(
        root=codegen_root,
        go_package=codegen_params.go_package,
        dependency_module=codegen_params.dependency_module,
        content=_read_source(codegen_params.file),
        line=codegen_params.line,
    )
    generated_code = generator.generate()

    generated_filename = os.path.join(
        codegen_params.file.rsplit('/')[0],
        f'{codegen_params.object_type}.{_FILE_EXTENSION}',
    )
    _save_code(generated_filename, generated_code)
