# Exportador de JSON a Cualquier Formato

📘 La documentación está disponible en otros idiomas:

* [Inglés](README.md)
* [简体中文](README.zh.md)
* [Русский](README.ru.md)
* [Deutsch](README.de.md)

Herramienta de línea de comandos para convertir entre formatos JSON y Excel (.xlsx), CSV y TXT, con soporte para temas, formateo de tipos de datos, procesamiento paralelo y conversión bidireccional.

---

## Características

* **Exportar**: Convertir archivos JSON a formatos `.xlsx`, `.csv` y `.txt`.
* **Importar**: Convertir archivos `.xlsx`, `.csv` y `.txt` a formato JSON.
* Soporte para temas: `black`, `green`, `red`, `purple`, `none`.
* Formateo correcto de números, fechas, cadenas y valores booleanos.
* Procesamiento paralelo con cantidad configurable de hilos de trabajo.
* Métricas opcionales de rendimiento al finalizar.

---

## Compilación

```bash
go build -o json2xlsx ./cmd/app/main.go
```

---

## Instalación (para acceso global)

```bash
go install github.com/KaoriEl/json2any@latest
```

---

## Uso

### Exportar JSON a otros formatos

Convertir datos JSON a formatos `.xlsx`, `.csv` o `.txt` con parámetros personalizables.

#### Ejemplo: Exportar a XLSX

```bash
json2xlsx export -i example.json -o result.xlsx --format=xlsx --theme=green --max_workers=100 --show_metrics=true
```

#### Ejemplo: Exportar a CSV

```bash
json2xlsx export -i example.json -o result.csv --format=csv --max_workers=10
```

#### Ejemplo: Exportar a TXT

```bash
json2xlsx export -i example.json -o result.txt --format=txt --max_workers=5
```

### Importar desde otros formatos a JSON

Convertir archivos `.xlsx`, `.csv` o `.txt` a formato JSON.

#### Ejemplo: Importar desde XLSX a JSON

```bash
json2xlsx import -i example.xlsx -o result.json --format=xlsx --max_workers=10
```

#### Ejemplo: Importar desde CSV a JSON

```bash
json2xlsx import -i example.csv -o result.json --format=csv --max_workers=10
```

#### Ejemplo: Importar desde TXT a JSON

```bash
json2xlsx import -i example.txt -o result.json --format=txt --max_workers=10
```

---

## Ayuda

```bash
json2xlsx --help
```

---

## Banderas de la CLI

| Bandera          | Descripción                                                                                                                                                               |
| ---------------- | ------------------------------------------------------------------------------------------------------------------------------------------------------------------------- |
| `--input, -i`    | **(Obligatorio)** Ruta del archivo de entrada (JSON para exportar, XLSX/CSV/TXT para importar).                                                                           |
| `--output, -o`   | Ruta del archivo de salida. Por defecto: `random.xlsx` (para exportar) o `output.json` (para importar).                                                                   |
| `--format`       | Formato de salida para exportar: `xlsx`, `csv` o `txt`. Formato de entrada para importar: `xlsx`, `csv` o `txt`. Por defecto: `xlsx` para exportar, `xlsx` para importar. |
| `--theme`        | Tema de la tabla: `black`, `green`, `red`, `purple`, `none`. Por defecto: `black`. (Solo para exportar)                                                                   |
| `--max_workers`  | Número de hilos de trabajo paralelos. Número entero > 0. Por defecto: `20`.                                                                                               |
| `--show_metrics` | Mostrar métricas de procesamiento al finalizar. Por defecto: `false`.                                                                                                     |
