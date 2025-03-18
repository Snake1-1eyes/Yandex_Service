FROM golang:1.23.4-alpine3.21

WORKDIR /app

COPY . .

RUN go mod download

RUN CGO=ENABLED=0 GOOS=linux go build -o /service ./cmd/main.go

EXPOSE 8081
EXPOSE 50051

CMD [ "/service" ]