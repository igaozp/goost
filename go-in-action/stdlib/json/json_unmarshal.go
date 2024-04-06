package main

import (
	"encoding/json"
	"fmt"
	"log"
)

type Contract struct {
	Name     string `json:"name"`
	Title    string `json:"title"`
	Contract struct {
		Home string `json:"home"`
		Cell string `json:"cell"`
	} `json:"contract"`
}

var JSON = `{
	"name": "John Doe",
	"title": "Software Developer",
	"contract": {
		"home": "123-456-7890",
		"cell": "098-765-4321"
	}
}`

func main() {
	var c Contract
	err := json.Unmarshal([]byte(JSON), &c)
	if err != nil {
		log.Println("ERROR:", err)
		return
	}

	fmt.Println(c)

	var a map[string]interface{}
	err = json.Unmarshal([]byte(JSON), &a)
	if err != nil {
		log.Println("ERROR:", err)
		return
	}

	fmt.Println("Name:", a["name"])
	fmt.Println("Title:", a["title"])
	fmt.Println("Contract")
	fmt.Println("H:", a["contract"].(map[string]interface{})["home"])
	fmt.Println("C:", a["contract"].(map[string]interface{})["cell"])
}
