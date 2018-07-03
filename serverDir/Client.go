package serverDir

import (
	"net"
	"strconv"
	"bufio"
	"syscall"

	"fmt"
	"time"
)




type Client struct{
	port int
	address string
	connection net.Conn
	br *bufio.Reader
	pw *bufio.Writer
	DC bool
}


func (clien *Client) SetReadDeadline(){

	clien.connection.(*net.TCPConn).SetKeepAlive(true)
	clien.connection.(*net.TCPConn).SetReadDeadline(time.Now().Add(time.Second*15))
}

func StartClient(port int, address string) *Client{
	conn,error := net.DialTimeout("tcp", address+":"+strconv.Itoa(port),time.Second*2)
	if(error!=nil){
		panic("DC")
	}
	initialize()

	return &Client{port,address,conn, bufio.NewReader(conn), bufio.NewWriter(conn),false}
}

func (ct *Client) Close(){
	ct.connection.Close()
}
//Use of bufio.readline() because java puts \r\n which has carriage return that fucks comparison
//so to input compatibility instead of writing raw bytes in java which is better but more work...
func (clien *Client) ReadLine_new() string{
	content,_,err := clien.br.ReadLine()
	if(err!=nil){
		panic("DC")
	}

	return string(content)

}

func (clien *Client) KeepReadingLinesUntilDelim(delim string) string{

	var total string = ""

	for newLine:=clien.ReadLine_new();newLine!=delim && !clien.DC;newLine=clien.ReadLine_new(){
		total+=newLine
	}

	return total
}

//a dangerous method
func (clien *Client) Write(data []byte){

	clien.pw.Write(data)
	clien.pw.Flush();
}

func (clien *Client) WriteLine(data string) {

	write:=clien.pw
	_,err:=write.WriteString(data+"\n")
	if(err!=nil){
		fmt.Println(err)
		panic("DC")
	}
	err=write.Flush()
	if(err!=nil){
		panic("DC")
	}


}

func (clien *Client) WriteLineWithDelim(data string, delim string){

	clien.WriteLine(delim+"\n"+data+"\n"+delim)

}


























func (clien *Client) Read_deprecated() []byte {
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


















