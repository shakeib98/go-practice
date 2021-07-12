package main

import "fmt"

var s = []int{1, 2, 3, 4, 5}

func main() {
	for i, v := range s {
		fmt.Println(i, v)
	}
}
