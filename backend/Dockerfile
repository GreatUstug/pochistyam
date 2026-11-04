FROM golang:1.26 AS builder
LABEL authors="GreatUstug"

WORKDIR /backend

COPY go.mod ./
RUN go mod download

COPY . .
ENV CGO_ENABLED=0 GOOS=linux 
RUN go mod tidy && go build -o main

FROM alpine:latest

WORKDIR /backend

COPY --from=builder /backend/main .

EXPOSE 8080

CMD ["./main"]