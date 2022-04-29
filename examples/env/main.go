package main

import (
	"fmt"

	"github.com/xuender/go-sign"
)

func main() {
	if err := sign.CheckEnv("SECRET_KEY"); err != nil {
		panic(err)
	}

	fmt.Println("Hello Word.")
	fmt.Println("Run on safe environment.")
}
