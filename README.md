**DOC**

This is simple WebSocket client.<br/>
It can connect to a server by WebSocket protocol, send and get messages and perform scenarios.<br/>
Also, here is header customization.

**Install**

`GO111MODULE=on go get github.com/voltento/ws_tool/cmd/ws_tool`<br/> <br/>
The command will install the binary `ws_tool` to the `$GOPATH/bin` folder<br/><br/>
You must have installed golang and set the `$GOPATH` f.e. `export GOPATH=$HOME/go` Set `$GOPATH` if it wasn't set before. F.e. `export GOPATH=$HOME/go`

Check for more information about golang install and configuration
- https://golang.org/doc/install 
- https://github.com/golang/go/wiki/SettingGOPATH

**Remove**

`go clean -i github.com/voltento/ws_tool`

**Dependencies**
- Golang 1.11.2 or above
- Set `$GOPATH` before call installation script

**Usage**

Usage example: `./ws_tool ws://localhost:3000/echo/websocket -H "host:ws" -C "userId=1"`

Use flag `--help` for more information

**Supported commands**

You can specify file with scenario 
Usage example: `./ws_tool ws://localhost:3000/echo/websocket commands.txt -H "host:ws" -C "userId=1"`
After processing all commands from a scenario file ws_tool will still listening if exit command wasn't call explicitly.

Supporting commands:

- `<` - will read any message from ws connection and print it on screen
- `> msg` - will send msg to ws connection and print it on screen
- `exit` - will terminate ws_tool

Commands file example:

```
<
> {"message": "foo"}
<
```


**Testing**

`cd $GOPATH/github.com/voltento/ws_tool`<br/>
`go test ./...`
