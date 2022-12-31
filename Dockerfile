FROM golang:1.19-alpine3.17

WORKDIR /app

# pre-copy/cache go.mod for pre-downloading dependencies and only redownloading them in subsequent builds if they change
COPY go.mod go.sum ./
RUN go mod download && go mod verify

COPY . .

EXPOSE 8070
RUN go build 

CMD ["go", "run", "main.go"]
