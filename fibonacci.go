package main

import "fmt"

func fibonacci() func(first, second int) int {
	return func(first, second int) int {
		return first + second
	}
}

func swap(answer, second int) (int, int) {
	return second, answer
}

var start int = 0
var second int = 1

func main() {
	f := fibonacci()
	fmt.Println(start)
	fmt.Println(second)
	for i := 0; i < 10; i++ {
		answer := f(start, second)
		start, second = swap(answer, second)
		fmt.Println(answer)
	}
}
