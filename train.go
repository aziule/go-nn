package go_nn

import (
	"fmt"
	"math"
)

// We assume that the last value of set is the expected output
// Format: [[[input1.1, input1.2], [expected1]], [[input2.1, input2.2], [expected2]]], and so on...
func (nn *NeuralNetwork) Train(set [][][]float64) {
	for _, row := range set {
		nn.processInputs(row[0])
		nn.backPropagation(row[1])
	}
}

func (nn *NeuralNetwork) backPropagation(expected []float64) {
	fmt.Printf("Expected: %v \n", expected)
	var errorsSum float64

	for i, n := range nn.OutputLayer.Neurons {
		errorsSum += calculateError(expected[i], n.Out)
	}
}

func calculateError(expected, actual float64) float64 {
	return 0.5 * math.Pow(expected - actual, 2)
}
