package main

import (
	"fmt"
	"log"

	"example.com/greetings"
)

func main() {
	log.SetPrefix("greetings: ")
	log.SetFlags(0)

	names := []string{"Matthew", "Mark", "Luke", "John"}
	messages, err := greetings.Hellos(names)
	// message, err := greetings.Hello("John")
	// message, err := greetings.Hello("")
	if err != nil {
		log.Fatal(err)
	}

	for _, message := range messages {
		fmt.Println(message)
	}
}
