# Build the manager binary
FROM golang:1.12.1 as builder

# Copy in the go src
WORKDIR /go/src/github.com/frodenas/pks-k8s-api
COPY pkg/    pkg/
COPY cmd/    cmd/
COPY vendor/ vendor/

# Build
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -a -o manager github.com/frodenas/pks-k8s-api/cmd/manager

# Copy the controller-manager into a thin image
FROM ubuntu:latest
WORKDIR /

RUN apt-get update && apt-get install -y ca-certificates

COPY --from=builder /go/src/github.com/frodenas/pks-k8s-api/manager .
ENTRYPOINT ["/manager"]
