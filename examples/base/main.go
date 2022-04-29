package main

import (
	"fmt"

	"github.com/xuender/go-sign"
)

func main() {
	if err := sign.Check("secret_key"); err != nil {
		panic(err)
	}

	fmt.Println("Hello Word.")
	fmt.Println("This file integrity.")
}
