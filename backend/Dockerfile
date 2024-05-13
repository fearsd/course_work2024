FROM golang:latest

WORKDIR /app

COPY ./backend /app

RUN go mod download

RUN go get github.com/githubnemo/CompileDaemon

ENTRYPOINT CompileDaemon --build="go build server.go" --command=./server