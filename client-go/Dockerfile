FROM golang:1.11.2 as builder
WORKDIR /go/src/github.com/chechiachang/kafka-playground/client-go/
COPY . .
RUN go get ./... \
  && GOOS=linux GOARCH=amd64 CGO_ENABLED=0 go build -a -o app .

FROM golang:1.11.2-alpine
WORKDIR /root/
COPY --from=builder /go/src/github.com/chechiachang/kafka-playground/client-go/app .
EXPOSE 8080
ENTRYPOINT ["./app"]
