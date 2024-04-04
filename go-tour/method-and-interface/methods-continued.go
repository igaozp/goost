package main

import (
	"fmt"
	"math"
)

func main() {
	f := MyFloat(-math.Sqrt2)
	fmt.Println(f.Abs())
}
