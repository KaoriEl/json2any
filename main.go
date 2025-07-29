package main

import (
	"fmt"
	"github.com/KaoriEl/json2xlsx/internal/definitions"
	"os"

	"github.com/alperdrsnn/clime"

	"github.com/KaoriEl/json2xlsx/internal/action"
	"github.com/urfave/cli/v2"
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
					&cli.StringFlag{Name: "output", Aliases: []string{"o"}, Usage: "Path to output XLSX file", Value: definitions.DefaultOutputXLSX},
					&cli.StringFlag{Name: "format", Usage: "Output format: xlsx, csv, txt", Value: definitions.DefaultFormatXLSX},
					&cli.StringFlag{Name: "theme", Usage: "Theme: black, green, red, purple, none | only for xlsx", Value: definitions.DefaultTheme},
					&cli.IntFlag{Name: "max_workers", Usage: "Max goroutines", Value: definitions.DefaultMaxWorkers},
					&cli.BoolFlag{Name: "show_metrics", Usage: "Show metrics after completion", Value: definitions.DefaultShowMetrics},
				},
				Action: func(c *cli.Context) error {
					format := c.String("format")
					switch format {
					case definitions.DefaultFormatXLSX:
						return action.RunExport(c)
					case definitions.DefaultFormatCSV:
						return action.RunExportCSV(c)
					case definitions.DefaultFormatTXT:
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
					&cli.StringFlag{Name: "output", Aliases: []string{"o"}, Usage: "Path to output JSON file", Value: definitions.DefaultOutputJSON},
					&cli.StringFlag{Name: "format", Usage: "Input format: xlsx, csv, txt", Value: definitions.DefaultFormatXLSX},
					&cli.IntFlag{Name: "max_workers", Usage: "Max goroutines", Value: definitions.DefaultMaxWorkers},
					&cli.BoolFlag{Name: "show_metrics", Usage: "Show metrics after completion", Value: definitions.DefaultShowMetrics},
				},
				Action: func(c *cli.Context) error {
					format := c.String("format")
					switch format {
					case definitions.DefaultFormatXLSX:
						return action.RunImportXLSX(c)
					case definitions.DefaultFormatCSV:
						return action.RunImportCSV(c)
					case definitions.DefaultFormatTXT:
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
