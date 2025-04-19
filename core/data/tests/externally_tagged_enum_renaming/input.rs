/// Tests renaming variants
#[typeshare]
#[serde(rename_all = "camelCase")]
pub enum CamelCaseEnum {
    UnitVariant,
    SimpleValue(String),
    ComplexValue { field: i32 },
}

/// Tests specific renames on variants
#[typeshare]
pub enum SpecificRenameEnum {
    Regular,
    #[serde(rename = "custom_name")]
    CustomName(String),
    #[serde(rename = "custom_struct")]
    CustomStruct { value: i32 },
}

/// Mixed kebab and specific renames
#[typeshare]
#[serde(rename_all = "kebab-case")]
pub enum MixedRenameEnum {
    UnitValue,
    #[serde(rename = "CUSTOM")]
    CustomValue(String),
    ComplexValue { field: i32 },
}