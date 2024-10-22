# :umbrella: gzap

<div align="center">

Pack up the [Zap](https://github.com/uber-go/zap), out of the box. ：）


[![Go](https://github.com/Flyingmn/gzap/actions/workflows/go.yml/badge.svg)](https://github.com/Flyingmn/gzap/actions/workflows/go.yml) [![codecov](https://codecov.io/github/Flyingmn/gzap/graph/badge.svg?token=I829UGIO29)](https://codecov.io/github/Flyingmn/gzap) [![Go Report Card](https://goreportcard.com/badge/github.com/Flyingmn/gzap)](https://goreportcard.com/report/github.com/Flyingmn/gzap) [![Go Reference](https://pkg.go.dev/badge/github.com/Flyingmn/gzap.svg)](https://pkg.go.dev/github.com/Flyingmn/gzap) ![Static Badge](https://img.shields.io/badge/License-MIT-blue)

</div>

## Installation

```bash
go get -u github.com/Flyingmn/gzap
```

## Quick Start

##### Zap Common (*zap.Logger)
```go
// High performance: Debug, Info, Warn, Error, DPanic, Panic, Fatal

gzap.Info("hello world", zap.String("name", "zhangsan"), zap.Int("age", 18))
/*  Output the logs on the same line. Here, in order to demonstrate the formatting of JSON
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

##### Zap with grammar sugar (*zap.SugaredLogger)
```go
// Usage in performance insensitive scenarios: Debugw, Infow, Warnw, Errorw, DPanicw, Panicw, Fatalw

gzap.Infow("hello world", "name", "zhangsan", "age", 18)
// {"level":"info","msg":"hello world","name":"zhangsan","age":18}
```


##### Zap in Printf format (*zap.SugaredLogger)
```go
// Debugf, Infof, Warnf, Errorf, DPanicf, Panicf, Fatalf

gzap.Infof("hello world; name:%s; age:%d", "zhangsan", 18)
// {"level":"info","msg":"hello world; name:zhangsan; age:18"}
```

##### Multi level nesting
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



## Other settings (if needed)

##### Set log level
```go
// Default info level, if you want to customize the level (note that SetZapCfg needs to be set before using logs)

gzap.SetZapCfg(gzap.ZapLevel("info"))
```

##### Preset fields
```go
// (Note that SetZapCfg needs to be set before using zap）

gzap.SetZapCfg(gzap.SetPresetFields(map[string]any{"service": "myservice"}))
gzap.Info("hello world")
// {"level":"info","msg":"hello world","service":"myservice"}
```

##### Set the log out file
```go
//(Note that SetZapCfg needs to be set before using zap）

gzap.SetZapCfg(
    gzap.ZapOutFile(
        "./log/test.log",               // file location
        gzap.ZapOutFileMaxSize(128),    // the maximum size of the log file (in MB)
        gzap.ZapOutFileMaxAge(7),       // maximum number of days to retain old files
        gzap.ZapOutFileMaxBackups(30),  // maximum number of old files retained
    ),
)

gzap.Info("hello world", zap.String("name", "zhangsan"), zap.Int("age", 18))
```

##### Deep customization of configuration
```go
// 自定义配置后传入  (Note that SetZapCfg needs to be set before using zap）

config := zap.Config{
    Level:       zap.NewAtomicLevelAt(zap.InfoLevel), // Log Level
    Development: true,                                // development mode, stack trace
    Encoding:    "json",                              // Output format console or JSON

    EncoderConfig: zapcore.EncoderConfig{
        TimeKey:        "time",
        LevelKey:       "level",
        NameKey:        "name",
        CallerKey:      "line",
        MessageKey:     "msg",
        FunctionKey:    "func",
        StacktraceKey:  "stacktrace",
        LineEnding:     zapcore.DefaultLineEnding,
        EncodeLevel:    zapcore.LowercaseLevelEncoder,                         //Lowercase encoder
        EncodeTime:     zapcore.TimeEncoderOfLayout("2006-01-02 15:04:05.000"),//Define time format
        EncodeDuration: zapcore.SecondsDurationEncoder,
        EncodeCaller:   zapcore.FullCallerEncoder, // Full path encoder
    }, // Encoder configuration
    InitialFields: map[string]interface{}{
        "app": "test",
    },
}

gzap.SetZapCfg(gzap.ZapConf(conf))

gzap.Info("hello world", zap.String("name", "zhangsan"), zap.Int("age", 18))
// {"level":"info","msg":"hello world","app":"test","name":"zhangsan","age":18}
```

##### Get ZapClient
```go
// Get logger:gzap.Zap(); 

gzap.Zap().Log(
    zap.InfoLevel, 
    "hello world", 
    zap.String("name", "zhangsan"), 
    zap.Int("age", 18)
)
// {"level":"info","msg":"hello world","name":"zhangsan","age":18}

// Get sugaredLogger: gzap.Sap()
gzap.Sap().Infoln("hello world", "name", "zhangsan", "age", 18)
// {"level":"info","msg":"hello world name zhangsan age 18"}
```
