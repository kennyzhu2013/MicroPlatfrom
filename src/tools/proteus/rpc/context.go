package rpc

import (
	"fmt"
	protobuf2 "github.com/kennyzhu/go-os/src/tools/proteus/protobuf"
	"go/types"
	"strings"
)

type context struct {
	implName        string
	constructorName string
	proto           *protobuf2.Package
	pkg             *types.Package
	imports         []string
}

func (c *context) isNameDefined(name string) bool {
	for _, n := range c.pkg.Scope().Names() {
		if n == name {
			return true
		}
	}
	return false
}

func (c *context) findMessage(name string) *protobuf2.Message {
	for _, m := range c.proto.Messages {
		if m.Name == name {
			return m
		}
	}
	return nil
}

func (c *context) findSignature(rpc *protobuf2.RPC) *types.Signature {
	var fn types.Object
	if rpc.Recv != "" {
		recv := c.pkg.Scope().Lookup(rpc.Recv)
		fn, _, _ = types.LookupFieldOrMethod(recv.Type(), true, c.pkg, rpc.Method)
	} else {
		fn = c.pkg.Scope().Lookup(rpc.Method)
	}

	return fn.Type().(*types.Signature)
}

func (c *context) argumentType(rpc *protobuf2.RPC) string {
	signature := c.findSignature(rpc)
	skip := 0
	if rpc.HasCtx {
		skip++
	}
	obj := firstTypeName(skip, signature.Params())
	c.addImport(obj.Pkg().Path())

	return c.objectNameInContext(obj)
}

func (c *context) returnType(rpc *protobuf2.RPC) string {
	signature := c.findSignature(rpc)
	obj := firstTypeName(0, signature.Results())
	c.addImport(obj.Pkg().Path())

	return c.objectNameInContext(obj)
}

// objectNameInContext returns the name of the object prefixed by its package name
// if needed
func (c *context) objectNameInContext(obj types.Object) string {
	if c.pkg.Path() == obj.Pkg().Path() {
		return obj.Name()
	} else {
		return fmt.Sprintf("%s.%s", obj.Pkg().Name(), obj.Name())
	}
}

func firstTypeName(skip int, tuple *types.Tuple) types.Object {
	t := tuple.At(skip).Type()
	if inner, ok := t.(*types.Pointer); ok {
		t = inner.Elem()
	}
	return t.(*types.Named).Obj()
}

func (c *context) pkgPath() string {
	if c.pkg != nil {
		return c.pkg.Path()
	}

	return ""
}

func (c *context) addImport(path string) {
	if path == c.pkgPath() {
		return
	}

	for _, i := range c.imports {
		if i == path {
			return
		}
	}
	c.imports = append(c.imports, path)
}

//
func serviceImplName(pkg *protobuf2.Package) string {
	n := pkg.ServiceName()
	return strings.ToLower(string(n[0])) + n[1:] + "Server"
}


func constructorName(pkg *protobuf2.Package) string {
	return fmt.Sprintf("New%sServer", pkg.ServiceName())
}
