package serverDir

import (
	"net"
	"strconv"
	"bufio"
	"syscall"

)



type client struct{
	port int
	address string
	connection net.Conn
	br *bufio.Reader
	pw *bufio.Writer
}

var br *bufio.Reader


func StartClient(port int, address string) *client{
	conn,error := net.Dial("tcp", address+":"+strconv.Itoa(port))
	if(error!=nil){
		return nil
	}
	initialize()

	return &client{port,address,conn, bufio.NewReader(conn), bufio.NewWriter(conn)}
}

func (ct *client) Close(){
	ct.connection.Close()
}
//Use of bufio.readline() because java puts \r\n which has carriage return that fucks comparison
//so to input compatibility instead of writing raw bytes in java which is better but more work...
func (clien *client) ReadLine_new() string{
	content,_,err := clien.br.ReadLine()
	if(err!=nil){
		syscall.Exit(0)
	}

	return string(content)

}

func (clien *client) KeepReadingLinesUntilDelim(delim string) string{
	var total string = ""

	for newLine:=clien.ReadLine_new();newLine!=delim;newLine=clien.ReadLine_new(){
		total+=newLine


	}
	return total
}

//a dangerous method
func (clien *client) Write(data []byte){
	clien.pw.Write(data)
	clien.pw.Flush();
}

func (clien *client) WriteLine(data string){
	write:=clien.pw
	write.WriteString(data+"\n")
	write.Flush()

}

func (clien *client) WriteLineWithDelim(data string, delim string){

	clien.WriteLine(delim+"\n"+data+"\n"+delim)

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
func (clien *client) ReadLine_deprecated() string{
	if(br==nil){
		br=bufio.NewReader(clien.connection)
	}
	result,err:=br.ReadString('\n')
	if(err!=nil){
		syscall.Exit(0)
	}
	return result[:len(result)-1]

}
















