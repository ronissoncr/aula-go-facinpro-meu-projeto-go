// ===========================================
// SCRIPT DE INICIALIZAÇÃO - PRODUÇÃO
// ===========================================

print('🚀 Inicializando MongoDB para PRODUÇÃO...');

// Criar usuário de produção
db = db.getSiblingDB('app_production');

db.createUser({
  user: 'prod_user',
  pwd: 'prod_super_secure_password789',
  roles: [
    {
      role: 'readWrite',
      db: 'app_production'
    }
  ]
});

print('✅ Usuário de produção criado!');

// Criar coleções com configurações de produção
db.createCollection('users', {
  validator: {
    $jsonSchema: {
      bsonType: 'object',
      required: ['name', 'email', 'age'],
      properties: {
        name: {
          bsonType: 'string',
          description: 'Nome é obrigatório e deve ser string'
        },
        email: {
          bsonType: 'string',
          pattern: '^.+@.+$',
          description: 'Email deve ter formato válido'
        },
        age: {
          bsonType: 'int',
          minimum: 0,
          maximum: 150,
          description: 'Idade deve ser um número entre 0 e 150'
        }
      }
    }
  }
});

db.createCollection('logs');
db.createCollection('audit');
db.createCollection('metrics');

print('📊 Coleções criadas com validação!');

// Criar índices otimizados para produção
db.users.createIndex({ email: 1 }, { unique: true });
db.users.createIndex({ created_at: -1 });
db.users.createIndex({ name: 'text', email: 'text' });

db.logs.createIndex({ timestamp: -1 });
db.audit.createIndex({ action: 1, timestamp: -1 });

print('🔍 Índices otimizados criados!');

// NÃO inserir dados de teste em produção
print('⚠️  Nenhum dado de teste inserido em produção');
print('🎯 Inicialização do MongoDB para PRODUÇÃO concluída!');

// Log de configuração
db.logs.insertOne({
  environment: 'production',
  action: 'database_initialized',
  timestamp: new Date(),
  message: 'MongoDB inicializado para ambiente de produção',
  security_level: 'high'
});