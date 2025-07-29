package exporter_test

import (
	"os"
	"path/filepath"
	"strings"
	"testing"
	"time"

	"github.com/KaoriEl/json2any/v2/internal/exporter"
	"github.com/KaoriEl/json2any/v2/internal/models"
)

func TestCSVExporter_Export(t *testing.T) {
	exp := exporter.NewCSVExporter()

	tmpFile := filepath.Join(os.TempDir(), "test.csv")
	defer os.Remove(tmpFile)

	data := []models.ProcessedRow{
		{"Name": "Alice", "Age": 30, "Active": true},
		{"Name": "Bob", "Age": 25, "Active": false},
	}

	keys := []string{"Name", "Age", "Active"}

	progressCalls := 0
	err := exp.Export(tmpFile, keys, data, func(current, total int) {
		progressCalls++
		if current > total {
			t.Errorf("progress current %d cannot be > total %d", current, total)
		}
	})
	if err != nil {
		t.Fatalf("Export failed: %v", err)
	}

	if progressCalls != len(data) {
		t.Errorf("expected progressCalls %d, got %d", len(data), progressCalls)
	}

	content, err := os.ReadFile(tmpFile)
	if err != nil {
		t.Fatalf("failed to read exported file: %v", err)
	}

	text := string(content)
	if !strings.Contains(text, "Name,Age,Active") {
		t.Error("CSV header not found")
	}
	if !strings.Contains(text, "Alice,30,true") {
		t.Error("CSV content for Alice missing or incorrect")
	}
}

func TestJSONExporter_Export(t *testing.T) {
	exp := exporter.NewJSONExporter()

	tmpFile := filepath.Join(os.TempDir(), "test.json")
	defer os.Remove(tmpFile)

	data := []models.ProcessedRow{
		{"Name": "Alice", "Age": 30},
		{"Name": "Bob", "Age": 25},
	}

	err := exp.Export(tmpFile, data)
	if err != nil {
		t.Fatalf("Export failed: %v", err)
	}

	content, err := os.ReadFile(tmpFile)
	if err != nil {
		t.Fatalf("failed to read exported file: %v", err)
	}

	text := string(content)
	if !strings.Contains(text, `"Name": "Alice"`) {
		t.Error("JSON content missing Alice")
	}
	if !strings.Contains(text, `"Age": 25`) {
		t.Error("JSON content missing Bob's age")
	}
}

func TestTXTExporter_Export(t *testing.T) {
	exp := exporter.NewTXTExporter()

	tmpFile := filepath.Join(os.TempDir(), "test.txt")
	defer os.Remove(tmpFile)

	now := time.Now()
	data := []models.ProcessedRow{
		{"Name": "Alice", "Registered": now},
		{"Name": "Bob", "Registered": now.Add(-24 * time.Hour)},
	}

	progressCalls := 0
	err := exp.Export(tmpFile, data, func(current, total int) {
		progressCalls++
		if current > total {
			t.Errorf("progress current %d cannot be > total %d", current, total)
		}
	})
	if err != nil {
		t.Fatalf("Export failed: %v", err)
	}

	if progressCalls != len(data) {
		t.Errorf("expected progressCalls %d, got %d", len(data), progressCalls)
	}

	content, err := os.ReadFile(tmpFile)
	if err != nil {
		t.Fatalf("failed to read exported file: %v", err)
	}

	text := string(content)
	if !strings.Contains(text, "Name=Alice") {
		t.Error("TXT content missing Alice")
	}
	if !strings.Contains(text, "Registered=") {
		t.Error("TXT content missing Registered field")
	}
}
