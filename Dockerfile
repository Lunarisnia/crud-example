FROM golang:1.25.3-alpine3.22
WORKDIR /app

COPY . .
RUN go mod tidy

RUN GOOS=linux GOARCH=amd64 go build -ldflags "-s -w" -o ./dist/crud-example ./cmd/demo/

EXPOSE 3210

CMD ["./dist/crud-example"]
