FROM golang:1.22.2-alpine3.18 AS builder
WORKDIR /go/src/app

ARG VERSION="n/a"

RUN go install github.com/cespare/reflex@latest

COPY . .

RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 \
    go build \
    -o /score-board \
    -ldflags "-X 'main.Version=$VERSION'" \
    /go/src/app

FROM alpine:3.18.6 AS certs
RUN apk add --no-cache ca-certificates=20240226-r0

FROM scratch
COPY --from=certs /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/ca-certificates.crt
COPY --from=builder /score-board /
USER 9000
ENTRYPOINT [ "/score-board" ]