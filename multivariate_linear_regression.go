package main

import (
	"math/rand/v2"
)

func CreateRandomVector(variableLength int, vectorLength int) [][]float64 {
	// Create Multivariate Vector
	vec := make([][]float64, vectorLength)
	for i := range vec {
		vec[i] = make([]float64, variableLength)
		for j := 0; j < variableLength; j++ {
			vec[i][j] = float64(j) + (2*rand.Float64() - 1) // monotonic increasing for testing
		}

	}

	return vec
}
