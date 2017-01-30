# websocket
Test scripts for web sockets implementation in GOLANG

## Features
* server-side implementation in Go that responds with requested date and extra set of fields
* client-side (web browser) implemenation of websockets that connect to server and send JSON requests
* reconnect logic after disconnection
* client-side (web browser) command line interface emulator via web-browser

## TODO
* support of WSS (over HTTPS) protocol
* proper handling of various specifics of web browser clients handling web sockets
* add tests


## Example

1. Run websockets server
```sh
go build
./websocket
```

2. Run html tests
```sh
cd web/
python -m SimpleHTTPServer
```
 
3. Open links to [send JSON](http://localhost:8000/test.html) or [send CLI commands](http://localhost:8000/test_cmd.html) via websocket clients


### Command Line Emulator
![CMD Websockets Emulator](http://take.ms/3Y9vQ)

### JSON Requester
![JSON Websockets Requester](http://take.ms/c7R0v)