FROM ghcr.io/stanislav-zeman/go-present:latest

COPY *.slide go.mod go.sum /usr/present/
COPY assets /usr/present/assets

CMD present -http=:8080 -play=false
