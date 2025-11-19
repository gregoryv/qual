# version should match go.mod
FROM golang:1.24-bookworm

# install tools you need
RUN go install github.com/gregoryv/gocolor/cmd/gocolor@latest && \
    go install golang.org/x/tools/cmd/goimports@latest && \
    go install mvdan.cc/gofumpt@latest

# clear the cache before downloading dependencies to minimize final
# size of image
RUN go clean -cache

# download dependencies
WORKDIR /app
COPY . /app
RUN cd /app/ && \
    go get && \
    go test -run=^$ ./... && \
    cd .. && \
    rm -rf app
