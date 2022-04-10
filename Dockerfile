FROM alpine:3.14

MAINTAINER coderj001

RUN apk add --no-cache git make musl-dev go

# Configure Go
ENV GOROOT /usr/lib/go
ENV GOPATH /go
ENV PATH /go/bin:$PATH

RUN go version

ADD . ${GOPATH}/src/github.com/coderj001/go-bookstore
WORKDIR ${GOPATH}/src/github.com/coderj001/go-bookstore

RUN go get -d -v ./...
RUN go install -v ./...
RUN go build cmd/main.go

CMD [ "./main" ]
