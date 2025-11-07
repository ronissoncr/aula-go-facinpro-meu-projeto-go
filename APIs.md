### 🔧 Usando Docker Compose diretamente:

```bash
# Desenvolvimento
docker-compose --profile dev up -d

# Homologação  
docker-compose --profile hml up -d

# Produção
docker-compose --profile prod up -d

# Todos os ambientes
docker-compose --profile dev --profile hml --profile prod up -d
```

## 🌐 Portas e Acessos:

| Ambiente | Aplicação | MongoDB | Mongo Express |
|----------|-----------|---------|---------------|
| **DEV**  | :8080     | :27017  | :8090         |
| **HML**  | :8081     | :27018  | -             |
| **PROD** | :8082     | :27019  | -             |

## 🛠️ Endpoints da API:


URL: localhost/

- `GET /` - Página inicial
- `GET /health` - Status da aplicação
- `GET /config` - Configurações (sem senhas)
- `GET /users` - Lista usuários
- `POST /users` - Cria usuário
