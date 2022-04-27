package main

import (
	"fmt"

	"github.com/xuender/gosign"
)

func main() {
	if err := gosign.Check("secret_key"); err != nil {
		panic(err)
	}

	fmt.Println("Hello Word.")
	fmt.Println("This file integrity.")
}
