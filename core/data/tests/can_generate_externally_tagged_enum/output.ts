/** Simple helper struct */
export interface ItemValue {
	field: string;
}

/** Basic externally tagged enum */
export type BasicExternalEnum =
	/** Unit variant */
	| "Unit"
	/** String variant */
	| { String: string }
	/** Number variant */
	| { Number: number }
	/** Struct variant */
	| { Struct: { field1: string; field2: number; } }
	/** Nested variant */
	| { Nested: ItemValue };