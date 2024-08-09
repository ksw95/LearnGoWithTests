package main

import (
	"fmt"
)

const (
	spanish             = "Spanish"
	french              = "French"
	japanese            = "Japanese"
	englishHelloPrefix  = "Hello, "
	spanishHelloPrefix  = "Hola, "
	frenchHelloPrefix   = "Bonjour, "
	japaneseHelloPrefix = "Konnichiwa, "
)

func Hello(name, language string) string {
	if name == "" {
		name = "World"
	}

	return getLanguagePrefix(language) + name
}

func getLanguagePrefix(language string) (prefix string) {
	switch language {
	case spanish:
		prefix = spanishHelloPrefix
	case french:
		prefix = frenchHelloPrefix
	case japanese:
		prefix = japaneseHelloPrefix
	default:
		prefix = englishHelloPrefix
	}
	return
}

func main() {
	//fmt.Println(quote.Go())
	fmt.Println(Hello("world", ""))
}
