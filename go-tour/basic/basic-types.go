package main

import (
	"fmt"
	"math/cmplx"
)

/*
*

	basic types
	bool
	string
	int  int8  int16  int32  int64
	uint uint8 uint16 uint32 uint64 uintptr
	byte // uint8 的别名
	rune // int32 的别名, 代表一个Unicode码
	float32 float64
	complex64 complex128
*/
var (
	ToBe   bool       = false
	MaxInt uint64     = 1<<64 - 1
	z      complex128 = cmplx.Sqrt(-5 + 12i)
)

func main() {
	const f = "qwerty"
	fmt.Println(f, ToBe, ToBe)
	fmt.Println(f, MaxInt, MaxInt)
	fmt.Println(f, z, z)
}
