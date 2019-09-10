package main

import (
	"fmt"

	"github.com/chilts/sid"
)

func main() {
	fmt.Println("Go mudules")
	id := sid.Id()
	fmt.Println("Generated Id: ", id)
}
