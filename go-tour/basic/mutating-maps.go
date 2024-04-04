package main

import "fmt"

func main() {
	m := make(map[string]int)
	m["Answer"] = 42
	fmt.Println("The value:", m["Answer"])
	// 修改元素
	m["Answer"] = 48
	fmt.Println("The value:", m["Answer"])
	// 删除元素
	delete(m, "Answer")
	fmt.Println("The value:", m["Answer"])
	// 通过双赋值检测某个键是否存在
	v, ok := m["Answer"]
	fmt.Println("The value:", v, "Present?", ok)
}
