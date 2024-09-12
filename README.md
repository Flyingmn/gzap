## 简介

包装一下zap，方便使用

## 安装

```bash
# 使用go modules
go get -u github.com/Flyingmn/gzap
```

## 使用
```go

// 默认info级别，如果要自定义配置，则执行（注意要在Zap()以及Sap()之前）
// gzap.SetZapCfg(gzap.ZapLevel("debug"), gzap.ZapOutFile("./log/test.log", gzap.ZapOutFileMaxSize(128), gzap.ZapOutFileMaxAge(7)))

//普通zap *zap.Logger
gzap.Zap().Info("hello world", zap.String("name", "zhang"), zap.Any("age", 18))

//带语法糖的zap *zap.SugaredLogger
gzap.Sap().Infow("hello world", "name", "zhang", "age", 18)

// {"level":"info","time":"2024-05-21 16:29:20.927","line":"...testing/testing.go:1595","func":"testing.tRunner","msg":"hello world","name":"zhang","age":18}
```