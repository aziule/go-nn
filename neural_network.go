package go_nn

import (
	"fmt"
	"math"
	"time"
	"math/rand"
)

type NeuralNetwork struct {
	Inputs []*Input
	Layers []*Layer
	Outputs []float64
}

func NewNeuralNetwork(nbInputs int, nbHiddenLayers int, nbOutputs int) *NeuralNetwork {
	nn := &NeuralNetwork{}

	nbNeuronsInHiddenLayers := int(math.Ceil(float64(nbInputs + nbOutputs) / 2))

	nn.initInputs(nbInputs)
	nn.initHiddenLayers(nbHiddenLayers, nbNeuronsInHiddenLayers)
	nn.wire()
	nn.print()

	nn.randomiseWeights()

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

func (nn *NeuralNetwork) initHiddenLayers(nbLayers, nbNeurons int) {
	for i := 0; i < nbLayers; i++ {
		nn.Layers = append(nn.Layers, NewLayer(nbNeurons, sigmoid()))
	}
}

func (nn *NeuralNetwork) wire() {
	for i := len(nn.Layers) - 1; i > 0; i-- {
		LinkLayers(nn.Layers[i - 1], nn.Layers[i])
	}

	ConnectInputs(nn.Inputs, nn.Layers[0])
}

func (nn *NeuralNetwork) randomiseWeights() {
	rand.Seed(time.Now().UTC().UnixNano())

	for _, layer := range nn.Layers {
		for _, n := range layer.Neurons {
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

func (nn *NeuralNetwork) Process(inputs []float64) {
	nn.setupInputs(inputs)

	for _, i := range nn.Inputs {
		i.Send()
	}

	for _, l := range nn.Layers {
		l.Process()
	}
}
