# Jogo da Memória em Go (iniciante)

**Objetivo:** o programa mostra uma sequência de **5 números**. Eu memorizo, aperto **ENTER**, a tela “limpa”, eu digito os 5 números **na ordem**, e o código confere e **pontua** (0–5).

## 1) Pré‑requisitos

- Ter o Go instalado (qualquer versão recente).
- Saber abrir um terminal.

## 2) Crie o arquivo

Crie uma pasta (por exemplo, `memoria-go` em internal) e dentro dela um arquivos`games.go`.

## 3) Cole o código completo

```go
package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"strconv"
	"strings"
	"time"
)

func limparTela() {
	// Tenta limpar via ANSI + garante espaço em branco como fallback
	fmt.Print("\033[H\033[2J")
	fmt.Print(strings.Repeat("\n", 10))
}

func gerarSequencia(tamanho, digitoMax int) []int {
	seq := make([]int, tamanho)
	for i := 0; i < tamanho; i++ {
		seq[i] = rand.Intn(digitoMax + 1) // 0..digitoMax (ex.: 0..9)
	}
	return seq
}

func juntar(nums []int) string {
	partes := make([]string, len(nums))
	for i, v := range nums {
		partes[i] = strconv.Itoa(v)
	}
	return strings.Join(partes, " ")
}

func lerPalpite(reader *bufio.Reader, tamanho int) []int {
	for {
		fmt.Printf("Digite a sequência (%d números de 0 a 9, separados por espaço): ", tamanho)
		linha, _ := reader.ReadString('\n')
		linha = strings.TrimSpace(linha)

		partes := strings.Fields(linha)
		if len(partes) != tamanho {
			fmt.Printf("Ops! São exatamente %d números. Tenta de novo.\n", tamanho)
			continue
		}

		palpite := make([]int, tamanho)
		ok := true
		for i, p := range partes {
			v, err := strconv.Atoi(p)
			if err != nil || v < 0 || v > 9 {
				fmt.Println("Use apenas dígitos de 0 a 9.")
				ok = false
				break
			}
			palpite[i] = v
		}
		if ok {
			return palpite
		}
	}
}

func pontuar(seq, palpite []int) int {
	acertos := 0
	for i := range seq {
		if seq[i] == palpite[i] {
			acertos++
		}
	}
	return acertos
}

func main() {
	rand.Seed(time.Now().UnixNano())

	const tamanho = 5
	const digitoMax = 9

	reader := bufio.NewReader(os.Stdin)

	fmt.Println("=== Jogo da Memória (Go) ===")
	fmt.Println("Objetivo: memorize a sequência e digite na mesma ordem.")
	fmt.Println()

	seq := gerarSequencia(tamanho, digitoMax)

	// Mostra a sequência para memorizar
	fmt.Println("Sequência (memorize):", juntar(seq))
	fmt.Println("Pressione ENTER quando estiver pronto para responder...")
	_, _ = reader.ReadString('\n')

	// "Limpa" a tela (ou empurra para cima)
	limparTela()

	// Pede o palpite
	palpite := lerPalpite(reader, tamanho)

	// Calcula pontuação
	acertos := pontuar(seq, palpite)

	// Mostra resultado
	fmt.Println()
	fmt.Println("Sequência correta:", juntar(seq))
	fmt.Println("Sua resposta     :", juntar(palpite))
	fmt.Printf("Você fez %d/%d ponto(s). ", acertos, tamanho)

	switch acertos {
	case 5:
		fmt.Println("Perfeito! Mandou muito bem!")
	case 3, 4:
		fmt.Println("Quase lá! Bora de novo?")
	case 1, 2:
		fmt.Println("Tá no caminho. Treina mais um pouco!")
	default:
		fmt.Println("Agora vai! Tenta outra vez!")
	}
}

```

## 4) Como executar

No terminal, dentro da pasta do projeto:

```bash
go run ./cmd/app

```

## 5) Exemplo de execução

```
=== Jogo da Memória (Go) ===
Objetivo: memorize a sequência e digite na mesma ordem.

Sequência (memorize): 3 8 1 6 2
Pressione ENTER quando estiver pronto para responder...

Digite a sequência (5 números de 0 a 9, separados por espaço): 3 8 1 6 9

Sequência correta: 3 8 1 6 2
Sua resposta     : 3 8 1 6 9
Você fez 4/5 ponto(s). Quase lá! Bora de novo?
```