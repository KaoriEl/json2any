# JSON zu Beliebigem Exporter

📘 Dokumentation ist in anderen Sprachen verfügbar:

* [English](README.md)
* [Русский](README.ru.md)
* [Español](README.es.md)
* [简体中文](README.zh.md)

Kommandozeilenwerkzeug zum Konvertieren zwischen JSON und Excel (.xlsx), CSV und TXT mit Unterstützung für Themen, Datenformatierung, parallele Verarbeitung und bidirektionale Konvertierung.

---

## Funktionen

* **Export**: Konvertieren von JSON-Dateien in die Formate `.xlsx`, `.csv` und `.txt`.
* **Import**: Konvertieren von `.xlsx`, `.csv` und `.txt` Dateien in das JSON-Format.
* Unterstützung für Themen: `black`, `green`, `red`, `purple`, `none`.
* Korrekte Formatierung von Zahlen, Daten, Zeichenketten und booleschen Werten.
* Parallele Verarbeitung mit konfigurierbarer Anzahl von Arbeitsthreads.
* Optionale Performance-Metriken nach Abschluss.

---

## Build

```bash
go build -o json2any ./cmd/app/main.go
```

---

## Installation (für globalen Zugriff)

```bash
go install github.com/KaoriEl/json2any@latest
```

---

## Nutzung

### Export von JSON in andere Formate

Konvertieren von JSON-Daten in die Formate `.xlsx`, `.csv` oder `.txt` mit anpassbaren Parametern.

#### Beispiel: Export nach XLSX

```bash
json2any export -i example.json -o result.xlsx --format=xlsx --theme=green --max_workers=100 --show_metrics=true
```

#### Beispiel: Export nach CSV

```bash
json2any export -i example.json -o result.csv --format=csv --max_workers=10
```

#### Beispiel: Export nach TXT

```bash
json2any export -i example.json -o result.txt --format=txt --max_workers=5
```

### Import von anderen Formaten nach JSON

Konvertieren von `.xlsx`, `.csv` oder `.txt` Dateien in das JSON-Format.

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

## Hilfe

```bash
json2any --help
```

---

## CLI-Flags

| Flag             | Beschreibung                                                                                                                                                                |
| ---------------- | --------------------------------------------------------------------------------------------------------------------------------------------------------------------------- |
| `--input, -i`    | **(Erforderlich)** Pfad zur Eingabedatei (JSON für den Export, XLSX/CSV/TXT für den Import).                                                                                |
| `--output, -o`   | Pfad zur Ausgabedatei. Standardwert: `random.xlsx` (für den Export) oder `output.json` (für den Import).                                                                    |
| `--format`       | Ausgabeformat für den Export: `xlsx`, `csv` oder `txt`. Eingabeformat für den Import: `xlsx`, `csv` oder `txt`. Standardwert: `xlsx` für den Export, `xlsx` für den Import. |
| `--theme`        | Tabellenthema: `black`, `green`, `red`, `purple`, `none`. Standardwert: `black`. (Nur für den Export)                                                                       |
| `--max_workers`  | Anzahl der parallelen Arbeitsthreads. Ganzzahl > 0. Standardwert: `20`.                                                                                                     |
| `--show_metrics` | Zeigt Leistungsmetriken nach Abschluss der Verarbeitung an. Standardwert: `false`.                                                                                          |