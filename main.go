package main

import (
	"./serverDir"
	"fmt"
)
func main(){
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
		server.Write([]byte("Beer "+content))
	}
}


