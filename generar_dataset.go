package main

import (
	"encoding/csv"
	"fmt"
	"math/rand"
	"os"
	"strconv"
	"time"
)

func main() {
	original, err := os.Open("diabetes.csv")
	if err != nil {
		fmt.Println("Error al abrir el archivo original:", err)
		return
	}
	defer original.Close()

	reader := csv.NewReader(original)
	header, _ := reader.Read()
	lines, err := reader.ReadAll()
	if err != nil {
		fmt.Println("Error al leer CSV:", err)
		return
	}

	out, err := os.Create("diabetes_big.csv")
	if err != nil {
		fmt.Println("Error al crear archivo nuevo:", err)
		return
	}
	defer out.Close()

	writer := csv.NewWriter(out)
	writer.Write(header)

	rand.Seed(time.Now().UnixNano())

	target := 1000000
	count := 0

	for count < target {
		for _, row := range lines {
			if count >= target {
				break
			}
			var newRow []string
			for i := 0; i < len(row)-1; i++ {
				val, _ := strconv.ParseFloat(row[i], 64)
				val += rand.NormFloat64() * 2.0
				newRow = append(newRow, fmt.Sprintf("%.2f", val))
			}
			newRow = append(newRow, row[len(row)-1])
			writer.Write(newRow)
			count++
		}
	}

	writer.Flush()
	fmt.Println("Archivo 'diabetes_big.csv' generado con", count, "registros.")
}
