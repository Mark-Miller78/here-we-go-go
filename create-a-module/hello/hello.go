package main

import (
	"fmt"
	"here-we-go-go/greetings"
	"log"
)

func main() {
	// Set properties of the predefined Logger, including
	// the log entry prefix and a flag to disable printing
	// the time, source file, and line number.
	log.SetPrefix("greetings: ")
	log.SetFlags(0)

	//A slice of names
	names := []string{"mark", "kevin", "carter"}

	//request greeting messages for the names
	messages, err := greetings.Hellos(names)

	//if error was returned, print it to the consol and exit program
	if err != nil {
		log.Fatal(err)
	}

	//if no error was returned, print the returned message to the console
	fmt.Println(messages)
}
