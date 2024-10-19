# syntax=docker/dockerfile:1
FROM golang:1.23-bullseye@sha256:ecb3fe70e1fd6cef4c5c74246a7525c3b7d59c48ea0589bbb0e57b1b37321fb9 AS build
WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY ./ ./
RUN make build && make build-batch-download

FROM debian:bullseye-slim@sha256:9bec46ecd98ce4bf8305840b021dda9b3e1f8494a0768c407e2b233180fa1466
WORKDIR /app

RUN <<EOF
  apt-get update
  apt-get install -yqq --no-install-recommends ca-certificates
  rm -rf /var/lib/apt/lists/*
EOF

RUN groupadd -g 1000 app && useradd -u 1000 -g app app

USER app
COPY --from=build /app/server ./
COPY --from=build /app/batch-download ./
CMD ["./server"]
