# 🧠 Manual Didático – Jogo da Memória em Go (CLI)

Este manual mostra passo a passo como adicionar e executar um **jogo da memória (match de pares)** dentro deste projeto Go.

---
## 🎯 Objetivo
Criar um tabuleiro (ex: 4x4) onde as cartas começam ocultas. A cada jogada o jogador escolhe 2 posições. Se os símbolos forem iguais, o par fica descoberto. Caso contrário, são ocultados novamente. O jogo termina quando todos os pares forem encontrados.

---
## 🗂️ Arquivos Criados
```
internal/memorygame/game.go      # Lógica do jogo (estado, embaralhar, virar cartas)
cmd/memorygame/main.go           # Interface CLI para o usuário jogar
MANUAL_JOGO_MEMORIA.md           # Este manual
```

---
## 🧱 Estruturas Principais
1. struct Card
   - Value: símbolo (rune) da carta (A, B, C ...)
   - Revealed: se está temporariamente virada
   - Matched: se já foi encontrada
2. struct Game
   - Board: matriz de *Card
   - Moves: número de tentativas
   - PairsFound / TotalPairs
   - Métodos: NewGame, FlipPair, HideNonMatched, Render, GameOver

---
## 🔄 Fluxo do Jogo (Resumo)
1. Criar jogo com `NewGame(4,4)` (gera 16 cartas = 8 pares).
2. Embaralhar os pares.
3. Exibir tabuleiro com `?` para cartas não reveladas.
4. Jogador digita: `r1 c1 r2 c2` (ex: `0 0 1 0`).
5. Verifica se formou par:
   - Sim → marca `Matched` e mantém reveladas.
   - Não → mostra por curto período e oculta novamente.
6. Quando `PairsFound == TotalPairs` → vitória.

---
## 📦 Código Principal (Lógica)
Arquivo: `internal/memorygame/game.go` (já incluído no repositório). Nele estão:
- Geração e embaralhamento de cartas
- Validações de jogada
- Impressão do tabuleiro
- Controle de estado

---
## 🖥️ CLI do Jogo
Arquivo: `cmd/memorygame/main.go`
- Lê dados do usuário via stdin
- Mostra o tabuleiro a cada jogada
- Aceita comando `sair` para encerrar

---
## ▶️ Executar o Jogo
Dentro da raiz do projeto:
```bash
go run ./cmd/memorygame
```
Exemplo de interação:
```
🧠 Jogo da Memória - CLI em Go
Encontre todos os pares! Formato de entrada: r1 c1 r2 c2 (ex: 0 0 1 0)
Digite 'sair' para encerrar antecipadamente.

    0  1  2  3
    ────────────
 0 │ ?  ?  ?  ?
 1 │ ?  ?  ?  ?
 2 │ ?  ?  ?  ?
 3 │ ?  ?  ?  ?

Sua jogada (r1 c1 r2 c2): 0 0 1 0
...
```

---
## 🧪 Possíveis Extensões
- Ler tamanho do tabuleiro via flags (`flag` package)
- Persistir placar (arquivo JSON)
- Adicionar limite de tempo
- Adicionar modo "mostrar tudo" com argumento `--debug`

---
## 🛠️ Debug Rápido
Para mostrar sempre os valores (alterar chamada de `Render(true)` em vez de `Render(false)` durante desenvolvimento).

---
## ❗ Tratamento de Erros
- Coordenadas inválidas
- Carta já revelada ou já encontrada
- Seleção da mesma carta duas vezes

Todos retornam mensagens guiando o jogador.

---
## 📚 Conceitos de Go Vistos
- Structs e ponteiros (`*Card`)
- Slices bidimensionais
- Funções com múltiplos retornos `(bool, error)`
- Uso de `rand.Shuffle`
- Controle de tempo (`time.Duration`)
- Entrada padrão (`bufio.Reader`)

---
## 🏁 Finalização
Quando terminar: registrar tentativas e tempo total. Fácil de adaptar para ranking / scoreboard.

Bom estudo e divirta-se expandindo o jogo! 🎮
