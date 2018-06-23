package serverDir

import (
	"net"
	"strconv"
	"fmt"
	"bufio"
	"syscall"
)

type server struct{
	port int
	connection net.Conn
}

type printable interface {
	ToString() string
}


func Start(port int) *server {
	result,_ := net.Listen("tcp",":"+strconv.Itoa(port))
	connection,_ := result.Accept()
	fmt.Println("Accepted")
	return &server{port,connection}

}


//serves as a writer stream of bytes - can be used to make raw inputstream but following c style getchar()
func (serv *server) Read_deprecated() []byte {
	buffer:=make([]byte,1000)
	i:=0
	reader :=bufio.NewReader(serv.connection)
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
func (serv *server) Read() []byte{
	br:=bufio.NewReader(serv.connection)
	result,_:=br.ReadString('\n')
	return []byte(result)[:len(result)-1]
}

func (serv *server) Write(data []byte){
	new_data :=serv.connection.LocalAddr().String()+" says "+string(data)+"\n"
	write:=bufio.NewWriter(serv.connection)
	write.Write([]byte(new_data))
	write.Flush()

}