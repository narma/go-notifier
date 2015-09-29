package utils

import (
	"path/filepath"
	"runtime"
	"strings"
)

func ProjectRoot() string {
	_, file, _, _ := runtime.Caller(0)
	filelist := strings.Split(filepath.Clean(file), string(filepath.Separator))

	if file[0] == filepath.Separator {
		filelist[0] = string(filepath.Separator) + filelist[0]
	}

	l := len(filelist) - 2
	return filepath.Join(filelist[:l]...)
}
