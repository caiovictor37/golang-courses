package main

import (
	"fmt"
	"os"
)

func main() {
	showOptions()
	command := chooseOption()

	// if command == 1 {
	// 	fmt.Println("Start monitoring")
	// } else if command == 2 {
	// 	fmt.Println("Show logs")
	// } else if command == 0 {
	// 	fmt.Println("Exit")
	// 	os.Exit(0)
	// } else {
	// 	fmt.Println("Invalid option")
	// 	os.Exit(-1)
	// }

	switch command {
	case 1:
		fmt.Println("Type name and surname")
		name, surname := nameAndSurname()
		fmt.Print(name, surname)
	case 2:
		fmt.Println("Type name and surname")
		name, _ := nameAndSurname()
		fmt.Print(name)
	case 0:
		fmt.Println("Exit code succesfully")
		os.Exit(0)
	default:
		fmt.Println("Invalid option")
		os.Exit(-1)
	}
}

func showOptions() {
	fmt.Println("1- Start monitoring")
	fmt.Println("2- Show logs")
	fmt.Println("0- Exit")
}

func chooseOption() int {
	var option int
	// fmt.Scanf("%d", &command)
	fmt.Scan(&option)
	fmt.Println("You chose option", option)
	return option
}

func nameAndSurname() (string, string) {
	var name string
	var surname string
	fmt.Scan(&name, &surname)
	return name, surname
}
