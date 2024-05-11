package rpc

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"strconv"
)

type EncodeExample struct {
	Testing bool
}

func EncodeMessage(message any) string{

	content,err:=json.Marshal(message)
	if err != nil{
		panic(err)
	}
	return fmt.Sprintf("Content-Length: %d\r\n\r\n%s",len(content),content)
}
type BaseMessage struct{
	Method string `"json:method"`
}
func DecodeMessage(msg []byte) (string,[]byte,error){
	header,content,found:=bytes.Cut(msg,[]byte{'\r','\n','\r','\n'})
	if !found{
	return "",nil,errors.New("did not find separator")
	}

	// Content-Length: <number>
	contentLenghtByte := header[len("Content-Length: "):]
	contentLength,err := strconv.Atoi(string(contentLenghtByte))
	if err!=nil{
		return "",nil, err
	}
	//TODO we will get to this
	var baseMessage BaseMessage
	if err:=json.Unmarshal(content[:contentLength],&baseMessage); err !=nil{
		return "",nil, err
	}

	return baseMessage.Method,content[:contentLength],nil
}