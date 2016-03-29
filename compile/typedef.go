// Copyright (c) 2016 Uber Technologies, Inc.
//
// Permission is hereby granted, free of charge, to any person obtaining a copy
// of this software and associated documentation files (the "Software"), to deal
// in the Software without restriction, including without limitation the rights
// to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
// copies of the Software, and to permit persons to whom the Software is
// furnished to do so, subject to the following conditions:
//
// The above copyright notice and this permission notice shall be included in
// all copies or substantial portions of the Software.
//
// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
// IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
// FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
// AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
// LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
// OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN
// THE SOFTWARE.

package compile

import (
	"github.com/thriftrw/thriftrw-go/ast"
	"github.com/thriftrw/thriftrw-go/wire"
)

// TypedefSpec represents an alias to another type in the Thrift file.
type TypedefSpec struct {
	linkOnce

	Name   string
	File   string
	Target TypeSpec
}

// compileTypedef compiles the given Typedef AST into a TypedefSpec.
func compileTypedef(file string, src *ast.Typedef) *TypedefSpec {
	return &TypedefSpec{
		Name:   src.Name,
		File:   file,
		Target: compileTypeReference(src.Type),
	}
}

// TypeCode gets the wire type for the typedef.
func (t *TypedefSpec) TypeCode() wire.Type {
	return t.Target.TypeCode()
}

// Link links the Target TypeSpec for this typedef in the given scope.
func (t *TypedefSpec) Link(scope Scope) (TypeSpec, error) {
	if t.linked() {
		return t, nil
	}

	var err error
	t.Target, err = t.Target.Link(scope)
	return t, err
}

// ThriftName is the name of the typedef as it appears in the Thrift file.
func (t *TypedefSpec) ThriftName() string {
	return t.Name
}

// ThriftFile for TypedefSpec.
func (t *TypedefSpec) ThriftFile() string {
	return t.File
}

// ForEachTypeReference for TypedefSpec
func (t *TypedefSpec) ForEachTypeReference(f func(TypeSpec) error) error {
	return f(t.Target)
}
