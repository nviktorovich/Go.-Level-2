package main

import(
	"testing"
)

func TestGetNumberFromSliceToIndex(t *testing.T) {
	var idx int
	idx = GetNumberFromSliceToIndex(1, []int{1, 2, 3, 4})
	if idx != 2 {
		t.Error("expected 3, got ", idx)
	}
}