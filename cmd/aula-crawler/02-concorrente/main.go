package main

import (
	"fmt"
	"net/http"
	"sync"
	"time"

	"github.com/PuerkitoBio/goquery"
)

/*
===========================================
EXEMPLO 2: CRAWLER CONCORRENTE (GOROUTINES)
===========================================

Este exemplo usa GOROUTINES para processar múltiplas URLs ao mesmo tempo!

🔥 CONCEITOS DE GOROUTINES:

1. O QUE SÃO GOROUTINES?
   - São "threads leves" do Go
   - Muito mais baratas que threads do sistema operacional
   - Você pode ter milhares rodando ao mesmo tempo
   - Gerenciadas automaticamente pelo Go runtime

2. COMO FUNCIONAM?
   - Thread normal: ~1-2MB de memória
   - Goroutine: ~2KB de memória (500x menor!)
   - O Go usa um "scheduler" que distribui goroutines entre threads

3. COMO USAR?
   - Basta colocar "go" antes de uma função
   - Exemplo: go minhaFuncao()
   - A função roda em paralelo, sem bloquear

4. SINCRONIZAÇÃO:
   - WaitGroup: espera goroutines terminarem
   - Channels: comunicação entre goroutines
   - Mutex: proteção de dados compartilhados
*/

type Result struct {
	URL   string
	Title string
	H1    string
	Error error
}

// fetchAndParse busca uma URL e extrai informações
func fetchAndParse(url string) Result {
	client := http.Client{
		Timeout: 10 * time.Second,
	}

	resp, err := client.Get(url)
	if err != nil {
		return Result{URL: url, Error: err}
	}
	defer resp.Body.Close()

	doc, err := goquery.NewDocumentFromReader(resp.Body)
	if err != nil {
		return Result{URL: url, Error: err}
	}

	title := doc.Find("title").First().Text()
	h1 := doc.Find("h1").First().Text()

	return Result{
		URL:   url,
		Title: title,
		H1:    h1,
	}
}

// worker é uma função que processa URLs de um canal
func worker(id int, jobs <-chan string, results chan<- Result, wg *sync.WaitGroup) {
	// defer garante que Done() será chamado quando a função terminar
	defer wg.Done()

	// Loop infinito: processa URLs enquanto houver
	for url := range jobs {
		fmt.Printf("🤖 Worker %d processando: %s\n", id, url)
		result := fetchAndParse(url)
		results <- result // Envia resultado para o canal
	}

	fmt.Printf("👋 Worker %d finalizou!\n", id)
}

func main() {
	fmt.Println("╔══════════════════════════════════════════╗")
	fmt.Println("║   WEB CRAWLER CONCORRENTE (GOROUTINES)   ║")
	fmt.Println("╚══════════════════════════════════════════╝")
	fmt.Println()

	urls := []string{
		"https://golang.org",
		"https://go.dev",
		"https://github.com",
		"https://stackoverflow.com",
		"https://www.reddit.com",
		"https://news.ycombinator.com",
		"https://medium.com",
		"https://dev.to",
	}

	// ═══════════════════════════════════════
	// PASSO 1: CRIAR OS CANAIS (CHANNELS)
	// ═══════════════════════════════════════
	// Canais são como "tubos" para passar dados entre goroutines

	// Canal de jobs: envia URLs para os workers
	jobs := make(chan string, len(urls))

	// Canal de results: recebe resultados dos workers
	results := make(chan Result, len(urls))

	// ═══════════════════════════════════════
	// PASSO 2: CRIAR WAITGROUP
	// ═══════════════════════════════════════
	// WaitGroup conta quantas goroutines ainda estão rodando
	var wg sync.WaitGroup

	// ═══════════════════════════════════════
	// PASSO 3: CRIAR OS WORKERS (GOROUTINES)
	// ═══════════════════════════════════════
	workerCount := 4 // Número de workers simultâneos
	fmt.Printf("🚀 Iniciando %d workers...\n\n", workerCount)

	startTime := time.Now()

	// Cria e inicia os workers
	for i := 1; i <= workerCount; i++ {
		wg.Add(1) // Incrementa o contador do WaitGroup
		go worker(i, jobs, results, &wg) // 🔥 AQUI ESTÁ A MÁGICA!
		// A palavra "go" faz a função rodar em paralelo!
	}

	// ═══════════════════════════════════════
	// PASSO 4: ENVIAR JOBS PARA OS WORKERS
	// ═══════════════════════════════════════
	for _, url := range urls {
		jobs <- url // Envia URL para o canal
	}
	close(jobs) // Fecha o canal (não haverá mais jobs)

	// ═══════════════════════════════════════
	// PASSO 5: COLETAR RESULTADOS
	// ═══════════════════════════════════════
	// Esta goroutine fecha o canal de results quando todos os workers terminarem
	go func() {
		wg.Wait()      // Espera todos os workers terminarem
		close(results) // Fecha o canal de results
	}()

	// ═══════════════════════════════════════
	// PASSO 6: PROCESSAR RESULTADOS
	// ═══════════════════════════════════════
	fmt.Println("\n📊 RESULTADOS:")
	fmt.Println("═══════════════════════════════════════════════")

	successCount := 0
	errorCount := 0

	// Range em canal: itera até o canal ser fechado
	for result := range results {
		if result.Error != nil {
			fmt.Printf("❌ %s\n   Erro: %v\n\n", result.URL, result.Error)
			errorCount++
		} else {
			fmt.Printf("✅ %s\n", result.URL)
			fmt.Printf("   📄 Título: %s\n", result.Title)
			fmt.Printf("   📌 H1: %s\n\n", result.H1)
			successCount++
		}
	}

	elapsed := time.Since(startTime)

	// ═══════════════════════════════════════
	// ESTATÍSTICAS
	// ═══════════════════════════════════════
	fmt.Println("╔══════════════════════════════════════════╗")
	fmt.Printf("║   TEMPO TOTAL: %-24s  ║\n", elapsed)
	fmt.Printf("║   URLs PROCESSADAS: %-19d  ║\n", len(urls))
	fmt.Printf("║   Workers: %-28d  ║\n", workerCount)
	fmt.Printf("║   Sucessos: %-27d  ║\n", successCount)
	fmt.Printf("║   Erros: %-30d  ║\n", errorCount)
	fmt.Println("╚══════════════════════════════════════════╝")
	fmt.Println()

	// ═══════════════════════════════════════
	// EXPLICAÇÃO DA VELOCIDADE
	// ═══════════════════════════════════════
	fmt.Println("💡 POR QUE É MAIS RÁPIDO?")
	fmt.Println("   Exemplo com 4 workers e 8 URLs:")
	fmt.Println()
	fmt.Println("   ⏱️  Síncrono (01-basico):")
	fmt.Println("      URL1 → URL2 → URL3 → URL4 → URL5 → URL6 → URL7 → URL8")
	fmt.Println("      Se cada uma leva 2s = 16 segundos total")
	fmt.Println()
	fmt.Println("   🚀 Concorrente (4 workers):")
	fmt.Println("      Worker 1: URL1 → URL5")
	fmt.Println("      Worker 2: URL2 → URL6")
	fmt.Println("      Worker 3: URL3 → URL7")
	fmt.Println("      Worker 4: URL4 → URL8")
	fmt.Println("      Todas rodando ao mesmo tempo = ~4 segundos!")
	fmt.Println()
	fmt.Println("   📈 GANHO: ~4x mais rápido com 4 workers!")
	fmt.Println()
	fmt.Println("⚠️  PRÓXIMO PASSO:")
	fmt.Println("    Adicionar rate limiting para não sobrecarregar servidores")
	fmt.Println("    Execute: go run cmd/aula-crawler/03-completo/main.go")
}
