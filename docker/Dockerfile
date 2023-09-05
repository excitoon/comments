FROM library/golang:1.21.0

WORKDIR /app/src
ENV GOPATH=

COPY src /app/src

RUN go build . # TODO go get without code dependencies
