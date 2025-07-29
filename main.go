package main

import (
	"log"
	"os"

	"github.com/KaoriEl/json2xlsx/internal/action"
	"github.com/urfave/cli/v2"
)

func main() {
	app := &cli.App{
		Name:  "json2xlsx",
		Usage: "JSON → Excel (.xlsx) converter with themes and parallel processing",
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:     "input",
				Aliases:  []string{"i"},
				Usage:    "Path to input JSON file (required)",
				Required: true,
			},
			&cli.StringFlag{
				Name:    "output",
				Aliases: []string{"o"},
				Usage:   "Path to output XLSX file (default random.xlsx)",
			},
			&cli.StringFlag{
				Name:  "theme",
				Usage: "Theme: black, green, red, purple, none",
				Value: "black",
			},
			&cli.IntFlag{
				Name:  "max_workers",
				Usage: "Maximum number of goroutines (int > 0)",
				Value: 20,
			},
			&cli.BoolFlag{
				Name:  "show_metrics",
				Usage: "Display metrics after completion",
				Value: false,
			},
		},
		Action: action.Run,
	}

	if err := app.Run(os.Args); err != nil {
		log.Fatal(err)
	}
}
