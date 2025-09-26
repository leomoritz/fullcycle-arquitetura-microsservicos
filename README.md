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