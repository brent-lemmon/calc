BINARY_NAME = calc

build:
	GOOS=linux GOARCH=amd64 go build -o bin/${BINARY_NAME}-linux calc.go
	GOOS=windows GOARCH=amd64 go build -o bin/${BINARY_NAME}-windows calc.go
	GOOS=darwin GOARCH=amd64 go build -o bin/${BINARY_NAME}-darwin calc.go
