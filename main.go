package main

import (
	"./serverDir"
	"fmt"
	"os"
	"bufio"
)
func main(){
	//start()


}

func testScreenShot(){

	file,err:=os.Create("./sc.png")
	if(err!=nil){
		fmt.Println("failed")
	}
	b :=bufio.NewWriter(file)
	serverDir.PNGScreenShotToBytes(b)
	b.Flush()


	file.Close()

}

func start(){
	server:=serverDir.Start(2334)
	for{
		x:=server.Read()
		content :=string(x)

		if(content=="done"){
			fmt.Println("Exitted")
			server.Write([]byte("<exit>\n"))
			break
		}
		fmt.Println("received "+content)
		//server.WriteDelim([]byte("Beer "+content),"\n")
		server.Write([]byte("rules"))
	}
}

var(
	specialFuncsAccess map[string]func()
)

func start_client(){

	client:=serverDir.StartClient(2334, "localhost")
	specialFuncsAccess["ping"] = client.Ping

	for{
		content:=string(client.ReadUntilDelimLine(serverDir.KNOWN_DELIM))
		if val,exists:=specialFuncsAccess[content]; exists{
			val()
		}

	}
}
