package main

import "fmt"

//Constants are defined like so
const englishHelloPrefix = "Hello, "
const spanishHelloPrefix = "Hola, "
const frenchHelloPrefix = "Bonjour, "

func hello(input string, language string) string {

	if input == "" {
		input = "World"
	}

	if language == "Spanish" {
		return spanishHelloPrefix + input
	}

	if language == "French" {
		return frenchHelloPrefix + input
	}

	return englishHelloPrefix + input
}

func main() {

	fmt.Println(hello("", ""))

}
