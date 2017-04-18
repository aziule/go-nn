package go_nn

import (
	"fmt"
	"time"
	"math/rand"
)

type NeuralNetwork struct {
	Inputs []*Input
	Layers []*Layer
	OutputLayer *Layer
}

func NewNeuralNetwork(nbInputs int, hiddenLayers []int, nbOutputs int) *NeuralNetwork {
	nn := &NeuralNetwork{}

	nn.initInputs(nbInputs)
	nn.initHiddenLayers(hiddenLayers)
	nn.initOutputs(nbOutputs)
	nn.wire()

	nn.randomise()

	return nn
}

func (nn *NeuralNetwork) initInputs(nb int) {
	for i := 0; i < nb; i++ {
		nn.Inputs = append(nn.Inputs, NewInput())
	}
}

func (nn *NeuralNetwork) initHiddenLayers(hiddenLayers []int) {
	if len(hiddenLayers) == 0 {
		panic(fmt.Sprintf("You must specify at least one hidden layer"))
	}

	for _, nbNeurons := range hiddenLayers {
		if nbNeurons <= 0 {
			panic(fmt.Sprintf("Hidden layers must have at least 1 neuron (%v given)", nbNeurons))
		}

		nn.Layers = append(nn.Layers, NewLayer(nbNeurons, sigmoid()))
	}
}

func (nn *NeuralNetwork) initOutputs(nbOutputs int) {
	nn.OutputLayer = NewLayer(nbOutputs, sigmoid())
}

func (nn *NeuralNetwork) wire() {
	for i := len(nn.Layers) - 1; i > 0; i-- {
		LinkLayers(nn.Layers[i - 1], nn.Layers[i])
	}

	LinkLayers(nn.Layers[len(nn.Layers) - 1], nn.OutputLayer)

	ConnectInputs(nn.Inputs, nn.Layers[0])
}

func (nn *NeuralNetwork) randomise() {
	rand.Seed(time.Now().UTC().UnixNano())

	layersToInit := nn.Layers
	layersToInit = append(layersToInit, nn.OutputLayer)

	for _, layer := range layersToInit {
		for _, n := range layer.Neurons {
			n.RandomiseBias()

			for _, l := range n.LinksIn {
				l.RandomiseWeight()
			}
		}
	}
}

func (nn *NeuralNetwork) setupInputs(values []float64) {
	nbInputs := len(nn.Inputs)
	nbValues := len(values)

	if nbInputs != nbValues {
		panic(fmt.Sprintf("Invalid number of inputs: %v waited (%v given)", nbInputs, nbValues))
	}

	for index, input := range nn.Inputs {
		input.Value = values[index]
	}
}

func (nn *NeuralNetwork) processInputs(inputs []float64) {
	nn.setupInputs(inputs)

	for _, i := range nn.Inputs {
		i.Send()
	}

	for _, l := range nn.Layers {
		l.Process()
	}

	nn.OutputLayer.Process()
}
