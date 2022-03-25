FROM golang:1.18-alpine as server-build

# コンテナ内に作業ディレクトリを作成
RUN mkdir /go/src/work
# コンテナログイン時のディレクトリ指定
WORKDIR /go/src/work
# ホストのファイルをコンテナの作業ディレクトリに移行
ADD . /go/src/work
EXPOSE 443

COPY ./ ./go/src/work
RUN go mod download

RUN go build -o main
# CMD ["go", "run", "main.go"]

FROM alpine:latest

WORKDIR /app

ENV DOCKERIZE_VERSION v0.6.1

RUN apk --update add tzdata \
  # && cp /usr/share/zoneinfo/Asia/Tokyo /etc/localtime \
  && wget https://github.com/jwilder/dockerize/releases/download/$DOCKERIZE_VERSION/dockerize-alpine-linux-amd64-$DOCKERIZE_VERSION.tar.gz \
  && tar -C /usr/local/bin -xzvf dockerize-alpine-linux-amd64-$DOCKERIZE_VERSION.tar.gz \
  && rm dockerize-alpine-linux-amd64-$DOCKERIZE_VERSION.tar.gz \ 
  && apk add --update ca-certificates \
  && update-ca-certificates \
  && rm -rf /var/cache/apk/*

COPY --from=server-build /go/src/work/main ./
ENV PORT 443
CMD ["./main"]