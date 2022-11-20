import abc
import os
from typing import Dict, Iterable

from generator.utils import read_content


class CodeGeneratorBase(abc.ABC):
    _BASE_TEMPLATE_PATH = 'template'

    _TEMPLATE_PATH: str

    def __init__(
        self,
        root: str,
        go_package: str,
        dependency_module: str,
        content: Iterable[str],
        line: int,
    ):
        self._root = root
        self._go_package = go_package
        self._content = content
        self._line = line
        self._dependency_module = dependency_module
        self._templates = self._init_templates()

    def _init_templates(self) -> Dict[str, str]:
        templates = {}
        path = os.path.join(self._root, self._BASE_TEMPLATE_PATH, self._TEMPLATE_PATH)
        for _, _, filenames in os.walk(path):
            for filename in filenames:
                template_name = filename.rsplit('.', maxsplit=1)[0]
                templates[template_name] = read_content(os.path.join(path, filename))

        return templates

    @abc.abstractmethod
    def generate(self) -> str:
        pass
