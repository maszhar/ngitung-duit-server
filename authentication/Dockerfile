FROM golang:1.19-alpine

WORKDIR /app
COPY go.mod ./
COPY go.sum ./
RUN go mod download
COPY . .

RUN mkdir -p dist
RUN go build -o /dist/app

EXPOSE 8000

CMD ["/dist/app"]