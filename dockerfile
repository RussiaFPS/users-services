FROM golang:1.18 as builder

WORKDIR ./bin/users-services
COPY . .
RUN go mod download
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build ./cmd/main.go

FROM alpine:latest
WORKDIR .
COPY --from=builder /go/bin/users-services/main .
COPY --from=builder /go/bin/users-services/envs/ envs/
CMD ["./main"]