package gzap

import "go.uber.org/zap"

// 刷新缓存
func Sync() {
	Zap().Sync()
}
func Syncw() {
	Sap().Sync()
}

func Syncf() {
	Sap().Sync()
}

/*
*- 高性能方法

	// Structured context as strongly typed Field values.
	gzap.Info("failed to fetch URL",
		zap.String("url", url),
		zap.Int("attempt", 3),
		zap.Duration("backoff", time.Second),
	)
*/
func Debug(msg string, fields ...zap.Field) {
	Zap().Debug(msg, fields...)
}

func Info(msg string, fields ...zap.Field) {
	Zap().Info(msg, fields...)
}

func Warn(msg string, fields ...zap.Field) {
	Zap().Warn(msg, fields...)
}

func Error(msg string, fields ...zap.Field) {
	Zap().Error(msg, fields...)
}
func DPanic(msg string, fields ...zap.Field) {
	Zap().DPanic(msg, fields...)
}
func Panic(msg string, fields ...zap.Field) {
	Zap().Panic(msg, fields...)
}
func Fatal(msg string, fields ...zap.Field) {
	Zap().Fatal(msg, fields...)
}

/*
*- 语法糖 - printf 方法

	gzap.Infof("Failed to fetch URL: %s", err)
*/
func Debugf(template string, args ...interface{}) {
	Sap().Debugf(template, args...)
}
func Infof(template string, args ...interface{}) {
	Sap().Infof(template, args...)
}
func Warnf(template string, args ...interface{}) {
	Sap().Warnf(template, args...)
}
func Errorf(template string, args ...interface{}) {
	Sap().Errorf(template, args...)
}
func DPanicf(template string, args ...interface{}) {
	Sap().DPanicf(template, args...)
}
func Panicf(template string, args ...interface{}) {
	Sap().Panicf(template, args...)
}
func Fatalf(template string, args ...interface{}) {
	Sap().Fatalf(template, args...)
}

/*
*语法糖 - 松散键值对 方法

	// Structured context as loosely typed key-value pairs.
	gzap.Infow("failed to fetch URL",
		"url", url,
		"attempt", 3,
		"backoff", time.Second,
	)
*/
func Debugw(msg string, keysAndValues ...interface{}) {
	Sap().Debugw(msg, keysAndValues...)
}
func Infow(msg string, keysAndValues ...interface{}) {
	Sap().Infow(msg, keysAndValues...)
}
func Warnw(msg string, keysAndValues ...interface{}) {
	Sap().Warnw(msg, keysAndValues...)
}
func Errorw(msg string, keysAndValues ...interface{}) {
	Sap().Errorw(msg, keysAndValues...)
}
func DPanicw(msg string, keysAndValues ...interface{}) {
	Sap().DPanicw(msg, keysAndValues...)
}
func Panicw(msg string, keysAndValues ...interface{}) {
	Sap().Panicw(msg, keysAndValues...)
}
func Fatalw(msg string, keysAndValues ...interface{}) {
	Sap().Fatalw(msg, keysAndValues...)
}
