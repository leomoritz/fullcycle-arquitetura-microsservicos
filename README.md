# Arquitetura de Microsserviços - Wallet Core e Balance API

> Projeto desenvolvido a partir das aulas de **Microsserviços** e **Event-Driven Architecture (EDA)** da FullCycle.

## 📋 Visão Geral

Este é um projeto de arquitetura de microsserviços que demonstra a implementação de uma aplicação com dois serviços independentes que se comunicam através de eventos (Event-Driven Architecture):

- **Wallet Core** (Go): Microsserviço responsável pela criação de clientes, contas e transações
- **Balance API** (NestJS): Microsserviço que escuta eventos de transações e mantém atualizado o saldo das contas

### 🔄 Fluxo de Comunicação

```
┌─────────────────┐                         ┌──────────────────┐
│                 │                         │                  │
│  Wallet Core    │─── BalanceUpdated ──►  │  Balance API     │
│  (Go Service)   │      Event (Kafka)      │  (NestJS)        │
│                 │                         │                  │
└─────────────────┘                         └──────────────────┘
       │                                            │
       ▼                                            ▼
   MySQL DB                                    MySQL DB
  (wallet)                                    (balance)
```

## 🛠️ Pré-requisitos

Antes de executar o projeto, certifique-se de ter os seguintes requisitos instalados:

### Sistema Operacional
- Linux, macOS ou Windows (com WSL2)

### Softwares Necessários

1. **Docker** (v20.10+)
   - [Documentação oficial](https://docs.docker.com/get-docker/)
   - Verificar instalação: `docker --version`

2. **Docker Compose** (v1.29+)
   - [Documentação oficial](https://docs.docker.com/compose/install/)
   - Verificar instalação: `docker-compose --version`

3. **Git** (para clonar o repositório)
   - [Documentação oficial](https://git-scm.com/downloads)
   - Verificar instalação: `git --version`

### Requisitos de Memória
- Mínimo: 4GB de RAM disponível
- Recomendado: 8GB de RAM para melhor desempenho

### Portas Necessárias
Certifique-se de que as seguintes portas estão disponíveis:

| Serviço | Porta | Descrição |
|---------|-------|-----------|
| Wallet Core API | 8080 | API HTTP do Wallet Core |
| Balance API | 3003 | API HTTP do Balance API |
| MySQL (wallet-core) | 3307 | Database do Wallet Core |
| MySQL (balance-api) | 3308 | Database do Balance API |
| Kafka Broker | 9092 | Message Broker |
| Zookeeper | 2181 | Coordenador do Kafka |
| Kafka Control Center | 9021 | UI para monitoramento do Kafka |

## 🚀 Guia de Execução

### Passo 1: Clonar o Repositório

```bash
git clone https://github.com/leomoritz/fullcycle-arquitetura-microsservicos.git
cd fullcycle-arquitetura-microsservicos
```

### Passo 2: Executar os Serviços com Docker Compose

O projeto utiliza **Docker Compose** para orquestrar todos os serviços necessários:

```bash
docker-compose up -d
```

Este comando irá:
1. Construir as imagens do Wallet Core e Balance API
2. Iniciar os containers de:
   - **Wallet Core** (aplicação Go)
   - **Balance API** (aplicação NestJS)
   - **MySQL** (duas instâncias para cada microsserviço)
   - **Zookeeper** (coordenador do Kafka)
   - **Kafka** (message broker para comunicação entre serviços)
   - **Kafka Control Center** (interface para monitoramento)

### Passo 3: Verificar Status dos Serviços

Para verificar se todos os containers estão rodando:

```bash
docker-compose ps
```

Você deve ver uma saída similar a:

```
NAME                 COMMAND                  SERVICE         STATUS
balance-api          ...                      nestjsapp       Up
kafka                ...                      kafka           Up
mysql                ...                      mysql           Up
mysql-balance        ...                      mysql-balance   Up
zookeeper            ...                      zookeeper       Up
control-center       ...                      control-center  Up
goapp                ...                      goapp           Up (healthy)
```

### Passo 4: Aguardar Inicialização Completa

Os serviços podem levar alguns minutos para iniciar completamente. Para monitorar o progresso:

```bash
# Ver logs da aplicação Wallet Core
docker-compose logs -f goapp

# Ver logs da aplicação Balance API
docker-compose logs -f nestjsapp

# Ver logs de todos os serviços
docker-compose logs -f
```

Aguarde até que você veja mensagens como:
- Wallet Core: `Server running on port 8080`
- Balance API: `Balance API running on port 3003` e `Kafka consumer active`

## 📡 API Endpoints

### Wallet Core (Porta 8080)

#### Criar Cliente
```http
POST http://localhost:8080/clients
Content-Type: application/json

{
  "name": "João Silva",
  "email": "joao@example.com",
  "cpf": "12345678901"
}
```

#### Criar Conta
```http
POST http://localhost:8080/accounts
Content-Type: application/json

{
  "client_id": "uuid-do-cliente",
  "initial_balance": 1000.00
}
```

#### Criar Transação (dispara evento BalanceUpdated)
```http
POST http://localhost:8080/transactions
Content-Type: application/json

{
  "account_id_from": "uuid-da-conta-origem",
  "account_id_to": "uuid-da-conta-destino",
  "amount": 500.00
}
```

### Balance API (Porta 3003)

#### Consultar Saldo da Conta
```http
GET http://localhost:3003/balances/:account_id
```

**Resposta esperada:**
```json
{
  "account_id": "uuid-da-conta",
  "balance": 500.00,
  "updated_at": "2025-01-09T10:30:00Z"
}
```

## 🏗️ Arquitetura do Projeto

### Estrutura de Diretórios

```
fullcycle-arquitetura-microsservicos/
├── wallet-core/                          # Microsserviço em Go
│   ├── cmd/walletcore/main.go            # Entrada da aplicação
│   ├── internal/
│   │   ├── entity/                       # Entidades de domínio
│   │   ├── usecase/                      # Casos de uso
│   │   ├── gateway/                      # Interfaces de persistência
│   │   ├── database/                     # Implementações de banco
│   │   ├── event/                        # Eventos de domínio
│   │   └── web/                          # Handlers HTTP
│   ├── pkg/
│   │   ├── events/                       # Event dispatcher
│   │   ├── kafka/                        # Kafka producer
│   │   └── uow/                          # Unit of Work pattern
│   ├── go.mod
│   ├── Dockerfile
│   └── README.md
│
├── balance-api/                          # Microsserviço em NestJS
│   ├── src/
│   │   ├── main.ts                       # Configuração Kafka + Bootstrap
│   │   ├── app/
│   │   │   ├── dto/                      # Data Transfer Objects
│   │   │   ├── use-cases/                # Casos de uso
│   │   │   ├── ports/                    # Portas da aplicação
│   │   │   └── mappers/                  # Mapeadores de dados
│   │   ├── domain/
│   │   │   ├── entities/                 # Entidades de domínio
│   │   │   ├── events/                   # Eventos de domínio
│   │   │   └── value-objects/            # Value objects
│   │   ├── infra/
│   │   │   ├── config/                   # Configurações (Kafka, DB)
│   │   │   ├── database/                 # Entidades e repositórios ORM
│   │   │   └── messaging/                # Produtores de eventos
│   │   ├── presentation/
│   │   │   ├── controllers/              # Controllers HTTP
│   │   │   └── consumers/                # Consumidores Kafka
│   │   └── shared/
│   │       └── exceptions/               # Exceções customizadas
│   ├── test/
│   ├── package.json
│   ├── tsconfig.json
│   ├── Dockerfile
│   └── README.md
│
├── docker-compose.yaml                   # Orquestração de containers
└── README.md                             # Este arquivo
```

### Padrões de Design Utilizados

#### Wallet Core (Go)
- **Clean Architecture**: Separação clara entre camadas (entity, usecase, gateway, web)
- **Unit of Work Pattern**: Gerenciamento de transações de banco de dados
- **Event Dispatcher Pattern**: Publicação de eventos de domínio
- **Repository Pattern**: Abstração de acesso a dados

#### Balance API (NestJS)
- **Hexagonal Architecture (Ports & Adapters)**: Aplicação independente de frameworks
- **CQRS** (implícito): Separação entre leitura e escrita de dados
- **Dependency Injection**: Uso nativo do NestJS para injeção de dependências
- **Mapper Pattern**: Transformação entre DTOs e entidades de domínio
- **Event-Driven Consumer**: Consumidor de eventos Kafka

## 🔌 Fluxo de Eventos

### Criação de Transação

```
1. Cliente faz POST em /transactions no Wallet Core
   └─► CreateTransaction UseCase é executado

2. A transação é criada no banco de dados do Wallet Core
   └─► Um evento "BalanceUpdated" é disparado

3. Kafka Producer publica o evento no tópico "balances"
   └─► O evento é persistido no Kafka

4. Balance API (Kafka Consumer) recebe o evento
   └─► WalletEventConsumer processa a mensagem

5. Balance API executa UpdateBalance UseCase
   └─► O saldo é atualizado no banco de dados do Balance API

6. Cliente pode consultar o saldo via GET /balances/:account_id
   └─► A resposta vem do banco de dados do Balance API
```

## 📊 Monitoramento

### Kafka Control Center

Acesse a interface web do Kafka em:

```
http://localhost:9021
```

Aqui você pode:
- Monitorar tópicos Kafka
- Acompanhar mensagens em tempo real
- Verificar grupos de consumidores
- Analisar throughput e latência

### Logs dos Serviços

Para acompanhar os logs em tempo real:

```bash
# Todos os logs
docker-compose logs -f

# Logs específicos
docker-compose logs -f goapp
docker-compose logs -f nestjsapp
docker-compose logs -f kafka
```

## 🐛 Troubleshooting

### Problema: Containers não iniciam

**Solução:**
```bash
# Verifique se as portas estão disponíveis
lsof -i :8080
lsof -i :3003
lsof -i :9092

# Remova volumes e containers antigos
docker-compose down -v

# Reconstrua e reinicie
docker-compose up -d --build
```

### Problema: Kafka não conecta

**Solução:**
```bash
# Aguarde alguns segundos para o Zookeeper e Kafka iniciarem
# Verifique os logs do Kafka
docker-compose logs kafka

# Verifique a conectividade
docker-compose exec kafka kafka-broker-api-versions --bootstrap-server kafka:29092
```

### Problema: Banco de dados não inicializa

**Solução:**
```bash
# Verifique permissões do diretório .docker
ls -la .docker/

# Remova volumes e recrie
docker-compose down -v
docker-compose up -d
```

### Problema: Balance API não consome eventos

**Solução:**
```bash
# Verifique os logs da aplicação NestJS
docker-compose logs -f nestjsapp

# Verifique a variável KAFKA_BROKER no docker-compose.yaml
# Deve ser "kafka:29092" para conexão interna

# Verifique se o tópico "balances" foi criado
docker-compose exec kafka kafka-topics --list --bootstrap-server kafka:29092
```

## 🔄 Parar os Serviços

Para parar todos os containers sem remover volumes:

```bash
docker-compose down
```

Para parar e remover volumes (cuidado - dados serão perdidos):

```bash
docker-compose down -v
```

## 📚 Documentação Adicional

- [Docker Compose Documentation](https://docs.docker.com/compose/)
- [Kafka Documentation](https://kafka.apache.org/documentation/)
- [NestJS Documentation](https://docs.nestjs.com/)
- [Go Fundamentals](https://golang.org/doc/)
- [MySQL Documentation](https://dev.mysql.com/doc/)

## 🎓 Conceitos de Microsserviços

Este projeto demonstra importantes conceitos de arquitetura de microsserviços:

### 1. **Independência de Serviços**
- Cada serviço possui seu próprio banco de dados
- Comunicação via eventos (não chamadas síncronas)
- Deploy independente

### 2. **Event-Driven Architecture (EDA)**
- Serviços se comunicam através de eventos
- Desacoplamento entre produtores e consumidores
- Escalabilidade horizontal

### 3. **Separação de Responsabilidades**
- **Wallet Core**: Responsável por transações e movimentação de valores
- **Balance API**: Responsável por consulta e atualização de saldos

### 4. **Comunicação Assíncrona**
- Kafka como message broker
- Processamento de eventos em background
- Garantia de entrega de mensagens

## 📝 Notas Importantes

1. **Persistência de Dados**: Os dados são armazenados em volumes Docker. Ao executar `docker-compose down -v`, todos os dados serão perdidos.

2. **Primeira Execução**: A primeira execução pode levar alguns minutos enquanto as imagens são construídas e os containers iniciam.

3. **Tópicos Kafka**: O Kafka está configurado para criar tópicos automaticamente quando necessário.

4. **Variáveis de Ambiente**: Você pode customizar portas e configurações editando o arquivo `docker-compose.yaml`.

## 🤝 Contribuindo

Este é um projeto educacional. Contribuições são bem-vindas!

## 📄 Licença

Este projeto é fornecido como material educacional da FullCycle.

---

**Desenvolvido com ❤️ para aprender sobre Microsserviços e Event-Driven Architecture**
