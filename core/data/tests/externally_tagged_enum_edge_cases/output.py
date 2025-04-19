from __future__ import annotations

from enum import Enum
from typing import Union, Dict, List, Optional, Any, Literal
from pydantic import BaseModel, Field


# Nested struct
class NestedStruct(BaseModel):
    nested_field: str


# Complex fields structure
class ComplexContent(BaseModel):
    a: str
    b: Optional[int] = None
    c: List[NestedStruct]
    d: Optional[List[str]] = None


# Tests various edge cases
class EdgeCaseEnum(BaseModel):
    __root__: Union[
        Literal["Empty"],  # Empty variant
        Dict[Literal["OptionalString"], Optional[str]],  # Optional string
        Dict[Literal["OptionalNested"], Optional[NestedStruct]],  # Optional nested struct
        Dict[Literal["Array"], List[int]],  # Array of values
        Dict[Literal["StructArray"], List[NestedStruct]],  # Array of structs
        Dict[Literal["Complex"], ComplexContent],  # Complex nested structure
    ]

    @classmethod
    def empty(cls) -> EdgeCaseEnum:
        """Empty variant"""
        return cls(__root__="Empty")

    @classmethod
    def optional_string(cls, value: Optional[str]) -> EdgeCaseEnum:
        """Optional string"""
        return cls(__root__={"OptionalString": value})

    @classmethod
    def optional_nested(cls, value: Optional[NestedStruct]) -> EdgeCaseEnum:
        """Optional nested struct"""
        return cls(__root__={"OptionalNested": value})

    @classmethod
    def array(cls, value: List[int]) -> EdgeCaseEnum:
        """Array of values"""
        return cls(__root__={"Array": value})

    @classmethod
    def struct_array(cls, value: List[NestedStruct]) -> EdgeCaseEnum:
        """Array of structs"""
        return cls(__root__={"StructArray": value})

    @classmethod
    def complex(
        cls, a: str, b: Optional[int], c: List[NestedStruct], d: Optional[List[str]]
    ) -> EdgeCaseEnum:
        """Complex nested structure"""
        return cls(__root__={"Complex": ComplexContent(a=a, b=b, c=c, d=d)})

    def is_empty(self) -> bool:
        return self.__root__ == "Empty"

    def is_optional_string(self) -> bool:
        return isinstance(self.__root__, dict) and "OptionalString" in self.__root__

    def is_optional_nested(self) -> bool:
        return isinstance(self.__root__, dict) and "OptionalNested" in self.__root__

    def is_array(self) -> bool:
        return isinstance(self.__root__, dict) and "Array" in self.__root__

    def is_struct_array(self) -> bool:
        return isinstance(self.__root__, dict) and "StructArray" in self.__root__

    def is_complex(self) -> bool:
        return isinstance(self.__root__, dict) and "Complex" in self.__root__

    def get_optional_string(self) -> Optional[str]:
        if self.is_optional_string():
            return self.__root__["OptionalString"]
        return None

    def get_optional_nested(self) -> Optional[NestedStruct]:
        if self.is_optional_nested():
            return self.__root__["OptionalNested"]
        return None

    def get_array(self) -> Optional[List[int]]:
        if self.is_array():
            return self.__root__["Array"]
        return None

    def get_struct_array(self) -> Optional[List[NestedStruct]]:
        if self.is_struct_array():
            return self.__root__["StructArray"]
        return None

    def get_complex(self) -> Optional[ComplexContent]:
        if self.is_complex():
            return self.__root__["Complex"]
        return None