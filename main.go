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

func MultivariateLinearRegression(inputs [][]float64, outputs [][]float64) ([][]float64, error) {
	// Calculate Inputs Transpose
	inputsTranspose := Transpose(inputs)

	// Calculate Product of Inputs and Inputs Transpose
	product, err := Multiply(inputsTranspose, inputs)

	if err != nil {
		fmt.Printf("Error in Calculating Product of Inputs and Inputs Transpose: %s", err.Error())
		return [][]float64{}, err
	}

	// Calculate Inverse of Product
	inverse, err := Inverse(product)

	if err != nil {
		fmt.Printf("Error in Calculating Inverse of Product: %s", err.Error())
		return [][]float64{}, err
	}

	// Calculate Product of Inputs Transpose and Outputs
	product_output, err := Multiply(inputsTranspose, outputs)

	if err != nil {
		fmt.Printf("Error in Calculating Product of Inputs Transpose and Output: %s", err.Error())
		return [][]float64{}, err
	}

	// Calculate Coefficients
	coefficients, err := Multiply(inverse, product_output)

	if err != nil {
		fmt.Printf("Error in Calculating Coefficients: %s", err.Error())
		return [][]float64{}, err
	}

	return coefficients, err

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
	outputs := make([][]float64, l)
	for i := range outputs {
		outputs[i] = CreateRandomVector(1)
	}

	// Create Input Vectors
	inputs := make([][]float64, l)
	for i := range inputs {
		inputs[i] = CreateRandomVector(n + 1)
		inputs[i][0] = 1
	}

	coefficients, err := MultivariateLinearRegression(inputs, outputs)

	if err != nil {
		return
	}

	fmt.Printf("Outputs: %v", outputs)
	fmt.Println()
	fmt.Printf("Inputs: %v", inputs)
	fmt.Println()
	fmt.Printf("Coefficients: %v", coefficients)
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
