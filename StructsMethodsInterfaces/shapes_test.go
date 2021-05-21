package main

import "testing"

func TestPerimeter(t *testing.T) {
	got := Perimeter(10.0, 10.0)
	want := 40.0

	if got != want {
		t.Errorf("got %.2f want %.2f", got, want)
	}
}

func TestArea(t *testing.T) {
	got := Area(9.0, 3.0)
	want := 27.0

	if got != want {
		t.Errorf("got %.2f want %.2f", got, want)
	}
}
