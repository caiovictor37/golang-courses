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
		fmt.Println("Start monitoring")
	case 2:
		fmt.Println("Show logs")
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
