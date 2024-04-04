package main

import "fmt"

func main() {
	sum := 1
	// 部分省略的 for 循环
	for sum < 1000 {
		sum += sum
	}
	fmt.Println(sum)
}
