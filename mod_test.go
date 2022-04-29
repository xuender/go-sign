package sign_test

import (
	"runtime/debug"
	"testing"

	"github.com/xuender/go-sign"
)

func TestGetMod(t *testing.T) {
	t.Parallel()

	if mod := sign.GetMod(nil, false); mod != nil {
		t.Errorf("GetMod() return= %v, wantErr %v", mod, nil)
	}

	module := debug.Module{Replace: &debug.Module{}}
	info := &debug.BuildInfo{Main: module, Deps: []*debug.Module{{
		Path:    "xuender/go-sign",
		Replace: &debug.Module{},
	}}}

	if mod := sign.GetMod(info, true); mod == nil {
		t.Errorf("GetMod() return= %v, wantErr %v", mod, module)
	}
}
