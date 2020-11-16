package proteus

import (
	protobuf2 "github.com/kennyzhu/go-os/src/tools/proteus/protobuf"
	resolver2 "github.com/kennyzhu/go-os/src/tools/proteus/resolver"
	rpc2 "github.com/kennyzhu/go-os/src/tools/proteus/rpc"
	scanner2 "github.com/kennyzhu/go-os/src/tools/proteus/scanner"
)

// Options are all the available options to configure proto generation.
type Options struct {
	BasePath string
	Packages []string
}

// store every proto file names  for every package...
// Todo:
type PackageFiles struct {
	Files map[string][]string // key = package name, value = files...
}

type generator func(*scanner2.Package, *protobuf2.Package) error

// generate for rpc generator...
// can't find import: "github.com/golang/protobuf/proto"
// go get github.com/golang/protobuf/proto
func transformToProtobuf(packages []string, generate generator) error {
	scanner, err := scanner2.New(packages...)
	if err != nil {
		return err
	}

	pkgs, err := scanner.Scan()
	if err != nil {
		return err
	}

	r := resolver2.New()
	r.Resolve(pkgs)

	t := protobuf2.NewTransformer()
	t.SetStructSet(createStructTypeSet(pkgs))
	t.SetEnumSet(createEnumTypeSet(pkgs))
	for _, p := range pkgs {
		pkg := t.Transform(p)
		if err := generate(p, pkg); err != nil {
			return err
		}
	}

	return nil
}

func createStructTypeSet(pkgs []*scanner2.Package) protobuf2.TypeSet {
	ts := protobuf2.NewTypeSet()
	for _, p := range pkgs {
		for _, s := range p.Structs {
			ts.Add(p.Path, s.Name)
		}
	}
	return ts
}

func createEnumTypeSet(pkgs []*scanner2.Package) protobuf2.TypeSet {
	ts := protobuf2.NewTypeSet()
	for _, p := range pkgs {
		for _, e := range p.Enums {
			ts.Add(p.Path, e.Name)
		}
	}
	return ts
}

// GenerateProtos generates proto files for the given options.
// BasePath = -f path...
func GenerateProtos(options Options) error {
	g := protobuf2.NewGenerator(options.BasePath)
	return transformToProtobuf(options.Packages, func(_ *scanner2.Package, pkg *protobuf2.Package) error {
		// return g.GeneratorDb(pkg)
		return g.Generator( pkg )
	})
}

// GenerateRPCServer generates the gRPC server implementation of the given
// packages.
func GenerateRPCServer(packages []string) error {
	// g := rpc.NewGenerator()
	g := rpc2.NewMicroGenerator() // change to micro service ....
	return transformToProtobuf(packages, func(p *scanner2.Package, pkg *protobuf2.Package) error {
		return g.Generate(pkg, p.Path)
	})
}
