package main

import "fmt"

func main() {
	// array
	var array [5]int
	array = [...]int{10, 20, 30, 40, 50}
	array[2] = 35
	for i := range array {
		fmt.Println(array[i])
	}

	array = [5]int{1: 10, 2: 20}
	for _, i2 := range array {
		fmt.Println(i2)
	}

	var pointerArray [2]*int
	pointerArray = [2]*int{0: new(int), 1: new(int)}
	*pointerArray[0] = 10
	*pointerArray[1] = 20
	for _, i2 := range pointerArray {
		fmt.Println(*i2)
	}

	var strArray [5]string
	anotherStrArray := [5]string{"Red", "Blue", "Green", "Yellow", "Pink"}
	strArray = anotherStrArray
	for _, s := range strArray {
		fmt.Println(s)
	}

	var pointerArray1 [3]*string
	pointerArray2 := [3]*string{new(string), new(string), new(string)}
	*pointerArray2[0] = "Red"
	*pointerArray2[1] = "Blue"
	*pointerArray2[2] = "Green"
	pointerArray1 = pointerArray2
	for _, s := range pointerArray1 {
		fmt.Println(*s)
	}

	var twoDArray [4][2]int
	twoDArray = [4][2]int{{10, 11}, {20, 21}, {30, 31}, {40, 41}}
	for _, ints := range twoDArray {
		for i, i2 := range ints {
			fmt.Printf("%d %d\n", i, i2)
		}
	}
	fmt.Println("=============")
	twoDArray = [4][2]int{1: {20, 21}, 3: {40, 41}}
	for _, ints := range twoDArray {
		for i, i2 := range ints {
			fmt.Printf("%d %d\n", i, i2)
		}
	}
	fmt.Println("=============")
	twoDArray = [4][2]int{1: {0: 20}, 3: {1: 4}}
	for _, ints := range twoDArray {
		for i, i2 := range ints {
			fmt.Printf("%d %d\n", i, i2)
		}
	}
	fmt.Println("=============")

	var array1 [2][2]int
	var array2 [2][2]int
	array2[0][0] = 10
	array2[0][1] = 20
	array2[1][0] = 30
	array2[1][1] = 40
	array1 = array2
	for i, ints := range array1 {
		for i2 := range ints {
			fmt.Printf("%d\n", array2[i][i2])
		}
	}

	var bigArray [1e6]int
	foo(&bigArray)

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

func foo(array *[1e6]int) {

}

func sliceFoo(slice []int) []int {
	return slice
}
