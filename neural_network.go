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
	nn.print()

	nn.randomise()

	return nn
}

func (nn *NeuralNetwork) print() {
	fmt.Printf("Number of layers: %d \n", len(nn.Layers))
	nbLinksInputs := 0

	for _, i := range nn.Inputs {
		nbLinksInputs += len(i.LinksOut)
	}

	fmt.Printf("Number of input connections: %d \n", nbLinksInputs)

	for i, l := range nn.Layers {
		fmt.Printf("Layer %d has %d neurons \n", i + 1, len(l.Neurons))
		nbLinks := 0
		for _, n := range l.Neurons {
			nbLinks += len(n.LinksOut)
		}
		fmt.Printf("Nb links %d \n", nbLinks)
	}
}

func (nn *NeuralNetwork) initInputs(nb int) {
	for i := 0; i < nb; i++ {
		nn.Inputs = append(nn.Inputs, NewInput())
	}
}

func (nn *NeuralNetwork) initHiddenLayers(hiddenLayers []int) {
	for _, nbInputs := range hiddenLayers {
		nn.Layers = append(nn.Layers, NewLayer(nbInputs, sigmoid()))
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

	for i, n := range nn.OutputLayer.Neurons {
		fmt.Printf("Output %v = %v \n", i, n.Out)
	}
}
