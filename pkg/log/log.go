package log

import (
	"fmt"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"gopkg.in/natefinch/lumberjack.v2"
	"os"
	"time"
)

func TimeEncoder(t time.Time, enc zapcore.PrimitiveArrayEncoder) {
	enc.AppendString(t.Format("2006-01-02 15:04:05.000"))
}

var logger *zap.Logger

func Debug(tag string, msg string) {
	logger.Debug(msg, zap.String("tag", tag))
}

func Info(tag string, msg string) {
	logger.Info(msg, zap.String("tag", tag))
}

func Error(tag string, msg string) {
	logger.Error(msg, zap.String("tag", tag))
}

func Warn(tag string, msg string) {
	logger.Warn(msg, zap.String("tag", tag))
}

func InitLogger(logpath string, loglevel string, maxSize int, maxBackups int, maxAge int, compress bool, serviceName string) {
	var level zapcore.Level
	switch loglevel {
	case "debug":
		level = zap.DebugLevel
	case "info":
		level = zap.InfoLevel
	case "error":
		level = zap.ErrorLevel
	default:
		level = zap.InfoLevel
	}
	logPath := fmt.Sprintf("%s/%s.log", logpath, loglevel)
	// 日志分割
	hook := lumberjack.Logger{
		Filename:   logPath,    // 日志文件路径，默认 os.TempDir()
		MaxSize:    maxSize,    // 每个日志文件保存10M，默认 100M
		MaxBackups: maxBackups, // 保留30个备份，默认不限
		MaxAge:     maxAge,     // 保留7天，默认不限
		Compress:   compress,   // 是否压缩，默认不压缩
	}
	//write := zapcore.AddSync(&hook)
	// 设置日志级别
	// debug 可以打印出 info debug warn
	// info  级别可以打印 warn info
	// warn  只能打印 warn
	// debug->info->warn->error

	encoderConfig := zapcore.EncoderConfig{
		TimeKey:        "time",
		LevelKey:       "level",
		NameKey:        "logger",
		CallerKey:      "linenum",
		MessageKey:     "msg",
		StacktraceKey:  "stacktrace",
		LineEnding:     zapcore.DefaultLineEnding,
		EncodeLevel:    zapcore.LowercaseLevelEncoder,  // 小写编码器
		EncodeTime:     TimeEncoder,                    // ISO8601 UTC 时间格式
		EncodeDuration: zapcore.SecondsDurationEncoder, //
		EncodeCaller:   zapcore.ShortCallerEncoder,     // 全路径编码器 zapcore.FullCallerEncoder
		EncodeName:     zapcore.FullNameEncoder,
	}
	// 设置日志级别
	atomicLevel := zap.NewAtomicLevel()
	atomicLevel.SetLevel(level)
	core := zapcore.NewCore(
		//zapcore.NewConsoleEncoder(encoderConfig),
		zapcore.NewJSONEncoder(encoderConfig),
		zapcore.NewMultiWriteSyncer(zapcore.AddSync(&hook), zapcore.AddSync(os.Stdout)), // 打印到控制台和文件
		//write,
		level,
	)
	// 开启开发模式，堆栈跟踪
	caller := zap.AddCaller()
	// 开启文件及行号
	development := zap.Development()
	// 设置初始化字段,如：添加一个服务器名称
	filed := zap.Fields(zap.String("service", serviceName))
	// 构造日志
	logger = zap.New(core, caller, development, filed)
}
