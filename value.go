package quad

import (
	"strconv"
)

// IsValidValue checks if the value is valid. It returns false if the value is nil, an empty IRI or an empty BNode.
func IsValidValue(v Value) bool {
	if v == nil {
		return false
	}
	switch v := v.(type) {
	case RawValue:
		return v != ""
	}
	return true
}

// Value is a type used by all quad directions.
type Value interface {
	String() string
	// Native converts Value to a closest native Go type.
	//
	// If type has no analogs in Go, Native return an object itself.
	Native() interface{}
}

var (
	_ Identifier = RawValue("")
)

// Identifier is a union of IRI and BNode.
type Identifier interface {
	Value
	isIdentifier()
}

// StringOf safely call v.String, returning empty string in case of nil Value.
func StringOf(v Value) string {
	if v == nil {
		return ""
	}
	return v.String()
}

// NativeOf safely call v.Native, returning nil in case of nil Value.
func NativeOf(v Value) interface{} {
	if v == nil {
		return nil
	}
	return v.Native()
}

// RawValue is a value contained in a string with no special formatting defined.
type RawValue string

// isIdentifier implements Identifier.
func (s RawValue) isIdentifier() {}

// String prints the raw value.
func (s RawValue) String() string { return string(s) }

// GoString overrides IRI's %#v printing behaviour to include the type name.
func (s RawValue) GoString() string {
	return "quad.RawValue(" + strconv.Quote(string(s)) + ")"
}

// Native returns an IRI value unchanged (to not collide with String values).
func (s RawValue) Native() interface{} {
	return string(s)
}

// Native support for basic types

// StringConversion is a function to convert string values with a
// specific IRI type to their native equivalents.
type StringConversion func(string) (Value, error)
