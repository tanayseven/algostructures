package main

import "fmt"

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

func main() {
	fmt.Println(fmt.Sprintf("%d", []int{1, 2, 3, 4, 5}))
}
