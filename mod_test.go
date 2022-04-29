package gosign_test

import (
	"runtime/debug"
	"testing"

	"github.com/xuender/gosign"
)

func TestGetMod(t *testing.T) {
	t.Parallel()

	if mod := gosign.GetMod(nil, false); mod != nil {
		t.Errorf("GetMod() return= %v, wantErr %v", mod, nil)
	}

	module := debug.Module{Replace: &debug.Module{}}
	info := &debug.BuildInfo{Main: module, Deps: []*debug.Module{{
		Path:    "xuender/gosign",
		Replace: &debug.Module{},
	}}}

	if mod := gosign.GetMod(info, true); mod == nil {
		t.Errorf("GetMod() return= %v, wantErr %v", mod, module)
	}
}
