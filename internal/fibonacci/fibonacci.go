// Pacote fibonacci: implementa funções relacionadas à sequência de Fibonacci.
// A sequência de Fibonacci é definida como:
// F(0) = 0, F(1) = 1 e F(n) = F(n-1) + F(n-2) para n >= 2.
// Esta implementação usa abordagem iterativa (O(n) tempo, O(1) memória para o n-ésimo termo)
// e uma função auxiliar para gerar a lista até um certo n.
package fibonacci

// Importa o pacote fmt para formatação de strings e impressão (necessário para PrintSequence)
import "fmt"

// Fibonacci retorna o n-ésimo número da sequência de Fibonacci.
// Caso n seja negativo, ocorre pânico (panic) pois é entrada inválida.
// Complexidade: tempo O(n), espaço O(1).
func Fibonacci(n int) int {
	if n < 0 {
		panic("n não pode ser negativo")
	}
	if n == 0 {
		return 0
	}
	if n == 1 {
		return 1
	}
	// a = F(i-2), b = F(i-1)
	a, b := 0, 1
	for i := 2; i <= n; i++ {
		a, b = b, a+b // avança mantendo somente dois últimos valores
	}
	return b
}

// Sequence gera um slice contendo os valores de F(0) até F(n).
// Útil para exibir a série completa. Complexidade: tempo O(n), espaço O(n).
func Sequence(n int) []int {
	if n < 0 {
		panic("n não pode ser negativo")
	}
	seq := make([]int, n+1)
	if n >= 0 {
		seq[0] = 0
	}
	if n >= 1 {
		seq[1] = 1
	}
	for i := 2; i <= n; i++ {
		seq[i] = seq[i-1] + seq[i-2]
	}
	return seq
}

// PrintSequence imprime a sequência de F(0) até F(n) formatada.
// Esta função é opcional e demonstra reutilização das funções acima.
func PrintSequence(n int) {
	seq := Sequence(n)
	fmt.Printf("Sequência de Fibonacci até F(%d): %v\n", n, seq)
}
