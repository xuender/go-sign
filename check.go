package gosign

import (
	"errors"
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/denisbrodbeck/machineid"
)

func CheckEnv(env string) error {
	return Check(os.Getenv(env))
}

func GetMachineSecret() (string, error) {
	return machineid.ProtectedID("github.com/xuender/gosign")
}

func CheckMachine() error {
	id, err := GetMachineSecret()
	if err != nil {
		return err
	}

	return Check(id)
}

func Check(secret string) error {
	file := os.Args[0]
	if strings.Contains(file, "go-build") {
		return nil
	}

	err := CheckFile(file, secret)
	if errors.Is(ErrSignFailed, err) {
		// nolint
		return fmt.Errorf("%s Sign FAILED", filepath.Base(file))
	}

	return err
}

func CheckFile(file, secret string) error {
	sum := NewSign(file, []byte(secret))
	if sum.Error != nil {
		return sum.Error
	}

	return sum.Check()
}
