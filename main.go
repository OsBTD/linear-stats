package main

import (
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

func main() {
	sl := Read()
	a, b := LinearRegressionLine(sl)
	fmt.Printf("%.6f %.6f\n", a, b)
	fmt.Printf("%.10f\n", PearsonCorrelationCoefficient(sl))
}

func Read() []float64 {
	file := "data.txt"
	data, err := os.ReadFile(file)
	if err != nil {
		fmt.Println("couldn't read", err)
	}

	lines := strings.Split(string(data), "\n")

	var sl []float64
	for _, value := range lines {

		num, err := strconv.ParseFloat(value, 64)
		if err != nil {
			fmt.Println("couldn't convert", err)
		}
		sl = append(sl, (num))
	}

	return sl
}

func Sum(sl []float64) (float64, float64, float64, float64, float64) {
	var sumX, sumY, sumXY, sumXX, sumYY float64
	for index, value := range sl {
		sumX += float64(index)
		sumY += value
		sumXY += float64(index) * value
		sumXX += float64(index) * float64(index)
		sumYY += value * value
	}
	return sumX, sumY, sumXY, sumXX, sumYY
}

func LinearRegressionLine(sl []float64) (float64, float64) {
	sumX, sumY, sumXY, sumXX, _ := Sum(sl)
	n := float64(len(sl))
	a := (n*sumXY - sumX*sumY) / (n*sumXX - sumX*sumX)
	b := (sumY - a*sumX) / n
	// fmt.Println("nis", n)
	// fmt.Println("ais", n*sumXX-sumX*sumX)
	// fmt.Println("bis", b)

	return a, b
}

func PearsonCorrelationCoefficient(sl []float64) float64 {
	sumX, sumY, sumXY, sumXX, sumYY := Sum(sl)
	n := float64(len(sl))

	result := ((n * sumXY) - (sumX * sumY)) / math.Sqrt(((n*sumXX)-(sumX*sumX))*(n*sumYY-sumY*sumY))
	
	return result
}
