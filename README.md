# 📘 Aula – Meu Projeto em Go

Este repositório é o **projeto base da disciplina**.  
Aqui vamos aprender a criar, rodar, buildar e publicar um projeto em Go, já usando uma **estrutura de pastas padrão de mercado**.

---

## 📂 Estrutura do Projeto
```
meu-projeto-go/
├── cmd/app/          -> ponto de entrada da aplicação (main.go)
├── internal/hello/   -> código interno, não exportável
├── internal/fibonacci/ -> código interno, não exportável
├── go.mod            -> arquivo de módulo Go
└── README.md         -> instruções do projeto
```

- **cmd/app/** → onde fica o `main.go`, ponto de entrada do programa.  
- **internal/** → pacotes internos, só podem ser usados dentro do projeto.  
- **go.mod** → define que esse diretório é um módulo Go.  
- **README.md** → instruções passo a passo da aula.  

---

## 🚀 Passo a Passo da Aula

### ✅ Pré-requisitos
- Instalar Go (versão 1.22 ou superior) → https://go.dev/dl/
- Verificar instalação:
```bash
go version
```
- (Opcional) Configurar GOPATH e adicionar `~/go/bin` ao PATH.

### 🔄 Sincronizar dependências (caso necessário)
```bash
go mod tidy
```

### 1. Clonar ou baixar este repositório
```bash
git clone https://github.com/ronissoncr/aula-go-facinpro-meu-projeto-go 
cd meu-projeto-go
```

Se estiver usando o ZIP entregue, basta descompactar e entrar na pasta.

### 2. Rodar o projeto
```bash
go run ./cmd/app
```

➡️ Saída (exemplo):
```
🚀 Meu primeiro projeto em Go com estrutura de mercado!
Olá, mundo! 🇺🇿! 👋
F(10) = 55
Sequência de Fibonacci até F(10): [0 1 1 2 3 5 8 13 21 34 55]
```

### 3. Gerar um executável (build)
```bash
go build -o meuapp ./cmd/app
./meuapp
```

### 3.1 Build para outro sistema operacional
Exemplo: gerar binário Linux a partir do macOS/Windows (cross-compilation):
```bash
GOOS=linux GOARCH=amd64 go build -o meuapp-linux ./cmd/app
```
Ou gerar para Windows a partir de Linux/macOS:
```bash
GOOS=windows GOARCH=amd64 go build -o meuapp.exe ./cmd/app
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

---

## 📐 Fibonacci – Explicação

Implementado em `internal/fibonacci/fibonacci.go`:

- `Fibonacci(n int) int`: versão iterativa eficiente (evita recursão e stack overflow em n grandes). Mantém somente dois últimos valores.
- `Sequence(n int) []int`: gera um slice com todos os valores de F(0) até F(n).
- `PrintSequence(n int)`: função utilitária para exibir a lista formatada.

Complexidade:
- Cálculo de um termo: O(n) tempo, O(1) espaço.
- Geração da sequência: O(n) tempo, O(n) espaço.

Validação: `panic` se `n < 0` (entrada inválida), simplificando o exemplo didático.


### Benchmarks (exercício extra)
Adicionar (futuramente) funções `Benchmark...` em `*_test.go` e rodar:
```bash
go test -bench=. -benchmem ./internal/fibonacci
```

---

## 🧪 Rodando no Windows

PowerShell:
```powershell
go run ./cmd/app
go build -o meuapp.exe ./cmd/app
./meuapp.exe
go test ./...
```

CMD (Prompt de Comando):
```bat
go run .\cmd\app
go build -o meuapp.exe .\cmd\app
meuapp.exe
go test ./...
```

Cross-compilation (gerar Linux a partir do Windows):
```powershell
$env:GOOS="linux"; $env:GOARCH="amd64"; go build -o meuapp-linux ./cmd/app
```

## 🐧 Rodando no Linux
```bash
go run ./cmd/app
go build -o meuapp ./cmd/app
./meuapp
go test ./...
```

## 💡 Dicas
- Use `go fmt ./...` para formatar.
- Use `go vet ./...` para detectar possíveis problemas.
- Variáveis de ambiente `GOOS`/`GOARCH` permitem builds multiplataforma.
- Verificar módulos não usados: `go mod tidy`.
- Ver cobertura de testes:
  ```bash
  go test -cover ./...
  ```
- Gerar perfil de cobertura:
  ```bash
  go test -coverprofile=coverage.out ./...
  go tool cover -html=coverage.out
  ```

---

## 🔍 Estrutura com Fibonacci
```
meu-projeto-go/
├── cmd/app/main.go
├── internal/hello/hello.go
├── internal/fibonacci/fibonacci.go
├── internal/fibonacci/fibonacci_test.go
└── go.mod
```


## 🤝 Entregas 😀
1. Fork do repositório
2. Criar branch: `git checkout -b feature/nova-funcionalidade`
3. Commit: `git commit -m "feat: descrição curta"`
4. Push: `git push origin feature/nova-funcionalidade`
5. Abrir Pull Request

Commits: siga padrão convencional (`feat:`, `fix:`, `docs:`, `test:`, etc.).
- Exemplo: `git commit -m "feat: adicionar função Fibonacci🫡"`
- Exemplo: `git commit -m "fix: corrigir bug na função SayHello 😥"`
- Exemplo: `git commit -m "docs: atualizar README com instruções"`
- Exemplo: `git commit -m "test: adicionar testes para Fibonacci 🧪"`


<!-- variaves de ambiente em windows -->
## Como definir variáveis de ambiente no Windows
No PowerShell:
```powershell
$env:VARIAVEL="valor"
```
No CMD (Prompt de Comando):
```bat
set VARIAVEL=valor
```

## Como seta variaveis do GO Lang no Windows apos a instalação do ClI do Go
No PowerShell:
```powershell
$env:GO111MODULE="on"
$env:GOPATH="$HOME\go"
$env:PATH="$env:PATH;$env:GOPATH\bin"
```
No CMD (Prompt de Comando):
```bat
set GO111MODULE=on
set GOPATH=%HOMEPATH%\go
set PATH=%PATH%;%GOPATH%\bin
```


