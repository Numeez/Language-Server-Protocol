package main

import (
	"bufio"
	"encoding/json"
	"log"
	"os"

	"github.com/Numeez/Language-Server-Protocol/lsp"
	"github.com/Numeez/Language-Server-Protocol/rpc"
)
	

func main(){
logger:= getLogger("/Users/numeezbaloch17/Documents/Language-Server-Protocol/log.txt")
logger.Println("Hey I started !")
scanner := bufio.NewScanner(os.Stdin)
scanner.Split(rpc.Split)
for scanner.Scan(){
	msg:=scanner.Bytes()
	method,contents,err:=rpc.DecodeMessage(msg)
	if err !=nil{
		logger.Printf("Got an error :%s",err)
	}

	handleMessage(logger,method,contents)
}
}

func handleMessage(logger *log.Logger,method string, content []byte){
	logger.Printf("Received a message with method : %s",method )

	switch method {
	case "initialize":
		var request lsp.InitializeRequest
		if err := json.Unmarshal(content,&request); err!=nil{
			logger.Printf("We could not parse this : %s",err)
		}
		logger.Printf("Connected to %s %s",request.Params.ClientInfo.Name,request.Params.ClientInfo.Version)
		msg:= lsp.NewInitializeResponse(request.ID)
		reply:= rpc.EncodeMessage(msg)
		writer:=os.Stdout
		writer.Write([]byte(reply))

		logger.Print("Reply sent ........")
	}
}

func getLogger(filename string) *log.Logger{
	logFile,err:= os.OpenFile(filename,os.O_CREATE|os.O_TRUNC|os.O_WRONLY,0666)
	if err !=nil{
		panic("Hey you did not give me the right file")
	}

	return log.New(logFile,"LSP : ",log.Ldate|log.Lshortfile|log.Ltime)

}
