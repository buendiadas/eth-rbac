FROM golang:1.10

RUN mkdir -p /go/src/github.com/buendiadas/eth-rbac
WORKDIR /go/src/github.com/buendiadas/eth-rbac

ADD . /go/src/github.com/buendiadas/eth-rbac

RUN go get -v

EXPOSE 5050

CMD ["go", "run", "main.go"]