FROM golang:1.21 as builder
ENV GOPROXY=https://goproxy.io
WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download
COPY . .
RUN go clean -modcache && go build -o main .

FROM ubuntu:latest
WORKDIR /app
COPY --from=builder /app/main .
CMD ["./main"]
