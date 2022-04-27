package main

import (
	"fmt"
	"os"

	"github.com/xuender/gosign"
)

func main() {
	mid := gosign.GetMachineSecret(os.Getenv("SECRET_KEY"))
	if err := gosign.Check(mid); err != nil {
		panic(err)
	}

	fmt.Println("Hello Word.")
	fmt.Println("Run on sign machine and has env.")
}
