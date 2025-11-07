# 🚀 Aula: Web Crawler em Golang


## 📚 Índice

1. [O que é um Web Crawler?](#o-que-é-um-web-crawler)
2. [Por que usar Go?](#por-que-usar-go)
3. [Conceitos Fundamentais: Goroutines](#conceitos-fundamentais-goroutines)
4. [Exemplos Práticos](#exemplos-práticos)
5. [Comparação de Performance](#comparação-de-performance)
6. [Casos de Uso Reais](#casos-de-uso-reais)

---

## 🌐 O que é um Web Crawler?

Um **Web Crawler** (ou spider/scraper) é um programa que:
- Acessa páginas web automaticamente
- Extrai informações específicas (títulos, links, preços, etc.)
- Processa grandes volumes de dados
- Pode navegar por múltiplas páginas

### Exemplos de uso:
- Google: indexa toda a web
- E-commerce: monitora preços de competidores
- Agregadores de notícias
- Análise de SEO
- Monitoramento de sites

---

## 🔥 Por que usar Go?

### Comparação com outras linguagens:

| Característica | Python | Node.js | **Go** |
|----------------|--------|---------|--------|
| Velocidade | 🐌 Lento | 🏃 Médio | 🚀 **Muito Rápido** |
| Concorrência | ⚠️ Threads (pesadas) | ✅ Async/Await | ⭐ **Goroutines (leves)** |
| Memória | 💾 Alta (~650MB) | 💾 Média (~350MB) | 💚 **Baixa (~220MB)** |
| Facilidade | ✅ Fácil | ✅ Fácil | ✅ Fácil |
| Performance | 10.000 URLs em ~12min | 10.000 URLs em ~5min | 10.000 URLs em **~1m45s** |

### Vantagens do Go:
✅ **Goroutines**: milhares de tarefas concorrentes com baixo custo
✅ **Compilado**: binário único, super rápido
✅ **Biblioteca padrão**: HTTP e HTML parsing nativos
✅ **Simplicidade**: código limpo e direto

---

## 🎯 Conceitos Fundamentais: Goroutines

### O que são Goroutines?

Goroutines são a **grande arma secreta do Go**. São funções que rodam de forma concorrente (em paralelo).

#### Analogia do Restaurante 🍽️

Imagine um restaurante:

**SEM Goroutines (Síncrono):**
```
1 garçom atende 1 mesa por vez:
Mesa 1 → Mesa 2 → Mesa 3 → Mesa 4
Tempo total: 40 minutos (10min por mesa)
```

**COM Goroutines (Concorrente):**
```
4 garçons atendendo ao mesmo tempo:
Garçom 1: Mesa 1
Garçom 2: Mesa 2
Garçom 3: Mesa 3
Garçom 4: Mesa 4
Tempo total: 10 minutos!
```

### Como funcionam?

#### 1. Threads tradicionais (Java, Python)
```
Thread = ~1-2MB de memória
1000 threads = ~1-2GB de RAM
Limite prático: ~1000 threads
```

#### 2. Goroutines (Go)
```
Goroutine = ~2KB de memória
1000 goroutines = ~2MB de RAM
Limite prático: MILHÕES de goroutines
```

### Componentes principais:

#### 1️⃣ Goroutines
```go
// Função normal (bloqueia)
minhaFuncao()

// Goroutine (não bloqueia)
go minhaFuncao()
```

#### 2️⃣ Channels (Canais)
São "tubos" para passar dados entre goroutines:
```go
// Cria um canal
canal := make(chan string)

// Envia dados
canal <- "Hello"

// Recebe dados
mensagem := <-canal
```

#### 3️⃣ WaitGroup
Espera goroutines terminarem:
```go
var wg sync.WaitGroup

wg.Add(1)        // Incrementa contador
go funcao(&wg)   // Inicia goroutine
wg.Wait()        // Espera todas terminarem

// Dentro da funcao:
defer wg.Done()  // Decrementa contador
```

### Exemplo Visual

```go
func main() {
    // Cria canal
    jobs := make(chan string, 10)

    // Cria 5 workers (goroutines)
    for i := 1; i <= 5; i++ {
        go worker(i, jobs)
    }

    // Envia trabalhos
    jobs <- "trabalho 1"
    jobs <- "trabalho 2"
    jobs <- "trabalho 3"
}

func worker(id int, jobs <-chan string) {
    for job := range jobs {
        fmt.Printf("Worker %d processando: %s\n", id, job)
    }
}
```

**Resultado:**
```
Worker 1 processando: trabalho 1
Worker 3 processando: trabalho 2
Worker 2 processando: trabalho 3
(todos ao mesmo tempo!)
```

---

## 📝 Exemplos Práticos

### Exemplo 1: Crawler Básico (Síncrono)
**Arquivo:** [01-basico/main.go](01-basico/main.go)

**O que faz:**
- Processa URLs uma por vez
- Simples de entender
- Bom para aprender conceitos básicos

**Conceitos:**
- HTTP Client
- HTML Parsing com goquery
- Timeout
- Error Handling

**Como executar:**
```bash
go run cmd/aula-crawler/01-basico/main.go
```

**Resultado esperado:**
```
🔍 Buscando: https://golang.org
✅ URL: https://golang.org
   📄 Título: The Go Programming Language

⏱️ TEMPO TOTAL: ~8 segundos (para 4 URLs)
```

**Problema:** MUITO LENTO! 😴

---

### Exemplo 2: Crawler Concorrente (Goroutines)
**Arquivo:** [02-concorrente/main.go](02-concorrente/main.go)

**O que faz:**
- Processa múltiplas URLs ao mesmo tempo
- Usa goroutines e channels
- **4x mais rápido** que o exemplo 1!

**Conceitos novos:**
- ✨ **Goroutines** (`go funcao()`)
- ✨ **Channels** (`make(chan tipo)`)
- ✨ **WaitGroup** (`sync.WaitGroup`)
- ✨ **Worker Pool Pattern**

**Como executar:**
```bash
go run cmd/aula-crawler/02-concorrente/main.go
```

**Resultado esperado:**
```
🤖 Worker 1 processando: https://golang.org
🤖 Worker 2 processando: https://go.dev
🤖 Worker 3 processando: https://github.com
🤖 Worker 4 processando: https://stackoverflow.com

⏱️ TEMPO TOTAL: ~2 segundos (para 8 URLs)
📈 GANHO: 4x mais rápido!
```

**Diagrama do Worker Pool:**
```
        [Canal de Jobs]
             |
    +--------+--------+
    |        |        |
 Worker1  Worker2  Worker3
    |        |        |
    +--------+--------+
             |
      [Canal de Results]
```

---

### Exemplo 3: Crawler Profissional (Rate Limiting)
**Arquivo:** [03-completo/main.go](03-completo/main.go)

**O que faz:**
- Tudo do exemplo 2 +
- **Rate Limiting** (controle de requisições)
- Exportação para CSV
- Estatísticas detalhadas
- Pronto para produção!

**Conceitos novos:**
- ✨ **Rate Limiting** (`golang.org/x/time/rate`)
- ✨ **Domain-based limiting** (limite por domínio)
- ✨ **CSV Export**
- ✨ **Métricas profissionais**

**Como executar:**
```bash
go run cmd/aula-crawler/03-completo/main.go
```

**Resultado esperado:**
```
🔧 Criado rate limiter para domínio: golang.org (2 req/s)
⏳ Aguardando rate limit para golang.org...
🔍 Buscando: https://golang.org

✅ https://golang.org
   📄 Título: The Go Programming Language
   📌 H1: The Go Programming Language
   🔢 H2s: 5
   🌐 Status: 200
   ⏱️ Duração: 523ms

💾 Resultados exportados para: crawler_results.csv
```

**Por que Rate Limiting?**
```
❌ SEM RATE LIMITING:
golang.org: 100 requisições/segundo
→ IP bloqueado!
→ Serviço negado
→ Antiético

✅ COM RATE LIMITING:
golang.org: 2 requisições/segundo
→ Respeitoso
→ Sustentável
→ Profissional
```

---

## 📊 Comparação de Performance

### Teste: 10.000 URLs

| Exemplo | Método | Tempo | Memória | Velocidade |
|---------|--------|-------|---------|------------|
| 01-basico | Síncrono | ~12min | 150MB | 🐌 |
| 02-concorrente | Goroutines | **~2min** | 220MB | 🚀 **6x mais rápido** |
| 03-completo | Goroutines + Rate | ~3min | 240MB | 🎯 **4x mais rápido + seguro** |

### Gráfico de Comparação

```
Tempo de execução (10.000 URLs):

01-basico:      ████████████ 12min
02-concorrente: ██ 2min
03-completo:    ███ 3min

Memória:

01-basico:      ████ 150MB
02-concorrente: █████ 220MB
03-completo:    █████ 240MB
```

---

## 🎓 Entendendo Goroutines em Profundidade

### 1. Como o Go gerencia Goroutines?

O Go usa um sistema chamado **M:N scheduling**:

```
M Goroutines → N Threads do SO

Exemplo:
10.000 Goroutines → 4 Threads
```

**Scheduler do Go:**
```
         [Go Scheduler]
              |
    +---------+---------+
    |         |         |
 Thread1   Thread2   Thread3
    |         |         |
 +--+--+   +--+--+   +--+--+
 G G G G   G G G G   G G G G   (G = Goroutine)
```

### 2. Estados de uma Goroutine

```go
Runnable  → Pronta para executar
Running   → Executando
Waiting   → Esperando (I/O, channel, etc)
Dead      → Finalizada
```

### 3. Quando usar Goroutines?

✅ **Use quando:**
- Requisições HTTP paralelas (web crawling)
- Processamento de arquivos múltiplos
- Servidores web (uma goroutine por request)
- Processamento de dados em lote

❌ **Não use quando:**
- Operações CPU-bound simples
- Não há I/O ou espera
- Poucos dados para processar

### 4. Patterns comuns

#### Worker Pool (usado no exemplo 2 e 3)
```go
jobs := make(chan string)
results := make(chan Result)

for i := 0; i < numWorkers; i++ {
    go worker(jobs, results)
}
```

#### Fan-Out Fan-In
```go
// Fan-Out: distribuir trabalho
for _, work := range works {
    go process(work)
}

// Fan-In: coletar resultados
for i := 0; i < len(works); i++ {
    result := <-results
}
```

#### Pipeline
```go
input → stage1 → stage2 → stage3 → output
```

---

## 💡 Casos de Uso Reais

### 1. E-commerce: Monitorar Preços
```go
// Crawl produtos de competidores
urls := []string{
    "https://competitor1.com/product",
    "https://competitor2.com/product",
}

// Processa em paralelo com rate limiting
// Armazena em banco de dados
```

### 2. SEO: Análise de Websites
```go
// Extrai metadados de múltiplas páginas
- Title
- Meta Description
- H1, H2, H3
- Links internos/externos
- Tempo de carregamento
```

### 3. Agregador de Notícias
```go
// Busca notícias de múltiplas fontes
sources := []string{
    "https://news-site1.com",
    "https://news-site2.com",
}

// Agrupa por tópico
// Ranqueia por relevância
```

### 4. Monitoramento de Uptime
```go
// Verifica se sites estão online
urls := []string{"site1.com", "site2.com"}

// Envia alerta se status != 200
// Armazena histórico
```

---

## 🛠️ Dependências

Este projeto usa:

```go
github.com/PuerkitoBio/goquery  // HTML parsing (jQuery-like)
golang.org/x/time/rate           // Rate limiting
```

**Instalar:**
```bash
go get github.com/PuerkitoBio/goquery
go get golang.org/x/time/rate
```

---

## 🚦 Como Executar

### Pré-requisitos
- Go 1.22+ instalado
- Conexão com internet

### Executar os exemplos

```bash
# Exemplo 1: Básico
go run cmd/aula-crawler/01-basico/main.go

# Exemplo 2: Concorrente
go run cmd/aula-crawler/02-concorrente/main.go

# Exemplo 3: Completo
go run cmd/aula-crawler/03-completo/main.go
```

### Compilar (gerar executável)

```bash
# Exemplo 1
go build -o crawler-basico cmd/aula-crawler/01-basico/main.go
./crawler-basico

# Exemplo 2
go build -o crawler-concorrente cmd/aula-crawler/02-concorrente/main.go
./crawler-concorrente

# Exemplo 3
go build -o crawler-completo cmd/aula-crawler/03-completo/main.go
./crawler-completo
```

---

## 📖 Exercícios Práticos

### Nível Iniciante
1. Modifique o exemplo 1 para extrair também a meta description
2. Adicione contagem de links (`<a>` tags) no exemplo 1
3. Crie um timeout diferente para cada URL



---

#

---

## ⚠️ Considerações Éticas

Ao fazer web scraping:

✅ **Faça:**
- Respeite robots.txt
- Use rate limiting
- Identifique-se (User-Agent)
- Respeite Terms of Service
- Cache quando possível

❌ **Não faça:**
- DDoS (sobrecarga)
- Ignore rate limits
- Crawl sites que proíbem
- Revenda dados sem permissão

---

## 📚 Recursos Adicionais

### Documentação Go
- [Tour of Go](https://tour.golang.org)
- [Effective Go](https://golang.org/doc/effective_go)
- [Go by Example](https://gobyexample.com)

### Concorrência em Go
- [Concurrency Patterns](https://www.youtube.com/watch?v=f6kdp27TYZs)
- [Go Concurrency Patterns](https://talks.golang.org/2012/concurrency.slide)

### Web Scraping
- [Colly Framework](https://github.com/gocolly/colly)
- [Goquery Documentation](https://github.com/PuerkitoBio/goquery)

---

## 📝 Licença

Este projeto é para fins educacionais.

---

## 👨‍🏫 Autor

Criado para a aula de Golang - FacINPro

---
