package main

import (
	"bufio"
	"fmt"
	"github.com/voltento/WsTool/Utils"
	"github.com/voltento/WsTool/WebSocketClient"
	"github.com/voltento/WsTool/command"
	"github.com/voltento/WsTool/handlers"
	"os"
)

func main() {

	filePath := "/Users/vladimir.dumanovskiy/go/src/github.com/voltento/WsTool/commands"
	commands := command.CreateReaderFromFile(filePath)
	adress, headers := Utils.ParseArgs()

	ws := new(WebSocketClient.WebSocket)
	if er := ws.Connect(string(adress), headers); er != nil {
		println("Error on connection. Reason: " + er.Error())
		os.Exit(1)
	}

	mgr := handlers.CreateHandlerManager(ws)
	for cmd := range commands {
		if er := mgr.Handle(cmd); er != nil {
			fmt.Printf("Error occured. Error: %v", er.Error())
			os.Exit(1)
		}
	}

	go printMessageFromWs(ws)

	readFromConsoleAndSendToWs(ws)
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

func readFromConsoleAndSendToWs(ws *WebSocketClient.WebSocket) {
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
