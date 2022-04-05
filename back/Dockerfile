FROM golang:1.18 AS build

WORKDIR /work

COPY go.mod .
COPY go.sum .
RUN go mod download

COPY . .
RUN \
    GOOS=linux \
    go build -o ./app

FROM ubuntu:20.04

WORKDIR /work

COPY --from=build /work/app .

ENTRYPOINT [ "/work/app" ]