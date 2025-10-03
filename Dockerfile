# ===========================================
# DOCKERFILE PARA APLICAÇÃO GO COM MONGODB
# ===========================================

# Usar Go oficial como base
FROM golang:1.22-alpine AS builder

# Instalar ca-certificates para conexões HTTPS
RUN apk --no-cache add ca-certificates git

# Definir diretório de trabalho
WORKDIR /app

# Copiar go.mod e go.sum para cache de dependências
COPY go.mod go.sum ./

# Baixar dependências
RUN go mod download

# Copiar código fonte
COPY . .

# Compilar a aplicação
RUN CGO_ENABLED=0 GOOS=linux go build -o main cmd/docker-mongo-app/main.go

# ===========================================
# IMAGEM FINAL (MULTI-STAGE BUILD)
# ===========================================
FROM alpine:latest

# Instalar ca-certificates para conexões HTTPS
RUN apk --no-cache add ca-certificates tzdata

# Criar diretório para a aplicação
WORKDIR /root/

# Copiar o binário da aplicação do builder
COPY --from=builder /app/main .

# Expor a porta da aplicação
EXPOSE 8080

# Comando para executar a aplicação
CMD ["./main"]