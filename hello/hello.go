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

	names := []string{"foo", "baz", "bar"}

	messages, error := greetings.Hellos(names)

	if error != nil {
		log.Fatal(error)
	}
	
	fmt.Println(messages, error)
}