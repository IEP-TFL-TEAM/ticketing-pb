FROM alpine:latest

ARG PB_VERSION=0.22.20

RUN apk add --no-cache \
  ca-certificates \ 
  go \
  openssh

WORKDIR /app

COPY . .
RUN go mod download
RUN go build -o /app/pb ./main.go

EXPOSE 8090

CMD ["/app/pb", "serve", "--http=0.0.0.0:8090"]
