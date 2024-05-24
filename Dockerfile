FROM golang:1.18-alpine
LABEL authors="Alireza"

WORKDIR /app

COPY . .

RUN go mod download

# Install Golang migration
RUN go install -tags 'mysql' github.com/golang-migrate/migrate/v4/cmd/migrate@latest

RUN go build -o /app/main_app cmd/main.go

RUN chmod +x /app/main_app

CMD ["./main_app"]
