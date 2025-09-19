package anamnese

import "fmt"

// CalcularIMC realiza o cálculo do Índice de Massa Corporal.
// A função é exportável porque começa com letra maiúscula.
func CalcularIMC(peso float64, altura float64) float64 {
	// A fórmula é: Peso / (Altura * Altura)
	// altua somente em metros
	resultado := peso / (altura * altura)
	return resultado
}

// ExecutarAnamnese é uma função q  imprime um relatório completo.
func ExecutarAnamnese(nome string, idade int, peso float64, altura float64) {
	fmt.Println("--- Ficha de Anamnese ---")
	fmt.Printf("Nome: %s\n", nome)
	fmt.Printf("Idade: %d anos\n", idade)
	fmt.Printf("Peso: %.2f kg\n", peso)
	fmt.Printf("Altura: %.2f m\n", altura)

	// Chama a outra função para obter o resultado
	imc := CalcularIMC(peso, altura)

	// Imprime o resultado formatado com 2 casas decimais
	fmt.Printf("Resultado (IMC): %.2f\n", imc)
	fmt.Println("-------------------------")
}
