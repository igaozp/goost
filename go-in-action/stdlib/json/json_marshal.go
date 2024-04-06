package main

import (
	"encoding/json"
	"fmt"
	"log"
)

func main() {
	c := make(map[string]interface{})
	c["name"] = "John Doe"
	c["title"] = "Software Developer"
	c["contract"] = map[string]string{
		"home": "123-456-7890",
		"cell": "098-765-4321",
	}

	data, err := json.MarshalIndent(c, "", "  ")
	if err != nil {
		log.Println("ERROR:", err)
		return
	}

	fmt.Println(string(data))
}
