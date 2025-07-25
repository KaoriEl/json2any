package loader

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/iancoleman/orderedmap"
)

type DataLoader interface {
	Load(filePath string) ([]*orderedmap.OrderedMap, error)
}

type jsonLoader struct{}

func NewJSONLoader() DataLoader {
	return &jsonLoader{}
}

func (j *jsonLoader) Load(filePath string) ([]*orderedmap.OrderedMap, error) {
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
