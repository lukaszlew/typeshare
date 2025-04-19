public struct NestedStruct: Codable, Hashable {
	public var nested_field: String

	public init(nested_field: String) {
		self.nested_field = nested_field
	}
}

/** Tests various edge cases */
public enum EdgeCaseEnum: Codable, Hashable {
	/** Empty variant */
	case empty
	
	/** Optional string */
	case optionalString(String?)
	
	/** Optional nested struct */
	case optionalNested(NestedStruct?)
	
	/** Array of values */
	case array([Int32])
	
	/** Array of structs */
	case structArray([NestedStruct])
	
	/** Complex nested structure */
	case complex(Complex)
	
	public struct Complex: Codable, Hashable {
		public var a: String
		public var b: Int32?
		public var c: [NestedStruct]
		public var d: [String]?
		
		public init(a: String, b: Int32?, c: [NestedStruct], d: [String]?) {
			self.a = a
			self.b = b
			self.c = c
			self.d = d
		}
	}
	
	private enum CodingKeys: String, CodingKey {
		case type
	}
	
	private enum TypeKeys: String, CodingKey {
		case empty = "Empty"
		case optionalString = "OptionalString"
		case optionalNested = "OptionalNested"
		case array = "Array"
		case structArray = "StructArray"
		case complex = "Complex"
	}
	
	public init(from decoder: Decoder) throws {
		let container = try decoder.container(keyedBy: CodingKeys.self)
		if let type = try? container.decodeNil(forKey: .type) {
			self = .empty
			return
		}
		
		let type = try decoder.singleValueContainer()
		
		if let value = try? type.decode(String.self) {
			if value == "Empty" {
				self = .empty
				return
			}
		}
		
		if let nestedContainer = try? decoder.container(keyedBy: TypeKeys.self) {
			if let value = try? nestedContainer.decode(String?.self, forKey: .optionalString) {
				self = .optionalString(value)
				return
			}
			
			if let value = try? nestedContainer.decode(NestedStruct?.self, forKey: .optionalNested) {
				self = .optionalNested(value)
				return
			}
			
			if let value = try? nestedContainer.decode([Int32].self, forKey: .array) {
				self = .array(value)
				return
			}
			
			if let value = try? nestedContainer.decode([NestedStruct].self, forKey: .structArray) {
				self = .structArray(value)
				return
			}
			
			if let value = try? nestedContainer.decode(Complex.self, forKey: .complex) {
				self = .complex(value)
				return
			}
		}
		
		throw DecodingError.dataCorrupted(
			DecodingError.Context(
				codingPath: decoder.codingPath,
				debugDescription: "Unable to decode EdgeCaseEnum"
			)
		)
	}
	
	public func encode(to encoder: Encoder) throws {
		var container: KeyedEncodingContainer<TypeKeys>
		
		switch self {
		case .empty:
			var container = encoder.singleValueContainer()
			try container.encode("Empty")
			return
		case .optionalString(let value):
			container = encoder.container(keyedBy: TypeKeys.self)
			try container.encode(value, forKey: .optionalString)
		case .optionalNested(let value):
			container = encoder.container(keyedBy: TypeKeys.self)
			try container.encode(value, forKey: .optionalNested)
		case .array(let value):
			container = encoder.container(keyedBy: TypeKeys.self)
			try container.encode(value, forKey: .array)
		case .structArray(let value):
			container = encoder.container(keyedBy: TypeKeys.self)
			try container.encode(value, forKey: .structArray)
		case .complex(let value):
			container = encoder.container(keyedBy: TypeKeys.self)
			try container.encode(value, forKey: .complex)
		}
	}
}