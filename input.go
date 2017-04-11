package go_nn

type Input struct {
	Value float64
	LinksOut []*Link
}

func NewInput() *Input {
	return &Input{}
}

func ConnectInputs(inputs []*Input, layer *Layer) {
	for _, i := range inputs {
		for _, n := range layer.Neurons {
			ConnectInput(i, n)
		}
	}
}
