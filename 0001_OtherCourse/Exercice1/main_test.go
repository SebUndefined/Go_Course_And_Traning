package main

import "testing"

func TestGetSliceOfInts(t *testing.T) {
	number := 100
	sliceOfInts := getSliceOfInts(number)
	if len(sliceOfInts) != number+1 {
		t.Errorf("Expected lengh of %v, but got %v", number+1, len(sliceOfInts))
	}
}
