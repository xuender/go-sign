package sign

import (
	"errors"
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/denisbrodbeck/machineid"
)

// nolint
var Safe = "weak"

func Check(secret string) error {
	if file := os.Args[0]; IsBuild(file) {
		return Error(file, CheckFile(file, secret))
	}

	return nil
}

func IsBuild(file string) bool {
	if Safe != "weak" {
		return true
	}

	dir := filepath.Dir(file)
	base := filepath.Base(file)

	if dir == "/tmpfs" && base == "play" {
		return false
	}

	if strings.HasPrefix(dir, os.TempDir()) && strings.Contains(dir, "go-build") {
		if filepath.Base(dir) == "exe" {
			return false
		}

		if filepath.Ext(base) == ".test" {
			return false
		}
	}

	return true
}

func Error(file string, err error) error {
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
