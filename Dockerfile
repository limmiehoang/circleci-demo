FROM golang:1.7.4

ADD . /go/src/github.com/limmiehoang/circleci-demo

WORKDIR /go/src/github.com/limmiehoang/circleci-demo
RUN go install github.com/limmiehoang/circleci-demo/cmd/demo

COPY cmd/demo/demo.toml /etc/circleci-demo/demo.toml

CMD ["/go/bin/demo","-cfg","/etc/circleci-demo/demo.toml"]