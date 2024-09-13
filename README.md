## 简介

包装一下zap，方便使用

## 安装

```bash
# 使用go modules
go get -u github.com/Flyingmn/gzap
```

## 使用

#### 设置日志级别
```go
// 默认info级别，如果要自定义级别(注意SetZapCfg要在使用日志之前设置）

gzap.SetZapCfg(gzap.ZapLevel("info"))
```

#### 普通zap *zap.Logger
```go
// 高性能: Debug, Info, Warn, Error, DPanic, Panic, Fatal

gzap.Info("hello world", zap.String("name", "zhangsan"), zap.Int("age", 18))
/* 
{
    "level":"info",
    "time":"2024-10-01 12:00:00.000",
    "line":"/gzap/zap_test.go:143",
    "func":"github.com/Flyingmn/gzap_test.TestFunc",
    "msg":"hello world",
    "name":"zhangsan",
    "age":18
}
*/
```

#### 带语法糖的zap *zap.SugaredLogger
```go
// 性能不敏感场景使用: Debugw, Infow, Warnw, Errorw, DPanicw, Panicw, Fatalw

gzap.Infow("hello world", "name", "zhangsan", "age", 18)
// {"level":"info","msg":"hello world","name":"zhangsan","age":18}
```


#### Printf的方式 *zap.SugaredLogger
```go
// Debugf, Infof, Warnf, Errorf, DPanicf, Panicf, Fatalf

gzap.Infof("hello world; name:%s; age:%d", "zhangsan", 18)
// {"level":"info","msg":"hello world; name:zhangsan; age:18"}
```

#### 预设字段
```go
// (注意SetZapCfg要在使用日志之前设置）

gzap.SetZapCfg(gzap.SetPresetFields(map[string]any{"service": "myservice"}))
gzap.Info("hello world")
// {"level":"info","msg":"hello world","service":"myservice"}
```

#### 多层次嵌套
```go
gzap.Info(
    "hello world", 
    zap.Namespace("user1"), 
    zap.String("name", "zhangsan"), 
    zap.Int("age", 18), 
    zap.Namespace("user2"), 
    zap.String("name", "lisi"), 
    zap.Int("age", 19)
)
// {"level":"info","msg":"hello world","user1":{"name":"zhangsan","age":18}}
```

#### 设置日志输出方式 
```go
//（注意SetZapCfg要在使用日志之前设置）

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

#### 配置深度定制
```go
// 自定义配置后传入  （注意SetZapCfg要在使用日志之前设置）

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
        EncodeLevel:    zapcore.LowercaseLevelEncoder,                         //小写编码器
        EncodeTime:     zapcore.TimeEncoderOfLayout("2006-01-02 15:04:05.000"),//定义时间格式
        EncodeDuration: zapcore.SecondsDurationEncoder,
        EncodeCaller:   zapcore.FullCallerEncoder, // 全路径编码器
    }, // 编码器配置
    InitialFields: map[string]interface{}{
        "app": "test",
    },
}

gzap.SetZapCfg(gzap.ZapConf(conf))

gzap.Info("hello world", zap.String("name", "zhangsan"), zap.Int("age", 18))
// {"level":"info","msg":"hello world","app":"test","name":"zhangsan","age":18}
```

#### 获取zapClient
```go
// 获取logger:gzap.Zap(); 

gzap.Zap().Log(
    zap.InfoLevel, 
    "hello world", 
    zap.String("name", "zhangsan"), 
    zap.Int("age", 18)
)
// {"level":"info","msg":"hello world","name":"zhangsan","age":18}

// 获取sugaredLogger: gzap.Sap()
gzap.Sap().Infoln("hello world", "name", "zhangsan", "age", 18)
// {"level":"info","msg":"hello world name zhangsan age 18"}
```