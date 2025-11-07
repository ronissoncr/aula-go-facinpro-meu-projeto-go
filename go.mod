// Arquivo go.mod define o módulo e suas dependências
//module github.com/seu-usuario/meu-projeto-go

//local do módulo
module meu-projeto-go

go 1.24.0

toolchain go1.24.1

// Dependências para o exemplo Docker + MongoDB
require (
	github.com/gorilla/mux v1.8.1
	go.mongodb.org/mongo-driver v1.13.1
)

// Dependências indiretas
require (
	github.com/golang/snappy v0.0.1 // indirect
	github.com/klauspost/compress v1.13.6 // indirect
	github.com/montanaflynn/stats v0.0.0-20171201202039-1bf9dbcd8cbe // indirect
	github.com/xdg-go/pbkdf2 v1.0.0 // indirect
	github.com/xdg-go/scram v1.1.2 // indirect
	github.com/xdg-go/stringprep v1.0.4 // indirect
	github.com/youmark/pkcs8 v0.0.0-20181117223130-1be2e3e5546d // indirect
	golang.org/x/crypto v0.37.0 // indirect
	golang.org/x/sync v0.13.0 // indirect
	golang.org/x/text v0.24.0 // indirect
)

require (
	github.com/PuerkitoBio/goquery v1.10.3 // indirect
	github.com/andybalholm/cascadia v1.3.3 // indirect
	golang.org/x/net v0.39.0 // indirect
	golang.org/x/time v0.14.0 // indirect
)

// Indica que o módulo é compatível com a versão 1.22 do Go
