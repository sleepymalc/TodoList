FROM golang:1.14.1

WORKDIR /go/src/app

COPY . .

RUN go get -u

RUN go build ./main.go

EXPOSE 8080

CMD ["./main"]