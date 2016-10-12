package main

import "testing"

func TestTrecimate(t *testing.T) {
	var expected = 1
	var actual = trecimate(33)
	if actual != expected {
		t.Errorf("%d was not equal to %d", actual, expected)
	}
}
