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