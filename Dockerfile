FROM golang:1.22-alpine AS builder
WORKDIR /app
COPY . .
RUN go test ./... && CGO_ENABLED=0 GOOS=linux go build -o hello_world_app .

FROM scratch
COPY --from=builder /app/hello_world_app /hello_world_app
EXPOSE 8080
ENTRYPOINT ["/hello_world_app"]