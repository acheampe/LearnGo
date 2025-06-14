package main

import "fmt" // imports a package that contains the println func

const englishHelloPrefix string = "Hello, "

func hello(name string) string {
	return englishHelloPrefix + name
}

func main() {
	fmt.Println(hello("Chris"))
}
