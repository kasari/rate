package rate

import (
	"fmt"
	"math/rand"
	"time"
)

type RateValue struct {
	Rate  float64
	Value interface{}
}

type Option struct {
	NoMaxRate bool
}

type Rate struct {
	Option
	MaxRate      float64
	RateValues   []RateValue
	DefaultValue interface{}
	RandFunc     func() float64 // return number in [0.0,1.0)
}

func NewRate() *Rate {
	return NewRateWithOption(Option{})
}

func NewRateWithOption(o Option) *Rate {
	return &Rate{Option: o}
}

func (r *Rate) Add(rate float64, value interface{}) error {
	if r.Option.NoMaxRate {
		r.MaxRate += rate
	}

	var totalRate float64
	for _, rv := range r.RateValues {
		totalRate += rv.Rate
	}

	if totalRate+rate > r.MaxRate {
		return fmt.Errorf("error exceed MaxRate, current:%f MaxRate:%f", totalRate+rate, r.MaxRate)
	}

	r.RateValues = append(r.RateValues, RateValue{rate, value})

	return nil
}

func (r *Rate) Generate() interface{} {
	if r.RandFunc == nil {
		rand.Seed(time.Now().UnixNano())
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

	return r.DefaultValue
}
