package calculator

// CalcularIMC calcula o Índice de Massa Corporal.
// A função foi renomeada para começar com letra maiúscula, tornando-a
// "exportada" (pública) e visível para outros pacotes, como o 'main'.
func CalcularIMC(peso float64, altura float64) float64 {
	// Prevenção de divisão por zero.
	if altura <= 0 {
		return 0
	}
	// A fórmula do IMC é peso / (altura * altura)
	return peso / (altura * altura)
}
