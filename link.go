package go_nn

import (
	"math/rand"
	"fmt"
)

type Link struct {
	InValue float64
	OutValue float64
	Weight float64
}

func (l *Link) Trigger(inValue float64) {
	l.InValue = inValue
	l.OutValue = l.InValue * l.Weight
}

func (l *Link) RandomiseWeight() {
	l.Weight = 2 * rand.Float64()
	fmt.Printf("Weight: %v \n", l.Weight)
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
