package main

import (
	"testing"
	"time"
)

func Test_main(t *testing.T) {
	eatTime = 0 * time.Second

	for i := 0; i < 10; i++ {
		order = order
		main()
		if len(*order.list) != 5 {
			t.Errorf("incorrect length of slice; expected 5 but got %d", len(*order.list))
		}
	}
}

func Test_mainWithVaryingDelays(t *testing.T) {
	var theTests = []struct {
		name  string
		delay time.Duration
	}{
		{"zeroDelay", time.Second * 0},
		{"quater second delay", time.Millisecond * 250},
		{"half second delay", time.Millisecond * 500},
	}

	for _, e := range theTests {
		order = order
		eatTime = e.delay
		main()

		if len(*order.list) != 5 {
			t.Errorf("incorrect length of slice; expected 5 but got %d", len(*order.list))
		}
	}
}
