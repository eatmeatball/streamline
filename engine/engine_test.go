package engine

import (
	"fmt"
	"runtime"
	"testing"
)

func TestVt2(t *testing.T) {
	a := new(vt2)
	a.v = 1
	a.v = "1231"
	a.v = 123.123
	fmt.Println(a)
}

func TestVt(t *testing.T) {
	var m runtime.MemStats
	runtime.ReadMemStats(&m)
	fmt.Println(m.Alloc / 8)
	runtime.ReadMemStats(&m)
	fmt.Println("print")
	fmt.Println(m.Alloc / 8)
	runtime.ReadMemStats(&m)
	fmt.Println(m.Alloc / 8)
	fmt.Println("new")
	v := new(vt)
	runtime.ReadMemStats(&m)
	fmt.Println(m.Alloc / 8)
	fmt.Println(v.vInt)
	runtime.ReadMemStats(&m)
	fmt.Println(m.Alloc / 8)
	fmt.Println("set vint")
	vint := 1
	vints := [10]int{}
	v.vInt = &vint
	v.vIntS = &vints
	runtime.ReadMemStats(&m)
	fmt.Println(m.Alloc / 8)
}
