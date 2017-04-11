package go_nn

type Neuron struct {
	LinksIn []*Link
	LinksOut []*Link
	Out float64
}

func NewNeuron() *Neuron {
	return &Neuron{}
}
