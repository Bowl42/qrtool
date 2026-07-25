FROM golang:1.26-alpine AS build

WORKDIR /src

COPY go.mod go.sum ./
RUN go mod download

COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -trimpath -ldflags="-s -w" -o /out/qrtool ./cmd/qrtool

FROM alpine:3.22

RUN addgroup -S qrtool && adduser -S -G qrtool qrtool

COPY --from=build /out/qrtool /usr/local/bin/qrtool

USER qrtool
EXPOSE 8080

HEALTHCHECK --interval=30s --timeout=3s --start-period=5s --retries=3 \
	CMD wget -qO- http://127.0.0.1:8080/healthz || exit 1

ENTRYPOINT ["/usr/local/bin/qrtool"]
