package main

import (
	"bufio"   // Importando o pacote bufio para leitura de entrada do usuário
	"fmt"     // Importando o pacote fmt para formatação de entrada e saída
	"os"      // Importando o pacote os para interações com o sistema operacional
	"strings" // Importando o pacote strings para manipulação de strings
)

// PASSO 1: Definindo uma STRUCT (estrutura personalizada)
// Uma struct é como uma "classe" que agrupa dados relacionados
type Pessoa struct {
	Nome  string
	Idade int
	Email string
}

// PASSO 2: Definindo outra struct para demonstrar mais conceitos
// Aqui temos uma struct para representar um produto, structs são muito úteis para organizar dados complexos
type Produto struct {
	ID    int
	Nome  string
	Preco float64
}

func main() {
	fmt.Println("=== AULA: ESTRUTURAS DE CONTROLE, ARRAYS, MAPS E STRUCTS ===")

	// PASSO 3: Vamos começar com um IF/ELSE simples
	fmt.Println("PASSO 3: Exemplo de IF/ELSE")
	idade := 20
	if idade >= 18 {
		fmt.Printf("Pessoa com %d anos é maior de idade\n", idade)
	} else {
		fmt.Printf("Pessoa com %d anos é menor de idade\n", idade)
	}
	fmt.Println()

	// PASSO 4: Demonstrando ARRAYS
	fmt.Println("PASSO 4: Trabalhando com ARRAYS")

	// Array de tamanho fixo (5 elementos)
	var numeros [5]int
	numeros[0] = 10
	numeros[1] = 20
	numeros[2] = 30
	numeros[3] = 40
	numeros[4] = 50

	fmt.Println("Array criado:", numeros)

	// Array inicializado diretamente
	cores := [3]string{"vermelho", "verde", "azul"}
	fmt.Println("Array de cores:", cores)
	fmt.Println()

	// PASSO 5: Demonstrando MAPS (como dicionários)
	fmt.Println("PASSO 5: Trabalhando com MAPS")

	// Map: chave string, valor int
	idades := make(map[string]int)
	idades["João"] = 25
	idades["Maria"] = 30
	idades["Pedro"] = 22

	fmt.Println("Map de idades:", idades)
	fmt.Println("Idade do João:", idades["João"])

	// Map inicializado diretamente
	capitais := map[string]string{
		"Brasil":    "Brasília",
		"Argentina": "Buenos Aires",
		"Chile":     "Santiago",
	}
	fmt.Println("Map de capitais:", capitais)
	fmt.Println()

	// PASSO 6: Trabalhando com STRUCTS
	fmt.Println("PASSO 6: Trabalhando com STRUCTS")

	// Criando uma pessoa
	pessoa1 := Pessoa{
		Nome:  "Ana",
		Idade: 28,
		Email: "ana@email.com",
	}
	fmt.Printf("Pessoa criada: %+v\n", pessoa1)
	fmt.Printf("Nome: %s, Idade: %d\n", pessoa1.Nome, pessoa1.Idade)

	// Criando um slice de produtos (array dinâmico)
	produtos := []Produto{
		{ID: 1, Nome: "Notebook", Preco: 2500.00},
		{ID: 2, Nome: "Mouse", Preco: 50.00},
		{ID: 3, Nome: "Teclado", Preco: 150.00},
	}
	fmt.Println("Lista de produtos:", produtos)
	fmt.Println()

	// PASSO 7: SWITCH CASE para escolher tipo de laço
	fmt.Println("PASSO 7: SWITCH CASE para escolher tipo de laço")

	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Escolha o tipo de laço de repetição:")
	fmt.Println("1 - For básico (contador)")
	fmt.Println("2 - For como while")
	fmt.Println("3 - For range (arrays)")
	fmt.Println("4 - For range (maps)")
	fmt.Println("5 - For range (structs)")
	fmt.Print("Digite sua opção (1-5): ")

	opcao, _ := reader.ReadString('\n')
	opcao = strings.TrimSpace(opcao)

	// Aqui está o SWITCH CASE principal!
	switch opcao {
	case "1":
		fmt.Println("\n=== FOR BÁSICO (CONTADOR) ===")
		exemploForBasico()

	case "2":
		fmt.Println("\n=== FOR COMO WHILE ===")
		exemploForWhile()

	case "3":
		fmt.Println("\n=== FOR RANGE (ARRAYS) ===")
		exemploForRangeArray()

	case "4":
		fmt.Println("\n=== FOR RANGE (MAPS) ===")
		exemploForRangeMap()

	case "5":
		fmt.Println("\n=== FOR RANGE (STRUCTS) ===")
		exemploForRangeStruct()

	default:
		fmt.Println("Opção inválida! Executando exemplo básico...")
		exemploForBasico()
	}

	fmt.Println("\n=== FIM DA DEMONSTRAÇÃO ===")
}

// PASSO 8: Função para demonstrar for básico
func exemploForBasico() {
	fmt.Println("Contando de 1 a 5:")
	for i := 1; i <= 5; i++ {
		fmt.Printf("Contador: %d\n", i)
	}
}

// PASSO 9: Função para demonstrar for como while
func exemploForWhile() {
	fmt.Println("Contando até chegar a 10:")
	numero := 1
	for numero <= 10 {
		if numero%2 == 0 {
			fmt.Printf("Número par: %d\n", numero)
		}
		numero++
	}
}

// PASSO 10: Função para demonstrar for range com arrays
func exemploForRangeArray() {
	linguagens := []string{"Go", "Python", "JavaScript", "Java", "C++"}

	fmt.Println("Iterando sobre array de linguagens:")
	for index, linguagem := range linguagens {
		fmt.Printf("Posição %d: %s\n", index, linguagem)
	}

	fmt.Println("\nApenas os valores (sem índice):")
	for _, linguagem := range linguagens {
		fmt.Printf("- %s\n", linguagem)
	}
}

// PASSO 11: Função para demonstrar for range com maps
func exemploForRangeMap() {
	notas := map[string]float64{
		"João":  8.5,
		"Maria": 9.2,
		"Pedro": 7.8,
		"Ana":   9.5,
	}

	fmt.Println("Notas dos alunos:")
	for nome, nota := range notas {
		status := "Reprovado"
		if nota >= 7.0 {
			status = "Aprovado"
		}
		fmt.Printf("%s: %.1f - %s\n", nome, nota, status)
	}
}

// PASSO 12: Função para demonstrar for range com slice de structs
func exemploForRangeStruct() {
	funcionarios := []Pessoa{
		{Nome: "Carlos", Idade: 35, Email: "carlos@empresa.com"},
		{Nome: "Lucia", Idade: 28, Email: "lucia@empresa.com"},
		{Nome: "Roberto", Idade: 42, Email: "roberto@empresa.com"},
	}

	fmt.Println("Lista de funcionários:")
	for i, funcionario := range funcionarios {
		categoria := "Júnior"
		if funcionario.Idade >= 30 {
			categoria = "Sênior"
		}
		fmt.Printf("%d. %s (%d anos) - %s - Categoria: %s\n",
			i+1, funcionario.Nome, funcionario.Idade, funcionario.Email, categoria)
	}
}
