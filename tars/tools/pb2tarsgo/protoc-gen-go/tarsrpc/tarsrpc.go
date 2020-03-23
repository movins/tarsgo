// Go support for Protocol Buffers - Google's data interchange format
//
// Copyright 2015 The Go Authors.  All rights reserved.
// https://github.com/golang/protobuf
//
// Redistribution and use in source and binary forms, with or without
// modification, are permitted provided that the following conditions are
// met:
//
//     * Redistributions of source code must retain the above copyright
// notice, this list of conditions and the following disclaimer.
//     * Redistributions in binary form must reproduce the above
// copyright notice, this list of conditions and the following disclaimer
// in the documentation and/or other materials provided with the
// distribution.
//     * Neither the name of Google Inc. nor the names of its
// contributors may be used to endorse or promote products derived from
// this software without specific prior written permission.
//
// THIS SOFTWARE IS PROVIDED BY THE COPYRIGHT HOLDERS AND CONTRIBUTORS
// "AS IS" AND ANY EXPRESS OR IMPLIED WARRANTIES, INCLUDING, BUT NOT
// LIMITED TO, THE IMPLIED WARRANTIES OF MERCHANTABILITY AND FITNESS FOR
// A PARTICULAR PURPOSE ARE DISCLAIMED. IN NO EVENT SHALL THE COPYRIGHT
// OWNER OR CONTRIBUTORS BE LIABLE FOR ANY DIRECT, INDIRECT, INCIDENTAL,
// SPECIAL, EXEMPLARY, OR CONSEQUENTIAL DAMAGES (INCLUDING, BUT NOT
// LIMITED TO, PROCUREMENT OF SUBSTITUTE GOODS OR SERVICES; LOSS OF USE,
// DATA, OR PROFITS; OR BUSINESS INTERRUPTION) HOWEVER CAUSED AND ON ANY
// THEORY OF LIABILITY, WHETHER IN CONTRACT, STRICT LIABILITY, OR TORT
// (INCLUDING NEGLIGENCE OR OTHERWISE) ARISING IN ANY WAY OUT OF THE USE
// OF THIS SOFTWARE, EVEN IF ADVISED OF THE POSSIBILITY OF SUCH DAMAGE.

// Package tarsrpc outputs tars service descriptions in Go code.
// It runs as a plugin for the Go protocol buffer compiler plugin.
// It is linked in to protoc-gen-go.
package tarsrpc

import (
	"fmt"
	"strings"

	pb "github.com/golang/protobuf/protoc-gen-go/descriptor"
	"github.com/golang/protobuf/protoc-gen-go/generator"
)

// Paths for packages used by code generated in this file,
// relative to the import_prefix of the generator.Generator.
const (
	modelPkgPath   = "gitee.com/bee-circle/tarsgo/tars/model"
	requestPkgPath = "gitee.com/bee-circle/tarsgo/tars/protocol/res/requestf"
	tarsPkgPath    = "gitee.com/bee-circle/tarsgo/tars"
	toolsPath      = "gitee.com/bee-circle/tarsgo/tars/util/tools"
)

func init() {
	generator.RegisterPlugin(new(tarsrpc))
}

// tarsrpc is an implementation of the Go protocol buffer compiler's
// plugin architecture.  It generates bindings for tars rpc support.
type tarsrpc struct {
	gen *generator.Generator
}

//Name returns the name of this plugin
func (t *tarsrpc) Name() string {
	return "tarsrpc"
}

//Init initializes the plugin.
func (t *tarsrpc) Init(gen *generator.Generator) {
	t.gen = gen
}

// upperFirstLatter make the fisrt charater of given string  upper class
func upperFirstLatter(s string) string {
	if len(s) == 0 {
		return ""
	}
	if len(s) == 1 {
		return strings.ToUpper(string(s[0]))
	}
	return strings.ToUpper(string(s[0])) + s[1:]
}

// Given a type name defined in a .proto, return its object.
// Also record that we're using it, to guarantee the associated import.
func (t *tarsrpc) objectNamed(name string) generator.Object {
	t.gen.RecordTypeUse(name)
	return t.gen.ObjectNamed(name)
}

// Given a type name defined in a .proto, return its name as we will print it.
func (t *tarsrpc) typeName(str string) string {
	return t.gen.TypeName(t.objectNamed(str))
}

// GenerateImports generates the import declaration for this file.
func (t *tarsrpc) GenerateImports(file *generator.FileDescriptor) {
}

// P forwards to g.gen.P.
func (t *tarsrpc) P(args ...interface{}) { t.gen.P(args...) }

// Generate generates code for the services in the given file.
func (t *tarsrpc) Generate(file *generator.FileDescriptor) {
	if len(file.FileDescriptorProto.Service) == 0 {
		return
	}
	_ = t.gen.AddImport(modelPkgPath)
	_ = t.gen.AddImport(requestPkgPath)
	_ = t.gen.AddImport(tarsPkgPath)
	_ = t.gen.AddImport(toolsPath)
	_ = t.gen.AddImport("context")
	for i, service := range file.FileDescriptorProto.Service {
		t.generateService(file, service, i)
	}
}

// generateService generates all the code for the named service
func (t *tarsrpc) generateService(file *generator.FileDescriptor, service *pb.ServiceDescriptorProto, index int) {
	originServiceName := service.GetName()
	serviceName := upperFirstLatter(originServiceName)
	t.P("// This following code was generated by tarsrpc")
	t.P(fmt.Sprintf("// Gernerated from %s", file.GetName()))
	t.P(fmt.Sprintf(`type  %s struct {
		s model.Servant
	}
	`, serviceName))
	t.P()

	//generate SetServant
	t.P(fmt.Sprintf(`//SetServant is required by the servant interface.
	func (obj *%s) SetServant(s model.Servant){
		obj.s = s
	}
	`, serviceName))
	t.P()
	//generate AddServant
	t.P(fmt.Sprintf(`//AddServant is required by the servant interface
	func (obj *%s) AddServant(imp imp%s, objStr string){
		tars.AddServant(obj, imp, objStr)
	}`, serviceName, serviceName))

	//generate TarsSetTimeout
	t.P(fmt.Sprintf(`//TarsSetTimeout is required by the servant interface. t is the timeout in ms. 
	func (obj *%s) TarsSetTimeout(t int){
		obj.s.TarsSetTimeout(t)
	}
	`, serviceName))

	//generate TarsSetHashCode
	t.P(fmt.Sprintf(`//TarsSetHashCode sets the hash code for client calling the server , which is for Message hash code.
	func (obj *%s) TarsSetHashCode(code int64){
        s := _obj.s.(*tars.ServantProxy)
        s.TarsSetHashCode(code)
	}
	`, serviceName))
	t.P()

	//generate the interface
	t.P(fmt.Sprintf("type imp%s interface{", serviceName))
	for _, method := range service.Method {
		t.P(fmt.Sprintf("%s (input %s) (output %s, err error)",
			upperFirstLatter(method.GetName()), t.typeName(method.GetInputType()), t.typeName(method.GetOutputType())))
	}
	t.P("}")
	t.P()

	//gernerate the dispathcer
	t.generateDispatch(service)

	for _, method := range service.Method {
		t.generateClientCode(service, method)
	}
}
func (t *tarsrpc) generateClientCode(service *pb.ServiceDescriptorProto, method *pb.MethodDescriptorProto) {
	methodName := upperFirstLatter(method.GetName())
	serviceName := upperFirstLatter(service.GetName())
	inType := t.typeName(method.GetInputType())
	outType := t.typeName(method.GetOutputType())
	t.P(fmt.Sprintf(`// %s is client rpc method as defined
		func (obj *%s) %s(input %s)(output %s, err error){
			var _status map[string]string
			var _context map[string]string
			var inputMarshal []byte
			inputMarshal, err = proto.Marshal(&input)
			if err != nil {
				return output, err
			}
			resp := new(requestf.ResponsePacket)
			ctx := context.Background()
			err = obj.s.Tars_invoke(ctx, 0, "%s", inputMarshal, _status, _context, resp)
			if err != nil {
				return output, err
			}
			if err = proto.Unmarshal(tools.Int8ToByte(resp.SBuffer), &output); err != nil{
				return output, err
			}
			return output, nil
		}
	`, methodName, serviceName, methodName, inType, outType, method.GetName()))
}
func (t *tarsrpc) generateDispatch(service *pb.ServiceDescriptorProto) {
	serviceName := upperFirstLatter(service.GetName())
	t.P(fmt.Sprintf(`//Dispatch is used to call the user implement of the defined method.
	func (obj *%s) Dispatch(ctx context.Context, val interface{}, req * requestf.RequestPacket, resp *requestf.ResponsePacket, withContext bool)(err error){
		input := tools.Int8ToByte(req.SBuffer)
		var output []byte
		imp := val.(imp%s)
		funcName := req.SFuncName
		switch funcName {
	`, serviceName, serviceName))
	for _, method := range service.Method {
		t.P(fmt.Sprintf(`case "%s":
			inputDefine := %s{}
			if err = proto.Unmarshal(input,&inputDefine); err != nil{
				return err
			}
			res, err := imp.%s(inputDefine)
			if err != nil {
				return err
			}
			output , err = proto.Marshal(&res)
			if err != nil {
				return err
			}
		`, method.GetName(), t.typeName(method.GetInputType()), upperFirstLatter(method.GetName())))
	}
	t.P(`default:
			return fmt.Errorf("func mismatch")
	}
	var status map[string]string
	*resp = requestf.ResponsePacket{
		IVersion:     1,
		CPacketType:  0,
		IRequestId:   req.IRequestId,
		IMessageType: 0,
		IRet:         0,
		SBuffer:      tools.ByteToInt8(output),
		Status:       status,
		SResultDesc:  "",
		Context:      req.Context,
	}
	return nil
}
	`)
	t.P()
}
