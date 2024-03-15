FROM golang:latest

WORKDIR /usr/lectures
COPY *.slide .
COPY assets assets
COPY go.mod go.mod
COPY go.sum go.sum

RUN go install golang.org/x/tools/cmd/present@latest

CMD present -http=:8080 -play=false
