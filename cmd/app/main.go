// Arquivo principal do programa (entrypoint) 🫡
// Convenção de mercado: colocar em cmd/<nome-app>/main.go
package main

// Importa os pacotes necessários
import (
	"bufio"
	"fmt"
	"os"
	"strings"

	// Mantendo todos os seus pacotes internos
	"github.com/seu-usuario/meu-projeto-go/internal/fibonacci"
	"github.com/seu-usuario/meu-projeto-go/internal/hello"
	"github.com/seu-usuario/meu-projeto-go/internal/memoriago"
)

// Função principal do programa
func main() {
	reader := bufio.NewReader(os.Stdin)

	// Loop infinito para manter o menu ativo até o usuário decidir sair
	for {
		// Limpa a tela para uma melhor experiência de menu
		fmt.Print("\033[H\033[2J")

		fmt.Println("🚀 Meu Projeto em Go 🚀")
		fmt.Println("---------------------------")
		fmt.Println("1. Jogo da Memória")
		fmt.Println("2. Demonstração Fibonacci")
		fmt.Println("3. Dizer Olá")
		fmt.Println("---------------------------")
		fmt.Println("Digite 'sair' para terminar.")
		fmt.Print("Escolha uma opção: ")

		// Lê a escolha do usuário
		escolha, _ := reader.ReadString('\n')
		escolha = strings.TrimSpace(escolha)

		// Executa a ação baseada na escolha
		switch escolha {
		case "1":
			// Chama o Jogo da Memória
			memoriago.Play()
		case "2":
			// Executa a demonstração do Fibonacci que você já tinha
			fmt.Println("\n--- Executando Demonstração Fibonacci ---")
			n := 10
			valor := fibonacci.Fibonacci(n)
			fmt.Printf("F(%d) = %d\n", n, valor)
			fibonacci.PrintSequence(n)
			fmt.Println("\nPressione ENTER para voltar ao menu...")
			_, _ = reader.ReadString('\n')
		case "3":
			// Executa a saudação
			fmt.Println("\n--- Executando Saudação ---")
			hello.SayHello()
			fmt.Println("\nPressione ENTER para voltar ao menu...")
			_, _ = reader.ReadString('\n')
		case "sair":
			fmt.Println("Até a próxima!")
			return // Encerra o programa
		default:
			fmt.Println("Opção inválida. Pressione ENTER para tentar novamente.")
			_, _ = reader.ReadString('\n')
		}
	}
}
