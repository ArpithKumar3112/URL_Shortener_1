FROM golang:1.19
ENV GO111MODULE=on
RUN mkdir /app
WORKDIR /app
COPY go.mod ./
COPY go.sum ./
RUN go mod download
COPY . .
RUN go build -o /url-shortener-1
EXPOSE  8080
CMD ["/url-shortener-1"]