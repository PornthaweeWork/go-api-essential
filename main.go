package main

import (
	"fmt"

	"github.com/google/uuid"
	// "github.com/obey/go-example/test"
)

func main() {
	id := uuid.New()
	fmt.Println("Hello world 2")
	fmt.Printf("UUID: %s\n", id)
	fullname := "Obey"
	age := 25
	switch age {
	case 25:
		fmt.Println("You are 25 years old")
	case 26:
		fmt.Println("You are 26 years old")
	default:
		fmt.Println("You are not 25 or 26 years old")
	}
	fmt.Printf("Hello %s Yay! age = %d\n", fullname, age)

	for i := 0; i < 5; i++ {
		fmt.Println(i)
	}
}
