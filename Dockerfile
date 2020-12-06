FROM golang:1.15.3-alpine as builder

RUN apk --no-cache add git

WORKDIR /dao/

ADD . .

ENV GOOS=linux \
    GO111MODULE=on \
    GOPROXY=https://goproxy.cn,direct

RUN cd rpc-mysql/server \
    && go build -o app .

FROM alpine:latest as prod

RUN apk --no-cache add ca-certificates

WORKDIR /root/

COPY --from=builder /dao/rpc-mysql/server/app .

CMD [ "./app", "-config", "/dao/build/config.yaml" ]