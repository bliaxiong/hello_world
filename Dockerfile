FROM golang:1.24-alpine AS builder
WORKDIR /app
COPY . .
RUN go test ./... && CGO_ENABLED=0 GOOS=linux go build -o hello_world_app .

FROM scratch
COPY --from=builder /app/hello_world_app /hello_world_app
EXPOSE 9000
ENTRYPOINT ["/hello_world_app"]