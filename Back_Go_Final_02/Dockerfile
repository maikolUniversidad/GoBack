FROM golang:1.21 as builder

WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download

COPY . .
RUN go build -o backend cmd/main.go

FROM gcr.io/distroless/static
COPY --from=builder /app/backend /backend
EXPOSE 8080
ENTRYPOINT ["/backend"]