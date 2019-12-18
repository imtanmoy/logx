package logx

// A global variable so that log functions can be directly accessed
var defaultLogger ILogger

//Fields Type to pass when we want to call WithFields for structured logging
type Fields map[string]interface{}

//ILogger is our contract for the logger
type ILogger interface {
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

	WithFields(keyValues Fields) ILogger
}

func initialize() {
	logger := NewZapLogger()
	defaultLogger = logger
}

func init() {
	initialize()
}

// Debug logs a message at DebugLevel
func Debug(args ...interface{}) {
	defaultLogger.Debug(args...)
}

// Debugf logs a message at DebugLevel with string format
func Debugf(format string, args ...interface{}) {
	defaultLogger.Debugf(format, args...)
}

// Info logs a message at InfoLevel
func Info(args ...interface{}) {
	defaultLogger.Info(args...)
}

// Infof logs a message at InfoLevel with string format
func Infof(format string, args ...interface{}) {
	defaultLogger.Infof(format, args...)
}

// Warn logs a message at WarnLevel
func Warn(args ...interface{}) {
	defaultLogger.Warn(args...)
}

// Warnf logs a message at WarnLevel with string format
func Warnf(format string, args ...interface{}) {
	defaultLogger.Warnf(format, args...)
}

// Error logs a message at ErrorLevel
func Error(args ...interface{}) {
	defaultLogger.Error(args...)
}

// Errorf logs a message at ErrorLevel with string format
func Errorf(format string, args ...interface{}) {
	defaultLogger.Errorf(format, args...)
}

// Fatal logs a message at FatalLevel
func Fatal(args ...interface{}) {
	defaultLogger.Fatal(args...)
}

// Fatalf logs a message at FatalLevel with string format
func Fatalf(format string, args ...interface{}) {
	defaultLogger.Fatalf(format, args...)
}

// Panic logs a message at PanicLevel
func Panic(args ...interface{}) {
	defaultLogger.Panic(args...)
}

// Panicf logs a message at PanicLevel with string format
func Panicf(format string, args ...interface{}) {
	defaultLogger.Panicf(format, args...)
}

// Print logs a message at InfoLevel
func Print(args ...interface{}) {
	defaultLogger.Info(args...)
}

// Printf logs a message at InfoLevel with string format
func Printf(format string, args ...interface{}) {
	defaultLogger.Infof(format, args...)
}

// Println logs a message at InfoLevel
func Println(args ...interface{}) {
	defaultLogger.Info(args...)
}

// WithFields returns a new logger instance which can Log a message at the given level with context key/value pairs
func WithFields(keyValues Fields) ILogger {
	return defaultLogger.WithFields(keyValues)
}
