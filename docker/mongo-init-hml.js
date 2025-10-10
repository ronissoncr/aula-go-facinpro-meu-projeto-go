// ===========================================
// SCRIPT DE INICIALIZAÇÃO - HOMOLOGAÇÃO
// ===========================================

print('🚀 Inicializando MongoDB para HOMOLOGAÇÃO...');

// Criar usuário de homologação
db = db.getSiblingDB('app_homologation');

db.createUser({
  user: 'hml_user',
  pwd: 'hml_password456',
  roles: [
    {
      role: 'readWrite',
      db: 'app_homologation'
    }
  ]
});

print('✅ Usuário de homologação criado!');

// Criar coleções
db.createCollection('users');
db.createCollection('logs');
db.createCollection('audit');

print('📊 Coleções criadas!');

// Inserir dados limitados para homologação
db.users.insertMany([
  {
    name: 'Usuário HML 1',
    email: 'user1.hml@example.com',
    age: 25,
    created_at: new Date(),
    updated_at: new Date()
  },
  {
    name: 'Usuário HML 2',
    email: 'user2.hml@example.com',
    age: 30,
    created_at: new Date(),
    updated_at: new Date()
  }
]);

print('👥 Usuários de homologação inseridos!');

// Configurar índices para performance
db.users.createIndex({ email: 1 }, { unique: true });
db.users.createIndex({ created_at: -1 });

print('🔍 Índices criados!');
print('🎯 Inicialização do MongoDB para HOMOLOGAÇÃO concluída!');

// Log de configuração
db.logs.insertOne({
  environment: 'homologation',
  action: 'database_initialized',
  timestamp: new Date(),
  message: 'MongoDB inicializado para ambiente de homologação'
});