package test

import (
	"fmt"

	"github.com/obey/go-example/test/internal/obey"
)

func SayHelloTest() {
	fmt.Println("Hello world")
	obey.SayObey()
}
