// Package memorygame implementa um simples jogo da memória em modo texto.
// Objetivo: encontrar todos os pares de cartas iguais com o menor número de tentativas.
//
// Conceitos didáticos cobertos:
// - Estruturas (struct)
// - Slices e matrizes (slice de slices)
// - Encapsulamento via funções e métodos
// - Aleatoriedade (embaralhar)
// - Leitura de entrada do usuário
// - Controle de estado do jogo
package memorygame

import (
	"errors"
	"fmt"
	"math/rand"
	"time"
)

// Card representa uma carta no tabuleiro.
type Card struct {
	Value    rune // valor exibido quando revelada (ex: 'A', 'B', ...)
	Revealed bool // se a carta está temporariamente voltada para cima
	Matched  bool // se já foi encontrada (faz parte de um par resolvido)
}

// Game contém todo o estado do jogo.
type Game struct {
	Rows, Cols int
	Board      [][]*Card
	Moves      int       // número de tentativas (pares virados)
	PairsFound int       // pares corretos encontrados
	TotalPairs int       // total de pares no jogo
	startTime  time.Time // usado para calcular duração
}

// NewGame cria um novo jogo com linhas/colunas dados.
// O número total de cartas (rows*cols) precisa ser par.
func NewGame(rows, cols int) (*Game, error) {
	if rows <= 0 || cols <= 0 {
		return nil, errors.New("linhas e colunas devem ser > 0")
	}
	total := rows * cols
	if total%2 != 0 {
		return nil, errors.New("o total de cartas deve ser par para formar pares")
	}

	g := &Game{Rows: rows, Cols: cols, TotalPairs: total / 2, startTime: time.Now()}
	g.initBoard()
	return g, nil
}

// initBoard gera os valores das cartas em pares e embaralha.
func (g *Game) initBoard() {
	// Gerar lista de runes para os pares (A, B, C ...)
	needed := g.TotalPairs
	values := make([]rune, 0, needed)
	for i := 0; i < needed; i++ {
		// Começa em 'A' e avança
		values = append(values, rune('A'+i))
	}
	// Duplicar para formar pares
	all := make([]rune, 0, needed*2)
	for _, v := range values {
		all = append(all, v, v)
	}
	// Embaralhar
	rand.Seed(time.Now().UnixNano())
	rand.Shuffle(len(all), func(i, j int) { all[i], all[j] = all[j], all[i] })

	// Preencher matriz
	g.Board = make([][]*Card, g.Rows)
	k := 0
	for r := 0; r < g.Rows; r++ {
		row := make([]*Card, g.Cols)
		for c := 0; c < g.Cols; c++ {
			row[c] = &Card{Value: all[k]}
			k++
		}
		g.Board[r] = row
	}
}

// InBounds verifica se posição é válida.
func (g *Game) InBounds(r, c int) bool {
	return r >= 0 && r < g.Rows && c >= 0 && c < g.Cols
}

// FlipPair tenta virar duas cartas.
// Retorna:
// - matched = true se formaram um par
// - erro em caso de coordenada inválida ou carta já revelada/match
func (g *Game) FlipPair(r1, c1, r2, c2 int) (matched bool, err error) {
	if !g.InBounds(r1, c1) || !g.InBounds(r2, c2) {
		return false, errors.New("coordenadas fora do tabuleiro")
	}
	if r1 == r2 && c1 == c2 {
		return false, errors.New("selecione cartas diferentes")
	}
	a := g.Board[r1][c1]
	b := g.Board[r2][c2]
	if a.Matched || b.Matched {
		return false, errors.New("uma das cartas já foi resolvida")
	}
	if a.Revealed || b.Revealed {
		return false, errors.New("uma das cartas já está revelada nesta jogada")
	}
	// Revelar temporariamente
	a.Revealed = true
	b.Revealed = true
	g.Moves++

	if a.Value == b.Value {
		a.Matched, b.Matched = true, true
		g.PairsFound++
		return true, nil
	}
	return false, nil
}

// HideNonMatched oculta cartas reveladas que não formaram par.
func (g *Game) HideNonMatched() {
	for r := 0; r < g.Rows; r++ {
		for c := 0; c < g.Cols; c++ {
			card := g.Board[r][c]
			if card.Revealed && !card.Matched {
				card.Revealed = false
			}
		}
	}
}

// GameOver retorna true se todos os pares foram encontrados.
func (g *Game) GameOver() bool {
	return g.PairsFound == g.TotalPairs
}

// Elapsed retorna duração em segundos.
func (g *Game) Elapsed() time.Duration {
	return time.Since(g.startTime)
}

// Render imprime o tabuleiro.
// showAll: força exibir todos os valores (debug / final).
func (g *Game) Render(showAll bool) {
	fmt.Print("    ")
	for c := 0; c < g.Cols; c++ {
		fmt.Printf("%2d ", c)
	}
	fmt.Println()
	fmt.Println("   ", repeat("──", g.Cols))
	for r := 0; r < g.Rows; r++ {
		fmt.Printf("%2d │", r)
		for c := 0; c < g.Cols; c++ {
			card := g.Board[r][c]
			ch := "?"
			if showAll || card.Revealed || card.Matched {
				ch = string(card.Value)
			}
			fmt.Printf(" %s ", ch)
		}
		fmt.Println()
	}
	fmt.Println()
}

func repeat(s string, n int) string {
	out := ""
	for i := 0; i < n; i++ {
		out += s
	}
	return out
}
