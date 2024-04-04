package main

import "fmt"

func main() {
	pow := make([]int, 10)
	// 去掉 “ , value ” 的部分, 只获取下标
	for i := range pow {
		pow[i] = 1 << uint(1)
	}
	// 通过赋值给 _ 来忽略下标
	for _, value := range pow {
		fmt.Printf("%d\n", value)
	}
}
