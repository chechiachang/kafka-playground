FROM golang:1.11.2 as builder
WORKDIR /go/src/github.com/chechiachang/kafka-playground/client-go/
COPY . .
RUN go get ./... \
  && go build -a -o app .

FROM scratch
WORKDIR /root/
COPY --from=builder /go/src/github.com/chechiachang/kafka-playground/client-go/app .
EXPOSE 8080
ENTRYPOINT ["./app"]
