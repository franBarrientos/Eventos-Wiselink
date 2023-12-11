FROM golang:1.21.5-bullseye

WORKDIR /go/src/app

COPY . .

RUN chmod +x wait-for-it.sh

RUN go mod tidy

RUN go build -o main .

EXPOSE 8080

CMD ["./wait-for-it.sh", "my-mysql-db:3306", "--timeout=30", "--", "go", "run", "main.go"]
