package exporter

import (
	"fmt"

	"github.com/KaoriEl/json2any/v2/internal/models"
	"github.com/alperdrsnn/clime"
	"github.com/xuri/excelize/v2"
)

type ExcelExporter interface {
	ExportWithTheme(filePath string, keys []string, rows []models.ProcessedRow, theme string, progress func(current, total int)) error
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

func (e *excelExporter) ExportWithTheme(
	filePath string,
	keys []string,
	rows []models.ProcessedRow,
	theme string,
	progress func(current, total int),
) error {
	f := excelize.NewFile()
	sheet := "Sheet1"

	streamWriter, err := f.NewStreamWriter(sheet)
	if err != nil {
		return fmt.Errorf("failed to create stream writer: %w", err)
	}

	if err := setColumnsWidth(streamWriter, keys, rows); err != nil {
		return fmt.Errorf("error setting column widths: %w", err)
	}

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
		return fmt.Errorf("failed to create header style: %w", err)
	}

	altRowStyle, normalRowStyle, _, err := createRowStyles(f)
	if err != nil {
		return fmt.Errorf("failed to create row styles: %w", err)
	}

	headers := make([]interface{}, len(keys))
	for i, key := range keys {
		headers[i] = excelize.Cell{StyleID: headerStyle, Value: key}
	}
	if err := streamWriter.SetRow("A1", headers); err != nil {
		return fmt.Errorf("failed to write header: %w", err)
	}

	total := len(rows)
	for i, row := range rows {
		rowData := make([]interface{}, len(keys))
		for j, key := range keys {
			rowData[j] = row[key]
		}
		cellName, _ := excelize.CoordinatesToCellName(1, i+2)
		if err := streamWriter.SetRow(cellName, rowData); err != nil {
			return fmt.Errorf("failed to write row %d: %w", i+2, err)
		}
		if progress != nil {
			progress(i+1, total)
		}
	}

	if err := streamWriter.Flush(); err != nil {
		return fmt.Errorf("error finalizing stream write: %w", err)
	}

	colsCount := len(keys)
	lastRow := total + 1
	for rowIdx := 2; rowIdx <= lastRow; rowIdx++ {
		startCell, _ := excelize.CoordinatesToCellName(1, rowIdx)
		endCell, _ := excelize.CoordinatesToCellName(colsCount, rowIdx)
		style := normalRowStyle
		if (rowIdx-1)%2 == 1 {
			style = altRowStyle
		}
		if err := f.SetCellStyle(sheet, startCell, endCell, style); err != nil {
			return fmt.Errorf("failed to apply style to row %d: %w", rowIdx, err)
		}
	}

	clime.SuccessLine("Excel export completed, saving file...")
	if err := f.SaveAs(filePath); err != nil {
		return fmt.Errorf("failed to save file: %w", err)
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

func setColumnsWidth(streamWriter *excelize.StreamWriter, keys []string, rows []models.ProcessedRow) error {
	colWidths := make([]float64, len(keys))

	for colIdx, key := range keys {
		maxWidth := float64(len([]rune(key)))
		for _, row := range rows {
			val := row[key]
			str := fmt.Sprintf("%v", val)
			if w := float64(len([]rune(str))); w > maxWidth {
				maxWidth = w
			}
		}

		colWidths[colIdx] = maxWidth*1.2 + 2
		if colWidths[colIdx] < 10 {
			colWidths[colIdx] = 10
		}

		if colWidths[colIdx] > 255 {
			colWidths[colIdx] = 255
		}
	}

	for i, width := range colWidths {
		if err := streamWriter.SetColWidth(i+1, i+1, width); err != nil {
			return fmt.Errorf("failed to set width for column %d: %w", i+1, err)
		}
	}
	return nil
}
