package gosign_test

import (
	"errors"
	"os"
	"testing"

	"github.com/xuender/gosign"
)

func TestSign(t *testing.T) {
	t.Parallel()

	writer, _ := os.CreateTemp(os.TempDir(), "test")
	data := make([]byte, 100)
	_, _ = writer.Write(data)
	_ = writer.Close()

	secret := []byte("testkey")
	sum := gosign.NewSign(writer.Name(), secret)

	if !errors.Is(sum.Check(), gosign.ErrSignFailed) {
		t.Errorf("Check() error = %v, wantErr %v", sum.Error, gosign.ErrSignFailed)
	}

	if err := sum.Sign(); err != nil {
		t.Errorf("Sign() error = %v, wantErr %v", err, nil)
	}

	sum = gosign.NewSign(writer.Name(), secret)
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

	writer, _ := os.CreateTemp(os.TempDir(), "test")
	data := make([]byte, 10)
	_, _ = writer.Write(data)
	_ = writer.Close()

	secret := []byte("testkey")
	if sum := gosign.NewSign(writer.Name(), secret); sum.Error == nil {
		t.Errorf("NewSign() error = %v, wantErr %v", sum.Error, gosign.ErrSignFailed)
	}
}

func TestNewSign_ReadError(t *testing.T) {
	t.Parallel()

	writer, _ := os.CreateTemp(os.TempDir(), "test")
	data := make([]byte, 100)
	_, _ = writer.Write(data)
	_ = writer.Close()
	os.Chmod(writer.Name(), 0o266)

	secret := []byte("testkey")
	if sum := gosign.NewSign(writer.Name(), secret); sum.Error == nil {
		t.Errorf("NewSign() error = %v, wantErr %v", sum.Error, "permission denied")
	}
}

func TestNewSign_WriteError(t *testing.T) {
	t.Parallel()

	writer, _ := os.CreateTemp(os.TempDir(), "test")
	data := make([]byte, 100)
	_, _ = writer.Write(data)
	_ = writer.Close()
	os.Chmod(writer.Name(), 0o466)

	secret := []byte("testkey")
	sum := gosign.NewSign(writer.Name(), secret)
	if err := sum.Sign(); err == nil {
		t.Errorf("NewSign() error = %v, wantErr %v", sum.Error, "permission denied")
	}
}
