FROM golang:1.24-alpine AS builder

RUN apk update --no-cache && apk add --no-cache tzdata
WORKDIR /build
ADD go.mod .
ADD go.sum .
RUN go env -w GO111MODULE=on
RUN go env -w GOPROXY=https://goproxy.cn,direct
RUN go mod download
COPY . .

RUN go build -ldflags="-s -w" -o /app/gateway-srv .

FROM alpine

COPY --from=builder /usr/share/zoneinfo/Asia/Shanghai /usr/share/zoneinfo/Asia/Shanghai
ENV TZ Asia/Shanghai

WORKDIR /app
COPY --from=builder /app/gateway-srv /app/gateway-srv
RUN chmod +x /app/gateway-srv

ENTRYPOINT ["/app/gateway-srv"]