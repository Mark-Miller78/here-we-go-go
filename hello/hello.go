package main

import (
	"fmt"
	"here-we-go-go/greetings"
)

func main() {
	//Get a greeting message and print it
	message := greetings.Hello("Gladys")
	fmt.Println(message)
}
