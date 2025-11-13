# 🚀 Quick Start - Web Crawler em Go

## 📦 Instalação Rápida

```bash
# 1. Instalar dependências
go get github.com/PuerkitoBio/goquery
go get golang.org/x/time/rate

# 2. Testar se está funcionando
go run cmd/aula-crawler/01-basico/main.go
go run cmd/aula-crawler/02-concorrente/main.go
go run cmd/aula-crawler/03-completo/main.go
```

---

## 🎓 Ordem de Aprendizado

### 1️⃣ Exemplo Básico (15 minutos)
**Arquivo:** `01-basico/main.go`

**O que você vai aprender:**
- Como fazer requisições HTTP
- Como parsear HTML com goquery
- Tratamento de erros
- Timeouts

**Execute:**
```bash
go run cmd/aula-crawler/01-basico/main.go
```

**Resultado:** Crawler simples, mas LENTO 🐌

---

### 2️⃣ Exemplo Concorrente (30 minutos)
**Arquivo:** `02-concorrente/main.go`

**O que você vai aprender:**
- 🔥 **GOROUTINES** (o superpoder do Go!)
- Channels (comunicação entre goroutines)
- WaitGroups (sincronização)
- Worker Pool Pattern

**Execute:**
```bash
go run cmd/aula-crawler/02-concorrente/main.go
```

**Resultado:** 4x mais rápido! 🚀

---

### 3️⃣ Exemplo Profissional (45 minutos)
**Arquivo:** `03-completo/main.go`

**O que você vai aprender:**
- Rate Limiting (controle por domínio)
- Exportação para CSV
- Métricas avançadas
- Pronto para produção!

**Execute:**
```bash
go run cmd/aula-crawler/03-completo/main.go
```

**Resultado:** Crawler profissional com responsabilidade! 🎯

---

## 📚 Documentação Completa

| Arquivo | Descrição |
|---------|-----------|
| [README.md](README.md) | Guia completo da aula |
| [GOROUTINES.md](GOROUTINES.md) | Tudo sobre Goroutines |
| Este arquivo | Quick Start |

---

## 🎯 Conceitos Principais

### Goroutines em 1 Minuto

```go
// Função normal (espera terminar)
minhaFuncao()

// Goroutine (roda em paralelo)
go minhaFuncao()
```

**Por que são incríveis?**
- Thread normal: ~1MB de memória
- Goroutine: ~2KB de memória (500x menor!)
- Você pode ter MILHÕES rodando ao mesmo tempo

### Worker Pool em 1 Minuto

```go
// 1. Criar canais
jobs := make(chan string)
results := make(chan Result)

// 2. Criar workers
for i := 0; i < 5; i++ {
    go worker(jobs, results)
}

// 3. Enviar trabalhos
for _, url := range urls {
    jobs <- url
}
```

---

## 💡 Comparação Visual

```
EXEMPLO 1 (Síncrono):
URL1 → URL2 → URL3 → URL4
⏱️  8 segundos

EXEMPLO 2 (4 Workers):
Worker1: URL1 → URL5
Worker2: URL2 → URL6
Worker3: URL3 → URL7
Worker4: URL4 → URL8
⏱️  2 segundos (4x mais rápido!)

EXEMPLO 3 (Com Rate Limit):
Igual ao exemplo 2, mas:
✅ Respeita servidores
✅ Evita bloqueios
✅ Profissional
```

---

## 🎓 Próximos Passos

Depois de terminar os 3 exemplos:

1. ✅ Leia [GOROUTINES.md](GOROUTINES.md) para aprofundar
2. ✅ Faça os exercícios do [README.md](README.md)
3. ✅ Modifique os exemplos
4. ✅ Crie seu próprio crawler!

---

## 🆘 Problemas Comuns

### Erro: "package not found"
```bash
go mod tidy
go get github.com/PuerkitoBio/goquery
```

### Erro: "timeout"
```bash
# Aumente o timeout no código
client := http.Client{
    Timeout: 30 * time.Second, // Era 10s
}
```

### Erro: "too many open files"
```bash
# Reduza o número de workers
workerCount := 3 // Era 5
```

---

## 📊 Estatísticas Esperadas

| Exemplo | URLs | Tempo | Memória |
|---------|------|-------|---------|
| 01-basico | 4 | ~8s | 150MB |
| 02-concorrente | 8 | ~2s | 220MB |
| 03-completo | 10 | ~3s | 240MB |

---
