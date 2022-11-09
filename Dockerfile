FROM golang:latest

WORKDIR /go/src/app

COPY . .
COPY init.sql /docker-entrypoint-initdb.d/

RUN go build -o ./bin/server .

CMD ./bin/server