# JSON 转换器

📘 文档可以在其他语言中查看：

* [English](README.md)
* [Русский](README.ru.md)
* [Español](README.es.md)
* [Deutsch](README.de.md)

命令行工具，用于在 JSON 和 Excel (.xlsx)、CSV 和 TXT 格式之间进行转换，支持主题、数据类型格式化、并行处理和双向转换。

---

## 特性

* **导出**：将 JSON 文件转换为 `.xlsx`、`.csv` 和 `.txt` 格式。
* **导入**：将 `.xlsx`、`.csv` 和 `.txt` 文件转换为 JSON 格式。
* 支持主题：`black`、`green`、`red`、`purple`、`none`。
* 正确格式化数字、日期、字符串和布尔值。
* 支持并行处理，并可配置线程数。
* 完成后可选显示性能指标。

---

## 编译

```bash
go build -o json2any ./cmd/app/main.go
```

---

## 安装（全局访问）

```bash
go install github.com/KaoriEl/json2any/v2@latest
```

---

## 使用

### 将 JSON 导出为其他格式

将 JSON 数据转换为 `.xlsx`、`.csv` 或 `.txt` 格式，支持自定义参数。

#### 示例：导出为 XLSX

```bash
json2any export -i example.json -o result.xlsx --format=xlsx --theme=green --max_workers=100 --show_metrics=true
```

#### 示例：导出为 CSV

```bash
json2any export -i example.json -o result.csv --format=csv --max_workers=10
```

#### 示例：导出为 TXT

```bash
json2any export -i example.json -o result.txt --format=txt --max_workers=5
```

### 从其他格式导入到 JSON

将 `.xlsx`、`.csv` 或 `.txt` 文件转换为 JSON 格式。

#### 示例：从 XLSX 导入到 JSON

```bash
json2any import -i example.xlsx -o result.json --format=xlsx --max_workers=10
```

#### 示例：从 CSV 导入到 JSON

```bash
json2any import -i example.csv -o result.json --format=csv --max_workers=10
```

#### 示例：从 TXT 导入到 JSON

```bash
json2any import -i example.txt -o result.json --format=txt --max_workers=10
```

---

## 帮助

```bash
json2any --help
```

---

## CLI 标志

| 标志               | 描述                                                                               |
| ---------------- | -------------------------------------------------------------------------------- |
| `--input, -i`    | **（必需）** 输入文件的路径（导出时为 JSON，导入时为 XLSX/CSV/TXT）。                                   |
| `--output, -o`   | 输出文件的路径。默认值：`random.xlsx`（导出时）或 `output.json`（导入时）。                              |
| `--format`       | 导出格式：`xlsx`、`csv` 或 `txt`。导入格式：`xlsx`、`csv` 或 `txt`。默认值：导出时为 `xlsx`，导入时为 `xlsx`。 |
| `--theme`        | 表格主题：`black`、`green`、`red`、`purple`、`none`。默认值：`black`。（仅导出时有效）                  |
| `--max_workers`  | 并行工作线程数。必须为大于 0 的整数。默认值：`20`。                                                    |
| `--show_metrics` | 完成后显示处理性能指标。默认值：`false`。                                                         |
