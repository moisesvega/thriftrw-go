// Code generated by thriftrw v1.6.0. DO NOT EDIT.
// @generated

// Copyright (c) 2017 Uber Technologies, Inc.
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

package api

import (
	"errors"
	"fmt"
	"go.uber.org/thriftrw/wire"
	"strings"
)

type ServiceGenerator_Generate_Args struct {
	Request *GenerateServiceRequest `json:"request,omitempty"`
}

func (v *ServiceGenerator_Generate_Args) ToWire() (wire.Value, error) {
	var (
		fields [1]wire.Field
		i      int = 0
		w      wire.Value
		err    error
	)
	if v.Request != nil {
		w, err = v.Request.ToWire()
		if err != nil {
			return w, err
		}
		fields[i] = wire.Field{ID: 1, Value: w}
		i++
	}
	return wire.NewValueStruct(wire.Struct{Fields: fields[:i]}), nil
}

func _GenerateServiceRequest_Read(w wire.Value) (*GenerateServiceRequest, error) {
	var v GenerateServiceRequest
	err := v.FromWire(w)
	return &v, err
}

func (v *ServiceGenerator_Generate_Args) FromWire(w wire.Value) error {
	var err error
	for _, field := range w.GetStruct().Fields {
		switch field.ID {
		case 1:
			if field.Value.Type() == wire.TStruct {
				v.Request, err = _GenerateServiceRequest_Read(field.Value)
				if err != nil {
					return err
				}
			}
		}
	}
	return nil
}

func (v *ServiceGenerator_Generate_Args) String() string {
	if v == nil {
		return "<nil>"
	}
	var fields [1]string
	i := 0
	if v.Request != nil {
		fields[i] = fmt.Sprintf("Request: %v", v.Request)
		i++
	}
	return fmt.Sprintf("ServiceGenerator_Generate_Args{%v}", strings.Join(fields[:i], ", "))
}

func (v *ServiceGenerator_Generate_Args) Equals(rhs *ServiceGenerator_Generate_Args) bool {
	if !((v.Request == nil && rhs.Request == nil) || (v.Request != nil && rhs.Request != nil && v.Request.Equals(rhs.Request))) {
		return false
	}
	return true
}

func (v *ServiceGenerator_Generate_Args) MethodName() string {
	return "generate"
}

func (v *ServiceGenerator_Generate_Args) EnvelopeType() wire.EnvelopeType {
	return wire.Call
}

var ServiceGenerator_Generate_Helper = struct {
	Args           func(request *GenerateServiceRequest) *ServiceGenerator_Generate_Args
	IsException    func(error) bool
	WrapResponse   func(*GenerateServiceResponse, error) (*ServiceGenerator_Generate_Result, error)
	UnwrapResponse func(*ServiceGenerator_Generate_Result) (*GenerateServiceResponse, error)
}{}

func init() {
	ServiceGenerator_Generate_Helper.Args = func(request *GenerateServiceRequest) *ServiceGenerator_Generate_Args {
		return &ServiceGenerator_Generate_Args{Request: request}
	}
	ServiceGenerator_Generate_Helper.IsException = func(err error) bool {
		switch err.(type) {
		default:
			return false
		}
	}
	ServiceGenerator_Generate_Helper.WrapResponse = func(success *GenerateServiceResponse, err error) (*ServiceGenerator_Generate_Result, error) {
		if err == nil {
			return &ServiceGenerator_Generate_Result{Success: success}, nil
		}
		return nil, err
	}
	ServiceGenerator_Generate_Helper.UnwrapResponse = func(result *ServiceGenerator_Generate_Result) (success *GenerateServiceResponse, err error) {
		if result.Success != nil {
			success = result.Success
			return
		}
		err = errors.New("expected a non-void result")
		return
	}
}

type ServiceGenerator_Generate_Result struct {
	Success *GenerateServiceResponse `json:"success,omitempty"`
}

func (v *ServiceGenerator_Generate_Result) ToWire() (wire.Value, error) {
	var (
		fields [1]wire.Field
		i      int = 0
		w      wire.Value
		err    error
	)
	if v.Success != nil {
		w, err = v.Success.ToWire()
		if err != nil {
			return w, err
		}
		fields[i] = wire.Field{ID: 0, Value: w}
		i++
	}
	if i != 1 {
		return wire.Value{}, fmt.Errorf("ServiceGenerator_Generate_Result should have exactly one field: got %v fields", i)
	}
	return wire.NewValueStruct(wire.Struct{Fields: fields[:i]}), nil
}

func _GenerateServiceResponse_Read(w wire.Value) (*GenerateServiceResponse, error) {
	var v GenerateServiceResponse
	err := v.FromWire(w)
	return &v, err
}

func (v *ServiceGenerator_Generate_Result) FromWire(w wire.Value) error {
	var err error
	for _, field := range w.GetStruct().Fields {
		switch field.ID {
		case 0:
			if field.Value.Type() == wire.TStruct {
				v.Success, err = _GenerateServiceResponse_Read(field.Value)
				if err != nil {
					return err
				}
			}
		}
	}
	count := 0
	if v.Success != nil {
		count++
	}
	if count != 1 {
		return fmt.Errorf("ServiceGenerator_Generate_Result should have exactly one field: got %v fields", count)
	}
	return nil
}

func (v *ServiceGenerator_Generate_Result) String() string {
	if v == nil {
		return "<nil>"
	}
	var fields [1]string
	i := 0
	if v.Success != nil {
		fields[i] = fmt.Sprintf("Success: %v", v.Success)
		i++
	}
	return fmt.Sprintf("ServiceGenerator_Generate_Result{%v}", strings.Join(fields[:i], ", "))
}

func (v *ServiceGenerator_Generate_Result) Equals(rhs *ServiceGenerator_Generate_Result) bool {
	if !((v.Success == nil && rhs.Success == nil) || (v.Success != nil && rhs.Success != nil && v.Success.Equals(rhs.Success))) {
		return false
	}
	return true
}

func (v *ServiceGenerator_Generate_Result) MethodName() string {
	return "generate"
}

func (v *ServiceGenerator_Generate_Result) EnvelopeType() wire.EnvelopeType {
	return wire.Reply
}
