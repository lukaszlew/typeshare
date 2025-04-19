import Foundation


/// Generated type representing the anonymous struct variant `ComplexValue` of the `CamelCaseEnum` Rust enum
public struct CamelCaseEnumComplexValueInner: Codable {
	public let field: Int32

	public init(field: Int32) {
		self.field = field
	}
}
/// Tests renaming variants
public enum CamelCaseEnum: Codable {
	case unitVariant
	case simpleValue(String)
	case complexValue(Struct)

	public struct Struct: Codable, Hashable {
		public var field: Int32
		
		public init(field: Int32) {
			self.field = field
		}
	}

	private enum CodingKeys: String, CodingKey {
		case type
	}

	private enum TypeKeys: String, CodingKey {
		case unitVariant = "unitVariant"
		case simpleValue = "simpleValue"
		case complexValue = "complexValue"
	}

	public init(from decoder: Decoder) throws {
		let container = try decoder.container(keyedBy: CodingKeys.self)
		if let type = try? container.decodeNil(forKey: .type) {
			self = .unitVariant
			return
		}
		
		let type = try decoder.singleValueContainer()
		
		if let value = try? type.decode(String.self) {
			if value == "unitVariant" {
				self = .unitVariant
				return
			}
		}
		
		if let nestedContainer = try? decoder.container(keyedBy: TypeKeys.self) {
			if let value = try? nestedContainer.decode(String.self, forKey: .simpleValue) {
				self = .simpleValue(value)
				return
			}
			if let value = try? nestedContainer.decode(Struct.self, forKey: .complexValue) {
				self = .complexValue(value)
				return
			}
		}
		
		throw DecodingError.dataCorrupted(
			DecodingError.Context(
				codingPath: decoder.codingPath,
				debugDescription: "Unable to decode CamelCaseEnum"
			)
		)
	}
	
	public func encode(to encoder: Encoder) throws {
		var container: KeyedEncodingContainer<TypeKeys>
		
		switch self {
		case .unitVariant:
			var container = encoder.singleValueContainer()
			try container.encode("unitVariant")
			return
		case .simpleValue(let value):
			container = encoder.container(keyedBy: TypeKeys.self)
			try container.encode(value, forKey: .simpleValue)
		case .complexValue(let value):
			container = encoder.container(keyedBy: TypeKeys.self)
			try container.encode(value, forKey: .complexValue)
		}
	}
}


/// Generated type representing the anonymous struct variant `ComplexValue` of the `MixedRenameEnum` Rust enum
public struct MixedRenameEnumComplexValueInner: Codable {
	public let field: Int32

	public init(field: Int32) {
		self.field = field
	}
}
/// Mixed kebab and specific renames
public enum MixedRenameEnum: Codable {
	case unitValue
	case customValue(String)
	case complexValue(Struct)

	public struct Struct: Codable, Hashable {
		public var field: Int32
		
		public init(field: Int32) {
			self.field = field
		}
	}

	private enum CodingKeys: String, CodingKey {
		case type
	}

	private enum TypeKeys: String, CodingKey {
		case unitValue = "unit-value"
		case customValue = "CUSTOM"
		case complexValue = "complex-value"
	}

	public init(from decoder: Decoder) throws {
		let container = try decoder.container(keyedBy: CodingKeys.self)
		if let type = try? container.decodeNil(forKey: .type) {
			self = .unitValue
			return
		}
		
		let type = try decoder.singleValueContainer()
		
		if let value = try? type.decode(String.self) {
			if value == "unit-value" {
				self = .unitValue
				return
			}
		}
		
		if let nestedContainer = try? decoder.container(keyedBy: TypeKeys.self) {
			if let value = try? nestedContainer.decode(String.self, forKey: .customValue) {
				self = .customValue(value)
				return
			}
			if let value = try? nestedContainer.decode(Struct.self, forKey: .complexValue) {
				self = .complexValue(value)
				return
			}
		}
		
		throw DecodingError.dataCorrupted(
			DecodingError.Context(
				codingPath: decoder.codingPath,
				debugDescription: "Unable to decode MixedRenameEnum"
			)
		)
	}
	
	public func encode(to encoder: Encoder) throws {
		var container: KeyedEncodingContainer<TypeKeys>
		
		switch self {
		case .unitValue:
			var container = encoder.singleValueContainer()
			try container.encode("unit-value")
			return
		case .customValue(let value):
			container = encoder.container(keyedBy: TypeKeys.self)
			try container.encode(value, forKey: .customValue)
		case .complexValue(let value):
			container = encoder.container(keyedBy: TypeKeys.self)
			try container.encode(value, forKey: .complexValue)
		}
	}
}


/// Generated type representing the anonymous struct variant `CustomStruct` of the `SpecificRenameEnum` Rust enum
public struct SpecificRenameEnumCustomStructInner: Codable {
	public let value: Int32

	public init(value: Int32) {
		self.value = value
	}
}
/// Tests specific renames on variants
public enum SpecificRenameEnum: Codable {
	case regular
	case customName(String)
	case customStruct(Struct)

	public struct Struct: Codable, Hashable {
		public var value: Int32
		
		public init(value: Int32) {
			self.value = value
		}
	}

	private enum CodingKeys: String, CodingKey {
		case type
	}

	private enum TypeKeys: String, CodingKey {
		case regular = "Regular"
		case customName = "custom_name"
		case customStruct = "custom_struct"
	}

	public init(from decoder: Decoder) throws {
		let container = try decoder.container(keyedBy: CodingKeys.self)
		if let type = try? container.decodeNil(forKey: .type) {
			self = .regular
			return
		}
		
		let type = try decoder.singleValueContainer()
		
		if let value = try? type.decode(String.self) {
			if value == "Regular" {
				self = .regular
				return
			}
		}
		
		if let nestedContainer = try? decoder.container(keyedBy: TypeKeys.self) {
			if let value = try? nestedContainer.decode(String.self, forKey: .customName) {
				self = .customName(value)
				return
			}
			if let value = try? nestedContainer.decode(Struct.self, forKey: .customStruct) {
				self = .customStruct(value)
				return
			}
		}
		
		throw DecodingError.dataCorrupted(
			DecodingError.Context(
				codingPath: decoder.codingPath,
				debugDescription: "Unable to decode SpecificRenameEnum"
			)
		)
	}
	
	public func encode(to encoder: Encoder) throws {
		var container: KeyedEncodingContainer<TypeKeys>
		
		switch self {
		case .regular:
			var container = encoder.singleValueContainer()
			try container.encode("Regular")
			return
		case .customName(let value):
			container = encoder.container(keyedBy: TypeKeys.self)
			try container.encode(value, forKey: .customName)
		case .customStruct(let value):
			container = encoder.container(keyedBy: TypeKeys.self)
			try container.encode(value, forKey: .customStruct)
		}
	}
}
