package main

import (
	"fmt"
	"math"
)

// "方法接收者"出现在 func 关键字和方法名之间的参数中
// 即可以调用该方法的的变量类型
func (v *Vertex) Sqrt() float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

func main() {
	v := &Vertex{3, 4}
	fmt.Println(v.Sqrt())
}
