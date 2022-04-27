package gosign_test

import (
	"testing"

	"github.com/xuender/gosign"
)

func TestGetMachineSecret(t *testing.T) {
	t.Parallel()

	mid, err := gosign.GetMachineSecret("test")
	if err != nil {
		t.Errorf("GetMachineSecret() error = %v, wantErr %v", err, nil)
	}

	if len(mid) != 64 {
		t.Errorf("GetMachineSecret() len(mid)= %v, wantErr %v", len(mid), 64)
	}
}
