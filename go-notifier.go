package gonotifier

import (
	"runtime"

	"github.com/narma/go-notifier/lib"
	"github.com/narma/go-notifier/notifiers/balloon"
	"github.com/narma/go-notifier/notifiers/growl"
	"github.com/narma/go-notifier/notifiers/notifysend"
	"github.com/narma/go-notifier/notifiers/osxnotify"
	"github.com/narma/go-notifier/notifiers/toaster"
	"github.com/narma/go-notifier/utils"
)

func New() (lib.Notifier, error) {
	switch runtime.GOOS {
	case "darwin":
		if utils.IsMountainLion() {
			return osxnotify.New("")
		}
	case "windows":
		if utils.IsWin8() {
			return toaster.New("")
		}
		return balloon.New("")
	}
	if runtime.GOOS != "linux" && utils.HasGrowl() {
		return growl.New("")
	}
	return notifysend.New("")
}
