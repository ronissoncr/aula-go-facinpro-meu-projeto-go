# 🚀 Guia Completo: Goroutines em Go

Este guia explica  o conceito de **Goroutines**. 

---

## 📚 Índice

1. [O que são Goroutines?](#o-que-são-goroutines)
2. [Como funcionam internamente?](#como-funcionam-internamente)
3. [Goroutines vs Threads](#goroutines-vs-threads)
4. [Channels: Comunicação entre Goroutines](#channels-comunicação-entre-goroutines)
5. [Sincronização com WaitGroup](#sincronização-com-waitgroup)
6. [Patterns Comuns](#patterns-comuns)
7. [Problemas Comuns e Soluções](#problemas-comuns-e-soluções)
8. [Exemplos Práticos](#exemplos-práticos)

---

## 🎯 O que são Goroutines?

### Definição Simples
Uma **goroutine** é uma função que roda de forma concorrente (ao mesmo tempo) com outras funções.

### Analogia: Cozinha de Restaurante 🍳

**SEM Goroutines (Sequencial):**
```
Chef único faz tudo sozinho:
1. Corta cebola (5min)
2. Ferve água (5min)
3. Cozinha arroz (10min)
4. Prepara salada (5min)

Total: 25 minutos
```

**COM Goroutines (Paralelo):**
```
4 chefs trabalhando ao mesmo tempo:
Chef 1: Corta cebola (5min)     |
Chef 2: Ferve água (5min)       | Todos ao
Chef 3: Cozinha arroz (10min)   | mesmo tempo!
Chef 4: Prepara salada (5min)   |

Total: 10 minutos (tempo do mais demorado)
```

### Sintaxe Básica

```go
// Função normal (bloqueia/espera terminar)
minhaFuncao()

// Goroutine (não bloqueia/roda em paralelo)
go minhaFuncao()
```

### Exemplo Mínimo

```go
package main

import (
    "fmt"
    "time"
)

func tarefa(nome string) {
    for i := 1; i <= 3; i++ {
        fmt.Printf("%s: passo %d\n", nome, i)
        time.Sleep(100 * time.Millisecond)
    }
}

func main() {
    // Roda duas tarefas em paralelo
    go tarefa("Goroutine 1")
    go tarefa("Goroutine 2")

    // Espera um pouco (não faça isso em produção!)
    time.Sleep(1 * time.Second)
}
```

**Saída:**
```
Goroutine 1: passo 1
Goroutine 2: passo 1
Goroutine 1: passo 2
Goroutine 2: passo 2
Goroutine 1: passo 3
Goroutine 2: passo 3
```

---

## ⚙️ Como funcionam internamente?

### Arquitetura M:N
-M:N é um modelo de agendamento onde **M** goroutines são mapeadas para **N** threads do sistema operacional.
Go usa um modelo **M:N scheduling**:
- **M** goroutines rodam em **N** threads do sistema operacional
- O **scheduler** do Go distribui goroutines entre threads

```
        [Go Runtime Scheduler]
                 |
       +---------+---------+--------+
       |         |         |        |
    Thread1   Thread2   Thread3   Thread4  (OS Threads)
       |         |         |        |
    +--+--+   +--+--+   +--+--+   +--+--+
    G G G G   G G G G   G G G G   G G G G   (Goroutines)
```

### Componentes do Runtime

1. **G (Goroutine)**
   - Representa uma goroutine
   - Contém stack, registradores, etc.
   - ~2KB de memória inicial

2. **M (Machine/Thread)**
   - Thread do sistema operacional
   - Executa goroutines
   - Por padrão: GOMAXPROCS = número de CPUs

3. **P (Processor)**
   - Contexto de execução
   - Fila local de goroutines
   - Recursos para executar código Go

### Diagrama Completo

```
┌─────────────────────────────────────────┐
│           Go Program                     │
│                                          │
│  Goroutine 1   Goroutine 2  Goroutine 3 │
│      |              |            |       │
└──────┼──────────────┼────────────┼───────┘
       │              │            │
       └──────┬───────┴────────────┘
              │
    ┌─────────▼──────────┐
    │   Go Scheduler     │
    │  (Runtime)         │
    └─────────┬──────────┘
              │
    ┌─────────┴──────────┐
    │ OS Thread Pool     │
    │ Thread1   Thread2  │
    └─────────┬──────────┘
              │
    ┌─────────▼──────────┐
    │   CPU Cores        │
    │  Core1    Core2    │
    └────────────────────┘
```

---

## 🆚 Goroutines vs Threads

### Tabela Comparativa

| Característica | Thread (Java/Python) | Goroutine (Go) |
|----------------|---------------------|----------------|
| **Tamanho** | 1-2 MB | 2 KB (~500x menor) |
| **Criação** | Pesada (~1ms) | Leve (~20µs) |
| **Gerenciamento** | Sistema Operacional | Go Runtime |
| **Limite prático** | ~1.000 threads | Milhões |
| **Troca de contexto** | Cara (~1-2µs) | Barata (~0.2µs) |
| **Stack** | Fixo | Cresce dinamicamente |

### Exemplo Visual: Memória

```
1000 Threads (Java/Python):
████████████████████████████████████  1-2 GB

1000 Goroutines (Go):
█  2 MB

1.000.000 Goroutines (Go):
████████████████  2 GB
```

### Código Comparativo

**Java (Threads):**
```java
// Cria 1000 threads
for (int i = 0; i < 1000; i++) {
    new Thread(() -> {
        // faz algo
    }).start();
}
// Problema: consumo alto de memória
```

**Go (Goroutines):**
```go
// Cria 1000 goroutines
for i := 0; i < 1000; i++ {
    go func() {
        // faz algo
    }()
}
// Sem problemas! Leve e eficiente
```

---

## 📡 Channels: Comunicação entre Goroutines

### O que são Channels?

Channels são **"tubos"** que permitem goroutines se comunicarem de forma segura.

### Analogia: Sistema de Esteira 📦

```
      Goroutine 1             Goroutine 2
          |                        |
          v                        v
    [Produz dados]  →  [Canal]  →  [Consome dados]
          ↓                            ↓
      channel <- data          data <- channel
```

### Tipos de Channels

#### 1. Channel sem buffer (bloqueante)
```go
ch := make(chan string)

// Envia (bloqueia até alguém receber)
ch <- "mensagem"

// Recebe (bloqueia até chegar mensagem)
msg := <-ch
```

#### 2. Channel com buffer
```go
ch := make(chan string, 3) // buffer de 3 elementos

// Envia (não bloqueia até buffer encher)
ch <- "msg1"
ch <- "msg2"
ch <- "msg3"
ch <- "msg4" // BLOQUEIA aqui (buffer cheio!)
```

### Operações com Channels

```go
// Criar
ch := make(chan int)

// Enviar
ch <- 42

// Receber
value := <-ch

// Receber e ignorar
<-ch

// Fechar
close(ch)

// Verificar se fechado
value, ok := <-ch
if !ok {
    fmt.Println("Canal fechado!")
}
```

### Exemplo Prático: Produtor-Consumidor

```go
package main

import "fmt"

func produtor(ch chan<- int) {
    for i := 1; i <= 5; i++ {
        fmt.Printf("Produzindo: %d\n", i)
        ch <- i // Envia para o canal
    }
    close(ch) // Fecha quando terminar
}

func consumidor(ch <-chan int) {
    for valor := range ch { // Itera até canal fechar
        fmt.Printf("Consumindo: %d\n", valor)
    }
}

func main() {
    canal := make(chan int, 2) // Buffer de 2

    go produtor(canal)
    consumidor(canal) // Roda no main (bloqueia)
}
```

**Saída:**
```
Produzindo: 1
Produzindo: 2
Produzindo: 3
Consumindo: 1
Consumindo: 2
Produzindo: 4
Consumindo: 3
Produzindo: 5
Consumindo: 4
Consumindo: 5
```

### Direções de Channels

```go
// Bidirecional (padrão)
var ch chan string

// Somente envio
var chSend chan<- string

// Somente recebimento
var chRecv <-chan string

// Uso em funções (type safety)
func enviar(ch chan<- string) {
    ch <- "mensagem"
}

func receber(ch <-chan string) string {
    return <-ch
}
```

---

## ⏳ Sincronização com WaitGroup

### O que é WaitGroup?

**WaitGroup** é um contador que espera goroutines terminarem.

### Analogia: Chamada Escolar 📝

```
Professor: "Vou esperar todos terminarem a prova"

Aluno 1: terminou! (contador: 3 → 2)
Aluno 2: terminou! (contador: 2 → 1)
Aluno 3: terminou! (contador: 1 → 0)

Professor: "Ok, todos terminaram!"
```

### API do WaitGroup

```go
import "sync"

var wg sync.WaitGroup

// Adiciona N goroutines ao contador
wg.Add(N)

// Decrementa contador (chamado pela goroutine)
wg.Done()

// Espera contador chegar a 0
wg.Wait()
```

### Exemplo Completo

```go
package main

import (
    "fmt"
    "sync"
    "time"
)

func trabalhador(id int, wg *sync.WaitGroup) {
    defer wg.Done() // Garante que Done() seja chamado

    fmt.Printf("Trabalhador %d: iniciando\n", id)
    time.Sleep(time.Second)
    fmt.Printf("Trabalhador %d: finalizando\n", id)
}

func main() {
    var wg sync.WaitGroup

    // Inicia 5 trabalhadores
    for i := 1; i <= 5; i++ {
        wg.Add(1)
        go trabalhador(i, &wg)
    }

    fmt.Println("Aguardando todos terminarem...")
    wg.Wait()
    fmt.Println("Todos finalizaram!")
}
```

### ⚠️ Erro Comum: Race Condition no Add

```go
// ❌ ERRADO
for i := 0; i < 5; i++ {
    go func() {
        wg.Add(1) // Race condition!
        defer wg.Done()
        // trabalho
    }()
}

// ✅ CORRETO
for i := 0; i < 5; i++ {
    wg.Add(1) // Add ANTES de iniciar goroutine
    go func() {
        defer wg.Done()
        // trabalho
    }()
}
```

---

## 🎨 Patterns Comuns

### 1. Worker Pool

O pattern mais usado em web crawlers!

```go
func workerPool() {
    jobs := make(chan int, 100)
    results := make(chan int, 100)

    // Criar workers
    numWorkers := 5
    var wg sync.WaitGroup

    for i := 1; i <= numWorkers; i++ {
        wg.Add(1)
        go worker(i, jobs, results, &wg)
    }

    // Enviar jobs
    for j := 1; j <= 20; j++ {
        jobs <- j
    }
    close(jobs)

    // Esperar workers
    go func() {
        wg.Wait()
        close(results)
    }()

    // Coletar resultados
    for result := range results {
        fmt.Println(result)
    }
}

func worker(id int, jobs <-chan int, results chan<- int, wg *sync.WaitGroup) {
    defer wg.Done()
    for job := range jobs {
        fmt.Printf("Worker %d processando job %d\n", id, job)
        results <- job * 2
    }
}
```

### 2. Fan-Out, Fan-In

Distribui trabalho e coleta resultados.

```go
func fanOutFanIn() {
    input := make(chan int)

    // Fan-Out: distribui para múltiplas goroutines
    c1 := processor(input)
    c2 := processor(input)
    c3 := processor(input)

    // Fan-In: combina resultados
    output := merge(c1, c2, c3)

    // Envia dados
    go func() {
        for i := 1; i <= 10; i++ {
            input <- i
        }
        close(input)
    }()

    // Recebe resultados
    for result := range output {
        fmt.Println(result)
    }
}

func processor(input <-chan int) <-chan int {
    output := make(chan int)
    go func() {
        for i := range input {
            output <- i * 2
        }
        close(output)
    }()
    return output
}

func merge(cs ...<-chan int) <-chan int {
    out := make(chan int)
    var wg sync.WaitGroup

    for _, c := range cs {
        wg.Add(1)
        go func(ch <-chan int) {
            defer wg.Done()
            for v := range ch {
                out <- v
            }
        }(c)
    }

    go func() {
        wg.Wait()
        close(out)
    }()

    return out
}
```

### 3. Pipeline

Processa dados em estágios.

```go
func pipeline() {
    // Stage 1: gera números
    nums := generator(1, 2, 3, 4, 5)

    // Stage 2: multiplica por 2
    doubled := multiply(nums, 2)

    // Stage 3: adiciona 10
    added := add(doubled, 10)

    // Resultado
    for result := range added {
        fmt.Println(result) // 12, 14, 16, 18, 20
    }
}

func generator(nums ...int) <-chan int {
    out := make(chan int)
    go func() {
        for _, n := range nums {
            out <- n
        }
        close(out)
    }()
    return out
}

func multiply(in <-chan int, factor int) <-chan int {
    out := make(chan int)
    go func() {
        for n := range in {
            out <- n * factor
        }
        close(out)
    }()
    return out
}

func add(in <-chan int, addend int) <-chan int {
    out := make(chan int)
    go func() {
        for n := range in {
            out <- n + addend
        }
        close(out)
    }()
    return out
}
```

### 4. Select (Multiplexação)

Espera múltiplos channels ao mesmo tempo.

```go
func selectExample() {
    ch1 := make(chan string)
    ch2 := make(chan string)

    go func() {
        time.Sleep(1 * time.Second)
        ch1 <- "um"
    }()

    go func() {
        time.Sleep(2 * time.Second)
        ch2 <- "dois"
    }()

    for i := 0; i < 2; i++ {
        select {
        case msg1 := <-ch1:
            fmt.Println("Recebido:", msg1)
        case msg2 := <-ch2:
            fmt.Println("Recebido:", msg2)
        case <-time.After(3 * time.Second):
            fmt.Println("Timeout!")
        }
    }
}
```

---

## ⚠️ Problemas Comuns e Soluções

### 1. Goroutine Leak (Vazamento)

**Problema:**
```go
// ❌ Goroutine nunca termina
func leak() {
    ch := make(chan int)
    go func() {
        val := <-ch // Espera para sempre!
        fmt.Println(val)
    }()
    // Esqueceu de enviar algo para ch
}
```

**Solução:**
```go
// ✅ Sempre garantir que goroutine termine
func noLeak() {
    ch := make(chan int)
    go func() {
        val := <-ch
        fmt.Println(val)
    }()
    ch <- 42 // Envia valor
}
```

### 2. Race Condition

**Problema:**
```go
// ❌ Múltiplas goroutines modificando mesmo dado
var counter int
for i := 0; i < 1000; i++ {
    go func() {
        counter++ // Race condition!
    }()
}
```

**Solução 1: Mutex**
```go
// ✅ Proteger com mutex
var (
    counter int
    mu      sync.Mutex
)

for i := 0; i < 1000; i++ {
    go func() {
        mu.Lock()
        counter++
        mu.Unlock()
    }()
}
```

**Solução 2: Channel**
```go
// ✅ Usar channel
counterChan := make(chan int)
go func() {
    count := 0
    for range counterChan {
        count++
    }
}()

for i := 0; i < 1000; i++ {
    counterChan <- 1
}
```

### 3. Deadlock

**Problema:**
```go
// ❌ Todas goroutines esperando
ch := make(chan int)
ch <- 42 // Bloqueia para sempre (ninguém recebendo)
```

**Solução:**
```go
// ✅ Buffer ou goroutine separada
ch := make(chan int, 1) // Com buffer
ch <- 42

// OU
ch := make(chan int)
go func() {
    ch <- 42
}()
val := <-ch
```

---

## 🎓 Quando usar Goroutines?

### ✅ Use quando:

1. **I/O-bound operations**
   ```go
   // HTTP requests
   for _, url := range urls {
       go fetchURL(url)
   }
   ```

2. **Tarefas independentes**
   ```go
   // Processar múltiplos arquivos
   for _, file := range files {
       go processFile(file)
   }
   ```

3. **Servidores web**
   ```go
   // Uma goroutine por request
   http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
       // Cada request roda em sua goroutine
   })
   ```

### ❌ NÃO use quando:

1. **CPU-bound pequeno**
   ```go
   // ❌ Overhead maior que ganho
   go calculateSmallSum(1, 2, 3)
   ```

2. **Operações sequenciais obrigatórias**
   ```go
   // ❌ Etapa 2 depende de etapa 1
   go step1()
   go step2() // Precisa de resultado de step1
   ```

---

## 📝 Resumo Visual

```
┌─────────────────────────────────────────────┐
│         GOROUTINES EM UMA IMAGEM            │
├─────────────────────────────────────────────┤
│                                             │
│  main()                                     │
│    │                                        │
│    ├─ go func1() ──────┐                   │
│    │                   │                   │
│    ├─ go func2() ──────┼─────┐             │
│    │                   │     │             │
│    ├─ go func3() ──────┼─────┼─────┐       │
│    │                   │     │     │       │
│    │                   ▼     ▼     ▼       │
│    │              [Goroutines rodando]     │
│    │                   │     │     │       │
│    │                   │     │     │       │
│    └─ wg.Wait() ◄──────┴─────┴─────┘       │
│         (espera todas terminarem)           │
│                                             │
└─────────────────────────────────────────────┘
```

---

## 🎯 Exercícios Práticos

### Nível 1: Básico
```go
// 1. Crie 10 goroutines que imprimem números
// 2. Use WaitGroup para esperar todas terminarem
// 3. Observe a ordem de execução

- [Effective Go - Concurrency](https://golang.org/doc/effective_go#concurrency)
- [Go Blog - Concurrency Patterns](https://blog.golang.org/pipelines)
- [Go Tour - Concurrency](https://tour.golang.org/concurrency/1)

---
