package main

import (
    "fmt"

    "github.com/antoniocarelli/GoLang_curso/tree/Tutorial_2/greetings"
)

func main() {
    // Get a greeting message and print it.
    message := greetings.Hello("Carelli")
    fmt.Println(message)
}
