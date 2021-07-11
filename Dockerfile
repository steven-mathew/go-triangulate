# --- Stage 1 ---
# Base Layer, contains Go build tools
# Name this stage 'wasm-builder' so we can refer to it later
FROM golang:alpine AS wasm-builder 

# Enable Go modules
ENV GO111MODULE=on

# Set working dir to /app/src
WORKDIR /app/src

# Copy dep files
COPY go.mod .
COPY go.sum .

# Install dependencies
RUN go mod download

# Install Git and network certificates
RUN apk add git ca-certificates

# Copy source code to image
COPY . /app/src

RUN cp "$(go env GOROOT)/misc/wasm/wasm_exec.js" .

RUN GOARCH=wasm GOOS=js go build -o lib.wasm /app/src/main.go

# RUN go run /app/src/server.go 
