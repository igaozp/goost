package main

import "fmt"

var maps = map[string]Vertex{
	"Bell Labs": Vertex{40.44, -66.666},
	"Google":    Vertex{37.1, -122.05},
}

func main() {
	fmt.Println(maps)
}
