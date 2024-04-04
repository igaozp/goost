package main

import "fmt"

type Person struct {
	Name string
	Age  int
}

/**
type Stringer interface {
	String() string
}
*/

func (p Person) String() string {
	return fmt.Sprintf("%v (%v years", p.Name, p.Age)
}

func main() {
	a := Person{"Arthur Dent", 42}
	z := Person{"Zaphod Beeblebrox", 9001}
	// Stringer 是一个可以用字符串描述自己的类型
	// "fmt"包使用这个来进行输出
	// java 中的 toString
	fmt.Println(a, z)
}
