package go_nn

import "fmt"

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
}
