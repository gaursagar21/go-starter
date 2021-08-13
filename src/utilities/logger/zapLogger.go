package logger

import (
	"context"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"gopkg.in/natefinch/lumberjack.v2"
	"os"
)

type zapLogger struct {
	sugaredLogger *zap.SugaredLogger
}

func getEncoder(isJSON bool) zapcore.Encoder {
	encoderConfig := zap.NewProductionEncoderConfig()
	encoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder
	encoderConfig.TimeKey = "time"
	if isJSON {
		return zapcore.NewJSONEncoder(encoderConfig)
	}
	return zapcore.NewConsoleEncoder(encoderConfig)
}

func getZapLevel(level LogLevel) zapcore.Level {
	zapLevelMapper := map[LogLevel]zapcore.Level{
		DEBUG: zapcore.DebugLevel,
		INFO:  zapcore.InfoLevel,
		ERROR: zapcore.ErrorLevel,
		WARN:  zapcore.WarnLevel,
		FATAL: zapcore.FatalLevel,
	}
	if zapLevel, ok := zapLevelMapper[level]; !ok {
		return zapcore.InfoLevel
	} else {
		return zapLevel
	}
}

func newZapLogger(config Configuration) (Logger, error) {
	var cores []zapcore.Core

	if config.Output == CONSOLE {
		level := getZapLevel(config.LogLevel)
		writer := zapcore.Lock(os.Stdout)
		core := zapcore.NewCore(getEncoder(config.EnableJSON), writer, level)
		cores = append(cores, core)
	}

	if config.Output == FILE {
		level := getZapLevel(config.LogLevel)
		writer := zapcore.AddSync(&lumberjack.Logger{
			Filename: config.FileLocation,
			MaxSize:  100,
			Compress: true,
			MaxAge:   28,
		})
		core := zapcore.NewCore(getEncoder(config.EnableJSON), writer, level)
		cores = append(cores, core)
	}

	combinedCore := zapcore.NewTee(cores...)

	// AddCallerSkip skips 2 number of callers, this is important else the file that gets
	// logged will always be the wrapped file. In our case zapLogger.go
	logger := zap.New(combinedCore,
		zap.AddCallerSkip(2),
		zap.AddCaller(),
	).Sugar()

	return &zapLogger{
		sugaredLogger: logger,
	}, nil
}

func (l *zapLogger) Debugf(ctx context.Context, facet Facet, format string, args ...interface{}) {
	l = l.WithFields(GetMetaFromContext(ctx)).WithFields(facet).(*zapLogger)
	l.sugaredLogger.Debugf(format, args...)
}

func (l *zapLogger) Debug(ctx context.Context, facet Facet, args ...interface{}) {
	l = l.WithFields(GetMetaFromContext(ctx)).WithFields(facet).(*zapLogger)
	l.sugaredLogger.Debug(args...)
}

func (l *zapLogger) Infof(ctx context.Context, facet Facet, format string, args ...interface{}) {
	l = l.WithFields(GetMetaFromContext(ctx)).WithFields(facet).(*zapLogger)
	l.sugaredLogger.Infof(format, args...)
}

func (l *zapLogger) Info(ctx context.Context, facet Facet, args ...interface{}) {
	l = l.WithFields(GetMetaFromContext(ctx)).WithFields(facet).(*zapLogger)
	l.sugaredLogger.Info(args...)
}

func (l *zapLogger) Warnf(ctx context.Context, facet Facet, format string, args ...interface{}) {
	l = l.WithFields(GetMetaFromContext(ctx)).WithFields(facet).(*zapLogger)
	l.sugaredLogger.Warnf(format, args...)
}

func (l *zapLogger) Warn(ctx context.Context, facet Facet, args ...interface{}) {
	l = l.WithFields(GetMetaFromContext(ctx)).WithFields(facet).(*zapLogger)
	l.sugaredLogger.Warn(args...)
}

func (l *zapLogger) Errorf(ctx context.Context, facet Facet, format string, args ...interface{}) {
	l = l.WithFields(GetMetaFromContext(ctx)).WithFields(facet).(*zapLogger)
	l.sugaredLogger.Errorf(format, args...)
}

func (l *zapLogger) Error(ctx context.Context, facet Facet, args ...interface{}) {
	l = l.WithFields(GetMetaFromContext(ctx)).WithFields(facet).(*zapLogger)
	l.sugaredLogger.Error(args...)
}

func (l *zapLogger) Fatalf(ctx context.Context, facet Facet, format string, args ...interface{}) {
	l = l.WithFields(GetMetaFromContext(ctx)).WithFields(facet).(*zapLogger)
	l.sugaredLogger.Fatalf(format, args...)
}

func (l *zapLogger) Fatal(ctx context.Context, facet Facet, args ...interface{}) {
	l = l.WithFields(GetMetaFromContext(ctx)).WithFields(facet).(*zapLogger)
	l.sugaredLogger.Fatal(args...)
}

func (l *zapLogger) WithFields(facet Facet) Logger {
	var f = make([]interface{}, 0)
	for k, v := range facet.GetFields() {
		f = append(f, k)
		f = append(f, v)
	}
	newLogger := l.sugaredLogger.With(f...)
	return &zapLogger{newLogger}
}
