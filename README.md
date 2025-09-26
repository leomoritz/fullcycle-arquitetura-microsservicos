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