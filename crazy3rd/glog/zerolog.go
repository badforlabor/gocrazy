/**
 * Auth :   liubo
 * Date :   2021/4/14 11:44
 * Comment:
 */

package glog



import (
	"fmt"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"golang.org/x/sys/windows/svc"
	"time"

	//"time"
	//"gopkg.in/natefinch/lumberjack.v2"
	"github.com/lestrrat-go/file-rotatelogs"
	"io"
	"os"
	"path"
	log2 "log"
)

var svcIsWindowsService func()(bool, error)

// Configuration for logging
type Config struct {
	// Enable console logging
	ConsoleLoggingEnabled bool
	PrettyConsoleLoggingEnabled bool

	// EncodeLogsAsJson makes the log framework log JSON
	EncodeLogsAsJson bool
	// FileLoggingEnabled makes the framework log to a file
	// the fields below can be skipped if this value is false!
	FileLoggingEnabled bool
	// Directory to log to to when filelogging is enabled
	Directory string
	// Filename is the name of the logfile which will be placed inside the directory
	Filename string
	// MaxSize the max size in MB of the logfile before it's rolled
	MaxSize int
	// MaxBackups the max number of rolled files to keep
	MaxBackups int
	// MaxAge the max age in days to keep a logfile
	MaxAge int
}

type Logger struct {
	Inside *zerolog.Logger
}
func (self *Logger) Infoln(str string) {
	if self.Inside == nil {
		log2.Println("[Info]", str)
	} else {
		self.Inside.Info().Msg(str)
	}
}
func (self *Logger) Warnln(str string) {
	if self.Inside == nil {
		log2.Println("[Warn]", str)
	} else {
		self.Inside.Info().Msg(str)
	}
}
func (self *Logger) Errorln(str string) {
	if self.Inside == nil {
		log2.Println("[Error]", str)
	} else {
		self.Inside.Info().Msg(str)
	}
}


// Configure sets up the logging framework
//
// In production, the container logs will be collected and file logging should be disabled. However,
// during development it's nicer to see logs as text and optionally write to a file when debugging
// problems in the containerized pipeline
//
// The output log file will be located at /var/log/service-xyz/service-xyz.log and
// will be rolled according to configuration set.
func Configure(config Config, structLog bool) *Logger {
	var writers []io.Writer

	if config.ConsoleLoggingEnabled {
		var ws = false
		if svcIsWindowsService != nil {
			ws, _ = svcIsWindowsService()
		}
		if !ws {
			writers = append(writers, os.Stdout)
		}
	}
	if config.PrettyConsoleLoggingEnabled {
		writers = append(writers, zerolog.ConsoleWriter{Out: os.Stderr})
	}
	if config.FileLoggingEnabled {
		writers = append(writers, newRollingFile(config))
	}
	mw := io.MultiWriter(writers...)

	var retLog = &Logger{Inside:nil}

	// golang自带的日志
	if !structLog {
		log2.SetOutput(mw)
		return retLog
	}

	// zerolog.SetGlobalLevel(zerolog.DebugLevel)
	logger := zerolog.New(mw).With().Timestamp().Logger()

	logger.Info().
		Bool("fileLogging", config.FileLoggingEnabled).
		Bool("jsonLogOutput", config.EncodeLogsAsJson).
		Str("logDirectory", config.Directory).
		Str("fileName", config.Filename).
		Int("maxSizeMB", config.MaxSize).
		Int("maxBackups", config.MaxBackups).
		Int("maxAgeInDays", config.MaxAge).
		Msg("logging configured")

	retLog.Inside = &logger

	return retLog
}

func newRollingFile(config Config) io.Writer {
	if err := os.MkdirAll(config.Directory, 0744); err != nil {
		log.Error().Err(err).Str("path", config.Directory).Msg("can't create log directory")
		return nil
	}

	var postfix = time.Now().Format("20060102-150405")
	//postfix = "%Y%m%d%H%M"

	var rot, e = rotatelogs.New(path.Join(config.Directory, config.Filename + "." + postfix + ".log"),
		rotatelogs.WithRotationSize(int64(config.MaxSize) * 1024 * 1024),
		rotatelogs.ForceNewFile(),
	)
	if e != nil {
		fmt.Println("create log failed：", e)
	}
	return rot

	//return &lumberjack.Logger{
	//	Filename:   path.Join(config.Directory, config.Filename),
	//	MaxBackups: config.MaxBackups,		// files
	//	MaxSize:    config.MaxSize,			// megabytes
	//	MaxAge:     config.MaxAge,			// days
	//}
}