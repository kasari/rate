package rate

import (
	"math/rand"
)

type RateValue struct {
	Rate  float64
	Value interface{}
}

type Rate struct {
	MaxRate    float64
	RateValues []RateValue
	RandFunc   func() float64 // return number in [0.0,1.0)
}

func NewRate() *Rate {
	return &Rate{}
}

func (r *Rate) Add(rate float64, value interface{}) {
	r.MaxRate += rate
	r.RateValues = append(r.RateValues, RateValue{rate, value})
}

func (r *Rate) Generate() interface{} {
	if r.RandFunc == nil {
		r.RandFunc = rand.Float64
	}

	index := r.RandFunc() * r.MaxRate

	var cursor float64
	for _, rv := range r.RateValues {
		cursor += rv.Rate
		if index < cursor {
			return rv.Value
		}
	}

	return nil // unreachable code
}
