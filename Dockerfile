FROM golang:alpine AS builder
WORKDIR /build
COPY /src .
RUN go mod download

ENV CGO_ENABLED=0
RUN go build -ldflags="-s -w" -o /main main.go

FROM scratch
COPY --from=builder main .
ENTRYPOINT ["./main"]
