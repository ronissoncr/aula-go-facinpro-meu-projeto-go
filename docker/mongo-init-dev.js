// ===========================================
// SCRIPT DE INICIALIZAÇÃO - DESENVOLVIMENTO
// ===========================================

print('🚀 Inicializando MongoDB para DESENVOLVIMENTO...');

// Criar usuário de desenvolvimento
db = db.getSiblingDB('app_development');

db.createUser({
  user: 'dev_user',
  pwd: 'dev_password123',
  roles: [
    {
      role: 'readWrite',
      db: 'app_development'
    }
  ]
});

print('✅ Usuário de desenvolvimento criado!');

// Criar coleções iniciais
db.createCollection('users');
db.createCollection('logs');

print('📊 Coleções criadas!');

// Inserir dados de teste para desenvolvimento
db.users.insertMany([
  {
    name: 'João Desenvolvedor',
    email: 'joao.dev@example.com',
    age: 25,
    created_at: new Date(),
    updated_at: new Date()
  },
  {
    name: 'Maria Testadora',
    email: 'maria.test@example.com',
    age: 28,
    created_at: new Date(),
    updated_at: new Date()
  },
  {
    name: 'Pedro QA',
    email: 'pedro.qa@example.com',
    age: 30,
    created_at: new Date(),
    updated_at: new Date()
  }
]);

print('👥 Usuários de teste inseridos!');
print('🎯 Inicialização do MongoDB para DESENVOLVIMENTO concluída!');

// Log de configuração
db.logs.insertOne({
  environment: 'development',
  action: 'database_initialized',
  timestamp: new Date(),
  message: 'MongoDB inicializado com dados de desenvolvimento'
});