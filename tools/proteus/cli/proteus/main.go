package main

import (
	"errors"
	"fmt"
	"os"
	"os/exec"
	"path/filepath"

	"github.com/kennyzhu/go-os/tools/proteus"
	"github.com/kennyzhu/go-os/tools/proteus/protobuf"
	"github.com/kennyzhu/go-os/tools/proteus/report"

	// Package cli provides a minimal framework for creating and organizing command line Go applications.
	"gopkg.in/urfave/cli.v1"
)

var (
	packages cli.StringSlice
	path     string
	verbose  bool
)

func main() {
	app := cli.NewApp()
	app.Name = "proteus"
	app.Description = "Proteus generates code and protobuffer 3 proto files while keeping your Go source code as the source of truth."
	app.Version = "1.0.0"

	baseFlags := []cli.Flag{
		cli.StringSliceFlag{
			Name:  "pkg, p",
			Usage: "Use `PACKAGE` as input for the generation. You can use this flag multiple times to specify more than one package.",
			Value: &packages,
		},
		cli.BoolFlag{
			Name:        "verbose",
			Usage:       "Print all warnings and info messages.",
			Destination: &verbose,
		},
	}

	folderFlag := cli.StringFlag{
		Name:        "folder, f",
		Usage:       "All generated .proto files will be written to `FOLDER`.",
		Destination: &path,
	}

	app.Flags = append(baseFlags, folderFlag)
	app.Commands = []cli.Command{
		{
			Name:        "proto",
			Description: "Generates .proto files from your Go source code.",
			Usage:       "Generates .proto files from Go packages",
			Action:      initCmd(genProtos),
			Flags:       append(baseFlags, folderFlag),
		},
		{
			Name:        "rpc",
			Description: "Generates the gRPC implementation of the gRPC server interface defined by your Go source code.",
			Usage:       "Generates gRPC server implementation",
			Action:      initCmd(genRPCServer),
			Flags:       baseFlags,
		},
		{  // add http json api...
			Name:        "http",
			Description: "Generates the json implementation of the http server interface defined by your Go source code.",
			Usage:       "Generates http server implementation",
			Action:      initCmd(genHttpServer),
			Flags:       baseFlags,
		},
	}
	app.Action = initCmd(genAll)

	if err := app.Run(os.Args); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

type action func(c *cli.Context) error

func initCmd(next action) func(c *cli.Context) error {
	return func(c *cli.Context) error {
		if len(packages) == 0 {
			return errors.New("no package provided, there is nothing to generate")
		}

		if !verbose {
			report.Silent()
		}

		return next(c)
	}
}

func genProtos(c *cli.Context) error {
	if path == "" {
		return errors.New("destination path cannot be empty")
	}

	if err := checkFolder(path); err != nil {
		return err
	}

	// generate generated.proto...
	return proteus.GenerateProtos(proteus.Options{
		BasePath: path,
		Packages: packages,
	})
}

func genRPCServer(c *cli.Context) error {
	return proteus.GenerateRPCServer(packages)
}

func genHttpServer(c *cli.Context) error {
	return nil
}

var (
	goSrc       = filepath.Join(os.Getenv("GOPATH"), "src")
	protobufSrc = filepath.Join(goSrc, "github.com", "gogo", "protobuf")
)

func genAll(c *cli.Context) error {
	protocPath, err := exec.LookPath("protoc")
	if err != nil {
		return fmt.Errorf("protoc is not installed: %s", err)
	}
	protocPath = protocPath

	if err := checkFolder(protobufSrc); err != nil {
		return fmt.Errorf("github.com/gogo/protobuf is not installed")
	}

	if err := genProtos(c); err != nil {
		return err
	}

	// for files, to proto c...:generate pb.go and micro.go
	// 或手动执行 : protoc --proto_path=$GOPATH/src:. --micro_out=. --go_out=. usercenter/t_*.proto
	// or go-kit -- proto file to generate
	// protoc --proto_path=$GOPATH/src::/opt/gopath/src/github.com/gogo/protobuf/protobuf:. --micro_out=. ./protos/gopkg.in/src-d/proteus.v1/example/generated.proto.
	// here to generator pb.go, 已经成功执行....
	// 注意: 同一个包里的公有结构无法引用，因此公有结构要单独在一个包里被应用....
	for _, p := range packages {
		outPath := goSrc
		proto := filepath.Join(path, p, "generated.proto") // -f path + packagename + "*.proto"

		if err := protocExec(protocPath, p, outPath, proto); err != nil {
			return fmt.Errorf("error generating Go files from %q: %s", proto, err)
		}

		matches, err := filepath.Glob(filepath.Join(path, p, "*.pb.go"))
		if err != nil {
			return fmt.Errorf("error moving Go files")
		}

		moveToDir := filepath.Join(outPath, p)
		for _, s := range matches {
			mv(s, moveToDir)
		}
	}

	// generator json gateway for gin using https://github.com/grpc-ecosystem/grpc-gateway...
	// Todo: ...
	genHttpServer(c)

	// generate grpc server handler for micro, to modify  ...
	// generate micro service example for micro...
	return genRPCServer(c)
}

// create fast pb.go,,
func protocExec(protocPath, pkg, outPath, protoFile string) error {
	protocArgs := fmt.Sprintf(
		"--proto_path=%s:%s:%s:%s:.",
		goSrc,
		path,
		filepath.Join(protobufSrc, "protobuf"),
		filepath.Join(path, pkg),
	)

	report.Info("executing protoc: %s %s %s %s %s", protocPath, protocArgs, genAllGoFastOutOption(outPath), genAllMicroOutOption(outPath), protoFile)

	// 不能通配，为啥...
	/*
	cmd := exec.Command(
		protocPath,
		protocArgs,
		genAllGoFastOutOption(outPath),
		protoFile,
	)*/

	// 通配符问题...
	commandLine := "`" + protocPath + " " + protocArgs + " " + genAllGoFastOutOption(outPath) + " " + genAllMicroOutOption(outPath) + " " + protoFile + "`"
	cmd := exec.Command(
		"/bin/bash", "-c",
		commandLine,
	)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	return cmd.Run()
}

func mv(from, to string) error {
	cmd := exec.Command("mv", from, to)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	return cmd.Run()
}

func genAllGoFastOutOption(outPath string) string {
	str := "--gofast_out=plugins=grpc"
	importMappings := protobuf.DefaultMappings.ToGoOutPath()

	if importMappings != "" {
		str += fmt.Sprintf(",%s", importMappings)
	}

	// output path..
	str += fmt.Sprintf(":%s", outPath)

	return str
}

func genAllMicroOutOption(outPath string) string {
	str := "--micro_out="
	importMappings := protobuf.DefaultMappings.ToGoOutPath()

	if importMappings != "" {
		str += fmt.Sprintf(",%s", importMappings)
	}

	// output path..
	str += fmt.Sprintf(":%s", outPath)

	return str
}

func checkFolder(p string) error {
	fi, err := os.Stat(p)
	switch {
	case os.IsNotExist(err):
		return errors.New("folder does not exist, please create it first")
	case err != nil:
		return err
	case !fi.IsDir():
		return fmt.Errorf("folder is not directory: %s", p)
	}
	return nil
}
