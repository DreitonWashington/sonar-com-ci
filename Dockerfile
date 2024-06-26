FROM golang:latest
WORKDIR /app
RUN go mod init test
COPY . .
RUN go build -o main
CMD ["/main"]