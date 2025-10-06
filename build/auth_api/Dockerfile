# Etapa 1: Build
FROM golang:1.25.1 AS builder

WORKDIR /app

# Copiar dependências primeiro para cache
COPY go.mod go.sum ./
RUN go mod download

# Copiar o restante do código
COPY . .

# Compilar binário a partir do main.go
RUN go build -o api ./cmd/auth_api

# Etapa 2: Imagem final mínima
FROM debian:bookworm-slim

RUN apt-get update && apt-get install -y ca-certificates && rm -rf /var/lib/apt/lists/*

WORKDIR /root/

# Copiar binário
COPY --from=builder /app/api .

EXPOSE 3000

CMD ["./api"]