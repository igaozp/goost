package main

import "fmt"

func main() {
	var z []int
	fmt.Println(z, len(z), cap(z))
	// slice 的零值是 nil
	if z == nil {
		fmt.Println("nil!")
	}
}
