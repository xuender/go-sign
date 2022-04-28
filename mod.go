package gosign

import (
	"runtime/debug"
)

// nolint
var (
	Mod *debug.Module
)

// nolint
func init() {
	info, ok := debug.ReadBuildInfo()
	if !ok {
		return
	}

	Mod = &info.Main

	if Mod.Replace != nil {
		Mod = Mod.Replace
	}

	if Mod.Path == "" && len(info.Deps) > 0 {
		Mod = info.Deps[len(info.Deps)-1]
	}

	if Mod.Replace != nil {
		Mod = Mod.Replace
	}
}
