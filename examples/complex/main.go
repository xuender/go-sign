package main

import (
	"fmt"
	"os"

	"github.com/xuender/go-sign"
)

func main() {
	mid := sign.GetMachineSecret(os.Getenv("SECRET_KEY"))
	if err := sign.Check(mid); err != nil {
		panic(err)
	}

	fmt.Println("Hello Word.")
	fmt.Println("Run on sign machine and has env.")
}
