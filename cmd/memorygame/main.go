// Programa executável do jogo da memória (CLI).
// Uso:
//
//	go run ./cmd/memorygame
//
// Depois siga instruções na tela.
package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/seu-usuario/meu-projeto-go/internal/memorygame"
)

func main() {
	fmt.Println("🧠 Jogo da Memória - CLI em Go")
	fmt.Println("Encontre todos os pares! Formato de entrada: r1 c1 r2 c2 (ex: 0 0 1 0)")
	fmt.Println("Digite 'sair' para encerrar antecipadamente.")
	fmt.Println("")

	rows, cols := 4, 4 // Tabuleiro 4x4 -> 8 pares
	game, err := memorygame.NewGame(rows, cols)
	if err != nil {
		fmt.Println("Erro criando jogo:", err)
		return
	}

	reader := bufio.NewReader(os.Stdin)
	for {
		// Mostrar tabuleiro atual
		game.Render(false)
		if game.GameOver() {
			fmt.Println("🎉 Parabéns! Você concluiu o jogo!")
			fmt.Printf("Tentativas: %d | Tempo: %v\n", game.Moves, game.Elapsed().Round(time.Millisecond))
			fmt.Println("Tabuleiro final:")
			game.Render(true)
			break
		}

		fmt.Print("Sua jogada (r1 c1 r2 c2): ")
		line, _ := reader.ReadString('\n')
		line = strings.TrimSpace(line)
		if line == "" {
			continue
		}
		if strings.EqualFold(line, "sair") {
			fmt.Println("Saindo... até a próxima!")
			return
		}
		parts := strings.Fields(line)
		if len(parts) != 4 {
			fmt.Println("Entrada inválida. Use: r1 c1 r2 c2")
			continue
		}
		nums := make([]int, 4)
		ok := true
		for i, p := range parts {
			v, err := strconv.Atoi(p)
			if err != nil {
				fmt.Println("Valor não numérico:", p)
				ok = false
				break
			}
			nums[i] = v
		}
		if !ok {
			continue
		}

		matched, err := game.FlipPair(nums[0], nums[1], nums[2], nums[3])
		if err != nil {
			fmt.Println("Erro:", err)
			continue
		}
		// Mostrar cartas viradas
		game.Render(false)
		if matched {
			fmt.Println("✅ Par encontrado!")
		} else {
			fmt.Println("❌ Não foi par. Cartas serão ocultadas...")
			time.Sleep(1500 * time.Millisecond)
			game.HideNonMatched()
		}
	}
}
