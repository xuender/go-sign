package gosign_test

import (
	"testing"

	"github.com/xuender/gosign"
)

func TestGetMod(t *testing.T) {
	t.Parallel()

	if mod := gosign.GetMod(); mod == nil {
		t.Errorf("GetMod() return= %v, wantErr %v", mod, nil)
	}
}
