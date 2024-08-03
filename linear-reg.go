package main

import (
	"math/rand/v2"
)

func square(x float64) float64 {
	return x * x
}

func mean(x []float64) float64 {
	var sum float64 = 0
	for i := 0; i < len(x); i++ {
		sum += x[i]
	}
	return sum / float64(len(x))
}

func createRandVec(size int) []float64 {
	var vec = make([]float64, size)
	for i := 0; i < size; i++ {
		vec[i] = float64(i) + (2*rand.Float64() - 1) // monotonic increasing for testing
	}
	return vec
}

func createZeroesVec(size int) []float64 {
	var vec = make([]float64, size)
	return vec
}

func LinearRegression(n int, x []float64, y []float64) (float64, float64) {
	var xMean float64 = mean(x)
	var yMean float64 = mean(y)

	var theta1_num float64 = 0.0
	var theta1_den float64 = 0.0

	for i := 0; i < n; i++ {
		theta1_num += (x[i] - xMean) * (y[i] - yMean)
	}
	for i := 0; i < n; i++ {
		theta1_den += square(x[i] - xMean)
	}

	var theta1 float64 = theta1_num / theta1_den
	var theta0 float64 = yMean - theta1*xMean

	return theta1, theta0
}
