FROM golang:1.24.1
WORKDIR /app
COPY go.mod ./
RUN go mod download
COPY . .
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o /key-server ./cmd/main.go
# Create and switch to a non-root user
RUN useradd -m appuser
USER appuser
CMD ["/key-server"]
