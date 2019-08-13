**DOC**

This is simple WebSocket client.<br/>
It can connect to a server by WebSocket protocol, send and get messages.<br/>
Also, here is header customization.

**Install**

`go get -u github.com/voltento/WsTool`<br/> <br/>
The command will install the binary `wstool` to the `$GOPATH/bin` folder<br/><br/>
You must have installed golang and set the `$GOPATH` f.e. `export GOPATH=$HOME/go` Set `$GOPATH` if it wasn't set before. F.e. `export GOPATH=$HOME/go`

Check for more information about golang install and configuration
- https://golang.org/doc/install 
- https://github.com/golang/go/wiki/SettingGOPATH

**Remove**

`go clean -i github.com/voltento/WsTool`

**Dependenciec**
- Golang 1.11.2 or above
- Set `$GOPATH` before call installation script

**Usage**

Usage example: `./WsTool ws://localhost:3000/echo/websocket -H "host:ws" -C "userId=1"`

Use flag `--help` for more information

**Supported commands**

You can specify file with commands 
Usage example: `./WsTool ws://localhost:3000/echo/websocket commands -H "host:ws" -C "userId=1"`
After processing all commands from file wstool will still listening 

Supporting commands:

- `<` - will read any message from ws connection and print it on screen
- `> msg` - will send msg to ws connection and print it on screen

Commands file example:

```
<
> {"message": "foo"}
<
```


**Testing**

`cd $GOPATH/github.com/voltento/WsTool`<br/>
`go test ./...`
