from generator.adapter import AdapterGenerator
from generator.factory import generator_factory


class TestFactory:
    def test_returns_generator_class_by_object_type(self):
        generator = generator_factory('adapter')

        assert generator is AdapterGenerator
