FROM golang:1.13-alpine

RUN mkdir /app

COPY . /app/

WORKDIR /app/

RUN apk --update add ca-certificates

RUN go build -mod=vendor -o main

CMD ["/app/main"]
