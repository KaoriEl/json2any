# JSON to XLSX Exporter

📘 Documentación disponible en otros idiomas:
- [English](README.md)
- [Русский](README.ru.md)
- [简体中文](README.zh.md)
- [Deutsch](README.de.md)


Herramienta de línea de comandos (CLI) para convertir archivos JSON en Excel (.xlsx) con soporte para temas visuales, tipado de datos y procesamiento paralelo.

---

## Funcionalidades

* Conversión de arreglos de objetos JSON a hojas de cálculo `.xlsx`.
* Soporte para temas de formato: `black`, `green`, `red`, `purple`, `none`.
* Formateo correcto de números, fechas, cadenas y valores booleanos.
* Procesamiento paralelo con opción para definir la cantidad de trabajadores.
* Salida opcional de métricas de rendimiento tras la ejecución.

---

## Compilación

```bash
go build -o json2xlsx ./cmd/app/main.go
```

---

## Instalación (para acceso global en el sistema)

```bash
sudo cp json2xlsx /usr/local/bin/
```

---

## Uso

### Con `go run` (sin compilar):

```bash
go run ./cmd/app/main.go -i example.json -o result.xlsx --theme=green --max_workers=100 --show_metrics=true
```

### Con binario compilado desde el directorio actual:

```bash
./json2xlsx -i example.json -o result.xlsx --theme=green --max_workers=10
```

### Desde cualquier ubicación (si está instalado globalmente):

```bash
json2xlsx -i example.json -o result.xlsx --theme=green --max_workers=10
```

---

## Ayuda

```bash
json2xlsx --help
```

---

## Parámetros CLI

| Parámetro        | Descripción                                                                                 |
| ---------------- | ------------------------------------------------------------------------------------------- |
| `--input, -i`    | **(Obligatorio)** Ruta al archivo JSON de entrada.                                          |
| `--output, -o`   | Ruta al archivo de salida XLSX. Por defecto: `random.xlsx`.                                 |
| `--theme`        | Tema visual para la tabla: `black`, `green`, `red`, `purple`, `none`. Por defecto: `black`. |
| `--max_workers`  | Número de trabajadores paralelos. Entero > 0. Por defecto: `20`.                            |
| `--show_metrics` | Mostrar métricas de rendimiento al finalizar. Por defecto: `false`.                         |
