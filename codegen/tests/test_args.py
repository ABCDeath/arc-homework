from unittest import mock

from args import CodeGenParams, _get_arg_key_value, _get_args, get_codegen_params


class TestArgs:
    def test_get_arg_key_value_returns_tuple_of_key_and_value_for_value_arg(self):
        assert _get_arg_key_value('--arg=val') == ('arg', 'val')

    def test_get_arg_key_value_returns_tuple_of_key_and_empty_string_for_flag_arg(self):
        assert _get_arg_key_value('--arg=val') == ('arg', 'val')

    def test_get_args_returns_dict(self):
        with mock.patch(
            'sys.argv',
            ('script_path_string', '--arg1=val1', '--arg2=val2', '--arg_flag'),
        ):
            args = _get_args()

            expected = {
                'arg1': 'val1',
                'arg2': 'val2',
                'arg_flag': '',
            }
            assert expected == args

    def test_get_codegen_params(self):
        with mock.patch(
            'os.environ',
            {
                'PWD': '/some/path',
                'GOFILE': 'filename',
                'GOPACKAGE': 'package_name',
                'GOLINE': '42',
            }
        ):
            with mock.patch(
                'sys.argv',
                (
                    'script_path_string',
                    '--object-type=object_type',
                    '--dependency-module=module_name',
                ),
            ):
                params = get_codegen_params()

                expected = CodeGenParams(
                    object_type='object_type',
                    dependency_module='module_name',
                    go_package='package_name',
                    file='/some/path/filename',
                    line=42,
                )
                assert expected == params
