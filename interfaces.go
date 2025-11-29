package sglogger

import "context"

// Level представляет уровень логирования
type Level int

const (
    LevelDebug Level = iota // Уровень отладки - детальная информация для разработчиков
    LevelInfo               // Информационный уровень - общая информация о работе приложения
    LevelWarn               // Уровень предупреждения - нештатные ситуации, которые не приводят к ошибкам
    LevelError              // Уровень ошибки - ошибки, которые влияют на работу приложения
    LevelFatal              // Критический уровень - ошибки, приводящие к остановке приложения
)

// LoggerProvider определяет интерфейс для провайдеров логирования.
// Провайдеры отвечают за запись логов в конкретные места назначения (консоль, файл, Loki и т.д.).
type LoggerProvider interface {
    // Write записывает лог-сообщение с указанным уровнем, текстом и дополнительными полями.
    // Возвращает ошибку в случае проблем при записи.
    Write(ctx context.Context, level Level, message string, fields Fields) error
    
    // ShouldLog проверяет, нужно ли логировать сообщение данного уровня.
    // Используется для фильтрации логов по уровню важности.
    ShouldLog(ctx context.Context, level Level) bool
    
    // Close освобождает ресурсы провайдера. Должен вызываться при завершении работы приложения.
    Close(ctx context.Context) error
}

// Logger определяет основной интерфейс для логирования в приложении.
// Предоставляет методы для логирования с различными комбинациями параметров:
// - интерполяция строк (форматирование)
// - обработка ошибок
// - структурированные поля
type Logger interface {
    // Debug логирует сообщение уровня отладки
    Debug(ctx context.Context, format string, args ...interface{})
    
    // Info логирует информационное сообщение
    Info(ctx context.Context, format string, args ...interface{})
    
    // Warning логирует предупреждение
    Warning(ctx context.Context, format string, args ...interface{})
    
    // Error логирует сообщение об ошибке
    Error(ctx context.Context, format string, args ...interface{})
    
    // Fatal логирует критическую ошибку и завершает приложение
    Fatal(ctx context.Context, format string, args ...interface{})
    
    // DebugErr логирует сообщение уровня отладки с ошибкой
    DebugErr(ctx context.Context, err error, format string, args ...interface{})
    
    // InfoErr логирует информационное сообщение с ошибкой
    InfoErr(ctx context.Context, err error, format string, args ...interface{})
    
    // WarningErr логирует предупреждение с ошибкой
    WarningErr(ctx context.Context, err error, format string, args ...interface{})
    
    // ErrorErr логирует сообщение об ошибке с дополнительной ошибкой
    ErrorErr(ctx context.Context, err error, format string, args ...interface{})
    
    // FatalErr логирует критическую ошибку с дополнительной ошибкой и завершает приложение
    FatalErr(ctx context.Context, err error, format string, args ...interface{})
    
    // DebugWithFields логирует сообщение уровня отладки с дополнительными полями
    DebugWithFields(ctx context.Context, fields Fields, format string, args ...interface{})
    
    // InfoWithFields логирует информационное сообщение с дополнительными полями
    InfoWithFields(ctx context.Context, fields Fields, format string, args ...interface{})
    
    // WarningWithFields логирует предупреждение с дополнительными полями
    WarningWithFields(ctx context.Context, fields Fields, format string, args ...interface{})
    
    // ErrorWithFields логирует сообщение об ошибке с дополнительными полями
    ErrorWithFields(ctx context.Context, fields Fields, format string, args ...interface{})
    
    // FatalWithFields логирует критическую ошибку с дополнительными полями и завершает приложение
    FatalWithFields(ctx context.Context, fields Fields, format string, args ...interface{})
    
    // DebugErrWithFields логирует сообщение уровня отладки с ошибкой и дополнительными полями
    DebugErrWithFields(ctx context.Context, err error, fields Fields, format string, args ...interface{})
    
    // InfoErrWithFields логирует информационное сообщение с ошибкой и дополнительными полями
    InfoErrWithFields(ctx context.Context, err error, fields Fields, format string, args ...interface{})
    
    // WarningErrWithFields логирует предупреждение с ошибкой и дополнительными полями
    WarningErrWithFields(ctx context.Context, err error, fields Fields, format string, args ...interface{})
    
    // ErrorErrWithFields логирует сообщение об ошибке с дополнительной ошибкой и полями
    ErrorErrWithFields(ctx context.Context, err error, fields Fields, format string, args ...interface{})
    
    // FatalErrWithFields логирует критическую ошибку с дополнительной ошибкой, полями и завершает приложение
    FatalErrWithFields(ctx context.Context, err error, fields Fields, format string, args ...interface{})
}