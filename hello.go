package main

import "fmt"

const (
	englishHelloPrefix = "Hello, "
	exclamationPostfix = "!"
)

func Hello(name string) string {
	if name == "" {
		name = "World"
	}
	return englishHelloPrefix + name + exclamationPostfix
}

func main() {
	fmt.Println(Hello("John"))
}
