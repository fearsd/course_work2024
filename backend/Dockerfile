FROM golang:latest

WORKDIR /app

COPY . .

RUN go mod download

RUN go install -mod=mod github.com/githubnemo/CompileDaemon

CMD CompileDaemon --build="go build server.go" --command=./server