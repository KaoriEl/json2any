# JSON to XLSX 导出工具

📘 文档支持多语言：
- [English](README.md)
- [Русский](README.ru.md)
- [Español](README.es.md)
- [Deutsch](README.de.md)


一个命令行工具（CLI），用于将 JSON 文件转换为 Excel (.xlsx) 文件，支持主题样式、数据类型格式化和并发处理。

---

## 功能特点

* 将 JSON 对象数组转换为 `.xlsx` 表格。
* 支持的主题样式：`black`、`green`、`red`、`purple`、`none`。
* 正确格式化数字、日期、字符串和布尔值。
* 支持并发处理，可自定义工作线程数。
* 可选的性能指标输出。

---

## 构建

```bash
go build -o json2xlsx ./cmd/app/main.go
```

---

## 安装（系统全局可用）

```bash
sudo cp json2xlsx /usr/local/bin/
```

---

## 使用方式

### 使用 `go run`（无需构建）：

```bash
go run ./cmd/app/main.go -i example.json -o result.xlsx --theme=green --max_workers=100 --show_metrics=true
```

### 使用已构建的可执行文件（当前目录）：

```bash
./json2xlsx -i example.json -o result.xlsx --theme=green --max_workers=10
```

### 系统安装后（任意目录）：

```bash
json2xlsx -i example.json -o result.xlsx --theme=green --max_workers=10
```

---

## 查看帮助

```bash
json2xlsx --help
```

---

## 命令行参数

| 参数               | 描述                                                       |
| ---------------- | -------------------------------------------------------- |
| `--input, -i`    | **（必填）** 输入的 JSON 文件路径。                                  |
| `--output, -o`   | 输出的 XLSX 文件路径。默认值：`random.xlsx`。                         |
| `--theme`        | 表格主题样式：`black`、`green`、`red`、`purple`、`none`。默认：`black`。 |
| `--max_workers`  | 并发工作线程数。必须为大于 0 的整数。默认值：`20`。                            |
| `--show_metrics` | 显示处理完成后的性能指标。默认值：`false`。                                |