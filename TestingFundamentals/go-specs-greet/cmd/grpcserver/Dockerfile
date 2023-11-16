FROM golang:1.21-alpine

WORKDIR /app

COPY go.mod ./

RUN go mod download

COPY . .

RUN go build -o svr cmd/grpcserver/*.go

EXPOSE 8080
CMD [ "./svr" ]