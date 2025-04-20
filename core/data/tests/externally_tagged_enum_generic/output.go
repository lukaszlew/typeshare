package proto

import "encoding/json"

// Helper for nested generic types
type Container[T any] struct {
	Value T `json:"value"`
}
