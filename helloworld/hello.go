package main

import "fmt"

const (
	spanish = "Spanish"
	french  = "French"
	english = "English"

	englishHelloPrefix = "Hello, "
	spanishHelloPrefix = "Hola, "
	frenchHelloPrefix  = "Bonjour, "

	exclamationPostfix = "!"
)

func Hello(name, language string) string {
	if name == "" {
		name = "World"
	}
	if language == "" {
		language = english
	}

	return greetingPrefix(language) + name + exclamationPostfix
}

func greetingPrefix(language string) (prefix string) {
	switch language {
	case spanish:
		prefix = spanishHelloPrefix
	case french:
		prefix = frenchHelloPrefix
	default:
		prefix = englishHelloPrefix
	}
	return
}

func main() {
	fmt.Println(Hello("John", "Spanish"))
}
