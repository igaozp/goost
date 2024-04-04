package main

import (
	"fmt"
	"math"
)

func pow(x, n, lim float64) float64 {
	// if 语句可以在条件之前执行一个简单语句
	if v := math.Pow(x, n); v < lim {
		return v
	} else {
		fmt.Printf("%g >= %g", v, lim)
	}
	return lim
}

func main() {
	fmt.Println(
		pow(3, 2, 10),
		pow(3, 3, 20),
	)
}
