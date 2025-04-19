//! Test for generic externally tagged enums

/// Helper for nested generic types
#[typeshare]
pub struct Container<T> {
    pub value: T,
}

/// Generic externally tagged enum
#[typeshare]
pub enum GenericExternalEnum<T, U> {
    /// Unit variant
    Unit,
    
    /// Generic tuple variant with first type
    First(T),
    
    /// Generic tuple variant with second type
    Second(U),
    
    /// Mixed generic tuple variant
    Mixed(T, U),
    
    /// Generic struct variant
    Struct {
        t_value: T,
        u_value: U,
    },
    
    /// Container with generic type
    Nested(Container<T>),
}