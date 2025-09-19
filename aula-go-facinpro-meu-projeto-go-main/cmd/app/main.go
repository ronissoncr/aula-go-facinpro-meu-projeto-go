// Arquivo principal do programa (entrypoint)
// Convenção de mercado: colocar em cmd/<nome-app>/main.go
package main

// Importa os pacotes necessários
import (
	"aulago/internal/anamnese"
	"aulago/internal/fibonacci"
	"aulago/internal/hello"
	"fmt"
)

// Função principal do programa
func main() {
	fmt.Println("🚀 Meu primeiro projeto em Go com estrutura de mercado!")
	hello.SayHello()
	fmt.Println("----")
	// --- Parte 2: Fibonacci ---
	// Define qual número da sequência queremos
	posicao := 10

	// Chama a função Fibonacci do nosso pacote e guarda o resultado
	resultadoFib := fibonacci.Fibonacci(posicao)

	// Imprime o resultado formatado
	fmt.Printf("O %dº número na sequência de Fibonacci é: %d\n", posicao, resultadoFib)

	// func anamnese

	nome := "Keven"
	idade := 28
	peso := 84.0
	altura := 1.75

	fmt.Println("Funcção Completa")
	anamnese.ExecutarAnamnese(nome, idade, peso, altura)
}
