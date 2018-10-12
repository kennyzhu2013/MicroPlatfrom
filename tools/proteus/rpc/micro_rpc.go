/*
@Time : 2018/10/12 10:58 
@Author : kenny zhu
@File : micro_rpc.go
@Software: GoLand
@Others:
*/
package rpc


import (
	"fmt"
	"go/ast"
	"go/printer"
	"go/token"
	"os"
	"path/filepath"
	"strings"

	"github.com/kennyzhu/go-os/tools/proteus/protobuf"
	"github.com/kennyzhu/go-os/tools/proteus/report"

	"gopkg.in/src-d/go-parse-utils.v1"
)

// Generator generates implementations of an RPC server for a package.
// It will declare a new type with the following name: {serviceName}Server
// with the first letter in lower case. In case the type already exists in the
// package, it will not be generated again. The purpose of that is that you can
// customize the type, even though the methods are automatically generated.
// Same happens with the constructor. If it does not exist, a function named
// new{ServiceName}Server with no parameters and a single result of the type
// {serviceName}Server will be generated. It can be defined, to avoid a default
// implementation, but the same signature is required, as it is used for
// registering the server implementation.
//
// So, if you have a service named FooService, you can implement
// `fooServiceServer` and `func newFooServiceServer() *fooServiceServer`.
//
// All generated methods will use as receiver a field in the server
// implementation with the same name as the type of the receiver.
// For example, the method generated for `func (*Foo) Bar()` will be require
// that our `fooServiceServer` had a field with that name.
//
// 	type fooServiceServer struct {
//		Foo *Foo
//	}
//
// For every method you want to generate an RPC method for, you have to
// implement its receiver by yourself in the server implementation type and the
// constructor.
//
// A single file per package will be generated containing all the RPC methods.
// The file will be written to the package path and it will be named
// "server.proteus.go"
type MicroGenerator struct {
	importer *parseutil.Importer
}

// NewGenerator creates a new Generator.
func NewMicroGenerator() *MicroGenerator {
	return &MicroGenerator{parseutil.NewImporter()}
}

// Generate creates a new file in the package at the given path and implements
// the server according to the given proto package.
// generate micro service...
func (g *MicroGenerator) Generate(proto *protobuf.Package, path string) error {
	if len(proto.RPCs) == 0 {
		report.Warn("no micro RPCs in the given proto file, not generating anything")
		return nil
	}

	// insert import, exclude files..
	pkg, err := g.importer.ImportWithFilters(
		path,
		parseutil.FileFilters{
			func(pkg, file string, typ parseutil.FileType) bool {
				return !strings.HasSuffix(file, ".pb.go")
			},
			func(pkg, file string, typ parseutil.FileType) bool {
				return !strings.HasSuffix(file, ".proteus.go")
			},
			func(pkg, file string, typ parseutil.FileType) bool {
				return !strings.HasSuffix(file, ".micro.go")
			},
		},
	)
	if err != nil {
		return err
	}

	// find context
	ctx := &context{
		implName:        serviceImplName(proto),
		constructorName: constructorName(proto),
		proto:           proto,
		pkg:             pkg,
	}

	var decls []ast.Decl
	if !ctx.isNameDefined(ctx.implName) {
		// type struct
		decls = append(decls, g.declImplType(ctx.implName))
	}

	if !ctx.isNameDefined(ctx.constructorName) {
		// constructor
		report.Warn("constructor %s for service %s is not implemented", ctx.implName, ctx.constructorName)
		decls = append(decls, g.declConstructor(ctx.implName, ctx.constructorName))
	}

	for _, rpc := range proto.RPCs {
		// function
		decls = append(decls, g.declMethod(ctx, rpc))
	}

	return g.writeFile(g.buildFile(ctx, decls), path)
}

func (g *MicroGenerator) declImplType(implName string) ast.Decl {
	return &ast.GenDecl{
		Tok: token.TYPE,
		Specs: []ast.Spec{
			&ast.TypeSpec{
				Name: ast.NewIdent(implName),
				Type: &ast.StructType{
					Fields: fields(),
				},
			},
		},
	}
}

func (g *MicroGenerator) declConstructor(implName, constructorName string) ast.Decl {
	return &ast.FuncDecl{
		Name: ast.NewIdent(constructorName),
		Type: &ast.FuncType{
			Params: fields(),
			Results: fields(&ast.Field{
				Type: ptr(ast.NewIdent(implName)),
			}),
		},
		Body: &ast.BlockStmt{
			List: []ast.Stmt{
				&ast.ReturnStmt{
					Results: []ast.Expr{
						&ast.UnaryExpr{
							Op: token.AND,
							X: &ast.CompositeLit{
								Type: ast.NewIdent(implName),
							},
						},
					},
				},
			},
		},
	}
}

// create function type for micro..
func (g *MicroGenerator) genMethodType(ctx *context, rpc *protobuf.RPC) *ast.FuncType {
	var in, out string

	if isGenerated(rpc.Input) {
		in = typeName(rpc.Input)
	} else {
		in = ctx.argumentType(rpc)
	}

	if isGenerated(rpc.Output) {
		out = typeName(rpc.Output)
	} else {
		out = ctx.returnType(rpc)
	}

	// func (s *Say) GetAlphaTime(ctx context.Context, in *GetAlphaTimeRequest, result *MyTime) error..
	return &ast.FuncType{
		Params: fields(
			field("ctx", ast.NewIdent("context.Context")),
			field("in", ptr(ast.NewIdent(in))),

			// add
			field("result", ptr(ast.NewIdent(out))),
		),
		Results: fields(
			// field("result", ptr(ast.NewIdent(out))),
			field("err", ast.NewIdent("error")),
		),
	}
}

// function call part..
func (g *MicroGenerator) genMethodCall(ctx *context, rpc *protobuf.RPC) ast.Expr {
	call := &ast.CallExpr{Fun: ast.NewIdent(rpc.Method)}
	if rpc.Recv != "" {
		call.Fun = ast.NewIdent(fmt.Sprintf("s.%s.%s", rpc.Recv, rpc.Method))
	}

	if rpc.HasCtx {
		call.Args = append(call.Args, ast.NewIdent("ctx"))
	}
	if rpc.IsVariadic {
		call.Ellipsis = token.Pos(1)
	}

	if !isGenerated(rpc.Input) {
		var in ast.Expr = ast.NewIdent("in")
		if !rpc.Input.IsNullable() {
			in = &ast.StarExpr{
				X: in,
			}
		}
		call.Args = append(call.Args, in)
	} else {
		msg := ctx.findMessage(typeName(rpc.Input))
		for i := range msg.Fields {
			call.Args = append(call.Args, ast.NewIdent(fmt.Sprintf(
				"in.Arg%d", i+1,
			)))
		}
	}

	return call
}

// result = new(MyDuration), change to temp = new(MyDuration)
func (g *MicroGenerator) genBaseMethodBody(methodType *ast.FuncType) *ast.BlockStmt {
	return &ast.BlockStmt{
		List: []ast.Stmt{
			&ast.AssignStmt{
				Tok: token.DEFINE,
				Lhs: []ast.Expr{ast.NewIdent("tempResult")},
				Rhs: []ast.Expr{
					&ast.CallExpr{
						Fun: ast.NewIdent("new"),
						Args: []ast.Expr{
							methodType.Params.List[2].Type.(*ast.StarExpr).X,
						},
					},
				},
			},
		},
	}
}

func (g *MicroGenerator) genResultOutputMethodBody(methodType *ast.FuncType) *ast.AssignStmt {
	return &ast.AssignStmt{
				Tok: token.ASSIGN,
				// Lhs: []ast.Expr{ast.NewIdent("result")},
				Lhs: []ast.Expr{ast.NewIdent("*result")},
				Rhs: []ast.Expr{
					&ast.UnaryExpr{
						Op: token.MUL,
						X:  ast.NewIdent("tempResult"),
					},
				},
			}
}

// micro function body
func (g *MicroGenerator) genMethodBody(ctx *context, rpc *protobuf.RPC, typ *ast.FuncType) *ast.BlockStmt {
	// whether the named type was generated by proteus ..
	if !isGenerated(rpc.Output) {
		return g.genMethodBodyForNotGeneratedOutput(ctx, rpc, typ)
	} else {
		return g.genMethodBodyForGeneratedOutput(ctx, rpc, typ)
	}
}

func (g *MicroGenerator) genMethodBodyAssignmentsForGeneratedOutput(ctx *context, rpc *protobuf.RPC, msg *protobuf.Message) (lhs []ast.Expr) {
	for i, f := range msg.Fields {
		if f == nil {
			lhs = append(lhs, ast.NewIdent("_"))
		} else {
			lhs = append(lhs, ast.NewIdent(fmt.Sprintf(
				"tempResult.Result%d", i+1,
			)))
		}
	}
	return
}

// 返回值需要封装的情况:...
func (g *MicroGenerator) genMethodBodyForGeneratedOutput(ctx *context, rpc *protobuf.RPC, typ *ast.FuncType) *ast.BlockStmt {
	body := g.genBaseMethodBody(typ)
	methodCall := g.genMethodCall(ctx, rpc)
	call := &ast.AssignStmt{
		Tok: token.ASSIGN,
		Rhs: []ast.Expr{methodCall},
	}

	// message struct of output
	msg := ctx.findMessage(typeName(rpc.Output))

	if len(msg.Fields) == 0 && !rpc.HasError {
		return emptyBodyForMethodCall(body, methodCall)
	} else if len(msg.Fields) == 0 {
		body.List = nil
	}

	body.List = append(body.List, call)
	lhs := g.genMethodBodyAssignmentsForGeneratedOutput(ctx, rpc, msg)
	call.Lhs = append(call.Lhs, lhs...)

	if rpc.HasError {
		call.Lhs = append(call.Lhs, ast.NewIdent("err"))
	}

	// add *result = *tempResult
	resultOutput := g.genResultOutputMethodBody(typ)
	body.List = append(body.List, resultOutput)

	body.List = append(body.List, new(ast.ReturnStmt))
	return body
}

// func with no output with aux := RandomCategory()...
func (g *MicroGenerator) genMethodBodyForNotGeneratedOutput(ctx *context, rpc *protobuf.RPC, typ *ast.FuncType) *ast.BlockStmt {
	body := g.genBaseMethodBody(typ) // temp = new(MyDuration)
	methodCall := g.genMethodCall(ctx, rpc) // func call part. like : GetDurationForLengthCtx(ctx, in.Arg1)..
	call := &ast.AssignStmt{Tok: token.ASSIGN}

	// whether rpc.Output is return *ptr value
	needToAddressOutput := !isGenerated(rpc.Output) && !rpc.Output.IsNullable()

	// Specific code ,if address ..
	if needToAddressOutput {
		call.Lhs = append(call.Lhs, ast.NewIdent("aux"))
		call.Tok = token.DEFINE
	} else {
		call.Lhs = append(call.Lhs, ast.NewIdent("tempResult"))
	}

	// have error return...
	if rpc.HasError {
		call.Lhs = append(call.Lhs, ast.NewIdent("err"))
	}

	//
	call.Rhs = append(call.Rhs, methodCall)
	body.List = append(body.List, call)

	// result = &aux
	if needToAddressOutput {
		body.List = append(body.List, &ast.AssignStmt{
			Tok: token.ASSIGN,
			Lhs: []ast.Expr{ast.NewIdent("tempResult")},
			Rhs: []ast.Expr{
				&ast.UnaryExpr{
					Op: token.AND,
					X:  ast.NewIdent("aux"),
				},
			},
		})
	}

	// add *result = *tempResult
	resultOutput := g.genResultOutputMethodBody(typ)
	body.List = append(body.List, resultOutput)

	body.List = append(body.List, new(ast.ReturnStmt))
	return body
}

//
func (g *MicroGenerator) declMethod(ctx *context, rpc *protobuf.RPC) ast.Decl {
	typ := g.genMethodType(ctx, rpc)
	return &ast.FuncDecl{
		Recv: fields(field("s", ptr(ast.NewIdent(ctx.implName)))), // (s *exampleServiceHandler)
		Name: ast.NewIdent(rpc.Name),
		Type: typ,
		Body: g.genMethodBody(ctx, rpc, typ),
	}
}

func (g *MicroGenerator) buildFile(ctx *context, decls []ast.Decl) *ast.File {
	f := &ast.File{
		Name: ast.NewIdent(ctx.pkg.Name()),
	}

	var specs = []ast.Spec{newImport("context")}
	for _, i := range ctx.imports {
		specs = append(specs, newImport(i))
	}

	f.Decls = append(f.Decls, &ast.GenDecl{
		Tok:    token.IMPORT,
		Lparen: token.Pos(1),
		Specs:  specs,
	})
	f.Decls = append(f.Decls, decls...)

	return f
}

func (g *MicroGenerator) writeFile(file *ast.File, path string) error {
	fileName := filepath.Join(goSrc, path, "server.micro.go")
	f, err := os.Create(fileName)
	if err != nil {
		return err
	}
	defer f.Close()

	return printer.Fprint(f, token.NewFileSet(), file)
}
