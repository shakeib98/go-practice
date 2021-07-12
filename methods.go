package main

import "fmt"

type User struct {
	name string
	age  int
}

func (u User) ChangeName() string {
	u.name = "Shahrukh"
	return u.name
}

func PrintName(u User) string {
	return u.name
}

func main() {
	u := User{"Shakeib", 10}
	fmt.Println(PrintName(u))
	fmt.Println(u.ChangeName())
}
