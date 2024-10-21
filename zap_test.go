package gzap_test

import (
	"fmt"
	"testing"

	"github.com/Flyingmn/gzap"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

func TestDebug(t *testing.T) {
	//默认info,所以先设置级别
	gzap.SetZapCfg(gzap.ZapLevel("debug"), gzap.ZapLevel("info"), gzap.ZapLevel("warn"), gzap.ZapLevel("error"), gzap.ZapLevel("dpanic"), gzap.ZapLevel("panic"), gzap.ZapLevel("fatal"))

	defer gzap.Syncw()
	defer gzap.Sync()
	defer gzap.Syncf()
	gzap.Debug("hello world", zap.String("name", "zhangsan"), zap.Int("age", 18)) //高性能

	gzap.Debugw("hello world", "name", "zhangsan", "age", 18) //性能普通

	gzap.Debugf("hello world; name:%s; age:%d", "zhangsan", 18) // printf
}

func TestInfo(t *testing.T) {
	defer gzap.Syncw()
	defer gzap.Sync()
	defer gzap.Syncf()
	gzap.Info("hello world", zap.String("name", "zhangsan"), zap.Int("age", 18)) //高性能

	gzap.Infow("hello world", "name", "zhangsan", "age", 18) //性能普通

	gzap.Infof("hello world; name:%s; age:%d", "zhangsan", 18) // printf
}

func TestWarn(t *testing.T) {
	defer gzap.Syncw()
	defer gzap.Sync()
	defer gzap.Syncf()
	gzap.Warn("hello world", zap.String("name", "zhangsan"), zap.Int("age", 18)) //高性能

	gzap.Warnw("hello world", "name", "zhangsan", "age", 18) //性能普通

	gzap.Warnf("hello world; name:%s; age:%d", "zhangsan", 18) // printf
}

func TestError(t *testing.T) {
	defer gzap.Syncw()
	defer gzap.Sync()
	defer gzap.Syncf()
	gzap.Error("hello world", zap.String("name", "zhangsan"), zap.Int("age", 18)) //高性能

	gzap.Errorw("hello world", "name", "zhangsan", "age", 18) //性能普通

	gzap.Errorf("hello world; name:%s; age:%d", "zhangsan", 18) // printf
}

func TestDPanic(t *testing.T) {
	defer gzap.Syncw()
	defer gzap.Sync()
	defer gzap.Syncf()

	go func() {
		// 捕获异常
		defer func() {
			if err := recover(); err != nil {
				fmt.Println(err)
			}
		}()

		gzap.DPanic("hello world", zap.String("name", "zhangsan"), zap.Int("age", 18)) //高性能
	}()

	go func() {
		// 捕获异常
		defer func() {
			if err := recover(); err != nil {
				fmt.Println(err)
			}
		}()

		gzap.DPanicw("hello world", "name", "zhangsan", "age", 18) //性能普通
	}()

	go func() {
		// 捕获异常
		defer func() {
			if err := recover(); err != nil {
				fmt.Println(err)
			}
		}()

		gzap.DPanicf("hello world; name:%s; age:%d", "zhangsan", 18) // printf
	}()
}

func TestPanic(t *testing.T) {
	defer gzap.Syncw()
	defer gzap.Sync()
	defer gzap.Syncf()

	go func() {
		// 捕获异常
		defer func() {
			if err := recover(); err != nil {
				fmt.Println(err)
			}
		}()

		gzap.Panic("hello world", zap.String("name", "zhangsan"), zap.Int("age", 18)) //高性能
	}()

	go func() {
		// 捕获异常
		defer func() {
			if err := recover(); err != nil {
				fmt.Println(err)
			}
		}()

		gzap.Panicw("hello world", "name", "zhangsan", "age", 18) //性能普通
	}()

	go func() {
		// 捕获异常
		defer func() {
			if err := recover(); err != nil {
				fmt.Println(err)
			}
		}()

		gzap.Panicf("hello world; name:%s; age:%d", "zhangsan", 18) // printf
	}()
}

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
	gzap.Sap().Warnln("hello world", "name", "zhangsan", "age", 18)
	gzap.Sap().Errorln("hello world", "name", "zhangsan", "age", 18)
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

func TestNamespace(t *testing.T) {
	gzap.Info("hello world", zap.Namespace("user1"), zap.String("name", "zhangsan"), zap.Int("age", 18))
}

func TestLoggerNil(t *testing.T) {
	gzap.Info("hello world")
	gzap.SetZapCfg(gzap.ZapLevel("info"))
}

func TestFatal(t *testing.T) {
	// 捕获异常
	defer func() {
		if err := recover(); err != nil {
			fmt.Println("fatal")
		}
	}()
	// go gzap.Fatal("hello world", zap.String("name", "zhangsan"), zap.Int("age", 18)) //高性能

	// go gzap.Fatalw("hello world", "name", "zhangsan", "age", 18) //性能普通

	// go gzap.Fatalf("hello world; name:%s; age:%d", "zhangsan", 18) // printf
}

func TestPreset(t *testing.T) {
	gzap.SetZapCfg(gzap.SetPresetFields(map[string]any{"service": "myservice"}), gzap.ZapDevelopment(true), gzap.ZapCallerSkip(1), gzap.ZapEncodering("console"))
	gzap.Zap().Error("hello world")
}
