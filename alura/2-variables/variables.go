package main

import (
	"fmt"
	"reflect"
)

func main() {
	// var name string = "Caio Baeta"
	// var age int = 31
	// var version float32 = 1.2
	// var name = "Caio Baeta"
	// var age = 31
	// var version = 1.2
	name := "Caio Baeta"
	age := 31
	version := 1.2
	fmt.Println("Hello,", name+"!")
	fmt.Println("You have", age, "years old!")
	fmt.Println("This program is in version", version, "!")
	fmt.Println("Type of my variables:", reflect.TypeOf(name), reflect.TypeOf(age), reflect.TypeOf(version))
}
