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
		Usage: "Конвертер JSON → Excel (.xlsx) с темами и параллельной обработкой",
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:     "input",
				Aliases:  []string{"i"},
				Usage:    "Путь к входному JSON файлу (обязательно)",
				Required: true,
			},
			&cli.StringFlag{
				Name:    "output",
				Aliases: []string{"o"},
				Usage:   "Путь к выходному XLSX файлу (по умолчанию random.xlsx)",
			},
			&cli.StringFlag{
				Name:  "theme",
				Usage: "Тема оформления: black, green, red, purple, none",
				Value: "black",
			},
			&cli.IntFlag{
				Name:  "max_workers",
				Usage: "Максимальное количество горутин (int > 0)",
				Value: 20,
			},
			&cli.BoolFlag{
				Name:  "show_metrics",
				Usage: "Выводить метрики по окончании работы",
				Value: false,
			},
		},
		Action: action.Run,
	}

	if err := app.Run(os.Args); err != nil {
		log.Fatal(err)
	}
}
