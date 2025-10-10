#!/bin/bash

# ===========================================
# SCRIPT DE INICIALIZAÇÃO - DOCKER + GO + MONGODB
# ===========================================

set -e

# Cores para output
RED='\033[0;31m'
GREEN='\033[0;32m'
YELLOW='\033[1;33m'
BLUE='\033[0;34m'
NC='\033[0m' # No Color

# Função para imprimir com cores
print_info() {
    echo -e "${BLUE}ℹ️  $1${NC}"
}

print_success() {
    echo -e "${GREEN}✅ $1${NC}"
}

print_warning() {
    echo -e "${YELLOW}⚠️  $1${NC}"
}

print_error() {
    echo -e "${RED}❌ $1${NC}"
}

# Função para verificar se Docker está rodando
check_docker() {
    if ! docker info > /dev/null 2>&1; then
        print_error "Docker não está rodando. Por favor, inicie o Docker primeiro."
        exit 1
    fi
    print_success "Docker está rodando"
}

# Função para verificar se Docker Compose está instalado
check_docker_compose() {
    if ! command -v docker-compose &> /dev/null; then
        print_error "Docker Compose não está instalado."
        exit 1
    fi
    print_success "Docker Compose encontrado"
}

# Função de ajuda
show_help() {
    echo "🐳 Script de Gerenciamento - Docker + Go + MongoDB"
    echo ""
    echo "Uso: $0 [COMANDO] [AMBIENTE]"
    echo ""
    echo "COMANDOS:"
    echo "  up       - Subir ambiente"
    echo "  down     - Parar ambiente"
    echo "  logs     - Ver logs"
    echo "  restart  - Reiniciar ambiente"
    echo "  clean    - Limpar volumes e imagens"
    echo "  test     - Testar API"
    echo "  mongo    - Conectar ao MongoDB"
    echo ""
    echo "AMBIENTES:"
    echo "  dev      - Desenvolvimento (porta 8080)"
    echo "  hml      - Homologação (porta 8081)"
    echo "  prod     - Produção (porta 8082)"
    echo "  all      - Todos os ambientes"
    echo ""
    echo "EXEMPLOS:"
    echo "  $0 up dev          # Subir ambiente de desenvolvimento"
    echo "  $0 logs hml        # Ver logs de homologação"
    echo "  $0 test dev        # Testar API de desenvolvimento"
    echo "  $0 mongo prod      # Conectar ao MongoDB de produção"
    echo ""
}

# Função para subir ambiente
up_environment() {
    local env=$1
    print_info "Subindo ambiente: $env"
    
    case $env in
        "dev"|"hml"|"prod")
            docker-compose --profile $env up -d
            print_success "Ambiente $env iniciado"
            show_urls $env
            ;;
        "all")
            docker-compose --profile dev --profile hml --profile prod up -d
            print_success "Todos os ambientes iniciados"
            show_urls "dev"
            show_urls "hml"
            show_urls "prod"
            ;;
        *)
            print_error "Ambiente inválido: $env"
            show_help
            exit 1
            ;;
    esac
}

# Função para parar ambiente
down_environment() {
    local env=$1
    print_info "Parando ambiente: $env"
    
    case $env in
        "dev"|"hml"|"prod")
            docker-compose --profile $env down
            print_success "Ambiente $env parado"
            ;;
        "all")
            docker-compose down
            print_success "Todos os ambientes parados"
            ;;
        *)
            print_error "Ambiente inválido: $env"
            show_help
            exit 1
            ;;
    esac
}

# Função para mostrar logs
show_logs() {
    local env=$1
    print_info "Mostrando logs do ambiente: $env"
    
    case $env in
        "dev")
            docker-compose logs -f app-dev
            ;;
        "hml")
            docker-compose logs -f app-hml
            ;;
        "prod")
            docker-compose logs -f app-prod
            ;;
        *)
            print_error "Ambiente inválido: $env"
            show_help
            exit 1
            ;;
    esac
}

# Função para reiniciar ambiente
restart_environment() {
    local env=$1
    print_info "Reiniciando ambiente: $env"
    
    down_environment $env
    sleep 2
    up_environment $env
}

# Função para mostrar URLs
show_urls() {
    local env=$1
    case $env in
        "dev")
            echo "🌐 Aplicação Dev: http://localhost:8080"
            echo "🗄️  MongoDB Dev: localhost:27017"
            echo "🖥️  Mongo Express: http://localhost:8090 (admin/admin123)"
            ;;
        "hml")
            echo "🌐 Aplicação HML: http://localhost:8081"
            echo "🗄️  MongoDB HML: localhost:27018"
            ;;
        "prod")
            echo "🌐 Aplicação Prod: http://localhost:8082"
            echo "🗄️  MongoDB Prod: localhost:27019"
            ;;
    esac
}

# Função para testar API
test_api() {
    local env=$1
    local port
    
    case $env in
        "dev") port=8080 ;;
        "hml") port=8081 ;;
        "prod") port=8082 ;;
        *)
            print_error "Ambiente inválido: $env"
            exit 1
            ;;
    esac
    
    print_info "Testando API do ambiente $env (porta $port)"
    
    # Aguardar API ficar disponível
    echo "Aguardando API ficar disponível..."
    for i in {1..30}; do
        if curl -s http://localhost:$port/health > /dev/null; then
            break
        fi
        echo -n "."
        sleep 1
    done
    echo ""
    
    # Testar endpoints
    echo ""
    print_info "=== TESTANDO ENDPOINTS ==="
    
    echo "1. Health Check:"
    curl -s http://localhost:$port/health | jq '.'
    
    echo ""
    echo "2. Configurações:"
    curl -s http://localhost:$port/config | jq '.'
    
    echo ""
    echo "3. Listar Usuários:"
    curl -s http://localhost:$port/users | jq '.'
    
    echo ""
    echo "4. Criar Usuário:"
    curl -s -X POST http://localhost:$port/users \
        -H "Content-Type: application/json" \
        -d '{"name":"Teste API","email":"teste@api.com","age":25}' | jq '.'
    
    print_success "Teste da API concluído!"
}

# Função para conectar ao MongoDB
connect_mongo() {
    local env=$1
    local container
    local user
    local password
    local database
    
    case $env in
        "dev")
            container="mongo-dev"
            user="dev_user"
            password="dev_password123"
            database="app_development"
            ;;
        "hml")
            container="mongo-hml"
            user="hml_user"
            password="hml_password456"
            database="app_homologation"
            ;;
        "prod")
            container="mongo-prod"
            user="prod_user"
            password="prod_super_secure_password789"
            database="app_production"
            ;;
        *)
            print_error "Ambiente inválido: $env"
            exit 1
            ;;
    esac
    
    print_info "Conectando ao MongoDB do ambiente $env"
    docker-compose exec $container mongosh -u $user -p $password $database
}

# Função para limpeza
clean_all() {
    print_warning "Esta operação irá remover todos os containers, volumes e imagens relacionadas."
    read -p "Deseja continuar? (y/N): " -n 1 -r
    echo
    if [[ $REPLY =~ ^[Yy]$ ]]; then
        print_info "Parando todos os containers..."
        docker-compose down
        
        print_info "Removendo volumes..."
        docker volume prune -f
        
        print_info "Removendo imagens não utilizadas..."
        docker image prune -f
        
        print_success "Limpeza concluída!"
    else
        print_info "Operação cancelada."
    fi
}

# ===========================================
# FUNÇÃO PRINCIPAL
# ===========================================

main() {
    local command=$1
    local environment=$2
    
    # Verificar Docker
    check_docker
    check_docker_compose
    
    case $command in
        "up")
            if [ -z "$environment" ]; then
                print_error "Ambiente não especificado"
                show_help
                exit 1
            fi
            up_environment $environment
            ;;
        "down")
            if [ -z "$environment" ]; then
                print_error "Ambiente não especificado"
                show_help
                exit 1
            fi
            down_environment $environment
            ;;
        "logs")
            if [ -z "$environment" ]; then
                print_error "Ambiente não especificado"
                show_help
                exit 1
            fi
            show_logs $environment
            ;;
        "restart")
            if [ -z "$environment" ]; then
                print_error "Ambiente não especificado"
                show_help
                exit 1
            fi
            restart_environment $environment
            ;;
        "test")
            if [ -z "$environment" ]; then
                print_error "Ambiente não especificado"
                show_help
                exit 1
            fi
            test_api $environment
            ;;
        "mongo")
            if [ -z "$environment" ]; then
                print_error "Ambiente não especificado"
                show_help
                exit 1
            fi
            connect_mongo $environment
            ;;
        "clean")
            clean_all
            ;;
        "help"|"--help"|"-h")
            show_help
            ;;
        *)
            if [ -z "$command" ]; then
                show_help
            else
                print_error "Comando inválido: $command"
                show_help
                exit 1
            fi
            ;;
    esac
}

# Executar função principal com todos os argumentos
main "$@"