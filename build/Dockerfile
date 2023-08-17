
FROM golang:1.20-alpine as builder
WORKDIR /go/src/github.com/WoodExplorer/workdir

ENV GOPROXY="https://goproxy.cn"
ENV GO111MODULE=on
ENV GOSUMDB=sum.golang.google.cn

RUN sed -i 's/dl-cdn.alpinelinux.org/mirrors.aliyun.com/g' /etc/apk/repositories
RUN apk update
RUN apk --no-cache add git pkgconfig build-base
RUN apk --no-cache add binutils-gold

COPY go.mod .
COPY go.sum .
RUN go mod download
ADD . .
USER root
RUN go build -buildmode=pie  -o /go/bin/user-auth main.go

FROM alpine:3.11

RUN sed -i 's/dl-cdn.alpinelinux.org/mirrors.aliyun.com/g' /etc/apk/repositories
RUN apk update
RUN apk add jq
RUN apk add curl

WORKDIR /root

COPY --from=builder /go/bin/user-auth /root/user-auth
CMD ["./user-auth", "run"]
