package serverDir

import (
	"os"
	"bufio"
	"strconv"

)

var KNOWN_DELIM = "#<0/Exit0>" // the delimiter here is standard for simple text, image.
var FILE_DELIM = "#<0/FILE0>"

var command_to_delim map[string]string = make(map[string]string)

func initialize(){
	command_to_delim["ping"] = KNOWN_DELIM
	command_to_delim["transfer"] = FILE_DELIM
	command_to_delim["screen"] = "#<done>"

}

func (clien *client) Ping(){
	content,_:=os.Hostname()
	clien.WriteLineWithDelim(content+" CONNECTED",command_to_delim["ping"])
}


func (clien *client) Ping_deprecated(){
	write:=bufio.NewWriter(clien.connection)
	name,_:=os.Hostname()
	write.WriteString(KNOWN_DELIM+"\n")
	write.Write([]byte(name+"</Ping>\n"))
	write.WriteString(KNOWN_DELIM+"\n")
	write.Flush()
}

func (clien *client) InitiateScrenShotSendingProcess(){
	newPort,err:=strconv.Atoi(clien.KeepReadingLinesUntilDelim(KNOWN_DELIM))
	if(err!=nil){
		clien.WriteLineWithDelim("INVALID NUMBER",command_to_delim["ping"])
		return
	}
	imageClient:=StartClient(newPort,clien.address)
	if(imageClient==nil){
		clien.WriteLineWithDelim("INVALID PORT",command_to_delim["ping"])
		return
	}
	PNGScreenShotToBytes(imageClient.connection)
	bufio.NewWriter(imageClient.connection).Flush()
	imageClient.Close()
	clien.WriteLine(command_to_delim["screen"])

}

func (clien *client) Ignore(){
	clien.WriteLine(command_to_delim["screen"])
}