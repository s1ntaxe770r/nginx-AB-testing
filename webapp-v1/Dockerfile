FROM golang:latest as builder

RUN mkdir /app

WORKDIR /app

COPY . . 

RUN  GOOS=linux GOARCH=amd64 go build -o server main.go

FROM alpine 

RUN apk add --no-cache ca-certificates

COPY --from=builder /app/server /app/server 

RUN mkdir -p /app/images

EXPOSE 8080

ENTRYPOINT ["/app/server"]
