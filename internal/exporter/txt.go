package exporter

import (
	"bufio"
	"fmt"
	"os"
	"time"

	"github.com/KaoriEl/json2any/v2/internal/models"
)

type TXTExporter interface {
	Export(filePath string, data []models.ProcessedRow, onProgress func(current, total int)) error
}

type txtExporter struct{}

func NewTXTExporter() TXTExporter {
	return &txtExporter{}
}

func (e *txtExporter) Export(filePath string, data []models.ProcessedRow, onProgress func(current, total int)) error {
	file, err := os.Create(filePath)
	if err != nil {
		return fmt.Errorf("failed to create TXT file: %w", err)
	}
	defer file.Close()

	writer := bufio.NewWriter(file)
	defer writer.Flush()

	total := len(data)
	for i, row := range data {
		line := stringifyRow(row)
		_, err := writer.WriteString(line + "\n")
		if err != nil {
			return fmt.Errorf("failed to write TXT record: %w", err)
		}

		if onProgress != nil {
			onProgress(i+1, total)
		}
	}

	return nil
}

func stringifyRow(row models.ProcessedRow) string {
	parts := make([]string, 0, len(row))
	for k, v := range row {
		parts = append(parts, fmt.Sprintf("%s=%s", k, stringifyValue(v)))
	}
	return joinWithSemicolon(parts)
}

func stringifyValue(v interface{}) string {
	switch val := v.(type) {
	case nil:
		return ""
	case string:
		return val
	case int, int64:
		return fmt.Sprintf("%d", val)
	case float64:
		return fmt.Sprintf("%f", val)
	case bool:
		if val {
			return "true"
		}
		return "false"
	case time.Time:
		return val.Format(time.RFC3339)
	default:
		return fmt.Sprintf("%v", val)
	}
}

func joinWithSemicolon(parts []string) string {
	result := ""
	for i, part := range parts {
		if i > 0 {
			result += "; "
		}
		result += part
	}
	return result
}
