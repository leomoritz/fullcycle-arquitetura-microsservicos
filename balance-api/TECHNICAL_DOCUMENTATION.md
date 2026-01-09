# Balance API - Documentação Técnica

> Microsserviço NestJS responsável por consumir eventos de transações e fornecer consultas de saldo de contas.

## 📋 Índice

1. [Visão Geral](#visão-geral)
2. [Arquitetura](#arquitetura)
3. [Tecnologias](#tecnologias)
4. [Pré-requisitos](#pré-requisitos)
5. [Instalação e Execução](#instalação-e-execução)
6. [Configuração](#configuração)
7. [Endpoints da API](#endpoints-da-api)
8. [Fluxo de Consumo de Eventos](#fluxo-de-consumo-de-eventos)
9. [Estrutura do Projeto](#estrutura-do-projeto)
10. [Padrões de Design](#padrões-de-design)
11. [Testes](#testes)
12. [Troubleshooting](#troubleshooting)

---

## 🎯 Visão Geral

O **Balance API** é um microsserviço desenvolvido em **NestJS** que faz parte de uma arquitetura baseada em microsserviços com Event-Driven Architecture (EDA).

### Responsabilidades Principais

- **Consumir Eventos Kafka**: Escuta o tópico `balances` para receber eventos de transações do Wallet Core
- **Atualizar Saldos**: Processa eventos e atualiza o saldo das contas no banco de dados local
- **Fornecer Consultas**: Disponibiliza endpoints REST para consultar o saldo de uma conta

### Características

- ✅ Consumo assíncrono de eventos via Kafka
- ✅ Banco de dados independente (MySQL)
- ✅ Arquitetura Hexagonal (Ports & Adapters)
- ✅ Separação clara de responsabilidades
- ✅ TypeScript para type safety
- ✅ Containerização com Docker

---

## 🏗️ Arquitetura

### Arquitetura Hexagonal (Ports & Adapters)

O Balance API segue o padrão de **Arquitetura Hexagonal**, também conhecida como "Ports & Adapters":

```
┌─────────────────────────────────────────────────────────┐
│                  External Systems                        │
│  (HTTP Clients, Kafka, MySQL Database)                   │
└──────────────────┬──────────────────────────────────────┘
                   │
        ┌──────────┴──────────┐
        │   Presentation      │
        │  (Controllers &     │
        │   Consumers)        │
        └──────────┬──────────┘
                   │
        ┌──────────┴──────────────────────┐
        │   Application Layer (Use Cases) │
        │  (GetBalance, UpdateBalance)    │
        └──────────┬─────────────────────┘
                   │
        ┌──────────┴──────────────────────┐
        │    Domain Layer (Business Logic)│
        │  (Account Entity, Value Objects)│
        └──────────┬─────────────────────┘
                   │
        ┌──────────┴──────────────────────┐
        │  Infrastructure (Gateways)      │
        │  (Repository, Kafka Producer)   │
        └──────────┴──────────────────────┘
                   │
        ┌──────────┴──────────────────────┐
        │   External Implementations      │
        │   (MySQL, Kafka, TypeORM)       │
        └─────────────────────────────────┘
```

### Fluxo de Dados

#### 1. Consumo de Eventos (Entrada)

```
Kafka Topic "balances"
        │
        ▼
WalletEventConsumer (Controller)
        │
        ▼
UpdateBalanceUseCase (Application)
        │
        ▼
Account Domain Entity (Validation)
        │
        ▼
AccountRepository (Persist to DB)
```

#### 2. Consulta de Saldos (Saída)

```
HTTP GET /balances/:accountId
        │
        ▼
BalanceController
        │
        ▼
GetBalanceUseCase (Application)
        │
        ▼
AccountRepository (Fetch from DB)
        │
        ▼
GetBalanceOutputDto (Response)
```

---

## 🛠️ Tecnologias

### Framework e Runtime

| Tecnologia | Versão | Descrição |
|-----------|--------|-----------|
| Node.js | 22 | Runtime JavaScript/TypeScript |
| NestJS | 11.0.1 | Framework web progressivo |
| TypeScript | 5.7.3 | Linguagem tipada |

### Banco de Dados

| Tecnologia | Versão | Descrição |
|-----------|--------|-----------|
| MySQL | 5.7 | Banco de dados relacional |
| TypeORM | 0.3.28 | ORM TypeScript |
| mysql2 | 3.16.0 | Driver MySQL |

### Mensageria

| Tecnologia | Versão | Descrição |
|-----------|--------|-----------|
| Kafka | 6.1.0 | Message Broker |
| KafkaJS | 2.2.4 | Cliente Kafka |
| @nestjs/microservices | 11.1.11 | Integração Kafka com NestJS |

### Ferramentas de Desenvolvimento

| Ferramenta | Versão | Propósito |
|-----------|--------|----------|
| Jest | 30.0.0 | Framework de testes |
| ESLint | 9.18.0 | Linter |
| Prettier | 3.4.2 | Formatador de código |

---

## 📋 Pré-requisitos

### Sistema Operacional
- Linux, macOS ou Windows (com WSL2)

### Softwares Obrigatórios

1. **Node.js** (v18+)
   ```bash
   node --version  # v18.0.0 ou superior
   ```

2. **npm** (v9+)
   ```bash
   npm --version   # v9.0.0 ou superior
   ```

3. **Docker** e **Docker Compose** (para executar com containers)
   ```bash
   docker --version
   docker-compose --version
   ```

### Serviços Externos Necessários

- **Kafka** (para consumir eventos)
- **Zookeeper** (coordenador do Kafka)
- **MySQL** (banco de dados local)

> **Nota**: Se utilizar Docker Compose do diretório raiz, todos esses serviços serão iniciados automaticamente.

---

## 🚀 Instalação e Execução

### Opção 1: Executar com Docker Compose (Recomendado)

A forma mais simples é executar desde o diretório raiz do projeto:

```bash
cd ..
docker-compose up -d
```

Isso iniciará automaticamente:
- Balance API na porta `3003`
- MySQL na porta `3308`
- Kafka e Zookeeper
- Wallet Core na porta `8080`

### Opção 2: Executar Localmente

#### 1. Instalar Dependências

```bash
npm install
```

#### 2. Configurar Variáveis de Ambiente

Crie ou edite o arquivo `.env` com as seguintes variáveis:

```bash
# Kafka Configuration
KAFKA_BROKER=localhost:9092
KAFKA_GROUP_ID=wallet

# Database Configuration
DB_HOST=localhost
DB_PORT=3306
DB_USER=root
DB_PASSWORD=root
DB_NAME=balance
DB_TYPE=mysql

# Application Configuration
APP_PORT=3003
```

#### 3. Executar Migrations do Banco de Dados

O TypeORM está configurado com `synchronize: true`, então as tabelas serão criadas automaticamente.

#### 4. Iniciar a Aplicação

**Modo desenvolvimento (com reload automático):**
```bash
npm run start:dev
```

**Modo produção:**
```bash
npm run build
npm run start:prod
```

**Modo debug:**
```bash
npm run start:debug
```

#### 5. Verificar Inicialização

```
[Nest] 12345   - 01/09/2026, 10:30:00 AM     LOG [NestFactory] Starting Nest application...
[Nest] 12345   - 01/09/2026, 10:30:01 AM     LOG [InstanceLoader] TypeOrmModule dependencies initialized +234ms
[Nest] 12345   - 01/09/2026, 10:30:02 AM     LOG [InstanceLoader] ConfigModule dependencies initialized +2ms
Balance API running on port 3003
Kafka consumer active - listening to configured topics
Kafka Config: Broker=localhost:9092, GroupId=wallet
```

---

## ⚙️ Configuração

### Variáveis de Ambiente

#### Kafka

```env
# Broker do Kafka (ou lista de brokers separados por vírgula)
KAFKA_BROKER=kafka:29092

# ID do grupo de consumidores
# Usado para agrupar múltiplas instâncias do Balance API
KAFKA_GROUP_ID=wallet
```

#### Database (MySQL)

```env
# Host do MySQL
DB_HOST=mysql-balance

# Porta do MySQL
DB_PORT=3306

# Usuário do MySQL
DB_USER=root

# Senha do MySQL
DB_PASSWORD=root

# Nome do banco de dados
DB_NAME=balance

# Tipo de banco (sempre mysql para este projeto)
DB_TYPE=mysql
```

#### Aplicação

```env
# Porta onde a aplicação HTTP será escutada
APP_PORT=3003
```

### Arquivo .env Padrão

Um arquivo `.env` pré-configurado está incluído no repositório:

```dotenv
KAFKA_BROKER=kafka:29092
KAFKA_GROUP_ID=wallet
DB_HOST=mysql-balance
DB_PORT=3306
DB_USER=root
DB_PASSWORD=root
DB_NAME=balance
DB_TYPE=mysql
APP_PORT=3003
```

---

## 🌐 Endpoints da API

### 1. Consultar Saldo de uma Conta

**Endpoint:**
```http
GET /balances/:accountId
```

**Parâmetros:**
- `accountId` (path) - UUID da conta

**Resposta de Sucesso (200):**
```json
{
  "accountId": "123e4567-e89b-12d3-a456-426614174000",
  "balance": 1500.50,
  "timestamp": "2025-01-09T10:30:00Z"
}
```

**Resposta de Erro (404):**
```json
{
  "statusCode": 404,
  "message": "Não foi possível recuperar o saldo atual para conta informada",
  "error": "Not Found"
}
```

**Resposta de Erro (400):**
```json
{
  "statusCode": 400,
  "message": "Descrição do erro de validação",
  "error": "Bad Request"
}
```

#### Exemplos de Uso

**cURL:**
```bash
curl -X GET http://localhost:3003/balances/123e4567-e89b-12d3-a456-426614174000
```

**JavaScript/Fetch:**
```javascript
const accountId = '123e4567-e89b-12d3-a456-426614174000';
fetch(`http://localhost:3003/balances/${accountId}`)
  .then(response => response.json())
  .then(data => console.log(data))
  .catch(error => console.error('Error:', error));
```

**HTTP Client (VSCode):**
```http
GET http://localhost:3003/balances/123e4567-e89b-12d3-a456-426614174000
```

---

## 🔄 Fluxo de Consumo de Eventos

### Processo Completo

```
1. Wallet Core gera evento de transação
   └─► BalanceUpdated event é publicado no tópico "balances"

2. Kafka persiste a mensagem no tópico

3. Balance API (Kafka Consumer) recebe a mensagem
   └─► WalletEventConsumer.handleBalanceUpdated() é executado

4. Mensagem é validada
   └─► Estrutura de envelope de mensagem é verificada

5. Dados são extraídos do payload
   └─► account_id_from, balance_account_id_from
   └─► account_id_to, balance_account_id_to

6. Use Case UpdateBalance é executado (2x - para ambas contas)
   └─► Conta "de" é atualizada com novo saldo
   └─► Conta "para" é atualizada com novo saldo

7. Dados são persistidos no MySQL
   └─► Operação UPSERT na tabela de contas

8. Evento de sucesso é logado
   └─► "Successfully updated balances for accounts..."
```

### Estrutura da Mensagem Kafka

O Balance API espera mensagens no seguinte formato:

```typescript
interface KafkaMessageEnvelope {
  name: string;                           // Ex: "BalanceUpdated"
  payload: WalletUpdateBalanceEventData;  // Dados do evento
}

interface WalletUpdateBalanceEventData {
  account_id_from: string;        // UUID da conta de origem
  balance_account_id_from: number; // Novo saldo da conta de origem
  account_id_to: string;          // UUID da conta de destino
  balance_account_id_to: number;   // Novo saldo da conta de destino
}
```

### Exemplo de Mensagem

```json
{
  "name": "BalanceUpdated",
  "payload": {
    "account_id_from": "123e4567-e89b-12d3-a456-426614174000",
    "balance_account_id_from": 500.00,
    "account_id_to": "987fcdeb-51a2-67c4-b789-012345678901",
    "balance_account_id_to": 1500.00
  }
}
```

### Tratamento de Erros

Se uma mensagem falhar no processamento:

1. **Mensagem inválida**: Mensagem é descartada com log de aviso
2. **Erro de processamento**: Exceção é lançada e Kafka reprocessará a mensagem
3. **Erro de persistência**: Erro é logado e mensagem pode ser reprocessada

---

## 📁 Estrutura do Projeto

### Árvore de Diretórios

```
balance-api/
├── src/
│   ├── main.ts                          # Entry point da aplicação
│   ├── app.module.ts                    # Módulo raiz do NestJS
│   │
│   ├── app/                             # Camada de Aplicação
│   │   ├── dto/                         # Data Transfer Objects
│   │   │   ├── input/
│   │   │   │   └── update-balance-input.dto.ts
│   │   │   └── output/
│   │   │       └── get-balance-output.dto.ts
│   │   │
│   │   ├── mappers/                     # Mapeadores de dados
│   │   │   └── account.mapper.ts        # Converte entre DTOs, Entities, DB
│   │   │
│   │   ├── ports/                       # Interfaces (Hexagonal Arch)
│   │   │   ├── in/                      # Portas de entrada (Use Cases)
│   │   │   │   ├── get-balance.usecase.ts
│   │   │   │   └── update-balance.usecase.ts
│   │   │   └── out/                     # Portas de saída (Gateways)
│   │   │       ├── accounts-balance.gateway.ts
│   │   │       └── event.publisher.gateway.ts
│   │   │
│   │   └── use-cases/                   # Implementações de Use Cases
│   │       ├── get-balance/
│   │       │   └── get-balance.service.ts
│   │       └── update-balance/
│   │           └── update-balance.service.ts
│   │
│   ├── domain/                          # Camada de Domínio (Business Logic)
│   │   ├── entities/
│   │   │   └── account.ts               # Entidade Account
│   │   │
│   │   ├── events/
│   │   │   └── balance-updated.event.ts # Evento de domínio
│   │   │
│   │   └── value-objects/               # Value Objects (imutáveis)
│   │       ├── account-id.ts            # ID da conta
│   │       └── balance.ts               # Saldo
│   │
│   ├── infra/                           # Camada de Infraestrutura
│   │   ├── config/                      # Configurações
│   │   │   ├── kafka.config.ts          # Configuração do Kafka
│   │   │   └── mysql.database.config.ts # Configuração do MySQL
│   │   │
│   │   ├── database/                    # Persistência
│   │   │   ├── entities/
│   │   │   │   └── account.entity.ts    # Entidade TypeORM
│   │   │   │
│   │   │   ├── migrations/
│   │   │   │   └── create-account-table.ts
│   │   │   │
│   │   │   └── repositories/
│   │   │       └── account.repository.ts # Implementação da gateway
│   │   │
│   │   └── messaging/                   # Event Publishing
│   │       ├── messages/
│   │       │   └── message.ts
│   │       │
│   │       └── producers/
│   │           ├── balance-updated.event.producer.ts
│   │           ├── event.producer.ts
│   │           └── kafka.producer.ts    # Implementação Kafka
│   │
│   ├── presentation/                    # Camada de Apresentação
│   │   ├── consumers/                   # Kafka Consumers (Controllers)
│   │   │   ├── wallet-event.consumer.ts # Processa eventos do Wallet Core
│   │   │   └── wallet.payload.data.ts   # Estrutura do payload
│   │   │
│   │   └── controllers/
│   │       └── balance.controller.ts    # Endpoint GET /balances
│   │
│   └── shared/                          # Código compartilhado
│       └── exceptions/
│           └── domain.exception.ts      # Exceções customizadas
│
├── test/                                # Testes End-to-End
│   └── jest-e2e.json                   # Configuração Jest E2E
│
├── api/
│   └── client.http                      # Cliente HTTP para testes
│
├── .env                                 # Variáveis de ambiente
├── .prettierrc                          # Configuração Prettier
├── Dockerfile                           # Imagem Docker
├── eslint.config.mjs                    # Configuração ESLint
├── nest-cli.json                        # Configuração NestJS CLI
├── package.json                         # Dependências e scripts
├── tsconfig.json                        # Configuração TypeScript
├── tsconfig.build.json                  # Configuração build
└── README.md                            # Documentação padrão NestJS
```

### Responsabilidades por Camada

#### 1. **Presentation Layer** (`/presentation`)
- Recebe requisições HTTP e eventos Kafka
- Controllers e Consumers
- Valida entrada de dados
- Retorna respostas ao cliente

#### 2. **Application Layer** (`/app`)
- Contém a lógica de aplicação (Use Cases)
- Orquestra operações entre domínio e infraestrutura
- Implementa Ports (Hexagonal Architecture)
- Utiliza DTOs para isolamento

#### 3. **Domain Layer** (`/domain`)
- Contém regras de negócio
- Entidades e Value Objects
- Validações de domínio
- Completamente independente de frameworks

#### 4. **Infrastructure Layer** (`/infra`)
- Implementações de persistência
- Configurações externas
- Integrações (Kafka, MySQL)
- TypeORM Entities

---

## 🎨 Padrões de Design

### 1. Hexagonal Architecture (Ports & Adapters)

O projeto implementa portas e adaptadores para desacoplar a aplicação da infraestrutura:

```typescript
// PORT (Interface)
export interface IAccountsBalanceGateway {
  saveBalance(account: Account): Promise<void>;
  getBalanceByAccountId(accountId: string): Promise<number>;
}

// ADAPTER (Implementação)
@Injectable()
export class AccountRepository implements IAccountsBalanceGateway {
  // ... implementação
}

// Injeção no módulo
{
  provide: 'IAccountsBalanceGateway',
  useClass: AccountRepository,
}
```

### 2. Dependency Injection

NestJS fornece DI nativo através de decoradores:

```typescript
constructor(
  @Inject('IGetBalanceUseCase')
  private readonly getBalanceUseCase: IGetBalanceUseCase,
) {}
```

### 3. Event-Driven Architecture

Consumo de eventos via Kafka:

```typescript
@EventPattern('balances')
async handleBalanceUpdated(@Payload() message: any): Promise<void> {
  // Processa evento de transação
}
```

### 4. Repository Pattern

Abstração de acesso a dados:

```typescript
@Injectable()
export class AccountRepository implements IAccountsBalanceGateway {
  async saveBalance(account: Account): Promise<void> { }
  async getBalanceByAccountId(accountId: string): Promise<number> { }
}
```

### 5. Mapper Pattern

Transformação entre objetos de diferentes camadas:

```typescript
AccountMapper.toUpdateBalanceInputDto(accountId, balance);
AccountMapper.toDBEntity(account);
AccountMapper.toGetBalanceOutputDTO(accountId, balance);
```

### 6. Entity Pattern

Entidades de domínio com lógica de negócio:

```typescript
export class Account {
  constructor(accountId: AccountId, balance: number) { }
  
  validate(): void {
    if (this._balance.getValue() < 0) {
      throw new DomainException("Balance cannot be negative");
    }
  }
}
```

### 7. Value Objects

Objetos imutáveis que representam valores:

```typescript
export class Balance {
  private readonly _value: number;
  
  constructor(value: number) {
    this._value = value;
  }
  
  getValue(): number {
    return this._value;
  }
}
```

---

## 🧪 Testes

### Executar Testes

**Testes unitários:**
```bash
npm test
```

**Testes em modo watch:**
```bash
npm run test:watch
```

**Testes com cobertura:**
```bash
npm run test:cov
```

**Testes End-to-End:**
```bash
npm run test:e2e
```

### Estrutura de Testes

Testes unitários devem estar no mesmo diretório do arquivo testado com sufixo `.spec.ts`:

```
src/
├── app/
│   └── use-cases/
│       └── get-balance/
│           ├── get-balance.service.ts
│           └── get-balance.service.spec.ts  # Teste unitário
```

### Configuração Jest

Configuração do Jest está em [package.json](package.json):

```json
{
  "jest": {
    "moduleFileExtensions": ["js", "json", "ts"],
    "rootDir": "src",
    "testRegex": ".*\\.spec\\.ts$",
    "transform": {
      "^.+\\.(t|j)s$": "ts-jest"
    },
    "coverageDirectory": "../coverage",
    "testEnvironment": "node"
  }
}
```

---

## 🐛 Troubleshooting

### Problema: "Failed to connect to Kafka"

**Causa**: Broker do Kafka não está acessível

**Solução**:
```bash
# Verifique se Kafka está rodando
docker-compose ps | grep kafka

# Verifique a variável KAFKA_BROKER
cat .env | grep KAFKA_BROKER

# Para execução local, use localhost:9092
# Para Docker Compose, use kafka:29092
```

### Problema: "ECONNREFUSED - MySQL connection refused"

**Causa**: Banco de dados não está disponível

**Solução**:
```bash
# Aguarde alguns segundos para o MySQL iniciar
docker-compose logs mysql-balance

# Ou configure manualmente
mysql -h localhost -u root -proot -e "CREATE DATABASE balance;"
```

### Problema: "Port 3003 is already in use"

**Causa**: Outra aplicação já está usando a porta

**Solução**:
```bash
# Encontre o processo na porta 3003
lsof -i :3003

# Mate o processo (se necessário)
kill -9 <PID>

# Ou configure outra porta no .env
APP_PORT=3004
```

### Problema: "Could not find task.json"

**Causa**: Arquivo de configuração VSCode não existe

**Solução**: Não é obrigatório para execução. Use os comandos npm normalmente:

```bash
npm run start:dev
```

### Problema: "Cannot find module '@nestjs/core'"

**Causa**: Dependências não foram instaladas

**Solução**:
```bash
npm install
```

### Problema: "No consumers found in group 'wallet'"

**Causa**: Aplicação não está rodando ou não conectou ao Kafka

**Solução**:
```bash
# Verifique os logs da aplicação
docker-compose logs -f nestjsapp

# Verifique conectividade com Kafka
docker-compose exec kafka kafka-consumer-groups \
  --bootstrap-server kafka:29092 \
  --list
```

### Problema: "Tópico 'balances' não existe ou está vazio"

**Causa**: Wallet Core não está gerando eventos ou tópico não foi criado

**Solução**:
```bash
# Listar tópicos disponíveis
docker-compose exec kafka kafka-topics \
  --bootstrap-server kafka:29092 \
  --list

# Criar tópico manualmente (se necessário)
docker-compose exec kafka kafka-topics \
  --bootstrap-server kafka:29092 \
  --create \
  --topic balances \
  --partitions 1 \
  --replication-factor 1

# Verificar mensagens no tópico
docker-compose exec kafka kafka-console-consumer \
  --bootstrap-server kafka:29092 \
  --topic balances \
  --from-beginning
```

---

## 📊 Monitoramento

### Logs da Aplicação

**Tempo real:**
```bash
docker-compose logs -f nestjsapp
```

**Últimas N linhas:**
```bash
docker-compose logs -n 100 nestjsapp
```

### Kafka Control Center

Acesse em: http://localhost:9021

- Visualize tópicos e partições
- Monitore grupos de consumidores
- Verifique throughput de mensagens

### Health Check

A aplicação expõe informações no console:

```
Balance API running on port 3003
Kafka consumer active - listening to configured topics
Kafka Config: Broker=kafka:29092, GroupId=wallet
```

---

## 📚 Referências

### Documentação Externa
- [NestJS Docs](https://docs.nestjs.com/)
- [TypeORM Docs](https://typeorm.io/)
- [Kafka Docs](https://kafka.apache.org/documentation/)
- [KafkaJS Docs](https://kafka.js.org/)

### Padrões e Arquitetura
- [Hexagonal Architecture](https://en.wikipedia.org/wiki/Hexagonal_architecture_(software))
- [Event-Driven Architecture](https://en.wikipedia.org/wiki/Event-driven_architecture)
- [Domain-Driven Design](https://en.wikipedia.org/wiki/Domain-driven_design)

---

## 🔐 Notas de Segurança

> **Atenção**: Este é um projeto educacional. As seguintes práticas devem ser implementadas em produção:

1. **Credenciais do Banco de Dados**
   - Não armazene senhas em `.env` versionado
   - Use secrets managers (Vault, AWS Secrets Manager, etc.)

2. **Validação de Entrada**
   - Validar UUIDs
   - Sanitizar inputs
   - Implementar rate limiting

3. **Autenticação e Autorização**
   - Adicionar JWT ou OAuth2
   - Validar tokens antes de processar

4. **Logging e Auditoria**
   - Registrar todas as alterações
   - Implementar log centralizado
   - Não logar dados sensíveis

5. **Tratamento de Erros**
   - Não expor stack traces ao cliente
   - Log detalhado apenas internamente
   - Retornar mensagens genéricas

---

## 📝 Changelog

### v0.0.1 (Inicial)
- ✅ Consumo de eventos Kafka
- ✅ Endpoints de consulta de saldo
- ✅ Persistência em MySQL
- ✅ Arquitetura Hexagonal

---

**Desenvolvido como parte do curso de Microsserviços da FullCycle** ❤️
