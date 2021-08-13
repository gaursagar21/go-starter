package logger

import (
	"context"
	"errors"
	"github.com/gaursagarMT/starter/src/helper"
)

var log Logger

type LogLevel int

const (
	DEBUG LogLevel = iota
	INFO
	WARN
	ERROR
	FATAL
)

const (
	ZapLogger int = iota
	LogrusLogger
	// extendable to any other custom logger which implements this interface
)

var (
	errInvalidLoggerInstance = errors.New("Invalid logger instance")
)

type Logger interface {
	Debugf(ctx context.Context, facet Facet, format string, args ...interface{})
	Debug(ctx context.Context, facet Facet, args ...interface{})
	Infof(ctx context.Context, facet Facet, format string, args ...interface{})
	Info(ctx context.Context, facet Facet, args ...interface{})
	Warnf(ctx context.Context, facet Facet, format string, args ...interface{})
	Warn(ctx context.Context, facet Facet, args ...interface{})
	Errorf(ctx context.Context, facet Facet, format string, args ...interface{})
	Error(ctx context.Context, facet Facet, args ...interface{})
	Fatalf(ctx context.Context, facet Facet, format string, args ...interface{})
	Fatal(ctx context.Context, facet Facet, args ...interface{})
	WithFields(keyValues Facet) Logger
}

type LoggerOutput int

const (
	FILE LoggerOutput = iota
	CONSOLE
)

type Configuration struct {
	LogLevel     LogLevel
	EnableJSON   bool
	Output       LoggerOutput
	FileLocation string
}

func NewLogger(config Configuration, loggerInstance int) error {
	switch loggerInstance {
	case ZapLogger:
		logger, err := newZapLogger(config)
		if err != nil {
			return err
		}
		log = logger
		return nil
	default:
		return errInvalidLoggerInstance
	}
}

func Debugf(ctx context.Context, facet Facet, format string, args ...interface{}) {
	log.Debugf(ctx, facet, format, args...)
}
func Debug(ctx context.Context, facet Facet, args ...interface{}) {
	log.Debug(ctx, facet, args...)
}
func Infof(ctx context.Context, facet Facet, format string, args ...interface{}) {
	log.Infof(ctx, facet, format, args...)
}
func Info(ctx context.Context, facet Facet, args ...interface{}) {
	log.Info(ctx, facet, args...)
}
func Warnf(ctx context.Context, facet Facet, format string, args ...interface{}) {
	log.Warnf(ctx, facet, format, args...)
}
func Warn(ctx context.Context, facet Facet, args ...interface{}) {
	log.Warn(ctx, facet, args...)
}
func Errorf(ctx context.Context, facet Facet, format string, args ...interface{}) {
	log.Errorf(ctx, facet, format, args...)
}
func Error(ctx context.Context, facet Facet, args ...interface{}) {
	log.Error(ctx, facet, args...)
}
func Fatalf(ctx context.Context, facet Facet, format string, args ...interface{}) {
	log.Fatalf(ctx, facet, format, args...)
}
func Fatal(ctx context.Context, facet Facet, args ...interface{}) {
	log.Fatal(ctx, facet, args...)
}
func WithFields(keyValues Facet) Logger {
	return log.WithFields(keyValues)
}

var MetaFields = []string{}

func GetMetaFromContext(ctx context.Context) Facet {
	facet := NewFacets()
	for _, field := range MetaFields {
		value, err := helper.ExtractFromContext(ctx, field)
		if err == nil {
			facet.AddField(field, value)
		}
	}
	return facet
}
