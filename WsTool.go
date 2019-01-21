package main

import (
	"bufio"
	"fmt"
	"github.com/voltento/WsTool/WebSocketClient"
	"net/http"
	"os"
	"strings"
)

type Adress string

func printHelp() {
	fmt.Println("Args: url [-H \"HeaderName: Header Value\"]")
}

func parseArgs() (Adress, http.Header) {
	if len(os.Args) == 1 || os.Args[1] == "--help" {
		printHelp()
		os.Exit(0)
	}

	headers := http.Header{}
	argIndex := 2
	for argIndex < len(os.Args) {
		if os.Args[argIndex] != "-H" {
			fmt.Printf("Can't parse arg value. Value: %s\n", os.Args[argIndex])
			printHelp()
			os.Exit(1)
		}
		argIndex += 1
		if argIndex == len(os.Args) {
			fmt.Printf("Value for header flag wasn't provided\n")
			printHelp()
			os.Exit(1)
		}
		header := strings.SplitAfterN(os.Args[argIndex], ":", 2)
		if len(header) != 2 {
			fmt.Printf("Wrong header value. Header: %s\n", os.Args[argIndex])
			printHelp()
			os.Exit(1)
		}

		headers.Add(header[0][:len(header[0])-1], header[1])
		argIndex += 1
	}

	return Adress(os.Args[1]), headers
}

func main() {
	var adress Adress
	var headers http.Header
	adress, headers = parseArgs()

	ws := new(WebSocketClient.WebSocket)
	if er := ws.Connect(string(adress), headers); er != nil {
		println("Error on connection. Reason: " + er.Error())
		return
	}

	go printMessageFromWs(ws)

	readFromConsloeAndSendToWs(ws)
}

func printMessageFromWs(ws *WebSocketClient.WebSocket) {
	var err error
	var msg string
	for {
		msg, err = ws.ReadOneMessage()
		if err != nil {
			fmt.Printf("Error occurred during read from socket. Error: %s\n", err.Error())
			os.Exit(0)
		}
		print("< : ", msg, "\n")
	}
}

func readFromConsloeAndSendToWs(ws *WebSocketClient.WebSocket) {
	for {
		reader := bufio.NewReader(os.Stdin)

		var messageToWs string
		var err error

		messageToWs, err = reader.ReadString('\n')
		if err != nil {
			fmt.Printf("Can't read from console. Error: %s\n", err.Error())
		}

		err = ws.SendMessage(messageToWs)
		if err != nil {
			panic("Error occurred during send message to ws. Error: " + err.Error())
		} else {
			fmt.Printf(" > : %s\n", messageToWs)
		}
	}
}
