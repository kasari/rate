package main

import (
	"fmt"
	"os"

	"github.com/kasari/rate"
)

const TryCount = 10000

func main() {
	data := []rate.RateValue{
		rate.RateValue{10, "rate-10"},
		rate.RateValue{30, "rate-30"},
		rate.RateValue{60, "rate-60"},
	}

	r := rate.NewRate()
	r.MaxRate = 200
	r.DefaultValue = "default"
	for _, d := range data {
		err := r.Add(d.Rate, d.Value)
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
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
