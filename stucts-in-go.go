package main

import "fmt"

type User struct {
	name string
	age  int
}

func main() {
	user := User{"shakeib", 22}
	fmt.Println(user.name, user.age)
}
