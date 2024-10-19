package gzap_test

import (
	"testing"

	"github.com/Flyingmn/gzap"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

func TestDebug(t *testing.T) {
	//默认info,所以先设置级别
	gzap.SetZapCfg(gzap.ZapLevel("debug"))

	defer gzap.Syncw()
	defer gzap.Sync()
	gzap.Debug("hello world", zap.String("name", "zhangsan"), zap.Int("age", 18)) //高性能

	gzap.Debugw("hello world", "name", "zhangsan", "age", 18) //性能普通

	gzap.Debugf("hello world; name:%s; age:%d", "zhangsan", 18) // printf
}

func TestInfo(t *testing.T) {
	defer gzap.Syncw()
	defer gzap.Sync()
	gzap.Info("hello world", zap.String("name", "zhangsan"), zap.Int("age", 18)) //高性能

	gzap.Infow("hello world", "name", "zhangsan", "age", 18) //性能普通

	gzap.Infof("hello world; name:%s; age:%d", "zhangsan", 18) // printf
}

func TestWarn(t *testing.T) {
	defer gzap.Syncw()
	defer gzap.Sync()
	gzap.Warn("hello world", zap.String("name", "zhangsan"), zap.Int("age", 18)) //高性能

	gzap.Warnw("hello world", "name", "zhangsan", "age", 18) //性能普通

	gzap.Warnf("hello world; name:%s; age:%d", "zhangsan", 18) // printf
}

func TestError(t *testing.T) {
	defer gzap.Syncw()
	defer gzap.Sync()
	gzap.Error("hello world", zap.String("name", "zhangsan"), zap.Int("age", 18)) //高性能

	gzap.Errorw("hello world", "name", "zhangsan", "age", 18) //性能普通

	gzap.Errorf("hello world; name:%s; age:%d", "zhangsan", 18) // printf
}

func TestDPanic(t *testing.T) {
	defer gzap.Syncw()
	defer gzap.Sync()
	gzap.DPanic("hello world", zap.String("name", "zhangsan"), zap.Int("age", 18)) //高性能

	gzap.DPanicw("hello world", "name", "zhangsan", "age", 18) //性能普通

	gzap.DPanicf("hello world; name:%s; age:%d", "zhangsan", 18) // printf
}

// func TestPanic(t *testing.T) {
// 	defer gzap.Syncw()
// 	defer gzap.Sync()
// 	gzap.Panic("hello world", zap.String("name", "zhangsan"), zap.Int("age", 18)) //高性能

// 	//不会执行
// 	gzap.Panicw("hello world", "name", "zhangsan", "age", 18) //性能普通
// 	//不会执行
// 	gzap.Panicf("hello world; name:%s; age:%d", "zhangsan", 18) // printf
// }

// func TestFatal(t *testing.T) {
// 	defer gzap.Syncw()
// 	defer gzap.Sync()
// 	gzap.Fatal("hello world", zap.String("name", "zhangsan"), zap.Int("age", 18)) //高性能
// 	//不会执行
// 	gzap.Fatalw("hello world", "name", "zhangsan", "age", 18) //性能普通
// 	//不会执行
// 	gzap.Fatalf("hello world; name:%s; age:%d", "zhangsan", 18) // printf
// }

func TestLog(t *testing.T) {
	gzap.Zap().Log(zap.DebugLevel, "hello world", zap.String("name", "zhangsan"), zap.Int("age", 18)) //不输出
	gzap.Zap().Log(zap.InfoLevel, "hello world", zap.String("name", "zhangsan"), zap.Int("age", 18))
	gzap.Zap().Log(zap.WarnLevel, "hello world", zap.String("name", "zhangsan"), zap.Int("age", 18))
	gzap.Zap().Log(zap.ErrorLevel, "hello world", zap.String("name", "zhangsan"), zap.Int("age", 18))
	gzap.Zap().Log(zap.DPanicLevel, "hello world", zap.String("name", "zhangsan"), zap.Int("age", 18))
}

func TestLn(t *testing.T) {
	gzap.Sap().Debugln("hello world", "name", "zhangsan", "age", 18) // 不输出
	gzap.Sap().Infoln("hello world", "name", "zhangsan", "age", 18)
}

func TestFile(t *testing.T) {
	gzap.SetZapCfg(
		gzap.ZapOutFile(
			"./log/test.log",
			gzap.ZapOutFileMaxSize(128),
			gzap.ZapOutFileMaxAge(7),
			gzap.ZapOutFileMaxBackups(30),
		),
	)

	gzap.Info("hello world", zap.String("name", "zhangsan"), zap.Int("age", 18))
}

func TestConf(t *testing.T) {
	conf := zap.Config{
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
		InitialFields: map[string]interface{}{
			"app": "test",
		},
	}

	gzap.SetZapCfg(gzap.ZapConf(conf))

	gzap.Info("hello world", zap.String("name", "zhangsan"), zap.Int("age", 18))
}

func TestPreset(t *testing.T) {
	gzap.SetZapCfg(gzap.SetPresetFields(map[string]any{"service": "myservice"}))
	gzap.Info("hello world")
}

func TestNamespace(t *testing.T) {
	gzap.Info("hello world", zap.Namespace("user1"), zap.String("name", "zhangsan"), zap.Int("age", 18))
}
