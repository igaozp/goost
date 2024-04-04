package main

import "fmt"

func main() {
	// slice
	slice := make([]string, 5)
	slice = make([]string, 3, 5)

	slice = []string{"Red", "Blue", "Green", "Yellow", "Pink"}
	for _, s := range slice {
		fmt.Printf("%s ", s)
	}
	fmt.Println()
	slice = []string{99: ""}
	for _, s := range slice {
		fmt.Printf("%s ", s)
	}

	// nil slice
	var _ []int
	// blank slice
	_ = make([]int, 0)
	_ = []int{}

	newSlice := slice[1:3]
	for _, s := range newSlice {
		fmt.Printf("%s ", s)
	}

	appendSlice := []string{"a", "b", "c"}
	appendSlice = append(appendSlice, "d")
	for _, s := range appendSlice {
		fmt.Printf("%s ", s)
	}
	fmt.Println()

	source := []string{"Apple", "Orange", "Plum", "Banana", "Grape"}
	fruitSlice := source[2:3:4]
	for _, s := range fruitSlice {
		fmt.Printf("%s ", s)
	}
	fmt.Println()

	slice = source[2:3:3]
	slice = append(slice, "Kiwi")
	for _, s := range slice {
		fmt.Printf("%s ", s)
	}

	s1 := []int{1, 2}
	s2 := []int{3, 4}
	fmt.Printf("%v\n", append(s1, s2...))

	numSlice := []int{10, 20, 30, 40}
	for i, i2 := range numSlice {
		fmt.Printf("Value: %d Value-Addr: %X ElemAddr: %X\n", i2, &i2, &numSlice[i])
	}

	twoDSlice := [][]int{{10}, {100, 200}}
	twoDSlice[0] = append(twoDSlice[0], 20)

	bigSlice := make([]int, 1e6)
	bigSlice = sliceFoo(bigSlice)
}

func sliceFoo(slice []int) []int {
	return slice
}
