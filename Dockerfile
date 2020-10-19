FROM golang:alpine

ADD ./ /go/src

WORKDIR /go/src

ENV GO111MODULE on

ENV GOPROXY https://goproxy.io,direct

ENV CGO_ENABLED=0 GOOS=linux CGO_ENABLED=0

RUN chmod +x ./build_linux.sh && ./build_linux.sh

EXPOSE 8889

ENTRYPOINT ["./go_api"]

