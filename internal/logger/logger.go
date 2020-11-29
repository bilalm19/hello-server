package logger

import (
	"os"

	"github.com/sirupsen/logrus"
)

// Logger logs logs to stderr. This is primarily used for setting log levels
// and making it easier to read them.
var Logger = &logrus.Logger{
	Out:       os.Stderr,
	Formatter: new(logrus.TextFormatter),
	Hooks:     make(logrus.LevelHooks),
	Level:     logrus.TraceLevel,
}
