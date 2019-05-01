package main

import "fmt"
import "os"
import "net/http"

func main() {

	for {
		showIntroduction()
		showMenu()
		command := readCommand()

		switch command {
		case 1:
			startMonitoring()
		case 2:
			fmt.Println("Showing Logs...")
		case 0:
			fmt.Println("Exiting program")
			os.Exit(0)
		default:
			fmt.Println("Doesn't reconize this command")
			os.Exit(-1)
		}
	}
}

func showIntroduction() {
	name := "Arthur"
	version := 1.1

	fmt.Println("Hello, sr.", name)
	fmt.Println("This program is in version", version)
}

func showMenu() {
	fmt.Println("1- Start Monitoring")
	fmt.Println("2- Show Logs")
	fmt.Println("0- Exit")
}

func readCommand() int {
	var command int
	fmt.Scan(&command)

	return command
}

func startMonitoring() {
	fmt.Println("Monitoring...")
	site := "https://random-status-code.herokuapp.com"
	res, _ := http.Get(site)

	if res.StatusCode == 200 {
		fmt.Println("Site: ", site, "was loaded with sucess!")
	} else {
		fmt.Println("Site: ", site, "has problems. Status Code:", res.StatusCode)
	}

}
