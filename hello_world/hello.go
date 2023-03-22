package main

import "fmt"

const russian = "Russian"
const french = "French"
const englishHelloPrefix = "Hello, "
const russianHelloPrefix = "Здорова, "
const frenchHelloPrefix = "Bonjour, "

func Hello(name string, language string) string {
	if name == "" {
		name = "World"
	}
	if language == russian {
		return russianHelloPrefix + name
	}
	if language == french {
		return frenchHelloPrefix + name
	}
	return englishHelloPrefix + name
}

func main() {
	fmt.Println(Hello("world", ""))
}
