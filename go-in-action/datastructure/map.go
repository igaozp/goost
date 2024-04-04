package main

import "fmt"

func main() {

	// map
	dict := make(map[string]string)
	dict = map[string]string{"Red": "#da1337", "Orange": "#e95a22"}
	for key, value := range dict {
		fmt.Printf("%s: %s\n", key, value)
	}
	dict["Red"] = "#da1337"

	// blank map
	var _ map[string]string

	value, exists := dict["Red"]
	if exists {
		fmt.Printf("%s", value)
	}

	delete(dict, "Blue")
}
