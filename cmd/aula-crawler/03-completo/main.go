package main

import (
	"context"
	"encoding/csv"
	"fmt"
	"net/http"
	"net/url"
	"os"
	"sync"
	"time"

	"github.com/PuerkitoBio/goquery"
	"golang.org/x/time/rate"
)

/*
===========================================
EXEMPLO 3: CRAWLER COMPLETO (PROFISSIONAL)
===========================================

Este é um crawler de nível profissional com:
✅ Concorrência (Goroutines)
✅ Rate Limiting (controle de requisições por domínio)
✅ Exportação para CSV
✅ Tratamento robusto de erros
✅ Estatísticas detalhadas

🔥 NOVO CONCEITO: RATE LIMITING

O que é Rate Limiting?
- Limita a quantidade de requisições por tempo
- Evita sobrecarregar servidores
- Previne bloqueios de IP
- Respeita políticas de uso dos sites

Como funciona?
- Usamos golang.org/x/time/rate
- Cada domínio tem seu próprio limiter
- Exemplo: máximo 2 requisições por segundo por domínio
*/

// ═══════════════════════════════════════
// ESTRUTURAS DE DADOS
// ═══════════════════════════════════════

type Result struct {
	URL        string
	Title      string
	H1         string
	H2Count    int
	StatusCode int
	Duration   time.Duration
	Error      error
}

// DomainLimiter gerencia rate limiting por domínio
type DomainLimiter struct {
	limiters map[string]*rate.Limiter
	mu       sync.Mutex
}

// NewDomainLimiter cria um novo gerenciador de limiters
func NewDomainLimiter() *DomainLimiter {
	return &DomainLimiter{
		limiters: make(map[string]*rate.Limiter),
	}
}

// GetLimiter retorna ou cria um limiter para um domínio
func (dl *DomainLimiter) GetLimiter(domain string) *rate.Limiter {
	dl.mu.Lock()
	defer dl.mu.Unlock()

	// Se já existe limiter para este domínio, retorna
	if limiter, ok := dl.limiters[domain]; ok {
		return limiter
	}

	// Cria novo limiter: 2 requisições por segundo, burst de 5
	// rate.Every(500*time.Millisecond) = 1 req a cada 500ms = 2 req/s
	limiter := rate.NewLimiter(rate.Every(500*time.Millisecond), 5)
	dl.limiters[domain] = limiter

	fmt.Printf("🔧 Criado rate limiter para domínio: %s (2 req/s)\n", domain)
	return limiter
}

// ═══════════════════════════════════════
// FUNÇÃO PRINCIPAL DE CRAWLING
// ═══════════════════════════════════════

func fetchAndParse(urlStr string, dl *DomainLimiter) Result {
	startTime := time.Now()

	// 1. PARSE DA URL
	parsedURL, err := url.Parse(urlStr)
	if err != nil {
		return Result{URL: urlStr, Error: err}
	}

	// 2. RATE LIMITING
	// Obtém o limiter para este domínio
	limiter := dl.GetLimiter(parsedURL.Host)

	// Espera até ter permissão para fazer a requisição
	fmt.Printf("⏳ Aguardando rate limit para %s...\n", parsedURL.Host)
	err = limiter.Wait(context.Background())
	if err != nil {
		return Result{URL: urlStr, Error: err}
	}

	// 3. REQUISIÇÃO HTTP
	client := http.Client{
		Timeout: 10 * time.Second,
	}

	fmt.Printf("🔍 Buscando: %s\n", urlStr)
	resp, err := client.Get(urlStr)
	if err != nil {
		return Result{URL: urlStr, Error: err, Duration: time.Since(startTime)}
	}
	defer resp.Body.Close()

	// 4. PARSING DO HTML
	doc, err := goquery.NewDocumentFromReader(resp.Body)
	if err != nil {
		return Result{URL: urlStr, Error: err, StatusCode: resp.StatusCode, Duration: time.Since(startTime)}
	}

	// 5. EXTRAÇÃO DE DADOS
	title := doc.Find("title").First().Text()
	h1 := doc.Find("h1").First().Text()
	h2Count := doc.Find("h2").Length()

	duration := time.Since(startTime)

	return Result{
		URL:        urlStr,
		Title:      title,
		H1:         h1,
		H2Count:    h2Count,
		StatusCode: resp.StatusCode,
		Duration:   duration,
	}
}

// ═══════════════════════════════════════
// WORKER COM RATE LIMITING
// ═══════════════════════════════════════

func worker(id int, jobs <-chan string, results chan<- Result, dl *DomainLimiter, wg *sync.WaitGroup) {
	defer wg.Done()

	for urlStr := range jobs {
		fmt.Printf("🤖 Worker %d processando: %s\n", id, urlStr)
		result := fetchAndParse(urlStr, dl)
		results <- result
	}

	fmt.Printf("👋 Worker %d finalizou!\n", id)
}

// ═══════════════════════════════════════
// EXPORTAÇÃO PARA CSV
// ═══════════════════════════════════════

func exportToCSV(results []Result, filename string) error {
	file, err := os.Create(filename)
	if err != nil {
		return err
	}
	defer file.Close()

	writer := csv.NewWriter(file)
	defer writer.Flush()

	// Cabeçalho
	header := []string{"URL", "Title", "H1", "H2_Count", "Status_Code", "Duration_ms", "Error"}
	if err := writer.Write(header); err != nil {
		return err
	}

	// Dados
	for _, r := range results {
		errorStr := ""
		if r.Error != nil {
			errorStr = r.Error.Error()
		}

		row := []string{
			r.URL,
			r.Title,
			r.H1,
			fmt.Sprintf("%d", r.H2Count),
			fmt.Sprintf("%d", r.StatusCode),
			fmt.Sprintf("%.0f", r.Duration.Seconds()*1000),
			errorStr,
		}

		if err := writer.Write(row); err != nil {
			return err
		}
	}

	return nil
}

// ═══════════════════════════════════════
// FUNÇÃO MAIN
// ═══════════════════════════════════════

func main() {
	fmt.Println("╔════════════════════════════════════════════════════╗")
	fmt.Println("║   WEB CRAWLER PROFISSIONAL (COM RATE LIMITING)    ║")
	fmt.Println("╚════════════════════════════════════════════════════╝")
	fmt.Println()

	// URLs de diferentes domínios
	urls := []string{
		// Mesmo domínio - será limitado
		"https://golang.org",
		"https://golang.org/doc",
		"https://golang.org/pkg",

		// Mesmo domínio - será limitado
		"https://go.dev",
		"https://go.dev/learn",
		"https://go.dev/solutions",

		// Domínios diferentes
		"https://github.com/golang",
		"https://stackoverflow.com/questions/tagged/go",
		"https://pkg.go.dev",
		"https://play.golang.org",
	}

	// ═══════════════════════════════════════
	// CONFIGURAÇÃO
	// ═══════════════════════════════════════
	workerCount := 5
	domainLimiter := NewDomainLimiter()

	jobs := make(chan string, len(urls))
	results := make(chan Result, len(urls))
	var wg sync.WaitGroup

	fmt.Printf("🚀 Iniciando %d workers...\n", workerCount)
	fmt.Printf("📊 URLs para processar: %d\n", len(urls))
	fmt.Printf("⚡ Rate limit: 2 requisições/segundo por domínio\n\n")

	startTime := time.Now()

	// ═══════════════════════════════════════
	// INICIAR WORKERS
	// ═══════════════════════════════════════
	for i := 1; i <= workerCount; i++ {
		wg.Add(1)
		go worker(i, jobs, results, domainLimiter, &wg)
	}

	// ═══════════════════════════════════════
	// ENVIAR JOBS
	// ═══════════════════════════════════════
	for _, url := range urls {
		jobs <- url
	}
	close(jobs)

	// ═══════════════════════════════════════
	// COLETAR RESULTADOS
	// ═══════════════════════════════════════
	go func() {
		wg.Wait()
		close(results)
	}()

	// ═══════════════════════════════════════
	// PROCESSAR RESULTADOS
	// ═══════════════════════════════════════
	var allResults []Result
	successCount := 0
	errorCount := 0
	totalDuration := time.Duration(0)

	fmt.Println("\n📊 RESULTADOS:")
	fmt.Println("═══════════════════════════════════════════════════════")

	for result := range results {
		allResults = append(allResults, result)

		if result.Error != nil {
			fmt.Printf("❌ %s\n   Erro: %v\n   Duração: %v\n\n",
				result.URL, result.Error, result.Duration)
			errorCount++
		} else {
			fmt.Printf("✅ %s\n", result.URL)
			fmt.Printf("   📄 Título: %s\n", result.Title)
			fmt.Printf("   📌 H1: %s\n", result.H1)
			fmt.Printf("   🔢 H2s: %d\n", result.H2Count)
			fmt.Printf("   🌐 Status: %d\n", result.StatusCode)
			fmt.Printf("   ⏱️  Duração: %v\n\n", result.Duration)
			successCount++
			totalDuration += result.Duration
		}
	}

	elapsed := time.Since(startTime)

	// ═══════════════════════════════════════
	// EXPORTAR PARA CSV
	// ═══════════════════════════════════════
	csvFilename := "crawler_results.csv"
	if err := exportToCSV(allResults, csvFilename); err != nil {
		fmt.Printf("❌ Erro ao exportar CSV: %v\n", err)
	} else {
		fmt.Printf("💾 Resultados exportados para: %s\n\n", csvFilename)
	}

	// ═══════════════════════════════════════
	// ESTATÍSTICAS FINAIS
	// ═══════════════════════════════════════
	avgDuration := time.Duration(0)
	if successCount > 0 {
		avgDuration = totalDuration / time.Duration(successCount)
	}

	fmt.Println("╔════════════════════════════════════════════════════╗")
	fmt.Printf("║   TEMPO TOTAL: %-35s  ║\n", elapsed)
	fmt.Printf("║   URLs PROCESSADAS: %-30d  ║\n", len(urls))
	fmt.Printf("║   Workers: %-39d  ║\n", workerCount)
	fmt.Printf("║   Sucessos: %-38d  ║\n", successCount)
	fmt.Printf("║   Erros: %-41d  ║\n", errorCount)
	fmt.Printf("║   Duração média por URL: %-25s  ║\n", avgDuration)
	fmt.Println("╚════════════════════════════════════════════════════╝")
	fmt.Println()

	// ═══════════════════════════════════════
	// EXPLICAÇÃO DO RATE LIMITING
	// ═══════════════════════════════════════
	fmt.Println("💡 COMO O RATE LIMITING FUNCIONOU?")
	fmt.Println("═══════════════════════════════════════════════════════")
	fmt.Println()
	fmt.Println("🌐 DOMÍNIOS PROCESSADOS:")
	fmt.Println("   golang.org: 3 URLs (limitadas a 2 req/s)")
	fmt.Println("   go.dev: 3 URLs (limitadas a 2 req/s)")
	fmt.Println("   Outros: 4 URLs (cada um com seu limiter)")
	fmt.Println()
	fmt.Println("⏱️  SEM RATE LIMIT:")
	fmt.Println("   - Risco de IP bloqueado")
	fmt.Println("   - Possível sobrecarga do servidor")
	fmt.Println("   - Comportamento antiético")
	fmt.Println()
	fmt.Println("✅ COM RATE LIMIT:")
	fmt.Println("   - Respeita o servidor")
	fmt.Println("   - Evita bloqueios")
	fmt.Println("   - Comportamento profissional")
	fmt.Println()
	fmt.Println("📈 BENEFÍCIOS DA CONCORRÊNCIA + RATE LIMITING:")
	fmt.Println("   - Processa múltiplos domínios em paralelo")
	fmt.Println("   - Respeita limites de cada domínio")
	fmt.Println("   - Máxima velocidade com responsabilidade")
	fmt.Println()
	fmt.Println("🎯 CASOS DE USO REAIS:")
	fmt.Println("   ✓ Web scraping de produtos")
	fmt.Println("   ✓ Monitoramento de SEO")
	fmt.Println("   ✓ Agregação de notícias")
	fmt.Println("   ✓ Análise de competidores")
	fmt.Println("   ✓ Indexação de conteúdo")
}
