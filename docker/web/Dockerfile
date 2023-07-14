FROM golang:1.18-alpine

WORKDIR /web

ENV GOPROXY https://goproxy.io

COPY go.mod ./
COPY go.sum ./
RUN go mod download

COPY cmd/ cmd/
COPY pkg/ pkg/

RUN CGO_ENABLED=0 GOOS=linux go build -a -o app cmd/web/main.go

EXPOSE 4000

ENTRYPOINT [ "/web/app" ]