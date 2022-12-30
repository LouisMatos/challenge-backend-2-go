FROM golang:1.19-alpine3.17

WORKDIR /app

COPY . ./

RUN go mod download

RUN go build -o /challenge-backend-2-go

EXPOSE 8070

CMD ["/challenge-backend-2-go", "run"]

