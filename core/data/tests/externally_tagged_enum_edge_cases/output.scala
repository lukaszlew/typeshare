package com.agilebits.onepassword

import io.circe._
import io.circe.generic.semiauto._
import io.circe.syntax._

case class NestedStruct(
	nested_field: String
)

object NestedStruct {
	implicit val decoder: Decoder[NestedStruct] = deriveDecoder
	implicit val encoder: Encoder[NestedStruct] = deriveEncoder
}

/** Tests various edge cases */
sealed trait EdgeCaseEnum

object EdgeCaseEnum {
	/** Empty variant */
	case object Empty extends EdgeCaseEnum
	
	/** Optional string */
	case class OptionalString(value: Option[String]) extends EdgeCaseEnum
	
	/** Optional nested struct */
	case class OptionalNested(value: Option[NestedStruct]) extends EdgeCaseEnum
	
	/** Array of values */
	case class Array(value: List[Int]) extends EdgeCaseEnum
	
	/** Array of structs */
	case class StructArray(value: List[NestedStruct]) extends EdgeCaseEnum
	
	/** Complex nested structure */
	case class Complex(
		a: String,
		b: Option[Int],
		c: List[NestedStruct],
		d: Option[List[String]]
	) extends EdgeCaseEnum
	
	implicit val encoder: Encoder[EdgeCaseEnum] = Encoder.instance {
		case Empty => Json.fromString("Empty")
		case OptionalString(value) => Json.obj("OptionalString" -> value.asJson)
		case OptionalNested(value) => Json.obj("OptionalNested" -> value.asJson)
		case Array(value) => Json.obj("Array" -> value.asJson)
		case StructArray(value) => Json.obj("StructArray" -> value.asJson)
		case Complex(a, b, c, d) => Json.obj(
			"Complex" -> Json.obj(
				"a" -> a.asJson,
				"b" -> b.asJson,
				"c" -> c.asJson,
				"d" -> d.asJson
			)
		)
	}
	
	implicit val decoder: Decoder[EdgeCaseEnum] = new Decoder[EdgeCaseEnum] {
		def apply(c: HCursor): Decoder.Result[EdgeCaseEnum] = {
			if (c.value.isString && c.value.asString.contains("Empty")) {
				Right(Empty)
			} else {
				val keys = c.keys.getOrElse(Nil)
				if (keys.isEmpty) {
					Left(DecodingFailure("Expected object with a single key", c.history))
				} else {
					val key = keys.head
					key match {
						case "OptionalString" => c.downField("OptionalString").as[Option[String]].map(OptionalString)
						case "OptionalNested" => c.downField("OptionalNested").as[Option[NestedStruct]].map(OptionalNested)
						case "Array" => c.downField("Array").as[List[Int]].map(Array)
						case "StructArray" => c.downField("StructArray").as[List[NestedStruct]].map(StructArray)
						case "Complex" => 
							for {
								a <- c.downField("Complex").downField("a").as[String]
								b <- c.downField("Complex").downField("b").as[Option[Int]]
								c <- c.downField("Complex").downField("c").as[List[NestedStruct]]
								d <- c.downField("Complex").downField("d").as[Option[List[String]]]
							} yield Complex(a, b, c, d)
						case _ => Left(DecodingFailure(s"Unknown key: $key", c.history))
					}
				}
			}
		}
	}
}