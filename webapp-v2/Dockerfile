FROM golang:latest as builder

RUN mkdir /app

WORKDIR /app

COPY . . 

RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o server main.go

FROM alpine 

RUN apk add --no-cache ca-certificates

COPY --from=builder /app/server /app/server 

RUN mkdir -p /app/images

EXPOSE 9090

ENTRYPOINT ["/app/server"]
