<p align="center">
  <img src="img/logo.jpg" alt="JSON to Any Exporter Logo" width="250"/>
</p>

# 🔧 JSON to Any Exporter

📘 Документация доступна на других языках:

* 🇬🇧 [English](README.md)
* 🇨🇳 [简体中文](README.zh.md)
* 🇪🇸 [Español](README.es.md)
* 🇩🇪 [Deutsch](README.de.md)

Это **CLI-инструмент** для конвертации между **JSON** и форматами **Excel (.xlsx), CSV и TXT** с поддержкой **темизации**, **форматирования типов данных**, **параллельной обработки** и **двунаправленной конвертации**.

---

## ✨ Возможности

* 🚀 **Экспорт**: преобразование JSON файлов в форматы `.xlsx`, `.csv` и `.txt`
* 🔄 **Импорт**: преобразование файлов `.xlsx`, `.csv` и `.txt` в формат JSON
* 🎨 Поддержка тем: `black`, `green`, `red`, `purple`, `none`
* 🔢 Корректное форматирование **чисел**, **дат**, **строк** и **логических значений**
* ⚙️ Параллельная обработка с настраиваемым количеством рабочих потоков
* 📊 Опциональный вывод метрик производительности после завершения

---

## 🛠️ Сборка

```bash
go build -o json2any ./cmd/app/main.go
```

---

## 🚀 Установка (для системного доступа)

```bash
go install github.com/KaoriEl/json2any/v2@latest
```

---

## 📋 Использование

### Экспорт JSON в другие форматы

![example.png](img/example_xlsx.png)

Конвертируйте JSON данные в форматы `.xlsx`, `.csv` или `.txt` с возможностью настройки.

#### Пример: экспорт в XLSX

```bash
json2any export -i example.json -o result.xlsx --format=xlsx --theme=green --max_workers=100 --show_metrics=true
```

#### Пример: экспорт в CSV

```bash
json2any export -i example.json -o result.csv --format=csv --max_workers=10
```

#### Пример: экспорт в TXT

```bash
json2any export -i example.json -o result.txt --format=txt --max_workers=5
```

---

### Импорт из других форматов в JSON

![example\_import\_xlsx.png](img/example_import_txt.png)

Конвертируйте файлы `.xlsx`, `.csv` или `.txt` в формат JSON.

#### Пример: импорт из XLSX в JSON

```bash
json2any import -i example.xlsx -o result.json --format=xlsx --max_workers=10
```

#### Пример: импорт из CSV в JSON

```bash
json2any import -i example.csv -o result.json --format=csv --max_workers=10
```

#### Пример: импорт из TXT в JSON

```bash
json2any import -i example.txt -o result.json --format=txt --max_workers=10
```

---

## ❓ Помощь

```bash
json2any --help
```

---

## ⚙️ Параметры CLI

| Флаг             | Описание                                                                                                                                               |
| ---------------- | ------------------------------------------------------------------------------------------------------------------------------------------------------ |
| `--input, -i`    | **(Обязательно)** Путь к входному файлу (JSON для экспорта, XLSX/CSV/TXT для импорта).                                                                 |
| `--output, -o`   | Путь к выходному файлу. По умолчанию: `random.xlsx` (для экспорта) или `output.json` (для импорта).                                                    |
| `--format`       | Формат вывода для экспорта: `xlsx`, `csv`, или `txt`. Формат ввода для импорта: `xlsx`, `csv`, или `txt`. По умолчанию: `xlsx` для экспорта и импорта. |
| `--theme`        | Тема таблицы: `black`, `green`, `red`, `purple`, `none`. По умолчанию: `black`. (Только для экспорта)                                                  |
| `--max_workers`  | Количество параллельных рабочих потоков. Целое число > 0. По умолчанию: `20`.                                                                          |
| `--show_metrics` | Показывать метрики производительности после завершения. По умолчанию: `false`.                                                                         |