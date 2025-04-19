#[typeshare]
pub struct NestedStruct {
    pub nested_field: String,
}

/// Tests various edge cases
#[typeshare]
pub enum EdgeCaseEnum {
    /// Empty variant
    Empty,

    /// Optional string
    OptionalString(Option<String>),

    /// Optional nested struct
    OptionalNested(Option<NestedStruct>),

    /// Array of values
    Array(Vec<i32>),

    /// Array of structs
    StructArray(Vec<NestedStruct>),

    /// Complex nested structure
    Complex {
        a: String,
        b: Option<i32>,
        c: Vec<NestedStruct>,
        d: Option<Vec<String>>,
    }
}