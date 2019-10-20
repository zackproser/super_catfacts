#/bin/bash
echo "Building Super Cafacts services.." && \
echo "Running tests..." && \
go test -v ./... && \
echo "Building Docker image..." && \
go build && \
docker build .