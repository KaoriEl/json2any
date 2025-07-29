package action

import (
	"fmt"
	"os"
	"sync"
	"time"

	"github.com/KaoriEl/json2xlsx/internal/exporter"

	"github.com/KaoriEl/json2xlsx/internal/loader"
	"github.com/KaoriEl/json2xlsx/internal/processor"
	"github.com/KaoriEl/json2xlsx/internal/utils"
	"github.com/alperdrsnn/clime"
	"github.com/urfave/cli/v2"
)

func RunExportTXT(c *cli.Context) error {
	clime.Clear()
	clime.NewBanner("JSON ➜ TXT Exporter", clime.BannerSuccess).
		WithStyle(clime.BannerStyleDouble).
		WithColor(clime.BrightCyanColor).
		WithBorderColor(clime.BrightBlueColor).
		Println()

	input := c.String("input")
	output := c.String("output")
	if output == "" {
		output = utils.RandomFileName() + ".txt"
	}
	maxWorkers := c.Int("max_workers")

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

	ld := loader.NewLoader()
	records, err := ld.LoadJson(input)
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

	processed, err := pr.ProcessOrderedMaps(records, maxWorkers, func(current int) {
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

	ex := exporter.NewTXTExporter() // <-- твой TXT экспортер (нужно реализовать)

	exportBar := clime.NewProgressBar(int64(len(processed))).
		WithLabel("Exporting to TXT").
		WithStyle(clime.ProgressStyleModern).
		WithColor(clime.BrightGreenColor).
		ShowRate(true)

	exportStart := time.Now()

	err = ex.Export(output, processed, func(current, total int) {
		mu.Lock()
		defer mu.Unlock()

		exportBar.Set(int64(current))
		if current%100 == 0 || current == total {
			exportBar.Print()
		}
	})
	if err != nil {
		clime.ErrorLine(err.Error())
		time.Sleep(2 * time.Second)
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
			len(records),
			len(processed),
			loadDuration,
			processDuration,
			exportDuration,
		)
	}

	return nil
}
