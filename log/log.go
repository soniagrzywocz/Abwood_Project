package log

import (
	"fmt"
	"go_server/config"
	"os"
	"strings"

	"github.com/go-gem/log"
)

var apiLogger *log.Logger
var apiLogPath string

var Debugf func(string, ...interface{})

var Errorf func(string, ...interface{})

var Error func(...interface{})

var Infof func(string, ...interface{})

var Printf func(string, ...interface{})

var Println func(...interface{})

var Warningf func(string, ...interface{})

var Fatalf func(string, ...interface{})

var Fatal func(...interface{})

func panicf(message string, args ...interface{}) {
	panic(fmt.Sprintf(message+"\n", args...))
}

func InitializeLog() {
	apiLogPath = config.C.Logging.ApiLogPath
	logLevel := strings.ToLower(config.C.Logging.LogLevel)
	var level int

	switch logLevel {
	case "all":
		level = log.LevelAll
	case "debug":
		level = log.LevelDebug | log.LevelInfo | log.LevelWarning | log.LevelError | log.LevelFatal
	case "info":
		level = log.LevelInfo | log.LevelWarning | log.LevelError | log.LevelFatal
	case "warn":
		level = log.LevelWarning | log.LevelError | log.LevelFatal
	case "error":
		level = log.LevelError | log.LevelFatal
	case "fatal":
		level = log.LevelFatal
	default:
		level = log.LevelDebug | log.LevelInfo | log.LevelWarning | log.LevelError | log.LevelFatal
	}

	if apiLogPath != "" {
		out, err := os.OpenFile(apiLogPath, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0600)
		if err != nil {
			panicf("%s: %s\n", apiLogPath, err)
		}
		apiLogger = log.New(out, log.Lshortfile|log.LstdFlags, level)
	} else {
		apiLogger = log.New(os.Stderr, log.Lshortfile|log.LstdFlags, level)
	}

	Debugf = apiLogger.Debugf
	Errorf = apiLogger.Errorf
	Error = apiLogger.Error
	Infof = apiLogger.Infof
	Printf = apiLogger.Printf
	Warningf = apiLogger.Warningf
	Println = apiLogger.Println
}
