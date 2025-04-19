package proto

import (
	"encoding/json"
	"fmt"
)

// ItemValue - Simple helper struct
type ItemValue struct {
	Field string `json:"field"`
}

// BasicExternalEnumType - Enum of possible types for BasicExternalEnum
type BasicExternalEnumType string

const (
	// BasicExternalEnumTypeUnit - Unit variant
	BasicExternalEnumTypeUnit BasicExternalEnumType = "Unit"
	// BasicExternalEnumTypeString - String variant
	BasicExternalEnumTypeString BasicExternalEnumType = "String"
	// BasicExternalEnumTypeNumber - Number variant
	BasicExternalEnumTypeNumber BasicExternalEnumType = "Number"
	// BasicExternalEnumTypeStruct - Struct variant
	BasicExternalEnumTypeStruct BasicExternalEnumType = "Struct"
	// BasicExternalEnumTypeNested - Nested variant
	BasicExternalEnumTypeNested BasicExternalEnumType = "Nested"
)

// BasicExternalEnum - Basic externally tagged enum
type BasicExternalEnum struct {
	Type   BasicExternalEnumType
	String *string     `json:"String,omitempty"`
	Number *int32      `json:"Number,omitempty"`
	Struct *struct {
		Field1 string `json:"field1"`
		Field2 int32  `json:"field2"`
	} `json:"Struct,omitempty"`
	Nested *ItemValue `json:"Nested,omitempty"`
}

// MarshalJSON implements json.Marshaler
func (e BasicExternalEnum) MarshalJSON() ([]byte, error) {
	switch e.Type {
	case BasicExternalEnumTypeUnit:
		return json.Marshal("Unit")
	case BasicExternalEnumTypeString:
		if e.String == nil {
			return nil, fmt.Errorf("BasicExternalEnum.String cannot be nil when Type is String")
		}
		return json.Marshal(map[string]string{"String": *e.String})
	case BasicExternalEnumTypeNumber:
		if e.Number == nil {
			return nil, fmt.Errorf("BasicExternalEnum.Number cannot be nil when Type is Number")
		}
		return json.Marshal(map[string]int32{"Number": *e.Number})
	case BasicExternalEnumTypeStruct:
		if e.Struct == nil {
			return nil, fmt.Errorf("BasicExternalEnum.Struct cannot be nil when Type is Struct")
		}
		return json.Marshal(map[string]interface{}{"Struct": e.Struct})
	case BasicExternalEnumTypeNested:
		if e.Nested == nil {
			return nil, fmt.Errorf("BasicExternalEnum.Nested cannot be nil when Type is Nested")
		}
		return json.Marshal(map[string]ItemValue{"Nested": *e.Nested})
	default:
		return nil, fmt.Errorf("unknown BasicExternalEnum type: %s", e.Type)
	}
}

// UnmarshalJSON implements json.Unmarshaler
func (e *BasicExternalEnum) UnmarshalJSON(data []byte) error {
	// Try to unmarshal as string (for Unit variant)
	var strValue string
	if err := json.Unmarshal(data, &strValue); err == nil {
		if strValue == "Unit" {
			e.Type = BasicExternalEnumTypeUnit
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
		case "String":
			var stringValue string
			if err := json.Unmarshal(value, &stringValue); err != nil {
				return err
			}
			e.Type = BasicExternalEnumTypeString
			e.String = &stringValue
		case "Number":
			var numberValue int32
			if err := json.Unmarshal(value, &numberValue); err != nil {
				return err
			}
			e.Type = BasicExternalEnumTypeNumber
			e.Number = &numberValue
		case "Struct":
			var structValue struct {
				Field1 string `json:"field1"`
				Field2 int32  `json:"field2"`
			}
			if err := json.Unmarshal(value, &structValue); err != nil {
				return err
			}
			e.Type = BasicExternalEnumTypeStruct
			e.Struct = &structValue
		case "Nested":
			var nestedValue ItemValue
			if err := json.Unmarshal(value, &nestedValue); err != nil {
				return err
			}
			e.Type = BasicExternalEnumTypeNested
			e.Nested = &nestedValue
		default:
			return fmt.Errorf("unknown variant %q in BasicExternalEnum", key)
		}
	}
	return nil
}