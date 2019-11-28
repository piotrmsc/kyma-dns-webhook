# Build the manager binary
FROM golang:1.13.4 as builder

WORKDIR /dns-challenger

ENV CGO_ENABLED=0
ENV GOOS=linux
ENV GOARCH=amd64
ENV GO111MODULE=on

# Copy the go source
COPY . .

# Fetch dependencies
RUN go mod download

# Build
RUN go build -a -o dns-challenger ./cmd/server/main.go

# Use distroless as minimal base image to package the manager binary
# Refer to https://github.com/GoogleContainerTools/distroless for more details
# the debug image has busybox shell added
FROM gcr.io/distroless/base:debug
WORKDIR /
COPY --from=builder /dns-challenger/dns-challenger .
ENTRYPOINT ["/dns-challenger"]