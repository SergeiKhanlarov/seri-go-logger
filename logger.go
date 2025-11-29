package sglogger

import (
	"context"
	"fmt"
	"log"
	"sync"
)

// logger является основной структурой для логирования, управляющей несколькими провайдерами.
// Обеспечивает потокобезопасное логирование через multiple providers.
type logger struct {
	providers     []LoggerProvider
	config        LoggerConfig
	fieldsHandler FieldsHandler
	mu            sync.RWMutex
}

// NewLoggerDefault создает логгер с конфигурацией по умолчанию.
// Использует fmtProvider как единственный провайдер вывода.
// Удобен для быстрого старта и разработки.
func NewLoggerDefault(config ProviderConfig, fieldsHandler FieldsHandler) Logger {
	return &logger{
		providers: []LoggerProvider{
			NewFmtProvider(config),
		},
		config:        config.LoggerConfig,
		fieldsHandler: fieldsHandler,
	}
}

// NewLogger создает кастомный логгер с указанными провайдерами.
// Позволяет гибко настраивать вывод логов через multiple providers.
// Пример: файловый провайдер + провайдер для Sentry + stdout провайдер.
func NewLogger(config LoggerConfig, fieldsHandler FieldsHandler, providers ...LoggerProvider) Logger {
	return &logger{
		providers:     providers,
		config:        config,
		fieldsHandler: fieldsHandler,
	}
}

func (l *logger) Debug(ctx context.Context, format string, args ...interface{}) {
    message := fmt.Sprintf(format, args...)
    l.writeLog(ctx, LevelDebug, message, nil)
}

func (l *logger) Info(ctx context.Context, format string, args ...interface{}) {
    message := fmt.Sprintf(format, args...)
    l.writeLog(ctx, LevelInfo, message, nil)
}

func (l *logger) Warning(ctx context.Context, format string, args ...interface{}) {
    message := fmt.Sprintf(format, args...)
    l.writeLog(ctx, LevelWarn, message, nil)
}

func (l *logger) Error(ctx context.Context, format string, args ...interface{}) {
    message := fmt.Sprintf(format, args...)
    l.writeLog(ctx, LevelError, message, nil)
}

func (l *logger) Fatal(ctx context.Context, format string, args ...interface{}) {
    message := fmt.Sprintf(format, args...)
    l.writeLog(ctx, LevelFatal, message, nil)
    log.Fatal(message)
}

func (l *logger) DebugErr(ctx context.Context, err error, format string, args ...interface{}) {
    message := fmt.Sprintf(format, args...)
    fields := Fields{"error": err.Error()}
    l.writeLog(ctx, LevelDebug, message, fields)
}

func (l *logger) InfoErr(ctx context.Context, err error, format string, args ...interface{}) {
    message := fmt.Sprintf(format, args...)
    fields := Fields{"error": err.Error()}
    l.writeLog(ctx, LevelInfo, message, fields)
}

func (l *logger) WarningErr(ctx context.Context, err error, format string, args ...interface{}) {
    message := fmt.Sprintf(format, args...)
    fields := Fields{"error": err.Error()}
    l.writeLog(ctx, LevelWarn, message, fields)
}

func (l *logger) ErrorErr(ctx context.Context, err error, format string, args ...interface{}) {
    message := fmt.Sprintf(format, args...)
    fields := Fields{"error": err.Error()}
    l.writeLog(ctx, LevelError, message, fields)
}

func (l *logger) FatalErr(ctx context.Context, err error, format string, args ...interface{}) {
    message := fmt.Sprintf(format, args...)
    fields := Fields{"error": err.Error()}
    l.writeLog(ctx, LevelFatal, message, fields)
    log.Fatalf("%s: %v", message, err)
}

func (l *logger) DebugWithFields(ctx context.Context, fields Fields, format string, args ...interface{}) {
    message := fmt.Sprintf(format, args...)
    l.writeLog(ctx, LevelDebug, message, fields)
}

func (l *logger) InfoWithFields(ctx context.Context, fields Fields, format string, args ...interface{}) {
    message := fmt.Sprintf(format, args...)
    l.writeLog(ctx, LevelInfo, message, fields)
}

func (l *logger) WarningWithFields(ctx context.Context, fields Fields, format string, args ...interface{}) {
    message := fmt.Sprintf(format, args...)
    l.writeLog(ctx, LevelWarn, message, fields)
}

func (l *logger) ErrorWithFields(ctx context.Context, fields Fields, format string, args ...interface{}) {
    message := fmt.Sprintf(format, args...)
    l.writeLog(ctx, LevelError, message, fields)
}

func (l *logger) FatalWithFields(ctx context.Context, fields Fields, format string, args ...interface{}) {
    message := fmt.Sprintf(format, args...)
    l.writeLog(ctx, LevelFatal, message, fields)
    log.Fatal(message)
}

func (l *logger) DebugErrWithFields(ctx context.Context, err error, fields Fields, format string, args ...interface{}) {
    message := fmt.Sprintf(format, args...)
    allFields := l.mergeFields(fields, Fields{"error": err.Error()})
    l.writeLog(ctx, LevelDebug, message, allFields)
}

func (l *logger) InfoErrWithFields(ctx context.Context, err error, fields Fields, format string, args ...interface{}) {
    message := fmt.Sprintf(format, args...)
    allFields := l.mergeFields(fields, Fields{"error": err.Error()})
    l.writeLog(ctx, LevelInfo, message, allFields)
}

func (l *logger) WarningErrWithFields(ctx context.Context, err error, fields Fields, format string, args ...interface{}) {
    message := fmt.Sprintf(format, args...)
    allFields := l.mergeFields(fields, Fields{"error": err.Error()})
    l.writeLog(ctx, LevelWarn, message, allFields)
}

func (l *logger) ErrorErrWithFields(ctx context.Context, err error, fields Fields, format string, args ...interface{}) {
    message := fmt.Sprintf(format, args...)
    allFields := l.mergeFields(fields, Fields{"error": err.Error()})
    l.writeLog(ctx, LevelError, message, allFields)
}

func (l *logger) FatalErrWithFields(ctx context.Context, err error, fields Fields, format string, args ...interface{}) {
    message := fmt.Sprintf(format, args...)
    allFields := l.mergeFields(fields, Fields{"error": err.Error()})
    l.writeLog(ctx, LevelFatal, message, allFields)
    log.Fatalf("%s: %v", message, err)
}

func (l *logger) writeLog(ctx context.Context, level Level, message string, fields Fields) {
    l.mu.RLock()
    defer l.mu.RUnlock()

    allFields := l.extractFieldsFromContext(ctx, fields)

    for _, provider := range l.providers {
        if provider.ShouldLog(ctx, level) {
            provider.Write(ctx, level, message, allFields)
        }
    }
}

func (l *logger) extractFieldsFromContext(ctx context.Context, fields Fields) Fields {
    return l.fieldsHandler.ExtractFieldsFromContext(ctx, fields)
}

func (l *logger) mergeFields(fields1, fields2 Fields) Fields {    
    return l.fieldsHandler.MergeFields(fields1, fields2)
}