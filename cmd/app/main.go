// Arquivo principal do programa (entrypoint) 🫡
// Convenção de mercado: colocar em cmd/<nome-app>/main.go
package main

// Importa os pacotes necessários
import (
	"fmt"

	// "github.com/seu-usuario/meu-projeto-go/internal/fibonacci"
	// "github.com/seu-usuario/meu-projeto-go/internal/hello"
	"meu-projeto-go/internal/fibonacci"
	"meu-projeto-go/internal/hello"
)

// Função principal do programa
func main() {
	// Mensagem inicial da aplicação
	fmt.Println("🚀 Meu primeiro projeto em Go com estrutura de mercado!")

	// Chamada para função de saudação
	hello.SayHello()

	// Demonstração: cálculo do 10º número de Fibonacci
	n := 10
	// Chama a função Fibonacci do pacote fibonacci
	// fibonacci // importado acima
	// Fibonacci(n) // retorna o n-ésimo número da sequência
	// := é usado para declarar e inicializar a variável
	valor := fibonacci.Fibonacci(n)
	// Imprime o resultado com formatação
	fmt.Printf("F(%d) = %d\n", n, valor)

	// Demonstração: imprimir a sequência completa até n
	fibonacci.PrintSequence(n)
}
