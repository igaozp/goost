package main

import "fmt"

type vertexInt struct {
	X, Y int
}

func main() {
	v := vertexInt{1, 2}
	// 结构体字段使用 "." 来访问
	v.X = 4
	fmt.Println(v.X)
}
