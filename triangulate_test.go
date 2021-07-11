package main

import (
	"testing"
)

func TestCcw(t *testing.T) {
	a := Point{
		x: 1,
		y: 2,
	}

	b := Point{
		x: 2,
		y: 3,
	}

	c := Point{
		x: 3,
		y: 4,
	}

	if ans := Ccw(a, b, c); ans == 0 {
		t.Log("wrong")
		t.Fail()
	}
}
