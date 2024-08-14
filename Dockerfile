FROM golang:1.23-alpine AS build

WORKDIR /app

COPY go.mod ./
COPY go.sum ./

RUN go mod download

COPY . .

RUN go build -o myserver .

FROM alpine:3.20 AS run

EXPOSE 8000

WORKDIR /app

COPY --from=build --chown=1000:1000 /app/myserver /app/myserver

USER 1000

CMD ["/app/myserver"]
