package loader

import (
	"bufio"
	"encoding/csv"
	"encoding/json"
	"fmt"
	"os"
	"strings"

	"github.com/iancoleman/orderedmap"
	"github.com/xuri/excelize/v2"
)

type DataLoader interface {
	LoadJson(filePath string) ([]*orderedmap.OrderedMap, error)
	LoadXLSX(filePath string) ([]map[string]string, error)
	LoadCSV(filePath string) ([]map[string]string, error)
	LoadTXT(filePath string) ([]map[string]string, error)
}

type dataLoader struct{}

func NewLoader() DataLoader {
	return &dataLoader{}
}

func (d *dataLoader) LoadJson(filePath string) ([]*orderedmap.OrderedMap, error) {
	data, err := os.ReadFile(filePath)
	if err != nil {
		return nil, fmt.Errorf("failed to read file: %w", err)
	}

	var records []*orderedmap.OrderedMap
	if err := json.Unmarshal(data, &records); err != nil {
		return nil, fmt.Errorf("failed to unmarshal JSON: %w", err)
	}

	return records, nil
}

func (d *dataLoader) LoadXLSX(filePath string) ([]map[string]string, error) {
	if _, err := os.Stat(filePath); os.IsNotExist(err) {
		return nil, fmt.Errorf("file not found: %s", filePath)
	}

	f, err := excelize.OpenFile(filePath)
	if err != nil {
		return nil, fmt.Errorf("failed to open XLSX: %w", err)
	}

	sheetName := f.GetSheetName(0)
	if sheetName == "" {
		return nil, fmt.Errorf("no sheets found in file")
	}

	rows, err := f.GetRows(sheetName)
	if err != nil {
		return nil, fmt.Errorf("failed to read rows: %w", err)
	}

	if len(rows) < 1 {
		return nil, fmt.Errorf("sheet is empty")
	}

	headers := rows[0]
	records := make([]map[string]string, 0, len(rows)-1)

	for _, row := range rows[1:] {
		entry := make(map[string]string)
		for i, cell := range row {
			if i < len(headers) {
				entry[headers[i]] = cell
			}
		}
		records = append(records, entry)
	}

	return records, nil
}

func (d *dataLoader) LoadCSV(filePath string) ([]map[string]string, error) {
	file, err := os.Open(filePath)
	if err != nil {
		return nil, fmt.Errorf("failed to open CSV file: %w", err)
	}
	defer file.Close()

	reader := csv.NewReader(file)
	rows, err := reader.ReadAll()
	if err != nil {
		return nil, fmt.Errorf("failed to read CSV file: %w", err)
	}

	if len(rows) < 1 {
		return nil, fmt.Errorf("CSV file is empty")
	}

	headers := rows[0]

	// Предварительное выделение памяти для среза records
	records := make([]map[string]string, 0, len(rows)-1) // Заголовок не учитывается

	for _, row := range rows[1:] {
		entry := make(map[string]string)
		for i, cell := range row {
			if i < len(headers) {
				entry[headers[i]] = cell
			}
		}
		records = append(records, entry)
	}

	return records, nil
}

func (d *dataLoader) LoadTXT(filePath string) ([]map[string]string, error) {
	file, err := os.Open(filePath)
	if err != nil {
		return nil, fmt.Errorf("failed to open TXT file: %w", err)
	}
	defer file.Close()

	var records []map[string]string
	scanner := bufio.NewScanner(file)
	headers := []string{}
	for scanner.Scan() {
		line := scanner.Text()
		if len(headers) == 0 {
			headers = strings.Split(line, "\t")
			continue
		}
		fields := strings.Split(line, "\t")
		entry := make(map[string]string)
		for i, field := range fields {
			if i < len(headers) {
				entry[headers[i]] = field
			}
		}
		records = append(records, entry)
	}

	if err := scanner.Err(); err != nil {
		return nil, fmt.Errorf("failed to read TXT file: %w", err)
	}

	return records, nil
}
