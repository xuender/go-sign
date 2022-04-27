package main

import (
	"fmt"

	"github.com/xuender/gosign"
)

func main() {
	if err := gosign.CheckMachine(); err != nil {
		panic(err)
	}

	fmt.Println("Hello Word.")
	fmt.Println("Run on sign machine.")
}
