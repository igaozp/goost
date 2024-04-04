package main

import (
	_ "github.com.goost.go-in-action.fastlook/matchers"
	"github.com.goost.go-in-action.fastlook/search"
	"log"
	"os"
)

func init() {
	log.SetOutput(os.Stdout)
}

func main() {
	search.Run("president")
}
