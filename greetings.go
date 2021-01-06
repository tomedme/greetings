package greetings

import (
	"errors"
	"fmt"
)

// Hello returns a greeting for the named person
func Hello(name string) (string, error) {
	// If no name given, return error with a message
	if name == "" {
		return "", errors.New("error: empty name")
	}

	// Return a greeting that embeds the name in a message
	message := fmt.Sprintf("Hi %v!", name)
	return message, nil
}
