package main

import "testing"

func TestMaxCalory(t *testing.T) {

	data := `1000
2000
3000

4000

5000
6000

7000
8000
9000

10000`

	max := MaxCalory(data)
	expected := []int{24000, 11000, 10000}

	for i := range expected {
		if max[i] != expected[i] {
			t.Errorf("got %v, expected %v", max[i], expected[i])
		}
	}
}
