package com.agilebits.onepassword.colorsModule

import kotlinx.serialization.*
import kotlinx.serialization.json.*
import kotlinx.serialization.descriptors.*
import kotlinx.serialization.encoding.*

/** Tests renaming variants */
@Serializable(CamelCaseEnumSerializer::class)
sealed class CamelCaseEnum {
	@Serializable
	object UnitVariant : CamelCaseEnum()
	
	@Serializable
	data class SimpleValue(val value: String) : CamelCaseEnum()
	
	@Serializable
	data class ComplexValue(val field: Int) : CamelCaseEnum()
}

class CamelCaseEnumSerializer : JsonContentPolymorphicSerializer<CamelCaseEnum>(CamelCaseEnum::class) {
	override fun selectDeserializer(element: JsonElement): DeserializationStrategy<out CamelCaseEnum> {
		return when {
			element is JsonPrimitive && element.isString && element.content == "unitVariant" -> CamelCaseEnum.UnitVariant.serializer()
			element is JsonObject && "simpleValue" in element -> CamelCaseEnum.SimpleValue.serializer()
			element is JsonObject && "complexValue" in element -> CamelCaseEnum.ComplexValue.serializer()
			else -> throw SerializationException("Unknown element: $element")
		}
	}
}

/** Tests specific renames on variants */
@Serializable(SpecificRenameEnumSerializer::class)
sealed class SpecificRenameEnum {
	@Serializable
	object Regular : SpecificRenameEnum()
	
	@Serializable
	data class CustomName(val value: String) : SpecificRenameEnum()
	
	@Serializable
	data class CustomStruct(val value: Int) : SpecificRenameEnum()
}

class SpecificRenameEnumSerializer : JsonContentPolymorphicSerializer<SpecificRenameEnum>(SpecificRenameEnum::class) {
	override fun selectDeserializer(element: JsonElement): DeserializationStrategy<out SpecificRenameEnum> {
		return when {
			element is JsonPrimitive && element.isString && element.content == "Regular" -> SpecificRenameEnum.Regular.serializer()
			element is JsonObject && "custom_name" in element -> SpecificRenameEnum.CustomName.serializer()
			element is JsonObject && "custom_struct" in element -> SpecificRenameEnum.CustomStruct.serializer()
			else -> throw SerializationException("Unknown element: $element")
		}
	}
}

/** Mixed kebab and specific renames */
@Serializable(MixedRenameEnumSerializer::class)
sealed class MixedRenameEnum {
	@Serializable
	object UnitValue : MixedRenameEnum()
	
	@Serializable
	data class CustomValue(val value: String) : MixedRenameEnum()
	
	@Serializable
	data class ComplexValue(val field: Int) : MixedRenameEnum()
}

class MixedRenameEnumSerializer : JsonContentPolymorphicSerializer<MixedRenameEnum>(MixedRenameEnum::class) {
	override fun selectDeserializer(element: JsonElement): DeserializationStrategy<out MixedRenameEnum> {
		return when {
			element is JsonPrimitive && element.isString && element.content == "unit-value" -> MixedRenameEnum.UnitValue.serializer()
			element is JsonObject && "CUSTOM" in element -> MixedRenameEnum.CustomValue.serializer()
			element is JsonObject && "complex-value" in element -> MixedRenameEnum.ComplexValue.serializer()
			else -> throw SerializationException("Unknown element: $element")
		}
	}
}