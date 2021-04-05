ARG GO_VERSION=1.15.6
FROM golang:${GO_VERSION}-alpine AS builder

RUN apk update && apk add alpine-sdk git && rm -rf /var/cache/apk/*

RUN mkdir -p /src
WORKDIR /src

COPY src/go.mod .
COPY src/go.sum .
RUN go mod download

COPY src/ .
RUN go build -o ./api ./main.go


FROM alpine:latest

RUN apk update && apk add ca-certificates && rm -rf /var/cache/apk/*

RUN mkdir -p /src
WORKDIR /src
COPY --from=builder /src/api .

EXPOSE 8080

ENTRYPOINT ["./api"]
