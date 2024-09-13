## 简介

包装一下zap，方便使用

## 安装

```bash
# 使用go modules
go get -u github.com/Flyingmn/gzap
```

## 使用

### 默认info级别，如果要自定义级别(注意SetZapCfg要在使用日志之前设置）
```go
gzap.SetZapCfg(gzap.ZapLevel("info"))
```

### 普通zap *zap.Logger; 高性能: Debug, Info, Warn, Error, DPanic, Panic, Fatal
```go
gzap.Info("hello world", zap.String("name", "zhangsan"), zap.Int("age", 18))
// {"level":"info","time":"2024-00-00 00:00:00.000","line":"/gzap/zap_test.go:26","func":"github.com/Flyingmn/gzap_test.Test","msg":"hello world","name":"zhangsan","age":18}
```

### 带语法糖的zap *zap.SugaredLogger; 性能不敏感场景使用: Debugw, Infow, Warnw, Errorw, DPanicw, Panicw, Fatalw
```go
gzap.Infow("hello world", "name", "zhangsan", "age", 18)
// {"level":"info","time":"2024-00-00 00:00:00.000","line":"/gzap/zap_test.go:26","func":"github.com/Flyingmn/gzap_test.Test","msg":"hello world","name":"zhangsan","age":18}
```


### Printf的方式 *zap.SugaredLogger: Debugf, Infof, Warnf, Errorf, DPanicf, Panicf, Fatalf
```go
gzap.Infof("hello world; name:%s; age:%d", "zhangsan", 18)
// {"level":"info","time":"2024-00-00 00:00:00.000","line":"/gzap/zap_test.go:26","func":"github.com/Flyingmn/gzap_test.Test","msg":"hello world; name:zhangsan; age:18"}
```

### 预设字段(注意SetZapCfg要在使用日志之前设置）
```go
gzap.SetZapCfg(gzap.SetPresetFields(map[string]any{"service": "myservice"}))
gzap.Info("hello world")
// {"level":"info","time":"2024-00-00 00:00:00.000","line":"/gzap/zap_test.go:26","func":"github.com/Flyingmn/gzap_test.Test","msg":"hello world","service":"myservice"}
```

### 多层次嵌套
```go
gzap.Info("hello world", zap.Namespace("user1"), zap.String("name", "zhangsan"), zap.Int("age", 18), zap.Namespace("user2"), zap.String("name", "lisi"), zap.Int("age", 19))
// {"level":"info","time":"2024-00-00 00:00:00.000","line":"/gzap/zap_test.go:26","func":"github.com/Flyingmn/gzap_test.Test","msg":"hello world","user1":{"name":"zhangsan","age":18}}
```

### 其他方法, 请获取zapClient调用; 获取logger:gzap.Zap(); 获取sugaredLogger: gzap.Sap()
```go
gzap.Zap().Log(zap.InfoLevel, "hello world", zap.String("name", "zhangsan"), zap.Int("age", 18))
// {"level":"info","time":"2024-00-00 00:00:00.000","line":"/gzap/zap_test.go:26","func":"github.com/Flyingmn/gzap_test.Test","msg":"hello world","name":"zhangsan","age":18}
gzap.Sap().Infoln("hello world", "name", "zhangsan", "age", 18)
// {"level":"info","time":"2024-00-00 00:00:00.000","line":"/gzap/zap_test.go:26","func":"github.com/Flyingmn/gzap_test.Test","msg":"hello world name zhangsan age 18"}
```

### 设置日志输出方式 （注意SetZapCfg要在使用日志之前设置）
```go
gzap.SetZapCfg(
    gzap.ZapOutFile(
        "./log/test.log",               //文件位置
        gzap.ZapOutFileMaxSize(128),    // 日志文件的最大大小(MB为单位)
        gzap.ZapOutFileMaxAge(7),       //保留旧文件的最大天数量
        gzap.ZapOutFileMaxBackups(30),  //保留旧文件的最大个数
    ),
)

gzap.Info("hello world", zap.String("name", "zhangsan"), zap.Int("age", 18))
```

### 配置深度定制, 自定义配置后传入  （注意SetZapCfg要在使用日志之前设置）
```go
config := zap.Config{
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
// {"level":"info","time":"2024-00-00 00:00:00.000","line":"/gzap/zap_test.go:26","func":"github.com/Flyingmn/gzap_test.Test","msg":"hello world","app":"test","name":"zhangsan","age":18}
```