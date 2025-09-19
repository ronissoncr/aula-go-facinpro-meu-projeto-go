package fibonacci

// Fibonacci retorna o n-ésimo número da sequência de Fibonacci.

func Fibonacci(n int) int {
	// Casos base: Fibonacci de 0 é 0, e de 1 é 1.
	if n <= 1 {
		return n
	}

	// Inicializa as duas primeiras variáveis da sequência.
	var n2, n1 = 0, 1

	// Itera de 2 até n, calculando o próximo número.
	for i := 2; i <= n; i++ {
		n2, n1 = n1, n2+n1
	}

	// n1 conterá o resultado final.
	return n1
}
