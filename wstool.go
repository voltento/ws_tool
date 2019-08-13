package main

import (
	"bufio"
	"fmt"
	"github.com/voltento/WsTool/command"
	"github.com/voltento/WsTool/handlers"
	"github.com/voltento/WsTool/utils"
	"github.com/voltento/WsTool/web_socket_client"
	"os"
)

func main() {

	filePath := "/Users/vladimir.dumanovskiy/go/src/github.com/voltento/WsTool/commands"
	commands := command.CreateReaderFromFile(filePath)
	adress, headers := utils.ParseArgs()

	ws := new(web_socket_client.WebSocket)
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

func printMessageFromWs(ws *web_socket_client.WebSocket) {
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

func readFromConsoleAndSendToWs(ws *web_socket_client.WebSocket) {
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
