// Code generated by thriftrw v1.28.0. DO NOT EDIT.
// @generated

package other_constants

import (
	structs "go.uber.org/thriftrw/gen/internal/tests/structs"
	thriftreflect "go.uber.org/thriftrw/thriftreflect"
)

var ListOfInts []int32 = []int32{
	1,
	2,
	3,
}

var SomePoint *structs.Point = &structs.Point{
	X: 1,
	Y: 2,
}

// ThriftModule represents the IDL file used to generate this package.
var ThriftModule = &thriftreflect.ThriftModule{
	Name:     "other_constants",
	Package:  "go.uber.org/thriftrw/gen/internal/tests/other_constants",
	FilePath: "other_constants.thrift",
	SHA1:     "578e8e5aafda10921bb99b58be9a3714c78e31fc",
	Includes: []*thriftreflect.ThriftModule{
		structs.ThriftModule,
	},
	Raw: rawIDL,
}

const rawIDL = "include \"./structs.thrift\"\n\nconst list<i32> listOfInts = [1, 2, 3]\n\nconst structs.Point some_point = {\"x\": 1, \"y\": 2.0}\n"
