build:
	go build -o $(GOPATH)/bin/gitraw main.go

compile:
	echo "Compiling for every OS and Platform"
	GOOS=linux GOARCH=arm go build -o $(GOPATH)/bin/gitraw-linux-arm main.go
	GOOS=linux GOARCH=arm64 go build -o $(GOPATH)/bin/gitraw-linux-arm64 main.go
	GOOS=freebsd GOARCH=386 go build -o $(GOPATH)/bin/gitraw-freebsd-386 main.go