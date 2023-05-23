compile:
	echo "Compiling for every OS and Platform"
	GOOS=freebsd GOARCH=386 go build -o bin/main-freebsd-386 main.go
	GOOS=linux GOARCH=386 go build -o bin/main-linux-386 main.go
	GOOS=windows GOARCH=386 go build -o bin/main-windows-386 main.go

	GOOS=freebsd GOARCH=amd64 go build -o bin/main-freebsd-amd64 main.go
	GOOS=linux GOARCH=amd64 go build -o bin/main-linux-amd64 main.go
	GOOS=windows GOARCH=amd64 go build -o bin/main-windows-amd64 main.go
	GOOS=darwin GOARCH=amd64 go build -o bin/main-darwin-amd64 main.go

	GOOS=linux GOARCH=arm GOARM=6 go build -o bin/main-linux-armv6 main.go
	GOOS=linux GOARCH=arm GOARM=7 go build -o bin/main-linux-armv7 main.go
	GOOS=linux GOARCH=arm64 go build -o bin/main-linux-arm64 main.go

	GOOS=darwin GOARCH=arm64 go build -o bin/main-darwin-arm64 main.go
