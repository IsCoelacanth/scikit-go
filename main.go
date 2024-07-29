package main

import "fmt"

// Run Linear Regression
func simulateLinearRegression() {
	fmt.Println("Linear Regression")
	// user input for number of data points
	var n int
	fmt.Println("Enter the number of data points: ")
	fmt.Scanln(&n)
	var x = createRandVec(n)
	var y = createRandVec(n)
	for i := 0; i < n; i++ {
		y[i] = y[i] * y[i]
	}
	fmt.Println("x = ", x)
	fmt.Println("y = ", y)
	var slope, intercept = LinearRegression(n, x, y)
	fmt.Println("Slope: ", slope)
	fmt.Println("Intercept: ", intercept)
}

// Run Multivariate Linear regression
func simulateMultivariateLinearRegression() {
	fmt.Println("Multivariate Linear Regression")
	// user input for number of data points
	var n int
	var l int
	fmt.Println("Enter the number of variables: ")
	fmt.Scanln(&n)
	fmt.Println("Enter the length of vectors: ")
	fmt.Scanln(&l)

	// Create Ouput Vector
	var outputs []float64 = CreateRandomVector(l)
	for i := range outputs {
		outputs[i] = outputs[i] * outputs[i]
	}

	// Create Input Vectors
	inputs := make([][]float64, n+1)
	for i := range inputs {
		if i == 0 {
			inputs[i] = CreateOnesVector(l)
		}
		inputs[i] = CreateRandomVector(l)
	}

	fmt.Printf("Outputs: %v", outputs)
	fmt.Printf("Inputs: %v", inputs)
}

func main() {
	var n int
	fmt.Println("Enter 1 for Linear Regression and 2 for Multivariate Linear Regression: ")
	fmt.Scanln(&n)

	if n == 1 {
		simulateLinearRegression()
	} else {
		simulateMultivariateLinearRegression()
	}
}
