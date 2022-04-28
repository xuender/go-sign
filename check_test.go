package gosign_test

import (
	"fmt"
	"testing"

	"github.com/xuender/gosign"
)

func ExampleCheck() {
	if err := gosign.Check("secret_key"); err != nil {
		panic(err)
	}

	fmt.Println("Hello Word.")
	fmt.Println("This file integrity.")

	// Output:
	// Hello Word.
	// This file integrity.
}

func ExampleCheckEnv() {
	if err := gosign.CheckEnv("SECRET_KEY"); err != nil {
		panic(err)
	}

	fmt.Println("Hello Word.")
	fmt.Println("Run on safe environment.")

	// Output:
	// Hello Word.
	// Run on safe environment.
}

func ExampleCheckMachine() {
	if err := gosign.CheckMachine(); err != nil {
		panic(err)
	}

	fmt.Println("Hello Word.")
	fmt.Println("Run on sign machine.")

	// Output:
	// Hello Word.
	// Run on sign machine.
}

func TestGetMachineSecret(t *testing.T) {
	t.Parallel()

	mid := gosign.GetMachineSecret("test")

	if len(mid) != 64 && len(mid) != 4 {
		t.Errorf("GetMachineSecret() len(mid)= %v, wantErr %v", len(mid), 64)
	}
}
