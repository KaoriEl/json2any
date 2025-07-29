# JSON to XLSX Exporter

📘 Dokumentation in weiteren Sprachen verfügbar:
- [English](README.md)
- [Русский](README.ru.md)
- [简体中文](README.zh.md)
- [Español](README.es.md)


Ein CLI-Tool zur Konvertierung von JSON-Dateien in Excel-Dateien (.xlsx) mit Unterstützung für Designs, Datentypformatierung und parallele Verarbeitung.

---

## Funktionen

* Konvertierung von Arrays mit JSON-Objekten in `.xlsx`-Tabellen.
* Unterstützung für Tabellen-Designs: `black`, `green`, `red`, `purple`, `none`.
* Korrekte Formatierung von Zahlen, Datumsangaben, Zeichenketten und Booleschen Werten.
* Parallele Verarbeitung mit konfigurierbarer Anzahl von Workern.
* Optional: Anzeige von Leistungsmetriken nach Abschluss.

---

## Kompilierung

```bash
go build -o json2xlsx ./cmd/app/main.go
```

---

## Installation (systemweit verfügbar machen)

```bash
sudo cp json2xlsx /usr/local/bin/
```

---

## Verwendung

### Mit `go run` (ohne Kompilierung):

```bash
go run ./cmd/app/main.go -i example.json -o result.xlsx --theme=green --max_workers=100 --show_metrics=true
```

### Mit kompiliertem Binary im aktuellen Verzeichnis:

```bash
./json2xlsx -i example.json -o result.xlsx --theme=green --max_workers=10
```

### Von überall aus (wenn systemweit installiert):

```bash
json2xlsx -i example.json -o result.xlsx --theme=green --max_workers=10
```

---

## Hilfe anzeigen

```bash
json2xlsx --help
```

---

## CLI-Parameter

| Parameter        | Beschreibung                                                                   |
| ---------------- | ------------------------------------------------------------------------------ |
| `--input, -i`    | **(Erforderlich)** Pfad zur Eingabedatei im JSON-Format.                       |
| `--output, -o`   | Pfad zur Ausgabedatei im XLSX-Format. Standard: `random.xlsx`.                 |
| `--theme`        | Tabellen-Design: `black`, `green`, `red`, `purple`, `none`. Standard: `black`. |
| `--max_workers`  | Anzahl paralleler Worker. Ganze Zahl > 0. Standard: `20`.                      |
| `--show_metrics` | Zeigt nach Abschluss Leistungsmetriken an. Standard: `false`.                  |
