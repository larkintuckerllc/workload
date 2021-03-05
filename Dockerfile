FROM golang:1.16 AS builder
WORKDIR /go/src/app
COPY go.mod go.sum ./
RUN go mod download
COPY . .
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go install -v ./cmd/workload

FROM alpine:3.13
COPY --from=builder /go/bin/workload /usr/local/bin/
EXPOSE 50051
USER 1000:1000
ENV PORT=50051
CMD ["workload"]

