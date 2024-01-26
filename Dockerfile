FROM golang:1.21.1 as build

WORKDIR /app
COPY . .

RUN cd /app && GOOS=linux CGO_ENABLED=0 go build -o temp-by-cep cmd/main.go

FROM scratch

WORKDIR /app

COPY --from=build /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/
COPY --from=build /app/temp-by-cep ./
COPY --from=build /app/.env ./

EXPOSE 8080

ENTRYPOINT ["./temp-by-cep"]