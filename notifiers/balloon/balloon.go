package balloon

import (
	"os/exec"
	"path/filepath"

	"github.com/imdario/mergo"
	"github.com/narma/go-notifier/lib"
	"github.com/narma/go-notifier/utils"
)

type Options struct {
	Title   string `arg:"p"`
	Message string `arg:"m"`
	Icon    string `arg:"i"`
	Quiet   bool   `arg:"q,omitvalue"`
	Wait    bool   `arg:"w,omitvalue"`
	Xp      bool   `arg:xp,omitvalue"`
	Type    string `arg:t"`
}

var projectRoot string

func init() {
	projectRoot = utils.ProjectRoot()
}

type Notifier struct {
	Cmd           string
	argsFormatter lib.ArgsFormatter
}

func New(cmd string) (*Notifier, error) {
	if cmd == "" {
		cmd = filepath.Join(projectRoot, "vendor", "toaster", "notifu.exe")
	}
	path, err := exec.LookPath(cmd)
	if err != nil {
		return nil, err
	}
	return &Notifier{
		Cmd: path,
		argsFormatter: lib.ArgsFormatter{
			KeyPrefix: "/",
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
		Title:   options.Title,
		Message: options.Message,
		Icon:    options.Icon,
		Wait:    options.Wait,
	}

	for _, opts := range aopts {
		if nopts, ok := opts.(Options); ok {
			mergo.Merge(&nativeOpts, nopts)
		}
	}
	return n.NativePush(nativeOpts)
}
