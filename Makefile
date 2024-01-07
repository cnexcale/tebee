BINARY_NAME=main.out

build:
	go build -o ${BINARY_NAME} src/main.go

run:
	go build -o ${BINARY_NAME} main.go
	./${BINARY_NAME}

clean:
	go clean
	rm ${BINARY_NAME}

test:
	go test teebee/src/bot