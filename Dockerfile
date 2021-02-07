FROM golang:latest

WORKDIR /backend

RUN go get -u github.com/mattn/go-sqlite3

COPY go.mod go.sum ./
RUN go mod download

COPY . .
RUN go build -o main .

EXPOSE 9010

CMD ["./main"]