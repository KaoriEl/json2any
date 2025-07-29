package loader

import (
	"github.com/xuri/excelize/v2"
	"os"
	"path/filepath"
	"testing"
)

func createTempFile(t *testing.T, content string, ext string) string {
	t.Helper()
	tmpFile, err := os.CreateTemp("", "testfile_*."+ext)
	if err != nil {
		t.Fatalf("failed to create temp file: %v", err)
	}
	defer tmpFile.Close()

	if _, err := tmpFile.WriteString(content); err != nil {
		t.Fatalf("failed to write to temp file: %v", err)
	}
	return tmpFile.Name()
}

func TestLoadJson(t *testing.T) {
	loader := NewLoader()

	jsonContent := `[{"key1": "value1", "key2": 123}]`
	file := createTempFile(t, jsonContent, "json")
	defer os.Remove(file)

	data, err := loader.LoadJson(file)
	if err != nil {
		t.Fatalf("LoadJson failed: %v", err)
	}
	if len(data) != 1 {
		t.Errorf("expected 1 record, got %d", len(data))
	}
	val, ok := data[0].Get("key1")
	if !ok || val != "value1" {
		t.Errorf("expected key1 to be 'value1', got %v", val)
	}
}

func TestLoadJson_BadFile(t *testing.T) {
	loader := NewLoader()

	_, err := loader.LoadJson("nonexistent.json")
	if err == nil {
		t.Errorf("expected error for nonexistent file")
	}
}

func TestLoadCSV(t *testing.T) {
	loader := NewLoader()

	csvContent := "name,age\nAlice,30\nBob,25\n"
	file := createTempFile(t, csvContent, "csv")
	defer os.Remove(file)

	records, err := loader.LoadCSV(file)
	if err != nil {
		t.Fatalf("LoadCSV failed: %v", err)
	}

	if len(records) != 2 {
		t.Errorf("expected 2 records, got %d", len(records))
	}

	if records[0]["name"] != "Alice" || records[0]["age"] != "30" {
		t.Errorf("unexpected first record: %+v", records[0])
	}
}

func TestLoadCSV_EmptyFile(t *testing.T) {
	loader := NewLoader()

	file := createTempFile(t, "", "csv")
	defer os.Remove(file)

	_, err := loader.LoadCSV(file)
	if err == nil {
		t.Errorf("expected error for empty CSV file")
	}
}

func TestLoadTXT(t *testing.T) {
	loader := NewLoader()

	txtContent := "col1\tcol2\nval1\tval2\nval3\tval4\n"
	file := createTempFile(t, txtContent, "txt")
	defer os.Remove(file)

	records, err := loader.LoadTXT(file)
	if err != nil {
		t.Fatalf("LoadTXT failed: %v", err)
	}

	if len(records) != 2 {
		t.Errorf("expected 2 records, got %d", len(records))
	}

	if records[1]["col2"] != "val4" {
		t.Errorf("unexpected second record: %+v", records[1])
	}
}

func TestLoadTXT_BadFile(t *testing.T) {
	loader := NewLoader()

	_, err := loader.LoadTXT("nonexistent.txt")
	if err == nil {
		t.Errorf("expected error for nonexistent file")
	}
}

func TestLoadXLSX(t *testing.T) {
	loader := NewLoader()

	// Создадим временный xlsx-файл с помощью excelize
	file := filepath.Join(os.TempDir(), "testfile.xlsx")
	f := excelize.NewFile()
	index, _ := f.NewSheet("Sheet1")
	f.SetCellValue("Sheet1", "A1", "name")
	f.SetCellValue("Sheet1", "B1", "age")
	f.SetCellValue("Sheet1", "A2", "John")
	f.SetCellValue("Sheet1", "B2", "45")
	f.SetActiveSheet(index)
	if err := f.SaveAs(file); err != nil {
		t.Fatalf("failed to save xlsx: %v", err)
	}
	defer os.Remove(file)

	records, err := loader.LoadXLSX(file)
	if err != nil {
		t.Fatalf("LoadXLSX failed: %v", err)
	}

	if len(records) != 1 {
		t.Errorf("expected 1 record, got %d", len(records))
	}
	if records[0]["name"] != "John" || records[0]["age"] != "45" {
		t.Errorf("unexpected record: %+v", records[0])
	}
}

func TestLoadXLSX_FileNotFound(t *testing.T) {
	loader := NewLoader()

	_, err := loader.LoadXLSX("nonexistent.xlsx")
	if err == nil {
		t.Errorf("expected error for nonexistent xlsx file")
	}
}
