// Code generated by thriftrw v1.28.0. DO NOT EDIT.
// @generated

package hyphenated_file

import (
	errors "errors"
	fmt "fmt"
	multierr "go.uber.org/multierr"
	non_hyphenated "go.uber.org/thriftrw/gen/internal/tests/non_hyphenated"
	thriftreflect "go.uber.org/thriftrw/thriftreflect"
	wire "go.uber.org/thriftrw/wire"
	zapcore "go.uber.org/zap/zapcore"
	strings "strings"
)

type DocumentStruct struct {
	Second *non_hyphenated.Second `json:"second,required"`
}

// ToWire translates a DocumentStruct struct into a Thrift-level intermediate
// representation. This intermediate representation may be serialized
// into bytes using a ThriftRW protocol implementation.
//
// An error is returned if the struct or any of its fields failed to
// validate.
//
//   x, err := v.ToWire()
//   if err != nil {
//     return err
//   }
//
//   if err := binaryProtocol.Encode(x, writer); err != nil {
//     return err
//   }
func (v *DocumentStruct) ToWire() (wire.Value, error) {
	var (
		fields [1]wire.Field
		i      int = 0
		w      wire.Value
		err    error
	)

	if v.Second == nil {
		return w, errors.New("field Second of DocumentStruct is required")
	}
	w, err = v.Second.ToWire()
	if err != nil {
		return w, err
	}
	fields[i] = wire.Field{ID: 1, Value: w}
	i++

	return wire.NewValueStruct(wire.Struct{Fields: fields[:i]}), nil
}

func _Second_Read(w wire.Value) (*non_hyphenated.Second, error) {
	var v non_hyphenated.Second
	err := v.FromWire(w)
	return &v, err
}

// FromWire deserializes a DocumentStruct struct from its Thrift-level
// representation. The Thrift-level representation may be obtained
// from a ThriftRW protocol implementation.
//
// An error is returned if we were unable to build a DocumentStruct struct
// from the provided intermediate representation.
//
//   x, err := binaryProtocol.Decode(reader, wire.TStruct)
//   if err != nil {
//     return nil, err
//   }
//
//   var v DocumentStruct
//   if err := v.FromWire(x); err != nil {
//     return nil, err
//   }
//   return &v, nil
func (v *DocumentStruct) FromWire(w wire.Value) error {
	var err error

	secondIsSet := false

	for _, field := range w.GetStruct().Fields {
		switch field.ID {
		case 1:
			if field.Value.Type() == wire.TStruct {
				v.Second, err = _Second_Read(field.Value)
				if err != nil {
					return err
				}
				secondIsSet = true
			}
		}
	}

	if !secondIsSet {
		return errors.New("field Second of DocumentStruct is required")
	}

	return nil
}

// String returns a readable string representation of a DocumentStruct
// struct.
func (v *DocumentStruct) String() string {
	if v == nil {
		return "<nil>"
	}

	var fields [1]string
	i := 0
	fields[i] = fmt.Sprintf("Second: %v", v.Second)
	i++

	return fmt.Sprintf("DocumentStruct{%v}", strings.Join(fields[:i], ", "))
}

// Equals returns true if all the fields of this DocumentStruct match the
// provided DocumentStruct.
//
// This function performs a deep comparison.
func (v *DocumentStruct) Equals(rhs *DocumentStruct) bool {
	if v == nil {
		return rhs == nil
	} else if rhs == nil {
		return false
	}
	if !v.Second.Equals(rhs.Second) {
		return false
	}

	return true
}

// MarshalLogObject implements zapcore.ObjectMarshaler, enabling
// fast logging of DocumentStruct.
func (v *DocumentStruct) MarshalLogObject(enc zapcore.ObjectEncoder) (err error) {
	if v == nil {
		return nil
	}
	err = multierr.Append(err, enc.AddObject("second", v.Second))
	return err
}

// GetSecond returns the value of Second if it is set or its
// zero value if it is unset.
func (v *DocumentStruct) GetSecond() (o *non_hyphenated.Second) {
	if v != nil {
		o = v.Second
	}
	return
}

// IsSetSecond returns true if Second is not nil.
func (v *DocumentStruct) IsSetSecond() bool {
	return v != nil && v.Second != nil
}

// ThriftModule represents the IDL file used to generate this package.
var ThriftModule = &thriftreflect.ThriftModule{
	Name:     "hyphenated-file",
	Package:  "go.uber.org/thriftrw/gen/internal/tests/hyphenated-file",
	FilePath: "hyphenated-file.thrift",
	SHA1:     "8513913ac76a3ba1c6b2b3b6fb241462e162c446",
	Includes: []*thriftreflect.ThriftModule{
		non_hyphenated.ThriftModule,
	},
	Raw: rawIDL,
}

const rawIDL = "include \"./non_hyphenated.thrift\"\n\nstruct DocumentStruct {\n 1: required non_hyphenated.Second second\n}"
