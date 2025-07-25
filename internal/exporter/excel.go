package exporter

import (
	"fmt"
	"time"

	"github.com/KaoriEl/json2xlsx/internal/models"
	"github.com/xuri/excelize/v2"
)

type ExcelExporter interface {
	ExportWithTheme(filePath string, keys []string, rows []models.ProcessedRow, theme string) error
}

type excelExporter struct{}

func NewExcelExporter() ExcelExporter {
	return &excelExporter{}
}

const (
	colorBlack       = "#000000"
	colorWhite       = "#FFFFFF"
	colorGreen       = "#006400"
	colorLightGreen  = "#E0FFE0"
	colorDarkRed     = "#8B0000"
	colorLightRed    = "#FFD0D0"
	colorIndigo      = "#4B0082"
	colorLightPurple = "#E6E6FA"
	colorAltRow      = "#F0F0F0"
	colorBorder      = "#666666"
	colorNormalRow   = "#FFFFFF"
)

func (e *excelExporter) ExportWithTheme(filePath string, keys []string, rows []models.ProcessedRow, theme string) error {
	f := excelize.NewFile()
	sheet := "Sheet1"

	headerFillColor, headerFontColor := getHeaderColors(theme)

	headerStyle, err := f.NewStyle(&excelize.Style{
		Font:      &excelize.Font{Bold: true, Color: headerFontColor},
		Fill:      excelize.Fill{Type: "pattern", Color: []string{headerFillColor}, Pattern: 1},
		Alignment: &excelize.Alignment{Horizontal: "center", Vertical: "center"},
		Border: []excelize.Border{
			{Type: "left", Color: colorBorder, Style: 1},
			{Type: "right", Color: colorBorder, Style: 1},
			{Type: "top", Color: colorBorder, Style: 1},
			{Type: "bottom", Color: colorBorder, Style: 1},
		},
	})
	if err != nil {
		return fmt.Errorf("создание стиля заголовка: %w", err)
	}

	altRowStyle, normalRowStyle, dateStyle, err := createRowStyles(f)
	if err != nil {
		return err
	}

	if err := setHeaders(f, sheet, keys, headerStyle); err != nil {
		return err
	}

	if err := fillRows(f, sheet, keys, rows, altRowStyle, normalRowStyle, dateStyle); err != nil {
		return err
	}

	if err := setColumnsWidth(f, sheet, len(keys)); err != nil {
		return err
	}

	if err := f.SaveAs(filePath); err != nil {
		return fmt.Errorf("сохранение файла: %w", err)
	}

	return nil
}

func getHeaderColors(theme string) (fillColor, fontColor string) {
	switch theme {
	case "black":
		return colorBlack, colorWhite
	case "green":
		return colorGreen, colorLightGreen
	case "red":
		return colorDarkRed, colorLightRed
	case "purple":
		return colorIndigo, colorLightPurple
	case "none":
		return colorWhite, colorBlack
	default:
		return colorBlack, colorWhite
	}
}

func createRowStyles(f *excelize.File) (altRowStyle, normalRowStyle, dateStyle int, err error) {
	altRowStyle, err = f.NewStyle(&excelize.Style{
		Fill: excelize.Fill{Type: "pattern", Color: []string{colorAltRow}, Pattern: 1},
		Font: &excelize.Font{Color: colorBlack},
	})
	if err != nil {
		return
	}
	normalRowStyle, err = f.NewStyle(&excelize.Style{
		Fill: excelize.Fill{Type: "pattern", Color: []string{colorNormalRow}, Pattern: 1},
		Font: &excelize.Font{Color: colorBlack},
	})
	if err != nil {
		return
	}
	dateFormat := "DD.MM.YYYY"
	dateStyle, err = f.NewStyle(&excelize.Style{
		CustomNumFmt: &dateFormat,
	})
	return
}

func setHeaders(f *excelize.File, sheet string, keys []string, headerStyle int) error {
	for colIdx, header := range keys {
		cell, err := excelize.CoordinatesToCellName(colIdx+1, 1)
		if err != nil {
			return fmt.Errorf("получение имени ячейки: %w", err)
		}
		if err := f.SetCellValue(sheet, cell, header); err != nil {
			return fmt.Errorf("установка значения заголовка: %w", err)
		}
		if err := f.SetCellStyle(sheet, cell, cell, headerStyle); err != nil {
			return fmt.Errorf("установка стиля заголовка: %w", err)
		}
	}
	return nil
}

func fillRows(f *excelize.File, sheet string, keys []string, rows []models.ProcessedRow, altRowStyle, normalRowStyle, dateStyle int) error {
	for rowIdx, row := range rows {
		style := normalRowStyle
		if rowIdx%2 == 1 {
			style = altRowStyle
		}

		for colIdx, key := range keys {
			cell, err := excelize.CoordinatesToCellName(colIdx+1, rowIdx+2)
			if err != nil {
				return fmt.Errorf("получение имени ячейки: %w", err)
			}

			val := row[key]
			if err := setCellValueWithStyle(f, sheet, cell, val, style, dateStyle); err != nil {
				return fmt.Errorf("установка значения/стиля в ячейку %s: %w", cell, err)
			}
		}
	}
	return nil
}

func setCellValueWithStyle(f *excelize.File, sheet, cell string, val interface{}, style, dateStyle int) error {
	switch v := val.(type) {
	case int64, int32, int:
		if err := f.SetCellValue(sheet, cell, v); err != nil {
			return fmt.Errorf("SetCellValue: %w", err)
		}
		if err := f.SetCellStyle(sheet, cell, cell, style); err != nil {
			return fmt.Errorf("SetCellStyle: %w", err)
		}
	case float64, float32:
		if err := f.SetCellValue(sheet, cell, v); err != nil {
			return fmt.Errorf("SetCellValue: %w", err)
		}
		if err := f.SetCellStyle(sheet, cell, cell, style); err != nil {
			return fmt.Errorf("SetCellStyle: %w", err)
		}
	case string:
		if err := f.SetCellValue(sheet, cell, v); err != nil {
			return fmt.Errorf("SetCellValue: %w", err)
		}
		if err := f.SetCellStyle(sheet, cell, cell, style); err != nil {
			return fmt.Errorf("SetCellStyle: %w", err)
		}
	case bool:
		if err := f.SetCellValue(sheet, cell, v); err != nil {
			return fmt.Errorf("SetCellValue: %w", err)
		}
		if err := f.SetCellStyle(sheet, cell, cell, style); err != nil {
			return fmt.Errorf("SetCellStyle: %w", err)
		}
	case time.Time:
		if err := f.SetCellValue(sheet, cell, v); err != nil {
			return fmt.Errorf("SetCellValue: %w", err)
		}
		if err := f.SetCellStyle(sheet, cell, cell, dateStyle); err != nil {
			return fmt.Errorf("SetCellStyle: %w", err)
		}
	case nil:
		if err := f.SetCellValue(sheet, cell, ""); err != nil {
			return fmt.Errorf("SetCellValue: %w", err)
		}
		if err := f.SetCellStyle(sheet, cell, cell, style); err != nil {
			return fmt.Errorf("SetCellStyle: %w", err)
		}
	default:
		if err := f.SetCellValue(sheet, cell, fmt.Sprintf("%v", v)); err != nil {
			return fmt.Errorf("SetCellValue: %w", err)
		}
		if err := f.SetCellStyle(sheet, cell, cell, style); err != nil {
			return fmt.Errorf("SetCellStyle: %w", err)
		}
	}
	return nil
}

func setColumnsWidth(f *excelize.File, sheet string, colsCount int) error {
	for colIdx := 0; colIdx < colsCount; colIdx++ {
		column, err := excelize.ColumnNumberToName(colIdx + 1)
		if err != nil {
			return fmt.Errorf("получение имени столбца: %w", err)
		}
		if err := f.SetColWidth(sheet, column, column, 25); err != nil {
			return fmt.Errorf("установка ширины столбца %s: %w", column, err)
		}
	}
	return nil
}
