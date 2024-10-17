# syntax=docker/dockerfile:1
FROM golang:1.21-bullseye@sha256:26c7537d6ac3827eb4638034d16edc64de57bb011c8cc8fe301ac13a6568f6f4 AS build
WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY ./ ./
RUN make build && make build-batch-download

FROM debian:bullseye-slim@sha256:610b4c7ad241e66f6e2f9791e3abdf0cc107a69238ab21bf9b4695d51fd6366a
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
