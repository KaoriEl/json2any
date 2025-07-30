<p align="center">
  <img src="img/logo.jpg" alt="JSON to Any Exporter Logo" width="250"/>
</p>

# 🔧 JSON to Any Exporter

📘 Dokumentation in weiteren Sprachen verfügbar:

* 🇬🇧 [English](README.md)
* 🇨🇳 [简体中文](README.zh.md)
* 🇷🇺 [Русский](README.ru.md)
* 🇪🇸 [Español](README.es.md)

Dies ist ein **CLI-Tool** zum Konvertieren zwischen **JSON** und den Formaten **Excel (.xlsx), CSV und TXT** mit Unterstützung für **Themen**, **Datentypformatierung**, **Parallelverarbeitung** und **bidirektionale Konvertierung**.

---

## ✨ Funktionen

* 🚀 **Exportieren**: JSON-Dateien in `.xlsx`-, `.csv`- und `.txt`-Formate umwandeln
* 🔄 **Importieren**: `.xlsx`-, `.csv`- und `.txt`-Dateien in JSON umwandeln
* 🎨 Unterstützt Themen: `black`, `green`, `red`, `purple`, `none`
* 🔢 Korrekte Formatierung für **Zahlen**, **Datum**, **Strings** und **Boolean**
* ⚙️ Parallele Verarbeitung mit konfigurierbarer Anzahl von Arbeitern
* 📊 Optionale Anzeige von Leistungsmetriken nach Abschluss

---

## 🛠️ Build

```bash
go build -o json2any ./main.go
```

---

## 🚀 Installation (für systemweiten Zugriff)

```bash
go install github.com/KaoriEl/json2any/v2@latest
```

---

## 📋 Verwendung

### JSON in andere Formate exportieren

![example.png](img/example_xlsx.png)

Konvertieren Sie JSON-Daten in `.xlsx`-, `.csv`- oder `.txt`-Formate mit anpassbaren Optionen.

#### Beispiel: Export als XLSX

```bash
json2any export -i example.json -o result.xlsx --format=xlsx --theme=green --max_workers=100 --show_metrics=true
```

#### Beispiel: Export als CSV

```bash
json2any export -i example.json -o result.csv --format=csv --max_workers=10
```

#### Beispiel: Export als TXT

```bash
json2any export -i example.json -o result.txt --format=txt --max_workers=5
```

---

### Importieren von anderen Formaten nach JSON

![example\_import\_xlsx.png](img/example_import_txt.png)

Konvertieren Sie `.xlsx`-, `.csv`- oder `.txt`-Dateien in JSON.

#### Beispiel: Import von XLSX nach JSON

```bash
json2any import -i example.xlsx -o result.json --format=xlsx --max_workers=10
```

#### Beispiel: Import von CSV nach JSON

```bash
json2any import -i example.csv -o result.json --format=csv --max_workers=10
```

#### Beispiel: Import von TXT nach JSON

```bash
json2any import -i example.txt -o result.json --format=txt --max_workers=10
```

---

## ❓ Hilfe

```bash
json2any --help
```

---

## ⚙️ CLI-Flags

| Flag             | Beschreibung                                                                                                                         |
| ---------------- | ------------------------------------------------------------------------------------------------------------------------------------ |
| `--input, -i`    | **(Erforderlich)** Pfad zur Eingabedatei (JSON für Export, XLSX/CSV/TXT für Import).                                                 |
| `--output, -o`   | Pfad zur Ausgabedatei. Standard: `random.xlsx` (für Export) oder `output.json` (für Import).                                         |
| `--format`       | Ausgabeformat für Export: `xlsx`, `csv` oder `txt`. Eingabeformat für Import: `xlsx`, `csv` oder `txt`. Standard: `xlsx` für beides. |
| `--theme`        | Tabellen-Thema: `black`, `green`, `red`, `purple`, `none`. Standard: `black`. (Nur Export)                                           |
| `--max_workers`  | Anzahl paralleler Arbeiter. Ganzzahl > 0. Standard: `20`.                                                                            |
| `--show_metrics` | Leistungsmetriken nach Abschluss anzeigen. Standard: `false`.                                                                        |