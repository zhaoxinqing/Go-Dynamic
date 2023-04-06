package logger

import (
	"fmt"
	"os"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"gopkg.in/natefinch/lumberjack.v2"
)

const loggerKey = iota

// Logger 全局 Logger 对象
var Logger *zap.Logger

type LoggerOptions struct {
	Filepath string `json:"file_path" mapstructure:"file_path"`
	Channel  string `json:"channel" mapstructure:"channel"`
}

type Log interface {
	Channel(filename string) *zap.Logger
	WithContext(ctx *gin.Context) *zap.Logger
}

var _ Log

func Channel(channel string) *zap.Logger {
	//if userID := ctx.Value("userID"); userID != 0 {
	//	Logger.With(zap.Any("userID", userID))
	//}

	return Logger.With(zap.Any("channel", channel))
}

func NewContext(ctx *gin.Context, fields ...zapcore.Field) *zap.Logger {
	ctx.Set(strconv.Itoa(loggerKey), WithContext(ctx).With(fields...))
	return Logger
}

func WithContext(ctx *gin.Context) *zap.Logger {
	if ctx == nil {
		return Logger
	}
	l, _ := ctx.Get(strconv.Itoa(loggerKey))
	ctxLogger, ok := l.(*zap.Logger)
	if ok {
		return ctxLogger
	}
	return Logger
}

// NewLogger 日志初始化
func NewLogger(o *LoggerOptions) *zap.Logger {
	// 获取日志写入介质
	writeSyncer := getLogWriter(fmt.Sprintf("%s/logs.log", o.Filepath))
	logLevel := new(zapcore.Level)
	// 初始化 core
	core := zapcore.NewCore(getEncoder(), writeSyncer, logLevel)
	// 初始化 Logger
	Logger = zap.New(core,

		zap.AddCaller(),      // 调用文件和行号，内部使用 runtime.Caller
		zap.AddCallerSkip(1), // 封装了一层，调用文件去除一层(runtime.Caller(1))
		// zap.AddStacktrace(zap.ErrorLevel), // Error 时才会显示 stacktrace
	)
	// 将自定义的 logger 替换为全局的 logger
	// zap.L().Fatal() 调用时，就会使用我们自定的 Logger
	zap.ReplaceGlobals(Logger)
	return Logger
}

// getEncoder 设置日志存储格式
func getEncoder() zapcore.Encoder {
	// 日志格式规则
	encoderConfig := zapcore.EncoderConfig{
		TimeKey:  "time",
		LevelKey: "level",
		//NameKey:        "logger",
		CallerKey:      "caller", // 代码调用，如 paginator/paginator.go:148
		FunctionKey:    zapcore.OmitKey,
		MessageKey:     "message",
		StacktraceKey:  "stacktrace",
		LineEnding:     zapcore.DefaultLineEnding,      // 每行日志的结尾添加 "\n"
		EncodeLevel:    zapcore.CapitalLevelEncoder,    // 日志级别名称大写，如 ERROR、INFO
		EncodeTime:     customTimeEncoder,              // 时间格式，我们自定义为 2006-01-02 15:04:05
		EncodeDuration: zapcore.SecondsDurationEncoder, // 执行时间，以秒为单位
		EncodeCaller:   zapcore.ShortCallerEncoder,     // Caller 短格式，如：types/converter.go:17，长格式为绝对路径
	}

	logFormat := viper.GetString("logger.format")
	if logFormat == "json" {
		return zapcore.NewJSONEncoder(encoderConfig)
	} else if logFormat == "log" {
		encoderConfig.EncodeLevel = zapcore.CapitalColorLevelEncoder
		return zapcore.NewConsoleEncoder(encoderConfig)
	}
	return zapcore.NewConsoleEncoder(encoderConfig)

}

// customTimeEncoder 自定义友好的时间格式
func customTimeEncoder(t time.Time, enc zapcore.PrimitiveArrayEncoder) {
	enc.AppendString(t.Format("2006-01-02 15:04:05"))
}

// getLogWriter 日志记录介质
func getLogWriter(filename string) zapcore.WriteSyncer {

	// 如果配置了按照日期记录日志文件
	//if viper.GetString("logger.driver") == "daily" {
	//	logname := time.Now().UTC().Format("2006-01-02.log")
	//	filename = strings.ReplaceAll(filename, "logs.log", logname)
	//}
	// 滚动日志
	lumberJackLogger := &lumberjack.Logger{
		Filename: filename,
		MaxSize:  viper.GetInt("logger.max_size"),
	}

	//每日零点定时分割日志
	go func() {
		for {
			nowTime := time.Now().UTC()
			nowTimeStr := nowTime.Format("2006-01-02")
			t2, _ := time.ParseInLocation("2006-01-02", nowTimeStr, time.UTC)
			// 第二天零点时间戳
			next := t2.AddDate(0, 0, 1)
			after := next.UnixNano() - nowTime.UnixNano() - 1
			<-time.After(time.Duration(after) * time.Nanosecond)
			lumberJackLogger.Rotate()
		}
	}()

	return zapcore.NewMultiWriteSyncer(zapcore.AddSync(os.Stdout), zapcore.AddSync(lumberJackLogger))
}

func LogOptions(path string) *LoggerOptions {
	return &LoggerOptions{Filepath: path}
}
