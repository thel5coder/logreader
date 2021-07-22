build:
	go build -o analytics analytics.go

run:
	go run analytics.go

compile:
	echo "Compiling for every OS and Platform"
	GOOS=linux GOARCH=arm go build -o bin/analytics-arm analytics.go
	GOOS=linux GOARCH=arm64 go build -o bin/analytics-arm64 analytics.go
	GOOS=linux GOARCH=amd64 go build -o bin/analytics analytics.go
	GOOS=freebsd GOARCH=386 go build -o bin/analytics-freebsd-386 analytics.go