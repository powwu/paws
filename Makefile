BINARY_NAME=paws

build:
	 GOARCH=amd64 GOOS=darwin go build -o ${BINARY_NAME}-darwin main.go classroom.go database.go priority.go table.go
	 GOARCH=amd64 GOOS=linux go build -o ${BINARY_NAME}-linux main.go classroom.go database.go priority.go table.go
	 GOARCH=amd64 GOOS=windows go build -o ${BINARY_NAME}-windows main.go classroom.go database.go priority.go table.go

clean:
	 go clean
	 rm ${BINARY_NAME}-darwin
	 rm ${BINARY_NAME}-linux
	 rm ${BINARY_NAME}-windows
