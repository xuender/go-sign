package main

import (
	"fmt"
	"os"

	"github.com/xuender/gosign"
)

func main() {
	if len(os.Args) < 2 {
		panic("Miss licence.")
	}

	if err := gosign.Check(os.Args[1]); err != nil {
		panic("Licence FAILED.")
	}

	fmt.Println("Hello Word.")
	fmt.Println("Licence is OK.")
}
