package com.agilebits.onepassword

import kotlinx.serialization.Serializable
import kotlinx.serialization.SerialName

/// Helper for nested generic types
@Serializable
data class Container<T> (
	val value: T
)

