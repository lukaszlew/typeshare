package com.agilebits.onepassword.colorsModule

import kotlinx.serialization.*
import kotlinx.serialization.json.*
import kotlinx.serialization.descriptors.*
import kotlinx.serialization.encoding.*

/** Simple helper struct */
@Serializable
data class ItemValue(
	val field: String
)

/** Basic externally tagged enum */
@Serializable(BasicExternalEnumSerializer::class)
sealed class BasicExternalEnum {
	/** Unit variant */
	@Serializable
	object Unit : BasicExternalEnum()

	/** String variant */
	@Serializable
	data class String(val value: kotlin.String) : BasicExternalEnum()

	/** Number variant */
	@Serializable
	data class Number(val value: Int) : BasicExternalEnum()

	/** Struct variant */
	@Serializable
	data class Struct(
		val field1: kotlin.String,
		val field2: Int
	) : BasicExternalEnum()

	/** Nested variant */
	@Serializable
	data class Nested(val value: ItemValue) : BasicExternalEnum()
}

class BasicExternalEnumSerializer : JsonContentPolymorphicSerializer<BasicExternalEnum>(BasicExternalEnum::class) {
	override fun selectDeserializer(element: JsonElement): DeserializationStrategy<out BasicExternalEnum> {
		return when {
			element is JsonPrimitive && element.isString && element.content == "Unit" -> BasicExternalEnum.Unit.serializer()
			element is JsonObject && "String" in element -> BasicExternalEnum.String.serializer()
			element is JsonObject && "Number" in element -> BasicExternalEnum.Number.serializer()
			element is JsonObject && "Struct" in element -> BasicExternalEnum.Struct.serializer()
			element is JsonObject && "Nested" in element -> BasicExternalEnum.Nested.serializer()
			else -> throw SerializationException("Unknown element: $element")
		}
	}
}