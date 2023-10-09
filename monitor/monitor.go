package main

import (
	"fmt"
)

func main() {
	fmt.Println("1- Start monitoring")
	fmt.Println("2- Show logs")
	fmt.Println("0- Exit")
	var command int
	// fmt.Scanf("%d", &command)
	fmt.Scan(&command)
	fmt.Println("You chose option", command)
}
