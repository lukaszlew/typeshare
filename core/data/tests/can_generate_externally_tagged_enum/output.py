from __future__ import annotations

from enum import Enum
from typing import Literal, Union, Dict, Any, Optional
from pydantic import BaseModel, Field


# Simple helper struct
class ItemValue(BaseModel):
    field: str


# Struct variant inner type
class StructContent(BaseModel):
    field1: str
    field2: int


# Basic externally tagged enum
class BasicExternalEnum(BaseModel):
    __root__: Union[
        Literal["Unit"],  # Unit variant
        Dict[Literal["String"], str],  # String variant
        Dict[Literal["Number"], int],  # Number variant
        Dict[Literal["Struct"], StructContent],  # Struct variant
        Dict[Literal["Nested"], ItemValue],  # Nested variant
    ]

    @classmethod
    def unit(cls) -> BasicExternalEnum:
        """Unit variant"""
        return cls(__root__="Unit")

    @classmethod
    def string(cls, value: str) -> BasicExternalEnum:
        """String variant"""
        return cls(__root__={"String": value})

    @classmethod
    def number(cls, value: int) -> BasicExternalEnum:
        """Number variant"""
        return cls(__root__={"Number": value})

    @classmethod
    def struct(cls, field1: str, field2: int) -> BasicExternalEnum:
        """Struct variant"""
        return cls(__root__={"Struct": StructContent(field1=field1, field2=field2)})

    @classmethod
    def nested(cls, value: ItemValue) -> BasicExternalEnum:
        """Nested variant"""
        return cls(__root__={"Nested": value})

    def is_unit(self) -> bool:
        return self.__root__ == "Unit"

    def is_string(self) -> bool:
        return isinstance(self.__root__, dict) and "String" in self.__root__

    def is_number(self) -> bool:
        return isinstance(self.__root__, dict) and "Number" in self.__root__

    def is_struct(self) -> bool:
        return isinstance(self.__root__, dict) and "Struct" in self.__root__

    def is_nested(self) -> bool:
        return isinstance(self.__root__, dict) and "Nested" in self.__root__

    def get_string(self) -> Optional[str]:
        if self.is_string():
            return self.__root__["String"]
        return None

    def get_number(self) -> Optional[int]:
        if self.is_number():
            return self.__root__["Number"]
        return None

    def get_struct(self) -> Optional[StructContent]:
        if self.is_struct():
            return self.__root__["Struct"]
        return None

    def get_nested(self) -> Optional[ItemValue]:
        if self.is_nested():
            return self.__root__["Nested"]
        return None