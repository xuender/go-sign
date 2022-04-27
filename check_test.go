package gosign_test

import (
	"testing"

	"github.com/xuender/gosign"
)

func TestGetMachineSecret(t *testing.T) {
	t.Parallel()

	mid := gosign.GetMachineSecret("test")

	if len(mid) != 64 {
		t.Errorf("GetMachineSecret() len(mid)= %v, wantErr %v", len(mid), 64)
	}
}
