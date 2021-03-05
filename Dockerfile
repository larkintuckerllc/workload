FROM golang:1.16 AS grpc-health-probe-builder
RUN go get github.com/grpc-ecosystem/grpc-health-probe

FROM golang:1.16 AS grpcurl-builder
RUN go get github.com/fullstorydev/grpcurl/...
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go install github.com/fullstorydev/grpcurl/cmd/grpcurl@latest

FROM golang:1.16 AS builder
WORKDIR /go/src/app
COPY go.mod go.sum ./
RUN go mod download
COPY . .
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go install -v ./cmd/workload

FROM alpine:3.13
COPY --from=grpc-health-probe-builder /go/bin/grpc-health-probe /usr/local/bin/
COPY --from=grpcurl-builder /go/bin/grpcurl /usr/local/bin/
COPY --from=builder /go/bin/workload /usr/local/bin/
EXPOSE 50051
USER 1000:1000
ENV PORT=50051
CMD ["workload"]

