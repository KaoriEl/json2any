package processor

import (
	"fmt"
	"strconv"
	"sync"
	"time"

	"github.com/KaoriEl/json2xlsx/internal/models"
	"github.com/iancoleman/orderedmap"
)

type DataProcessor interface {
	ProcessOrderedMaps(records []*orderedmap.OrderedMap, maxWorkers int, onProgress func(current int)) ([]models.ProcessedRow, error)
	ProcessStringMaps(records []map[string]string, maxWorkers int, onProgress func(current int)) ([]map[string]interface{}, error)
}

type processor struct{}

func NewProcessor() DataProcessor {
	return &processor{}
}

var dateFormats = []string{
	time.RFC3339,
	"2006-01-02",
	"02.01.2006",
	"2006-01-02 15:04:05",
	"02.01.2006 15:04:05",
}

func tryParseDate(input string) (time.Time, error) {
	for _, layout := range dateFormats {
		if t, err := time.Parse(layout, input); err == nil {
			return t, nil
		}
	}
	return time.Time{}, fmt.Errorf("failed to parse date from string: %q", input)
}

func (p *processor) ProcessOrderedMaps(
	records []*orderedmap.OrderedMap,
	maxWorkers int,
	onProgress func(current int),
) ([]models.ProcessedRow, error) {
	if maxWorkers <= 0 {
		maxWorkers = 20
	}

	result := make([]models.ProcessedRow, len(records))
	var wg sync.WaitGroup
	var mu sync.Mutex
	var firstErr error
	sem := make(chan struct{}, maxWorkers)

	processRecord := func(record *orderedmap.OrderedMap) models.ProcessedRow {
		row := make(models.ProcessedRow)
		for _, key := range record.Keys() {
			rawValue, _ := record.Get(key)
			row[key] = normalizeValue(rawValue)
		}
		return row
	}

	wg.Add(len(records))
	for i, record := range records {
		sem <- struct{}{}
		go func(i int, record *orderedmap.OrderedMap) {
			defer wg.Done()
			defer func() { <-sem }()

			row := processRecord(record)

			mu.Lock()
			result[i] = row
			mu.Unlock()

			if onProgress != nil {
				onProgress(i + 1)
			}
		}(i, record)
	}

	wg.Wait()
	return result, firstErr
}

func (p *processor) ProcessStringMaps(records []map[string]string, maxWorkers int, onProgress func(current int)) ([]map[string]interface{}, error) {
	if maxWorkers <= 0 {
		maxWorkers = 20
	}

	result := make([]map[string]interface{}, len(records))
	var wg sync.WaitGroup
	var mu sync.Mutex
	sem := make(chan struct{}, maxWorkers)

	processRecord := func(record map[string]string) map[string]interface{} {
		row := make(map[string]interface{})
		for key, rawValue := range record {
			row[key] = normalizeStringValue(rawValue)
		}
		return row
	}

	wg.Add(len(records))
	for i, record := range records {
		sem <- struct{}{}
		go func(i int, record map[string]string) {
			defer wg.Done()
			defer func() { <-sem }()

			row := processRecord(record)

			mu.Lock()
			result[i] = row
			mu.Unlock()

			if onProgress != nil {
				onProgress(i + 1)
			}
		}(i, record)
	}

	wg.Wait()
	return result, nil
}

func normalizeValue(rawValue interface{}) interface{} {
	switch v := rawValue.(type) {
	case float64:
		if v == float64(int64(v)) {
			return int64(v)
		}
		return v
	case string:
		if parsedTime, err := tryParseDate(v); err == nil {
			return parsedTime
		}
		if num, err := strconv.ParseFloat(v, 64); err == nil {
			if num == float64(int64(num)) {
				return int64(num)
			}
			return num
		}
		return v
	case bool:
		return v
	case nil:
		return ""
	default:
		return fmt.Sprintf("%v", v)
	}
}

func normalizeStringValue(rawValue string) interface{} {
	if rawValue == "" {
		return ""
	}

	if parsedTime, err := tryParseDate(rawValue); err == nil {
		return parsedTime
	}
	if num, err := strconv.ParseFloat(rawValue, 64); err == nil {
		if num == float64(int64(num)) {
			return int64(num)
		}
		return num
	}

	return rawValue
}
