package com.agilebits.onepassword

import io.circe._
import io.circe.generic.semiauto._
import io.circe.syntax._

/** Tests renaming variants */
sealed trait CamelCaseEnum

object CamelCaseEnum {
	case object UnitVariant extends CamelCaseEnum
	case class SimpleValue(value: String) extends CamelCaseEnum
	case class ComplexValue(field: Int) extends CamelCaseEnum
	
	implicit val encoder: Encoder[CamelCaseEnum] = Encoder.instance {
		case UnitVariant => Json.fromString("unitVariant")
		case SimpleValue(value) => Json.obj("simpleValue" -> value.asJson)
		case ComplexValue(field) => Json.obj("complexValue" -> Json.obj("field" -> field.asJson))
	}
	
	implicit val decoder: Decoder[CamelCaseEnum] = new Decoder[CamelCaseEnum] {
		def apply(c: HCursor): Decoder.Result[CamelCaseEnum] = {
			if (c.value.isString && c.value.asString.contains("unitVariant")) {
				Right(UnitVariant)
			} else {
				val keys = c.keys.getOrElse(Nil)
				if (keys.isEmpty) {
					Left(DecodingFailure("Expected object with a single key", c.history))
				} else {
					val key = keys.head
					key match {
						case "simpleValue" => c.downField("simpleValue").as[String].map(SimpleValue)
						case "complexValue" => c.downField("complexValue").downField("field").as[Int].map(ComplexValue)
						case _ => Left(DecodingFailure(s"Unknown key: $key", c.history))
					}
				}
			}
		}
	}
}

/** Tests specific renames on variants */
sealed trait SpecificRenameEnum

object SpecificRenameEnum {
	case object Regular extends SpecificRenameEnum
	case class CustomName(value: String) extends SpecificRenameEnum
	case class CustomStruct(value: Int) extends SpecificRenameEnum
	
	implicit val encoder: Encoder[SpecificRenameEnum] = Encoder.instance {
		case Regular => Json.fromString("Regular")
		case CustomName(value) => Json.obj("custom_name" -> value.asJson)
		case CustomStruct(value) => Json.obj("custom_struct" -> Json.obj("value" -> value.asJson))
	}
	
	implicit val decoder: Decoder[SpecificRenameEnum] = new Decoder[SpecificRenameEnum] {
		def apply(c: HCursor): Decoder.Result[SpecificRenameEnum] = {
			if (c.value.isString && c.value.asString.contains("Regular")) {
				Right(Regular)
			} else {
				val keys = c.keys.getOrElse(Nil)
				if (keys.isEmpty) {
					Left(DecodingFailure("Expected object with a single key", c.history))
				} else {
					val key = keys.head
					key match {
						case "custom_name" => c.downField("custom_name").as[String].map(CustomName)
						case "custom_struct" => c.downField("custom_struct").downField("value").as[Int].map(CustomStruct)
						case _ => Left(DecodingFailure(s"Unknown key: $key", c.history))
					}
				}
			}
		}
	}
}

/** Mixed kebab and specific renames */
sealed trait MixedRenameEnum

object MixedRenameEnum {
	case object UnitValue extends MixedRenameEnum
	case class CustomValue(value: String) extends MixedRenameEnum
	case class ComplexValue(field: Int) extends MixedRenameEnum
	
	implicit val encoder: Encoder[MixedRenameEnum] = Encoder.instance {
		case UnitValue => Json.fromString("unit-value")
		case CustomValue(value) => Json.obj("CUSTOM" -> value.asJson)
		case ComplexValue(field) => Json.obj("complex-value" -> Json.obj("field" -> field.asJson))
	}
	
	implicit val decoder: Decoder[MixedRenameEnum] = new Decoder[MixedRenameEnum] {
		def apply(c: HCursor): Decoder.Result[MixedRenameEnum] = {
			if (c.value.isString && c.value.asString.contains("unit-value")) {
				Right(UnitValue)
			} else {
				val keys = c.keys.getOrElse(Nil)
				if (keys.isEmpty) {
					Left(DecodingFailure("Expected object with a single key", c.history))
				} else {
					val key = keys.head
					key match {
						case "CUSTOM" => c.downField("CUSTOM").as[String].map(CustomValue)
						case "complex-value" => c.downField("complex-value").downField("field").as[Int].map(ComplexValue)
						case _ => Left(DecodingFailure(s"Unknown key: $key", c.history))
					}
				}
			}
		}
	}
}