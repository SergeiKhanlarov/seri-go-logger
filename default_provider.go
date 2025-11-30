package sglogger

import (
	"context"
	"fmt"
	"strings"
	"time"
)

// fmtProvider реализует LoggerProvider для вывода логов в стандартный вывод
// с использованием пакета fmt. Подходит для разработки и отладки.
type fmtProvider struct {
	config ProviderConfig
}

// NewFmtProvider создает новый экземпляр fmtProvider с заданной конфигурацией.
// Возвращает интерфейс LoggerProvider для использования в системе логирования.
func NewFmtProvider(config ProviderConfig) LoggerProvider {
	return &fmtProvider{
		config: config,
	}
}

// Write записывает лог-сообщение в стандартный вывод, если уровень логирования
// соответствует конфигурации провайдера.
func (p *fmtProvider) Write(ctx context.Context, level Level, message string, fields Fields) error {
	if !p.ShouldLog(ctx, level) {
		return nil
	}

	var levelStr string
	switch level {
	case LevelDebug:
		levelStr = "debug"
	case LevelInfo:
		levelStr = "info"
	case LevelWarn:
		levelStr = "warning"
	case LevelError:
		levelStr = "error"
	case LevelFatal:
		levelStr = "critical"
	}

	fmt.Printf("[%s] %s \"%s\" %s\n", 
		time.Now().Format("2006-01-02 15:04:05"),
		levelStr,
		message, 
		serializeFields(fields),
	)

	return nil
}

// ShouldLog определяет, нужно ли логировать сообщение данного уровня.
// Использует минимальный уровень логирования из конфигурации провайдера.
func (p *fmtProvider) ShouldLog(ctx context.Context, level Level) bool {
	return level >= p.config.Level
}

// Close реализует метод закрытия провайдера. 
// В данной реализации не выполняет никаких действий, так как вывод в stdout
// не требует очистки ресурсов.
func (p *fmtProvider) Close(ctx context.Context) error {
	return nil
}

// serializeFields преобразует map полей в строку формата "key1=value1 key2=value2".
// Строковые значения заключаются в кавычки, остальные выводятся как есть.
func serializeFields(fields map[string]interface{}) string {
	if len(fields) == 0 {
		return ""
	}
	
	var pairs []string
	for k, v := range fields {
		switch val := v.(type) {
		case string:
			pairs = append(pairs, fmt.Sprintf("%s=%q", k, val))
		default:
			pairs = append(pairs, fmt.Sprintf("%s=%v", k, val))
		}
	}
	return "{" + strings.Join(pairs, " ") + "}"
}