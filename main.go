package main

import (
	"fmt"

	"github.com/google/uuid"
	"github.com/obey/go-example/test"
)

func main() {
	id := uuid.New()
	fmt.Println("Hello world 2")
	fmt.Printf("UUID: %s\n", id)
	test.SayHelloTest()
	test.SayHelloTest2()
}
