package main

import "testing"

func TestSearchArray(t *testing.T) {
	arr := []int{8, 3, 6, 1, 4, 7, 2}
	expectedKey := 1
	expectedIndex := 3
	actualIndex := FindInArray(arr, expectedKey)
	if expectedIndex != actualIndex {
		t.Errorf("Got the wrong key")
	}
}
