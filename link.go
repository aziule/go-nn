package go_nn

type Link struct {
	InValue float64
	OutValue float64
	Weight float64
}

func ConnectInput(input *Input, neuron *Neuron) {
	l := &Link{}

	input.LinksOut = append(input.LinksOut, l)
	neuron.LinksIn = append(neuron.LinksIn, l)
}

func LinkNeurons(neuronIn, neuronOut *Neuron) {
	l := &Link{}

	neuronIn.LinksOut = append(neuronIn.LinksOut, l)
	neuronOut.LinksIn = append(neuronIn.LinksIn, l)
}
