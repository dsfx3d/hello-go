package main

import (
	"fmt"
	"log"

	"dsfx3d/greetings"

	"rsc.io/quote"
)

func main() {
	log.SetPrefix("greetings: ")
	log.SetFlags(0)

	fmt.Println(quote.Glass())

	message, error := greetings.Hello("foo")
	fmt.Println(message, error)

	m, err := greetings.Hello("")

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println((m))
}