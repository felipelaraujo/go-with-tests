package main

import (
	"fmt"
)

const (
	portuguese = "Portuguese"
	spanish    = "Spanish"
	french     = "French"

	portugueseHelloPrefix = "Olá, "
	englishHelloPrefix    = "Hello, "
	spanishHelloPrefix    = "Hola, "
	frenchHelloPrefix     = "Bonjour, "
)

func Hello(name string, language string) string {
	if name == "" {
		name = "world"
	}

	return greetingPrefix(language) + name + "!"
}

func greetingPrefix(language string) (prefix string) {
	switch language {
	case portuguese:
		prefix = portugueseHelloPrefix
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
	fmt.Println(Hello("Felipe", ""))
}
