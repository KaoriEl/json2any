package action

import (
	"fmt"
	"github.com/KaoriEl/json2any/internal/definitions"
	"os"
	"strings"
	"sync"
	"time"

	"github.com/KaoriEl/json2any/internal/loader"

	"github.com/KaoriEl/json2any/internal/exporter"
	"github.com/KaoriEl/json2any/internal/processor"
	"github.com/KaoriEl/json2any/internal/utils"
	"github.com/alperdrsnn/clime"
	"github.com/urfave/cli/v2"
)

type LoadFunc func(string) ([]map[string]string, error)

func RunImportXLSX(c *cli.Context) error {
	return runImportGeneric(c, loader.NewLoader().LoadXLSX, strings.ToUpper(definitions.DefaultFormatXLSX))
}

func RunImportCSV(c *cli.Context) error {
	return runImportGeneric(c, loader.NewLoader().LoadCSV, strings.ToUpper(definitions.DefaultFormatCSV))
}

func RunImportTXT(c *cli.Context) error {
	return runImportGeneric(c, loader.NewLoader().LoadTXT, strings.ToUpper(definitions.DefaultFormatTXT))
}

func runImportGeneric(c *cli.Context, loadFunc LoadFunc, format string) error {
	clime.Clear()
	clime.NewBanner(fmt.Sprintf("Import %s to JSON", format), clime.BannerSuccess).
		WithStyle(clime.BannerStyleDouble).
		WithColor(clime.BrightGreenColor).
		WithBorderColor(clime.BrightBlueColor).
		Println()

	input := c.String("input")
	output := c.String("output")
	if output == "" {
		output = utils.RandomFileName() + ".json"
	}

	maxWorkers := c.Int("max_workers")
	if maxWorkers <= 0 {
		maxWorkers = 20
	}

	start := time.Now()

	if _, err := os.Stat(input); os.IsNotExist(err) {
		clime.ErrorLine(fmt.Sprintf("File not found: %s", input))
		return cli.Exit("", 1)
	}

	clime.InfoLine("Loading input file...")
	loadStart := time.Now()

	records, err := loadFunc(input)
	if err != nil {
		clime.ErrorLine(fmt.Sprintf("Failed to load input: %v", err))
		return cli.Exit("", 1)
	}
	if len(records) == 0 {
		clime.ErrorLine("Input file is empty")
		return cli.Exit("", 1)
	}
	loadDuration := time.Since(loadStart)

	clime.SuccessLine(fmt.Sprintf("Loaded %d records", len(records)))

	pr := processor.NewProcessor()
	bar := clime.NewProgressBar(int64(len(records))).
		WithLabel("Processing rows").
		WithStyle(clime.ProgressStyleModern).
		WithColor(clime.BrightGreenColor).
		ShowRate(true)

	var mu sync.Mutex
	total := len(records)
	step := total / 1000
	if step == 0 {
		step = 1
	}

	processStart := time.Now()

	processed, err := pr.ProcessStringMaps(records, maxWorkers, func(current int) {
		mu.Lock()
		defer mu.Unlock()

		bar.Set(int64(current))
		if current%step == 0 || current == total {
			bar.Print()
		}
	})
	if err != nil {
		clime.ErrorLine(fmt.Sprintf("Failed to process data: %v", err))
		return cli.Exit("", 1)
	}
	processDuration := time.Since(processStart)

	xp := exporter.NewJSONExporter()
	exportStart := time.Now()

	clime.InfoLine("Writing JSON file...")
	if err := xp.Export(output, processed); err != nil {
		clime.ErrorLine(fmt.Sprintf("Failed to write JSON: %v", err))
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
		showImportMetrics(
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

func showImportMetrics(output string, start time.Time, recordsIn, recordsOut int, loadDuration, processDuration, exportDuration time.Duration) {
	fileSize, err := utils.FileSizeMB(output)
	if err != nil {
		fileSize = "failed to determine file size"
	}

	clime.NewTable().
		AddColumn("Metric").
		AddColumn("Value").
		SetColumnColor(1, clime.Info).
		AddRow("Total time", time.Since(start).Truncate(time.Millisecond).String()).
		AddRow("XLSX load time", loadDuration.String()).
		AddRow("Processing time", processDuration.String()).
		AddRow("Export time", exportDuration.String()).
		AddRow("Records count (input)", fmt.Sprintf("%d", recordsIn)).
		AddRow("Records count (processed)", fmt.Sprintf("%d", recordsOut)).
		AddRow("File size", fileSize).
		AddRow("Processing speed", fmt.Sprintf("%.2f records/sec", float64(recordsOut)/processDuration.Seconds())).
		Print()
}
