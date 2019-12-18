package logx

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"time"
)

func SyslogTimeEncoder(t time.Time, enc zapcore.PrimitiveArrayEncoder) {
	enc.AppendString(t.Format("Jan  2 15:04:05"))
}

func CustomLevelEncoder(level zapcore.Level, enc zapcore.PrimitiveArrayEncoder) {
	enc.AppendString("[" + level.CapitalString() + "]")
}

func NewZapLogger() ILogger {
	cfg := zap.Config{
		Encoding:    "console",
		OutputPaths: []string{"stderr"},
		EncoderConfig: zapcore.EncoderConfig{
			MessageKey:  "message",
			TimeKey:     "time",
			EncodeTime:  zapcore.ISO8601TimeEncoder,
			LevelKey:    "level",
			EncodeLevel: zapcore.CapitalLevelEncoder,
		},
	}
	cfg.Level = zap.NewAtomicLevelAt(zapcore.DebugLevel)
	cfg.EncoderConfig.EncodeTime = SyslogTimeEncoder
	cfg.EncoderConfig.EncodeLevel = CustomLevelEncoder

	zLogger, _ := cfg.Build()
	defer zLogger.Sync()
	log := *zLogger.Sugar()

	return &zapLogger{
		sugaredLogger: &log,
	}
}

var _ ILogger = (*zapLogger)(nil)

type zapLogger struct {
	sugaredLogger *zap.SugaredLogger
}

// Print logs a message at InfoLevel
func (l *zapLogger) Print(args ...interface{}) {
	l.sugaredLogger.Info(args...)
}

// Printf logs a message at InfoLevel with string format
func (l *zapLogger) Printf(format string, args ...interface{}) {
	l.sugaredLogger.Infof(format, args...)
}

// Println logs a message at InfoLevel
func (l *zapLogger) Println(args ...interface{}) {
	l.sugaredLogger.Info(args...)
}

// Debugf logs a message at DebugLevel with string format
func (l *zapLogger) Debugf(format string, args ...interface{}) {
	l.sugaredLogger.Debugf(format, args...)
}

// Debug logs a message at DebugLevel
func (l *zapLogger) Debug(args ...interface{}) {
	l.sugaredLogger.Debug(args...)
}

// Infof logs a message at InfoLevel with string format
func (l *zapLogger) Infof(format string, args ...interface{}) {
	l.sugaredLogger.Infof(format, args...)
}

// Info logs a message at InfoLevel
func (l *zapLogger) Info(args ...interface{}) {
	l.sugaredLogger.Info(args...)
}

// Warnf logs a message at WarnLevel with string format
func (l *zapLogger) Warnf(format string, args ...interface{}) {
	l.sugaredLogger.Warnf(format, args...)
}

// Warn logs a message at WarnLevel
func (l *zapLogger) Warn(args ...interface{}) {
	l.sugaredLogger.Warn(args...)
}

// Errorf logs a message at ErrorLevel with string format
func (l *zapLogger) Errorf(format string, args ...interface{}) {
	l.sugaredLogger.Errorf(format, args...)
}

// Error logs a message at ErrorLevel
func (l *zapLogger) Error(args ...interface{}) {
	l.sugaredLogger.Error(args...)
}

// Fatalf logs a message at FatalLevel with string format
func (l *zapLogger) Fatalf(format string, args ...interface{}) {
	l.sugaredLogger.Fatalf(format, args...)
}

// Fatal logs a message at FatalLevel
func (l *zapLogger) Fatal(args ...interface{}) {
	l.sugaredLogger.Fatal(args...)
}

// Panicf logs a message at PanicLevel with string format
func (l *zapLogger) Panicf(format string, args ...interface{}) {
	l.sugaredLogger.Fatalf(format, args...)
}

// Panic logs a message at PanicLevel
func (l *zapLogger) Panic(args ...interface{}) {
	l.sugaredLogger.Fatal(args...)
}

// WithFields returns a new logger instance which can Log a message at the given level with context key/value pairs
func (l *zapLogger) WithFields(fields Fields) ILogger {
	var f = make([]interface{}, 0)
	for k, v := range fields {
		f = append(f, k)
		f = append(f, v)
	}
	newLogger := l.sugaredLogger.With(f...)
	return &zapLogger{newLogger}
}
