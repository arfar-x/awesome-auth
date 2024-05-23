FROM golang:1.18-alpine
LABEL authors="alireza"

WORKDIR /app

COPY . .

RUN go mod download

# Install Golang migration
RUN go install -tags 'mysql' github.com/golang-migrate/migrate/v4/cmd/migrate@latest

RUN go build cmd/main.go

CMD ["main"]
