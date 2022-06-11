FROM golang:1.18-buster as builder 
ENV GO111MODULE on
RUN apt-get update && apt-get install -y make git ca-certificates

WORKDIR /app
COPY ./backend/ ./
RUN go mod download
RUN go build -o grpc-server -ldflags="-s -w" server.go

FROM debian:buster-slim
RUN apt-get update && apt-get install -y ca-certificates
COPY --from=builder /app/grpc-server /app/grpc-server

EXPOSE 9090
EXPOSE 443
EXPOSE 8080
CMD ["/app/grpc-server"]