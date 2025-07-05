FROM golang:1.23-alpine AS builder

COPY . /github.com/SemenTretyakov/chat_service/app
WORKDIR /github.com/SemenTretyakov/chat_service/app

RUN go mod download
RUN go build -o ./bin/chat_server cmd/app/main.go

FROM alpine:latest

WORKDIR /root/
COPY --from=builder /github.com/SemenTretyakov/chat_service/app/bin/chat_server .

CMD ["./chat_server"]