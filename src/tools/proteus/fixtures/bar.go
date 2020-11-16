package foo

import (
	subpkg2 "github.com/kennyzhu/go-os/src/tools/proteus/fixtures/subpkg"
)

// Bar ...
//proteus:generate
type Bar struct {
	Bar uint64
	Baz Baz
}

// Baz ...
//proteus:generate
type Baz byte

// Saz ...
//proteus:generate
type Saz struct {
	Point subpkg2.Point
	Foo   float64
}

const (
	// ABaz ...
	ABaz Baz = iota
	// BBaz ...
	BBaz
	// CBaz ...
	CBaz
	// DBaz ...
	DBaz
)
