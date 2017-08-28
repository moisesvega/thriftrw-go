// Code generated by thriftrw v1.6.0. DO NOT EDIT.
// @generated

package unions

import (
	"go.uber.org/thriftrw/gen/testdata/typedefs"
	"go.uber.org/thriftrw/thriftreflect"
)

var ThriftModule = &thriftreflect.ThriftModule{Name: "unions", Package: "go.uber.org/thriftrw/gen/testdata/unions", FilePath: "unions.thrift", SHA1: "ec5d7ff25ced107d4166f19645b5e492391f21dd", Includes: []*thriftreflect.ThriftModule{typedefs.ThriftModule}, Raw: rawIDL}

const rawIDL = "include \"./typedefs.thrift\"\n\nunion EmptyUnion {}\n\nunion Document {\n    1: typedefs.PDF pdf\n    2: string plainText\n}\n\nunion ArbitraryValue {\n    1: bool boolValue\n    2: i64 int64Value\n    3: string stringValue\n    4: list<ArbitraryValue> listValue\n    5: map<string, ArbitraryValue> mapValue\n}\n"
