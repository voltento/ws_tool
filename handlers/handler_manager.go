package handlers

import (
	"errors"
	"fmt"
	"github.com/voltento/WsTool/command"
	"github.com/voltento/WsTool/web_socket_client"
)

type Handler func(*web_socket_client.WebSocket, command.Command) error

type HandlerManager struct {
	handlers map[string]Handler
	ws       *web_socket_client.WebSocket
}

func (mgr *HandlerManager) AddHandler(command string, h Handler) error {
	if _, exists := mgr.handlers[command]; exists {
		return errors.New(fmt.Sprintf("Handler for the command with name `%v` already exists", command))
	}

	mgr.handlers[command] = h
	return nil
}

func CreateHandlerManager(ws *web_socket_client.WebSocket) HandlerManager {
	mgr := HandlerManager{make(map[string]Handler), ws}
	InitMgrHandlers(&mgr)
	return mgr
}

func (mgr *HandlerManager) Handle(command *command.Command) error {
	h, ok := mgr.handlers[command.Name]
	if !ok {
		return errors.New(fmt.Sprintf("Can't find handler for the command `%v`", command.Name))
	}

	return h(mgr.ws, *command)
}
