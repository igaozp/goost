package main

import "fmt"

func main() {
	sum := 1
	// 相当于其他语言的 while 循环
	for sum < 1000 {
		sum += sum
	}
	fmt.Println(sum)
}
