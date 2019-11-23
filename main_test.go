package main

import (
	"fmt"
	"math/rand"
	"testing"
)

type LinearSearchTests struct {
	array         []int
	searchKey     int
	expectedIndex int
}

func GenerateSearchTests() []LinearSearchTests {
	const totalArrays = 100
	linearSearchTests := make([]LinearSearchTests, totalArrays)
	r := rand.New(rand.NewSource(0))
	for i := 0; i < totalArrays; i++ {
		arrayLen := r.Intn(10) + 2
		linearSearchTests[i].array = make([]int, arrayLen)
		for arrayIndex := 0; arrayIndex < arrayLen; arrayIndex++ {
			linearSearchTests[i].array[arrayIndex] = r.Intn(1000)
		}
		expectedIndex := arrayLen / 2
		linearSearchTests[i].expectedIndex = expectedIndex
		linearSearchTests[i].searchKey = linearSearchTests[i].array[expectedIndex]
	}
	return linearSearchTests
}

func FormattedTestName(testData LinearSearchTests) string {
	result := "Test for, array: " + fmt.Sprintf("%d", testData.array) + ", "
	result += "with the search key: " + fmt.Sprintf("%d", testData.searchKey) + ", "
	result += "with the expected index: " + fmt.Sprintf("%d", testData.expectedIndex) + "\n"
	return result
}

func TestSearchArray(t *testing.T) {
	for _, tt := range GenerateSearchTests() {
		t.Run(FormattedTestName(tt), func(t *testing.T) {
			actualIndex := LinearSearch(tt.array, tt.searchKey)
			if tt.expectedIndex != actualIndex {
				t.Errorf("Got the wrong key")
			}
		})
	}
}

func TestParallelSearchArray(t *testing.T) {
	for _, tt := range GenerateSearchTests() {
		t.Run(FormattedTestName(tt), func(t *testing.T) {
			actualIndex := LinearSearchParallel(tt.array, tt.searchKey)
			if tt.expectedIndex != actualIndex {
				t.Errorf("Got the wrong key " + fmt.Sprintf("%d", actualIndex))
			}
		})
	}

}
