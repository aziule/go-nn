package go_nn

import (
	"fmt"
	"math/rand"
)

type Neuron struct {
	LinksIn []*Link
	LinksOut []*Link
	Bias float64
	Out float64
	ActivationFunction ActivationFunction
}

func NewNeuron(activationFunction ActivationFunction) *Neuron {
	n := &Neuron{}
	n.ActivationFunction = activationFunction

	return n
}

func (n *Neuron) RandomiseBias() {
	n.Bias = rand.Float64()
	fmt.Printf("Bias: %v \n", n.Bias)
}

func (n *Neuron) Activate() {
	var sum float64

	for _, l := range n.LinksIn {
		sum += (l.InValue * l.Weight)
	}

	sum += n.Bias

	n.Out = n.ActivationFunction(sum)

	fmt.Printf("Input %v produces output %v \n", sum, n.Out)

	n.feedForward()
}

func (n *Neuron) feedForward() {
	for _, l := range n.LinksOut {
		l.Trigger(n.Out)
	}
}
