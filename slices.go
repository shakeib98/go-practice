package main

import "fmt"

//slices are like reference to arrays

func main() {
	arr := [5]int{1, 2, 3, 4, 5}

	sl := arr[1:4]

	fmt.Println(sl)

	functionTwo()
	functionThree()
	structSlic()
}

func functionTwo() {

	names := [4]string{
		"John",
		"Paul",
		"George",
		"Ringo",
	}
	fmt.Println(names)

	a := names[0:2]
	b := names[1:3]
	fmt.Println(a, b)

	b[0] = "XXX"
	fmt.Println(a, b)
	fmt.Println(names)

}

func structSlic() {
	a := []struct {
		name string
	}{
		{"Shakeib"},
	}

	fmt.Println(a)
}

func functionThree() {
	s := []int{2, 3, 5, 7, 11, 13}
	printSlice(s)

	// Slice the slice to give it zero length.
	s = s[:0]
	printSlice(s)

	// Extend its length.
	s = s[:4]
	printSlice(s)

	// Drop its first two values.
	s = s[2:]
	printSlice(s)

	var x []int
	printSlice(x)
}

func printSlice(s []int) {
	if s == nil {
		fmt.Println("nil!")
	} else {
		fmt.Printf("len=%d cap=%d %v\n", len(s), cap(s), s)
	}

}
