// Arquivo principal do programa (entrypoint) 🫡
// Convenção de mercado: colocar em cmd/<nome-app>/main.go
package main

// Importa os pacotes necessários
import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"

	// Mantendo todos os seus pacotes internos

	calculator "github.com/seu-usuario/meu-projeto-go/internal/calcular-imc"
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
		fmt.Println("4. Demonstração de Funções (Saudação e Anônima)") // <-- NOVA OPÇÃO
		fmt.Println("Calcular o IMC:")
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
		case "4": // <-- LÓGICA DA ATIVIDADE MOVIDA PARA CÁ
			fmt.Println("\n--- Demonstração de Funções (Quadro Branco) ---")
			// Parte 1: Usando a função nomeada do pacote 'hello'
			mensagemNomeada := hello.Saudacao("Galera da Aula")
			fmt.Println("Função Nomeada:", mensagemNomeada)

			// Parte 2: Usando uma Função Anônima
			mensagemAnonima := func(nome string) string {
				return fmt.Sprintf("Até mais, %s!", nome)
			}("Mundo")
			fmt.Println("Função Anônima:", mensagemAnonima)

			fmt.Println("\nPressione ENTER para voltar ao menu...")
			_, _ = reader.ReadString('\n')
		case "5":
			fmt.Println("\n--- Calculadora de IMC ---")

			// Pede o peso
			fmt.Print("Digite seu peso em kg (ex: 82.5): ")
			pesoStr, _ := reader.ReadString('\n')
			peso, err := strconv.ParseFloat(strings.TrimSpace(pesoStr), 64)
			if err != nil {
				fmt.Println("Valor de peso inválido. Pressione ENTER para tentar novamente.")
				_, _ = reader.ReadString('\n')
				continue // Volta para o início do menu
			}

			// Pede a altura
			fmt.Print("Digite sua altura em metros (ex: 1.79): ")
			alturaStr, _ := reader.ReadString('\n')
			altura, err := strconv.ParseFloat(strings.TrimSpace(alturaStr), 64)
			if err != nil {
				fmt.Println("Valor de altura inválido. Pressione ENTER para tentar novamente.")
				_, _ = reader.ReadString('\n')
				continue // Volta para o início do menu
			}

			// Chama a função do pacote 'calculator'
			imc := calculator.CalcularIMC(peso, altura)

			// Exibe o resultado
			fmt.Printf("\nSeu IMC é: %.2f\n", imc)

			fmt.Println("\nPressione ENTER para voltar ao menu...")
			_, _ = reader.ReadString('\n')
		// --- FIM DO NOVO BLOCO ---
		case "sair":
			fmt.Println("Até a próxima!")
			return // Encerra o programa
		default:
			fmt.Println("Opção inválida. Pressione ENTER para tentar novamente.")
			_, _ = reader.ReadString('\n')
		}
	}
}
