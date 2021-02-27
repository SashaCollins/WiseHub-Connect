# Futer work: fix timeout issues and go plugin error

FROM golang:alpine

ENV GO111MODULE=on  CGO_ENABLED=1 GOOS=linux GOARCH=amd64

RUN set -e; \
        apk add --update --no-cache --virtual .build-deps \
                gcc \
                libc-dev \
                linux-headers \
                mariadb-dev \
                postgresql-dev \
                pcre-dev \
                musl-dev \
                go \
                git \
                util-linux-dev \
                ca-certificates

RUN update-ca-certificates


WORKDIR /go/src/github/SashaCollins/Wisehub-Connect

COPY . .

RUN go get -u github.com/mattn/go-sqlite3
RUN go mod download
RUN go mod tidy && go mod verify

RUN go build -buildmode=plugin -o model/plugins/github/github.so model/plugins/github/Github.go
RUN go build -buildmode=plugin -o model/plugins/drone/drone.so model/plugins/drone/Drone.go

RUN go build -o main

EXPOSE 9010

CMD ["./main"]
#CMD ["go", "run", "main.go"]
