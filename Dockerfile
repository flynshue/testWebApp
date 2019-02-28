FROM golang:1.9 AS builder

ENV DEP_URL="https://github.com/golang/dep/releases/download/v0.5.0/dep-linux-amd64"
ENV GOPATH=/go
EXPOSE 3000/tcp
COPY src/ /go/src/github.com/flynshuePersonal/testWebApp/

WORKDIR /go/src/github.com/flynshuePersonal/testWebApp

RUN curl -fsSL -o /usr/local/bin/dep ${DEP_URL} && \
    chmod +x /usr/local/bin/dep

RUN dep ensure
RUN go build -o bin/webapp

CMD bin/webapp
