package main

import (
	"fmt"
	"math/rand"
)

func main() {

	fmt.Println("random number is: -> ", rand.Intn(100))
	fmt.Println("seed: -> ", rand.Seed)
}
