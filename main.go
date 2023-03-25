package main

import (
	"fmt"
	"os"
)

func main() {
	exampleInput := os.Getenv("MESSAGE")

	if exampleInput == "" {
		fmt.Println("No input provided")
		return
	}

	fmt.Println("[DEBUG] message: ", exampleInput)
}