package main

import (
    "fmt"

    "github.com/antoniocarelli/GoLang_curso/greetings"
)

func main() {
    // Get a greeting message and print it.
    message := greetings.Hello("Gladys")
    fmt.Println(message)
}
