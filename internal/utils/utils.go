package utils

import (
	"errors"
	"fmt"
	"net/http"
	"os"
	"strings"
)

func ParseHeaderKeyValue(s string) (string, string, error) {
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

type Address string

func printHelp() {
	fmt.Println("Args: url [commands_file] [-H \"HeaderName: Header Value\"] [-C \"CookieName: Cookie Value\"]")
}

func ProcessError(er error) {
	if er != nil {
		fmt.Printf("Error occured. Error: %v", er.Error())
		printHelp()
		os.Exit(1)
	}
}

const (
	header = iota
	cookie
	undefined
)

func ParseArgs() (Address, http.Header, string) {
	var commandsFile string
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
			argType = cookie
		} else if argIndex == 2 {
			commandsFile = os.Args[argIndex]
			argIndex += 1
			continue
		}

		if argType == undefined {
			er := errors.New(fmt.Sprintf("Can't parse arg value. Value: %s\n", os.Args[argIndex]))
			ProcessError(er)
		}

		argIndex += 1
		if argIndex == len(os.Args) {
			er := errors.New(fmt.Sprintf("Value for header flag wasn't provided. Value: %s\n", os.Args[argIndex]))
			ProcessError(er)
		}

		key, value, er := ParseHeaderKeyValue(os.Args[argIndex])
		ProcessError(er)

		if argType == header {
			headers.Add(key, value)
		} else if argType == cookie {
			cookie := http.Cookie{}
			cookie.Name = key
			cookie.Value = value
			headers.Add("Cookie", fmt.Sprintf("%v=%v;", key, value))
		}
		argIndex += 1
	}

	return Address(os.Args[1]), headers, commandsFile
}
