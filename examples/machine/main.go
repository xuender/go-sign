package main

import (
	"fmt"

	"github.com/xuender/go-sign"
)

func main() {
	if err := sign.CheckMachine(); err != nil {
		panic(err)
	}

	fmt.Println("Hello Word.")
	fmt.Println("Run on sign machine.")
}
