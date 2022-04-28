package gosign

import (
	"runtime/debug"
	"strings"
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

	if Mod.Path == "" {
		for _, m := range info.Deps {
			if strings.HasSuffix(m.Path, "gosign") {
				Mod = m

				break
			}
		}
	}

	if Mod.Replace != nil {
		Mod = Mod.Replace
	}
}
