package main

import (
	"bufio"
	"encoding/json"

	"log"
	"os"

	"github.com/Numeez/Language-Server-Protocol/analysis"
	"github.com/Numeez/Language-Server-Protocol/lsp"
	"github.com/Numeez/Language-Server-Protocol/rpc"
)
	

func main(){
logger:= getLogger("/Users/numeezbaloch17/Documents/Language-Server-Protocol/log.txt")
logger.Println("Hey I started !")
scanner := bufio.NewScanner(os.Stdin)
scanner.Split(rpc.Split)
state := analysis.NewState()
for scanner.Scan(){
	msg:=scanner.Bytes()
	method,contents,err:=rpc.DecodeMessage(msg)
	if err !=nil{
		logger.Printf("Got an error :%s",err)
	}

	handleMessage(logger,state,method,contents)
}
}

func handleMessage(logger *log.Logger,state analysis.State,method string, content []byte){
	logger.Printf("Received a message with method : %s",method)

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
	case "textDocument/didOpen":
		var request lsp.DidOpenTextDocumentNotification
		if err := json.Unmarshal(content,&request); err!=nil{
			logger.Printf("did not able to open : %s",err)
		}
		logger.Printf("Opened : %s ",request.Params.TextDocument.URI)
		state.OpenDocument(request.Params.TextDocument.URI,request.Params.TextDocument.Text)
	case "textDocument/didChange":
		var request lsp.TextDocumentDidChangeNotification
		
		if err := json.Unmarshal(content,&request); err!=nil{
			logger.Printf("document change error  : %s",err)
		}
		logger.Printf("Changed : %s ",request.Params.TextDocument.URI)
		for _,change := range request.Params.ContentChanges{
			state.UpdateDocument(request.Params.TextDocument.URI,change.Text)
		}
		

	}


}

func getLogger(filename string) *log.Logger{
	logFile,err:= os.OpenFile(filename,os.O_CREATE|os.O_TRUNC|os.O_WRONLY,0666)
	if err !=nil{
		panic("Hey you did not give me the right file")
	}

	return log.New(logFile,"LSP : ",log.Ldate|log.Lshortfile|log.Ltime)

}
