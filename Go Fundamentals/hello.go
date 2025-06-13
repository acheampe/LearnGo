package main

import "fmt" // imports a package that contains the println func

func hello(name string) string {
	return "Hello, " + name
}

func main() {
	fmt.Println(hello("Chris"))
}
