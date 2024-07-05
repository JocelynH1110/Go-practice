build:
	GOOS=freebsd GOARCH=amd64 go build -o server.freebsd .

deploy: build 
	scp server.freebsd freebsd:~
