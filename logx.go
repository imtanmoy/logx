package logx

// A global variable so that log functions can be directly accessed
var defaultLogger Logger

//Fields Type to pass when we want to call WithFields for structured logging
type Fields map[string]interface{}

//Logger is our contract for the logger
type Logger interface {
	Error(args ...interface{})
	Errorf(format string, args ...interface{})

	Fatal(args ...interface{})
	Fatalf(format string, args ...interface{})

	Info(args ...interface{})
	Infof(format string, args ...interface{})

	Warn(args ...interface{})
	Warnf(format string, args ...interface{})

	Debug(args ...interface{})
	Debugf(format string, args ...interface{})

	Panic(args ...interface{})
	Panicf(format string, args ...interface{})

	Print(args ...interface{})
	Printf(format string, args ...interface{})
	Println(args ...interface{})

	WithFields(keyValues Fields) Logger
}

func initialize() error {
	logger, err := NewZapLogger()
	if err != nil {
		return err
	}
	defaultLogger = logger
	return nil
}

func init() {
	_ = initialize()
}

func Debugf(format string, args ...interface{}) {
	defaultLogger.Debugf(format, args...)
}

func Debug(args ...interface{}) {
	defaultLogger.Debug(args...)
}

func Infof(format string, args ...interface{}) {
	defaultLogger.Infof(format, args...)
}

func Info(args ...interface{}) {
	defaultLogger.Info(args...)
}

func Warnf(format string, args ...interface{}) {
	defaultLogger.Warnf(format, args...)
}

func Warn(args ...interface{}) {
	defaultLogger.Warn(args...)
}

func Errorf(format string, args ...interface{}) {
	defaultLogger.Errorf(format, args...)
}

func Error(args ...interface{}) {
	defaultLogger.Error(args...)
}

func Fatalf(format string, args ...interface{}) {
	defaultLogger.Fatalf(format, args...)
}

func Fatal(args ...interface{}) {
	defaultLogger.Fatal(args...)
}

func Panicf(format string, args ...interface{}) {
	defaultLogger.Panicf(format, args...)
}

func Panic(args ...interface{}) {
	defaultLogger.Panic(args...)
}

func WithFields(keyValues Fields) Logger {
	return defaultLogger.WithFields(keyValues)
}
