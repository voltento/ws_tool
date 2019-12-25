package handlers

import (
	"errors"
	"fmt"
	"github.com/voltento/ws_tool/pkg/command"
	"github.com/voltento/ws_tool/pkg/web_socket_client"
	"os"
)

func InitMgrHandlers(mgr *HandlerManager) {
	var er error
	addHandler := func(command string, h Handler) {
		er = mgr.AddHandler(command, h)

		if er != nil {
			fmt.Printf("Error occured on read message. Reason: `%v`", er)
			os.Exit(2)
		}
	}

	addHandler("<", readMessage)
	addHandler(">", writeMessage)
	addHandler("exit", exit)
}

func readMessage(ws *web_socket_client.WebSocket, _ command.Command) error {
	msg, er := ws.ReadOneMessage()
	if er != nil {
		return errors.New(fmt.Sprintf("Error occured on in `readMessage` handler. Reason: `%v`", er))
	}
	fmt.Printf("< %v\n", msg)
	return nil
}

func writeMessage(ws *web_socket_client.WebSocket, cmd command.Command) error {
	er := ws.SendMessage(cmd.Args)
	if er != nil {
		return errors.New(fmt.Sprintf("Error occured on in `writeMessage` handler. Reason: `%v`", er))
	}
	fmt.Printf("> %v\n", cmd.Args)
	return nil
}

func exit(ws *web_socket_client.WebSocket, cmd command.Command) error {
	_ = ws
	_ = cmd
	fmt.Printf("Exit command was called.")
	os.Exit(0)
	return nil
}
