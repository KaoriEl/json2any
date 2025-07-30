<p align="center">
  <img src="img/logo.jpg" alt="JSON to Any Exporter Logo" width="250"/>
</p>

# 🔧 JSON to Any Exporter

📘 Documentation available in other languages:

* 🇷🇺 [Русский](README.ru.md)
* 🇨🇳 [简体中文](README.zh.md)
* 🇪🇸 [Español](README.es.md)
* 🇩🇪 [Deutsch](README.de.md)

A **CLI tool** for converting between **JSON** and **Excel (.xlsx), CSV, and TXT** formats with support for **theming**, **data type formatting**, **parallel processing**, and **bidirectional conversion**.

---

## ✨ Features

* 🚀 **Export**: Convert JSON files into `.xlsx`, `.csv`, and `.txt` formats
* 🔄 **Import**: Convert `.xlsx`, `.csv`, and `.txt` files into JSON format
* 🎨 Supports themes: `black`, `green`, `red`, `purple`, `none`
* 🔢 Correct formatting for **numbers**, **dates**, **strings**, and **booleans**
* ⚙️ Parallel processing with configurable worker count
* 📊 Optional performance metrics output after completion

---

## 🛠️ Build

```bash
go build -o json2any ./main.go
```

---

## 🚀 Installation (for system-wide access)

```bash
go install github.com/KaoriEl/json2any/v2@latest
```

---

## 📋 Usage

### Exporting JSON to Other Formats

![example.png](img/example_xlsx.png)

Convert JSON data into `.xlsx`, `.csv`, or `.txt` formats with customizable options.

#### Example: Export to XLSX

```bash
json2any export -i example.json -o result.xlsx --format=xlsx --theme=green --max_workers=100 --show_metrics=true
```

#### Example: Export to CSV

```bash
json2any export -i example.json -o result.csv --format=csv --max_workers=10
```

#### Example: Export to TXT

```bash
json2any export -i example.json -o result.txt --format=txt --max_workers=5
```

---

### Importing from Other Formats to JSON

![example\_import\_xlsx.png](img/example_import_txt.png)

Convert `.xlsx`, `.csv`, or `.txt` files into JSON format.

#### Example: Import from XLSX to JSON

```bash
json2any import -i example.xlsx -o result.json --format=xlsx --max_workers=10
```

#### Example: Import from CSV to JSON

```bash
json2any import -i example.csv -o result.json --format=csv --max_workers=10
```

#### Example: Import from TXT to JSON

```bash
json2any import -i example.txt -o result.json --format=txt --max_workers=10
```

---

## ❓ Help

```bash
json2any --help
```

---

## ⚙️ CLI Flags

| Flag             | Description                                                                                                                             |
| ---------------- | --------------------------------------------------------------------------------------------------------------------------------------- |
| `--input, -i`    | **(Required)** Path to the input file (JSON for export, XLSX/CSV/TXT for import).                                                       |
| `--output, -o`   | Path to the output file. Default: `random.xlsx` (for export) or `output.json` (for import).                                             |
| `--format`       | Output format for export: `xlsx`, `csv`, or `txt`. Input format for import: `xlsx`, `csv`, or `txt`. Default: `xlsx` for export/import. |
| `--theme`        | Table theme: `black`, `green`, `red`, `purple`, `none`. Default: `black`. (Export only)                                                 |
| `--max_workers`  | Number of parallel workers. Integer > 0. Default: `20`.                                                                                 |
| `--show_metrics` | Show processing metrics after completion. Default: `false`.                                                                             |
