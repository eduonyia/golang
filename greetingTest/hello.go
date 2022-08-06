package main

import (
	"fmt"

	"github.com/eduonyia/golang/greetings"
)

func main() {
	// Get a greeting message and print it.
	message := greetings.Hello("Gladys")
	fmt.Println(message)
}
