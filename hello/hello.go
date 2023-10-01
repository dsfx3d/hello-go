package main

import (
	"fmt"

	"dsfx3d/greetings"

	"rsc.io/quote"
)

func main() {
	fmt.Println(quote.Glass())

	message := greetings.Hello("foo")
	fmt.Println(message)
}