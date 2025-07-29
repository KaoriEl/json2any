package exporter

import (
	"encoding/csv"
	"fmt"
	"os"
	"strconv"
	"time"

	"github.com/KaoriEl/json2any/v2/internal/models"
)

type CSVExporter interface {
	Export(filePath string, keys []string, data []models.ProcessedRow, onProgress func(current, total int)) error
}

type csvExporter struct{}

func NewCSVExporter() CSVExporter {
	return &csvExporter{}
}

func (e *csvExporter) Export(filePath string, keys []string, data []models.ProcessedRow, onProgress func(current, total int)) error {
	file, err := os.Create(filePath)
	if err != nil {
		return fmt.Errorf("failed to create CSV file: %w", err)
	}
	defer file.Close()

	writer := csv.NewWriter(file)
	defer writer.Flush()

	// Записываем заголовки
	if err := writer.Write(keys); err != nil {
		return fmt.Errorf("failed to write CSV header: %w", err)
	}

	total := len(data)
	for i, row := range data {
		record := make([]string, len(keys))
		for j, key := range keys {
			val := row[key]
			record[j] = stringify(val)
		}
		if err := writer.Write(record); err != nil {
			return fmt.Errorf("failed to write CSV record: %w", err)
		}

		if onProgress != nil {
			onProgress(i+1, total)
		}
	}

	return nil
}

func stringify(v interface{}) string {
	switch val := v.(type) {
	case nil:
		return ""
	case string:
		return val
	case int, int64:
		return fmt.Sprintf("%d", val)
	case float64:
		return strconv.FormatFloat(val, 'f', -1, 64)
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
