package gosign_test

import (
	"errors"
	"os"
	"testing"

	"github.com/xuender/gosign"
)

func TestSign(t *testing.T) {
	t.Parallel()

	file, _ := os.CreateTemp(os.TempDir(), "sing")
	defer os.Remove(file.Name())

	_, _ = file.Write(make([]byte, 100))
	_ = file.Close()

	secret := []byte("key")
	sum := gosign.NewSign(file.Name(), secret)

	if !errors.Is(sum.Check(), gosign.ErrSignFailed) {
		t.Errorf("Check() error = %v, wantErr %v", sum.Error, gosign.ErrSignFailed)
	}

	if err := sum.Sign(); err != nil {
		t.Errorf("Sign() error = %v, wantErr %v", err, nil)
	}

	sum = gosign.NewSign(file.Name(), secret)
	if err := sum.Sign(); !errors.Is(err, gosign.ErrSigned) {
		t.Errorf("Sign() error = %v, wantErr %v", err, gosign.ErrSigned)
	}
}

func TestNewSign_NotFile(t *testing.T) {
	t.Parallel()

	if sum := gosign.NewSign("/ffff", []byte("xx")); sum.Error == nil {
		t.Errorf("NewSign() error = %v, wantErr %v", sum.Error, nil)
	}
}

func TestNewSign_SignFailed(t *testing.T) {
	t.Parallel()

	file, _ := os.CreateTemp(os.TempDir(), "sign")
	defer os.Remove(file.Name())

	_, _ = file.Write(make([]byte, 10))
	_ = file.Close()

	if sum := gosign.NewSign(file.Name(), []byte("key")); sum.Error == nil {
		t.Errorf("NewSign() error = %v, wantErr %v", sum.Error, gosign.ErrSignFailed)
	}
}

func TestNewSign_ReadError(t *testing.T) {
	t.Parallel()

	file, _ := os.CreateTemp(os.TempDir(), "sign")
	defer os.Remove(file.Name())

	_, _ = file.Write(make([]byte, 100))
	_ = file.Close()
	_ = os.Chmod(file.Name(), 0o266)

	if sum := gosign.NewSign(file.Name(), []byte("key")); sum.Error == nil {
		t.Errorf("NewSign() error = %v, wantErr %v", sum.Error, "permission denied")
	}
}

func TestNewSign_WriteError(t *testing.T) {
	t.Parallel()

	file, _ := os.CreateTemp(os.TempDir(), "sign")
	defer os.Remove(file.Name())

	_, _ = file.Write(make([]byte, 100))
	_ = file.Close()
	_ = os.Chmod(file.Name(), 0o466)

	sum := gosign.NewSign(file.Name(), []byte("key"))
	if err := sum.Sign(); err == nil {
		t.Errorf("NewSign() error = %v, wantErr %v", sum.Error, "permission denied")
	}
}

func TestSign_Hash_ReadErr(t *testing.T) {
	t.Parallel()

	sum := gosign.NewSign("/tmp", []byte("key"))
	file, _ := os.CreateTemp(os.TempDir(), "sign")
	_, _ = file.Write(make([]byte, 100))

	defer os.Remove(file.Name())

	file.Close()
	_ = os.Chmod(file.Name(), 0o266)
	file, _ = os.Open(file.Name())

	if _, err := sum.Hash(file); err == nil {
		t.Errorf("Sign() error = %v, wantErr %v", err, "permission denied")
	}
}

func TestSign_Sign_Err(t *testing.T) {
	t.Parallel()

	file, _ := os.CreateTemp(os.TempDir(), "sign")
	_, _ = file.Write(make([]byte, 100))

	defer os.Remove(file.Name())

	sum := gosign.NewSign(file.Name(), []byte("key"))
	_ = os.Chmod(file.Name(), 0o466)

	if err := sum.Sign(); err == nil {
		t.Errorf("Sign() error = %v, wantErr %v", err, "permission denied")
	}
}
