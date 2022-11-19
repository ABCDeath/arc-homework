from typing import Type

from generator.adapter import AdapterGenerator
from generator.generator import CodeGeneratorBase

_generator_by_type = {
    'adapter': AdapterGenerator,
}


def generator_factory(object_type: str) -> Type[CodeGeneratorBase]:
    return _generator_by_type[object_type]
