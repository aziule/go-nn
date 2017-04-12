package go_nn

type Layer struct {
	Neurons []*Neuron
}

func NewLayer(nbNeurons int) *Layer {
	layer := &Layer{}

	for i := 0; i < nbNeurons; i++ {
		n :=  NewNeuron()
		layer.Neurons = append(layer.Neurons, n)
	}

	return layer
}

func LinkLayers(layerIn, layerOut *Layer) {
	for _, nIn := range layerIn.Neurons {
		for _, nOut := range layerOut.Neurons {
			LinkNeurons(nIn, nOut)
		}
	}
}
