package main

import "fmt"

const (
	englishHelloPrefix = "Hello, "
	spanishHelloPrefix = "Hola, "
	frenchHelloPrefix  = "Bonjour, "
	russianHelloPrefix = "Привет, "
	spanish            = "Spanish"
	french             = "French"
	russian            = "Russian"
	defaultName        = "World!"
)

func Hello(name, language string) string {
	if name == "" {
		name = defaultName
	}

	return greetingPrefix(language) + name
}

func main() {
	fmt.Println(Hello("mf", ""))
}

func greetingPrefix(language string) (prefix string) {
	switch language {
	case spanish:
		prefix = spanishHelloPrefix
	case french:
		prefix = frenchHelloPrefix
	case russian:
		prefix = russianHelloPrefix
	default:
		prefix = englishHelloPrefix
	}
	return
}
