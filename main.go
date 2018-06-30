package main

import (
	"./serverDir"
)
func main(){
	//start()
	start_client()

}


var(
	specialFuncsAccess map[string]func() = make(map[string]func())
)

func start_client(){

	client:=serverDir.StartClient(2334, "LAPTOP-ARK617N3")

	specialFuncsAccess["ping"] = client.Ping
	specialFuncsAccess["screen"] = client.InitiateScrenShotSendingProcess

	for{
		content:=string(client.KeepReadingLinesUntilDelim(serverDir.KNOWN_DELIM))
		if(content=="exit"){
			break;
		}
		if val,exists:=specialFuncsAccess[content];exists{
			val()
		}else{
			client.Ignore()
		}

	}
}
