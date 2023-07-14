FROM golang:1.18-alpine

WORKDIR /judge

ENV GOPROXY https://goproxy.io

COPY go.mod ./
COPY go.sum ./
RUN go mod download

COPY cmd/ cmd/
COPY pkg/ pkg/

RUN CGO_ENABLED=0 GOOS=linux go build -a -o app cmd/judge/main.go

ENTRYPOINT [ "/judge/app" ]