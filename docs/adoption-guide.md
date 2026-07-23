# Guia de adoção

Use o kit quando o projeto será mantido por humanos e agentes, mudanças precisam ser auditáveis e o time quer specs, diagnosis specs, reports e gates explícitos. Evite para spikes descartáveis ou projetos que não desejam enforcement em CI.

## Sequência verificável de adoção

Cada passo tem um comando e uma evidência esperada.

### 1. Instalar a CLI

```bash
go install ./cmd/gakit
gakit version
```

Evidência: versão impressa. Pré-requisito: Go 1.26.4+.

### 2. Criar o projeto

```bash
gakit create ./novo-projeto --slug novoprojeto --title "Novo Projeto" --module github.com/acme/novo-projeto --description "..." --author "Acme"
```

Evidência: árvore renderizada sem placeholders, git inicializado. Use `gakit template list --tree` antes, se quiser auditar o que será gerado.

### 3. Executar os checks iniciais

```bash
cd novo-projeto
make setup
make check-compliance
make test
```

Evidência: `check-compliance: ok` e testes verdes. Esses são os mesmos gates do CI gerado.

### 4. Diagnosticar e interpretar

```bash
gakit diagnose --path ./novo-projeto --report-only
```

Como interpretar o relatório de um starter recém-gerado:

- AI-first, Governança, DX, Security e Documentação devem pontuar alto: medem os artefatos de governança renderizados.
- Hexagonal Architecture e OpenAPI reportam ausências por design: `pkg/`, `internal/` e contrato OpenAPI só nascem sob spec aprovada. Não trate esses findings como defeito do bootstrap.
- Score alto mede baseline de governança, não maturidade de produto.

Para gate objetivo em CI: `gakit diagnose --path . --min-score 80 --report-only` (exit code 1 abaixo do mínimo).

### 5. Escolher a primeira jornada

No projeto gerado, abra `docs/journeys/README.md` e escolha:

- primeira contribuição e configuração da baseline;
- executar a fase 0 do roadmap (`phase_autopilot`);
- registrar um pedido no backlog e abrir uma trilha SDD;
- preparar release, ADR e changelog.

## Adotar em repositório existente

`gakit diagnose --path .` funciona em qualquer repositório Go. Use o relatório para priorizar lacunas por pilar e `--plan-prompt-only --out ./reports` para gerar um prompt de plano de correção SDD dual-spec a partir dos achados. O catálogo do que pode ser portado está em [capabilities.md](capabilities.md).

## O que a adoção não entrega

Os projetos gerados vêm com backlog governado, ADRs, políticas de release e autopilot dual prontos para enforcement — mas tudo isso é processo operado por agentes com evidência versionada. Nenhuma capability de produto, serviço hospedado ou stack de infraestrutura é instalada; ver a seção "Fora de escopo" de [capabilities.md](capabilities.md).
