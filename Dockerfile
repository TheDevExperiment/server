# syntax=docker/dockerfile:1

## Build
FROM golang:1.20.2-buster AS build

WORKDIR /app

COPY go.mod ./
COPY go.sum ./
RUN go mod download

COPY . ./

RUN go build -o /server

## Deploy
FROM gcr.io/distroless/base-debian11

WORKDIR /

COPY --from=build /server /server
COPY ./config.yml /config.yml

EXPOSE 8080

USER nonroot:nonroot

ENTRYPOINT ["/server"]