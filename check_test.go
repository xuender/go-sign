package sign_test

import (
	"errors"
	"fmt"
	"os"
	"path/filepath"
	"testing"

	"github.com/xuender/go-sign"
)

func ExampleCheck() {
	if err := sign.Check("secret_key"); err != nil {
		panic(err)
	}

	fmt.Println("Hello Word.")
	fmt.Println("This file integrity.")

	// Output:
	// Hello Word.
	// This file integrity.
}

func ExampleCheckEnv() {
	if err := sign.CheckEnv("SECRET_KEY"); err != nil {
		panic(err)
	}

	fmt.Println("Hello Word.")
	fmt.Println("Run on safe environment.")

	// Output:
	// Hello Word.
	// Run on safe environment.
}

func ExampleCheckMachine() {
	if err := sign.CheckMachine(); err != nil {
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

	mid := sign.GetMachineSecret("test")

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

	if err := sign.CheckFile("/tmp", "key"); err == nil {
		t.Errorf("CheckFile() Error= %v, wantErr %v", err, nil)
	}

	sum := sign.NewSign(file.Name(), []byte("key"))
	_ = sum.Sign()

	if err := sign.CheckFile(file.Name(), "key"); err != nil {
		t.Errorf("CheckFile() Error= %v, wantErr %v", err, nil)
	}
}

func TestError(t *testing.T) {
	t.Parallel()

	if err := sign.Error("test", nil); err != nil {
		t.Errorf("Error() Error= %v, wantErr %v", err, nil)
	}

	if err := sign.Error("test", sign.ErrSignFailed); errors.Is(err, sign.ErrSignFailed) {
		t.Errorf("Error() Error= %v, wantErr %v", err, sign.ErrSignFailed)
	}

	if err := sign.Error("test", sign.ErrSigned); !errors.Is(err, sign.ErrSigned) {
		t.Errorf("Error() Error= %v, wantErr %v", err, sign.ErrSigned)
	}
}

func TestIsBuild(t *testing.T) {
	t.Parallel()

	if sign.IsBuild("/tmpfs/play") {
		t.Errorf("IsBuild() = false")
	}

	if sign.IsBuild(filepath.Join(os.TempDir(), "go-build123", "exe", "main")) {
		t.Errorf("IsBuild() = false")
	}

	if sign.IsBuild(filepath.Join(os.TempDir(), "go-build123", "main.test")) {
		t.Errorf("IsBuild() = false")
	}

	if !sign.IsBuild(filepath.Join("aa", "bb")) {
		t.Errorf("IsBuild() = true")
	}
}
