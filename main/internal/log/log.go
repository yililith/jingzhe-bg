package log // 定义 log 包，供其他地方引入使用

import (
	"go.uber.org/zap"                  // 引入 zap 日志库主包
	"go.uber.org/zap/zapcore"          // 引入 zapcore，控制日志编码、级别等底层实现
	"gopkg.in/natefinch/lumberjack.v2" // 引入 lumberjack，用于日志文件自动切割
	"os"                               // 用于输出到标准输出（控制台）
)

var (
	Logger *zap.Logger // 全局可用的 zap.Logger 实例
)

func InitLog() {
	// 1. 配置日志切割
	lumberjackLogger := &lumberjack.Logger{
		Filename:   "logs/app.log", // 指定日志文件路径和名称
		MaxSize:    100,            // 单个日志文件最大 100MB
		MaxBackups: 30,             // 最多保留 30 个旧日志文件
		MaxAge:     30,             // 日志最多保留 30 天
		Compress:   true,           // 超出后旧日志是否进行 gzip 压缩
	}

	// 2. 配置编码器（即日志输出格式定义）
	encoderConfig := zapcore.EncoderConfig{
		TimeKey:        "time",                        // 时间字段名称
		LevelKey:       "level",                       // 日志级别字段名称
		NameKey:        "logger",                      // 日志器名称字段
		CallerKey:      "caller",                      // 调用源文件位置字段
		FunctionKey:    zapcore.OmitKey,               // 忽略函数名称字段
		MessageKey:     "message",                     // 日志正文字段
		StacktraceKey:  "stacktrace",                  // 堆栈信息字段（仅 Error+ 显示）
		LineEnding:     zapcore.DefaultLineEnding,     // 换行符定义
		EncodeLevel:    zapcore.CapitalLevelEncoder,   // 日志级别大写输出（INFO、ERROR）
		EncodeTime:     zapcore.ISO8601TimeEncoder,    // ISO8601 格式时间（如 2025-06-25T14:00:00Z）
		EncodeDuration: zapcore.StringDurationEncoder, // 输出 duration 为字符串格式
		EncodeCaller:   zapcore.ShortCallerEncoder,    // 调用位置为简写格式（如 main.go:42）
	}

	// 3. 创建多个输出端
	// 文件输出(JSON格式)
	fileEncoder := zapcore.NewJSONEncoder(encoderConfig) // 文件输出使用 JSON 编码
	fileWriter := zapcore.AddSync(lumberjackLogger)      // 输出目标为 lumberjack 的文件处理器
	fileLevel := zap.LevelEnablerFunc(func(lvl zapcore.Level) bool {
		return lvl >= zapcore.DebugLevel // 文件记录所有级别（Debug及以上）
	})

	// 控制台输出(美化格式)
	consoleEncoder := zapcore.NewConsoleEncoder(zapcore.EncoderConfig{
		TimeKey:        "ts",                                               // 控制台时间字段使用 "ts"
		LevelKey:       "level",                                            // 日志级别字段
		NameKey:        "logger",                                           // 日志器字段
		CallerKey:      "caller",                                           // 调用位置字段
		FunctionKey:    zapcore.OmitKey,                                    // 不输出函数名
		MessageKey:     "msg",                                              // 日志正文字段为 "msg"
		StacktraceKey:  "stacktrace",                                       // 堆栈字段
		LineEnding:     zapcore.DefaultLineEnding,                          // 默认换行符
		EncodeLevel:    zapcore.CapitalColorLevelEncoder,                   // 彩色大写日志级别（INFO、ERROR）
		EncodeTime:     zapcore.TimeEncoderOfLayout("2006-01-02 15:04:05"), // 自定义时间格式（常规时间格式）
		EncodeDuration: zapcore.StringDurationEncoder,                      // duration 输出为字符串
		EncodeCaller:   zapcore.ShortCallerEncoder,                         // 简写调用位置
	})
	consoleWriter := zapcore.AddSync(os.Stdout) // 控制台输出目标为标准输出
	consoleLevel := zap.LevelEnablerFunc(func(lvl zapcore.Level) bool {
		return lvl >= zapcore.InfoLevel // 控制台只记录 Info 及以上级别日志
	})

	// 4. 创建核心（zapcore.Core 组合输出）
	core := zapcore.NewTee( // 将多个 Core 组合在一起
		zapcore.NewCore(fileEncoder, fileWriter, fileLevel),          // 文件日志核心
		zapcore.NewCore(consoleEncoder, consoleWriter, consoleLevel), // 控制台日志核心
	)

	// 5. 创建 Logger 实例
	Logger = zap.New(core, // 使用组合后的 Core
		zap.AddCaller(),                       // 添加调用者信息（caller字段）
		zap.AddCallerSkip(1),                  // 跳过 1 层调用栈，防止记录 log 包内部
		zap.AddStacktrace(zapcore.ErrorLevel), // 错误级别及以上添加堆栈信息
		zap.WithCaller(true),                  // 明确启用 caller 输出（加上保险）
		//zap.AddCallerSkip(5), // （注释掉的）可调整跳过层级数，适用于更复杂封装场景
	)
}
