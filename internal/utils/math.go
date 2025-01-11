package utils

import "math/rand"

type Choice[T any] struct {
	Weight  int
	Element T
}

func WeightedChoice[T any](choices ...Choice[T]) T {
	// XXX: is there a better way to do this
	// than iterating over the weights twice?
	var t int
	for _, c := range choices {
		t += c.Weight
	}

	r := rand.Intn(t)
	for _, c := range choices {
		if r <= c.Weight {
			return c.Element
		}
		r = r - c.Weight
	}

	return choices[len(choices)-1].Element
}
