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
	"os/exec"
	"strings"
)

type responseMsgType struct {
	Id      uint32 `json:"transactionId"`
	Data    interface{} `json:"data"`
	Server  string `json:"server"`
	Created string `json:"created"`
}

type requestMsgType struct {
	Message string `json:"message"`
}

// wait seconds before send response
const sleepSecs = 10
const serverName = `WebSocket Go Server v.1.0`
const serverAddr = `:1234`

var TransactionId uint32 = 0

func main() {
	http.Handle("/api", websocket.Handler(Echo))
	http.Handle("/cmd", websocket.Handler(EchoCmd))

	log.Printf(`Server listening at: "%s"%s`, serverAddr, "\n")

	if err := http.ListenAndServe(serverAddr, nil); err != nil {
		log.Fatal("ListenAndServe:", err)
	}
}

func ParseRequestMsg(ws *websocket.Conn) (requestMsg requestMsgType, err error) {
	var request string

	if err = websocket.Message.Receive(ws, &request); err != nil {
		return requestMsg, errors.New(fmt.Sprint("Can't receive: ", err))
	}

	fmt.Println("Received from client: " + request)

	if err = json.Unmarshal([]byte(request), &requestMsg); err != nil {
		return requestMsg, errors.New(fmt.Sprintf("Failed to decode %s with message: %s\n", request, err))
	}

	fmt.Println(`Parsed to`, requestMsg)

	return requestMsg, nil
}

func GenerateResponseMsg(requestMsg *requestMsgType) (responseMsg responseMsgType, err error) {
	atomic.AddUint32(&TransactionId, 1)
	responseMsg.Id = TransactionId
	responseMsg.Data = requestMsg
	responseMsg.Server = serverName
	responseMsg.Created = time.Now().Format(time.RFC3339)

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

func EchoCmd(ws *websocket.Conn) {
	for {
		requestMsg, err := ParseCommand(ws)

		if err != nil {
			log.Println(err)
			continue
		}

		responseMsg, err := GenerateResponseCmd(&requestMsg)

		if err != nil {
			log.Println(`Command error:`, err)
			//continue // send response with error to client
		}

		msg, err := json.Marshal(responseMsg)
		if err != nil {
			log.Println(`Encoding error:`, err)
			continue
		}

		if sleepSecs > 0 {
			//time.Sleep(sleepSecs * time.Second)
		}

		if err = websocket.Message.Send(ws, string(msg)); err != nil {
			log.Println("Can't send:", err)
		}
	}
}

type cmdType struct {
	Command string `json:"command"`
}

func ParseCommand(ws *websocket.Conn) (cmd cmdType, err error) {
	var request string

	if err = websocket.Message.Receive(ws, &request); err != nil {
		return cmd, errors.New(fmt.Sprint("Can't receive: ", err))
	}

	fmt.Println("Received from client: " + request)

	if err = json.Unmarshal([]byte(request), &cmd); err != nil {
		return cmd, errors.New(fmt.Sprintf("Failed to decode \"%s\" with message: %s\n", request, err))
	}

	fmt.Println(`Parsed to`, cmd)

	return cmd, nil
}

type cmdResponseType struct {
	Id      uint32 `json:"transactionId"`
	Command cmdType `json:"request"`
	Output  string `json:"output"`
	Created string `json:"created"`
	Server  string `json:"server"`
	Error   string `json:"error"`
}

func GenerateResponseCmd(cmd *cmdType) (cmdResponse cmdResponseType, err error) {
	atomic.AddUint32(&TransactionId, 1)
	cmdResponse.Id = TransactionId
	cmdResponse.Server = serverName
	cmdResponse.Created = time.Now().Format(time.RFC3339)
	cmdResponse.Command = *cmd

	arrChunks := strings.Split(cmd.Command, ` `)
	cmdObj := exec.Command(arrChunks[0])
	if len(arrChunks) > 1 {
		cmdObj.Args = arrChunks
		log.Println(`Adding chunks:`, cmdObj.Args)
	}
	log.Println(`Command:`, cmdResponse)
	output, err := cmdObj.Output()

	log.Println(`Command executed:`, cmdObj)
	if err != nil {
		cmdResponse.Error = err.Error()
		return cmdResponse, err
	}

	cmdResponse.Output = string(output)

	msg, err := json.Marshal(cmdResponse)
	if err != nil {
		cmdResponse.Error = err.Error()
		return cmdResponse, errors.New(fmt.Sprintf("Failed to encode \"%v\" with message: %s\n", cmdResponse, err))
	}

	fmt.Println("Sending to client: " + string(msg))

	return cmdResponse, nil
}
