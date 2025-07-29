package exporter

import (
	"encoding/json"
	"fmt"
	"os"
)

type JSONExporter interface {
	Export(filePath string, data interface{}) error
}

type jsonExporter struct{}

func NewJSONExporter() JSONExporter {
	return &jsonExporter{}
}

func (e *jsonExporter) Export(filePath string, data interface{}) error {
	jsonBytes, err := json.MarshalIndent(data, "", "  ")
	if err != nil {
		return fmt.Errorf("failed to marshal JSON: %w", err)
	}

	if err := os.WriteFile(filePath, jsonBytes, 0o600); err != nil {
		return fmt.Errorf("failed to write JSON file: %w", err)
	}

	return nil
}
