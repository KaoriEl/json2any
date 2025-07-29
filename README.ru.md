# JSON to XLSX Exporter

📘 Documentation available in other languages: [Русский](README.ru.md)

A CLI tool for converting JSON files into Excel (.xlsx) format with support for theming, data type formatting, and parallel processing.

---

## Features

* Convert arrays of JSON objects into `.xlsx` spreadsheets.
* Supports themes: `black`, `green`, `red`, `purple`, `none`.
* Correct formatting for numbers, dates, strings, and booleans.
* Parallel processing with configurable worker count.
* Optional performance metrics output after completion.

---

## Build

```bash
go build -o json2xlsx ./cmd/app/main.go
```

---

## Installation (for system-wide access)

```bash
sudo cp json2xlsx /usr/local/bin/
```

---

## Usage

### With `go run`:

```bash
go run ./cmd/app/main.go -i example.json -o result.xlsx --theme=green --max_workers=100 --show_metrics=true
```

### With compiled binary in current directory:

```bash
./json2xlsx -i example.json -o result.xlsx --theme=green --max_workers=10
```

### From anywhere (if installed system-wide):

```bash
json2xlsx -i example.json -o result.xlsx --theme=green --max_workers=10
```

---

## Help

```bash
json2xlsx --help
```

---

## CLI Flags

| Flag             | Description                                                               |
| ---------------- | ------------------------------------------------------------------------- |
| `--input, -i`    | **(Required)** Path to the input JSON file.                               |
| `--output, -o`   | Path to the output XLSX file. Default: `random.xlsx`.                     |
| `--theme`        | Table theme: `black`, `green`, `red`, `purple`, `none`. Default: `black`. |
| `--max_workers`  | Number of parallel workers. Integer > 0. Default: `20`.                   |
| `--show_metrics` | Show processing metrics after completion. Default: `false`.               |

