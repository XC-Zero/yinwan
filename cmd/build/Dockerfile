FROM golang:alpine3.13

ADD yiwan /go/src/yinwan

ENV GO111MODULE=on \
    CGO_ENABLED=0 \
    GOOS=linux \
    GOARCH=amd64 \
    GOPROXY="https://goproxy.cn,direct"

EXPOSE 3776

WORKDIR /go/src


CMD ["./yinwan"]