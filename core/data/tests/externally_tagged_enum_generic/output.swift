import Foundation

/// Helper for nested generic types
public struct Container<T: Codable>: Codable {
	public let value: T

	public init(value: T) {
		self.value = value
	}
}


/// Generated type representing the anonymous struct variant `Struct` of the `GenericExternalEnum` Rust enum
public struct GenericExternalEnumStructInner<T: Codable, U: Codable>: Codable {
	public let t_value: T
	public let u_value: U

	public init(t_value: T, u_value: U) {
		self.t_value = t_value
		self.u_value = u_value
	}
}
/// Generic externally tagged enum
public enum GenericExternalEnum<T: Codable, U: Codable>: Codable {
	/// Unit variant
	case unit
	/// Generic tuple variant with first type
	case first(T)
	/// Generic tuple variant with second type
	case second(U)
	/// Generic struct variant
	case `struct`(Struct)

	public struct Struct: Codable, Hashable {
		public var t_value: T
		public var u_value: U
		
		public init(t_value: T, u_value: U) {
			self.t_value = t_value
			self.u_value = u_value
		}
	}
	/// Container with generic type
	case nested(Container<T>)

	private enum CodingKeys: String, CodingKey {
		case type
	}

	private enum TypeKeys: String, CodingKey {
		case unit = "Unit"
		case first = "First"
		case second = "Second"
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
			if let value = try? nestedContainer.decode(T.self, forKey: .first) {
				self = .first(value)
				return
			}
			if let value = try? nestedContainer.decode(U.self, forKey: .second) {
				self = .second(value)
				return
			}
			if let value = try? nestedContainer.decode(Struct.self, forKey: .`struct`) {
				self = .`struct`(value)
				return
			}
			if let value = try? nestedContainer.decode(Container<T>.self, forKey: .nested) {
				self = .nested(value)
				return
			}
		}
		
		throw DecodingError.dataCorrupted(
			DecodingError.Context(
				codingPath: decoder.codingPath,
				debugDescription: "Unable to decode GenericExternalEnum"
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
		case .first(let value):
			container = encoder.container(keyedBy: TypeKeys.self)
			try container.encode(value, forKey: .first)
		case .second(let value):
			container = encoder.container(keyedBy: TypeKeys.self)
			try container.encode(value, forKey: .second)
		case .`struct`(let value):
			container = encoder.container(keyedBy: TypeKeys.self)
			try container.encode(value, forKey: .`struct`)
		case .nested(let value):
			container = encoder.container(keyedBy: TypeKeys.self)
			try container.encode(value, forKey: .nested)
		}
	}
}
