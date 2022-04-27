package gosign_test

import (
	"errors"
	"os"
	"path/filepath"
	"testing"

	"github.com/xuender/gosign"
)

func TestSign(t *testing.T) {
	t.Parallel()

	_ = os.Mkdir("dist", gosign.DirMode)

	file := filepath.Join("dist", "tmp")
	data := make([]byte, 100)
	_ = os.WriteFile(file, data, gosign.FileMode)

	secret := []byte("testkey")
	sum := gosign.NewSign(file, secret)

	if !errors.Is(sum.Check(), gosign.ErrSignFailed) {
		t.Errorf("Check() error = %v, wantErr %v", sum.Error, gosign.ErrSignFailed)
	}

	if err := sum.Sign(); err != nil {
		t.Errorf("Sign() error = %v, wantErr %v", err, nil)
	}

	sum = gosign.NewSign(file, secret)
	if err := sum.Sign(); !errors.Is(err, gosign.ErrSigned) {
		t.Errorf("Sign() error = %v, wantErr %v", err, gosign.ErrSigned)
	}
}
