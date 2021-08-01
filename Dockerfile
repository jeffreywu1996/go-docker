FROM golang:1.16-alpine
WORKDIR /app

RUN apk update
RUN apk add --no-cache bash

COPY * ./
RUN go build
CMD ["./Go-naive-user-service"]

EXPOSE 8000
