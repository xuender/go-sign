package gosign

import (
	"runtime/debug"
)

// nolint
var (
	ModVersion string
	ModSum     string
	ModPath    string
)

// nolint
func init() {
	info, ok := debug.ReadBuildInfo()
	if !ok {
		return
	}

	mod := &info.Main

	if mod.Path == "" {
		mod = info.Deps[len(info.Deps)-1]
	}

	if mod.Replace != nil {
		mod = mod.Replace
	}

	ModVersion = mod.Version
	ModSum = mod.Sum
	ModPath = mod.Path
}
