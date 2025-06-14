package main

import "fmt" // imports a package that contains the println func

const englishHelloPrefix string = "Hello, "

func Hello(name string) string {

	if name == "" {
		name = "World"
	}
	return englishHelloPrefix + name
}

func main() {
	fmt.Println(Hello("Chris"))
}
