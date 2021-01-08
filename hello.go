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

	varHello := englishHelloPrefix

	switch language {
	case "French":
		varHello = frenchHelloPrefix
	case "Spanish":
		varHello = spanishHelloPrefix
	}

	return varHello + input
}

func add(a int, b int) int {

	result := a + b
	return result

}

func main() {

	fmt.Println(hello("", ""))
	fmt.Println(add(2, 3))

}
