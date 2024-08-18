FROM golang:1.22.5 as dev

WORKDIR /api

RUN go install github.com/air-verse/air@latest && \
    go install github.com/sqldef/sqldef/cmd/mysqldef@latest && \
    go install github.com/k1LoW/tbls@latest

CMD ["air"]

