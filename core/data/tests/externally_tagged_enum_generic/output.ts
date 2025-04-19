/** Helper for nested generic types */
export interface Container<T> {
	value: T;
}

/** Generic externally tagged enum */
export type GenericExternalEnum<T, U> = 
	/** Unit variant */
	| "Unit"
	/** Generic tuple variant with first type */
	| { First: T }
	/** Generic tuple variant with second type */
	| { Second: U }
	/** Generic struct variant */
	| { Struct: { t_value: T; u_value: U } }
	/** Container with generic type */
	| { Nested: Container<T> };

