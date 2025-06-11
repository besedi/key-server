FROM golang:1.24.1 AS build
WORKDIR /app
COPY go.mod ./
RUN go mod download
COPY . .
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o /key-server ./cmd/main.go

FROM gcr.io/distroless/static-debian12
COPY --from=build /key-server /
CMD ["/key-server"]
