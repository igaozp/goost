package main

import "fmt"

type Vertex struct {
	Lat, Long float64
}

func main() {
	var m = map[string]Vertex{
		"Bell Labs": {40.1, -74.2},
		"Google":    {37.23, -122.05},
	}
	fmt.Println(m)
}
