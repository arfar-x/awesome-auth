FROM golang
LABEL authors="alireza"

# Install Golang migration
RUN go install -tags 'mysql' github.com/golang-migrate/migrate/v4/cmd/migrate@latest

RUN go build cmd/main.go

ENTRYPOINT ["main"]