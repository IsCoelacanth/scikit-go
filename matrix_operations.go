package main

import (
	"fmt"
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

func CreateIdentityMatrix(size int) [][]float64 {
	identity := make([][]float64, size)
	for i := range identity {
		identity[i] = make([]float64, size)
		identity[i][i] = 1
	}

	return identity
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

func Transpose(matrix [][]float64) [][]float64 {
	transpose := make([][]float64, len(matrix[0]))
	for i := range transpose {
		transpose[i] = make([]float64, len(matrix))
	}

	for i := range matrix {
		for j := range matrix[i] {
			transpose[j][i] = matrix[i][j]
		}
	}

	return transpose
}

func Inverse(matrix [][]float64) ([][]float64, error) {
	n := len(matrix)
	identity := CreateIdentityMatrix(n)

	// Augment the original matrix with the identity matrix
	augmented := make([][]float64, n)
	for i := range matrix {
		augmented[i] = append(matrix[i], identity[i]...)
	}

	// Perform Gauss-Jordan elimination
	for i := range augmented {
		// Make the diagonal contain all 1's
		if augmented[i][i] == 0 {
			// Swap rows if the diagonal element is zero
			for j := i + 1; j < n; j++ {
				if augmented[j][i] != 0 {
					augmented[i], augmented[j] = augmented[j], augmented[i]
					break
				}
			}
			if augmented[i][i] == 0 {
				return nil, fmt.Errorf("matrix is singular and cannot be inverted")
			}
		}

		diag := augmented[i][i]
		for j := range augmented[i] {
			augmented[i][j] /= diag
		}

		// Make the other elements in the column 0
		for j := range augmented {
			if i != j {
				factor := augmented[j][i]
				for k := range augmented[j] {
					augmented[j][k] -= factor * augmented[i][k]
				}
			}
		}
	}

	// Extract the inverse matrix from the augmented matrix
	inverseMatrix := make([][]float64, n)
	for i := range inverseMatrix {
		inverseMatrix[i] = augmented[i][n:]
	}

	return inverseMatrix, nil
}

func Multiply(matrix1 [][]float64, matrix2 [][]float64) ([][]float64, error) {
	n := len(matrix1)
	m := len(matrix1[0])
	p := len(matrix2[0])

	// Check if multiplication is possible
	if len(matrix2) != m {
		return nil, fmt.Errorf("number of columns in A must equal number of rows in B")
	}

	// Create the result matrix
	C := make([][]float64, n)
	for i := range C {
		C[i] = make([]float64, p)
	}

	// Perform matrix multiplication
	for i := 0; i < n; i++ {
		for j := 0; j < p; j++ {
			for k := 0; k < m; k++ {
				C[i][j] += matrix1[i][k] * matrix2[k][j]
			}
		}
	}

	return C, nil
}
