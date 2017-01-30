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
![CMD Websockets Emulator](https://d1ro8r1rbfn3jf.cloudfront.net/ms_64293/zFIEMZQyPDDPER6puRxROGIzqxZ9SB/Testing%2BWebsocket%2BCommand%2BLine%2BEmulator%2B2017-01-30%2B2%2BAM2-08-37.png?Expires=1485821552&Signature=jkVLuZ2Dh1vawwBq30pPquEfFR61NKww-XcIN3nFScG65xpG75yslpcTxb0KbwnUvSu0L4DSXb6jmfHo0~EYQBfHmHIy1RcDegpkxhYI~M2Uo~HnVR2G3I2IhaUHQkX7EPd-c9nQ9reDOA9IhamKcprEcSKGAbHI4Q-hGo-Tt2f9pGoYW00VUHY44AOzL-JGBFrSvSyYsyHt25~-vOGUDOR9JR5~XCn6gqKoA9iTNaxDbiffsEmhXJ-d44YW2emG2NTX3IPTBIRkN9oCJgCEB11b3P3CqlajVUQrv2VqCB6stpb-1Q5J2p0oBwX07eCBqz3o8CSSNaAsubNx5-r1-Q__&Key-Pair-Id=APKAJHEJJBIZWFB73RSA)

### JSON Requester
![JSON Websockets Requester](https://d1ro8r1rbfn3jf.cloudfront.net/ms_64293/9FNrn5kV8xEQkjN07Kd8nr3m1l4zUj/Testing%2BWebsocket%2B2017-01-30%2B2%2BAM2-11-05.png?Expires=1485821604&Signature=Pt9RIEhomzr6OqnICrjwSqDIdzB~ePvAug7oTFxEjhc9klNzlm3XIM0~sSN-1ozVOCByR5R4Q~nFtEolIKiO63mcjP2CuJVd1TSUeDEO5AY6TKjny0q5-MJc1-c6JRkUCblWR3QjXdS6mTs-w1pPJ7PMk9AkIiAK7eFmFYPoGA1P6VaM723oTcQacAnbTQQjLfe1aQ1rBT55cg0gvIQLlUpGMPhz2gHv-4TIYXIiyTV3-QvlbERfSbzLSBFyLWLz2hAjsZ8gPbAuc6UPvgHm9XukZYw7ux6Ga0eHxSsezaSDMJfHXfDyUWDt1iHTq4CBUwbHChIaCu37Yq7lyDH0tQ__&Key-Pair-Id=APKAJHEJJBIZWFB73RSA)