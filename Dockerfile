FROM golang:1.22.5-alpine3.20 AS build-dev

WORKDIR /go/src/app

COPY . .

RUN go env -w GO111MODULE=on && \
    go env -w GOPROXY=https://goproxy.cn,direct && \
    go mod download && \
    apk add --no-cache upx && \
    go build -trimpath  -ldflags="-s -w" -o sensitive-api . && \
    [ -e /usr/bin/upx ] && upx sensitive-api || echo

FROM alpine

COPY --from=build-dev /go/src/app/sensitive-api ./holiday
COPY dict ./dict

EXPOSE 9210

CMD ["./sensitive-api"]