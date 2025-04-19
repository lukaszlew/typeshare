package proto

import (
	"encoding/json"
	"fmt"
)

type NestedStruct struct {
	NestedField string `json:"nested_field"`
}

// EdgeCaseEnumType - Enum of possible types for EdgeCaseEnum
type EdgeCaseEnumType string

const (
	// EdgeCaseEnumTypeEmpty - Empty variant
	EdgeCaseEnumTypeEmpty EdgeCaseEnumType = "Empty"
	// EdgeCaseEnumTypeOptionalString - Optional string
	EdgeCaseEnumTypeOptionalString EdgeCaseEnumType = "OptionalString"
	// EdgeCaseEnumTypeOptionalNested - Optional nested struct
	EdgeCaseEnumTypeOptionalNested EdgeCaseEnumType = "OptionalNested"
	// EdgeCaseEnumTypeArray - Array of values
	EdgeCaseEnumTypeArray EdgeCaseEnumType = "Array"
	// EdgeCaseEnumTypeStructArray - Array of structs
	EdgeCaseEnumTypeStructArray EdgeCaseEnumType = "StructArray"
	// EdgeCaseEnumTypeComplex - Complex nested structure
	EdgeCaseEnumTypeComplex EdgeCaseEnumType = "Complex"
)

// ComplexFields - Fields for Complex variant
type ComplexFields struct {
	A string         `json:"a"`
	B *int32         `json:"b,omitempty"`
	C []NestedStruct `json:"c"`
	D []string       `json:"d,omitempty"`
}

// EdgeCaseEnum - Tests various edge cases
type EdgeCaseEnum struct {
	Type          EdgeCaseEnumType
	OptionalString *string        `json:"OptionalString,omitempty"`
	OptionalNested *NestedStruct  `json:"OptionalNested,omitempty"`
	Array         []int32         `json:"Array,omitempty"`
	StructArray   []NestedStruct  `json:"StructArray,omitempty"`
	Complex       *ComplexFields  `json:"Complex,omitempty"`
}

// MarshalJSON implements json.Marshaler
func (e EdgeCaseEnum) MarshalJSON() ([]byte, error) {
	switch e.Type {
	case EdgeCaseEnumTypeEmpty:
		return json.Marshal("Empty")
	case EdgeCaseEnumTypeOptionalString:
		return json.Marshal(map[string]interface{}{"OptionalString": e.OptionalString})
	case EdgeCaseEnumTypeOptionalNested:
		return json.Marshal(map[string]interface{}{"OptionalNested": e.OptionalNested})
	case EdgeCaseEnumTypeArray:
		if e.Array == nil {
			return nil, fmt.Errorf("EdgeCaseEnum.Array cannot be nil when Type is Array")
		}
		return json.Marshal(map[string][]int32{"Array": e.Array})
	case EdgeCaseEnumTypeStructArray:
		if e.StructArray == nil {
			return nil, fmt.Errorf("EdgeCaseEnum.StructArray cannot be nil when Type is StructArray")
		}
		return json.Marshal(map[string][]NestedStruct{"StructArray": e.StructArray})
	case EdgeCaseEnumTypeComplex:
		if e.Complex == nil {
			return nil, fmt.Errorf("EdgeCaseEnum.Complex cannot be nil when Type is Complex")
		}
		return json.Marshal(map[string]ComplexFields{"Complex": *e.Complex})
	default:
		return nil, fmt.Errorf("unknown EdgeCaseEnum type: %s", e.Type)
	}
}

// UnmarshalJSON implements json.Unmarshaler
func (e *EdgeCaseEnum) UnmarshalJSON(data []byte) error {
	// Try to unmarshal as string (for Empty variant)
	var strValue string
	if err := json.Unmarshal(data, &strValue); err == nil {
		if strValue == "Empty" {
			e.Type = EdgeCaseEnumTypeEmpty
			return nil
		}
	}

	// Try to unmarshal as an object with a single key
	var objMap map[string]json.RawMessage
	if err := json.Unmarshal(data, &objMap); err != nil {
		return err
	}

	if len(objMap) != 1 {
		return fmt.Errorf("expected exactly one field in externally tagged enum")
	}

	for key, value := range objMap {
		switch key {
		case "OptionalString":
			var stringValue *string
			if err := json.Unmarshal(value, &stringValue); err != nil {
				return err
			}
			e.Type = EdgeCaseEnumTypeOptionalString
			e.OptionalString = stringValue
		case "OptionalNested":
			var nestedValue *NestedStruct
			if err := json.Unmarshal(value, &nestedValue); err != nil {
				return err
			}
			e.Type = EdgeCaseEnumTypeOptionalNested
			e.OptionalNested = nestedValue
		case "Array":
			var arrayValue []int32
			if err := json.Unmarshal(value, &arrayValue); err != nil {
				return err
			}
			e.Type = EdgeCaseEnumTypeArray
			e.Array = arrayValue
		case "StructArray":
			var structArrayValue []NestedStruct
			if err := json.Unmarshal(value, &structArrayValue); err != nil {
				return err
			}
			e.Type = EdgeCaseEnumTypeStructArray
			e.StructArray = structArrayValue
		case "Complex":
			var complexValue ComplexFields
			if err := json.Unmarshal(value, &complexValue); err != nil {
				return err
			}
			e.Type = EdgeCaseEnumTypeComplex
			e.Complex = &complexValue
		default:
			return fmt.Errorf("unknown variant %q in EdgeCaseEnum", key)
		}
	}
	return nil
}