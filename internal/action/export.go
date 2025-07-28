package action

import (
	"fmt"
	"strings"
	"sync"
	"time"

	"github.com/KaoriEl/json2xlsx/internal/exporter"
	"github.com/KaoriEl/json2xlsx/internal/loader"
	"github.com/KaoriEl/json2xlsx/internal/processor"
	"github.com/KaoriEl/json2xlsx/internal/utils"
	"github.com/alperdrsnn/clime"
	"github.com/urfave/cli/v2"
)

func Run(c *cli.Context) error {
	clime.Clear()
	clime.NewBanner("JSON ➜ XLSX Exporter", clime.BannerSuccess).
		WithStyle(clime.BannerStyleDouble).
		WithColor(clime.BrightCyanColor).
		WithBorderColor(clime.BrightBlueColor).
		Println()

	input := c.String("input")
	output := c.String("output")
	if output == "" {
		output = utils.RandomFileName() + ".xlsx"
	}
	theme := strings.ToLower(c.String("theme"))
	maxWorkers := c.Int("max_workers")

	if err := validateTheme(theme); err != nil {
		return cli.Exit(err.Error(), 1)
	}
	if maxWorkers <= 0 {
		return cli.Exit("max_workers должен быть положительным числом", 1)
	}

	start := time.Now()

	clime.InfoLine("Загрузка JSON-файла...")
	loadStart := time.Now()

	ld := loader.NewJSONLoader()
	records, err := ld.Load(input)
	if err != nil {
		clime.ErrorBox("Ошибка загрузки", err.Error())
		return cli.Exit("", 1)
	}
	if len(records) == 0 {
		clime.ErrorLine("JSON-файл пуст")
		return cli.Exit("", 1)
	}
	loadDuration := time.Since(loadStart)

	clime.SuccessLine(fmt.Sprintf("Загружено %d записей", len(records)))

	bar := clime.NewProgressBar(int64(len(records))).
		WithLabel("Обработка данных").
		WithStyle(clime.ProgressStyleModern).
		WithColor(clime.BrightCyanColor).
		ShowRate(true)

	pr := processor.NewProcessor()

	var mu sync.Mutex
	total := len(records)
	step := total / 1000
	if step == 0 {
		step = 1
	}

	processStart := time.Now()

	processed, err := pr.Process(records, maxWorkers, func(current int) {
		mu.Lock()
		defer mu.Unlock()

		bar.Set(int64(current))
		if current%step == 0 || current == total {
			bar.Print()
		}
	})
	if err != nil {
		clime.ErrorBox("Ошибка обработки", err.Error())
		return cli.Exit("", 1)
	}
	processDuration := time.Since(processStart)

	ex := exporter.NewExcelExporter()
	keys := records[0].Keys()

	exportBar := clime.NewProgressBar(int64(len(processed))).
		WithLabel("Экспорт в Excel").
		WithStyle(clime.ProgressStyleModern).
		WithColor(clime.BrightGreenColor).
		ShowRate(true)

	exportStart := time.Now()

	err = ex.ExportWithTheme(output, keys, processed, theme, func(current, total int) {
		mu.Lock()
		defer mu.Unlock()

		exportBar.Set(int64(current))
		if current%100 == 0 || current == total {
			exportBar.Print()
		}
	})
	if err != nil {
		clime.ErrorBox("Ошибка экспорта", err.Error())
		return cli.Exit("", 1)
	}
	exportDuration := time.Since(exportStart)

	clime.NewBox().
		WithTitle("✅ Успех").
		WithBorderColor(clime.GreenColor).
		WithStyle(clime.BoxStyleRounded).
		AddLine("Файл успешно создан: " + output).
		Println()

	if c.Bool("show_metrics") {
		showMetric(
			output,
			start,
			len(records),   // recordsIn
			len(processed), // recordsOut
			loadDuration,
			processDuration,
			exportDuration,
		)
	}

	return nil
}

func validateTheme(theme string) error {
	validThemes := map[string]bool{
		"black":  true,
		"green":  true,
		"red":    true,
		"purple": true,
		"none":   true,
	}
	if !validThemes[theme] {
		return fmt.Errorf("некорректная тема: %s. Допустимые: black, green, red, purple, none", theme)
	}
	return nil
}

func showMetric(output string, start time.Time, recordsIn, recordsOut int, loadDuration, processDuration, exportDuration time.Duration) {
	fileSize, err := utils.FileSizeMB(output)
	if err != nil {
		fileSize = "не удалось определить размер"
	}

	clime.NewTable().
		AddColumn("Метрика").
		AddColumn("Значение").
		SetColumnColor(1, clime.Info).
		AddRow("Общее время", time.Since(start).Truncate(time.Millisecond).String()).
		AddRow("Время загрузки JSON", loadDuration.String()).
		AddRow("Время обработки", processDuration.String()).
		AddRow("Время экспорта", exportDuration.String()).
		AddRow("Кол-во записей (вход)", fmt.Sprintf("%d", recordsIn)).
		AddRow("Кол-во записей (обработано)", fmt.Sprintf("%d", recordsOut)).
		AddRow("Размер файла", fileSize).
		AddRow("Скорость обработки", fmt.Sprintf("%.2f записей/сек", float64(recordsOut)/processDuration.Seconds())).
		Print()
}
