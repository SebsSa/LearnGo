package main

import "fmt"

//Constants are defined like so
const englishHelloPrefix = "Hello, "

func hello(input string) string {
	if input == "" {
		input = "World"
	}
	return englishHelloPrefix + input
}

func main() {

	fmt.Println(hello(""))

}
