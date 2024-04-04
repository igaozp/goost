package main

import "fmt"

var (
	v1 = vertexInt{1, 2}
	v2 = vertexInt{X: 1}
	v3 = vertexInt{}
	p  = &vertexInt{1, 2}
)

func main() {
	fmt.Println(v1, p, v2, v3)
}
