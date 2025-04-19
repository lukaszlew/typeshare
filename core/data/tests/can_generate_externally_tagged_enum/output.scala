package com.agilebits.onepassword

import io.circe._
import io.circe.generic.semiauto._
import io.circe.syntax._

/** Simple helper struct */
case class ItemValue(
  field: String
)

object ItemValue {
  implicit val decoder: Decoder[ItemValue] = deriveDecoder
  implicit val encoder: Encoder[ItemValue] = deriveEncoder
}

/** Basic externally tagged enum */
sealed trait BasicExternalEnum

object BasicExternalEnum {
  /** Unit variant */
  case object Unit extends BasicExternalEnum

  /** String variant */
  case class String(value: java.lang.String) extends BasicExternalEnum

  /** Number variant */
  case class Number(value: Int) extends BasicExternalEnum

  /** Struct variant */
  case class Struct(
    field1: java.lang.String,
    field2: Int
  ) extends BasicExternalEnum

  /** Nested variant */
  case class Nested(value: ItemValue) extends BasicExternalEnum

  implicit val encoder: Encoder[BasicExternalEnum] = Encoder.instance {
    case Unit => Json.fromString("Unit")
    case String(value) => Json.obj("String" -> value.asJson)
    case Number(value) => Json.obj("Number" -> value.asJson)
    case Struct(field1, field2) => Json.obj("Struct" -> Json.obj(
      "field1" -> field1.asJson,
      "field2" -> field2.asJson
    ))
    case Nested(value) => Json.obj("Nested" -> value.asJson)
  }

  implicit val decoder: Decoder[BasicExternalEnum] = new Decoder[BasicExternalEnum] {
    def apply(c: HCursor): Decoder.Result[BasicExternalEnum] = {
      if (c.value.isString && c.value.asString.contains("Unit")) {
        Right(Unit)
      } else {
        val keys = c.keys.getOrElse(Nil)
        if (keys.isEmpty) {
          Left(DecodingFailure("Expected object with a single key", c.history))
        } else {
          val key = keys.head
          key match {
            case "String" => c.downField("String").as[java.lang.String].map(String)
            case "Number" => c.downField("Number").as[Int].map(Number)
            case "Struct" =>
              for {
                field1 <- c.downField("Struct").downField("field1").as[java.lang.String]
                field2 <- c.downField("Struct").downField("field2").as[Int]
              } yield Struct(field1, field2)
            case "Nested" => c.downField("Nested").as[ItemValue].map(Nested)
            case _ => Left(DecodingFailure(s"Unknown key: $key", c.history))
          }
        }
      }
    }
  }
}