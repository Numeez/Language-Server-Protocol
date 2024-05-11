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

func DecodeMessage(msg []byte) (int,error){
	header,content,found:=bytes.Cut(msg,[]byte{'\r','\n','\r','\n'})
	if !found{
	return 0,errors.New(`Did not find separator`)
	}

	// Content-Length: <number>
	contentLenghtByte := header[len("Content-Length: "):]
	contentLength,err := strconv.Atoi(string(contentLenghtByte))
	if err!=nil{
		return 0, err
	}
	//TODO we will get to this
	_ = content

	return contentLength,nil
}