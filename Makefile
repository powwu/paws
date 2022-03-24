BINARY_NAME=paws

build:
	
	 GOARCH=amd64 GOOS=linux go build -o ${BINARY_NAME}-linux ./src/main.go ./src/classroom.go ./src/database.go ./src/priority.go ./src/table.go
	 GOARCH=amd64 GOOS=darwin go build -o ${BINARY_NAME}-darwin ./src/main.go ./src/classroom.go ./src/database.go ./src/priority.go ./src/table.go
	 GOARCH=amd64 GOOS=windows go build -o ${BINARY_NAME}-windows ./src/main.go ./src/classroom.go ./src/database.go ./src/priority.go ./src/table.go

	 
	 

clean:
	 go clean
	 rm ${BINARY_NAME}-darwin
	 rm ${BINARY_NAME}-linux
	 rm ${BINARY_NAME}-windows
