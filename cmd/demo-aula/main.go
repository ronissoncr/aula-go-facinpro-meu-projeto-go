package demoaula

import "fmt"

func main() {
	// ===== DEMONSTRAÇÃO PASSO A PASSO PARA A AULA =====

	fmt.Println("=== ESTRUTURAS DE CONTROLE, ARRAYS, MAPS E STRUCTS ===")
	fmt.Println()

	// PASSO 1: IF/ELSE - Condicionais
	fmt.Println("📋 PASSO 1: IF/ELSE")
	idade := 17
	fmt.Printf("Verificando idade: %d anos\n", idade)

	if idade >= 18 {
		fmt.Println("✅ Maior de idade - pode dirigir!")
	} else {
		fmt.Println("❌ Menor de idade - não pode dirigir ainda")
	}
	fmt.Println()

	// PASSO 2: ARRAYS - Coleções de tamanho fixo
	fmt.Println("📋 PASSO 2: ARRAYS")

	// Array de tamanho fixo [5]
	var notas [5]float64
	notas[0] = 8.5
	notas[1] = 7.2
	notas[2] = 9.1
	notas[3] = 6.8
	notas[4] = 8.9

	fmt.Printf("Notas do aluno: %v\n", notas)
	fmt.Printf("Primeira nota: %.1f\n", notas[0])
	fmt.Printf("Última nota: %.1f\n", notas[4])
	fmt.Println()

	// PASSO 3: MAPS - Chave-valor (como dicionários)
	fmt.Println("📋 PASSO 3: MAPS")

	// Map de string para int
	pontuacao := make(map[string]int)
	pontuacao["João"] = 85
	pontuacao["Maria"] = 92
	pontuacao["Pedro"] = 78

	fmt.Printf("Pontuações: %v\n", pontuacao)
	fmt.Printf("Pontuação da Maria: %d\n", pontuacao["Maria"])
	fmt.Println()

	// PASSO 4: STRUCTS - Estruturas personalizadas
	fmt.Println("📋 PASSO 4: STRUCTS")

	// Definindo uma struct inline (normalmente seria definida fora da main)
	type Aluno struct {
		Nome  string
		Nota  float64
		Turma string
	}

	// Criando um aluno
	aluno1 := Aluno{
		Nome:  "Carlos",
		Nota:  8.5,
		Turma: "A",
	}

	fmt.Printf("Aluno: %+v\n", aluno1)
	fmt.Printf("Nome: %s, Nota: %.1f, Turma: %s\n",
		aluno1.Nome, aluno1.Nota, aluno1.Turma)
	fmt.Println()

	// PASSO 5: SWITCH CASE - Menu de opções
	fmt.Println("📋 PASSO 5: SWITCH CASE")

	// Vamos simular diferentes opções
	opcoes := []string{"segunda", "terca", "quarta", "quinta", "sexta"}

	for _, dia := range opcoes {
		fmt.Printf("Dia: %s - ", dia)

		switch dia {
		case "segunda":
			fmt.Println("Início da semana! 💪")
		case "terca", "quarta", "quinta":
			fmt.Println("Meio da semana... 😐")
		case "sexta":
			fmt.Println("SEXTOU! 🎉")
		default:
			fmt.Println("Fim de semana! 😴")
		}
	}
	fmt.Println()

	// PASSO 6: TIPOS DE LAÇOS FOR
	fmt.Println("📋 PASSO 6: TIPOS DE LAÇOS FOR")

	// 6.1 For básico (contador)
	fmt.Println("6.1 - For básico (contador):")
	for i := 1; i <= 3; i++ {
		fmt.Printf("  Contagem: %d\n", i)
	}

	// 6.2 For como while
	fmt.Println("6.2 - For como while:")
	numero := 1
	for numero <= 3 {
		fmt.Printf("  Número: %d\n", numero)
		numero++
	}

	// 6.3 For range com array/slice
	fmt.Println("6.3 - For range com array:")
	frutas := []string{"maçã", "banana", "laranja"}
	for i, fruta := range frutas {
		fmt.Printf("  %d: %s\n", i, fruta)
	}

	// 6.4 For range com map
	fmt.Println("6.4 - For range com map:")
	for nome, pontos := range pontuacao {
		status := "Reprovado"
		if pontos >= 80 {
			status = "Aprovado"
		}
		fmt.Printf("  %s: %d pontos - %s\n", nome, pontos, status)
	}

	// 6.5 For range com slice de structs
	fmt.Println("6.5 - For range com structs:")
	turma := []Aluno{
		{Nome: "Ana", Nota: 9.0, Turma: "A"},
		{Nome: "Bruno", Nota: 7.5, Turma: "B"},
		{Nome: "Clara", Nota: 8.8, Turma: "A"},
	}

	for i, aluno := range turma {
		conceito := "C"
		if aluno.Nota >= 9.0 {
			conceito = "A"
		} else if aluno.Nota >= 8.0 {
			conceito = "B"
		}
		fmt.Printf("  %d: %s (Turma %s) - Nota: %.1f - Conceito: %s\n",
			i+1, aluno.Nome, aluno.Turma, aluno.Nota, conceito)
	}

	fmt.Println()
	fmt.Println("🎯 FIM DA DEMONSTRAÇÃO!")
	fmt.Println("Todos os conceitos foram demonstrados: IF/ELSE, ARRAYS, MAPS, STRUCTS, SWITCH e FOR!")
}
