@ECHO OFF
echo Make sure you're running this as an administrator. Press any key if you are :)
PAUSE
cd /D "%~dp0"

echo Building for Linux amd64...
set GOARCH=amd64
set GOOS=linux
go build -o bin/simpledns_linux_amd64

echo Building for Linux arm64...
set GOARCH=arm64
set GOOS=linux
go build -o bin/simpledns_linux_arm64

echo Building for Windows amd64...
set GOARCH=amd64
set GOOS=windows
go build -o bin/simpledns_windows_amd64.exe

PAUSE