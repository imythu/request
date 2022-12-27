go env -w GO111MODULE=on
go env -w GOPROXY=https://goproxy.cn,direct
@echo off
for %%I in (.) do set CurrDirName=%%~nxI
set name=%CurrDirName%

rem amd64 windows
SET CGO_ENABLED=0
SET GOOS=windows
SET GOARCH=amd64
go build -o %name%-windows-amd64.exe .

rem amd64 mac
SET CGO_ENABLED=0
SET GOOS=darwin
SET GOARCH=amd64
go build -o %name%-darwin-amd64 .

rem arm64 mac
SET CGO_ENABLED=0
SET GOOS=darwin
SET GOARCH=arm64
go build -o %name%-darwin-arm64 .

rem amd64 linux
SET CGO_ENABLED=0
SET GOOS=linux
SET GOARCH=amd64
go build -o %name%-linux-amd64 .

rem arm64 linux
SET CGO_ENABLED=0
SET GOOS=linux
SET GOARCH=arm64
go build -o %name%-linux-arm64 .
