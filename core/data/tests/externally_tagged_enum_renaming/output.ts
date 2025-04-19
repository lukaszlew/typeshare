/** Tests renaming variants */
export type CamelCaseEnum =
	| "unitVariant"
	| { simpleValue: string }
	| { complexValue: { field: number } };

/** Mixed kebab and specific renames */
export type MixedRenameEnum =
	| "unit-value"
	| { CUSTOM: string }
	| { complex-value: { field: number } };

/** Tests specific renames on variants */
export type SpecificRenameEnum =
	| "Regular"
	| { custom_name: string }
	| { custom_struct: { value: number } };
