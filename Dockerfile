FROM golang:1.21-bookworm as builder
WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download
COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o main .

FROM alpine:latest as runtime
WORKDIR /root
COPY --from=builder /app/main .
EXPOSE 8040
CMD ["./main"]