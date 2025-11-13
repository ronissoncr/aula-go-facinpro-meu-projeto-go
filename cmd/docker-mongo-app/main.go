package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// ===========================================
// ESTRUTURAS DE DADOS
// ===========================================

// Config representa as configurações da aplicação carregadas do .env
type Config struct {
	Environment   string
	AppName       string
	AppPort       string
	AppHost       string
	MongoURI      string
	MongoHost     string
	MongoPort     string
	MongoDatabase string
	Debug         string
	LogLevel      string
	APITimeout    string
	EnableCORS    string
	AllowOrigins  string
}

// User representa um usuário no MongoDB
type User struct {
	ID        primitive.ObjectID `json:"id" bson:"_id,omitempty"`
	Name      string             `json:"name" bson:"name"`
	Email     string             `json:"email" bson:"email"`
	Age       int                `json:"age" bson:"age"`
	CreatedAt time.Time          `json:"created_at" bson:"created_at"`
	UpdatedAt time.Time          `json:"updated_at" bson:"updated_at"`
}

// App representa nossa aplicação com suas dependências
type App struct {
	Config *Config
	DB     *mongo.Database
	Router *mux.Router
}

// ===========================================
// CONFIGURAÇÃO E INICIALIZAÇÃO
// ===========================================

// LoadConfig carrega as configurações das variáveis de ambiente
func LoadConfig() *Config {
	return &Config{
		Environment:   getEnv("ENV", "development"),
		AppName:       getEnv("APP_NAME", "go-mongo-app"),
		AppPort:       getEnv("APP_PORT", "8080"),
		AppHost:       getEnv("APP_HOST", "0.0.0.0"),
		MongoURI:      getEnv("MONGO_URI", "mongodb://localhost:27017/app_development"),
		MongoHost:     getEnv("MONGO_HOST", "localhost"),
		MongoPort:     getEnv("MONGO_PORT", "27017"),
		MongoDatabase: getEnv("MONGO_DATABASE", "app_development"),
		Debug:         getEnv("DEBUG", "true"),
		LogLevel:      getEnv("LOG_LEVEL", "debug"),
		APITimeout:    getEnv("API_TIMEOUT", "30s"),
		EnableCORS:    getEnv("ENABLE_CORS", "true"),
		AllowOrigins:  getEnv("ALLOW_ORIGINS", "*"),
	}
}

// getEnv retorna o valor da variável de ambiente ou o valor padrão
func getEnv(key, defaultValue string) string {
	if value := os.Getenv(key); value != "" {
		return value
	}
	return defaultValue
}

// ConnectMongoDB conecta com o MongoDB usando as configurações
func ConnectMongoDB(config *Config) (*mongo.Database, error) {
	// Context com timeout para conexão
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	// Configurações de conexão
	clientOptions := options.Client().ApplyURI(config.MongoURI)

	// Conectar ao MongoDB
	client, err := mongo.Connect(ctx, clientOptions)
	if err != nil {
		return nil, fmt.Errorf("erro ao conectar com MongoDB: %v", err)
	}

	// Verificar a conexão
	err = client.Ping(ctx, nil)
	if err != nil {
		return nil, fmt.Errorf("erro ao fazer ping no MongoDB: %v", err)
	}

	// Log da conexão bem-sucedida
	log.Printf("✅ Conectado ao MongoDB no ambiente: %s", config.Environment)
	log.Printf("📊 Database: %s", config.MongoDatabase)
	log.Printf("🏠 Host: %s:%s", config.MongoHost, config.MongoPort)

	return client.Database(config.MongoDatabase), nil
}

// ===========================================
// HANDLERS HTTP
// ===========================================

// HealthHandler verifica a saúde da aplicação
func (a *App) HealthHandler(w http.ResponseWriter, r *http.Request) {
	health := map[string]interface{}{
		"status":      "OK",
		"environment": a.Config.Environment,
		"app_name":    a.Config.AppName,
		"timestamp":   time.Now().Format(time.RFC3339),
		"database":    a.Config.MongoDatabase,
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(health)
}

// ConfigHandler mostra as configurações da aplicação (sem senhas)
func (a *App) ConfigHandler(w http.ResponseWriter, r *http.Request) {
	config := map[string]interface{}{
		"environment":    a.Config.Environment,
		"app_name":       a.Config.AppName,
		"app_port":       a.Config.AppPort,
		"mongo_host":     a.Config.MongoHost,
		"mongo_port":     a.Config.MongoPort,
		"mongo_database": a.Config.MongoDatabase,
		"debug":          a.Config.Debug,
		"log_level":      a.Config.LogLevel,
		"api_timeout":    a.Config.APITimeout,
		"enable_cors":    a.Config.EnableCORS,
		"timestamp":      time.Now().Format(time.RFC3339),
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(config)
}

// CreateUserHandler cria um novo usuário
func (a *App) CreateUserHandler(w http.ResponseWriter, r *http.Request) {
	var user User
	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		http.Error(w, "JSON inválido", http.StatusBadRequest)
		return
	}

	// Adicionar timestamps
	user.CreatedAt = time.Now()
	user.UpdatedAt = time.Now()

	// Inserir no MongoDB
	collection := a.DB.Collection("users")
	result, err := collection.InsertOne(context.Background(), user)
	if err != nil {
		log.Printf("Erro ao inserir usuário: %v", err)
		http.Error(w, "Erro interno do servidor", http.StatusInternalServerError)
		return
	}

	user.ID = result.InsertedID.(primitive.ObjectID)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(user)
}

// GetUsersHandler lista todos os usuários
func (a *App) GetUsersHandler(w http.ResponseWriter, r *http.Request) {
	collection := a.DB.Collection("users")
	cursor, err := collection.Find(context.Background(), bson.D{})
	if err != nil {
		log.Printf("Erro ao buscar usuários: %v", err)
		http.Error(w, "Erro interno do servidor", http.StatusInternalServerError)
		return
	}
	defer cursor.Close(context.Background())

	var users []User
	if err = cursor.All(context.Background(), &users); err != nil {
		log.Printf("Erro ao decodificar usuários: %v", err)
		http.Error(w, "Erro interno do servidor", http.StatusInternalServerError)
		return
	}

	response := map[string]interface{}{
		"users":       users,
		"total":       len(users),
		"environment": a.Config.Environment,
		"timestamp":   time.Now().Format(time.RFC3339),
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

// ===========================================
// MIDDLEWARE
// ===========================================

// LoggingMiddleware faz log das requisições
func LoggingMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()
		next.ServeHTTP(w, r)
		log.Printf("🌐 %s %s %v", r.Method, r.URL.Path, time.Since(start))
	})
}

// CORSMiddleware adiciona headers CORS se habilitado
func (a *App) CORSMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if a.Config.EnableCORS == "true" {
			w.Header().Set("Access-Control-Allow-Origin", a.Config.AllowOrigins)
			w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
			w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")
		}

		if r.Method == "OPTIONS" {
			return
		}

		next.ServeHTTP(w, r)
	})
}

// ===========================================
// SETUP DE ROTAS
// ===========================================

func (a *App) SetupRoutes() {
	// Middleware
	a.Router.Use(LoggingMiddleware)
	a.Router.Use(a.CORSMiddleware)

	// Rotas da API
	a.Router.HandleFunc("/health", a.HealthHandler).Methods("GET")
	a.Router.HandleFunc("/config", a.ConfigHandler).Methods("GET")
	a.Router.HandleFunc("/users", a.CreateUserHandler).Methods("POST")
	a.Router.HandleFunc("/users", a.GetUsersHandler).Methods("GET")

	// Rota raiz
	a.Router.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		welcome := map[string]interface{}{
			"message":     fmt.Sprintf("🚀 Bem-vindo ao %s!", a.Config.AppName),
			"environment": a.Config.Environment,
			"endpoints": []string{
				"GET /health - Status da aplicação",
				"GET /config - Configurações (sem senhas)",
				"GET /users - Lista usuários",
				"POST /users - Cria usuário",
			},
			"timestamp": time.Now().Format(time.RFC3339),
		}
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(welcome)
	}).Methods("GET")
}

// ===========================================
// FUNÇÃO PRINCIPAL
// ===========================================

func main() {
	// Carregar configurações
	config := LoadConfig()

	// Log das configurações iniciais
	log.Printf("🚀 Iniciando %s", config.AppName)
	log.Printf("🔧 Ambiente: %s", config.Environment)
	log.Printf("🏠 Servidor: %s:%s", config.AppHost, config.AppPort)
	log.Printf("🐛 Debug: %s", config.Debug)
	log.Printf("📊 Log Level: %s", config.LogLevel)

	// Conectar ao MongoDB
	db, err := ConnectMongoDB(config)
	if err != nil {
		log.Fatalf("❌ Falha ao conectar com MongoDB: %v", err)
	}

	// Criar instância da aplicação
	app := &App{
		Config: config,
		DB:     db,
		Router: mux.NewRouter(),
	}

	// Configurar rotas
	app.SetupRoutes()

	// Iniciar servidor
	addr := fmt.Sprintf("%s:%s", config.AppHost, config.AppPort)
	log.Printf("🌐 Servidor rodando em http://%s", addr)
	log.Printf("📝 Acesse http://%s para ver os endpoints disponíveis", addr)

	log.Fatal(http.ListenAndServe(addr, app.Router))
}
