package obj

import (
	"fmt"
	"reflect"

	. "github.com/polydawn/refmt/tok"
)

// ErrInvalidUnmarshalTarget describes an invalid argument passed to UnmarshalDriver.Bind.
// (Unmarshalling must target a non-nil pointer so that it can address the value.)
type ErrInvalidUnmarshalTarget struct {
	Type reflect.Type
}

func (e ErrInvalidUnmarshalTarget) Error() string {
	if e.Type == nil {
		return "invalid unmarshal target (nil)"
	}
	if e.Type.Kind() != reflect.Ptr {
		return "invalid unmarshal target (non-pointer " + e.Type.String() + ")"
	}
	return "invalid unmarshal target: (nil " + e.Type.String() + ")"
}

// ErrUnmarshalIncongruent is the error returned when unmarshalling cannot
// coerce the tokens in the stream into the kind of variables the unmarshal is targetting,
// for example if a map open token comes when an int is expected,
// or an int token comes when a string is expected.
type ErrUnmarshalIncongruent struct {
	Token Token
	Value reflect.Value
}

func (e ErrUnmarshalIncongruent) Error() string {
	return fmt.Sprintf("cannot assign %s to %s field", e.Token, e.Value.Kind())
}

// ErrUnexpectedTokenType is the error returned when unmarshalling recieves a
// completely invalid transition, such as when a map value is expected, but the
// map suddenly closes, or an array close is recieved with no matching array open.
type ErrUnexpectedTokenType struct {
	Got      TokenType // Token in the stream that triggered the error.
	Expected string    // Freeform string describing valid token types.  Often a summary like "array close or start of value", or "map close or key".
}

func (e ErrUnexpectedTokenType) Error() string {
	return fmt.Sprintf("unexpected %s token; expected %s", e.Got, e.Expected)
}
