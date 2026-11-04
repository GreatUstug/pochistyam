FROM golang:1.26 AS builder
LABEL authors="GreatUstug"

WORKDIR /backend

COPY go.mod ./
RUN go mod download

COPY . .

RUN CGO_ENABLED=0 GOOS=linux go mod tidy && go build -o main

FROM alpine:latest

WORKDIR /app

COPY --from=builder /backend/main .

EXPOSE 8080

CMD ["./main"]