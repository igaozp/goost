package main

import "fmt"

/*
*

	在定义一个变量却并不显式指定其类型时（使用 := 语法或者 var = 表达式语法），
	变量的类型由（等号）右侧的值推导得出.
*/
func main() {
	v := 42
	fmt.Printf("v is of type %T\n", v)
}
