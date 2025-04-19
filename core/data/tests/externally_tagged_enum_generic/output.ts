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
	/** Mixed generic tuple variant */
	| { Mixed: [T, U] }
	/** Generic struct variant */
	| { Struct: { t_value: T; u_value: U; } }
	/** Container with generic type */
	| { Nested: Container<T> };

export const ReviverFunc = (key: string, value: unknown): unknown => {
    // Handle externally tagged enum
if (typeof value === "string" || (typeof value === "object" && value !== null && Object.keys(value).length === 1)) {
    return value;
}
    return value;
};

export const ReplacerFunc = (key: string, value: unknown): unknown => {
    // Handle externally tagged enum
if (typeof value === "string" || (typeof value === "object" && value !== null && Object.keys(value).length === 1)) {
    return value;
}
    return value;
};