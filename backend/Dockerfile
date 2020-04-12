# 使用するGolangのイメージを指定する
FROM golang:1.14.2-alpine
ENV GO111MODULE=on

# 必要なパッケージなどなどをインストールする
RUN apk update \
    && apk add --no-cache git

EXPOSE 8080
