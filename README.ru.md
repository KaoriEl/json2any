# JSON to Any Exporter

📘 Документация доступна на других языках:
* [English](README.md)
* [简体中文](README.zh.md)
* [Español](README.es.md)
* [Deutsch](README.de.md)

Инструмент командной строки для конвертации между форматами JSON и Excel (.xlsx), CSV и TXT с поддержкой тем, форматирования типов данных, параллельной обработки и двусторонней конверсии.

---

## Особенности

* **Экспорт**: Конвертирование JSON-файлов в форматы `.xlsx`, `.csv` и `.txt`.
* **Импорт**: Конвертирование файлов `.xlsx`, `.csv` и `.txt` в формат JSON.
* Поддержка тем: `black`, `green`, `red`, `purple`, `none`.
* Правильное форматирование чисел, дат, строк и булевых значений.
* Параллельная обработка с настраиваемым количеством рабочих потоков.
* Опциональный вывод производственных метрик по завершении.

---

## Сборка

```bash
go build -o json2xlsx ./cmd/app/main.go
```

---

## Установка (для глобального доступа)

```bash
go install github.com/KaoriEl/json2xlsx@latest
```

---

## Использование

### Экспорт JSON в другие форматы

Конвертирование данных JSON в форматы `.xlsx`, `.csv` или `.txt` с настраиваемыми параметрами.

#### Пример: Экспорт в XLSX

```bash
json2xlsx export -i example.json -o result.xlsx --format=xlsx --theme=green --max_workers=100 --show_metrics=true
```

#### Пример: Экспорт в CSV

```bash
json2xlsx export -i example.json -o result.csv --format=csv --max_workers=10
```

#### Пример: Экспорт в TXT

```bash
json2xlsx export -i example.json -o result.txt --format=txt --max_workers=5
```

### Импорт из других форматов в JSON

Конвертирование файлов `.xlsx`, `.csv` или `.txt` в формат JSON.

#### Пример: Импорт из XLSX в JSON

```bash
json2xlsx import -i example.xlsx -o result.json --format=xlsx --max_workers=10
```

#### Пример: Импорт из CSV в JSON

```bash
json2xlsx import -i example.csv -o result.json --format=csv --max_workers=10
```

#### Пример: Импорт из TXT в JSON

```bash
json2xlsx import -i example.txt -o result.json --format=txt --max_workers=10
```

---

## Помощь

```bash
json2xlsx --help
```

---

## Флаги CLI

| Флаг             | Описание                                                                                                                                                                  |
| ---------------- | ------------------------------------------------------------------------------------------------------------------------------------------------------------------------- |
| `--input, -i`    | **(Обязательный)** Путь к входному файлу (JSON для экспорта, XLSX/CSV/TXT для импорта).                                                                                   |
| `--output, -o`   | Путь к выходному файлу. По умолчанию: `random.xlsx` (для экспорта) или `output.json` (для импорта).                                                                       |
| `--format`       | Формат вывода для экспорта: `xlsx`, `csv`, или `txt`. Формат входных данных для импорта: `xlsx`, `csv`, или `txt`. По умолчанию: `xlsx` для экспорта, `xlsx` для импорта. |
| `--theme`        | Тема таблицы: `black`, `green`, `red`, `purple`, `none`. По умолчанию: `black`. (Только для экспорта)                                                                     |
| `--max_workers`  | Количество параллельных рабочих потоков. Целое число > 0. По умолчанию: `20`.                                                                                             |
| `--show_metrics` | Показать метрики обработки по завершению. По умолчанию: `false`.                                                                                                          |
