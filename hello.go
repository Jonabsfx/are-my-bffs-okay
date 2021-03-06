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

func main() {

	showIntroduction()

	for {

		showMenu()

		command := readCommand()

		switch command {
		case 1:
			startMonitor()
		case 2:
			fmt.Println("Showing Logs...")
			logPrint()
		case 0:
			fmt.Println("Exiting...")
			os.Exit(0)
		default:
			fmt.Println("That command doesn't exist")
			os.Exit(-1)
		}
	}

}

func showIntroduction() {
	name := "Are my best friends okay?"
	version := 1.1

	fmt.Println(name)
	fmt.Println("This build is in version", version)
}

func showMenu() {
	fmt.Println("1 - Check if your best friends are ok")
	fmt.Println("2 - Show the logs")
	fmt.Println("0 - Exit")
}

func readCommand() int {
	var command int
	fmt.Scanf("%d", &command)
	fmt.Println("The command that you choose was:", command)

	return command
}

func startMonitor() {
	var times int
	var period int

	fmt.Println("How many times do you want to check if your best friends are ok?")
	fmt.Scanf("%d", &times)

	fmt.Println("OK. How often? In seconds, please.")
	fmt.Scanf("%d", &period)

	urls := readUrlFile()
	names := []string{"Your scientist friend", "Your libertarian friend", "Your smart freind",
		"Your lifeguard friend", "Your relaxing friend"}

	for i := 0; i <= times; i++ {
		for i, url := range urls {
			urlTest(url, names[i])
		}

		time.Sleep(time.Duration(period) * time.Second)
		fmt.Println("")
	}
}

func urlTest(url string, name string) {
	resp, err := http.Get(url)

	if err != nil {
		fmt.Println("WARNING! ERROR:", err)
	}

	if resp.StatusCode == 200 {
		fmt.Println(name, "is OK")
		logRegister(url, true)
	} else {
		fmt.Println(name, "is not OK. Status Code:", resp.StatusCode)
		logRegister(url, false)
	}
}

func readUrlFile() []string {
	var urls []string

	file, err := os.Open("urls.txt")

	if err != nil {
		fmt.Println("WARNING! ERROR:", err)
	}

	reader := bufio.NewReader(file)
	for {
		line, err := reader.ReadString('\n')
		line = strings.TrimSpace(line)

		urls = append(urls, line)

		if err == io.EOF {
			break
		}

	}

	file.Close()
	return urls
}

func logRegister(url string, status bool) {

	file, err := os.OpenFile("log.txt", os.O_CREATE|os.O_RDWR|os.O_APPEND, 0666)

	if err != nil {
		fmt.Println("WARNING! ERROR:", err)
	}
	file.WriteString(time.Now().Format("02/01/2006 15:04:05") + " - " + url +
		" - online: " + strconv.FormatBool(status) + "\n")

	file.Close()
}

func logPrint() {

	file, err := ioutil.ReadFile("log.txt")

	if err != nil {
		fmt.Println("WARNING! ERROR:", err)
	}

	fmt.Println(string(file))
}
