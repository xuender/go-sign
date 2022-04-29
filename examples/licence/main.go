package main

import (
	"fmt"
	"os"

	"github.com/xuender/go-sign"
)

func main() {
	if len(os.Args) < 2 {
		panic("Miss licence.")
	}

	if err := sign.Check(os.Args[1]); err != nil {
		panic("Licence FAILED.")
	}

	fmt.Println("Hello Word.")
	fmt.Println("Licence OK.")
}
