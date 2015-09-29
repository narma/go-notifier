package growl

import (
	"os/exec"

	"github.com/imdario/mergo"
	"github.com/narma/go-notifier/lib"
)

type Options struct {
	Name       string
	NoteName   string
	Message    string
	Sticky     bool
	AppIcon    string
	Icon       string
	IconPath   string
	Priority   string
	Identifier string
	Host       string
	Password   string
	Wait       bool `arg:"wait,omitvalue"`
	Url        string
}

var projectRoot string

type Notifier struct {
	Cmd           string
	argsFormatter lib.ArgsFormatter
}

func New(cmd string) (*Notifier, error) {
	if cmd == "" {
		cmd = "growlnotify"
	}
	path, err := exec.LookPath(cmd)
	if err != nil {
		return nil, err
	}
	return &Notifier{
		Cmd: path,
		argsFormatter: lib.ArgsFormatter{
			KeyPrefix:              "-",
			DisableSplitMultiwords: true,
		},
	}, nil
}

func (n Notifier) NativePush(opts Options) error {
	args := n.argsFormatter.FormatArgs(opts)
	cmd := exec.Command(n.Cmd, args...)
	return cmd.Run()
}

func (n Notifier) Push(options lib.Options, aopts ...interface{}) error {
	nativeOpts := Options{
		Name:    options.Title,
		Message: options.Message,
		Icon:    options.Icon,
		Wait:    options.Wait,
	}

	for _, opts := range aopts {
		if nopts, ok := opts.(Options); ok {
			mergo.Merge(&nativeOpts, nopts)
		}
	}
	return n.Push(nativeOpts).Run()
}
