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

WORKDIR /root

RUN apk --no-cache add ca-certificates \
    && mkdir keys

COPY build/config.yaml .
COPY keys keys/

COPY --from=builder /dao/rpc-mysql/server/app .

CMD [ "./app", "-config", "config.yaml" ]
