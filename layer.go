package go_nn

type Layer struct {
	Neurons []*Neuron
}

func NewLayer(nbNeurons int, activationFunction ActivationFunction) *Layer {
	layer := &Layer{}

	for i := 0; i < nbNeurons; i++ {
		n :=  NewNeuron(activationFunction)
		layer.Neurons = append(layer.Neurons, n)
	}

	return layer
}

func (l *Layer) Process() {
	for _, n := range l.Neurons {
		n.Activate()
	}
}

func LinkLayers(layerIn, layerOut *Layer) {
	for _, nIn := range layerIn.Neurons {
		for _, nOut := range layerOut.Neurons {
			LinkNeurons(nIn, nOut)
		}
	}
}
