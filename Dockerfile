FROM golang:1.24-alpine AS builder
WORKDIR /app
COPY . .
RUN go env -w GO111MODULE=on
RUN go env -w GOPROXY=https://goproxy.cn,direct
RUN go mod tidy && go build -o /app/gateway-srv ./

FROM alpine:3.20
WORKDIR /app
COPY --from=builder /app/gateway-srv /app/gateway-srv
RUN chmod +x /app/gateway-srv

ENTRYPOINT ["/app/gateway-srv"]