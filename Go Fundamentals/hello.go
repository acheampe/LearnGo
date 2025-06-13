package main

import "fmt" // imports a package that contains the println func

func hello() string {
	return "Hello, world"
}

func main() {
	fmt.Println(hello())
}
