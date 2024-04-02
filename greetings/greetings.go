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

// Hellos returns a map that associates each of the named people with a greeting message.
func Hellos(names []string) (map[string]string, error) {
	// A map to associate name with messages.
	// In GO a map is initialized with the following syntax: make(map[key-type]value-type).
	messages := make(map[string]string)

	// Loop through the received slice of names, calling the Hello function to get a message for each name.
	// The range returns two values: the index of the current item in the loop and a copy of the item's value.
	// Since the index is not used, you can ignore it by using the blank identifier (_).
	for _, name := range names {
		message, err := Hello(name)
		if err != nil {
			return nil, err
		}
		// In the map, associate the retrieved massage with the name.
		// This syntax is map[key] = value.
		messages[name] = message
	}
	return messages, nil
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
