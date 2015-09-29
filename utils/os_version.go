package utils

import (
	"github.com/blang/semver"
	"runtime"
)

var (
	win8ver, _         = semver.Make("6.2.9200")
	mountainLionVer, _ = semver.Make("12.0.0")
)

func IsWin8() bool {
	osver, err := semver.Make(OSRelease())
	if err != nil {
		return false
	}
	return runtime.GOOS == "windows" && osver.GTE(win8ver)
}

func IsMountainLion() bool {
	osver, err := semver.Make(OSRelease())
	if err != nil {
		return false
	}
	return runtime.GOOS == "darwin" && osver.GTE(mountainLionVer)
}
