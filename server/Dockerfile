FROM golang:latest

LABEL maintainer="Enes Karali <github.com/eneskarali>"

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod download

COPY . .

RUN go build -o main .

EXPOSE 8000

CMD ["./main"]
