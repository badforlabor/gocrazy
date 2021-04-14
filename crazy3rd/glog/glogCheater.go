package glog

import (
	"fmt"
	"os"
	"os/exec"
	"path"
	"path/filepath"
	"strings"
)

var dir string = ""

var myLogger *Logger

var LogFileName = "example.log"
var StructLog = false

func BaseInit() {
	dir := getCurPath()
	dir = path.Join(dir, "logs")
	os.Mkdir(dir, os.ModeDir)

	var cfg = Config{ConsoleLoggingEnabled:true, EncodeLogsAsJson:false,
		FileLoggingEnabled:true, Directory:dir, Filename:LogFileName,
		MaxBackups:100,MaxAge:7, MaxSize:32}
	myLogger = Configure(cfg, StructLog)

	Cheat(dir)

	//MaxSize = 1024 * 1024 * 16
}
func Close() {

}

func Cheat(tmpdir string) {
	dir = tmpdir
	// 设置输出路径
	//logDir = &dir
}

func getCurPath() string {
	fmt.Println(os.Args[0])
	file, _ := exec.LookPath(os.Args[0])

	if len(file) == 0 {
		file = os.Args[0]
	}

	//得到全路径，比如在windows下E:\\golang\\test\\a.exe
	path, _ := filepath.Abs(file)

	//将全路径用\\分割，得到4段，①E: ②golang ③test ④a.exe
	splitstring := strings.Split(path, string(os.PathSeparator))

	splitstring = splitstring[:len(splitstring) - 1]

	var rst = strings.Join(splitstring, string(os.PathSeparator))

	return rst
}

// Flush flushes all pending log I/O.
func Flush() {

}

func sprintf(args ...interface{}) string {
	var str = fmt.Sprintln(args...)
	return str[:len(str)-1]
}

// Info logs to the INFO log.
// Arguments are handled in the manner of fmt.Print; a newline is appended if missing.
func Info(args ...interface{}) {
	myLogger.Infoln(sprintf(args...))
}

// InfoDepth acts as Info but uses depth to determine which call frame to log.
// InfoDepth(0, "msg") is the same as Info("msg").
func InfoDepth(depth int, args ...interface{}) {
	myLogger.Infoln(sprintf(args...))
}

// Infoln logs to the INFO log.
// Arguments are handled in the manner of fmt.Println; a newline is appended if missing.
func Infoln(args ...interface{}) {
	myLogger.Infoln(sprintf(args...))
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
	myLogger.Warnln(sprintf(args...))
}

// Warningln logs to the WARNING and INFO logs.
// Arguments are handled in the manner of fmt.Println; a newline is appended if missing.
func Warningln(args ...interface{}) {
	myLogger.Warnln(sprintf(args...))
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
	myLogger.Errorln(sprintf(args...))
}

// Errorln logs to the ERROR, WARNING, and INFO logs.
// Arguments are handled in the manner of fmt.Println; a newline is appended if missing.
func Errorln(args ...interface{}) {
	myLogger.Errorln(sprintf(args...))
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
	myLogger.Errorln(sprintf(args...))
}

// Fatalln logs to the FATAL, ERROR, WARNING, and INFO logs,
// including a stack trace of all running goroutines, then calls os.Exit(255).
// Arguments are handled in the manner of fmt.Println; a newline is appended if missing.
func Fatalln(args ...interface{}) {
	myLogger.Errorln(sprintf(args...))
}

// Fatalf logs to the FATAL, ERROR, WARNING, and INFO logs,
// including a stack trace of all running goroutines, then calls os.Exit(255).
// Arguments are handled in the manner of fmt.Printf; a newline is appended if missing.
func Fatalf(format string, args ...interface{}) {
	myLogger.Errorln(fmt.Sprintf(format, args...))
}
