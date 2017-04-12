package go_nn

import "fmt"

type Neuron struct {
	LinksIn []*Link
	LinksOut []*Link
	Out float64
	ActivationFunction ActivationFunction
}

func NewNeuron(activationFunction ActivationFunction) *Neuron {
	n := &Neuron{}
	n.ActivationFunction = activationFunction

	return n
}

func (n *Neuron) Process() {
	var sum float64

	for _, l := range n.LinksIn {
		sum += (l.InValue * l.Weight)
	}

	n.Out = n.ActivationFunction(sum)

	fmt.Printf("Input %v produces output %v \n", sum, n.Out)

	for _, l := range n.LinksOut {
		l.Trigger(n.Out)
	}
}
