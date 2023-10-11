package main

import (
	"fmt"
	"net/http"
	"os"
)

func main() {
	for {
		showOptions()
		command := chooseOption()
		switch command {
		case 1:
			fmt.Println("Start monitoring")
			startMonitoring()
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

func startMonitoring() {

	sitesSlice := []string{"https://random-status-code.herokuapp.com",
		"https://www.alura.com.br",
		"https://www.caelum.com.br"}

	for _, site := range sitesSlice {
		resp, err := http.Get(site)
		if err == nil {
			if resp.StatusCode == 200 {
				fmt.Println("Site", site, "HTTP Status Successful", resp.StatusCode)
			} else {
				fmt.Println("Site", site, "HTTP Status Error", resp.StatusCode)
			}
		} else {
			fmt.Println("Site", site, "Error", err)
		}
	}

}
