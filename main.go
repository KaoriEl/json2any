package main

import (
	"fmt"
	"os"

	"github.com/alperdrsnn/clime"

	"github.com/KaoriEl/json2xlsx/internal/action"
	"github.com/urfave/cli/v2"
)

// Константы для значений флагов и форматов.
const (
	DefaultOutputXLSX  = "random.xlsx"
	DefaultOutputJSON  = "output.json"
	DefaultFormatXLSX  = "xlsx"
	DefaultFormatCSV   = "csv"
	DefaultFormatTXT   = "txt"
	DefaultTheme       = "black"
	DefaultMaxWorkers  = 20
	DefaultShowMetrics = false
)

func main() {
	app := &cli.App{
		Name:  "json2xlsx",
		Usage: "Convert between JSON and Excel|CSV|TXT formats",
		Commands: []*cli.Command{
			{
				Name:  "export",
				Usage: "Convert JSON to XLSX",
				Flags: []cli.Flag{
					&cli.StringFlag{Name: "input", Aliases: []string{"i"}, Usage: "Path to input JSON file", Required: true},
					&cli.StringFlag{Name: "output", Aliases: []string{"o"}, Usage: "Path to output XLSX file", Value: DefaultOutputXLSX},
					&cli.StringFlag{Name: "format", Usage: "Output format: xlsx, csv, txt", Value: DefaultFormatXLSX},
					&cli.StringFlag{Name: "theme", Usage: "Theme: black, green, red, purple, none", Value: DefaultTheme},
					&cli.IntFlag{Name: "max_workers", Usage: "Max goroutines", Value: DefaultMaxWorkers},
					&cli.BoolFlag{Name: "show_metrics", Usage: "Show metrics after completion", Value: DefaultShowMetrics},
				},
				Action: func(c *cli.Context) error {
					format := c.String("format")
					switch format {
					case DefaultFormatXLSX:
						return action.RunExport(c)
					case DefaultFormatCSV:
						return action.RunExportCSV(c)
					case DefaultFormatTXT:
						return action.RunExportTXT(c)
					default:
						return fmt.Errorf("unsupported format: %s", format)
					}
				},
			},
			{
				Name:  "import",
				Usage: "Convert between Excel|CSV|TXT and JSON formats",
				Flags: []cli.Flag{
					&cli.StringFlag{Name: "input", Aliases: []string{"i"}, Usage: "Path to input XLSX file", Required: true},
					&cli.StringFlag{Name: "output", Aliases: []string{"o"}, Usage: "Path to output JSON file", Value: DefaultOutputJSON},
					&cli.StringFlag{Name: "format", Usage: "Input format: xlsx, csv, txt", Value: DefaultFormatXLSX},
					&cli.IntFlag{Name: "max_workers", Usage: "Max goroutines", Value: DefaultMaxWorkers},
					&cli.BoolFlag{Name: "show_metrics", Usage: "Show metrics after completion", Value: DefaultShowMetrics},
				},
				Action: func(c *cli.Context) error {
					format := c.String("format")
					switch format {
					case DefaultFormatXLSX:
						return action.RunImportXLSX(c)
					case DefaultFormatCSV:
						return action.RunImportCSV(c)
					case DefaultFormatTXT:
						return action.RunImportTXT(c)
					default:
						return fmt.Errorf("unsupported format: %s", format)
					}
				},
			},
		},
	}

	if err := app.Run(os.Args); err != nil {
		clime.ErrorBanner(err.Error())
	}
}
