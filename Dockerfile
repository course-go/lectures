FROM golang:latest

WORKDIR /usr/lectures
COPY . .

RUN go install golang.org/x/tools/cmd/present@latest

CMD present -http=:8080 -play=false
