package array

import (
	"fmt"
	"runtime"
)

func LinearSearch(arr []int, key int) int {
	index := -1
	for i := 0; i < len(arr); i++ {
		if key == arr[i] {
			index = i
			break
		}
	}
	return index
}

func LinearSearchParallel(arr []int, key int) int {
	parallelThreads := runtime.NumCPU()
	searchResultIndex := make(chan int, parallelThreads)
	sliceLength := len(arr) / parallelThreads
	if sliceLength == 0 {
		sliceLength = 1
		parallelThreads = len(arr)
	}
	fmt.Println("Slice length is", sliceLength, " total threads are", parallelThreads, " length of the array is ", len(arr))
	for threadIndex := 0; threadIndex < parallelThreads; threadIndex++ {
		go func(threadIndex int) {
			subArray := arr[threadIndex*sliceLength : (threadIndex+1)*sliceLength]
			subArrayIndex := LinearSearch(subArray, key)
			if subArrayIndex != -1 {
				searchResultIndex <- (threadIndex*sliceLength + subArrayIndex)
				return
			}
			searchResultIndex <- -1
		}(threadIndex)
	}
	result := <-searchResultIndex
	for threadIndex := 1; threadIndex < parallelThreads; threadIndex++ {
		if sliceResult := <-searchResultIndex; sliceResult != -1 {
			if result < sliceResult {
				result = sliceResult
			}
		}
	}
	return result
}
