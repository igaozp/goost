package main

import (
	"fmt"
	"math"
)

// 接口类型是由一组方法定义的集合
type Abser interface {
	Abs() float64
}

func main() {
	// 接口类型的值可以存放实现这些方法的任何值
	var a Abser
	f := MyFloat(-math.Sqrt2)
	v := Vertex{3, 4}

	a = f  // a MyFloat 实现了 Abser
	a = &v // a *Vertex 实现了 Abser

	fmt.Println(a.Abs())
}

type MyFloat float64

// 可以对包中的"任意"类型定义任意方法
func (f MyFloat) Abs() float64 {
	if f < 0 {
		return float64(-f)
	}
	return float64(f)
}

type Vertex struct {
	X, Y float64
}

func (v *Vertex) Abs() float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}
