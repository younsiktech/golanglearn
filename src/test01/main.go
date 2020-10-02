package main

import (
	"fmt"
	animal "github.com/younsiktech/golanglearn/src/test01/modules/animal"
	hello "github.com/younsiktech/golanglearn/src/test01/modules/hello"
)

func main() {
	fmt.Println("test")
	fmt.Println(hello.Hello())
	fmt.Println(animal.DogBark())
	fmt.Println(animal.CatMeow())
}
