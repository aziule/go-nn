package go_nn

import "math"

type ActivationFunction func(float64) float64

func sigmoid() func(float64) float64 {
	return func(z float64) float64 {
		return (1 / 1 + math.Exp(-z))
	}
}
