package main

import (
	"fmt"
	"math"
)

func main() {
	// 在 Go 中，首字母大写的名称是被导出的，math.pi 是错误的
	fmt.Println(math.Pi)
}
