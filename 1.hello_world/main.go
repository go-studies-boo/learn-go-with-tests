package main

import "fmt"

func main() {
	fmt.Println(Hello("world", ""))
}

const englishHelloFixed = "Hello, "
const spanishHelloFixed = "Hola, "
const frenchHelloFixed = "Bonjour, "

func Hello(name string, language string) string {
	if name == "" {
		name = "World"
	}

	return greetingPrefix(language) + name
}

func greetingPrefix(language string) (prefix string) { // named return value; It will be assigned the "zero" value. 
	switch language {
	case "spanish":
		prefix = spanishHelloFixed
	case "french":
		prefix = frenchHelloFixed
	default:
		prefix = englishHelloFixed
	}

	return prefix
}