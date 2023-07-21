FROM golang:1.20 AS build

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod download

COPY main.go /app/main.go
COPY cmd /app/cmd

RUN CGO_ENABLED=0 GOOS=linux go build -o /out/d4bot

FROM debian AS certs

RUN apt-get update && apt-get install -y ca-certificates

FROM scratch AS deploy

COPY --from=certs /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/ca-certificates.crt

COPY --from=build /out/d4bot .

CMD ["./d4bot"]