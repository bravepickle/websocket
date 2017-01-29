package main

import (
	"golang.org/x/net/websocket"
	"fmt"
	"log"
	"net/http"
	"encoding/json"
	"errors"
	"sync/atomic"
	"time"
)

var TransactionId uint32

// wait seconds before send response
const sleepSecs = 10
const serverName = `WebSocket Go Server v.1.0`

func main() {
	http.Handle("/api", websocket.Handler(Echo))

	addr := ":1234"
	TransactionId = 0

	log.Printf(`Server listening at: "%s"%s`, addr, "\n")

	if err := http.ListenAndServe(addr, nil); err != nil {
		log.Fatal("ListenAndServe:", err)
	}
}

type requestMsgType struct {
	Message string `json:"message"`
}

func ParseRequestMsg(ws *websocket.Conn) (requestMsg requestMsgType, err error) {
	var reply string

	if err = websocket.Message.Receive(ws, &reply); err != nil {
		return requestMsg, errors.New(fmt.Sprint("Can't receive:", err))
	}

	fmt.Println("Received back from client: " + reply)

	if err = json.Unmarshal([]byte(reply), &requestMsg); err != nil {
		return requestMsg, errors.New(fmt.Sprintf("Failed to decode %s with message: %s\n", reply, err))
	}

	fmt.Println(`Parsed to`, requestMsg)

	return requestMsg, nil
}

type responseMsgType struct {
	Id     uint32 `json:"transactionId"`
	Data   interface{} `json:"data"`
	Server string `json:"server"`
	Created time.Time `json:"created"`
}

func GenerateResponseMsg(requestMsg *requestMsgType) (responseMsg responseMsgType, err error) {
	atomic.AddUint32(&TransactionId, 1)
	responseMsg.Id = TransactionId
	responseMsg.Data = requestMsg
	responseMsg.Server = serverName

	msg, err := json.Marshal(responseMsg)
	if err != nil {
		return responseMsg, errors.New(fmt.Sprintf("Failed to encode %v with message: %s\n", responseMsg, err.Error()))
	}

	fmt.Println("Sending to client: " + string(msg))

	return responseMsg, nil
}

func Echo(ws *websocket.Conn) {
	for {
		requestMsg, err := ParseRequestMsg(ws)

		if err != nil {
			log.Println(err)
			continue
		}

		responseMsg, err := GenerateResponseMsg(&requestMsg)

		if err != nil {
			log.Println(err)
			continue
		}

		msg, err := json.Marshal(responseMsg)
		if err != nil {
			log.Println(err)
			continue
		}

		if sleepSecs > 0 {
			time.Sleep(sleepSecs * time.Second)
		}

		if err = websocket.Message.Send(ws, string(msg)); err != nil {
			log.Println("Can't send:", err)
		}
	}
}