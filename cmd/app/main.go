package main

import (
	"flag"
	"fmt"
	"os"
	"strings"

	"github.com/KaoriEl/json2xlsx/internal/utils"

	"github.com/KaoriEl/json2xlsx/internal/exporter"
	"github.com/KaoriEl/json2xlsx/internal/loader"
	"github.com/KaoriEl/json2xlsx/internal/processor"
)

func main() {
	filePath, outputFile, theme, maxWorkers := parseArgs()

	ld := loader.NewJSONLoader()
	pr := processor.NewProcessor()
	ex := exporter.NewExcelExporter()

	records, err := ld.Load(filePath)
	if err != nil {
		panic(err)
	}
	if len(records) == 0 {
		fmt.Println("Empty JSON")
		return
	}

	processed, err := pr.Process(records, maxWorkers)
	if err != nil {
		panic(err)
	}

	keys := records[0].Keys()
	if err := ex.ExportWithTheme(outputFile, keys, processed, theme); err != nil {
		panic(err)
	}

	fmt.Printf("Файл сохранён: %s\n", outputFile)
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
		fmt.Println("Неверная тема. Допустимые значения: black, green, red, purple, none")
		printHelp()
		os.Exit(1)
	}

	if *maxWorkers <= 0 {
		fmt.Println("max_workers должен быть положительным числом")
		printHelp()
		os.Exit(1)
	}

	return *filePath, *outputFile, strings.ToLower(*theme), *maxWorkers
}

func printHelp() {
	fmt.Println(`Использование:
  --input       Путь к входному JSON файлу (обязательно)
  --output      Путь к выходному XLSX файлу (по умолчанию "output.xlsx")
  --theme       Тема оформления: black, green, red, purple, none (по умолчанию black)
  --max_workers Максимальное количество горутин (int > 0, по умолчанию 20)
  --help        Показать справку
Пример:
  go run main.go --input data.json --output result.xlsx --theme green --max_workers 30`)
}
