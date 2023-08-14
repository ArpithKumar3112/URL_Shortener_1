#Build Stage
FROM golang:alpine AS builder
ENV GO111MODULE=on
RUN mkdir /app
WORKDIR /app
COPY go.mod ./
COPY go.sum ./
RUN go mod download
COPY . .
RUN go build -o /app/url-shortener-1

#Deploy Stage
FROM alpine:latest
WORKDIR /
RUN mkdir /app
COPY --from=builder /app/url-shortener-1 /app/url-shortener-1
EXPOSE  8080
CMD ["/app/url-shortener-1"]