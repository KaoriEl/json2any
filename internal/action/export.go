package action

import (
	"fmt"
	"os"
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
		return cli.Exit("max_workers must be a positive number", 1)
	}

	start := time.Now()

	clime.InfoLine("Loading JSON file...")
	loadStart := time.Now()

	if _, err := os.Stat(input); os.IsNotExist(err) {
		clime.ErrorLine(fmt.Sprintf("File not found: %s", input))
		return cli.Exit("", 1)
	}

	ld := loader.NewJSONLoader()
	records, err := ld.Load(input)
	if err != nil {
		clime.ErrorLine(err.Error())
		return cli.Exit("", 1)
	}
	if len(records) == 0 {
		clime.ErrorLine("JSON file is empty")
		return cli.Exit("", 1)
	}
	loadDuration := time.Since(loadStart)

	clime.SuccessLine(fmt.Sprintf("Loaded %d records", len(records)))

	bar := clime.NewProgressBar(int64(len(records))).
		WithLabel("Processing data").
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
		clime.ErrorLine(err.Error())
		return cli.Exit("", 1)
	}
	processDuration := time.Since(processStart)

	ex := exporter.NewExcelExporter()
	keys := records[0].Keys()

	exportBar := clime.NewProgressBar(int64(len(processed))).
		WithLabel("Exporting to Excel").
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
		clime.ErrorLine(err.Error())
		time.Sleep(2 * time.Second) // Pause to show error
		return cli.Exit("", 1)
	}
	exportDuration := time.Since(exportStart)

	clime.NewBox().
		WithTitle("✅ Success").
		WithBorderColor(clime.GreenColor).
		WithStyle(clime.BoxStyleRounded).
		AddLine("File successfully created: " + output).
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
		return fmt.Errorf("invalid theme: %s. Allowed: black, green, red, purple, none", theme)
	}
	return nil
}

func showMetric(output string, start time.Time, recordsIn, recordsOut int, loadDuration, processDuration, exportDuration time.Duration) {
	fileSize, err := utils.FileSizeMB(output)
	if err != nil {
		fileSize = "failed to determine file size"
	}

	clime.NewTable().
		AddColumn("Metric").
		AddColumn("Value").
		SetColumnColor(1, clime.Info).
		AddRow("Total time", time.Since(start).Truncate(time.Millisecond).String()).
		AddRow("JSON load time", loadDuration.String()).
		AddRow("Processing time", processDuration.String()).
		AddRow("Export time", exportDuration.String()).
		AddRow("Records count (input)", fmt.Sprintf("%d", recordsIn)).
		AddRow("Records count (processed)", fmt.Sprintf("%d", recordsOut)).
		AddRow("File size", fileSize).
		AddRow("Processing speed", fmt.Sprintf("%.2f records/sec", float64(recordsOut)/processDuration.Seconds())).
		Print()
}
