FROM golang:latest AS builder
ADD . /app
WORKDIR /app/src
RUN go mod download
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -ldflags "-w" -a -o /main .
CMD /main
