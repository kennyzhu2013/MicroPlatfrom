/*
@Time : 2019/6/18 15:36
@Author : kenny zhu
@File : test_reflect
@Software: GoLand
@Others:
*/
package main

import (
	"reflect"
	"testing"
	"unsafe"
)

func incX(d interface{}) int64  {
	v := reflect.ValueOf(d).Elem()
	f := v.FieldByName("X")

	X := f.Int()
	X++
	f.SetInt(X)

	return X
}

var offset uintptr = 0xFFFF //
var cache = map[*uintptr]map[string]uintptr{}

func unsafeCacheIncX(d interface{}) int64 {
	itab := *(**uintptr)(unsafe.Pointer(&d))

	m, OK := cache[itab]
	if !OK {
		m = make(map[string]uintptr)
		cache[itab] = m
	}

	// get x filed
	offset, OK := m["X"]
	if !OK {
		t := reflect.TypeOf(d).Elem()
		x, _ := t.FieldByName("X")
		offset = x.Offset
		m["X"] = offset
	}

	p := (*[2]uintptr)(unsafe.Pointer(&d))
	px := (*int64)(unsafe.Pointer( p[1] + offset ))
	*px++

	return *px
}

func unsafeIncX(d interface{}) int64  {
	if offset == 0xFFFF {
		t := reflect.TypeOf(d).Elem()
		x,_ := t.FieldByName("X")
		offset = x.Offset
	}

	p := (*[2]uintptr)(unsafe.Pointer(&d))
	// print(unsafe.Sizeof(p[0]))
	// print("\n")
	// print(p)
	// print("\n")
	// print(p[0])
	// print("\n")
	// print(p[1])
	// print("\n")
	// print(&d)
	// print("\n")
	// p[1]是基址?...
	px := (*int64)(unsafe.Pointer( p[1] + offset ))
	*px++

	return *px
}

func BenchmarkIncX(b *testing.B)  {
	d := struct {
		X int
	}{100}

	for i := 0; i < b.N; i++ {
		incX(&d)
	}
}

func BenchmarkUnsafeIncX(b *testing.B)  {
	d := struct {
		X int
	}{100}

	for i := 0; i < b.N; i++ {
		unsafeIncX(&d)
	}
}

func BenchmarkCacheUnsafeIncX(b *testing.B)  {
	d := struct {
		X int
	}{100}

	for i := 0; i < b.N; i++ {
		unsafeCacheIncX(&d)
	}
}

func main()  {
	d := struct {
		X int
	}{100}

	print(unsafe.Sizeof(&d))
	print("\n")

	print( unsafeIncX(&d) )
}