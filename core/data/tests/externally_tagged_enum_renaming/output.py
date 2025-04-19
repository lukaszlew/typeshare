from __future__ import annotations

from enum import Enum
from typing import Union, Dict, Optional, Literal
from pydantic import BaseModel, Field


# Complex value struct for CamelCaseEnum
class ComplexValueContent(BaseModel):
    field: int


# Custom struct for SpecificRenameEnum
class CustomStructContent(BaseModel):
    value: int


# Complex value struct for MixedRenameEnum
class MixedComplexValueContent(BaseModel):
    field: int


# Tests renaming variants
class CamelCaseEnum(BaseModel):
    __root__: Union[
        Literal["unitVariant"],
        Dict[Literal["simpleValue"], str],
        Dict[Literal["complexValue"], ComplexValueContent],
    ]

    @classmethod
    def unit_variant(cls) -> CamelCaseEnum:
        """Unit variant"""
        return cls(__root__="unitVariant")

    @classmethod
    def simple_value(cls, value: str) -> CamelCaseEnum:
        """Simple value variant"""
        return cls(__root__={"simpleValue": value})

    @classmethod
    def complex_value(cls, field: int) -> CamelCaseEnum:
        """Complex value variant"""
        return cls(__root__={"complexValue": ComplexValueContent(field=field)})

    def is_unit_variant(self) -> bool:
        return self.__root__ == "unitVariant"

    def is_simple_value(self) -> bool:
        return isinstance(self.__root__, dict) and "simpleValue" in self.__root__

    def is_complex_value(self) -> bool:
        return isinstance(self.__root__, dict) and "complexValue" in self.__root__

    def get_simple_value(self) -> Optional[str]:
        if self.is_simple_value():
            return self.__root__["simpleValue"]
        return None

    def get_complex_value(self) -> Optional[ComplexValueContent]:
        if self.is_complex_value():
            return self.__root__["complexValue"]
        return None


# Tests specific renames on variants
class SpecificRenameEnum(BaseModel):
    __root__: Union[
        Literal["Regular"],
        Dict[Literal["custom_name"], str],
        Dict[Literal["custom_struct"], CustomStructContent],
    ]

    @classmethod
    def regular(cls) -> SpecificRenameEnum:
        """Regular variant"""
        return cls(__root__="Regular")

    @classmethod
    def custom_name(cls, value: str) -> SpecificRenameEnum:
        """Custom name variant"""
        return cls(__root__={"custom_name": value})

    @classmethod
    def custom_struct(cls, value: int) -> SpecificRenameEnum:
        """Custom struct variant"""
        return cls(__root__={"custom_struct": CustomStructContent(value=value)})

    def is_regular(self) -> bool:
        return self.__root__ == "Regular"

    def is_custom_name(self) -> bool:
        return isinstance(self.__root__, dict) and "custom_name" in self.__root__

    def is_custom_struct(self) -> bool:
        return isinstance(self.__root__, dict) and "custom_struct" in self.__root__

    def get_custom_name(self) -> Optional[str]:
        if self.is_custom_name():
            return self.__root__["custom_name"]
        return None

    def get_custom_struct(self) -> Optional[CustomStructContent]:
        if self.is_custom_struct():
            return self.__root__["custom_struct"]
        return None


# Mixed kebab and specific renames
class MixedRenameEnum(BaseModel):
    __root__: Union[
        Literal["unit-value"],
        Dict[Literal["CUSTOM"], str],
        Dict[Literal["complex-value"], MixedComplexValueContent],
    ]

    @classmethod
    def unit_value(cls) -> MixedRenameEnum:
        """Unit value variant"""
        return cls(__root__="unit-value")

    @classmethod
    def custom_value(cls, value: str) -> MixedRenameEnum:
        """Custom value variant"""
        return cls(__root__={"CUSTOM": value})

    @classmethod
    def complex_value(cls, field: int) -> MixedRenameEnum:
        """Complex value variant"""
        return cls(__root__={"complex-value": MixedComplexValueContent(field=field)})

    def is_unit_value(self) -> bool:
        return self.__root__ == "unit-value"

    def is_custom_value(self) -> bool:
        return isinstance(self.__root__, dict) and "CUSTOM" in self.__root__

    def is_complex_value(self) -> bool:
        return isinstance(self.__root__, dict) and "complex-value" in self.__root__

    def get_custom_value(self) -> Optional[str]:
        if self.is_custom_value():
            return self.__root__["CUSTOM"]
        return None

    def get_complex_value(self) -> Optional[MixedComplexValueContent]:
        if self.is_complex_value():
            return self.__root__["complex-value"]
        return None