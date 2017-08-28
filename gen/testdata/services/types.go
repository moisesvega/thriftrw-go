// Code generated by thriftrw v1.6.0. DO NOT EDIT.
// @generated

package services

import (
	"bytes"
	"errors"
	"fmt"
	"go.uber.org/thriftrw/wire"
	"strings"
)

type ConflictingNamesSetValueArgs struct {
	Key   string `json:"key,required"`
	Value []byte `json:"value,required"`
}

func (v *ConflictingNamesSetValueArgs) ToWire() (wire.Value, error) {
	var (
		fields [2]wire.Field
		i      int = 0
		w      wire.Value
		err    error
	)
	w, err = wire.NewValueString(v.Key), error(nil)
	if err != nil {
		return w, err
	}
	fields[i] = wire.Field{ID: 1, Value: w}
	i++
	if v.Value == nil {
		return w, errors.New("field Value of ConflictingNamesSetValueArgs is required")
	}
	w, err = wire.NewValueBinary(v.Value), error(nil)
	if err != nil {
		return w, err
	}
	fields[i] = wire.Field{ID: 2, Value: w}
	i++
	return wire.NewValueStruct(wire.Struct{Fields: fields[:i]}), nil
}

func (v *ConflictingNamesSetValueArgs) FromWire(w wire.Value) error {
	var err error
	keyIsSet := false
	valueIsSet := false
	for _, field := range w.GetStruct().Fields {
		switch field.ID {
		case 1:
			if field.Value.Type() == wire.TBinary {
				v.Key, err = field.Value.GetString(), error(nil)
				if err != nil {
					return err
				}
				keyIsSet = true
			}
		case 2:
			if field.Value.Type() == wire.TBinary {
				v.Value, err = field.Value.GetBinary(), error(nil)
				if err != nil {
					return err
				}
				valueIsSet = true
			}
		}
	}
	if !keyIsSet {
		return errors.New("field Key of ConflictingNamesSetValueArgs is required")
	}
	if !valueIsSet {
		return errors.New("field Value of ConflictingNamesSetValueArgs is required")
	}
	return nil
}

func (v *ConflictingNamesSetValueArgs) String() string {
	if v == nil {
		return "<nil>"
	}
	var fields [2]string
	i := 0
	fields[i] = fmt.Sprintf("Key: %v", v.Key)
	i++
	fields[i] = fmt.Sprintf("Value: %v", v.Value)
	i++
	return fmt.Sprintf("ConflictingNamesSetValueArgs{%v}", strings.Join(fields[:i], ", "))
}

func (v *ConflictingNamesSetValueArgs) Equals(rhs *ConflictingNamesSetValueArgs) bool {
	if !(v.Key == rhs.Key) {
		return false
	}
	if !bytes.Equal(v.Value, rhs.Value) {
		return false
	}
	return true
}

type InternalError struct {
	Message *string `json:"message,omitempty"`
}

func (v *InternalError) ToWire() (wire.Value, error) {
	var (
		fields [1]wire.Field
		i      int = 0
		w      wire.Value
		err    error
	)
	if v.Message != nil {
		w, err = wire.NewValueString(*(v.Message)), error(nil)
		if err != nil {
			return w, err
		}
		fields[i] = wire.Field{ID: 1, Value: w}
		i++
	}
	return wire.NewValueStruct(wire.Struct{Fields: fields[:i]}), nil
}

func (v *InternalError) FromWire(w wire.Value) error {
	var err error
	for _, field := range w.GetStruct().Fields {
		switch field.ID {
		case 1:
			if field.Value.Type() == wire.TBinary {
				var x string
				x, err = field.Value.GetString(), error(nil)
				v.Message = &x
				if err != nil {
					return err
				}
			}
		}
	}
	return nil
}

func (v *InternalError) String() string {
	if v == nil {
		return "<nil>"
	}
	var fields [1]string
	i := 0
	if v.Message != nil {
		fields[i] = fmt.Sprintf("Message: %v", *(v.Message))
		i++
	}
	return fmt.Sprintf("InternalError{%v}", strings.Join(fields[:i], ", "))
}

func _String_EqualsPtr(lhs, rhs *string) bool {
	if lhs != nil && rhs != nil {
		x := *lhs
		y := *rhs
		return (x == y)
	}
	return lhs == nil && rhs == nil
}

func (v *InternalError) Equals(rhs *InternalError) bool {
	if !_String_EqualsPtr(v.Message, rhs.Message) {
		return false
	}
	return true
}

func (v *InternalError) GetMessage() (o string) {
	if v.Message != nil {
		return *v.Message
	}
	return
}

func (v *InternalError) Error() string {
	return v.String()
}

type Key string

func (v Key) ToWire() (wire.Value, error) {
	x := (string)(v)
	return wire.NewValueString(x), error(nil)
}

func (v Key) String() string {
	x := (string)(v)
	return fmt.Sprint(x)
}

func (v *Key) FromWire(w wire.Value) error {
	x, err := w.GetString(), error(nil)
	*v = (Key)(x)
	return err
}

func (lhs Key) Equals(rhs Key) bool {
	return (lhs == rhs)
}
