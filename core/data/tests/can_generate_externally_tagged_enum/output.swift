/** Simple helper struct */
public struct ItemValue: Codable, Hashable {
	public var field: String

	public init(field: String) {
		self.field = field
	}
}

/** Basic externally tagged enum */
public enum BasicExternalEnum: Codable, Hashable {
	/** Unit variant */
	case unit

	/** String variant */
	case string(String)

	/** Number variant */
	case number(Int32)

	/** Struct variant */
	case `struct`(Struct)

	/** Nested variant */
	case nested(ItemValue)

	public struct Struct: Codable, Hashable {
		public var field1: String
		public var field2: Int32

		public init(field1: String, field2: Int32) {
			self.field1 = field1
			self.field2 = field2
		}
	}

	private enum CodingKeys: String, CodingKey {
		case type
	}

	private enum TypeKeys: String, CodingKey {
		case unit = "Unit"
		case string = "String"
		case number = "Number"
		case `struct` = "Struct"
		case nested = "Nested"
	}

	public init(from decoder: Decoder) throws {
		let container = try decoder.container(keyedBy: CodingKeys.self)
		if let type = try? container.decodeNil(forKey: .type) {
			self = .unit
			return
		}

		let type = try decoder.singleValueContainer()

		if let value = try? type.decode(String.self) {
			if value == "Unit" {
				self = .unit
				return
			}
		}

		if let nestedContainer = try? decoder.container(keyedBy: TypeKeys.self) {
			if let value = try? nestedContainer.decode(String.self, forKey: .string) {
				self = .string(value)
				return
			}

			if let value = try? nestedContainer.decode(Int32.self, forKey: .number) {
				self = .number(value)
				return
			}

			if let value = try? nestedContainer.decode(Struct.self, forKey: .struct) {
				self = .struct(value)
				return
			}

			if let value = try? nestedContainer.decode(ItemValue.self, forKey: .nested) {
				self = .nested(value)
				return
			}
		}

		throw DecodingError.dataCorrupted(
			DecodingError.Context(
				codingPath: decoder.codingPath,
				debugDescription: "Unable to decode BasicExternalEnum"
			)
		)
	}

	public func encode(to encoder: Encoder) throws {
		var container: KeyedEncodingContainer<TypeKeys>

		switch self {
		case .unit:
			var container = encoder.singleValueContainer()
			try container.encode("Unit")
			return
		case .string(let value):
			container = encoder.container(keyedBy: TypeKeys.self)
			try container.encode(value, forKey: .string)
		case .number(let value):
			container = encoder.container(keyedBy: TypeKeys.self)
			try container.encode(value, forKey: .number)
		case .struct(let value):
			container = encoder.container(keyedBy: TypeKeys.self)
			try container.encode(value, forKey: .struct)
		case .nested(let value):
			container = encoder.container(keyedBy: TypeKeys.self)
			try container.encode(value, forKey: .nested)
		}
	}
}