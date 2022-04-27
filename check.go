package gosign

import (
	"errors"
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/denisbrodbeck/machineid"
)

const modPath = "github.com/xuender/gosign"

func CheckEnv(env string) error {
	return Check(os.Getenv(env))
}

func GetMachineSecret(secret string) string {
	if secret == "" {
		secret = modPath
	}

	if mid, err := machineid.ProtectedID(secret); err == nil {
		return mid
	}

	return secret
}

func CheckMachine() error {
	return Check(GetMachineSecret(modPath))
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
