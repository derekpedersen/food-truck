# Start from a Debian image with the latest version of Go installed
# and a workspace (GOPATH) configured at /go.
FROM debian:stretch-slim

# Install wget and install/updates certificates
RUN apt-get update \
 && apt-get install -y -q --no-install-recommends \
    ca-certificates \
 && apt-get clean \
 && rm -r /var/lib/apt/lists/*

# Copy the local package files to the container's workspace.
COPY ./bin/food-truck.linux /go/bin/

# Run the command by default when the container starts.
CMD /go/bin/food-truck.linux