package greetings

import (
	"errors"
	"fmt"
	"math/rand"
)

// Note that functions that have a capital letter at the begining of their name are exported functions.
// function that have a lowercase letter at the begining of their name are private functions.

// Hello returns a greeting for the named person.
func Hello(name string) (string, error) {
	// If no name was given, return an error with a message.
	if name == "" {
		return "", errors.New("empty name")
	}
	// Return a greeting that embeds the name in a message.
	message := fmt.Sprintf(randomFormat(), name)
	return message, nil
}

// randomFormat returns one of a set of greeting messages.
// The returned message is selected at random.
func randomFormat() string {
	// A slice of message formats.
	// A slice is a dynamically-sized, flexible view into the elements of an array.
	formats := []string{
		"Hi, %v. Welcome!",
		"Great to see you, %v!",
		"Hail, %v! Well met!",
	}
	// Return a randomly selected message format by specifying a random index for the slice of formats.
	return formats[rand.Intn(len(formats))]
}