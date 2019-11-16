package main

import "fmt"

func FindInArray(arr []int, key int) int {
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
	fmt.Println("Hello World")
}
