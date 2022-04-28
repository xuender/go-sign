package gosign

import (
	"errors"
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/denisbrodbeck/machineid"
)

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

func CheckEnv(env string) error {
	return Check(os.Getenv(env))
}

func CheckMachine() error {
	return Check(GetMachineSecret(Mod.Path))
}

func CheckFile(file, secret string) error {
	sum := NewSign(file, []byte(secret))
	if sum.Error != nil {
		return sum.Error
	}

	return sum.Check()
}

func GetMachineSecret(secret string) string {
	if secret == "" {
		secret = Mod.Path
	}

	if mid, err := machineid.ProtectedID(secret); err == nil {
		return mid
	}

	return secret
}
