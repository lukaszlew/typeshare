package proto

import (
	"encoding/json"
	"fmt"
)

// CamelCaseEnumType - Enum of possible types for CamelCaseEnum
type CamelCaseEnumType string

const (
	// CamelCaseEnumTypeUnitVariant - Unit variant
	CamelCaseEnumTypeUnitVariant CamelCaseEnumType = "unitVariant"
	// CamelCaseEnumTypeSimpleValue - SimpleValue variant
	CamelCaseEnumTypeSimpleValue CamelCaseEnumType = "simpleValue"
	// CamelCaseEnumTypeComplexValue - ComplexValue variant
	CamelCaseEnumTypeComplexValue CamelCaseEnumType = "complexValue"
)

// ComplexValue - Fields for complex value
type ComplexValue struct {
	Field int32 `json:"field"`
}

// CamelCaseEnum - Tests renaming variants
type CamelCaseEnum struct {
	Type        CamelCaseEnumType
	SimpleValue *string      `json:"simpleValue,omitempty"`
	ComplexValue *ComplexValue `json:"complexValue,omitempty"`
}

// MarshalJSON implements json.Marshaler
func (e CamelCaseEnum) MarshalJSON() ([]byte, error) {
	switch e.Type {
	case CamelCaseEnumTypeUnitVariant:
		return json.Marshal("unitVariant")
	case CamelCaseEnumTypeSimpleValue:
		if e.SimpleValue == nil {
			return nil, fmt.Errorf("CamelCaseEnum.SimpleValue cannot be nil when Type is SimpleValue")
		}
		return json.Marshal(map[string]string{"simpleValue": *e.SimpleValue})
	case CamelCaseEnumTypeComplexValue:
		if e.ComplexValue == nil {
			return nil, fmt.Errorf("CamelCaseEnum.ComplexValue cannot be nil when Type is ComplexValue")
		}
		return json.Marshal(map[string]interface{}{"complexValue": e.ComplexValue})
	default:
		return nil, fmt.Errorf("unknown CamelCaseEnum type: %s", e.Type)
	}
}

// UnmarshalJSON implements json.Unmarshaler
func (e *CamelCaseEnum) UnmarshalJSON(data []byte) error {
	// Try to unmarshal as string (for UnitVariant)
	var strValue string
	if err := json.Unmarshal(data, &strValue); err == nil {
		if strValue == "unitVariant" {
			e.Type = CamelCaseEnumTypeUnitVariant
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
		case "simpleValue":
			var stringValue string
			if err := json.Unmarshal(value, &stringValue); err != nil {
				return err
			}
			e.Type = CamelCaseEnumTypeSimpleValue
			e.SimpleValue = &stringValue
		case "complexValue":
			var complexValue ComplexValue
			if err := json.Unmarshal(value, &complexValue); err != nil {
				return err
			}
			e.Type = CamelCaseEnumTypeComplexValue
			e.ComplexValue = &complexValue
		default:
			return fmt.Errorf("unknown variant %q in CamelCaseEnum", key)
		}
	}
	return nil
}

// SpecificRenameEnumType - Enum of possible types for SpecificRenameEnum
type SpecificRenameEnumType string

const (
	// SpecificRenameEnumTypeRegular - Regular variant
	SpecificRenameEnumTypeRegular SpecificRenameEnumType = "Regular"
	// SpecificRenameEnumTypeCustomName - CustomName variant
	SpecificRenameEnumTypeCustomName SpecificRenameEnumType = "custom_name"
	// SpecificRenameEnumTypeCustomStruct - CustomStruct variant
	SpecificRenameEnumTypeCustomStruct SpecificRenameEnumType = "custom_struct"
)

// CustomStruct - Fields for custom struct
type CustomStruct struct {
	Value int32 `json:"value"`
}

// SpecificRenameEnum - Tests specific renames on variants
type SpecificRenameEnum struct {
	Type        SpecificRenameEnumType
	CustomName  *string      `json:"custom_name,omitempty"`
	CustomStruct *CustomStruct `json:"custom_struct,omitempty"`
}

// MarshalJSON implements json.Marshaler
func (e SpecificRenameEnum) MarshalJSON() ([]byte, error) {
	switch e.Type {
	case SpecificRenameEnumTypeRegular:
		return json.Marshal("Regular")
	case SpecificRenameEnumTypeCustomName:
		if e.CustomName == nil {
			return nil, fmt.Errorf("SpecificRenameEnum.CustomName cannot be nil when Type is CustomName")
		}
		return json.Marshal(map[string]string{"custom_name": *e.CustomName})
	case SpecificRenameEnumTypeCustomStruct:
		if e.CustomStruct == nil {
			return nil, fmt.Errorf("SpecificRenameEnum.CustomStruct cannot be nil when Type is CustomStruct")
		}
		return json.Marshal(map[string]interface{}{"custom_struct": e.CustomStruct})
	default:
		return nil, fmt.Errorf("unknown SpecificRenameEnum type: %s", e.Type)
	}
}

// UnmarshalJSON implements json.Unmarshaler
func (e *SpecificRenameEnum) UnmarshalJSON(data []byte) error {
	// Try to unmarshal as string (for Regular)
	var strValue string
	if err := json.Unmarshal(data, &strValue); err == nil {
		if strValue == "Regular" {
			e.Type = SpecificRenameEnumTypeRegular
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
		case "custom_name":
			var stringValue string
			if err := json.Unmarshal(value, &stringValue); err != nil {
				return err
			}
			e.Type = SpecificRenameEnumTypeCustomName
			e.CustomName = &stringValue
		case "custom_struct":
			var structValue CustomStruct
			if err := json.Unmarshal(value, &structValue); err != nil {
				return err
			}
			e.Type = SpecificRenameEnumTypeCustomStruct
			e.CustomStruct = &structValue
		default:
			return fmt.Errorf("unknown variant %q in SpecificRenameEnum", key)
		}
	}
	return nil
}

// MixedRenameEnumType - Enum of possible types for MixedRenameEnum
type MixedRenameEnumType string

const (
	// MixedRenameEnumTypeUnitValue - Unit value variant
	MixedRenameEnumTypeUnitValue MixedRenameEnumType = "unit-value"
	// MixedRenameEnumTypeCustomValue - Custom value variant
	MixedRenameEnumTypeCustomValue MixedRenameEnumType = "CUSTOM"
	// MixedRenameEnumTypeComplexValue - Complex value variant
	MixedRenameEnumTypeComplexValue MixedRenameEnumType = "complex-value"
)

// MixedComplexValue - Fields for mixed complex value
type MixedComplexValue struct {
	Field int32 `json:"field"`
}

// MixedRenameEnum - Mixed kebab and specific renames
type MixedRenameEnum struct {
	Type         MixedRenameEnumType
	CustomValue  *string          `json:"CUSTOM,omitempty"`
	ComplexValue *MixedComplexValue `json:"complex-value,omitempty"`
}

// MarshalJSON implements json.Marshaler
func (e MixedRenameEnum) MarshalJSON() ([]byte, error) {
	switch e.Type {
	case MixedRenameEnumTypeUnitValue:
		return json.Marshal("unit-value")
	case MixedRenameEnumTypeCustomValue:
		if e.CustomValue == nil {
			return nil, fmt.Errorf("MixedRenameEnum.CustomValue cannot be nil when Type is CustomValue")
		}
		return json.Marshal(map[string]string{"CUSTOM": *e.CustomValue})
	case MixedRenameEnumTypeComplexValue:
		if e.ComplexValue == nil {
			return nil, fmt.Errorf("MixedRenameEnum.ComplexValue cannot be nil when Type is ComplexValue")
		}
		return json.Marshal(map[string]interface{}{"complex-value": e.ComplexValue})
	default:
		return nil, fmt.Errorf("unknown MixedRenameEnum type: %s", e.Type)
	}
}

// UnmarshalJSON implements json.Unmarshaler
func (e *MixedRenameEnum) UnmarshalJSON(data []byte) error {
	// Try to unmarshal as string (for UnitValue)
	var strValue string
	if err := json.Unmarshal(data, &strValue); err == nil {
		if strValue == "unit-value" {
			e.Type = MixedRenameEnumTypeUnitValue
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
		case "CUSTOM":
			var stringValue string
			if err := json.Unmarshal(value, &stringValue); err != nil {
				return err
			}
			e.Type = MixedRenameEnumTypeCustomValue
			e.CustomValue = &stringValue
		case "complex-value":
			var complexValue MixedComplexValue
			if err := json.Unmarshal(value, &complexValue); err != nil {
				return err
			}
			e.Type = MixedRenameEnumTypeComplexValue
			e.ComplexValue = &complexValue
		default:
			return fmt.Errorf("unknown variant %q in MixedRenameEnum", key)
		}
	}
	return nil
}