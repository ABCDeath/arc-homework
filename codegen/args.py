import os
import sys
from typing import Dict, NamedTuple, Tuple


class CodeGenParams(NamedTuple):
    object_type: str
    dependency_module: str
    go_package: str
    file: str
    line: int


def get_codegen_params() -> CodeGenParams:
    args = _get_args()

    return CodeGenParams(
        object_type=args['object-type'],
        dependency_module=args['dependency-module'],
        go_package=os.environ.get('GOPACKAGE'),
        file=os.path.join(os.environ.get('PWD'), os.environ.get('GOFILE')),
        line=int(os.environ.get('GOLINE'))
    )


def _get_args() -> Dict[str, str]:
    return dict(_get_arg_key_value(a) for a in sys.argv[1:])


def _get_arg_key_value(arg: str) -> Tuple[str, str]:
    kv = arg.replace('--', '').split('=', maxsplit=1)
    if len(kv) > 1:
        return kv[0], kv[1]

    return kv[0], ''
