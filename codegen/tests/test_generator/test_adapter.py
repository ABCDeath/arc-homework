from generator.adapter import AdapterGenerator


class TestAdapterGeneratorGetDependency:
    def test_returns_name_for_getter_method_signature(self):
        actual = AdapterGenerator._get_dependency('GetPosition')

        assert 'Position.Get' == actual

    def test_returns_name_for_setter_method_signature(self):
        actual = AdapterGenerator._get_dependency('SetPosition')

        assert 'Position.Set' == actual

    def test_returns_name_for_any_other_method_signature(self):
        method_name = 'DoSomething'
        actual = AdapterGenerator._get_dependency(method_name)

        assert method_name == actual


class TestAdapterGeneratorInit:
    def test_init_initializes_templates(self):
        generator = AdapterGenerator(
            root='', go_package='', dependency_module='', content='', line=0,
        )

        assert 'module' in generator._templates
        with open('template/adapter/module.go') as f:
            assert f.read() == generator._templates['module']

        assert 'method' in generator._templates
        with open('template/adapter/method.go') as f:
            assert f.read() == generator._templates['method']


class TestAdapterGeneratorBuildMethods:
    def test_parse_signature_for_args_returns_arg_names_with_types_as_dict(self):
        actual = AdapterGenerator._parse_signature_for_args(
            'FooBar(foo int, bar []byte, buzz map[int]SomeStructure) (int, error)'
        )
        expected = {
            'foo': 'int',
            'bar': '[]byte',
            'buzz': 'map[int]SomeStructure',
        }

        assert expected == actual

    def test_parse_signature_for_return_returns_tuple_of_return_types(self):
        actual = AdapterGenerator._parse_signature_for_return(
            'FooBar(foo int, bar []byte, buzz map[int]SomeStructure) (int, map[int]SomeStructure, error)'
        )
        expected = ('int', 'map[int]SomeStructure', 'error')

        assert expected == actual

    def test_parse_signature_returns_dict_of_args_and_tuple_of_return_types(self):
        generator = AdapterGenerator(
            root='', go_package='', dependency_module='', content='', line=0,
        )
        actual = generator._parse_signature(
            'FooBar(foo int, bar []byte, buzz map[int]SomeStructure) (int, map[int]SomeStructure, error)'
        )
        expected = (
            {
                'foo': 'int',
                'bar': '[]byte',
                'buzz': 'map[int]SomeStructure',
            },
            ('int', 'map[int]SomeStructure', 'error')
        )

        assert expected == actual

    def test_builds_method_string_from_template(self):
        generator = AdapterGenerator(
            root='', go_package='', dependency_module='', content='', line=0,
        )

        actual = generator._build_method(
            receiver_name='receiver',
            signature='FooMethod(arg1 int, arg2 []string, arg3 map[string]Bar) (vector.Vector, error)',
            dependency_module='SomeDependency',
        )

        expected = (
            'func (a *receiver) FooMethod(arg1 int, arg2 []string, arg3 map[string]Bar) (vector.Vector, error) {\n'
            '\tresolvedDependency, err := a.ioc.Resolve(ctx, "SomeDependency:FooMethod", a.obj, arg1, arg2, arg3)\n'
            '\tif err != nil {\n'
            '\t\treturn vector.Vector{}, err\n'
            '\t}\n'
            '\n'
            '\tresolvedDependencyCast, ok := resolvedDependency.(vector.Vector)\n'
            '\tif !ok {\n'
            '\t\treturn vector.Vector{}, err\n'
            '\t}\n'
            '\n'
            '\treturn resolvedDependencyCast, nil\n'
            '}\n'
        )
        assert expected == actual

    def test_builds_method_strings_from_template(self):
        generator = AdapterGenerator(
            root='', go_package='', dependency_module='', content='', line=0,
        )

        actual = generator._build_methods(
            struct_name='foo',
            dependency_module='SomeDependency',
            body=(
                'FooMethod(arg1 int, arg2 []string, arg3 map[string]Bar) (vector.Vector, error)',
                'BarMethod() (point.Point, error)'
            ),
        )

        expected = (
            (
                'func (a *foo) FooMethod(arg1 int, arg2 []string, arg3 map[string]Bar) (vector.Vector, error) {\n'
                '\tresolvedDependency, err := a.ioc.Resolve(ctx, "SomeDependency:FooMethod", a.obj, arg1, arg2, arg3)\n'
                '\tif err != nil {\n'
                '\t\treturn vector.Vector{}, err\n'
                '\t}\n'
                '\n'
                '\tresolvedDependencyCast, ok := resolvedDependency.(vector.Vector)\n'
                '\tif !ok {\n'
                '\t\treturn vector.Vector{}, err\n'
                '\t}\n'
                '\n'
                '\treturn resolvedDependencyCast, nil\n'
                '}\n'
            ),
            (
                'func (a *foo) BarMethod() (point.Point, error) {\n'
                '\tresolvedDependency, err := a.ioc.Resolve(ctx, "SomeDependency:BarMethod", a.obj)\n'
                '\tif err != nil {\n'
                '\t\treturn point.Point{}, err\n'
                '\t}\n'
                '\n'
                '\tresolvedDependencyCast, ok := resolvedDependency.(point.Point)\n'
                '\tif !ok {\n'
                '\t\treturn point.Point{}, err\n'
                '\t}\n'
                '\n'
                '\treturn resolvedDependencyCast, nil\n'
                '}\n'
            ),
        )
        assert expected == actual


class TestAdapterGeneratorBuildImports:
    def test_get_required_modules_returns_modules_from_interface_method_signatures(self):
        generator = AdapterGenerator(
            root='', go_package='foobar', dependency_module='', content='', line=0,
        )

        interface_body = (
            '\tFoo() (vector.Vector, error)\n',
            '\tBar(arg map[int]string) error\n',
            '\tGetSomething() point.Point\n',
            '\tSetSomething(v int)\n',
            '}\n',
            ''
        )

        actual = generator._get_required_modules(interface_body)

        expected = ('vector', 'point')
        assert expected == actual

    def test_get_current_imports_returns_import_strings_by_module_dict(self):
        content = (
            'package foo\n',
            '\n',
            'import (\n',
            '\t"errors"\n',
            '\t"math"\n',
            '\t"arc-homework/space-game/moving/object"\n',
            '\t"arc-homework/space-game/moving/vector"\n',
            ')\n',
        )
        generator = AdapterGenerator(
            root='', go_package='foobar', dependency_module='', content=content, line=999,
        )

        actual = generator._get_current_imports()
        expected = {
            'errors': '"errors"',
            'math': '"math"',
            'object': '"arc-homework/space-game/moving/object"',
            'vector': '"arc-homework/space-game/moving/vector"',
        }
        assert expected == actual

    def test_build_imports_returns_imports_string_with_imports_required_by_adapter(self):
        interface_body = (
            'type FooBar interface {\n',
            '\tFoo() (vector.Vector, error)\n',
            '\tBar(arg map[int]string) error\n',
            '\tSetSomething(v int)\n',
            '}\n',
            '\n'
        )
        content = (
            'package foo\n',
            '\n',
            'import (\n',
            '\t"errors"\n',
            '\t"math"\n',
            '\t"arc-homework/space-game/moving/object"\n',
            '\t"arc-homework/space-game/moving/vector"\n',
            ')\n',
            '\n',
            '//go:generate blah-blah-blah\n',
            *interface_body,
        )
        generator = AdapterGenerator(
            root='', go_package='foobar', dependency_module='', content=content, line=9,
        )

        actual = generator._build_imports(interface_body)

        expected = (
            'import (\n'
            '\t"context"\n'
            '\n'
            '\t"arc-homework/space-game/ioc"\n'
            '\t"arc-homework/space-game/moving/object"\n'
            '\t"arc-homework/space-game/moving/vector"\n'
            ')'
        )
        assert expected == actual


class TestAdapterGeneratorBuildFile:
    def test_builds_file_content(self):
        generator = AdapterGenerator(
            root='', go_package='foobar', dependency_module='', content='', line=0,
        )

        imports = (
            'import (\n'
            '\t"some/path/module"\n'
            ')'
        )

        implementation = (
            'func (a *foo) FooMethod(arg1 int, arg2 []string, arg3 map[string]Bar) (vector.Vector, error) {\n'
            '\tresolvedDependency, err := a.ioc.Resolve(ctx, "SomeDependency:FooMethod", a.obj, arg1, arg2, arg3)\n'
            '\tif err != nil {\n'
            '\t\treturn vector.Vector{}, err\n'
            '\t}\n'
            '\n'
            '\tresolvedDependencyCast, ok := resolvedDependency.(vector.Vector)\n'
            '\tif !ok {\n'
            '\t\treturn vector.Vector{}, err\n'
            '\t}\n'
            '\n'
            '\treturn resolvedDependencyCast, nil\n'
            '}\n'
            'func (a *foo) BarMethod() (point.Point, error) {\n'
            '\tresolvedDependency, err := a.ioc.Resolve(ctx, "SomeDependency:BarMethod", a.obj)\n'
            '\tif err != nil {\n'
            '\t\treturn point.Point{}, err\n'
            '\t}\n'
            '\n'
            '\tresolvedDependencyCast, ok := resolvedDependency.(point.Point)\n'
            '\tif !ok {\n'
            '\t\treturn point.Point{}, err\n'
            '\t}\n'
            '\n'
            '\treturn resolvedDependencyCast, nil\n'
            '}\n'
        )
        actual = generator._build_file(
            imports=imports,
            struct_name='fooBarAdapter',
            interface_name='FooBar',
            implementation=implementation,
        )

        expected = (
            '// auto-generated file, do not edit\n'
            'package foobar\n'
            '\n'
            f'{imports}\n'
            '\n'
            'type fooBarAdapter struct {\n'
            '\tioc ioc.IoC\n'
            '\tobj object.Object\n'
            '}\n'
            '\n'
            f'{implementation}\n'
            '\n'
            'func NewfooBarAdapter(ioc ioc.IoC, obj object.Object) FooBar {\n'
            '\treturn &fooBarAdapter{\n'
            '\t\tioc: ioc,\n'
            '\t\tobj: obj,\n'
            '\t}\n'
            '}\n'
        )
        assert expected == actual


class TestAdapterGeneratorGenerate:
    def test_generates_module_content(self):
        interface_body = (
            'type FooBar interface {\n',
            '\tFooMethod(arg1 int, arg2 []string, arg3 map[string]Bar) (vector.Vector, error)\n',
            '\tBarMethod(v vector.Vector) (vector.Vector, error)\n',
            '}\n',
            '\n',
        )
        content = (
            'package foo\n',
            '\n',
            'import (\n',
            '\t"errors"\n',
            '\t"math"\n',
            '\t"arc-homework/space-game/moving/object"\n',
            '\t"arc-homework/space-game/moving/vector"\n',
            ')\n',
            '\n',
            '//go:generate blah-blah-blah\n',
            *interface_body,
        )
        generator = AdapterGenerator(
            root='',
            go_package='foobar',
            dependency_module='FooBarDep',
            content=content,
            line=10,
        )

        actual = generator.generate()

        expected = (
            '// auto-generated file, do not edit\n'
            'package foobar\n'
            '\n'
            'import (\n'
            '\t"context"\n'
            '\n'
            '\t"arc-homework/space-game/ioc"\n'
            '\t"arc-homework/space-game/moving/object"\n'
            '\t"arc-homework/space-game/moving/vector"\n'
            ')\n'
            '\n'
            'type fooBarAdapter struct {\n'
            '\tioc ioc.IoC\n'
            '\tobj object.Object\n'
            '}\n'
            '\n'
            'func (a *fooBarAdapter) FooMethod(arg1 int, arg2 []string, arg3 map[string]Bar) (vector.Vector, error) {\n'
            '\tresolvedDependency, err := a.ioc.Resolve(ctx, "FooBarDep:FooMethod", a.obj, arg1, arg2, arg3)\n'
            '\tif err != nil {\n'
            '\t\treturn vector.Vector{}, err\n'
            '\t}\n'
            '\n'
            '\tresolvedDependencyCast, ok := resolvedDependency.(vector.Vector)\n'
            '\tif !ok {\n'
            '\t\treturn vector.Vector{}, err\n'
            '\t}\n'
            '\n'
            '\treturn resolvedDependencyCast, nil\n'
            '}\n'
            '\n'
            'func (a *fooBarAdapter) BarMethod(v vector.Vector) (vector.Vector, error) {\n'
            '\tresolvedDependency, err := a.ioc.Resolve(ctx, "FooBarDep:BarMethod", a.obj, v)\n'
            '\tif err != nil {\n'
            '\t\treturn vector.Vector{}, err\n'
            '\t}\n'
            '\n'
            '\tresolvedDependencyCast, ok := resolvedDependency.(vector.Vector)\n'
            '\tif !ok {\n'
            '\t\treturn vector.Vector{}, err\n'
            '\t}\n'
            '\n'
            '\treturn resolvedDependencyCast, nil\n'
            '}\n'
            '\n'
            'func NewfooBarAdapter(ioc ioc.IoC, obj object.Object) FooBar {\n'
            '\treturn &fooBarAdapter{\n'
            '\t\tioc: ioc,\n'
            '\t\tobj: obj,\n'
            '\t}\n'
            '}\n'
        )
        assert expected == actual
