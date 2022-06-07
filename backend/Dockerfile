FROM golang:1.18 AS build

ARG _TODO_PROJECT_ID
ARG _TODO_FROUTEND_URL
ENV TODO_PROJECT_ID ${_TODO_PROJECT_ID}
ENV TODO_FROUTEND_URL ${_TODO_FROUTEND_URL}

WORKDIR /work

COPY go.mod .
COPY go.sum .
RUN go mod download

COPY . .

RUN \
    GOOS=linux \
    GOARCH=amd64 \
    go build -mod=readonly -v -o app

FROM --platform=linux/amd64 ubuntu:22.04

RUN set -x \
    && apt-get update && apt-get install -y --no-install-recommends \
        ca-certificates \
    && apt-get clean && rm -rf /var/lib/apt/lists/*

COPY --from=build /work/app /app

ENTRYPOINT [ "/app" ]
