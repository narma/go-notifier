package osxnotify

import (
	"os/exec"

	"github.com/narma/go-notifier/lib"
	"github.com/narma/go-notifier/utils"

	"github.com/imdario/mergo"
)

type Options struct {
	Title        string
	Message      string
	Subtitle     string
	Sound        string
	Group        string
	Remove       string
	List         string
	Activate     string
	Sender       string
	AppIcon      string
	ContentImage string
	Open         string
	Execute      string
}

type Notifier struct {
	Cmd           string
	argsFormatter lib.ArgsFormatter
}

var projectRoot string

func init() {
	projectRoot = utils.ProjectRoot()
}

func New(cmd string) (*Notifier, error) {
	if cmd == "" {
		cmd = projectRoot + "/vendor/terminal-notifier.app/Contents/MacOS/terminal-notifier"
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

func (n Notifier) Push(opts Options) *exec.Cmd {
	args := n.argsFormatter.FormatArgs(opts)
	cmd := exec.Command(n.Cmd, args...)
	return cmd
}

func (n Notifier) GeneralPush(options lib.Options, aopts ...interface{}) error {
	nativeOpts := Options{
		Title:   options.Title,
		Message: options.Message,
		AppIcon: options.Icon,
	}

	for _, opts := range aopts {
		if nopts, ok := opts.(Options); ok {
			mergo.Merge(&nativeOpts, nopts)
		}
	}
	return n.Push(nativeOpts).Run()
}
