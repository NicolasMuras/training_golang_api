FROM golang:latest

WORKDIR /go/src/app

COPY . .

RUN go build -o ./bin/server .

CMD ./bin/server