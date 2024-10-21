package gzap

import (
	"os"
	"strings"
	"sync"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"gopkg.in/natefinch/lumberjack.v2"
)

type zapLogConfig struct {
	logger     *zap.Logger
	loggerOnce sync.Once
	zapSyncer  []zapcore.WriteSyncer
	zapConf    zap.Config
	CallerSkip int
}

// 默认的配置
var zapLogCfg = &zapLogConfig{
	CallerSkip: 1,
	logger:     nil,         // 返回的zap
	loggerOnce: sync.Once{}, // 懒加载

	zapSyncer: []zapcore.WriteSyncer{zapcore.AddSync(os.Stdout)}, //默认输出到控制台

	zapConf: zap.Config{
		Level:       zap.NewAtomicLevelAt(zap.InfoLevel), // 日志级别
		Development: true,                                // 开发模式，堆栈跟踪
		Encoding:    "json",                              // 输出格式 console 或 json

		EncoderConfig: zapcore.EncoderConfig{
			TimeKey:        "time",
			LevelKey:       "level",
			NameKey:        "name",
			CallerKey:      "line",
			MessageKey:     "msg",
			FunctionKey:    "func",
			StacktraceKey:  "stacktrace",
			LineEnding:     zapcore.DefaultLineEnding,
			EncodeLevel:    zapcore.LowercaseLevelEncoder,                          // 小写编码器
			EncodeTime:     zapcore.TimeEncoderOfLayout("2006-01-02 15:04:05.000"), // 自定义 时间格式
			EncodeDuration: zapcore.SecondsDurationEncoder,
			EncodeCaller:   zapcore.FullCallerEncoder, // 全路径编码器
		}, // 编码器配置
	},
}

/*
 * 默认Info级别，如果要自定义级别，请在此之前设置SetZapCfg(gzap.ZapLevel("debug"))
 * gzap.Zap().Info("hello world", zap.String("name", "zhangsir"), zap.Any("age", 18))
 */
func Zap() *zap.Logger {
	zapLazyInit()
	return zapLogCfg.logger
}

/*
 * 带语法糖的zap
 * gzap.Sap().Infow("hello world", "name", "zhangsir", "age", 18)
 */
func Sap() *zap.SugaredLogger {
	zapLazyInit()
	return zapLogCfg.logger.Sugar()
}

// 懒加载
func zapLazyInit() {
	zapLogCfg.loggerOnce.Do(func() {

		//日志格式
		encoder := zapcore.NewJSONEncoder(zapLogCfg.zapConf.EncoderConfig)

		switch zapLogCfg.zapConf.Encoding {
		case "console":
			encoder = zapcore.NewConsoleEncoder(zapLogCfg.zapConf.EncoderConfig)
		}

		//日志输出
		syncer := zapcore.NewMultiWriteSyncer(zapLogCfg.zapSyncer...)

		zapCore := zapcore.NewCore(encoder, syncer, zapLogCfg.zapConf.Level)

		//预设字段
		if len(zapLogCfg.zapConf.InitialFields) > 0 {
			fields := []zap.Field{}

			for k, v := range zapLogCfg.zapConf.InitialFields {
				fields = append(fields, zap.Any(k, v))
			}

			zapLogCfg.logger = zap.New(zapCore, zap.AddCaller(), zap.AddCallerSkip(zapLogCfg.CallerSkip), zap.Fields(fields...))
			return
		}

		zapLogCfg.logger = zap.New(zapCore, zap.AddCaller(), zap.AddCallerSkip(zapLogCfg.CallerSkip))
	})
}

/*
 * 自定义配置的时候需要（注意要在使用日志之前设置）
 * gzap.SetZapCfg(gzap.ZapLevel("debug"), gzap.ZapOutFile("./log/test.log", gzap.ZapOutFileMaxSize(128), gzap.ZapOutFileMaxAge(7)))
 */
func SetZapCfg(opts ...func(*zapLogConfig)) {
	for _, opt := range opts {
		opt(zapLogCfg)
	}
}

// 设置预设字段
func SetPresetFields(fields map[string]any) func(*zapLogConfig) {
	return func(z *zapLogConfig) {
		z.zapConf.InitialFields = fields
	}
}

// 设置日志等级
func ZapLevel(level string) func(*zapLogConfig) {
	level = strings.ToLower(level)

	return func(z *zapLogConfig) {
		if z.logger != nil {
			z.logger.Error("logger already init before SetLevel")
		}

		switch level {
		case "debug":
			z.zapConf.Level = zap.NewAtomicLevelAt(zap.DebugLevel)
		case "info":
			z.zapConf.Level = zap.NewAtomicLevelAt(zap.InfoLevel)
		case "warn":
			z.zapConf.Level = zap.NewAtomicLevelAt(zap.WarnLevel)
		case "error":
			z.zapConf.Level = zap.NewAtomicLevelAt(zap.ErrorLevel)
		case "dpanic":
			z.zapConf.Level = zap.NewAtomicLevelAt(zap.DPanicLevel)
		case "panic":
			z.zapConf.Level = zap.NewAtomicLevelAt(zap.PanicLevel)
		case "fatal":
			z.zapConf.Level = zap.NewAtomicLevelAt(zap.FatalLevel)
		}
	}
}

// 设置输出位置，带归档
func ZapOutFile(path string, opts ...func(*lumberjack.Logger)) func(*zapLogConfig) {

	return func(z *zapLogConfig) {
		if z.logger != nil {
			z.logger.Error("logger already init before SetOutFile")
		}

		lumberjackLogger := &lumberjack.Logger{
			Filename:   path,
			MaxSize:    128, // 文件最大MB
			MaxAge:     30,  // 保留30天
			MaxBackups: 30,  // 保留30个文件
			LocalTime:  true,
			Compress:   false,
		}

		for _, opt := range opts {
			opt(lumberjackLogger)
		}

		syncer := zapcore.AddSync(lumberjackLogger)

		z.zapSyncer = append(z.zapSyncer, syncer)
	}
}

// 进行切割之前,日志文件的最大大小(MB为单位)
func ZapOutFileMaxSize(maxSize int) func(ll *lumberjack.Logger) {
	return func(ll *lumberjack.Logger) {
		ll.MaxSize = maxSize
	}
}

// 保留旧文件的最大天数
func ZapOutFileMaxAge(maxAge int) func(ll *lumberjack.Logger) {
	return func(ll *lumberjack.Logger) {
		ll.MaxAge = maxAge
	}
}

// 保留旧文件的最大个数
func ZapOutFileMaxBackups(maxBackups int) func(ll *lumberjack.Logger) {
	return func(ll *lumberjack.Logger) {
		ll.MaxBackups = maxBackups
	}
}

// 设置堆栈跟踪
func ZapDevelopment(dev bool) func(*zapLogConfig) {
	return func(z *zapLogConfig) {
		if z.logger != nil {
			z.logger.Error("logger already init before SetDevelopment")
		}

		z.zapConf.Development = dev
	}
}

func ZapEncodering(encodering string) func(*zapLogConfig) {
	return func(z *zapLogConfig) {
		if z.logger != nil {
			z.logger.Error("logger already init before SetEncoder")
		}

		z.zapConf.Encoding = encodering
	}
}

// 设置callserskip
func ZapCallerSkip(skip int) func(*zapLogConfig) {
	return func(z *zapLogConfig) {
		if z.logger != nil {
			z.logger.Error("logger already init before SetCallerSkip")
		}

		z.CallerSkip = skip
	}
}

// 配置深度定制
func ZapConf(conf zap.Config) func(*zapLogConfig) {
	return func(z *zapLogConfig) {
		if z.logger != nil {
			z.logger.Error("logger already init before SetConf")
		}

		z.zapConf = conf
	}
}
