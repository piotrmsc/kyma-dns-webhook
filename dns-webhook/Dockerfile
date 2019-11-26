# Build the manager binary
FROM golang:1.13.4 as builder

WORKDIR /dns-webhook

ENV CGO_ENABLED=0
ENV GOOS=linux
ENV GOARCH=amd64
ENV GO111MODULE=on

# Copy the go source
COPY . .

# Fetch dependencies
RUN go mod download

# Build
RUN go build -a -o dns-webhook ./cmd/webhook/main.go

# Use distroless as minimal base image to package the manager binary
# Refer to https://github.com/GoogleContainerTools/distroless for more details
FROM gcr.io/distroless/static:latest
WORKDIR /
COPY --from=builder /dns-webhook .
ENTRYPOINT ["/dns-webhook"]