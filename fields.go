package sglogger

import (
	"context"
	"maps"
)

// Fields представляет дополнительные поля для структурированного логирования.
// Позволяет добавлять метаданные к логам для удобства поиска и анализа.
// Пример: Fields{"user_id": 123, "request_id": "abc-123"}
type Fields map[string]interface{}

// FieldsHandler определяет интерфейс для работы с дополнительными полями логов.
// Обеспечивает извлечение полей из контекста и объединение наборов полей.
type FieldsHandler interface {
	// ExtractFieldsFromContext извлекает поля из контекста и объединяет их с переданными полями.
	// Контекст может содержать стандартные поля (например, trace_id) для сквозной трассировки.
	ExtractFieldsFromContext(ctx context.Context, fields Fields) Fields
	
	// MergeFields объединяет два набора полей. При конфликте ключей
	// значения из fields2 перезаписывают значения из fields1.
	MergeFields(fields1, fields2 Fields) Fields
}

// fieldsHandler реализует интерфейс FieldsHandler для обработки дополнительных полей логов.
type fieldsHandler struct{}

// NewFieldsHandler создает новый экземпляр обработчика полей логов.
// Возвращает интерфейс FieldsHandler для использования в системе логирования.
func NewFieldsHandler() FieldsHandler {
	return &fieldsHandler{}
}

// ExtractFieldsFromContext извлекает поля из контекста и объединяет их с переданными полями.
// В текущей реализации извлекает только trace_id из контекста.
// Если контекст равен nil, возвращает исходные поля без изменений.
func (h *fieldsHandler) ExtractFieldsFromContext(ctx context.Context, fields Fields) Fields {
	if ctx == nil {
		return fields
	}

	result := make(Fields)
	maps.Copy(result, fields)

	// Извлекаем trace_id из контекста, если он присутствует
	if traceID, ok := ctx.Value(TraceIDKey).(string); ok {
		result["trace_id"] = traceID
	}

	return result
}

// MergeFields объединяет два набора полей. При совпадении ключей
// значения из fields2 имеют приоритет над значениями из fields1.
// Возвращает новый набор полей, содержащий объединенные данные.
func (h *fieldsHandler) MergeFields(fields1, fields2 Fields) Fields {
	result := make(Fields)
	
	maps.Copy(result, fields1)
	maps.Copy(result, fields2)
	
	return result
}