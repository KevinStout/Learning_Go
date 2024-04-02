package main

import (
	"fmt"
	"log"

	"example.com/greetings"
)

func main() {
	// Ser properties of the pedefined Logger, including the log entry prefix and a flag to disable printing
	// the time, source file, and line number.
	log.SetPrefix("greetings: ")
	log.SetFlags(0)

	// A slice of names.
	names := []string{"Mom", "Dad", "Brother"}

	// Request a greeting message for the names.
	messages, err := greetings.Hellos(names)
	// If an error was returned, print it to the console and exit the program.
	if err != nil {
		log.Fatal(err)
	}

	// If no error was returned, print the returned map of messages to the console.
	fmt.Println(messages)

	// Request a greeting message.
	message, err := greetings.Hello("Mom")
	// If an error was returned, print it to the console and exit the program.
	if err != nil {
		log.Fatal(err)
	}

	// If no error was returned, print the returned message to the console.
	fmt.Println(message)
}
