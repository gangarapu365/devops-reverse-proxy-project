FROM golang:1.20

WORKDIR /app

COPY . .

RUN go mod init service_1 || true
RUN go build -o main .

EXPOSE 5000

CMD ["./main"]

