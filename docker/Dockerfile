FROM library/golang:1.21.0

COPY src /app/src

ENV GOPATH=

WORKDIR /app/src

RUN go build .
