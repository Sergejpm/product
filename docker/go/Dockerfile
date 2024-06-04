FROM golang:bullseye AS build
RUN apt update && apt install -y bash curl gcc g++ git make
RUN BIN="/usr/local/bin" && \
VERSION="1.7.0" && \
  curl -sSL \
    "https://github.com/bufbuild/buf/releases/download/v${VERSION}/buf-$(uname -s)-$(uname -m)" \
    -o "${BIN}/buf" && \
  chmod +x "${BIN}/buf"
WORKDIR /app
COPY ../.. ./
RUN mkdir -p bin
RUN make build-api

FROM debian:bullseye
WORKDIR /app
COPY --from=build /app/bin/api ./api
RUN apt update && apt install -y curl && \
    apt install -y ca-certificates && \
    update-ca-certificates
ENTRYPOINT ["/app/api"]
