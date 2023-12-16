FROM golang:1.21.5

RUN go install github.com/cosmtrek/air@latest

WORKDIR /usr/src/app/

COPY . .

RUN go mod tidy