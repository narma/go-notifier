package notifysend

import (
	"os/exec"

	"github.com/imdario/mergo"
	"github.com/narma/go-notifier/lib"
)

type Options struct {
	Title      string `arg:"-"`
	Message    string `arg:"-"`
	Urgency    string
	AppName    string
	ExpireTime uint
	Icon       string
	Category   string
}

type Notifier struct {
	Cmd           string
	argsFormatter lib.ArgsFormatter
}

func New(cmd string) (*Notifier, error) {
	if cmd == "" {
		cmd = "notify-send"
	}
	path, err := exec.LookPath(cmd)
	if err != nil {
		return nil, err
	}
	return &Notifier{
		Cmd: path,
		argsFormatter: lib.ArgsFormatter{
			KeyPrefix: "--",
		},
	}, nil
}

func (n Notifier) Push(opts Options) *exec.Cmd {
	args := n.argsFormatter.FormatArgs(opts)
	cmd := exec.Command(n.Cmd, args...)
	return cmd
}

func (n Notifier) GeneralPush(options lib.Options, aopts ...interface{}) error {
	nativeOpts := Options{
		Title:   options.Title,
		Message: options.Message,
		Icon:    options.Icon,
	}

	for _, opts := range aopts {
		if nopts, ok := opts.(Options); ok {
			mergo.Merge(&nativeOpts, nopts)
		}
	}
	return n.Push(nativeOpts).Run()
}
