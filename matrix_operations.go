package main

import (
	"math"
	"math/rand/v2"
)

func CreateRandomVector(vectorLength int) []float64 {
	vector := make([]float64, vectorLength)
	for i := range vector {
		vector[i] = float64(i) + (2*rand.Float64() - 1) // monotonic increasing for testing
	}

	return vector
}

func CreateOnesVector(vectorLength int) []float64 {
	vector := make([]float64, vectorLength)
	for i := range vector {
		vector[i] = 1
	}

	return vector
}

func CreateZeroesVector(vectorLength int) []float64 {
	vector := make([]float64, vectorLength)
	for i := range vector {
		vector[i] = 0
	}

	return vector
}

func CreateInputVectors(vectorNumber int, vectorLength int) [][]float64 {
	// Create Multivariate Vector
	vec := make([][]float64, vectorNumber)
	for i := range vec {
		vec[i] = CreateRandomVector(vectorLength)
	}

	return vec
}

func Mean(vector []float64) float64 {
	var sum float64 = 0
	for i := range vector {
		sum += vector[i]
	}

	return sum / float64(len(vector))
}

func StdDeviation(vector []float64) float64 {
	mean := Mean(vector)
	var sum float64 = 0

	for i := range vector {
		sum += math.Pow(vector[i]-mean, 2)
	}

	return sum / float64(len(vector))
}

func Normalize(vector []float64) []float64 {
	mean := Mean(vector)
	stdDeviation := StdDeviation(vector)

	for i := range vector {
		vector[i] = (vector[i] - mean) / stdDeviation

	}

	return vector
}

func Transpose(vectors [][]float64) [][]float64 {
	return [][]float64{}
}

func Inverse(vectors [][]float64) [][]float64 {
	return [][]float64{}
}

func Multiply(vectors1 [][]float64, vectors2 [][]float64) [][]float64 {
	return [][]float64{}
}
