package main

import "fmt"

/*
*

	在函数中，:= 简洁赋值语句在明确类型的地方，可以用于替代 var 定义.
	函数外的每个语句都必须以关键字开始，:= 结构不能使用在函数外.
*/
func main() {
	var i, j int = 1, 2
	k := 3
	c, python, java := true, false, "no!"
	fmt.Println(i, j, k, c, python, java)
}
