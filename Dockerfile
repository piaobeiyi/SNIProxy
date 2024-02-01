FROM golang:alpine AS builder

LABEL stage=gobuilder

ENV CGO_ENABLED 0
ENV GOOS linux
# ENV GOPROXY https://goproxy.cn,direct

WORKDIR /build

ADD go.mod .
ADD go.sum .
RUN go mod download
COPY . .

RUN go build -ldflags="-s -w" -o /sniproxy fastgit.org/f-proxy-agent

FROM alpine

WORKDIR /
COPY --from=builder /sniproxy /sniproxy
CMD ["/sniproxy", "-c", "/config.yaml"]