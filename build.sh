go env -w GO111MODULE=on
go env -w GOPROXY=https://goproxy.cn,direct

# amd64 windows
CGO_ENABLED=0
GOOS=windows
GOARCH=amd64
go build -o rollback-address-finder-windows-amd64.exe .

# amd64 mac
CGO_ENABLED=0
GOOS=darwin
GOARCH=amd64
go build -o rollback-address-finder-darwin-amd64 .

# arm64 mac
CGO_ENABLED=0
GOOS=darwin
GOARCH=arm64
go build -o rollback-address-finder-darwin-arm64 .

# amd64 linux
CGO_ENABLED=0
GOOS=linux
GOARCH=amd64
go build -o rollback-address-finder-linux-amd64 .

# arm64 linux
CGO_ENABLED=0
GOOS=linux
GOARCH=arm64
go build -o rollback-address-finder-linux-arm64 .
