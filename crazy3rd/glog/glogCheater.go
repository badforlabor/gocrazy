package glog

import (
	"fmt"
	"github.com/badforlabor/gocrazy/crazyos"
	"os"
	"path"
	"strings"
)

var myLogger *Logger

const defaultLogFileName = "log.log"

var LogFileName = defaultLogFileName
var StructLog = false

func BaseInit() {

	if LogFileName == defaultLogFileName {
		LogFileName = crazyos.GetAppName()
	}

	if len(LogFileName) == 0 {
		LogFileName = defaultLogFileName
	}

	dir := crazyos.GetExecFolder()
	dir = path.Join(dir, "logs")

	var cfg = GetDefaultLogConfig()

	cfg.Filename = LogFileName
	cfg.Directory = dir


	BaseInitWithConfig(cfg)
}

var inited = false
func BaseInitWithConfig(cfg *Config) {
	if inited {
		return
	}

	inited = true

	os.Mkdir(cfg.Directory, 0666)

	myLogger = Configure(*cfg, StructLog)

	//MaxSize = 1024 * 1024 * 16
}

func GetDefaultLogConfig() *Config {

	var logfilename = crazyos.GetAppName()
	var logDir = ""

	logDir = crazyos.GetExecFolder()
	logDir = path.Join(logDir, "logs")

	var cfg = Config { ConsoleLoggingEnabled:true, EncodeLogsAsJson:false,
		FileLoggingEnabled:true, Directory:logDir, Filename:logfilename,
		MaxBackups:100,MaxAge:7, MaxSize:32}


	return &cfg
}

func Close() {

}

func Cheat(tmpdir string) {

	// 设置输出路径
	//logDir = &dir
}

func getCurPath() string {
	return crazyos.GetExecFolder()
}

// Flush flushes all pending log I/O.
func Flush() {

}

func sprintf(args ...interface{}) string {
	var str = fmt.Sprint(args...)
	str = strings.Trim(str, "\r\n")
	return str
}

// Info logs to the INFO log.
// Arguments are handled in the manner of fmt.Print; a newline is appended if missing.
func Info(args ...interface{}) {
	myLogger.Infoln(sprintf(args...))
}

// InfoDepth acts as Info but uses depth to determine which call frame to log.
// InfoDepth(0, "msg") is the same as Info("msg").
func InfoDepth(depth int, args ...interface{}) {
	Info(args...)
}

// Infoln logs to the INFO log.
// Arguments are handled in the manner of fmt.Println; a newline is appended if missing.
func Infoln(args ...interface{}) {
	Info(args...)
}

// Infof logs to the INFO log.
// Arguments are handled in the manner of fmt.Printf; a newline is appended if missing.
func Infof(format string, args ...interface{}) {
	myLogger.Infoln(fmt.Sprintf(format, args...))
}

// Warning logs to the WARNING and INFO logs.
// Arguments are handled in the manner of fmt.Print; a newline is appended if missing.
func Warning(args ...interface{}) {
	myLogger.Warnln(sprintf(args...))
}

// WarningDepth acts as Warning but uses depth to determine which call frame to log.
// WarningDepth(0, "msg") is the same as Warning("msg").
func WarningDepth(depth int, args ...interface{}) {
	Warning(args...)
}

// Warningln logs to the WARNING and INFO logs.
// Arguments are handled in the manner of fmt.Println; a newline is appended if missing.
func Warningln(args ...interface{}) {
	Warning(args...)
}

// Warningf logs to the WARNING and INFO logs.
// Arguments are handled in the manner of fmt.Printf; a newline is appended if missing.
func Warningf(format string, args ...interface{}) {
	myLogger.Warnln(fmt.Sprintf(format, args...))
}

// Error logs to the ERROR, WARNING, and INFO logs.
// Arguments are handled in the manner of fmt.Print; a newline is appended if missing.
func Error(args ...interface{}) {
	myLogger.Errorln(sprintf(args...))
}

// ErrorDepth acts as Error but uses depth to determine which call frame to log.
// ErrorDepth(0, "msg") is the same as Error("msg").
func ErrorDepth(depth int, args ...interface{}) {
	Error(args...)
}

// Errorln logs to the ERROR, WARNING, and INFO logs.
// Arguments are handled in the manner of fmt.Println; a newline is appended if missing.
func Errorln(args ...interface{}) {
	Error(args...)
}

// Errorf logs to the ERROR, WARNING, and INFO logs.
// Arguments are handled in the manner of fmt.Printf; a newline is appended if missing.
func Errorf(format string, args ...interface{}) {
	myLogger.Errorln(fmt.Sprintf(format, args...))
}

// Fatal logs to the FATAL, ERROR, WARNING, and INFO logs,
// including a stack trace of all running goroutines, then calls os.Exit(255).
// Arguments are handled in the manner of fmt.Print; a newline is appended if missing.
func Fatal(args ...interface{}) {
	myLogger.Errorln(sprintf(args...))
}

// FatalDepth acts as Fatal but uses depth to determine which call frame to log.
// FatalDepth(0, "msg") is the same as Fatal("msg").
func FatalDepth(depth int, args ...interface{}) {
	Fatal(args...)
}

// Fatalln logs to the FATAL, ERROR, WARNING, and INFO logs,
// including a stack trace of all running goroutines, then calls os.Exit(255).
// Arguments are handled in the manner of fmt.Println; a newline is appended if missing.
func Fatalln(args ...interface{}) {
	Fatal(args...)
}

// Fatalf logs to the FATAL, ERROR, WARNING, and INFO logs,
// including a stack trace of all running goroutines, then calls os.Exit(255).
// Arguments are handled in the manner of fmt.Printf; a newline is appended if missing.
func Fatalf(format string, args ...interface{}) {
	myLogger.Errorln(fmt.Sprintf(format, args...))
}
