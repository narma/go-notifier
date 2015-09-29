package lib

type Notifier interface {
	GeneralPush(options Options, otherOpts ...interface{}) error
}

type Options struct {
	Message  string
	Title    string
	Subtitle string
	Sound    string
	Icon     string
	Open     string
	Wait     bool
}
