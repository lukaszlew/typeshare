from __future__ import annotations

from pydantic import BaseModel
from typing import Generic, TypeVar

T = TypeVar("T")


class Container(BaseModel, Generic[T]):
    """
    Helper for nested generic types
    """
    value: T

