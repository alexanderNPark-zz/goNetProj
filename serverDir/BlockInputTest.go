package serverDir

import (
	"os"
	"fmt"
	"bufio"
)

func empty() {

}
/*
var(

	user32= syscall.NewLazyDLL("user32.dll")
	proc_blockInput= user32.NewProc("BlockInput")
	SM_CXSCREEN = 0
	SM_CYSCREEN = 1

)

func TestBlockInput(){
	administrative run ONLY
	var status_active int32 = 1
	result,_,err:=proc.Call(uintptr(status_active))
	fmt.Println(result,err)
	time.Sleep(time.Millisecond*10000)
	proc.Call(uintptr(0))

	proc_getSystemMetrics:=user32.NewProc("GetSystemMetrics")
	proc_getDesktopWindow:=user32.NewProc("GetDesktopWindow")
	proc_getDC:=user32.NewProc("GetDC")
	width,_,err:=proc_getSystemMetrics.Call(uintptr(SM_CXSCREEN))//second is an unintptr don't know why it is
	if(err!=nil){
		fmt.Println(err)
	}
	height,_,err:=proc_getSystemMetrics.Call(uintptr(SM_CYSCREEN))
	//width_dereference:=strconv.Itoa(int(width))
	/*system values always return some sort of pointer to some sort of number rather than the actual
	//fmt.Println(width_dereference)
	//fmt.Println(width,height)
	hDesktopWnd_ptr,_,err := proc_getDesktopWindow.Call()
	fmt.Println(hDesktopWnd_ptr,width,height)
	//hDesktopDC,_,err:= proc_getDC.Call(hDesktopWnd_ptr);



	// get the device context of the screen
	HDC hScreenDC = CreateDC("DISPLAY", NULL, NULL, NULL);
	// and a device context to put it in
	HDC hMemoryDC = CreateCompatibleDC(hScreenDC);

	int width = GetDeviceCaps(hScreenDC, HORZRES);
	int height = GetDeviceCaps(hScreenDC, VERTRES);

	// maybe worth checking these are positive values
	HBITMAP hBitmap = CreateCompatibleBitmap(hScreenDC, width, height);

	// get a new bitmap
	HBITMAP hOldBitmap = (HBITMAP) SelectObject(hMemoryDC, hBitmap);

	BitBlt(hMemoryDC, 0, 0, width, height, hScreenDC, 0, 0, SRCCOPY);
	hBitmap = (HBITMAP) SelectObject(hMemoryDC, hOldBitmap);

	// clean up
	DeleteDC(hMemoryDC);
	DeleteDC(hScreenDC);

	// "now your image is held in hBitmap. You can save it or do whatever with it"???????? - > even if I store it in device context
       I KNOW NOTHING ABOUT HOW TO EXTRACT THE BITMAP NOR THE BITMAP HEADER NOR ENCODING NOR EXTRACTION FORMAT



Your question is still not clear, but I assume you mean using SetWindowsHookEx()
to install a system wide hook, and so inject your dll into the address space of all apps.

The reason code will be DLL_PROCESS_ATTACH, since you are attaching to a process.


}
*/

func testScreenShot(){

	file,err:=os.Create("./sc.png")
	if(err!=nil){
		fmt.Println("failed")
	}
	b :=bufio.NewWriter(file)
	PNGScreenShotToBytes(b)
	b.Flush()


	file.Close()

}

func start(){
	server:=Start(2334)
	for{
		x:=server.Read()
		content :=string(x)

		if(content=="done"){
			fmt.Println("Exitted")
			server.Write([]byte("<exit>\n"))
			break
		}
		fmt.Println("received "+content)
		//server.WriteDelim([]byte("Beer "+content),"\n")
		server.Write([]byte("rules"))
	}
}