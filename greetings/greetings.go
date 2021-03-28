package greetings

import (
	"fmt"
)

// Hello returns a greeting for the named person.
func Hello(name string) string {
		//Declaração longa:
		//var message string
		//message = fmt.Sprintf("Hi, %v. Welcome!", name)

    // Return a greeting that embeds the name in a message.
    message := fmt.Sprintf("Hi, %v. Welcome!", name)
    return message
}
