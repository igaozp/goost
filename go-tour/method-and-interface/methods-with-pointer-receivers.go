package main

import (
	"fmt"
)

// 方法可以与命名类型或命名类型的指针关联
func (v *Vertex) Scale(f float64) {
	v.X = v.X * f
	v.Y = v.Y * f
}

func main() {
	v := &Vertex{3, 4}
	fmt.Printf("Before scalling: %+v, Abs: %v\n", v, v.Abs())
	v.Scale(5)
	fmt.Printf("After scalling: %+v, Abs: %v\n", v, v.Abs())
}
