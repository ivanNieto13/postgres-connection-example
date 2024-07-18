
FROM golang:1.22

WORKDIR /app

COPY go.* ./

RUN go mod download

COPY execute_simple_query ./execute_simple_query

COPY main.go ./

COPY .env ./

RUN go build -o app ./

ENTRYPOINT /app/app