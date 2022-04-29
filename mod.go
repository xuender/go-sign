package gosign

import (
	"runtime/debug"
	"strings"
)

// nolint
var (
	Mod = GetMod()
)

func GetMod() *debug.Module {
	info, ok := debug.ReadBuildInfo()
	if !ok {
		return nil
	}

	ret := &info.Main

	if ret.Replace != nil {
		ret = ret.Replace
	}

	if ret.Path == "" {
		for _, m := range info.Deps {
			if strings.HasSuffix(m.Path, "gosign") {
				return m
			}
		}
	}

	if ret.Replace != nil {
		return ret.Replace
	}

	return ret
}
