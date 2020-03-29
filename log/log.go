package log

import (
	"github.com/json-iterator/go/extra"
	"go.uber.org/zap"
	"time"
)

var (
	sugaredLogger *zap.SugaredLogger
	logger        *zap.Logger
	cfg           zap.Config
	IsDebug       = true
)

func GetDefaultLogger() *WarpLogger {
	return &WarpLogger{sugaredLogger}
}

type WarpLogger struct{ s *zap.SugaredLogger }

func init() {
	extra.RegisterFuzzyDecoders()
	extra.RegisterTimeAsInt64Codec(time.Microsecond)
	extra.SetNamingStrategy(extra.LowerCaseWithUnderscores)
	//extra.SupportPrivateFields()

	InitLogger(true)
}

func InitLogger(isDebug bool) *zap.SugaredLogger {
	cfg = zap.NewProductionConfig()
	IsDebug = isDebug
	if isDebug {
		cfg = zap.NewDevelopmentConfig()
	}
	logger, _ = cfg.Build(zap.AddCallerSkip(1), zap.AddStacktrace(zap.DPanicLevel))
	defer func() { _ = logger.Sync() }()
	sugaredLogger = logger.Named("main").Sugar()
	return sugaredLogger
}

func NewContextLogger(named string) *WarpLogger { return &WarpLogger{logger.Named(named).Sugar()} }

func (w WarpLogger) NewContextLogger(named string) *WarpLogger { return &WarpLogger{w.s.Named(named)} }

func (w WarpLogger) Debug(args ...interface{}) { w.s.Debug(args...) }

func (w WarpLogger) Info(args ...interface{}) { w.s.Info(args...) }

func (w WarpLogger) Warn(args ...interface{}) { w.s.Warn(args...) }

func (w WarpLogger) XWarn(err error) { w.s.Warnf("err: %+v", err) }

func (w WarpLogger) XError(err error) { w.s.Errorf("err: %+v", err) }

func (w WarpLogger) Error(args ...interface{}) { w.s.Error(args...) }

func (w WarpLogger) DPanic(args ...interface{}) { w.s.DPanic(args...) }

func (w WarpLogger) Panic(args ...interface{}) { w.s.Panic(args...) }

func (w WarpLogger) Fatal(args ...interface{}) { w.s.Fatal(args...) }

func (w WarpLogger) Debugf(template string, args ...interface{}) { w.s.Debugf(template, args...) }

func (w WarpLogger) Infof(template string, args ...interface{}) { w.s.Infof(template, args...) }

func (w WarpLogger) Warnf(template string, args ...interface{}) { w.s.Warnf(template, args...) }

func (w WarpLogger) Errorf(template string, args ...interface{}) { w.s.Errorf(template, args...) }

func (w WarpLogger) DPanicf(template string, args ...interface{}) { w.s.DPanicf(template, args...) }

func (w WarpLogger) Panicf(template string, args ...interface{}) { w.s.Panicf(template, args...) }

func (w WarpLogger) Fatalf(template string, args ...interface{}) { w.s.Fatalf(template, args...) }

func Debug(args ...interface{}) { sugaredLogger.Debug(args...) }

func Info(args ...interface{}) { sugaredLogger.Info(args...) }

func Warn(args ...interface{}) { sugaredLogger.Warn(args...) }

func XError(err error) { sugaredLogger.Errorf("err: %+v", err) }

func Error(args ...interface{}) { sugaredLogger.Error(args...) }

func DPanic(args ...interface{}) { sugaredLogger.DPanic(args...) }

func Panic(args ...interface{}) { sugaredLogger.Panic(args...) }

func Fatal(args ...interface{}) { sugaredLogger.Fatal(args...) }

func Debugf(template string, args ...interface{}) { sugaredLogger.Debugf(template, args...) }

func Infof(template string, args ...interface{}) { sugaredLogger.Infof(template, args...) }

func Warnf(template string, args ...interface{}) { sugaredLogger.Warnf(template, args...) }

func Errorf(template string, args ...interface{}) { sugaredLogger.Errorf(template, args...) }

func DPanicf(template string, args ...interface{}) { sugaredLogger.DPanicf(template, args...) }

func Panicf(template string, args ...interface{}) { sugaredLogger.Panicf(template, args...) }

func Fatalf(template string, args ...interface{}) { sugaredLogger.Fatalf(template, args...) }

func CheckFatal(err error) {
	if err != nil {
		sugaredLogger.Fatalf("error: %s", err)
	}
}
func CheckPanic(err error) {
	if err != nil {
		sugaredLogger.Panicf("error: %s", err)
	}
}
