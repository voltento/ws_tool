**DOC**

This is simple WebSocket client.<br/>
It can connect to a server by WebSocket protocol, send and get messages.<br/>
Also, here is header customization.

**Install**

Set `$GOPATH` if it wasn't set before. F.e. `export GOPATH=$HOME/go`  
For instalation run command: `go get -u github.com/voltento/WsTool`<br/>
Binary will be installed to the `$GOPATH/bin` folder<br/>
You must have installed golang and set the `$GOPATH` f.e. `export GOPATH=$HOME/go`

Check for more information about golang instal and configuration
- https://golang.org/doc/install 
- https://github.com/golang/go/wiki/SettingGOPATH

**Dependenciec**
- Golang 1.11.2 or above
- Set `$GOPATH` before call installation script

**Usage**

Usage example: `./WsTool ws://localhost:3000/echo/websocket -H "host:ws" -C "userId=1"`

Use flag `--help` for more information
