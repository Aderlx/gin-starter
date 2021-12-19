package log

import (
	"gin-starter/global"
	"os"

	"github.com/natefinch/lumberjack"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

func InitLog() *zap.Logger {
	var Logger *zap.Logger

	var logFilePath = global.Gol_CONFIG.GetString("log.logFilePath")

	var maxSize = global.Gol_CONFIG.GetInt64("log.maxSize")

	var maxBackups = global.Gol_CONFIG.GetInt64("log.maxBackups")

	var maxAge = global.Gol_CONFIG.GetInt64("log.maxAge")

	var compress = global.Gol_CONFIG.GetBool("log.compress")

	var formatJson = global.Gol_CONFIG.GetBool("log.formatJson")

	hook := lumberjack.Logger{
		Filename:   logFilePath,     // 日志文件路径
		MaxSize:    int(maxSize),    // 每个日志文件保存的最大尺寸 单位：M
		MaxBackups: int(maxBackups), // 日志文件最多保存多少个备份
		MaxAge:     int(maxAge),     // 文件最多保存多少天
		Compress:   compress,        // 是否压缩
	}

	encoderConfig := zapcore.EncoderConfig{
		TimeKey:        "time",
		LevelKey:       "level",
		NameKey:        "logger",
		CallerKey:      "linenum",
		MessageKey:     "msg",
		StacktraceKey:  "stacktrace",
		LineEnding:     zapcore.DefaultLineEnding,
		EncodeLevel:    zapcore.LowercaseLevelEncoder,  // 小写编码器
		EncodeTime:     zapcore.ISO8601TimeEncoder,     // ISO8601 UTC 时间格式
		EncodeDuration: zapcore.SecondsDurationEncoder, //
		EncodeCaller:   zapcore.FullCallerEncoder,      // 全路径编码器
		EncodeName:     zapcore.FullNameEncoder,
	}

	// 设置日志级别
	atomicLevel := zap.NewAtomicLevel()
	atomicLevel.SetLevel(zap.InfoLevel)

	// 日志输出格式
	var encoder zapcore.Encoder

	// 是否是json
	if formatJson {
		encoder = zapcore.NewJSONEncoder(encoderConfig)
	} else {
		encoder = zapcore.NewConsoleEncoder(encoderConfig)
	}

	core := zapcore.NewCore(
		encoder,
		// zapcore.NewJSONEncoder(encoderConfig),                                           // 编码器配置
		zapcore.NewMultiWriteSyncer(zapcore.AddSync(os.Stdout), zapcore.AddSync(&hook)), // 打印到控制台和文件
		atomicLevel, // 日志级别
	)

	// 开启开发模式，堆栈跟踪
	caller := zap.AddCaller()
	// 开启文件及行号
	development := zap.Development()
	// 设置初始化字段
	// filed := zap.Fields(zap.String("serviceName", "serviceName"))
	// 构造日志
	Logger = zap.New(core, caller, development)

	return Logger

}
