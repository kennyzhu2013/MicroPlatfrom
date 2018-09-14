package protobuf

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/suite"
)

func TestGenerator(t *testing.T) {
	suite.Run(t, new(GenSuite))
}

type GenSuite struct {
	suite.Suite
	path string
	g    *Generator
	buf  *bytes.Buffer
}

func (s *GenSuite) SetupTest() {
	var err error
	s.path, err = ioutil.TempDir("", "proteus")
	s.Nil(err)

	s.buf = bytes.NewBuffer(nil)
	s.g = NewGenerator(s.path)
}

func (s *GenSuite) TearDownTest() {
	s.Nil(os.RemoveAll(s.path))
}

const expectedOptionsIndented = `	option bar = true;
	option foo = "bar";
`

const expectedOptions = `option bar = true;
option foo = "bar";
`

func (s *GenSuite) TestWriteOptions() {
	options := Options{
		"foo": NewStringValue("bar"),
		"bar": NewLiteralValue("true"),
	}

	writeOptions(s.buf, options, true)
	s.Equal(expectedOptionsIndented, s.buf.String())

	s.buf.Reset()
	writeOptions(s.buf, options, false)
	s.Equal(expectedOptions, s.buf.String())
}

const expectedEnum = `// Possible pony races
enum PonyRace {
	option is_cute = true;
	// Pink cutie
	PINK_CUTIE = 0;
	RED_FURY = 1 [bar = "baz", foo = true];
}
`

var mockEnum = &Enum{
	Docs: []string{"Possible pony races"},
	Name: "PonyRace",
	Options: Options{
		"is_cute": NewLiteralValue("true"),
	},
	Values: []*EnumValue{
		{
			Docs:  []string{"Pink cutie"},
			Name:  "PINK_CUTIE",
			Value: 0,
		},
		{
			Name:  "RED_FURY",
			Value: 1,
			Options: Options{
				"foo": NewLiteralValue("true"),
				"bar": NewStringValue("baz"),
			},
		},
	},
}

func (s *GenSuite) TestWriteEnum() {
	writeEnum(s.buf, mockEnum)
	s.Equal(expectedEnum, s.buf.String())
}

const expectedMsg = `// Pony is so fancy
// and so fluffy
message Pony {
	option is_cute = true;
	reserved 5, 6;
	// Name of the pony
	string name = 1 [bar = "baz", foo = true];
	// Time the pony was born
	google.protobuf.Timestamp born_at = 2;
	// Pony's race
	foo.bar.PonyRace race = 3;
	// All the fancy nicknames the pony has
	repeated string nick_names = 4;
}
`

var mockMsg = &Message{
	Docs: []string{
		"Pony is so fancy",
		"and so fluffy",
	},
	Name: "Pony",
	Options: Options{
		"is_cute": NewLiteralValue("true"),
	},
	Reserved: []uint{5, 6},
	Fields: []*Field{
		{
			Docs: []string{
				"Name of the pony",
			},
			Name: "name",
			Type: NewBasic("string"),
			Pos:  1,
			Options: Options{
				"foo": NewLiteralValue("true"),
				"bar": NewStringValue("baz"),
			},
		},
		{
			Docs: []string{"Time the pony was born"},
			Name: "born_at",
			Type: NewNamed("google.protobuf", "Timestamp"),
			Pos:  2,
		},
		{
			Docs: []string{"Pony's race"},
			Name: "race",
			Type: NewNamed("foo.bar", "PonyRace"),
			Pos:  3,
		},
		{
			Docs:     []string{"All the fancy nicknames the pony has"},
			Name:     "nick_names",
			Repeated: true,
			Type:     NewBasic("string"),
			Pos:      4,
		},
	},
}

func (s *GenSuite) TestWriteMessage() {
	writeMessage(s.buf, mockMsg)
	s.Equal(expectedMsg, s.buf.String())
}

var mockRpcs = []*RPC{
	{
		Docs:   []string{"DoFoo does a lot of Foo"},
		Name:   "DoFoo",
		Input:  NewNamed("foo.bar", "DoFooRequest"),
		Output: NewNamed("foo.bar", "DoFooResponse"),
	},
	{
		Docs:   []string{"DoBar does a lot of Bar"},
		Name:   "DoBar",
		Input:  NewNamed("foo.bar", "DoBarRequest"),
		Output: NewNamed("foo.bar", "DoBarResponse"),
	},
}

const expectedService = `service BarService {
	// DoFoo does a lot of Foo
	rpc DoFoo (foo.bar.DoFooRequest) returns (foo.bar.DoFooResponse);
	// DoBar does a lot of Bar
	rpc DoBar (foo.bar.DoBarRequest) returns (foo.bar.DoBarResponse);
}

`

func (s *GenSuite) TestWriteService() {
	writeService(s.buf, &Package{
		Name: "foo.bar",
		RPCs: mockRpcs,
	})
	s.Equal(expectedService, s.buf.String())
}

var expectedProto = fmt.Sprintf(`syntax = "proto3";
package foo.bar;

import "google/protobuf/timestamp.proto";

option foo = true;

%s
%s
%s`, expectedMsg, expectedEnum, expectedService)

func (s *GenSuite) TestGenerate() {
	err := s.g.Generate(&Package{
		Name:     "foo.bar",
		Imports:  []string{"google/protobuf/timestamp.proto"},
		Messages: []*Message{mockMsg},
		Enums:    []*Enum{mockEnum},
		Options:  Options{"foo": NewLiteralValue("true")},
		RPCs:     mockRpcs,
	})
	s.Nil(err)

	bytes, err := ioutil.ReadFile(filepath.Join(s.path, "generated.proto"))
	s.Nil(err)

	s.Equal(expectedProto, string(bytes))
}
