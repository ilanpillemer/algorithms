package main

import (
	"testing"
)

func TestSimple(t *testing.T) {
	example := poly([]int{-1, 2, -6, 2})
	got := simple(3, example)
	if got != 5 {
		t.Errorf("want 5 got %d", got)
	}
}

func TestHorner(t *testing.T) {
	example := poly([]int{-1, 2, -6, 2})
	got := horner(3, example)
	if got != 5 {
		t.Errorf("want 5 got %d", got)
	}
}
