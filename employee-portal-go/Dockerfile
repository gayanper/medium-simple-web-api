FROM --platform=linux/arm64 golang:alpine AS builder

ENV GOOS linux
ENV GOARCH arm64

WORKDIR /build

COPY go.mod .
COPY go.sum .
COPY cmd ./cmd

RUN go mod download
RUN go build -o ./employee-portal-nix ./cmd/main.go

FROM scratch

ENV DB_HOST=

COPY --from=builder /build/employee-portal-nix /employee-portal-nix
ENTRYPOINT [ "/employee-portal-nix" ]