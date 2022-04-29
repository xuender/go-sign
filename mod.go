package sign

import (
	"runtime/debug"
	"strings"
)

// nolint
var Mod = GetMod(debug.ReadBuildInfo())

func GetMod(info *debug.BuildInfo, ok bool) *debug.Module {
	if !ok {
		return nil
	}

	ret := &info.Main

	if ret.Replace != nil {
		ret = ret.Replace
	}

	if ret.Path == "" {
		for _, m := range info.Deps {
			if strings.HasSuffix(m.Path, "go-sign") {
				ret = m

				break
			}
		}
	}

	if ret.Replace != nil {
		return ret.Replace
	}

	return ret
}
