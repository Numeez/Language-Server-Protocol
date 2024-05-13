package main

import (
	"bufio"
	"log"
	"os"

	"github.com/Numeez/Language-Server-Protocol/rpc"
)
	

func main(){
logger:= getLogger("/Users/numeezbaloch17/Documents/Language-Server-Protocol/log.txt")
logger.Println("Hey I started !")
scanner := bufio.NewScanner(os.Stdin)
scanner.Split(rpc.Split)
for scanner.Scan(){
	msg:=scanner.Text()
	handleMessage(logger,msg)
}
}

func handleMessage(logger *log.Logger,message any){
	logger.Println(message)
}

func getLogger(filename string) *log.Logger{
	logFile,err:= os.OpenFile(filename,os.O_CREATE|os.O_TRUNC|os.O_WRONLY,0666)
	if err !=nil{
		panic("Hey you did not give me the right file")
	}

	return log.New(logFile,"LSP : ",log.Ldate|log.Lshortfile|log.Ltime)

}
