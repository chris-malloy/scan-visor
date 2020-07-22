FROM golang:1.14.6 AS builder
WORKDIR /go/src/github.com/grnmnky/scan-visor/
COPY go.mod .
COPY go.sum .
COPY cmd/scan_visor/ .
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o cmd/scan_visor .

FROM alpine:latest
RUN apk --no-cache add ca-certificates
WORKDIR /root/

COPY --from=builder /go/src/github.com/grnmnky/scan-visor/cmd/scan_visor .
CMD ["./scan_visor"]