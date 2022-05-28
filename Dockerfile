FROM golang:1.18.1-alpine3.15 as builder
COPY . /usr/local/go/src/GoRss2Webhook
WORKDIR /usr/local/go/src/GoRss2Webhook
RUN CGO_ENABLED=1 GOOS=linux GOARCH=amd64 GO111MODULE=on GOPROXY=https://mirrors.aliyun.com/goproxy/,direct go build -o /usr/bin/GoRss2Webhook GoRss2Webhook

###
FROM alpine:3.15.0 as final
RUN apk update && apk add tzdata && cp /usr/share/zoneinfo/Asia/Shanghai /etc/localtime && mkdir /app
ENV LOG_HOME=/app/logs
ENV CONFIG_HOME=/app/etc
WORKDIR /app
ENV TZ=Asia/Shanghai
ENTRYPOINT ["/usr/bin/GoRss2Webhook"]
COPY --from=builder /usr/bin/GoRss2Webhook /usr/bin/