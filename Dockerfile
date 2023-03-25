FROM golang:1.17-alpine AS builder

WORKDIR /src
COPY . .
RUN go build -o /bin/action main.go

FROM alpine:latest
COPY --from=builder /bin/action /bin/action
ENTRYPOINT ["/bin/action"]