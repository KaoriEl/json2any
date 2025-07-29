package processor_test

import (
	"github.com/KaoriEl/json2any/v2/internal/processor"
	"testing"
	"time"

	"github.com/iancoleman/orderedmap"
	"github.com/stretchr/testify/assert"
)

func TestProcessOrderedMaps(t *testing.T) {
	p := processor.NewProcessor()

	om := orderedmap.New()
	om.Set("int", 42.0)
	om.Set("float", 3.14)
	om.Set("bool", true)
	om.Set("date", "2023-07-29")
	om.Set("string", "hello")
	om.Set("nil", nil)

	records := []*orderedmap.OrderedMap{om}

	results, err := p.ProcessOrderedMaps(records, 2, nil)
	assert.NoError(t, err)
	assert.Len(t, results, 1)

	row := results[0]
	assert.Equal(t, int64(42), row["int"])
	assert.Equal(t, 3.14, row["float"])
	assert.Equal(t, true, row["bool"])
	_, ok := row["date"].(time.Time)
	assert.True(t, ok)
	assert.Equal(t, "hello", row["string"])
	assert.Equal(t, "", row["nil"])
}

func TestProcessStringMaps(t *testing.T) {
	p := processor.NewProcessor()

	records := []map[string]string{
		{
			"int":    "100",
			"float":  "3.14",
			"bool":   "true",
			"date":   "2023-07-29",
			"string": "text",
			"empty":  "",
		},
	}

	results, err := p.ProcessStringMaps(records, 2, nil)
	assert.NoError(t, err)
	assert.Len(t, results, 1)

	row := results[0]
	assert.Equal(t, int64(100), row["int"])
	assert.Equal(t, 3.14, row["float"])
	assert.Equal(t, "true", row["bool"]) // bool в normalizeStringValue не парсится специально
	_, ok := row["date"].(time.Time)
	assert.True(t, ok)
	assert.Equal(t, "text", row["string"])
	assert.Equal(t, "", row["empty"])
}
