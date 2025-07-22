package log // 定义 log 包，供其他地方引入使用
import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"gopkg.in/natefinch/lumberjack.v2"
	"jingzhe-bg/main/global"
	"os"
	"strings"
	"time"
)

var lastTime string
var lastUnix int64

func cachedTime() string {
	now := time.Now().Unix()
	if now != lastUnix { // 每秒更新一次（精度可调）
		lastUnix = now
		lastTime = time.Now().Format("2006-01-02T15:04:05.000-0700")
	}
	return lastTime
}

// 自定义日志级别编码器
func customLevelEncoder(l zapcore.Level, enc zapcore.PrimitiveArrayEncoder) {
	enc.AppendString("[JZ-DEBUG] [" + strings.ToUpper(l.String()) + "]\t" + cachedTime())
}

// 初始化支持日志切割的 Zap Logger
func InitLogger() {
	// 1. 配置日志切割（文件输出）
	fileWriter := zapcore.AddSync(&lumberjack.Logger{
		Filename:   "logs/app.log", // 日志文件路径
		MaxSize:    100,            // 每个日志文件的最大大小（MB）
		MaxBackups: 10,             // 保留旧日志文件的最大数量
		MaxAge:     30,             // 保留旧日志文件的最大天数
		Compress:   true,           // 是否压缩旧日志文件
	})

	// 2. 配置控制台输出
	consoleWriter := zapcore.AddSync(os.Stdout)

	// 3. 配置编码器（统一格式）
	encoderConfig := zapcore.EncoderConfig{
		MessageKey:     "msg",
		TimeKey:        "",
		LevelKey:       "level",
		NameKey:        "logger",
		CallerKey:      "caller",
		FunctionKey:    zapcore.OmitKey,
		StacktraceKey:  "stacktrace",
		EncodeLevel:    customLevelEncoder, // 使用自定义级别编码器
		LineEnding:     zapcore.DefaultLineEnding,
		EncodeTime:     zapcore.ISO8601TimeEncoder,
		EncodeDuration: zapcore.StringDurationEncoder,
		EncodeCaller:   zapcore.ShortCallerEncoder,
	}

	// 4. 创建多输出核心
	core := zapcore.NewTee(
		// 文件核心（JSON 格式）
		zapcore.NewCore(
			zapcore.NewJSONEncoder(encoderConfig),
			fileWriter,
			zap.InfoLevel,
		),
		// 控制台核心（可读性更高的 Console 格式）
		zapcore.NewCore(
			zapcore.NewConsoleEncoder(encoderConfig),
			consoleWriter,
			zap.InfoLevel,
		),
	)

	// 5. 创建 Logger
	global.GVA_LOGGER = zap.New(core, zap.AddCaller())
}
