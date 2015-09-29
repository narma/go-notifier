package main

import (
	"flag"
	"log"
	"os"

	"github.com/narma/go-notifier"
	"github.com/narma/go-notifier/lib"
)

func ParseCmdLine() (options lib.Options) {
	flag.StringVar(&options.Message, "message", "", "The notification message")
	flag.StringVar(&options.Title, "title", "", "The notification title")
	flag.StringVar(&options.Subtitle, "subtitle", "", "The notification subtitle")
	flag.StringVar(&options.Sound, "sound", "", "The name of a sound to play when the notification appears. The names are listed")
	flag.StringVar(&options.Icon, "icon", "", "The notification icon")
	flag.StringVar(&options.Open, "open", "", "The URL of a resource to open when the user clicks the notification.")
	flag.BoolVar(&options.Wait, "wait", false, "Wait until the notification has been dismissed.")
	flag.Parse()
	return
}

func poe(err error) {
	if err != nil {
		panic(err)
	}
}

func main() {
	notifier, err := gonotifier.New()
	if err != nil {
		log.SetOutput(os.Stderr)
		log.Println("Can't create notifier", err)
		return
	}
	options := ParseCmdLine()
	notifier.GeneralPush(options)
}
