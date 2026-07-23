# Catálogo de capacidades do go-ai-first-kit

Este catálogo lista o que o kit e os projetos gerados realmente entregam, como ativar cada capacidade e qual evidência confirma seu uso. Ele existe para evitar dois erros comuns: tratar guidance como componente instalado e tratar processo governado como automação sem operador.

## Como ler este catálogo

Cada capacidade declara um tipo de verificação:

- `automático`: existe um comando ou script executável cuja saída confirma a capacidade.
- `processual`: a capacidade é um workflow governado operado por um agente (humano ou AI); a evidência são specs, reports e arquivos de estado versionados, não um daemon rodando.
- `guia`: a capacidade é orientação normativa (skills, políticas, checklists); nada de código ou infraestrutura foi instalado por ela.

E cada capacidade declara um estado:

- `baseline entregue`: o artefato existe no template e é renderizado em todo projeto criado.
- `evidência a produzir`: o starter entrega o processo; cada projeto precisa gerar sua própria evidência (reports, specs novas, itens de backlog).
- `fora de escopo`: intencionalmente não entregue; ver a seção final.

Dois avisos permanentes:

- Os autopilots são workflows governados operados por agentes seguindo runbooks versionados. Não existe runner hospedado, daemon ou serviço executando fases sozinho.
- `gakit diagnose` é um sinal de qualidade com score ponderado. Ele não certifica produção nem substitui specs, reports e decisão humana.

## Capacidades da ferramenta (kit)

### Criar projeto — `gakit create`

- Tipo: `automático`. Estado: `baseline entregue`.
- O que entrega: renderiza o `template/` embarcado, substitui placeholders em paths e conteúdos, renomeia `go.mod.tmpl` para `go.mod`, aplica permissões executáveis em scripts e inicializa git.
- Como começar: `gakit create ./myapp --slug myapp --title "My App" --module github.com/acme/myapp` (ver [cli-reference.md](cli-reference.md)).
- Evidência: diretório criado com a árvore completa; `make check-compliance` e `make test` passam no projeto gerado.
- Limites: não cria capability de produto; o projeto nasce com governança, não com API.

### Inspecionar o template — `gakit template list`

- Tipo: `automático`. Estado: `baseline entregue`.
- O que entrega: inventário do template embarcado em lista, árvore ou JSON.
- Como começar: `gakit template list --tree`.
- Evidência: saída do comando; útil para auditar o que será gerado antes do `create`.

### Diagnosticar maturidade — `gakit diagnose`

- Tipo: `automático`. Estado: `baseline entregue`.
- O que entrega: avaliação de sete pilares com score `0.6 * coverage + 0.4 * quality` por pilar e score global ponderado, com relatório terminal, Markdown e JSON.
- Pilares e pesos: DX (0.10), AI-first (0.20), Security (0.15), Hexagonal Architecture (0.15), OpenAPI (0.10), Documentação (0.15), Governança (0.15).
- Como começar: `gakit diagnose --path <dir> --report-only`; use `--plan-prompt-only --out ./reports` para gerar um prompt de plano de correção SDD.
- Evidência: relatório com score, findings por severidade e snippets.
- Limites: em um starter recém-gerado, os pilares Hexagonal e OpenAPI apontam ausências por design — `pkg/`, `internal/` e contrato OpenAPI só nascem sob spec aprovada. Score alto mede artefatos de governança, não maturidade de produto.

### Bootstrap por prompt — `prompt-bootstrap-go-ai-first.md`

- Tipo: `automático`. Estado: `baseline entregue`.
- O que entrega: prompt auto-contido com todo o template em blocos `<file path="...">`, para agentes sem acesso à CLI.
- Como começar: colar o conteúdo em um agente com permissão de escrita e informar os parâmetros.
- Evidência: o arquivo é gerado exclusivamente por `scripts/build-master-prompt.sh` e validado por teste (`--check`).

### Scripts legados de renderização

- Tipo: `automático`. Estado: `baseline entregue` (fallback).
- O que entrega: `scripts/init.sh` e `scripts/render-template.sh` renderizam o template sem a CLI.
- Limites: mantidos por compatibilidade; o fluxo primário é `gakit`.

## Capacidades entregues nos projetos gerados

O detalhamento operacional de cada uma vive no projeto gerado, em `docs/ai/capabilities.md` (renderizado a partir de [template/docs/ai/capabilities.md](../template/docs/ai/capabilities.md)). Resumo:

### Governança para agentes

- Tipo: `guia` com enforcement parcial `automático`. Estado: `baseline entregue`.
- Artefatos: `AGENTS.md` (constituição), `.cursor/rules/` (6 regras modulares), `.cursor/hooks/` (gofmt pós-edição), `skills/00-skill-index/SKILL.md` e 26 skills numeradas.
- Limites: as skills de infraestrutura (02–04, 06–12, 14, 18–19) são guidance herdada para quando o produto exigir aquela stack; nenhuma dependência de Redis, SQS, Gin, gRPC ou Postgres é instalada pelo starter.

### SDD dual-spec

- Tipo: `processual`. Estado: `baseline entregue` + `evidência a produzir`.
- Artefatos: `specs/000`, `001`, `010`, `020` e o par `680/681` da fase 0; skill 05 (spec architect) e skill 25 (ultra-rigid SDD).
- Evidência a produzir: cada trilha nova exige spec de construção e spec de diagnóstico antes de implementar.

### Autopilot de roadmap (`phase_autopilot`)

- Tipo: `processual`. Estado: `baseline entregue` + `evidência a produzir`.
- Artefatos: `automation/AUTOPILOT.md`, `automation/RUNBOOK.md`, `automation/ROADMAP.json` (fase 0), `automation/PHASE_STATE.json`, `automation/STOP_CONDITIONS.md`.
- Evidência a produzir: report da fase no caminho declarado pelo roadmap, com `classification` e `decision` exatas do gate.

### Autopilot interativo (`interactive_sdd_autopilot`)

- Tipo: `processual`. Estado: `baseline entregue` + `evidência a produzir`.
- Artefatos: `automation/INTERACTIVE_AUTOPILOT.md`, `automation/INTERACTIVE_RUNBOOK.md`, `automation/INTERACTIVE_STATE.json`, `docs/interactive-sdd-autopilot.md`, `docs/feature-request-lifecycle.md`, skill 23.
- Evidência a produzir: intake registrado, dual-spec da trilha, report com `PASS` e estado coerente.

### Backlog governado

- Tipo: `processual`. Estado: `baseline entregue` + `evidência a produzir`.
- Artefatos: `docs/backlog/Backlog.md` (canônico, vazio por design) e skill 26 com templates de item, regras de classificação e prompts de autopilot.

### ADR, release e changelog

- Tipo: `processual` com `guia`. Estado: `baseline entregue` + `evidência a produzir`.
- Artefatos: `docs/decisions/` (ADR seed 0001), `docs/release-versioning-policy.md`, `docs/release-notes-policy.md`, `docs/release-checklist.md`, `CHANGELOG.md` Keep a Changelog, skill 22.

### Qualidade, segurança e compliance

- Tipo: `automático`. Estado: `baseline entregue`.
- Artefatos: `Makefile` (fmt, lint, vet, test, race, coverage, test-security), `scripts/check-compliance.sh`, `check-secrets.sh`, `check-pr-body.sh`, `check-file-size.sh`, `.pre-commit-config.yaml`, `.github/workflows/ci.yml`, `test/security/`.
- Limites: a baseline de segurança é mínima (presença de suite, secrets scan e govulncheck); não é uma suite ampla de regressão de safety.

### Descoberta e inventário

- Tipo: `guia`. Estado: `baseline entregue`.
- Artefatos: `docs/README.md` (hub), `docs/ai/capabilities.md` (catálogo), `docs/journeys/` (jornadas humanas), `docs/ai/maturity-inventory.md` (baseline/evidência/limites), `docs/ai/ops-guide.md` (execução).

## Fora de escopo

Intencionalmente não entregues pelo kit nem pelo template:

- Capability de produto: API pública, `pkg/`, `internal/`, `examples/`, contrato OpenAPI. Nascem apenas sob spec aprovada no projeto.
- Serviços hospedados, control planes, dashboards, runners gerenciados ou deploy tooling obrigatório.
- Publicação de documentação (GitHub Pages ou site): divergência intencional registrada em [govolt-sync-diagnosis.md](govolt-sync-diagnosis.md).
- Stack de infraestrutura das skills herdadas (Redis, SQS, Postgres, Gin, gRPC, Viper, OTel): as skills orientam a implementação futura, não instalam nada.
