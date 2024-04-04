package main

import "fmt"

var m map[string]Vertex

func main() {
	// map 在使用之前必须用 make 来创建
	m = make(map[string]Vertex)
	// 值为 nil 的 map 是空的，并且不能对其赋值
	m["Bell Labs"] = Vertex{40.555, -73.23}
	fmt.Println(m["Bell Labs"])
}
