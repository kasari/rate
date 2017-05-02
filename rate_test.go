package rate

import "testing"

func TestAdd(t *testing.T) {
	r := NewRate()

	if len(r.RateValues) != 0 {
		t.Fatal("default RateValues length is 0")
	}

	r.Add(1.0, "1.0")

	if len(r.RateValues) != 1 {
		t.Fatal("Add add rate value to RateValues")
	}

}

func TestGenerate(t *testing.T) {
	r := NewRate()

	r.Add(30.0, "30")
	r.Add(70.0, "70")

	type Test struct {
		Rand  float64
		Value interface{}
	}
	tests := []Test{
		Test{
			Rand:  0.0,
			Value: "30",
		},
		Test{
			Rand:  0.29999,
			Value: "30",
		},
		Test{
			Rand:  0.3,
			Value: "70",
		},
		Test{
			Rand:  0.99999,
			Value: "70",
		},
	}

	for _, test := range tests {
		r.RandFunc = func() float64 {
			return test.Rand
		}

		value := r.Generate()
		if value != test.Value {
			t.Errorf("value %+v, want %+v", value, test.Value)
		}
	}
}
