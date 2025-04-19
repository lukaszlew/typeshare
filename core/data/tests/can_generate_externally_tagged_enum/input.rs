/// Simple helper struct
#[typeshare]
pub struct ItemValue {
    pub field: String,
}

/// Basic externally tagged enum
#[typeshare]
pub enum BasicExternalEnum {
    /// Unit variant
    Unit,

    /// String variant
    String(String),

    /// Number variant
    Number(i32),

    /// Struct variant
    Struct {
        field1: String,
        field2: i32,
    },

    /// Nested variant
    Nested(ItemValue),
}