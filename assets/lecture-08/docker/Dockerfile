FROM golang:alpine

ARG SERVICE
WORKDIR /usr/app

COPY . .
RUN go build -v -o /usr/app/bin/${SERVICE} cmd/${SERVICE}/main.go

EXPOSE 8080/tcp

ENTRYPOINT ["/usr/app/bin/${SERVICE}"]
