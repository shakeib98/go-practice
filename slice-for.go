package main

import "fmt"

func main() {
	var a = []int{1, 2, 3, 4, 5, 5}

	for i, v := range a {
		fmt.Printf("index %d , value %v \n", i, v)
	}
}
