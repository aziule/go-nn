package go_nn

type Input struct {
	Value float64
	LinksOut []*Link
}

func NewInput() *Input {
	return &Input{}
}

func (i *Input) Send() {
	for _, l := range i.LinksOut {
		l.Trigger(i.Value)
	}
}

func ConnectInputs(inputs []*Input, layer *Layer) {
	for _, i := range inputs {
		for _, n := range layer.Neurons {
			ConnectInput(i, n)
		}
	}
}
