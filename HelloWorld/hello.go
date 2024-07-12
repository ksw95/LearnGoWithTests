package main

import (
	"fmt"

	"rsc.io/quote"
)

func Hello(name string) string {
	return "Hello, " + name
}

func main() {
	fmt.Println(quote.Go())
	fmt.Println(Hello("world"))
}
