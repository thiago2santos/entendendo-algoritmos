package sorting

import (
	"testing"
)

func TestGetMin(t *testing.T) {
	arr := []int{5, 4, 3, 2, 1}
	min := GetMin(arr)
	if min != 4 {
		t.Errorf("Expected 1, but got %d", min)
	}
}
