#/bin/bash
echo "Building Super Cafacts services.."
go build && \
docker build .