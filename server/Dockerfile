FROM golang:1.20 AS builder
ENV CGO_ENABLED 0
ADD . /app
WORKDIR /app
RUN go build -ldflags "-s -w" -v -o server .

FROM alpine:3
RUN apk update && \
    apk add openssl && \
    rm -rf /var/cache/apk/* \
    && mkdir /app

WORKDIR /app

ADD Dockerfile /Dockerfile

COPY --from=builder /app/server /app/server

RUN chown nobody /app/server \
    && chmod 500 /app/server

USER nobody

ENTRYPOINT ["/app/server"]
