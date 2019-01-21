**DOC**

This is simple WebSocket client.<br/>
It can connect to a server by WebSocket protocol, send and get messages.<br/>
Also, here is header customization.

**Install**
  
For instalation run command: `go get -u github.com/voltento/WsTool`<br/>
Binary will be installed to `%GOPATH/bin`<br/>
You must have installed golang and set `%GOPATH` f.e. `export GOPATH=$HOME/go`

Check for more information about golang instal and configuration
- https://golang.org/doc/install 
- https://github.com/golang/go/wiki/SettingGOPATH

**Dependenciec**
- Golang 1.11.2 or above
- Set `%GOPATH` before call installation script

**Usage**

Usage example: `./WsTool ws://localhost:3000/echo/websocket -H "host:ws"`

Use flag `--help` for mo information