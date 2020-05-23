FROM golang:alpine

ENV PATH $GOPATH/bin:/usr/local/go/bin:$PATH

RUN apk update; apk upgrade; apk add -U --no-cache \
    ca-certificates \
		musl-dev \
		gcc \
    git \
    make \
		sqlite-dev \
    protoc \
    protobuf-dev

WORKDIR /app

ADD go.mod go.sum ./

RUN go mod download

ADD . .

RUN make deps

RUN make proto

CMD make run
