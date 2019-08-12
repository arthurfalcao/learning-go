package main

import (
	"bufio"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"strconv"
	"strings"
	"time"
)

const monitoring = 3
const delay = 10

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
			printLogs()
		case 0:
			fmt.Println("Exiting program")
			os.Exit(0)
		default:
			fmt.Println("Doesn't recognize this command")
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
	// sites := []string{"https://random-status-code.herokuapp.com", "http://www.flowcorp.com.br"}

	sites := readSitesFromFile()

	for i := 0; i < monitoring; i++ {
		for _, site := range sites {
			testSite(site)
		}
		time.Sleep(delay * time.Minute)
		fmt.Println("")
	}

	fmt.Println("")
}

func testSite(site string) {
	res, err := http.Get(site)

	if err != nil {
		fmt.Println("Ocorreu um erro:", err)
	}

	if res.StatusCode == 200 {
		fmt.Println("Site: ", site, "was loaded with sucess!")
		registryLog(site, true)
	} else {
		fmt.Println("Site: ", site, "has problems. Status Code:", res.StatusCode)
		registryLog(site, false)
	}
}

func readSitesFromFile() []string {
	var sites []string

	file, err := os.Open("files.txt")
	if err != nil {
		fmt.Println("Ocorreu um erro:", err)
	}

	reader := bufio.NewReader(file)
	for {
		line, err := reader.ReadString('\n')
		line = strings.TrimSpace(line)
		fmt.Println(line)

		sites = append(sites, line)

		if err == io.EOF {
			break
		}
	}

	_ = file.Close()
	return sites
}

func registryLog(site string, status bool) {
	file, err := os.OpenFile("log.txt", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)

	if err != nil {
		fmt.Println(err)
	}

	_, _ = file.WriteString(time.Now().Format("02/01/2006 15:04:05") + " - " + site + " - online: " + strconv.FormatBool(status) + "\n")

	_ = file.Close()
}

func printLogs() {

	file, err := ioutil.ReadFile("log.txt")

	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(string(file))
}
