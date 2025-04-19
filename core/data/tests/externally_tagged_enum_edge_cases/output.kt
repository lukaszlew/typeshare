package com.agilebits.onepassword.colorsModule

import kotlinx.serialization.*
import kotlinx.serialization.json.*
import kotlinx.serialization.descriptors.*
import kotlinx.serialization.encoding.*

@Serializable
data class NestedStruct(
	val nested_field: String
)

/** Tests various edge cases */
@Serializable(EdgeCaseEnumSerializer::class)
sealed class EdgeCaseEnum {
	/** Empty variant */
	@Serializable
	object Empty : EdgeCaseEnum()
	
	/** Optional string */
	@Serializable
	data class OptionalString(val value: String?) : EdgeCaseEnum()
	
	/** Optional nested struct */
	@Serializable
	data class OptionalNested(val value: NestedStruct?) : EdgeCaseEnum()
	
	/** Array of values */
	@Serializable
	data class Array(val value: List<Int>) : EdgeCaseEnum()
	
	/** Array of structs */
	@Serializable
	data class StructArray(val value: List<NestedStruct>) : EdgeCaseEnum()
	
	/** Complex nested structure */
	@Serializable
	data class Complex(
		val a: String,
		val b: Int?,
		val c: List<NestedStruct>,
		val d: List<String>?
	) : EdgeCaseEnum()
}

class EdgeCaseEnumSerializer : JsonContentPolymorphicSerializer<EdgeCaseEnum>(EdgeCaseEnum::class) {
	override fun selectDeserializer(element: JsonElement): DeserializationStrategy<out EdgeCaseEnum> {
		return when {
			element is JsonPrimitive && element.isString && element.content == "Empty" -> EdgeCaseEnum.Empty.serializer()
			element is JsonObject && "OptionalString" in element -> EdgeCaseEnum.OptionalString.serializer()
			element is JsonObject && "OptionalNested" in element -> EdgeCaseEnum.OptionalNested.serializer()
			element is JsonObject && "Array" in element -> EdgeCaseEnum.Array.serializer()
			element is JsonObject && "StructArray" in element -> EdgeCaseEnum.StructArray.serializer()
			element is JsonObject && "Complex" in element -> EdgeCaseEnum.Complex.serializer()
			else -> throw SerializationException("Unknown element: $element")
		}
	}
}