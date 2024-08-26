FROM golang:1.23 AS builder

WORKDIR /service/api-diff

COPY . /service/api-diff

ENV CGO_ENABLED=0 GOOS=linux GOARCH=amd64
RUN go build -o diff-service .

FROM alpine:latest

ARG ENVIRONMENT
ENV ENVIRONMENT=$ENVIRONMENT

WORKDIR /opt

EXPOSE 8000

COPY --from=builder /service/api-diff/diff-service diff-service

RUN apk update 
RUN apk upgrade
RUN apk --no-cache add ca-certificates

ENTRYPOINT ["/opt/diff-service", "serve"]
