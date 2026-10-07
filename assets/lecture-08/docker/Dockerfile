FROM golang:1.27-alpine

ARG SERVICE
WORKDIR /usr/app

COPY . .
RUN go build -v -o /usr/app/bin/app ./cmd/${SERVICE}

EXPOSE 8080/tcp

ENTRYPOINT ["/usr/app/bin/app"]
