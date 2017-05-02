package main

import (
	"fmt"

	"github.com/kasari/rate"
)

const TryCount = 10000

func main() {
	data := []rate.RateValue{
		rate.RateValue{0.1, "rate-0.1"},
		rate.RateValue{0.3, "rate-0.3"},
		rate.RateValue{0.6, "rate-0.6"},
	}

	r := rate.NewRate()

	for _, d := range data {
		r.Add(d.Rate, d.Value)
	}

	countMap := make(map[interface{}]int)
	for i := 0; i < TryCount; i++ {
		v := r.Generate()
		countMap[v]++
	}

	fmt.Printf("[TryCount: %d]\n", TryCount)
	for k, v := range countMap {
		fmt.Printf("%v: %d\n", k, v)
	}
}
