package main

import (
	"flag"
	"fmt"
)

func main() {
	fmt.Println(fmt.Sprintf("%d", []int{1, 2, 3, 4, 5}))
	var flagvar int
	flag.IntVar(&flagvar, "flagname", 1234, "help message for flagname")
}
