set GOARCH=amd64
set GOOS=linux
go env
go build -o %CD%/bin/FullStack.Service.exe
pause
