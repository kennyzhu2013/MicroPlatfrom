// +build linux,amd64

/*
@Time : 2019/6/19 10:46
@Author : kenny zhu
@File : test_readonly
@Software: GoLand
@Others: todo: 数据进一步控制:
身份验证：用 runtime.Caller 验证调用堆栈，仅允许指定函数调用。
内存锁定：用 syscall.Mlock 将数据锁定在物理内存页，禁止交换到硬盘。
*/
package main

import (
	"syscall"
	"unsafe"
)

type ProtectVar []byte

func newProtectVar(size int) (ProtectVar, error)  {
	b, err := syscall.Mmap(0, 0, size, syscall.PROT_READ | syscall.PROT_WRITE, syscall.MAP_ANON | syscall.MAP_PRIVATE)

	if err != nil {
		return nil, err
	}

	return ProtectVar(b), nil
}

func (p ProtectVar) Free() error  {
	return syscall.Munmap([]byte(p))
}

func (p ProtectVar) Readonly() error  {
	return syscall.Mprotect([]byte(p), syscall.PROT_READ)
}

func (p ProtectVar) ReadWrite() error  {
	return syscall.Mprotect([]byte(p), syscall.PROT_READ | syscall.PROT_WRITE)
}

func (p ProtectVar) Pointer() unsafe.Pointer  {
	return unsafe.Pointer(&p[0])
}

func main()  {
	pv, err := newProtectVar(8)
	if err != nil {
		println("error:", err)
	}

	defer pv.Free()

	p := (*int)(pv.Pointer())
	*p = 100
	println("1:", *p)

	pv.Readonly()
	println("2:", *p)

	pv.ReadWrite()
	*p += 100
	println("3:", *p)

	pv.Readonly()
	*p++
}