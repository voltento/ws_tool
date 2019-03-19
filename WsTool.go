package main

import (
	"bufio"
	"errors"
	"fmt"
	"github.com/voltento/WsTool/WebSocketClient"
	"net/http"
	"os"
	"strings"
)

type Adress string

func printHelp() {
	fmt.Println("Args: url [-H \"HeaderName: Header Value\"] [-C \"CookieName: Cookie Value\"]")
}

func parseHeaderKeyValue(s string) (string, string, error) {
	var (
		key   string
		value string
		er    error
	)
	header := strings.SplitAfterN(s, ":", 2)
	if len(header) != 2 {
		er = errors.New(fmt.Sprintf("Wrong header value. Header: %s\n", s))
	} else {
		key = header[0]
		key = key[:len(key)-1]
		value = header[1]
	}

	return key, value, er
}

func processError(er error) {
	if er != nil {
		fmt.Printf("Error occured. Error: ", er.Error())
		printHelp()
		os.Exit(1)
	}
}

const (
	header = iota
	coockie
	undefined
)

func parseArgs() (Adress, http.Header) {
	if len(os.Args) == 1 || os.Args[1] == "--help" {
		printHelp()
		os.Exit(0)
	}

	headers := http.Header{}
	argIndex := 2
	argType := undefined
	for argIndex < len(os.Args) {

		if os.Args[argIndex] == "-H" {
			argType = header
		} else if os.Args[argIndex] == "-C" {
			argType = coockie
		}

		if argType == undefined {
			er := errors.New(fmt.Sprintf("Can't parse arg value. Value: %s\n", os.Args[argIndex]))
			processError(er)
		}

		argIndex += 1
		if argIndex == len(os.Args) {
			er := errors.New(fmt.Sprintf("Value for header flag wasn't provided\n", os.Args[argIndex]))
			processError(er)
		}

		key, value, er := parseHeaderKeyValue(os.Args[argIndex])
		processError(er)

		if argType == header {
			headers.Add(key, value)
		} else if argType == coockie {
			ckoockie := http.Cookie{}
			ckoockie.Name = key
			ckoockie.Value = value
			headers.Add("Cookie", fmt.Sprintf("%v=%v;", key, value))
		}
		argIndex += 1
	}

	return Adress(os.Args[1]), headers
}

func main() {
	adress, headers := parseArgs()

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
