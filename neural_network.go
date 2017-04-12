package go_nn

import (
	"fmt"
	"math"
)

type NeuralNetwork struct {
	Inputs []*Input
	Layers []*Layer
	Outputs []float64
}

func NewNeuralNetwork(nbInputs int, nbHiddenLayers int, nbOutputs int) *NeuralNetwork {
	nn := &NeuralNetwork{}

	for i := 0; i < nbInputs; i++ {
		nn.Inputs = append(nn.Inputs, NewInput())
	}

	nbInputsInHiddenLayers := int(math.Ceil(float64(nbInputs + nbOutputs) / 2))

	for i := 0; i < nbHiddenLayers; i++ {
		nn.Layers = append(nn.Layers, NewLayer(nbInputsInHiddenLayers))
	}

	nn.init()

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

	return nn
}

func (nn *NeuralNetwork) init() {
	for i := len(nn.Layers) - 1; i > 0; i-- {
		LinkLayers(nn.Layers[i - 1], nn.Layers[i])
	}

	ConnectInputs(nn.Inputs, nn.Layers[0])
}
