export interface NestedStruct {
	nested_field: string;
}

/** Tests various edge cases */
export type EdgeCaseEnum =
	/** Empty variant */
	| "Empty"
	/** Optional string */
	| { OptionalString: string | null }
	/** Optional nested struct */
	| { OptionalNested: NestedStruct | null }
	/** Array of values */
	| { Array: number[] }
	/** Array of structs */
	| { StructArray: NestedStruct[] }
	/** Complex nested structure */
	| { Complex: { a: string; b: number | null; c: NestedStruct[]; d: string[] | null; } };