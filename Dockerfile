# syntax=docker/dockerfile:1
FROM golang:1.21-bullseye AS build
WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY ./ ./
RUN make build

FROM debian:bullseye-slim
WORKDIR /app

RUN <<EOF
  apt-get update
  apt-get install -yqq --no-install-recommends ca-certificates
  rm -rf /var/lib/apt/lists/*
EOF

RUN groupadd -g 1000 app && useradd -u 1000 -g app app

USER app
COPY --from=build /app/server ./
CMD ["./server"]
