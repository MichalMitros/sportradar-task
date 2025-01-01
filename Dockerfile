FROM golang:1.22.2-alpine3.18 AS builder
WORKDIR /go/src/app

ARG VERSION="n/a"

COPY . .

RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 \
    go build \
    -o /score-board \
    -ldflags "-X 'main.Version=$VERSION'" \
    /go/src/app

FROM scratch
COPY --from=builder /score-board /
USER 9000
ENTRYPOINT [ "/score-board" ]