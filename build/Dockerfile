FROM golang:1.14.6 AS builder
WORKDIR /usr/local/go/src/scan-visor/

COPY go.mod .
COPY go.sum .

COPY cmd/server/ .
COPY pkg pkg

RUN mkdir bin
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o ./bin/server
RUN chmod +x ./bin/*

FROM alpine:latest
RUN apk --no-cache add ca-certificates
WORKDIR /root/

COPY --from=builder /usr/local/go/src/scan-visor/bin .
CMD ["./server"]