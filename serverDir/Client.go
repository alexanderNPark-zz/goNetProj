package serverDir

import (
	"net"
	"strconv"
	"bufio"
	"syscall"
	"os"
	"fmt"
)

type client struct{
	port int
	address string
	connection net.Conn
	endCommand string
}

func StartClient(port int, address string,endComm string) *client{
	conn,error := net.Dial("tcp", address+":"+strconv.Itoa(port))
	if(error!=nil){
		return nil
	}

	return &client{port,address,conn, endComm}
}

func (ct *client) Close(){
	ct.connection.Close()
}

func (clien *client) Read_deprecated() []byte {
	buffer:=make([]byte,1000)
	i:=0
	reader :=bufio.NewReader(clien.connection)
	for {
		result, err := reader.ReadByte()
		if (err != nil) {
			syscall.Exit(0)
		}
		if(result==10){
			break;
		}
		if(i>cap(buffer)){
			new_buffer:=make([]byte,int(float64(i)*1.5))
			copy(new_buffer,buffer)
			buffer = new_buffer
		}
		buffer[i]=result
		i+=1
	}
	return buffer[:i]

}

//functions as a readline using buffer delimter of \n
func (clien *client) ReadLine() string{
	br:=bufio.NewReader(clien.connection)
	result,err:=br.ReadString('\n')
	if(err!=nil){
		syscall.Exit(0)
	}
	return result[:len(result)-1]

}

func (clien *client) ReadUntilDelimLine(delim string) string{
	var total string
	for newLine:=clien.ReadLine();newLine!=delim;newLine=clien.ReadLine(){
		total+=newLine
	}
	if(clien.endCommand!="" && total==clien.endCommand){
		fmt.Println("quit")
		syscall.Exit(0)
	}
	return total
}


func (clien *client) WriteLine(data string){

	write:=bufio.NewWriter(clien.connection)
	write.WriteString(data+"\n")
	write.Flush()

}

func (clien *client) WriteLineWithDelim(data string){

	clien.WriteLine(data+"\n"+KNOWN_DELIM)

}

var KNOWN_DELIM = "<0/Exit0>"
var in_use = make(chan string)

func (clien *client) InitiateScrenShotSendingProcess(){
	newPort,err:=strconv.Atoi(clien.ReadUntilDelimLine(KNOWN_DELIM))
	if(err!=nil){
		fmt.Println("Invalid Port")
	}
	imageClient:=StartClient(newPort,clien.address,"")
	PNGScreenShotToBytes(imageClient.connection)
	bufio.NewWriter(imageClient.connection).Flush()
	imageClient.Close()


}

//write no delimiter, just raw data
func (clien *client) Write(data []byte){
	osstream:=bufio.NewWriter(clien.connection)
	osstream.Write(data)
	osstream.Flush()
}

func (clien *client) Ping(){
	write:=bufio.NewWriter(clien.connection)
	name,_:=os.Hostname()
	write.Write([]byte(name+"</Ping>\n"))
	write.WriteString(KNOWN_DELIM+"\n")
	write.Flush()
}


