# Arquitetura Baseada em Microsserviços

🧠 Mindmap: https://whimsical.com/arquitetura-baseada-em-microsservicos-LxhyGPUwg3x5ghrc4s6WM4

## Conceitos Básicos

### Pilares Básicos sobre Microsserviços
- Aplicações comuns.
- Objetivos bem definidos.
- Faz parte de um ecossistema.
- São independentes ou autonomos.
- Se comunicam entre si o tempo todo.
- Devem ter bancos de dados independentes.
- Podem ser escritos em linguagens diferentes.

### Principais Diferenças entre Monolito e Microsserviços
- Objetivos / Domínio:
    - Monolito: Objetivo amplo. Toda aplicação. Todos os contextos dentro do mesmo sistema.
    - Microsserviços: Objetivo específico. Um contexto. Uma funcionalidade.
- Linguagem de Programação:
    - Monolito: Geralmente uma única linguagem/tecnologia.
    - Microsserviços: Diversas tecnologias; Pode ser escrita em linguagens diferentes.
- Deploy:
    - Monolito: Deploy único. Toda aplicação.
    - Microsserviços: Deploy independente. Cada serviço.
- Organização da Equipe:
    - Monolito: Equipe única. Todos trabalham no mesmo código.
    - Microsserviços: Equipes pequenas e independentes. Cada equipe pode ser responsável por um ou mais serviços.
- Começar um projeto / POC
    - Monolito: Mais rápido para começar. Menos complexidade inicial.
    - Microsserviços: Pode ser mais complexo no início. Requer planejamento de arquitetura.
- Escalabilidade:
    - Monolito: Escala a aplicação inteira. Pode ser ineficiente.
    - Microsserviços: Escala serviços individualmente. Mais eficiente.
- Manutenção:
    - Monolito: Pode ser mais difícil de manter com o tempo. Código pode se tornar complexo.
    - Microsserviços: Facilita a manutenção. Serviços menores e mais focados.
- Testes:
    - Monolito: Testes podem ser mais complexos. Todo o sistema precisa ser testado.
    - Microsserviços: Testes focados em serviços individuais. Pode ser mais simples.
- Comunicação:
    - Monolito: Comunicação interna. Chamadas de função/métodos.
    - Microsserviços: Comunicação via APIs/HTTP. Pode introduzir latência.
- Banco de Dados:
    - Monolito: Geralmente um único banco de dados.
    - Microsserviços: Bancos de dados independentes para cada serviço.
- Tolerância a Falhas:
    - Monolito: Uma falha pode afetar todo o sistema.
    - Microsserviços: Falhas isoladas. Outros serviços podem continuar funcionando.

### Quando usar microsserviços
- Projetos grandes e complexos.
- Necessidade de escalabilidade.
- Contextos de negócio bem definidos.
- Quando você possui maturidade nos processos de entrega.
- Quando você possui maturidade técnica dos times.
- Quando eu tenho a necessidade de escala de apenas uma parte do meu sistema.

### Quando Monolito é melhor
- POC (prova de conceito)
- Projetos pequenos e simples.
- Novos projetos onde não conhecemos todo o domínio.
- Governança simplificada sobre tecnologias.
- Facilidade na contratação de profissionais.
- Facilidade no treinamento dos devs.
- Tudo em um único lugar.
- Compartilhamento claro de libs (shared kernel).

### Migração de Monolito para Microsserviços
- Identificar os limites do domínio / Separação de contextos (Domain-Driven Design).
- Evitar excesso de granularidade.
- Verifique dependências.
    - Quando um serviço depende muito de outro, talvez eles devam estar juntos (monolito distribuído).
- Planeje o processo de migração dos bancos de dados.
- Utilizar eventos para comunicação entre serviços e para desacoplar dependências.
- Não ter medo de duplicação de dados.
- Consistência eventual.
- Maturidade de processos: CI/CD, Testes, Ambientes.
- Comece pelas beiradas do sistema.
    - Padrão de estrangulamento (Strangler Fig Pattern).

## Características

@Martin Fowler: https://martinfowler.com/articles/microservices.html

### Componentização
Componentização via serviços: Um componente é um pedaço de software que encapsula um conjunto de funcionalidades e expõe uma interface bem definida para interagir com outras partes do sistema. A componentização é o processo de dividir um sistema em componentes menores e independentes, cada um responsável por uma funcionalidade específica que pode ser substituída ou atualizada sem afetar o restante do sistema. A Biblioteca de certa forma é um componente, mas que é considerada "in memory", ou seja, roda dentro do mesmo processo da aplicação. Agora quando estamos falando de serviços, estamos falando de componentes que rodam em processos separados, ou seja, são executados de forma independente (out of process).

### Organização através das áreas de negócio
Cada serviço é responsável por uma capacidade ou funcionalidade específica do negócio, permitindo que equipes multidisciplinares trabalhem de forma independente e ágil.

### Produtos e não projetos
Cada serviço é tratado como um produto independente, com ciclo de vida próprio, desde o desenvolvimento até a manutenção e evolução. Isso permite uma maior autonomia das equipes e uma melhor adaptação às necessidades do mercado.

### Smart endpoints and dumb pipes
Os microsserviços devem ser projetados para serem "smart endpoints", ou seja, eles devem conter a lógica de negócio e serem responsáveis por processar as requisições de forma eficiente. Já os "dumb pipes" referem-se aos mecanismos de comunicação entre os serviços, que devem ser simples e leves, como filas de mensagens ou APIs RESTful. Isso ajuda a manter a simplicidade e a escalabilidade do sistema como um todo. Diferente de um ESB (Enterprise Service Bus) que é um "smart pipe" que contém lógica de negócio e regras de roteamento, o que pode levar a um acoplamento indesejado entre os serviços.

### Governança descentralizada
Cada equipe é responsável por suas próprias decisões técnicas e de arquitetura, permitindo uma maior autonomia e agilidade na entrega de valor ao negócio. Isso inclui a escolha de tecnologias, ferramentas e práticas de desenvolvimento que melhor atendam às necessidades do serviço. No entanto, é importante manter um certo nível de padronização e alinhamento entre as equipes para garantir a interoperabilidade e a coesão do sistema como um todo. Isso pode ser alcançado através de diretrizes e boas práticas compartilhadas, sem impor uma governança centralizada rígida. Sendo assim, deve-se pensar em "consumer driven contract", ou seja, o serviço que consome a API deve definir o contrato (contrato de consumo) e o serviço que provê a API deve seguir esse contrato.

### Dados descentralizados
Cada serviço deve ter seu próprio banco de dados ou esquema de dados, evitando o compartilhamento direto de dados entre os serviços. Isso ajuda a manter a independência e a autonomia dos serviços, permitindo que cada um evolua e escale de forma independente. Além disso, a descentralização dos dados pode melhorar a resiliência do sistema, já que uma falha em um serviço não afeta diretamente os dados de outros serviços. No entanto, é importante considerar estratégias para garantir a consistência e a integridade dos dados em todo o sistema, como o uso de eventos para sincronizar informações entre os serviços quando necessário. Isso pode levar a uma abordagem de consistência eventual, onde os dados podem estar temporariamente inconsistentes, mas eventualmente se tornam consistentes através de mecanismos de sincronização assíncrona. Referente a duplicação de dados, é importante entender que em uma arquitetura de microsserviços, a duplicação de dados pode ser uma prática aceitável e até mesmo necessária para garantir a independência e a autonomia dos serviços. Cada serviço deve ser responsável por seu próprio conjunto de dados, o que pode levar à duplicação de informações entre serviços diferentes. No entanto, é importante gerenciar essa duplicação de forma cuidadosa, utilizando estratégias como eventos para sincronizar dados entre serviços quando necessário, e garantindo que cada serviço mantenha a integridade e a consistência dos seus próprios dados. A duplicação de dados deve ser vista como uma trade-off entre a independência dos serviços e a complexidade adicional na gestão dos dados. Além disso, a duplicação de dados pode melhorar a performance e a escalabilidade do sistema, já que cada serviço pode acessar seus próprios dados localmente, sem depender de chamadas remotas para outros serviços. Portanto, é importante avaliar cuidadosamente as necessidades do sistema e as características dos serviços ao decidir sobre a duplicação de dados em uma arquitetura de microsserviços. Em resumo, a duplicação de dados em uma arquitetura de microsserviços é uma prática aceitável e pode ser benéfica para garantir a independência e a autonomia dos serviços, desde que seja gerenciada de forma cuidadosa e estratégica e sem duplicar informações desnecessariamente.

### Automação de infraestrutura
A automação de infraestrutura é um aspecto crucial em uma arquitetura de microsserviços, pois permite a criação, configuração e gerenciamento de ambientes de forma rápida, eficiente e consistente. Isso inclui a automação de tarefas como provisionamento de servidores, configuração de redes, implantação de aplicações e monitoramento de serviços. Ferramentas como Ansible, Terraform, Kubernetes e Docker são comumente utilizadas para automatizar a infraestrutura em ambientes de microsserviços. A automação ajuda a reduzir erros humanos, aumentar a agilidade na entrega de software e melhorar a escalabilidade e a resiliência do sistema como um todo. Além disso, a automação de infraestrutura facilita a implementação de práticas de DevOps, permitindo que as equipes de desenvolvimento e operações trabalhem de forma mais integrada e colaborativa. Em resumo, a automação de infraestrutura é essencial para o sucesso de uma arquitetura de microsserviços, garantindo que os serviços possam ser implantados, gerenciados e escalados de forma eficiente e confiável.

### Desenhado para falhar
Em uma arquitetura de microsserviços, é fundamental projetar os serviços para serem resilientes e tolerantes a falhas. Isso significa que cada serviço deve ser capaz de lidar com falhas de forma graciosa, sem afetar o funcionamento do sistema como um todo. Algumas estratégias comuns para alcançar essa resiliência incluem:
- **Circuit Breaker**: Implementar padrões de circuit breaker para evitar que falhas em um serviço se propaguem para outros serviços. Quando um serviço detecta que outro serviço está falhando, ele pode interromper temporariamente as chamadas para esse serviço, permitindo que ele se recupere.
- **Retries**: Implementar mecanismos de retry para tentar novamente chamadas que falharam, com backoff exponencial para evitar sobrecarregar o serviço que está falhando.
- **Timeouts**: Definir timeouts para chamadas entre serviços, garantindo que um serviço não fique esperando indefinidamente por uma resposta de outro serviço.
- **Fallbacks**: Implementar estratégias de fallback, onde um serviço pode fornecer uma resposta alternativa ou degradada quando outro serviço está indisponível.
- **Monitoramento e Alertas**: Monitorar continuamente o estado dos serviços e configurar alertas para detectar falhas rapidamente e tomar ações corretivas.
- **Isolamento de Falhas**: Projetar os serviços de forma que uma falha em um serviço não afete diretamente outros serviços, utilizando técnicas como filas de mensagens para desacoplar a comunicação entre serviços.
- **Testes de Resiliência**: Realizar testes de resiliência, como o Chaos Engineering, para identificar pontos fracos na arquitetura e melhorar a capacidade de recuperação dos serviços.
Ao adotar essas estratégias, é possível criar uma arquitetura de microsserviços que seja robusta e capaz de lidar com falhas de forma eficaz, garantindo a continuidade do serviço e uma melhor experiência para os usuários finais.

### Design Evolutivo
A arquitetura de microsserviços deve ser projetada para evoluir ao longo do tempo, permitindo que novos serviços sejam adicionados, serviços existentes sejam modificados ou removidos, e tecnologias sejam atualizadas sem causar interrupções significativas no sistema. Isso envolve a adoção de práticas como:
- **Versionamento de APIs**: Implementar versionamento nas APIs dos serviços para permitir mudanças sem quebrar a compatibilidade com clientes existentes.
- **Desenvolvimento Incremental**: Adotar uma abordagem de desenvolvimento incremental, onde novas funcionalidades são adicionadas gradualmente, permitindo feedback contínuo e ajustes conforme necessário.
- **Refatoração Contínua**: Estar aberto a refatorações e melhorias constantes na arquitetura e no código dos serviços, garantindo que eles permaneçam eficientes e alinhados com as necessidades do negócio.
- **Desacoplamento**: Manter os serviços o mais desacoplados possível, facilitando a substituição ou atualização de serviços individuais sem afetar outros serviços.
- **Automação de Testes**: Implementar testes automatizados para garantir que as mudanças nos serviços não introduzam regressões ou quebras de funcionalidade.
- **Monitoramento Contínuo**: Monitorar o desempenho e a saúde dos serviços para identificar áreas que precisam de melhorias ou ajustes.
- **Cultura de Aprendizado**: Fomentar uma cultura de aprendizado e adaptação, onde as equipes estão sempre buscando maneiras de melhorar a arquitetura e os processos de desenvolvimento.
Ao adotar um design evolutivo, a arquitetura de microsserviços pode se adaptar às mudanças nas necessidades do negócio e nas tecnologias, garantindo sua relevância e eficácia ao longo do tempo.

## Resilência

**Princípio da Resiliência** 
> "Em algum momento todo sistema vai falhar. Falhas são inevitáveis, mas o sistema deve continuar funcionando de forma aceitável mesmo quando partes dele falham."

### O que é resiliência?
- Resiliência é um conjunto de estratégias adotadas **intencionalmente** para a **adaptação** de um sistema quando uma falha ocorre. 
- A principal questão é: Você se dobra ou quebra?
- Ter estratégias de resiliência nos possibilita minimizar os riscos de perda de dados e transaçlões importantes para o negócio.

**Proteger e ser Protegido**
- Um sistema em uma arquitetura distribuída precisa adotar mecanismos de autopreservação para garantir ao máximo sua operação com **qualidade**.
- Um sistema precisa não pode ser "egoísta" ao ponto de realizar mais requisições em um sistema que está falhando.
- Um sistema lento no ar muitas vezes é pior do que um sistema fora do ar (efeito dominó).

### Estratégias de Resiliência

**Health Check**
- Sem sinais vitais, não é possível saber a "saúde" de um sistema.
- Um sistema que não está saudável possui uma chance de se recuperar caso o tráfego pare de ser direcionado a ele temporariamente.
- Health check de qualidade.
    - Deve ser simples e rápido.
    - Deve ser específico para o serviço.
    - Deve ser acessível publicamente (sem autenticação).
    - Deve retornar um status HTTP adequado (200 OK, 503 Service Unavailable).
    - Deve incluir verificações de dependências críticas (banco de dados, serviços externos).
    - Deve ser monitorado continuamente.
    - Deve ser utilizado por orquestradores (Kubernetes, AWS ECS, etc.)
- Health check passivo: As informações de integridade dependem de um estimulo externo (chamadas aos pontos de extemidade) para serem persitidas e/ou obtidas. Isso não deve ser um problema se estiver publicando em um ambiente que já possui um orquestrador (Kubernetes, AWS ECS, etc.) que faz chamadas periódicas ao endpoint de health check (se devidamente configurado) para monitorar a integridade de cada pod.
- Health check ativo: A aplicação é configurada para executar sozinha as vaklidações de integridade periodicamente e persistir o resultado em memória ou em um banco de dados rápido (Redis, Memcached, etc.). Isso é útil quando não se tem um orquestrador que faça chamadas periódicas ao endpoint de health check. Em cenários mais complexos, talvez seja mais interessante usar o modo ativo. Neste modo a execução dos HealthChecks são executados em background utilizando parâmetros globais de 'Delay' e 'Period' além do 'Predicate'. Os parâmetros 'Delay' e 'Period' podem ser perosnalizados para cada HealthCheck.  Em alguns cenários é mais interessante usar o modo ativo aliado a um cache de resultados, persistindo os resultados em memória informando os resultados para as requisições que chegam. Claro que a adoção modo ativo aliado a um cache implica em aceitar um grande **trade-off** de latência entre uma solicitação externa e a periodicidade que foi determinada para execução automática. Esta abordagem de modo ativo aliado a um cache,  pode reduzir o custo de execução das validações de integridade, melhorando a responsividade dos pontos de extremidade de validações de integridade e evitar abusos de requisições (com o custo de execução das validações de integridade) comum em ambientes complexos com muitos micros serviço que fazem chamadas para os pontos de extremidade de validações de integridade na sua aplicação.

**Rate Limiting**
- Protege o sistema baseado no que ele foi projetado para suportar.
- Preferência programada por tipo de client.
- Pode ser implementado em diferentes níveis:
    - Nível de aplicação (dentro do código do serviço).
    - Nível de API Gateway (antes de chegar ao serviço).
    - Nível de infraestrutura (firewall, load balancer).
- Estratégias comuns:
    - Token Bucket
    - Leaky Bucket
    - Fixed Window
    - Sliding Log
    - Sliding Window
- Definir políticas claras de rate limiting.
    - Limites por usuário, IP, API key, etc.
    - Diferentes limites para diferentes tipos de usuários (ex: free vs premium).
    - Limites por endpoint (ex: endpoints críticos podem ter limites mais baixos).

**Circuit Breaker**
- Protege o sistema fazendo com que as requisições feitas para ele sejam negadas. Ex: 500.
- Circuito fechado = Requisições chegam normalmente.
- Circuito aberto = Requisições são negadas imediatamente. Requisições não chegam ao sistema. Erro instantâneo ao client.
- Circuito meio aberto = Requisições são enviadas de forma controlada para testar se o sistema já está saudável. Em outras palavras, ele permite uma quantidade limitada de requisições para verificação se o sistema tem condições de voltar ao ar integralmente.

**API Gateway**
- Garante que requisições "inapropriadas" cheguem até o sistema: Ex: Usuário não autenticado.
- Implementa políticas de Rate Limiting, Health Check, Logging, Monitoramento, etc.
- Nos ajuda a organizar nossos microsserviços em contextos.
- Uma curiosidade importante a se destacar é que a API Gateway nos ajuda a evitar o problema da "estrela da morte" (starvation) por exemplo, onde um serviço depende de muitos outros serviços, criando um acoplamento indesejado e dificultando a resiliência do sistema como um todo. Isso é feito através da implementação de padrões como o "Backend for Frontend" (BFF), onde a API Gateway atua como um intermediário entre o cliente e os microsserviços, agregando e simplificando as chamadas para os serviços backend. Dessa forma, o cliente não precisa fazer múltiplas chamadas para diferentes serviços, reduzindo a complexidade e o acoplamento entre os serviços. Além disso, a API Gateway pode implementar políticas de resiliência, como circuit breakers e rate limiting, para proteger os serviços backend de falhas e sobrecarga. Em resumo, a API Gateway é uma peça fundamental na arquitetura de microsserviços, ajudando a organizar os serviços em contextos, melhorar a resiliência do sistema e evitar problemas como a "estrela da morte".

**Service Mesh**
- Controla o tráfego de rede entre os microsserviços por meio de proxies.
- Proxies podem ser sidecars, que são implantados junto com os microsserviços, ou podem ser proxies dedicados que gerenciam o tráfego entre os serviços.
- Permite a implementação de políticas de segurança, como autenticação e autorização, de forma centralizada.
- Facilita a observabilidade, permitindo monitorar e registrar o tráfego entre os microsserviços.
- Exemplos de Service Mesh incluem Istio, Linkerd e Consul.
- Evita implementações de proteção pelo próprio sistema.
- mTLS (Mutual TLS): Comunicação segura entre serviços, garantindo que ambos os lados da comunicação sejam autenticados.
- Circuit breaker, retry, timeout, fault injection, etc.

**Assíncronia**
- Evita perda de dados e transações importantes para o negócio.
- Evita o efeito dominó.
- Não há perda de dados no envio de uma transação de o server estiver fora.
- Servidor pode processar a transação em seu tempo quando estiver online.
- Entender com profundidade o message broker / sistema de stream.
- Exemplos: RabbitMQ, Apache Kafka, AWS SQS, Google Pub/Sub, etc.

**Garantias de entrega: Retry**
- Linear - Sem backoff.
![alt text](image.png)
- Exponencial - Com backoff: Faz com que o tempo entre as tentativas aumente exponencialmente para ajudar o servidor a se recuperar.
![alt text](image-1.png)
- Exponencial com jitter - Com backoff e variação aleatória: Adiciona uma variação aleatória ao tempo de espera para evitar que múltiplas tentativas sejam feitas ao mesmo tempo.
![alt text](image-2.png)

**Garantias de entrega: Kafka**
![alt text](image-3.png)
- Ack 0: O produtor não espera por nenhuma confirmação do broker. A mensagem é considerada enviada assim que é escrita na rede. Isso oferece a menor latência, mas não garante que a mensagem foi realmente recebida pelo broker. Geralmente conhecida como "fire and forget". Exemplos: 
    - Logs de eventos onde a perda ocasional de mensagens é aceitável.
    - Dados de telemetria onde a latência é crítica e a perda de algumas mensagens não afeta significativamente a análise.
- Ack 1: O produtor espera por uma confirmação do líder da partição antes de considerar a mensagem como enviada. Isso garante que a mensagem foi recebida pelo broker, mas pode aumentar a latência. Exemplos:
    - Sistemas de monitoramento onde é importante garantir que os dados foram recebidos, mas a latência não é tão crítica.
    - Aplicações de análise de dados onde a integridade dos dados é importante, mas a latência pode ser tolerada.
- Ack all (ou -1): O produtor espera por confirmações de todos os réplicas da partição antes de considerar a mensagem como enviada. Isso oferece a maior garantia de entrega, mas também a maior latência. Exemplos:
    - Sistemas financeiros onde a integridade dos dados é crítica e a perda de mensagens não é aceitável.
    - Aplicações de comércio eletrônico onde a consistência dos dados é essencial para evitar problemas como pedidos duplicados ou perda de informações de pagamento.

### Situações complexas
- O que acontece se o message broker cair? 
- Haverá perda de mensagens?
- Seu sistema ficará fora do ar?
- Como garantir resiliência em situações complexas?
- Estratégias:
    - Ter um cluster de message broker.
    - Ter mais de um message broker (federation).
    - Ter mais de um data center (disaster recovery).
    - Ter mais de uma região (disaster recovery).
    - Ter mais de uma nuvem (disaster recovery).
    - Padrão Transactional Outbox: Funciona como um buffer local para mensagens que precisam ser enviadas para um message broker. As mensagens são armazenadas em uma tabela de banco de dados local (a "outbox") e são enviadas para o message broker em lotes, garantindo que as mensagens sejam enviadas de forma confiável e consistente, mesmo em caso de falhas no sistema.
    ![alt text](image-4.png)

### Garantias de Recebimento
**Exemplo RabbitMQ**
- Auto Ack = false e commit manual: Garante que a mensagem só será removida da fila quando o consumidor confirmar que a processou com sucesso. Se o consumidor falhar antes de confirmar, a mensagem permanecerá na fila e poderá ser reprocessada por outro consumidor. Essa é uma excelente abordagem para garantir a resiliência e a integridade dos dados.
- Auto Ack = true: A mensagem é removida da fila assim que é entregue ao consumidor, independentemente de o consumidor processá-la com sucesso ou não. Isso pode levar à perda de mensagens se o consumidor falhar durante o processamento. Essa abordagem é mais rápida, mas menos confiável. Evite trabalhar com Auto Ack = true em sistemas onde a integridade dos dados é crítica.
- Prefetch Count: Limita o número de mensagens que um consumidor pode receber antes de confirmar o processamento. Isso ajuda a evitar que um consumidor fique sobrecarregado com muitas mensagens ao mesmo tempo, melhorando a resiliência e a eficiência do sistema.

### Idempotência e Políticas de Fallback
- Saber lidar com mensagens duplicadas.
- Garantir que o processamento de uma mensagem possa ser repetido sem efeitos colaterais indesejados.
- Estratégias:
    - Utilizar identificadores únicos para mensagens (por exemplo, UUIDs) para rastrear e deduplicar mensagens.
    - Implementar lógica de compensação para desfazer efeitos colaterais de mensagens processadas anteriormente.
    - Adotar padrões de design como o "Circuit Breaker" para lidar com falhas temporárias em serviços externos.
  - Independência. Ex: Banco de Dados
  - Políticas claras de fallback.
    - Ex: Retornar dados em cache quando o serviço principal está indisponível.
    - Ex: Retornar uma resposta padrão ou vazia quando o serviço está indisponível.
    - Ex: Redirecionar para um serviço alternativo ou de backup.

### Observabilidade
- Monitoramento: Coleta e análise de métricas para entender o desempenho e a saúde do sistema.
- Logging: Registro detalhado de eventos e atividades do sistema para facilitar a depuração e a análise de problemas.
- Tracing: Rastreio de requisições através dos diferentes serviços para entender o fluxo e identificar gargalos ou falhas.
- Ferramentas comuns: Prometheus, Grafana, ELK Stack (Elasticsearch, Logstash, Kibana), Jaeger, Zipkin, etc.
- Alertas: Configuração de alertas para notificar a equipe sobre problemas críticos no sistema.
- Dashboards: Criação de dashboards para visualizar métricas e logs em tempo real.
- Análise de causa raiz: Utilização de dados coletados para investigar e resolver problemas de forma eficaz.
- Cultura de observabilidade: Incentivar a equipe a adotar práticas de observabilidade como parte do desenvolvimento e operação do sistema.
- **APM (Application Performance Monitoring):** Ferramentas como New Relic, Datadog, AppDynamics, etc., que fornecem monitoramento detalhado do desempenho da aplicação, incluindo tempos de resposta, taxas de erro, e outros indicadores chave de performance (KPIs).
- **Tracing distribuído:** Ferramentas como Jaeger, Zipkin, OpenTelemetry, etc., que permitem rastrear requisições através de múltiplos serviços, ajudando a identificar gargalos e falhas em sistemas distribuídos.
![alt text](image-5.png)
- **Métricas personalizadas:** Coleta de métricas específicas do negócio ou da aplicação, como número de usuários ativos, transações processadas, etc., para entender melhor o impacto do sistema no negócio.
- **Spans personalizados:** Permitem adicionar informações contextuais adicionais às requisições rastreadas, facilitando a análise e depuração. É possível identificar qual a função/método específico dentro do serviço que está causando lentidão ou falhas, ajudando a localizar e resolver problemas de forma mais eficiente.
- **Open Telemetry:** Uma iniciativa de código aberto que fornece uma coleção de ferramentas, APIs e SDKs para instrumentação, geração, coleta e exportação de dados de telemetria (métricas, logs e traces) de aplicações. O OpenTelemetry é suportado por uma ampla comunidade e é adotado por muitas empresas e projetos de código aberto. Ele oferece suporte a várias linguagens de programação, incluindo Java, JavaScript, Python, Go, C#, Ruby, PHP, entre outras. O OpenTelemetry é compatível com várias ferramentas de monitoramento e análise, como Prometheus, Jaeger, Zipkin, Grafana, entre outras. Ele também suporta vários protocolos de exportação de dados, como OTLP (OpenTelemetry Protocol), Jaeger Thrift, Zipkin JSON, entre outros. O OpenTelemetry é uma ferramenta poderosa para melhorar a observabilidade e a resiliência de sistemas distribuídos e microsserviços.

### Referências
- Exponential backoff and jitter: https://aws.amazon.com/pt/blogs/architecture/exponential-backoff-and-jitter/
- Retry: Remédio ou Veneno? https://www.youtube.com/watch?v=1MkPpKPyBps
- OTEL - https://opentelemetry.io/

