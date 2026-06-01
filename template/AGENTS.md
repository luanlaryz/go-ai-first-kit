# AGENTS.md

## 1. Project mission

- `{{PROJECT_SLUG}}` existe para construir uma biblioteca em Go inspirada no nucleo do {{UPSTREAM_NAME}}.
- O objetivo principal e atingir paridade funcional com o framework base do {{UPSTREAM_NAME}}.
- A API publica deve ser idiomatica em Go, mesmo quando o conceito de origem vier do {{UPSTREAM_NAME}}.
- Serviços hospedados obrigatórios e control planes permanecem fora do escopo inicial.
- O repositorio segue Spec Driven Development rigoroso.

## 2. Source of truth

- As specs em `specs/` sao a fonte de verdade para requisitos por modulo.
- `specs/000-project-mission.md` define a missão e as regras globais do projeto.
- `specs/001-non-goals.md` define o que nao pode entrar no escopo.
- `specs/010-feature-matrix.md` deve ser usada como checklist de cobertura, prioridade, risco e paridade.
- `specs/020-repository-architecture.md` define as fronteiras entre `pkg/`, `internal/`, `test/` e `examples/`.
- Specs especificas de modulo, como `specs/030-agent.md` e `specs/031-app-instance.md`, governam o contrato daquele modulo.
- `specs/680-phase-25-interactive-sdd-autopilot-foundation.md` e `specs/681-phase-25-interactive-sdd-autopilot-foundation-diagnosis.md` governam o modo `interactive_sdd_autopilot`.
- `automation/AUTOPILOT.md`, `automation/RUNBOOK.md`, `automation/ROADMAP.json` e `automation/PHASE_STATE.json` governam o modo `phase_autopilot`.
- `automation/INTERACTIVE_AUTOPILOT.md`, `automation/INTERACTIVE_RUNBOOK.md` e `automation/INTERACTIVE_STATE.json` governam trilhas interativas de feature request.
- `docs/backlog/Backlog.md` e o backlog canonico para itens governados; fontes historicas em `docs/backlog/**` nao viram itens implementaveis sem triagem pela skill `skills/26-backlog-item-intake/SKILL.md`.
- `docs/decisions/` (ADRs) registra decisoes de arquitetura e governanca; `docs/release-versioning-policy.md`, `docs/release-notes-policy.md` e `docs/release-checklist.md` governam versionamento, changelog e release notes alinhados a `skills/22-release-versioning-governance/SKILL.md`.
- Se houver conflito entre implementacao existente e spec aprovada, a spec prevalece.

## 3. Required execution workflow

- Sempre ler as specs relevantes antes de implementar, refatorar ou revisar comportamento.
- Identificar explicitamente qual spec e qual secao cobrem a tarefa antes de alterar codigo.
- Nunca implementar feature fora da spec sem registrar a lacuna.
- Se a spec nao existir, estiver incompleta ou ambigua para a tarefa, parar a implementacao e registrar o gap de especificacao.
- Nunca alterar API publica sem atualizar a spec correspondente.
- Toda feature implementada deve ser rastreavel para pelo menos uma spec.
- Toda mudanca de comportamento deve incluir atualizacao ou adicao de testes.
- Toda mudanca relevante deve ser validada contra a feature matrix como checklist de cobertura.
- Toda trilha nova deve nascer com duas specs complementares: uma spec de construcao e uma spec de diagnostico.
- Nao considerar uma trilha nova suficientemente spec driven se existir apenas spec de construcao sem spec de diagnostico.
- Ao propor uma nova trilha, o agente deve verificar se ambas as specs existem. Se faltar qualquer uma delas, deve refinar o pedido ate fechar as duas antes de seguir para implementacao.
- Para solicitacoes interativas de feature, evolucao, bug, refatoracao ou docs, usar `interactive_sdd_autopilot`: intake, requisitos, decisao de spec, dual-spec, implementacao, diagnostico, report e gate.
- Para registrar pedidos, gaps, bugs, diagnosticos, recomendacoes ou ideias como backlog, usar `skills/26-backlog-item-intake/SKILL.md`, deduplicar em `docs/backlog/Backlog.md`, classificar escopo e specs, e nao promover fontes historicas em lote para itens `BLG-*`.
- Para decisoes de release/versionamento e changelog, usar `skills/22-release-versioning-governance/SKILL.md` com `docs/release-versioning-policy.md`, `docs/release-notes-policy.md` e `docs/release-checklist.md`; registrar decisoes de arquitetura/governanca nao-obvias como ADR em `docs/decisions/`.
- Nao usar `automation/ROADMAP.json` nem `automation/PHASE_STATE.json` para representar trilhas interativas; esses arquivos pertencem ao `phase_autopilot`.
- Em trilhas interativas, usar `skills/23-{{PROJECT_SLUG}}-sdd-autopilot/SKILL.md` alem da skill de dominio aplicavel.
- Para diagnosticos baseados em `@kodus/agent-readiness`, usar `skills/24-agent-readiness-governance/SKILL.md` e filtrar achados para o escopo de lib/framework Go.
- Para planejamento de fases SDD, geracao de prompts de autopilot, specs de construcao e diagnostico, reconciliacao, readiness/release gates ou qualquer continuidade ultra-rigida evidence-first, usar `skills/25-{{PROJECT_SLUG}}-ultra-rigid-sdd/SKILL.md`.
- Em trilhas interativas, o report da etapa e a fonte de verdade para avanco; testes verdes ou narrativa nao satisfazem gate.
- Parar imediatamente quando uma stop condition de `automation/STOP_CONDITIONS.md` ocorrer.
- Antes de commitar codigo Go, rodar `make fmt` para garantir formatacao `gofmt`. O CI rejeita PRs com arquivos fora do padrao gofmt.
- Ao abrir PR, usar `.github/PULL_REQUEST_TEMPLATE.md` como corpo base obrigatorio.
- Preferir `scripts/create-pr-from-template.sh` em vez de `gh pr create` direto para reduzir drift no corpo do PR.
- Se um PR for criado sem todas as secoes obrigatorias do template, o agente deve corrigi-lo antes de encerrar a tarefa.
- O corpo real do PR deve passar na validacao de `make check-pr-body` quando executado em `pull_request`.
- PRs Dependabot de atualizacao de dependencias sao excecao formal ao corpo completo do template, desde que a excecao esteja registrada em `docs/ai/compliance-exceptions.md` e os demais checks de CI continuem obrigatorios.

## 4. Architectural rules

- Manter separacao clara entre `pkg` e `internal`.
- `pkg/` contem contratos publicos estaveis e pontos de extensao importaveis.
- `internal/` contem runtime, wiring, validacao, state machine e detalhes operacionais nao exportados.
- Nao expor tipos de `internal/*` em assinaturas exportadas, erros publicos ou exemplos.
- Preferir interfaces pequenas e composicao.
- Evitar acoplamento ciclico entre pacotes.
- Evitar dependencias externas desnecessarias.
- Nao introduzir dependencia arquitetural em servicos hospedados para o runtime funcionar localmente.
- `pkg/app` e a unica ponte publica permitida para `internal/runtime`, conforme a arquitetura do repositorio.

## 5. Public API rules

- API publica deve ser pequena, intencional e orientada a contratos observaveis.
- Priorizar `context.Context` como fronteira padrao para execucao, cancelamento e lifecycle.
- Nomes, tipos e assinaturas em Go nao precisam copiar o {{UPSTREAM_NAME}} literalmente.
- Diferencas em relacao ao {{UPSTREAM_NAME}} so sao aceitaveis quando forem idiomaticas em Go e estiverem explicitadas na spec correspondente ou na feature matrix.
- Nunca alterar API publica sem atualizar a spec correspondente e os testes de contrato afetados.
- Sempre documentar erros observaveis, comportamento de cancelamento e ordem de execucao quando fizerem parte do contrato.
- Preferir construtores claros, options objetivas e tipos exportados apenas quando necessario.

## 6. Testing and conformance rules

- Sempre adicionar ou atualizar testes ao mudar comportamento.
- Testes de conformidade devem ser tratados como parte do contrato, nao como opcionais.
- Toda implementacao deve ter pelo menos um caminho de teste de compatibilidade conceitual ou comportamental.
- Casos de sucesso, falha, cancelamento e limites observaveis devem ser cobertos quando a spec exigir.
- Testes em `pkg/*` devem validar contratos publicos.
- Testes em `internal/*` devem validar detalhes de runtime e corretude operacional.
- Suites em `test/conformance` devem usar apenas a superficie publica em `pkg/*`.
- Se uma mudanca nao puder ser coberta por teste imediatamente, o gap deve ser declarado explicitamente com justificativa tecnica.

## Dual-spec rule for new tracks

- Toda trilha nova deve possuir duas specs normativas e rastreaveis.
- A spec de construcao define objetivo, escopo, fora de escopo, arquitetura, contratos, testes e criterios de aceite.
- A spec de diagnostico define sinais observaveis, sintomas de falha, hipoteses, metricas, logs, traces, health checks, consultas operacionais, troubleshooting e criterios de confirmacao.
- Nao iniciar implementacao de trilha nova sem as duas specs.
- Se existir spec de construcao sem spec de diagnostico, tratar como gap de especificacao.
- O agente deve refinar o pedido ate fechar ambas as specs antes de gerar prompts de execucao.

## 7. Spec compliance rules

- Nunca implementar feature fora da spec sem registrar a lacuna.
- Nunca assumir comportamento nao documentado quando a spec exigir semantica observavel.
- Se o codigo divergir da spec, corrigir o codigo ou atualizar a spec antes de seguir.
- Se a implementacao precisar de comportamento novo, primeiro atualizar ou propor a spec.
- Toda PR, patch ou entrega deve deixar claro qual spec foi atendida e quais criterios de aceitacao foram cobertos.
- Quando uma feature for parcialmente implementada, registrar explicitamente o que ficou pendente e qual secao da spec ainda nao foi atendida.

## 8. {{UPSTREAM_NAME}} parity rules

- O objetivo e paridade funcional, nao copia textual de API.
- Comparar comportamento por entrada, saida, efeitos colaterais controlados, erros e ordem de operacoes relevantes.
- Diferencas em relacao ao {{UPSTREAM_NAME}} so sao aceitaveis quando forem idiomaticas em Go e estiverem explicitadas na spec ou na feature matrix.
- Nao aceitar divergencia "conveniente" apenas para simplificar a implementacao.
- Nao introduzir nada de {{UPSTREAM_OPS_NAME}}.
- Quando houver diferenca intencional de comportamento, registrar a justificativa e o impacto na paridade.

## 9. Documentation rules

- Atualizar documentacao relevante sempre que a mudanca alterar comportamento observavel, API publica ou fluxo operacional importante.
- Nunca alterar API publica sem atualizar a spec correspondente.
- Usar `README.md` para contexto geral do projeto, nao para substituir specs normativas.
- Manter a feature matrix atual quando uma feature mudar de status, prioridade pratica ou cobertura.
- Documentar gaps restantes, riscos e divergencias intencionais.
- Relacionar codigo novo ou alterado com a spec correspondente de forma rastreavel.

## 10. Expected output format for code tasks

- Explicar quais specs foram lidas e quais secoes governaram a mudanca.
- Explicar quais arquivos foram alterados.
- Explicar as decisoes tomadas e os tradeoffs relevantes.
- Explicar quais testes foram adicionados, atualizados ou executados.
- Explicar gaps restantes, limitacoes conhecidas e o que ficou fora da entrega.
- Se a tarefa nao puder ser implementada por falta de spec, dizer isso explicitamente e apontar o gap de especificacao.
- Se houver divergencia intencional em relacao ao {{UPSTREAM_NAME}}, apontar a spec ou item da feature matrix que autoriza a diferenca.

## 11. Interactive SDD autopilot

- Use `interactive_sdd_autopilot` quando uma solicitacao inicial precisar virar trilha executavel de SDD.
- O pedido inicial e intake, nao spec aprovada.
- Primeiro normalize tipo, objetivo, escopo conhecido, incertezas, specs candidatas e skills candidatas.
- Se houver ambiguidade material, pergunte antes de editar.
- Decida entre amendar spec existente e criar nova trilha dual-spec usando `skills/05-{{PROJECT_SLUG}}-spec-architect/SKILL.md`.
- Crie ou refine spec de construcao e spec de diagnostico antes de implementar comportamento novo.
- Atualize `automation/INTERACTIVE_STATE.json` para registrar etapa, specs, report, classificacao, decisao, retries e bloqueio.
- Avance automaticamente somente quando o report da etapa declarar `classification = PASS` e decisao final autorizar avanco ou conclusao.
- Ao encerrar uma iteracao interativa, responder tambem com: `report lido`, `classificacao`, `decisao`, `estado atualizado` e `proxima etapa ou motivo de bloqueio`.


## Skill mapping

- Sempre começar por índice de skills: [`skills/00-skill-index/SKILL.md`](skills/00-skill-index/SKILL.md).
- Arquitetura/dependências/pastas: [`skills/01-hexagonal-architecture/SKILL.md`](skills/01-hexagonal-architecture/SKILL.md).
- Tenant/auth/quota: [`skills/02-tenant-auth-quotas/SKILL.md`](skills/02-tenant-auth-quotas/SKILL.md).
- ACK assíncrono/use cases de comando: [`skills/03-async-command-processing/SKILL.md`](skills/03-async-command-processing/SKILL.md).
- SSE/WS/replay por `seq`: [`skills/04-streaming-sse-ws/SKILL.md`](skills/04-streaming-sse-ws/SKILL.md).
- Refinamento de pedido em specs de construcao e diagnostico, arquitetura, decisao entre amend/new spec e geracao de prompt para Cursor/Codex: [`skills/05-{{PROJECT_SLUG}}-spec-architect/SKILL.md`](skills/05-{{PROJECT_SLUG}}-spec-architect/SKILL.md).
- Roteamento de provider/model por {{DOMAIN_ACTOR}}: [`skills/06-multi-provider-model-routing/SKILL.md`](skills/06-multi-provider-model-routing/SKILL.md).
- Postgres/pgx/migrations: [`skills/07-postgres-pgx-migrate/SKILL.md`](skills/07-postgres-pgx-migrate/SKILL.md).
- Redis cache/streams/TTL: [`skills/08-redis-cache-streams/SKILL.md`](skills/08-redis-cache-streams/SKILL.md).
- Worker/SQS FIFO/DLQ/idempotência: [`skills/09-sqs-fifo-idempotency/SKILL.md`](skills/09-sqs-fifo-idempotency/SKILL.md).
- HTTP Gin/OpenAPI: [`skills/10-http-gin-openapi/SKILL.md`](skills/10-http-gin-openapi/SKILL.md).
- gRPC/interceptors/health: [`skills/11-grpc-interceptors-health/SKILL.md`](skills/11-grpc-interceptors-health/SKILL.md).
- Logs/traces/metrics: [`skills/12-observability-zap-otel-prom/SKILL.md`](skills/12-observability-zap-otel-prom/SKILL.md).
- Padrões Go e concorrência segura: [`skills/13-go-idiomatic-effective-go/SKILL.md`](skills/13-go-idiomatic-effective-go/SKILL.md).
- Config/flags com Viper: [`skills/14-config-viper-flags/SKILL.md`](skills/14-config-viper-flags/SKILL.md).
- DevX/CI/pre-commit/changelog: [`skills/15-devx-ci-precommit-changelog/SKILL.md`](skills/15-devx-ci-precommit-changelog/SKILL.md).
- Object Calisthenics (Go): [`skills/16-object-calisthenics/SKILL.md`](skills/16-object-calisthenics/SKILL.md).
- SOLID em Go + ports/adapters: [`skills/17-solid-go-ports/SKILL.md`](skills/17-solid-go-ports/SKILL.md).
- OWASP API Security + CI checks: [`skills/18-security-owasp-api/SKILL.md`](skills/18-security-owasp-api/SKILL.md).
- Prompt Injection / LLM safety: [`skills/19-prompt-injection-llm-safety/SKILL.md`](skills/19-prompt-injection-llm-safety/SKILL.md).
- Estratégia de testes/regressão/carga/containers: [`skills/20-testing-strategy-regression-load-containers/SKILL.md`](skills/20-testing-strategy-regression-load-containers/SKILL.md).
- Documentação open source / godoc / README / CONTRIBUTING / CHANGELOG / examples: [`skills/21-documentation-open-source/SKILL.md`](skills/21-documentation-open-source/SKILL.md).
- Governança de release/versionamento: [`skills/22-release-versioning-governance/SKILL.md`](skills/22-release-versioning-governance/SKILL.md).
- Autopilot SDD interativo para feature requests: [`skills/23-{{PROJECT_SLUG}}-sdd-autopilot/SKILL.md`](skills/23-{{PROJECT_SLUG}}-sdd-autopilot/SKILL.md).
- Governança de diagnósticos `agent-readiness`: [`skills/24-agent-readiness-governance/SKILL.md`](skills/24-agent-readiness-governance/SKILL.md).
- SDD ultra-rígido / prompts / fases / gates evidence-first: [`skills/25-{{PROJECT_SLUG}}-ultra-rigid-sdd/SKILL.md`](skills/25-{{PROJECT_SLUG}}-ultra-rigid-sdd/SKILL.md).
- Backlog governado (intake, triagem, classificação, prompts autopilot): [`skills/26-backlog-item-intake/SKILL.md`](skills/26-backlog-item-intake/SKILL.md).
