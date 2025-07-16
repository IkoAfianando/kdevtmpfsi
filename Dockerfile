FROM golang:1.23-alpine AS builder

WORKDIR /app

COPY go.mod ./

COPY main.go .

RUN CGO_ENABLED=0 GOOS=linux go build -a -o /process-killer .

FROM alpine:latest

RUN apk --no-cache add procps

COPY --from=builder /process-killer /process-killer

CMD ["/process-killer"]