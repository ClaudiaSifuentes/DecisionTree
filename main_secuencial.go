package main

import (
	"encoding/csv"
	"fmt"
	"os"
	"strconv"
	"time"
)

// estrucutura datapoint
type DataPoint struct {
	Features []float64
	Label    int
}

// leer csv
func loadCSV(filename string) ([]DataPoint, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	reader := csv.NewReader(file)

	// Ignora la cabecera
	_, err = reader.Read()
	if err != nil {
		return nil, err
	}

	lines, err := reader.ReadAll()
	if err != nil {
		return nil, err
	}

	var data []DataPoint
	for _, line := range lines {
		var features []float64
		for i := 0; i < len(line)-1; i++ {
			val, err := strconv.ParseFloat(line[i], 64)
			if err != nil {
				return nil, err
			}
			features = append(features, val)
		}
		label, err := strconv.Atoi(line[len(line)-1])
		if err != nil {
			return nil, err
		}
		data = append(data, DataPoint{Features: features, Label: label})
	}

	return data, nil
}

// Nodo del arbol
type Node struct {
	FeatureIndex int
	Threshold    float64
	Left         *Node
	Right        *Node
	IsLeaf       bool
	Prediction   int
}

// Gini
func gini(data []DataPoint) float64 {
	count := make(map[int]int)
	for _, d := range data {
		count[d.Label]++
	}
	total := float64(len(data))
	impurity := 1.0
	for _, c := range count {
		prob := float64(c) / total
		impurity -= prob * prob
	}
	return impurity
}

// division de datos
func split(data []DataPoint, featureIndex int, threshold float64) ([]DataPoint, []DataPoint) {
	var left, right []DataPoint
	for _, d := range data {
		if d.Features[featureIndex] <= threshold {
			left = append(left, d)
		} else {
			right = append(right, d)
		}
	}
	return left, right
}

// mejor division
func bestSplit(data []DataPoint) (int, float64) {
	bestGini := 1.0
	bestFeature := 0
	bestThreshold := 0.0

	for featureIndex := range data[0].Features {
		for _, d := range data {
			threshold := d.Features[featureIndex]
			left, right := split(data, featureIndex, threshold)
			if len(left) == 0 || len(right) == 0 {
				continue
			}

			giniLeft := gini(left)
			giniRight := gini(right)
			weighted := (float64(len(left))*giniLeft + float64(len(right))*giniRight) / float64(len(data))

			if weighted < bestGini {
				bestGini = weighted
				bestFeature = featureIndex
				bestThreshold = threshold
			}
		}
	}

	return bestFeature, bestThreshold
}

// Clase mayoritaria
func majorityClass(data []DataPoint) int {
	count := make(map[int]int)
	for _, d := range data {
		count[d.Label]++
	}
	maxCount := -1
	majority := -1
	for label, c := range count {
		if c > maxCount {
			majority = label
			maxCount = c
		}
	}
	return majority
}

// Construir el arbol recursivamente
func buildTree(data []DataPoint, depth, maxDepth int) *Node {
	if len(data) == 0 || depth >= maxDepth {
		return &Node{IsLeaf: true, Prediction: majorityClass(data)}
	}

	feature, threshold := bestSplit(data)
	left, right := split(data, feature, threshold)

	if len(left) == 0 || len(right) == 0 {
		return &Node{IsLeaf: true, Prediction: majorityClass(data)}
	}

	return &Node{
		FeatureIndex: feature,
		Threshold:    threshold,
		Left:         buildTree(left, depth+1, maxDepth),
		Right:        buildTree(right, depth+1, maxDepth),
	}
}

// predecir
func predict(node *Node, features []float64) int {
	if node.IsLeaf {
		return node.Prediction
	}
	if features[node.FeatureIndex] <= node.Threshold {
		return predict(node.Left, features)
	}
	return predict(node.Right, features)
}

func main() {
	data, err := loadCSV("dataset/diabetes.csv")

	if err != nil {
		fmt.Println("Error al cargar CSV:", err)
		return
	}

	fmt.Printf("Se cargaron %d registros\n", len(data))

	start := time.Now()
	tree := buildTree(data, 0, 5)
	elapsed := time.Since(start)

	fmt.Printf("Árbol construido en: %s\n", elapsed)

	test := []float64{140, 80, 33, 115, 35.6, 0.45, 45}
	pred := predict(tree, test)
	fmt.Printf("Predicción para el paciente: Clase %d\n", pred)
}
