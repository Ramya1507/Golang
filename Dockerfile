#syntax=docker/dockerfile:1

FROM golang:1.16-alpine

WORKDIR /app

COPY go.mod ./

RUN go mod download

COPY *.go ./

RUN go build -v -o Golang .

EXPOSE 8087

CMD [ "./Golang" ]
