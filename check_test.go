package gosign_test

import (
	"errors"
	"fmt"
	"os"
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

func TestCheckFile(t *testing.T) {
	t.Parallel()

	file, _ := os.CreateTemp(os.TempDir(), "check")
	defer os.Remove(file.Name())

	_, _ = file.Write(make([]byte, 100))
	file.Close()

	if err := gosign.CheckFile("/tmp", "key"); err == nil {
		t.Errorf("CheckFile() Error= %v, wantErr %v", err, nil)
	}

	sum := gosign.NewSign(file.Name(), []byte("key"))
	_ = sum.Sign()

	if err := gosign.CheckFile(file.Name(), "key"); err != nil {
		t.Errorf("CheckFile() Error= %v, wantErr %v", err, nil)
	}
}

func TestError(t *testing.T) {
	t.Parallel()

	if err := gosign.Error("test", nil); err != nil {
		t.Errorf("Error() Error= %v, wantErr %v", err, nil)
	}

	if err := gosign.Error("test", gosign.ErrSignFailed); errors.Is(err, gosign.ErrSignFailed) {
		t.Errorf("Error() Error= %v, wantErr %v", err, gosign.ErrSignFailed)
	}

	if err := gosign.Error("test", gosign.ErrSigned); !errors.Is(err, gosign.ErrSigned) {
		t.Errorf("Error() Error= %v, wantErr %v", err, gosign.ErrSigned)
	}
}
