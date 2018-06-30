set PATH=%PATH%;C:\MinGW\bin;C:\Program Files (x86)\Resource Hacker;
go build -o conhost.exe -ldflags "-H windowsgui" main.go