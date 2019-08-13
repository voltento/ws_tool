package handlers

import (
	"errors"
	"fmt"
	"github.com/voltento/WsTool/WebSocketClient"
	"github.com/voltento/WsTool/command"
	"os"
)

func InitMgrHandlers(mgr *HandlerManager) {
	var er error
	er = mgr.AddHandler("<", Handler(readMessage))
	if er != nil {
		fmt.Printf("Error occured on read message. Reason: `%v`", er)
		os.Exit(2)
	}
	er = mgr.AddHandler(">", Handler(writeMessage))
}

func readMessage(ws *WebSocketClient.WebSocket, _ command.Command) error {
	msg, er := ws.ReadOneMessage()
	if er != nil {
		return errors.New(fmt.Sprintf("Error occured on in `readMessage` handler. Reason: `%v`", er))
	}
	fmt.Printf("< %v\n", msg)
	return nil
}

func writeMessage(ws *WebSocketClient.WebSocket, cmd command.Command) error {
	er := ws.SendMessage(cmd.Args)
	if er != nil {
		return errors.New(fmt.Sprintf("Error occured on in `writeMessage` handler. Reason: `%v`", er))
	}
	fmt.Printf("> %v\n", cmd.Args)
	return nil
}
