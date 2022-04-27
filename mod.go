package gosign

import (
	"runtime/debug"
)

// nolint
var (
	ModVersion = "(devel)"
	ModSum     = "00000000"
	ModPath    = ""
)

// nolint
func init() {
	if ModVersion != "(devel)" {
		return
	}

	info, ok := debug.ReadBuildInfo()
	if !ok {
		return
	}

	mod := &info.Main
	if mod.Replace != nil {
		mod = mod.Replace
	}

	ModVersion = mod.Version
	ModSum = mod.Sum
	ModPath = mod.Path
}
