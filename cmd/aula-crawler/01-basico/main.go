package main

import (
	"fmt"
	"net/http"
	"time"

	"github.com/PuerkitoBio/goquery"
)

/*
===========================================
EXEMPLO 1: CRAWLER BÁSICO (SÍNCRONO)
===========================================

Este é o exemplo mais simples de um web crawler.
Ele busca páginas uma por vez, de forma sequencial.

CONCEITOS IMPORTANTES:
- HTTP Client: Faz requisições HTTP
- Timeout: Define tempo máximo de espera
- HTML Parsing: Extrai dados do HTML
- Error Handling: Tratamento de erros
*/

// Result representa o resultado de uma requisição
type Result struct {
	URL   string // URL que foi acessada
	Title string // Título da página (<title>)
	H1    string // Primeiro H1 encontrado
	Error error  // Erro (se houver)
}

// fetchAndParse busca uma URL e extrai informações
func fetchAndParse(url string) Result {
	fmt.Printf("🔍 Buscando: %s\n", url)

	// 1. CRIANDO O CLIENT HTTP
	// Timeout: se não responder em 10 segundos, cancela
	client := http.Client{
		Timeout: 10 * time.Second,
	}

	// 2. FAZENDO A REQUISIÇÃO GET
	resp, err := client.Get(url)
	if err != nil {
		return Result{URL: url, Error: err}
	}
	defer resp.Body.Close() // IMPORTANTE: sempre fechar o body

	// 3. PARSEANDO O HTML
	// goquery é como jQuery para Go
	doc, err := goquery.NewDocumentFromReader(resp.Body)
	if err != nil {
		return Result{URL: url, Error: err}
	}

	// 4. EXTRAINDO DADOS
	// Find() busca elementos CSS, como em jQuery
	title := doc.Find("title").First().Text()
	h1 := doc.Find("h1").First().Text()

	return Result{
		URL:   url,
		Title: title,
		H1:    h1,
	}
}

func main() {
	fmt.Println("╔══════════════════════════════════════════╗")
	fmt.Println("║   WEB CRAWLER BÁSICO (SÍNCRONO)          ║")
	fmt.Println("╚══════════════════════════════════════════╝")
	fmt.Println()

	// Lista de URLs para crawlear
	urls := []string{
		"https://golang.org",
		"https://go.dev",
		"https://github.com",
		"https://stackoverflow.com",
	}

	// Marca o tempo de início
	startTime := time.Now()

	// PROCESSAMENTO SÍNCRONO
	// Uma URL por vez, esperando cada uma terminar
	for _, url := range urls {
		result := fetchAndParse(url)

		if result.Error != nil {
			fmt.Printf("❌ Erro ao buscar %s: %v\n\n", result.URL, result.Error)
		} else {
			fmt.Printf("✅ URL: %s\n", result.URL)
			fmt.Printf("   📄 Título: %s\n", result.Title)
			fmt.Printf("   📌 H1: %s\n\n", result.H1)
		}
	}

	// Calcula o tempo total
	elapsed := time.Since(startTime)

	fmt.Println("╔══════════════════════════════════════════╗")
	fmt.Printf("║   TEMPO TOTAL: %-24s  ║\n", elapsed)
	fmt.Printf("║   URLs PROCESSADAS: %-19d  ║\n", len(urls))
	fmt.Println("╚══════════════════════════════════════════╝")
	fmt.Println()
	fmt.Println("⚠️  PROBLEMA: Este crawler é LENTO!")
	fmt.Println("    Processa uma URL por vez, esperando cada uma terminar.")
	fmt.Println("    Para 4 URLs com 2s cada = 8 segundos totais")
	fmt.Println()
	fmt.Println("💡 SOLUÇÃO: Use concorrência com Goroutines!")
	fmt.Println("    Execute: go run cmd/aula-crawler/02-concorrente/main.go")
}
