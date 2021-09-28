/**
 * Auth :   liubo
 * Date :   2021/9/28 16:17
 * Comment:
 */

package crazylog

import (
	"fmt"
	"github.com/badforlabor/gocrazy/crazyio"
	"github.com/badforlabor/gocrazy/crazyos"
	rotatelogs "github.com/lestrrat-go/file-rotatelogs"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"io"
	"os"
	"path"
	"time"
)

//
//  DefaultLog
//  @Description: 提供一个默认的日志模板
//  @return *LogConfig
//  @return io.Writer
//
func DefaultLog() (*LogConfig, io.Writer){

	var logConfig = GetDefaultLogConfig()
	var logWriter = Configure(logConfig)

	// 删掉过期文件
	crazyio.DeleteFiles(logConfig.Directory, "^" + logConfig.Filename + `.*\.log`,
		time.Duration(logConfig.MaxAge) * time.Hour * 24)

	// 修改zerolog的
	log.Logger = log.Output(logWriter)

	// 使用zero格式打印日志
	log.Info().
		Bool("fileLogging", logConfig.FileLoggingEnabled).
		Bool("jsonLogOutput", logConfig.EncodeLogsAsJson).
		Str("logDirectory", logConfig.Directory).
		Str("fileName", logConfig.Filename).
		Int("maxSizeMB", logConfig.MaxSize).
		Msg("logging configured")

	// 使用常用的方式打印日志
	var globalLog = NewCategoryLogDefault("", logWriter)
	globalLog.WithTimestamp = "2006/01/02 15:04:05"

	globalLog.Println("123")
	globalLog.Printf("now:%s\n", time.Now().String())

	return logConfig, logWriter
}


// Configuration for logging
type LogConfig struct {
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
	// MaxAge the max age in days to keep a logfile
	MaxAge int
}

func GetDefaultLogConfig() *LogConfig {

	var logfilename = crazyos.GetAppName()
	var logDir = ""

	logDir = crazyos.GetExecFolder()
	logDir = path.Join(logDir, "logs")

	var cfg = LogConfig { ConsoleLoggingEnabled:true, EncodeLogsAsJson:false,
		FileLoggingEnabled:true, Directory:logDir, Filename:logfilename,
		MaxSize:32, MaxAge:7}


	return &cfg
}


var svcIsWindowsService func()(bool, error)

func Configure(config *LogConfig) io.Writer {
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

	return mw
}

func newRollingFile(config *LogConfig) io.Writer {
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
}
