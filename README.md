# 📘 Aula – Meu Projeto em Go

Este repositório é o **projeto base da disciplina**.  
Aqui vamos aprender a criar, rodar, buildar e publicar um projeto em Go, já usando uma **estrutura de pastas padrão de mercado**.

---

## 📂 Estrutura do Projeto
```
meu-projeto-go/
├── cmd/app/          -> ponto de entrada da aplicação (main.go)
├── internal/hello/   -> código interno, não exportável
├── go.mod            -> arquivo de módulo Go
└── README.md         -> instruções do projeto
```

- **cmd/app/** → onde fica o `main.go`, ponto de entrada do programa.  
- **internal/** → pacotes internos, só podem ser usados dentro do projeto.  
- **go.mod** → define que esse diretório é um módulo Go.  
- **README.md** → instruções passo a passo da aula.  

---

## 🚀 Passo a Passo da Aula

### 1. Clonar ou baixar este repositório
```bash
git clone URL_PROJETO
cd meu-projeto-go
```

Se estiver usando o ZIP entregue, basta descompactar e entrar na pasta.

### 2. Rodar o projeto
```bash
go run ./cmd/app
```

➡️ Saída esperada:
```
🚀 Meu primeiro projeto em Go com estrutura de mercado!
Hello Word in FacINpro! 👋
```

### 3. Gerar um executável (build)
```bash
go build -o meuapp ./cmd/app
./meuapp
```

### 4. Publicar no GitHub
```bash
git init
git add .
git commit -m "primeiro commit: projeto base em Go"
git branch -M main
git remote add origin https://github.com/<seu-usuario>/meu-projeto-go.git
git push -u origin main
```

### 5. Entrega
👉 Enviar o **link do repositório no GitHub** como resposta à atividade.

---

## 🎯 Desafio 
- Alterar a função `SayHello()` no arquivo `internal/hello/hello.go` para mostrar uma mensagem personalizada.
- Rodar novamente e ver a saída personalizada.
- Subir no GitHub com um novo commit.

- Parte 2 - Fibonacci:  
  - Criar uma nova função `Fibonacci(n int) int` no arquivo `internal/fibonacci/fibonacci.go` que retorna o n-ésimo número da sequência de Fibonacci.
  - Chamar essa função no `main.go` e imprimir o resultado.
  - Rodar, testar e subir no GitHub.
