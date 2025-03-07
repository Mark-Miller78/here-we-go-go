package greetings

import (
	"errors"
	"fmt"
	"math/rand"
)

// Hello returns a greeting from the named person
func Hello(name string) (string, error) {
	//if no name was give, return an error with message
	if name == "" {
		return "", errors.New("empty name")
	}
	//Return a greeting that embeds the name in the message
	message := fmt.Sprintf(randomFormat(), name)
	return message, nil
}

// Hellos returns a map that associates each of the named people with a greeting message
func Hellos(names []string) (map[string]string, error) {
	//a map to associate names with the messages.
	messages := make(map[string]string)
	//Loop through the recieved slice of names, calling the Hello function to get a message for each name
	for _, name := range names {
		message, err := Hello(name)
		if err != nil {
			return nil, err
		}
		//in the map, associate the retrieved message with the name
		messages[name] = message
	}
	return messages, nil
}

// randomFormat returns on of a set of greeting messages. The returned message is selected at random
func randomFormat() string {
	//a slice of message formats
	formats := []string{
		"Hi, %v. Welcome!",
		"Great to see you, %v!",
		"Hail, %v! Well met!",
	}

	//Return a randomly selected message format by specifying a random index for the slice of formats
	return formats[rand.Intn(len(formats))]
}
