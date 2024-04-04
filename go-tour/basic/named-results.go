package main

import "fmt"

/*
*

	没有参数的 return 语句返回各个返回变量的当前值
	这种用法被称作"裸"返回
*/
func split(sum int) (x, y int) {
	x = sum * 4 / 9
	y = sum - x
	return
}

func main() {
	fmt.Println(split(17))
}
