package command

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

type Command struct {
	Name string
	Args string
}

func CreateReaderFromFile(file string) chan *Command {
	reader := make(chan *Command, 100)
	data, er := ioutil.ReadFile(file)
	if er != nil {
		fmt.Printf("Can't read file %v. Reason: %v", file, er.Error())
		os.Exit(1)
	}

	go func() {
		rawCommands := strings.Split(string(data), "\n")
		for _, rawCommand := range rawCommands {
			var c Command
			i := strings.Index(rawCommand, " ")
			if i == -1 || i == len(rawCommand) {
				c.Name = rawCommand
			} else {
				c.Name = rawCommand[:i]
				c.Args = rawCommand[i+1:]
			}

			reader <- &c
		}
		close(reader)
	}()

	return reader
}
