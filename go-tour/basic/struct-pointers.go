package main

import "fmt"

func main() {
	v := vertexInt{1, 2}
	// 结构体字段可以通过结构体指针来访问
	p := &v
	p.X = 1e9
	fmt.Println(v)
}
