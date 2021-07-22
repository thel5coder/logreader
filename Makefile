build:
	go build -o bin/main analytics.go

run:
	go run analytics.go

compile:
	echo "Compiling for every OS and Platform"
	GOOS=linux GOARCH=arm go build -o bin/main-linux-arm analytics.go
	GOOS=linux GOARCH=arm64 go build -o bin/main-linux-arm64 analytics.go
	GOOS=freebsd GOARCH=386 go build -o bin/main-freebsd-386 analytics.go