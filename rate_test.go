package rate

import "testing"

func TestAdd(t *testing.T) {
	r := NewRate()
	r.MaxRate = 100

	if err := r.Add(30.0, "30.0"); err != nil {
		t.Error("if not exceed MaxRate, not return error")
	}

	if err := r.Add(80.0, "80.0"); err == nil {
		t.Error("if exceed MaxRate, return error")
	}
}

func TestGenerate(t *testing.T) {
	r := NewRate()
	r.DefaultValue = "default"
	r.MaxRate = 100
	r.RandFunc = func() float64 {
		return 1.0
	}

	r.Add(30.0, "30.0")

	if r.Generate() != "default" {
		t.Error("if not select set value, Generate returns DefaultValue")
	}
}
