package main

import (
	"flag"
	"fmt"
	"os"
	"strings"
	"time"

	"github.com/KaoriEl/json2xlsx/internal/exporter"
	"github.com/KaoriEl/json2xlsx/internal/loader"
	"github.com/KaoriEl/json2xlsx/internal/processor"
	"github.com/KaoriEl/json2xlsx/internal/utils"
)

func main() {
	start := time.Now()

	filePath, outputFile, theme, maxWorkers := parseArgs()

	ld := loader.NewJSONLoader()
	pr := processor.NewProcessor()
	ex := exporter.NewExcelExporter()

	records, err := ld.Load(filePath)
	if err != nil {
		printErrorAndExit(fmt.Errorf("ошибка загрузки файла: %w", err))
	}

	if len(records) == 0 {
		_, _ = fmt.Fprintln(os.Stderr, "JSON-файл пуст")
		os.Exit(1)
	}

	processed, err := pr.Process(records, maxWorkers)
	if err != nil {
		printErrorAndExit(fmt.Errorf("ошибка обработки данных: %w", err))
	}

	keys := records[0].Keys()
	if err := ex.ExportWithTheme(outputFile, keys, processed, theme); err != nil {
		printErrorAndExit(fmt.Errorf("ошибка экспорта в Excel: %w", err))
	}

	fmt.Printf("Файл сохранён: %s\n", outputFile)
	fmt.Printf("Время выполнения: %s\n", time.Since(start).Truncate(time.Millisecond))
}

func parseArgs() (string, string, string, int) {
	filePath := flag.String("input", "", "Путь к входному JSON файлу (обязательно)")
	outputFile := flag.String("output", utils.RandomFileName()+".xlsx", "Путь к выходному XLSX файлу")
	theme := flag.String("theme", "black", "Тема оформления: black, green, red, purple, none")
	maxWorkers := flag.Int("max_workers", 20, "Максимальное количество горутин (int > 0)")
	help := flag.Bool("help", false, "Показать справку")

	flag.Parse()

	if *help || *filePath == "" {
		printHelp()
		os.Exit(0)
	}

	switch strings.ToLower(*theme) {
	case "black", "green", "red", "purple", "none":
	default:
		fmt.Fprintln(os.Stderr, "Неверная тема. Допустимые значения: black, green, red, purple, none")
		printHelp()
		os.Exit(1)
	}

	if *maxWorkers <= 0 {
		fmt.Fprintln(os.Stderr, "max_workers должен быть положительным числом")
		printHelp()
		os.Exit(1)
	}

	return *filePath, *outputFile, strings.ToLower(*theme), *maxWorkers
}

func printHelp() {
	fmt.Println(`Использование:
  --input       Путь к входному JSON файлу (обязательно)
  --output      Путь к выходному XLSX файлу (по умолчанию random.xlsx)
  --theme       Тема оформления: black, green, red, purple, none (по умолчанию black)
  --max_workers Максимальное количество горутин (int > 0, по умолчанию 20)
  --help        Показать справку

Пример:
  json2xlsx --input=data.json --output=result.xlsx --theme=green --max_workers=30`)
}

func printErrorAndExit(err error) {
	_, _ = fmt.Fprintf(os.Stderr, "Ошибка: %v\n", err)
	os.Exit(1)
}
