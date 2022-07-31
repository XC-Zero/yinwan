FROM golang:alpine

#ADD yiwan /yinwan

ENV GO111MODULE=on \
    CGO_ENABLED=0 \
    GOOS=linux \
    GOARCH=amd64 \
    GOPROXY="https://goproxy.cn,direct"

EXPOSE 3776

WORKDIR /yinwan
COPY D:\go\yinwan\cmd\build\yinwan ./app

CMD ["./yinwan"]