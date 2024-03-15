FROM ghcr.io/stanislav-zeman/go-present:latest

COPY *.slide assets go.mod go.sum /usr/present/

CMD present -http=:8080 -play=false
