export interface NestedStruct {
	nested_field: string;
}

/** Tests various edge cases */
export type EdgeCaseEnum = 
	/** Empty variant */
	| "Empty"
	/** Optional string */
	| { OptionalString?: string }
	/** Optional nested struct */
	| { OptionalNested?: NestedStruct }
	/** Array of values */
	| { Array: number[] }
	/** Array of structs */
	| { StructArray: NestedStruct[] }
	/** Complex nested structure */
	| { Complex: { a: string; b?: number; c: NestedStruct[]; d?: string[] } };

