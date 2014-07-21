FROM google/golang
MAINTAINER George MacRorie github.com/GeorgeMac

WORKDIR /gopath/src/hello-server
ADD . /gopath/src/hello-server/

RUN pwd

RUN go test ./...

RUN go install ./...

CMD []
ENTRYPOINT ["/gopath/bin/hello-server"]
