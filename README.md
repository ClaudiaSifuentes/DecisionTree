# DecisionTree

# Árbol de Decisión en Go – Práctica Concurrente

Este proyecto implementa un algoritmo de Árbol de Decisión desde cero utilizando el lenguaje Go. Incluye dos versiones: una secuencial y una concurrente, con el objetivo de comparar su rendimiento al trabajar con grandes volúmenes de datos.

## 📁 Estructura del repositorio


```
dataset/
└── diabetes.csv           # Dataset original (768 registros)
generar_dataset.go         # Script para generar un dataset sintético grande
main_secuencial.go         # Implementación secuencial del árbol
main_concurrente.go        # Implementación concurrente con goroutines
.gitignore
```

---

## ⚙️ Requisitos

- Go 1.18 o superior
- Archivo `diabetes.csv` dentro de la carpeta `dataset`

---

## 📊 Dataset

Se utiliza el dataset Pima Indians Diabetes Dataset, que contiene atributos médicos relacionados con la diabetes.  
Se generó un nuevo archivo `diabetes_small.csv` con 1,000,000 de registros sintéticos a partir del original usando solo dos atributos: Glucosa y Edad.

---

## 🚀 Ejecución

### 1. Generar dataset grande

```bash
go run generar_dataset.go
````

Esto generará un archivo `diabetes_small.csv` con 1 millón de registros.

### 2. Ejecutar versión secuencial

```bash
go run main_secuencial.go
```

### 3. Ejecutar versión concurrente

```bash
go run main_concurrente.go
```

---

## 🧪 Comparación de rendimiento

| Dataset         | Registros  | Profundidad | Versión     | Tiempo estimado    |
| --------------- | ---------  | ----------- | ----------- | ------------------ |
| diabetes\_small | 20,000     | 2           | Secuencial  | 1 min 30 s         |
| diabetes\_small | 100,000    | 2           | Concurrente | 44m8.657083433s    |

---

## 📌 Conclusiones

* El uso de goroutines mejora significativamente el tiempo de entrenamiento en datasets grandes.
* Limitar a dos features y profundidad 2 permite mantener tiempos de ejecución manejables sin perder funcionalidad del árbol.
* Esta práctica demuestra el valor de la programación concurrente en algoritmos intensivos.

---

## 📎 Autor

Claudia Sifuentes
Práctica Calificada 2 – Programación Concurrente y Distribuida
UPC – 2025-1



