# Prompt Bootstrap Go AI-First

Você é um agente de engenharia criando um projeto Go novo com baseline AI-first completa.

## Parâmetros Obrigatórios

Antes de escrever arquivos, colete:

- `PROJECT_SLUG`: slug do projeto e pacote raiz Go seguro, por exemplo `myapp`.
- `PROJECT_TITLE`: nome humano do projeto, por exemplo `My App`.
- `MODULE_PATH`: módulo Go, por exemplo `github.com/acme/myapp`.
- `PROJECT_DESCRIPTION`: descrição objetiva do projeto.
- `AUTHOR_NAME`: autor/mantenedor inicial.
- `LICENSE_NAME`: licença, default `MIT`.
- `UPSTREAM_NAME`: referência/upstream opcional, default `none`.

## Regras Obrigatórias

1. Leia este prompt inteiro antes de criar arquivos.
2. Crie a árvore exatamente conforme os blocos `<file path="...">`.
3. Substitua placeholders `{{...}}` em paths e conteúdos.
4. Aplique permissão executável em scripts `.sh` e `.cursor/hooks/*.sh`.
5. Se `go.mod.tmpl` existir, renomeie para `go.mod` após renderizar.
6. Rode `make check-compliance`, `make fmt-check`, `make vet` e `make test` quando possível.
7. Responda com parâmetros usados, arquivos criados, comandos executados, validações passadas e gaps restantes.

## Arquivos

<file path=".cursor/hooks/gofmt-after-edit.sh">
#!/bin/bash
input=$(cat)
filepath=$(echo "$input" | jq -r '.path // empty')

if [[ -n "$filepath" && "$filepath" == *.go ]]; then
  gofmt -w "$filepath" 2>/dev/null
fi

echo '{}'
exit 0
</file>

<file path=".cursor/hooks.json">
{
  "version": 1,
  "hooks": {
    "afterFileEdit": [
      {
        "command": ".cursor/hooks/gofmt-after-edit.sh",
        "matcher": "\\.go$"
      }
    ]
  }
}
</file>

<file path=".cursor/rules/{{PROJECT_SLUG}}-autopilot-core.mdc">
---
description: {{PROJECT_SLUG}} autopilot core modes and evidence-first execution
alwaysApply: true
---

# {{PROJECT_SLUG}} autopilot core

- Trate `phase_autopilot` e `interactive_sdd_autopilot` como modos diferentes.
- Use `automation/ROADMAP.json` e `automation/PHASE_STATE.json` somente para o roadmap fixo.
- Use `automation/INTERACTIVE_AUTOPILOT.md`, `automation/INTERACTIVE_RUNBOOK.md` e `automation/INTERACTIVE_STATE.json` para trilhas interativas.
- Leia `AGENTS.md`, a spec governante e `skills/00-skill-index/SKILL.md` antes de editar.
- Para trilhas interativas, leia tambem `skills/23-{{PROJECT_SLUG}}-sdd-autopilot/SKILL.md`.
- Nunca avance por narrativa, intencao, plano futuro, TODO ou ausencia aparente de erro.
- Declare sempre `objetivo`, `specs lidas`, `skills aplicadas`, `arquivos alterados`, `comandos executados`, `testes executados` e `gaps restantes`.
</file>

<file path=".cursor/rules/{{PROJECT_SLUG}}-dual-spec-enforcement.mdc">
---
description: {{PROJECT_SLUG}} dual-spec enforcement for new tracks
alwaysApply: true
---

# {{PROJECT_SLUG}} dual-spec enforcement

- Exija uma spec de construcao e uma spec de diagnostico para toda trilha nova.
- Nao implemente uma trilha nova quando qualquer uma das duas specs estiver ausente.
- Se a mudanca pertencer a um contrato existente, amende a spec existente e registre as secoes governantes.
- Se a mudanca criar novo bounded context, contrato principal ou diagnostico separado, abra uma trilha dual-spec.
- Use `skills/05-{{PROJECT_SLUG}}-spec-architect/SKILL.md` para decidir entre amendar spec existente e criar nova trilha.
- Pare quando objetivo, escopo, fora de escopo, comportamento observavel, testes ou criterio diagnostico estiverem materialmente ambiguos.
- Nao trate pedido inicial, conversa, plano ou prompt como substituto de spec versionada.
</file>

<file path=".cursor/rules/{{PROJECT_SLUG}}-interactive-feature-intake.mdc">
---
description: {{PROJECT_SLUG}} interactive feature request intake
alwaysApply: true
---

# {{PROJECT_SLUG}} interactive feature intake

- Trate uma solicitacao inicial de feature, evolucao, bug, refatoracao ou docs como intake.
- Normalize o pedido antes de implementar: tipo, objetivo, escopo conhecido, incertezas, specs candidatas e skills candidatas.
- Pergunte antes de editar quando houver ambiguidade material em objetivo, escopo, contrato, diagnostico, teste ou criterio de aceite.
- Registre a decisao entre amendar spec existente e criar nova trilha dual-spec.
- Crie ou refine a spec de construcao antes de implementar comportamento novo.
- Crie a spec de diagnostico antes de executar diagnostico de uma trilha nova.
- Atualize `automation/INTERACTIVE_STATE.json` durante intake, implementacao, diagnostico, report review, bloqueio e conclusao.
- Conclua a solicitacao somente quando todas as etapas tiverem evidencia ou report com gate aprovado.
</file>

<file path=".cursor/rules/{{PROJECT_SLUG}}-phase-gates-state.mdc">
---
description: {{PROJECT_SLUG}} phase gates, reports and state files
alwaysApply: true
---

# {{PROJECT_SLUG}} phase gates and state

- Trate report versionado como fonte de verdade para avanco.
- Avance somente com `classification = PASS` e decisao final compativel com o gate atual.
- Para `phase_autopilot`, atualize apenas `automation/PHASE_STATE.json` conforme `automation/RUNBOOK.md`.
- Para `interactive_sdd_autopilot`, atualize apenas `automation/INTERACTIVE_STATE.json` conforme `automation/INTERACTIVE_RUNBOOK.md`.
- Nao altere `automation/ROADMAP.json` para inserir trilhas interativas.
- Nao altere `automation/PHASE_STATE.json` para representar feature request interativa.
- Preserve retries, bloqueio, ultimo report e ultima decisao ao parar por stop condition.
- Pare se report, classificacao, decisao ou estado estiverem ausentes, ambiguos ou inconsistentes.
</file>

<file path=".cursor/rules/{{PROJECT_SLUG}}-stop-conditions.mdc">
---
description: {{PROJECT_SLUG}} stop condition enforcement
alwaysApply: true
---

# {{PROJECT_SLUG}} stop conditions

- Pare imediatamente quando qualquer stop condition de `automation/STOP_CONDITIONS.md` ocorrer.
- Para roadmap fixo, mantenha `current_phase` na fase afetada e marque `automation/PHASE_STATE.json` como bloqueado.
- Para trilha interativa, mantenha `current_step` na etapa afetada e marque `automation/INTERACTIVE_STATE.json` como bloqueado.
- Nunca contorne stop condition removendo criterio de aceite, mudando gate, reordenando roadmap ou mascarando gap com texto narrativo.
- Pare quando testes obrigatorios falharem, report estiver ausente, report for `PARTIAL`, `FAIL` ou `BLOCKED`, ou houver decisao humana pendente.
- Pare quando a correcao exigir capability fora da spec, {{UPSTREAM_OPS_NAME}}, hosted service, dashboard, deploy tooling ou breaking change sem aprovacao.
- Ao parar, informe evidencia observada, stop condition acionada e acao humana recomendada.
</file>

<file path=".cursor/rules/{{PROJECT_SLUG}}.mdc">
---
description: {{PROJECT_SLUG}} spec-driven baseline for AI-assisted contributions
alwaysApply: true
---

# {{PROJECT_SLUG}} baseline

- Leia `AGENTS.md`, a spec governante e `skills/00-skill-index/SKILL.md` antes de propor edicao.
- Cite a spec e a secao governante antes de alterar comportamento, API publica ou docs operacionais.
- Nao implemente feature fora de spec. Se a spec faltar ou estiver ambigua, pare e registre o gap.
- Atualize testes no mesmo patch quando o comportamento mudar.
- Mantenha `pkg/*` livre de dependencias em `internal/*`, exceto pela ponte publica permitida em `pkg/app`.
- Nao declare capacidade operacional sem artefato real correspondente.
- Para feature requests interativas, aplique `interactive_sdd_autopilot`, `skills/23-{{PROJECT_SLUG}}-sdd-autopilot/SKILL.md` e as rules modulares de autopilot.
- Nao use `automation/ROADMAP.json` nem `automation/PHASE_STATE.json` para trilhas interativas.
- Responda sempre com: `objetivo`, `specs lidas`, `skills aplicadas`, `arquivos alterados`, `comandos executados`, `testes executados` e `gaps restantes`.
</file>

<file path=".editorconfig">
root = true

[*]
charset = utf-8
end_of_line = lf
insert_final_newline = true
trim_trailing_whitespace = true
indent_style = space
indent_size = 2

[*.go]
indent_style = tab
indent_size = 4
</file>

<file path=".file-size-exceptions">
# One relative Go file path per line, with justification in review/report.
</file>

<file path=".github/PULL_REQUEST_TEMPLATE.md">
## Objetivo

Descreva o resultado observavel desta mudanca.

## Specs lidas

- [ ] `AGENTS.md`
- [ ] spec(s) governante(s) citadas abaixo

Liste as specs e secoes aplicaveis:

- `specs/...`

## Skills aplicadas

- `skills/...`

## Agent Readiness

- Report: `none` ou `docs/reports/agent-readiness/...`
- Source: `none` ou comando/report `@kodus/agent-readiness`
- {{PROJECT_TITLE}}-filtered result: `READY`, `READY_WITH_WARNINGS`, `NOT_READY` ou `not applicable`
- PR impact: `BLOCKS_PR`, `DOES_NOT_BLOCK_PR` ou `not applicable`

Regra de bloqueio: `optional_for_{{PROJECT_SLUG}}` e `out_of_scope_for_{{PROJECT_SLUG}}` nao bloqueiam; apenas `worth_it_for_{{PROJECT_SLUG}}` que afete esta PR e tenha `PR impact = BLOCKS_PR` bloqueia merge.

Achados relevantes, recomendacoes opcionais e achados fora de escopo:

- `none`, ou liste apenas o que afeta esta PR

## Arquivos alterados

- liste os arquivos principais alterados neste PR

## Impacto em docs e superficie publica

- `none`, ou descreva o impacto relevante quando existir

## Comandos executados

- liste comandos executados alem dos checks padrao, quando houver

## Testes executados

- [ ] `make fmt-check`
- [ ] `make lint`
- [ ] `make vet`
- [ ] `make check-compliance`
- [ ] `make test`
- [ ] testes adicionais relevantes

## Excecoes formais

- `none`, ou referencie `docs/ai/compliance-exceptions.md`

## Gaps restantes

- descreva gaps, tradeoffs ou itens fora do escopo

## Checklist

- [ ] Li `AGENTS.md` e as specs governantes antes de editar
- [ ] Atualizei testes ou declarei excecao formal quando a mudanca tocou comportamento
- [ ] Atualizei docs e spec correspondente quando a mudanca tocou API publica ou fluxo operacional
- [ ] Nao deixei marcador vazio em assets operacionais
- [ ] Declarei impacto em docs ou superficie publica, ou marquei `none`
- [ ] A descricao deste PR lista gaps restantes de forma explicita
</file>

<file path=".github/dependabot.yml">
version: 2
updates:
  - package-ecosystem: gomod
    directory: "/"
    schedule:
      interval: weekly
    open-pull-requests-limit: 5
    commit-message:
      prefix: chore
      include: scope
    groups:
      go-dependencies:
        patterns:
          - "*"

  - package-ecosystem: github-actions
    directory: "/"
    schedule:
      interval: weekly
    open-pull-requests-limit: 5
    commit-message:
      prefix: chore
      include: scope
</file>

<file path=".github/workflows/ci.yml">
name: ci

on:
  pull_request:
  push:

jobs:
  validate:
    runs-on: ubuntu-latest

    steps:
      - name: Checkout
        uses: actions/checkout@v6

      - name: Setup Go
        uses: actions/setup-go@v6
        with:
          go-version-file: go.mod
          cache: true

      - name: Install staticcheck
        run: |
          go install honnef.co/go/tools/cmd/staticcheck@latest
          echo "$(go env GOPATH)/bin" >> "$GITHUB_PATH"

      - name: Fmt check
        run: make fmt-check

      - name: Lint
        run: make lint

      - name: Vet
        run: make vet

      - name: Check compliance
        run: make check-compliance

      - name: Check PR body
        if: github.event_name == 'pull_request'
        run: make check-pr-body

      - name: Security baseline
        run: make test-security

      - name: Check file size
        run: make check-file-size

      - name: Go test
        run: make test

      - name: Race
        run: make race
</file>

<file path=".gitignore">
# OS/editor
.DS_Store
.idea/
.vscode/

# Local env
.env
.env.*
!.env.example

# Go
bin/
*.test
coverage.out

# Generated
dist/
.tmp/
</file>

<file path=".pre-commit-config.yaml">
repos:
  - repo: local
    hooks:
      - id: fmt-fix
        name: gofmt auto-fix
        entry: bash -c 'make fmt && git diff --exit-code --quiet -- "*.go"'
        language: system
        pass_filenames: false
      - id: lint
        name: make lint
        entry: make lint
        language: system
        pass_filenames: false
      - id: vet
        name: make vet
        entry: make vet
        language: system
        pass_filenames: false
      - id: check-compliance
        name: make check-compliance
        entry: make check-compliance
        language: system
        pass_filenames: false
      - id: check-file-size
        name: make check-file-size
        entry: make check-file-size
        language: system
        pass_filenames: false
      - id: secrets-check
        name: make secrets-check
        entry: make secrets-check
        language: system
        pass_filenames: false
</file>

<file path="AGENTS.md">
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
</file>

<file path="CHANGELOG.md">
# Changelog

## [Unreleased]

### Added
- Bootstrap AI-first governance baseline.
</file>

<file path="CONTRIBUTING.md">
# Contributing

Use `docs/ai/task-input-format.md` para abrir tarefas para agentes. Rode `make setup` e `make test` localmente.
</file>

<file path="LICENSE">
{{LICENSE_NAME}}

Copyright (c) 2026 {{AUTHOR_NAME}}
</file>

<file path="Makefile">
GO ?= go
GOFMT ?= gofmt
STATICCHECK ?= $(shell command -v staticcheck 2>/dev/null)
STATICCHECK_VERSION ?= latest
GOVULNCHECK_VERSION ?= latest
GOFILES := $(shell find . -type f -name '*.go' -not -path './vendor/*')

.PHONY: setup bootstrap test test-security security-tests vulncheck secrets-check lint fmt fmt-check race coverage vet check-compliance check-pr-body check-file-size

setup:
	@command -v $(GO) >/dev/null 2>&1 || { echo "go not found; install Go 1.26.3+ from https://go.dev/dl/"; exit 1; }
	$(GO) mod download
	$(GO) install honnef.co/go/tools/cmd/staticcheck@$(STATICCHECK_VERSION)

bootstrap: setup

test:
	$(GO) test ./...

test-security: security-tests secrets-check vulncheck

security-tests:
	$(GO) test ./test/security/...

vulncheck:
	$(GO) run golang.org/x/vuln/cmd/govulncheck@$(GOVULNCHECK_VERSION) ./...

secrets-check:
	./scripts/check-secrets.sh

lint:
ifeq ($(STATICCHECK),)
	@echo "staticcheck not found; install with: go install honnef.co/go/tools/cmd/staticcheck@$(STATICCHECK_VERSION)"
	@exit 1
endif
	$(STATICCHECK) ./...

fmt:
	$(GOFMT) -w $(GOFILES)

fmt-check:
	@unformatted="$$(gofmt -l $(GOFILES))"; \
	if [ -n "$$unformatted" ]; then \
		echo "unformatted Go files:"; \
		echo "$$unformatted"; \
		exit 1; \
	fi

race:
	$(GO) test -race ./...

coverage:
	$(GO) test -coverprofile=coverage.out ./...

vet:
	$(GO) vet ./...

check-compliance:
	./scripts/check-compliance.sh

check-pr-body:
	./scripts/check-pr-body.sh

check-file-size:
	./scripts/check-file-size.sh
</file>

<file path="README.md">
# {{PROJECT_TITLE}}

{{PROJECT_DESCRIPTION}}

## AI-first governance

Antes de editar, leia `AGENTS.md`, `skills/00-skill-index/SKILL.md` e a spec governante em `specs/`.
</file>

<file path="SECURITY.md">
# Security

Este projeto roda `make test-security`, `make secrets-check` e `make vulncheck` como baseline mínima.
</file>

<file path="automation/AUTOPILOT.md">
# {{PROJECT_SLUG}} Autopilot

## 1. Missao

O autopilot do `{{PROJECT_SLUG}}` executa o roadmap fase por fase como uma maquina de estados controlada por evidencia.

A missao e:

1. executar o roadmap do `{{PROJECT_SLUG}}` fase por fase;
2. operar como maquina de estados baseada em `automation/PHASE_STATE.json`;
3. usar o report da fase como fonte de verdade para decisao de avanco;
4. nunca avancar por narrativa, intencao, resumo otimista ou ausencia de erro aparente;
5. preservar o fluxo Spec Driven Development definido em `AGENTS.md`;
6. respeitar a regra dual-spec: toda fase precisa de spec de construcao e spec de diagnostico;
7. parar imediatamente quando qualquer stop condition for encontrada.

O autopilot nao decide que uma fase esta pronta. O autopilot le o report da fase e so avanca quando o report contem exatamente o gate esperado para a fase atual.

## 1.1 Modos De Autopilot

O repositorio possui dois modos de automacao governada:

1. `phase_autopilot`: modo deste documento, fechado sobre `automation/ROADMAP.json` e `automation/PHASE_STATE.json`;
2. `interactive_sdd_autopilot`: modo interativo definido em `automation/INTERACTIVE_AUTOPILOT.md`, iniciado por solicitacao de feature ou evolucao e registrado em `automation/INTERACTIVE_STATE.json`.

O modo `interactive_sdd_autopilot` e aditivo. Ele nao pode reordenar, simplificar, remover ou reinterpretar o roadmap fixo deste documento.

Trilhas interativas nao devem ser registradas em `automation/ROADMAP.json`.

Solicitacoes interativas nao devem alterar `automation/PHASE_STATE.json`.

## 2. Fonte De Verdade

A execucao deve seguir esta hierarquia:

1. `AGENTS.md`;
2. spec de construcao da fase atual;
3. spec de diagnostico da fase atual;
4. `automation/ROADMAP.json`;
5. `automation/PHASE_STATE.json`;
6. `automation/STOP_CONDITIONS.md`;
7. report gerado para a fase atual.

Quando houver conflito entre implementacao e spec aprovada, a spec prevalece.

Quando houver conflito entre narrativa da conversa e report versionado, o report versionado prevalece para decisao de gate.

Quando o report estiver ausente, incompleto ou inconsistente com a fase atual, a fase fica bloqueada.

## 3. Roadmap Governado

O autopilot deve executar exatamente estas fases, nesta ordem:

1. Fase 20: `memory-adapter-expansion-and-storage-parity`
2. Fase 21: `agent-invocation-and-conversation-contract-parity`
3. Fase 22: `workflow-state-persistence-and-resume-parity`
4. Fase 23: `resumable-streams-and-playground-runtime-parity`
5. Fase 24: `hosting-integration-and-{{UPSTREAM_NAME_LOWER}}-app-migration-readiness`

O autopilot nao pode inventar fases, remover fases, reordenar fases, agrupar fases ou antecipar trabalho de fase futura.

## 4. Maquina De Estados

O estado persistente vive em `automation/PHASE_STATE.json`.

Campos obrigatorios:

1. `current_phase`: fase que deve ser executada agora;
2. `status`: estado operacional da fase atual;
3. `last_completed_phase`: ultima fase concluida com gate aprovado;
4. `retry_count`: numero de retries corretivos ja consumidos na fase atual;
5. `last_report`: caminho do ultimo report lido;
6. `last_decision`: decisao final extraida do ultimo report;
7. `blocked`: indica se o autopilot esta bloqueado;
8. `block_reason`: motivo objetivo do bloqueio.

Estados permitidos para `status`:

1. `pending`;
2. `running`;
3. `diagnosing`;
4. `blocked`;
5. `completed`;
6. `finished`.

Transicoes permitidas:

1. `pending` -> `running`
2. `running` -> `diagnosing`
3. `diagnosing` -> `completed`
4. `diagnosing` -> `blocked`
5. `running` -> `blocked`
6. `completed` -> `pending` da proxima fase
7. `completed` da Fase 24 -> `finished`

Qualquer transicao fora dessas deve ser tratada como inconsistencia de estado.

## 5. Ciclo Obrigatorio Por Fase

Para cada fase, o autopilot deve executar o ciclo completo abaixo, sem pular etapas.

### 5.1 Criar Ou Validar A Spec Da Fase

O autopilot deve localizar a spec de construcao definida em `automation/ROADMAP.json`.

Se a spec existir:

1. ler a spec completa;
2. identificar objetivo, escopo, fora de escopo, contratos, testes e criterios de aceite;
3. verificar se a spec cobre materialmente a implementacao esperada para a fase;
4. registrar no resultado da iteracao quais secoes governaram a fase.

Se a spec nao existir:

1. criar a spec no caminho exato definido em `automation/ROADMAP.json`;
2. respeitar `AGENTS.md`;
3. respeitar a fase e o nome definidos no roadmap;
4. nao criar fase nova;
5. nao alterar a ordem do roadmap;
6. nao implementar antes de a spec existir.

### 5.2 Implementar A Spec Da Fase

A implementacao deve ficar restrita ao escopo da fase atual.

O autopilot deve:

1. ler o codigo existente relevante antes de editar;
2. preservar contratos publicos salvo quando a spec exigir e justificar explicitamente;
3. atualizar testes no mesmo patch quando houver mudanca de comportamento;
4. atualizar documentacao operacional ou publica quando houver impacto observavel;
5. manter `pkg/*` livre de dependencia em `internal/*`, exceto pela ponte publica permitida em `pkg/app`;
6. manter rastreabilidade entre spec, codigo, testes e docs;
7. nao abrir capability nova fora da fase atual;
8. nao mascarar gap com documentacao narrativa.

### 5.3 Criar Ou Validar A Spec De Diagnostico

O autopilot deve localizar a spec de diagnostico definida em `automation/ROADMAP.json`.

Se a spec existir:

1. ler a spec completa;
2. identificar sinais observaveis, modos de falha, comandos, evidencias esperadas e criterio de confirmacao;
3. verificar que a auditoria consegue classificar a fase como `PASS`, `PARTIAL`, `FAIL` ou `BLOCKED`;
4. verificar que a auditoria exige decisao final explicita.

Se a spec nao existir:

1. criar a spec no caminho exato definido em `automation/ROADMAP.json`;
2. vincular a spec de diagnostico a fase atual;
3. incluir sinais observaveis, sintomas de falha, hipoteses, metricas, logs, traces, health checks quando aplicavel, comandos de verificacao, troubleshooting e criterio de confirmacao;
4. nao executar diagnostico antes de a spec existir.

### 5.4 Executar O Diagnostico

O diagnostico deve seguir a spec de diagnostico da fase atual.

O autopilot deve executar, no minimo:

1. verificacao de aderencia entre spec de construcao e implementacao;
2. verificacao de aderencia entre spec de diagnostico e evidencias coletadas;
3. `go test ./...`;
4. `go test -race ./...`;
5. validacao de OpenAPI quando a fase tocar API, runtime HTTP, playground, app hosting ou contrato documentado por OpenAPI;
6. testes especificos exigidos pela spec da fase;
7. leitura de docs, exemplos, feature matrix e reports relacionados quando a fase exigir;
8. classificacao objetiva da fase;
9. geracao do report no caminho exato definido em `automation/ROADMAP.json`.

Falha em qualquer validacao obrigatoria e bloqueio, exceto quando a spec de diagnostico autorizar explicitamente classificacao diferente e o report documentar a justificativa com evidencia.

### 5.5 Localizar O Report Gerado

O autopilot deve localizar o report no caminho exato da fase atual:

1. Fase 20: `docs/reports/phase-20-memory-adapter-expansion-and-storage-parity-report.md`
2. Fase 21: `docs/reports/phase-21-agent-invocation-and-conversation-contract-parity-report.md`
3. Fase 22: `docs/reports/phase-22-workflow-state-persistence-and-resume-parity-report.md`
4. Fase 23: `docs/reports/phase-23-resumable-streams-and-playground-runtime-parity-report.md`
5. Fase 24: `docs/reports/phase-24-hosting-integration-and-{{UPSTREAM_NAME_LOWER}}-app-migration-readiness-report.md`

Report ausente e stop condition.

Report em caminho diferente nao satisfaz o gate.

Report inconsistente com a fase atual e stop condition.

### 5.6 Ler A Decisao Final

O autopilot deve ler o report completo e extrair a classificacao e a decisao final.

A classificacao deve ser uma destas:

1. `PASS`
2. `PARTIAL`
3. `FAIL`
4. `BLOCKED`

A decisao deve ser exatamente a decisao esperada da fase atual.

Gates esperados:

1. Fase 20:
   - `classification = PASS`
   - `decision = READY FOR PHASE 21`

2. Fase 21:
   - `classification = PASS`
   - `decision = READY FOR PHASE 22`

3. Fase 22:
   - `classification = PASS`
   - `decision = READY FOR PHASE 23`

4. Fase 23:
   - `classification = PASS`
   - `decision = READY FOR PHASE 24`

5. Fase 24:
   - `classification = PASS`
   - `decision = READY FOR MIGRATION EXECUTION`

Decisao ausente e stop condition.

Classificacao ausente e stop condition.

Decisao ambigua e stop condition.

### 5.7 Atualizar `automation/PHASE_STATE.json`

O autopilot deve atualizar `automation/PHASE_STATE.json` ao final de cada tentativa.

Quando a fase inicia:

1. `status = "running"`
2. `blocked = false`
3. `block_reason = ""`

Quando a fase entra em diagnostico:

1. `status = "diagnosing"`

Quando a fase passa no gate:

1. `last_completed_phase` recebe a fase concluida;
2. `last_report` recebe o caminho do report lido;
3. `last_decision` recebe a decisao final do report;
4. `retry_count = 0`;
5. `blocked = false`;
6. `block_reason = ""`;
7. se houver proxima fase, `current_phase` recebe a proxima fase e `status = "pending"`;
8. se a fase concluida for a Fase 24, `current_phase = 24` e `status = "finished"`.

Quando a fase bloqueia:

1. `current_phase` permanece na fase atual;
2. `status = "blocked"`;
3. `last_report` recebe o caminho do report lido, se existir;
4. `last_decision` recebe a decisao extraida, se existir;
5. `blocked = true`;
6. `block_reason` recebe motivo objetivo;
7. `retry_count` reflete os retries corretivos consumidos.

### 5.8 Avancar Apenas Se O Gate For Satisfeito

O autopilot so pode avancar quando todas as condicoes forem verdadeiras:

1. report existe no caminho esperado;
2. report corresponde a fase atual;
3. `classification == PASS`;
4. `decision` e exatamente a decisao esperada para a fase atual;
5. `go test ./...` passou;
6. `go test -race ./...` passou;
7. validacao OpenAPI passou quando aplicavel;
8. nao ha divergencia material entre spec e implementacao;
9. nao ha stop condition ativa;
10. `automation/PHASE_STATE.json` foi atualizado de forma coerente.

## 6. Regras De Nao Avanco

O autopilot nunca deve avancar se o report trouxer qualquer uma destas condicoes:

1. `PARTIAL`;
2. `FAIL`;
3. `BLOCKED`;
4. `NOT READY`;
5. decisao ausente;
6. classificacao ausente;
7. report ausente;
8. report inconsistente com a fase atual;
9. report com decisao diferente da esperada;
10. report que declara evidencia insuficiente;
11. report que indica testes obrigatorios nao executados;
12. report que indica divergencia material entre spec e implementacao;
13. report que depende de decisao humana pendente.

O autopilot tambem nao deve avancar por inferencia baseada em testes verdes se o report nao satisfizer o gate.

## 7. Retry Policy

Cada fase permite no maximo `2` retries corretivos.

Retry corretivo significa uma tentativa de corrigir falhas identificadas durante implementacao, testes, diagnostico ou leitura do report da fase atual.

Politica:

1. ao encontrar falha corrigivel na fase atual, incrementar `retry_count`;
2. executar a correcao dentro do escopo da fase atual;
3. reexecutar as validacoes obrigatorias;
4. regenerar ou atualizar o report da fase;
5. reler o gate;
6. avancar somente se o gate for satisfeito.

Se `retry_count` exceder `2`, o autopilot deve parar.

Ao parar por retries excedidos, deve produzir um relatorio curto de bloqueio contendo:

1. fase afetada;
2. spec de construcao;
3. spec de diagnostico;
4. report esperado;
5. retries consumidos;
6. falhas persistentes;
7. acao humana recomendada.

## 8. Disciplina De Escopo

O autopilot deve manter disciplina estrita de escopo.

Regras obrigatorias:

1. nao abrir capability nova fora da fase atual;
2. nao implementar itens de fase futura;
3. nao reordenar o roadmap;
4. nao remover criterios de aceite;
5. nao mascarar gap com documentacao narrativa;
6. preferir reconciliacao documental e testes quando o codigo ja estiver correto;
7. nao redefinir API publica silenciosamente;
8. nao alterar naming publico sem justificativa explicita;
9. nao renomear contrato publico para satisfazer preferencia estetica;
10. nao introduzir dependencia externa desnecessaria;
11. nao introduzir hosted service como requisito para runtime local;
12. nao declarar capacidade operacional sem artefato real correspondente;
13. nao tratar intencao, plano, comentario ou TODO como entrega concluida.

## 9. Validacoes Obrigatorias

As validacoes abaixo sao obrigatorias em todas as fases, salvo quando a spec de diagnostico registrar explicitamente que uma validacao nao se aplica a fase atual e justificar com evidencia.

Bloqueios obrigatorios:

1. `go test ./...` falhar;
2. `go test -race ./...` falhar;
3. OpenAPI quebrar quando a fase tocar API, runtime HTTP, playground, app hosting ou contrato documentado por OpenAPI;
4. diagnostico nao gerar report;
5. report estar ausente;
6. report nao conter decisao final explicita;
7. existir divergencia material entre spec e implementacao;
8. existir breaking change sem aprovacao explicita em spec;
9. existir rename ou naming inconsistente afetando contrato publico;
10. existir necessidade de decisao humana de arquitetura;
11. `automation/PHASE_STATE.json` divergir do roadmap;
12. specs de construcao e diagnostico da fase nao existirem apos a etapa de criacao ou validacao.

## 10. Evidence-First

Toda conclusao do autopilot deve apontar para evidencia versionada ou comando executado.

Evidencias aceitas:

1. spec versionada;
2. codigo versionado;
3. teste versionado;
4. saida de comando executado;
5. documentacao versionada;
6. feature matrix;
7. report da fase;
8. diff revisavel;
9. diagnostico versionado.

Evidencias nao aceitas como conclusao de fase:

1. intencao;
2. resumo narrativo;
3. plano futuro;
4. comentario sem teste;
5. TODO;
6. ausencia de erro visual;
7. suposicao de compatibilidade;
8. report de fase diferente;
9. report sem classificacao;
10. report sem decisao final.

## 11. Encerramento Da Automacao

O autopilot termina com sucesso somente quando a Fase 24 tiver:

1. report no caminho `docs/reports/phase-24-hosting-integration-and-{{UPSTREAM_NAME_LOWER}}-app-migration-readiness-report.md`;
2. `classification = PASS`;
3. `decision = READY FOR MIGRATION EXECUTION`;
4. `go test ./...` verde;
5. `go test -race ./...` verde;
6. nenhuma stop condition ativa;
7. `automation/PHASE_STATE.json` com `status = "finished"`;
8. `last_completed_phase = 24`.

Qualquer outra parada deve ser tratada como bloqueio operacional.
</file>

<file path="automation/INTERACTIVE_AUTOPILOT.md">
# {{PROJECT_SLUG}} Interactive SDD Autopilot

## 1. Missao

O `interactive_sdd_autopilot` conduz uma solicitacao inicial de feature, evolucao, bug, refatoracao ou documentacao por um ciclo Spec Driven Development completo.

A missao e:

1. receber e normalizar o pedido inicial;
2. levantar requisitos e lacunas materiais;
3. decidir entre amendar specs existentes ou abrir nova trilha dual-spec;
4. criar ou refinar spec de construcao;
5. criar spec de diagnostico complementar;
6. implementar somente comportamento coberto por spec;
7. executar diagnostico;
8. ler report;
9. avancar automaticamente apenas quando o gate passar;
10. parar em stop condition ou conclusao completa.

Este modo e aditivo ao `phase_autopilot`. Ele nao substitui o roadmap fixo.

## 2. Fonte De Verdade

A execucao interativa deve seguir esta hierarquia:

1. `AGENTS.md`;
2. `specs/680-phase-25-interactive-sdd-autopilot-foundation.md`;
3. `specs/681-phase-25-interactive-sdd-autopilot-foundation-diagnosis.md`;
4. `skills/23-{{PROJECT_SLUG}}-sdd-autopilot/SKILL.md`;
5. spec de construcao da trilha atual;
6. spec de diagnostico da trilha atual;
7. `automation/INTERACTIVE_STATE.json`;
8. `automation/STOP_CONDITIONS.md`;
9. report da etapa atual.

Quando houver conflito entre pedido inicial e spec versionada, a spec versionada prevalece.

Quando houver conflito entre narrativa da conversa e report versionado, o report prevalece para decisao de avanco.

## 3. Separacao Do Roadmap Fixo

`automation/ROADMAP.json` continua sendo exclusivo do `phase_autopilot`.

`automation/PHASE_STATE.json` continua sendo exclusivo do `phase_autopilot`.

O modo interativo deve usar `automation/INTERACTIVE_STATE.json` para registrar solicitacao, etapa, specs, report, classificacao, decisao, retries e bloqueios.

Trilhas interativas nao devem ser adicionadas a `automation/ROADMAP.json`.

## 4. Maquina De Estados Interativa

Estados permitidos para `status`:

1. `idle`;
2. `intake`;
3. `requirements`;
4. `build_spec`;
5. `diagnosis_spec`;
6. `implementation`;
7. `diagnosis`;
8. `report_review`;
9. `next_step_decision`;
10. `blocked`;
11. `completed`.

Transicoes permitidas:

1. `idle` -> `intake`;
2. `intake` -> `requirements`;
3. `requirements` -> `build_spec`;
4. `requirements` -> `blocked`;
5. `build_spec` -> `diagnosis_spec`;
6. `diagnosis_spec` -> `implementation`;
7. `implementation` -> `diagnosis`;
8. `implementation` -> `blocked`;
9. `diagnosis` -> `report_review`;
10. `diagnosis` -> `blocked`;
11. `report_review` -> `next_step_decision`;
12. `report_review` -> `blocked`;
13. `next_step_decision` -> `requirements`;
14. `next_step_decision` -> `implementation`;
15. `next_step_decision` -> `completed`;
16. qualquer etapa -> `blocked` quando houver stop condition.

Qualquer transicao fora dessas deve ser tratada como inconsistencia de estado.

## 5. Ciclo Obrigatorio

Para cada solicitacao, o Cursor deve executar:

1. intake;
2. levantamento de requisitos;
3. decisao de spec;
4. criacao ou refinamento da spec de construcao;
5. criacao da spec de diagnostico;
6. implementacao;
7. diagnostico;
8. geracao ou leitura do report;
9. decisao de avanco, retry, bloqueio ou conclusao.

Nenhuma etapa implementavel pode ser considerada concluida sem evidencia.

## 6. Gate De Avanco

O Cursor so pode avancar quando todas as condicoes forem verdadeiras:

1. spec de construcao existe e cobre a etapa;
2. spec de diagnostico existe e define o diagnostico;
3. diagnostico foi executado ou justificado pela spec;
4. report existe e corresponde a etapa atual;
5. `classification = PASS`;
6. `decision` autoriza o avanco ou a conclusao;
7. testes obrigatorios passaram ou foram declarados nao aplicaveis pela spec diagnostica;
8. nenhuma stop condition esta ativa;
9. `automation/INTERACTIVE_STATE.json` foi atualizado.

## 7. Retry Policy

Cada etapa interativa permite no maximo `2` retries corretivos.

Retry corretivo significa uma tentativa de corrigir falhas identificadas durante implementacao, diagnostico ou leitura do report da etapa atual.

Se `retry_count` exceder `2`, o Cursor deve bloquear a trilha e registrar acao humana recomendada.

## 8. Encerramento

A solicitacao termina com sucesso somente quando:

1. todas as etapas necessarias foram concluidas;
2. todos os reports exigidos existem;
3. a classificacao final e `PASS`;
4. a decisao final declara conclusao da solicitacao;
5. nao ha stop condition ativa;
6. `automation/INTERACTIVE_STATE.json` registra `status = "completed"`.

Qualquer outro encerramento deve ser tratado como bloqueio ou entrega parcial explicitamente registrada.
</file>

<file path="automation/INTERACTIVE_RUNBOOK.md">
# {{PROJECT_SLUG}} Interactive SDD Autopilot Runbook

## 1. Objetivo

Este runbook instrui o Cursor a operar o `interactive_sdd_autopilot` para transformar uma solicitacao inicial em uma trilha Spec Driven Development completa.

Use este runbook somente para trilhas interativas. Para roadmap fixo, use `automation/RUNBOOK.md`.

## 2. Leitura Obrigatoria Inicial

Antes de qualquer edicao, leia:

1. `AGENTS.md`;
2. `skills/00-skill-index/SKILL.md`;
3. `skills/05-{{PROJECT_SLUG}}-spec-architect/SKILL.md`;
4. `skills/23-{{PROJECT_SLUG}}-sdd-autopilot/SKILL.md`;
5. `specs/680-phase-25-interactive-sdd-autopilot-foundation.md`;
6. `specs/681-phase-25-interactive-sdd-autopilot-foundation-diagnosis.md`;
7. `automation/INTERACTIVE_AUTOPILOT.md`;
8. `automation/INTERACTIVE_STATE.json`;
9. `automation/STOP_CONDITIONS.md`.

Depois identifique:

1. `status`;
2. `current_request.request_id`;
3. `current_request.current_step`;
4. specs governantes;
5. spec de construcao da trilha;
6. spec de diagnostico da trilha;
7. report esperado;
8. `retry_count`;
9. estado de `blocked`.

## 3. Intake

Normalize a solicitacao inicial:

1. preserve o texto original;
2. classifique o tipo;
3. identifique objetivo;
4. identifique escopo conhecido;
5. liste incertezas materiais;
6. liste specs candidatas;
7. liste skills candidatas;
8. decida se precisa perguntar.

Se houver ambiguidade material, pergunte antes de editar.

## 4. Decisao De Spec

Use `skills/05-{{PROJECT_SLUG}}-spec-architect/SKILL.md` para decidir:

1. amendar spec existente;
2. criar nova trilha dual-spec.

Registre a decisao em resposta, docs ou report conforme a etapa.

## 5. Criar Ou Refinar Specs

Antes de implementar:

1. garanta spec de construcao suficiente;
2. garanta spec de diagnostico complementar;
3. registre secoes governantes;
4. identifique arquivos provaveis;
5. identifique testes e diagnosticos exigidos.

Nao implemente se qualquer spec obrigatoria estiver ausente ou materialmente incompleta.

## 6. Implementar

Durante a implementacao:

1. edite somente arquivos no escopo da trilha;
2. preserve `pkg/` e `internal/`;
3. nao altere `automation/ROADMAP.json` nem `automation/PHASE_STATE.json`;
4. atualize testes se comportamento mudar;
5. atualize docs quando fluxo operacional ou publico mudar;
6. mantenha evidencia rastreavel.

## 7. Diagnostico

Execute o diagnostico definido pela spec diagnostica da trilha.

Validacoes comuns:

1. aderencia entre spec de construcao e implementacao;
2. aderencia entre spec de diagnostico e evidencias;
3. testes especificos da trilha;
4. `go test ./...`;
5. `go test -race ./...`;
6. OpenAPI quando aplicavel;
7. verificacao de escopo proibido.

## 8. Report

Depois do diagnostico:

1. gere ou localize o report esperado;
2. leia o report completo;
3. extraia `classification`;
4. extraia `decision`;
5. extraia gaps e skips;
6. compare com o gate da etapa.

Nao avance se qualquer campo estiver ausente, ambiguo ou divergente.

## 9. Atualizar `INTERACTIVE_STATE.json`

Ao iniciar intake:

1. `mode = "interactive_sdd_autopilot"`;
2. `status = "intake"`;
3. `blocked = false`;
4. `block_reason = ""`.

Ao bloquear:

1. preserve `current_request`;
2. preserve `current_step`;
3. defina `status = "blocked"`;
4. defina `blocked = true`;
5. registre `block_reason`;
6. preserve report, classificacao e decisao se existirem.

Ao concluir:

1. defina `status = "completed"`;
2. defina `blocked = false`;
3. registre report final;
4. registre `last_completed_request`;
5. preserve criterios de conclusao.

## 10. Prompt Mestre

```text
Execute o interactive_sdd_autopilot do {{PROJECT_SLUG}} para esta solicitacao.

Leia primeiro:
- AGENTS.md
- skills/00-skill-index/SKILL.md
- skills/05-{{PROJECT_SLUG}}-spec-architect/SKILL.md
- skills/23-{{PROJECT_SLUG}}-sdd-autopilot/SKILL.md
- specs/680-phase-25-interactive-sdd-autopilot-foundation.md
- specs/681-phase-25-interactive-sdd-autopilot-foundation-diagnosis.md
- automation/INTERACTIVE_AUTOPILOT.md
- automation/INTERACTIVE_STATE.json
- automation/STOP_CONDITIONS.md

Comece por intake.
Levante requisitos.
Decida entre amendar spec existente ou abrir nova trilha dual-spec.
Crie ou refine spec de construcao.
Crie spec de diagnostico.
Implemente somente depois das specs suficientes.
Execute diagnostico.
Leia o report.
Avance somente se classification = PASS e decision autorizar.
Pare em qualquer stop condition.

Preserve automation/ROADMAP.json e automation/PHASE_STATE.json para o phase_autopilot.

Responda com:
- objetivo
- specs lidas
- skills aplicadas
- arquivos alterados
- comandos executados
- testes executados
- report lido
- classificacao
- decisao
- estado atualizado
- gaps restantes
- proxima etapa ou motivo de bloqueio
```

## 11. Retomada Apos Bloqueio

```text
Retome o interactive_sdd_autopilot do {{PROJECT_SLUG}} apos bloqueio.

Leia automation/INTERACTIVE_STATE.json e identifique:
- request_id
- current_step
- retry_count
- last_report
- last_classification
- last_decision
- blocked
- block_reason

Leia a spec de construcao, a spec de diagnostico, o report e a stop condition relacionada.
Corrija somente o que estiver dentro do escopo aprovado.
Nao altere ROADMAP.json ou PHASE_STATE.json para contornar o bloqueio.
Gere ou atualize o report.
Avance somente se classification = PASS e decision autorizar.
```
</file>

<file path="automation/INTERACTIVE_STATE.json">
{
  "mode": "interactive_sdd_autopilot",
  "status": "idle",
  "max_retries_per_step": 2,
  "last_completed_request": "",
  "blocked": false,
  "block_reason": "",
  "current_request": null,
  "smoke_history": []
}
</file>

<file path="automation/PHASE_STATE.json">
{
  "current_phase": 0,
  "status": "pending",
  "last_completed_phase": null,
  "retry_count": 0,
  "last_report": "",
  "last_decision": "",
  "blocked": false,
  "block_reason": ""
}
</file>

<file path="automation/ROADMAP.json">
{
  "project": "{{PROJECT_SLUG}}",
  "mode": "phase_autopilot",
  "max_retries_per_phase": 2,
  "phases": [
    {
      "id": 0,
      "name": "bootstrap-foundation",
      "spec_file": "specs/680-phase-0-bootstrap-foundation.md",
      "diagnosis_spec_file": "specs/681-phase-0-bootstrap-foundation-diagnosis.md",
      "report_file": "docs/reports/phase-0-bootstrap-foundation-report.md",
      "advance_only_if": {
        "classification": "PASS",
        "decision": "READY FOR FIRST FEATURE TRAIL"
      }
    }
  ]
}
</file>

<file path="automation/RUNBOOK.md">
# {{PROJECT_SLUG}} Autopilot Runbook

## 1. Objetivo

Este runbook instrui o Cursor a executar o roadmap automatizado do `{{PROJECT_SLUG}}` como maquina de estados, uma fase por vez, usando reports como fonte de verdade e gates explicitos para avanco.

O Cursor deve executar somente a fase atual indicada por `automation/PHASE_STATE.json`.

Este runbook governa apenas o modo `phase_autopilot`.

Para solicitacoes interativas de feature, evolucao, bug, refatoracao ou documentacao, use `automation/INTERACTIVE_RUNBOOK.md`, `automation/INTERACTIVE_AUTOPILOT.md` e `automation/INTERACTIVE_STATE.json`.

O modo interativo nao deve alterar `automation/ROADMAP.json` nem `automation/PHASE_STATE.json` para representar uma trilha de feature request.

## 2. Leitura Obrigatoria Inicial

Antes de qualquer implementacao, diagnostico ou atualizacao de estado, o Cursor deve ler:

1. `AGENTS.md`;
2. `skills/00-skill-index/SKILL.md`;
3. `automation/AUTOPILOT.md`;
4. `automation/ROADMAP.json`;
5. `automation/PHASE_STATE.json`;
6. `automation/STOP_CONDITIONS.md`.

Depois deve identificar:

1. `current_phase`;
2. fase correspondente em `automation/ROADMAP.json`;
3. spec de construcao da fase atual;
4. spec de diagnostico da fase atual;
5. report esperado da fase atual;
6. gate esperado da fase atual;
7. `retry_count`;
8. estado de `blocked`.

Se `blocked = true`, seguir a rotina de retomada apos bloqueio antes de executar qualquer nova fase.

## 3. Execucao Somente Da Fase Atual

O Cursor deve executar somente a fase indicada por `current_phase`.

O Cursor nao deve:

1. executar fase futura;
2. reexecutar fase ja concluida;
3. pular fase;
4. reordenar roadmap;
5. alterar gates;
6. trocar nomes de arquivos;
7. substituir report esperado;
8. alterar `last_completed_phase` sem report aprovado.

## 4. Procedimento Da Fase Atual

Para a fase atual, executar obrigatoriamente:

1. criar ou validar spec;
2. implementar a spec;
3. criar ou validar diagnosis spec;
4. executar diagnostico;
5. localizar e ler o report;
6. atualizar `automation/PHASE_STATE.json`;
7. avancar somente quando o gate for satisfeito;
8. parar imediatamente quando houver stop condition.

## 5. Criar Ou Validar Spec

Localizar `spec_file` da fase atual em `automation/ROADMAP.json`.

Se existir:

1. ler a spec completa;
2. registrar secoes governantes;
3. confirmar objetivo, escopo, fora de escopo, contratos, testes e criterios de aceite;
4. identificar arquivos provaveis de implementacao;
5. identificar docs e testes exigidos.

Se nao existir:

1. criar a spec no caminho exato definido por `spec_file`;
2. manter o nome da fase exatamente como definido em `automation/ROADMAP.json`;
3. escrever objetivo, motivacao, pergunta principal, escopo, fora de escopo, contratos, arquitetura, testes, criterios de aceite e limites;
4. nao implementar antes de a spec existir.

## 6. Implementar

Durante a implementacao:

1. ler codigo existente antes de editar;
2. alterar somente arquivos necessarios para a fase atual;
3. preservar arquitetura `pkg/` e `internal/`;
4. preservar contratos publicos salvo autorizacao explicita da spec;
5. adicionar ou atualizar testes quando houver mudanca de comportamento;
6. atualizar docs quando houver comportamento observavel, API publica ou fluxo operacional afetado;
7. manter feature matrix coerente quando maturidade, cobertura, prioridade ou paridade mudarem;
8. registrar gaps se a spec exigir algo que nao pode ser implementado com seguranca.

## 7. Criar Ou Validar Diagnosis Spec

Localizar `diagnosis_spec_file` da fase atual em `automation/ROADMAP.json`.

Se existir:

1. ler a spec completa;
2. identificar checklist de diagnostico;
3. identificar comandos obrigatorios;
4. identificar sinais observaveis;
5. identificar modos de falha;
6. identificar criterio de classificacao;
7. identificar formato da decisao final.

Se nao existir:

1. criar a spec no caminho exato definido por `diagnosis_spec_file`;
2. vincular a diagnostico a fase atual;
3. definir sinais observaveis, modos de falha, comandos, evidencias, troubleshooting, classificacao e decisao final;
4. nao executar diagnostico antes de a spec existir.

## 8. Executar Diagnostico

Executar o diagnostico conforme a spec de diagnostico da fase atual.

Validacoes obrigatorias:

1. aderencia entre spec de construcao e implementacao;
2. aderencia entre spec de diagnostico e evidencias;
3. `go test ./...`;
4. `go test -race ./...`;
5. validacao OpenAPI quando aplicavel;
6. testes especificos da fase;
7. docs e feature matrix quando aplicavel;
8. ausencia de breaking change nao aprovado;
9. ausencia de divergencia material entre spec e codigo.

O diagnostico deve gerar o report no caminho exato definido por `report_file`.

## 9. Localizar E Ler Report

Depois do diagnostico, localizar `report_file`.

O Cursor deve:

1. confirmar que o arquivo existe;
2. confirmar que o report corresponde a fase atual;
3. ler o report completo;
4. extrair `classification`;
5. extrair `decision`;
6. comparar com `advance_only_if.classification`;
7. comparar com `advance_only_if.decision`.

Nao avancar se qualquer campo estiver ausente, ambiguo ou divergente.

## 10. Atualizar `PHASE_STATE.json`

### 10.1 Ao Iniciar Fase

Definir:

1. `status = "running"`;
2. `blocked = false`;
3. `block_reason = ""`.

### 10.2 Ao Entrar Em Diagnostico

Definir:

1. `status = "diagnosing"`.

### 10.3 Ao Passar No Gate

Se `classification` e `decision` forem exatamente os esperados:

1. definir `last_completed_phase` como a fase atual;
2. definir `last_report` como o report lido;
3. definir `last_decision` como a decisao lida;
4. definir `retry_count = 0`;
5. definir `blocked = false`;
6. definir `block_reason = ""`;
7. se houver proxima fase, definir `current_phase` como a proxima fase;
8. se houver proxima fase, definir `status = "pending"`;
9. se a fase atual for 24, definir `status = "finished"` e manter `current_phase = 24`.

### 10.4 Ao Bloquear

Se houver stop condition:

1. manter `current_phase` na fase atual;
2. definir `status = "blocked"`;
3. definir `blocked = true`;
4. definir `block_reason` com motivo objetivo;
5. preservar ou atualizar `retry_count`;
6. definir `last_report` se houver report lido;
7. definir `last_decision` se houver decisao lida;
8. parar.

## 11. Politica De Avanco

Avancar somente quando todas as condicoes forem verdadeiras:

1. fase atual foi implementada ou validada;
2. spec de construcao existe;
3. spec de diagnostico existe;
4. diagnostico foi executado;
5. report foi gerado no caminho correto;
6. report foi lido;
7. `classification` e exatamente a esperada;
8. `decision` e exatamente a esperada;
9. `go test ./...` passou;
10. `go test -race ./...` passou;
11. OpenAPI passou quando aplicavel;
12. nenhuma stop condition esta ativa;
13. `automation/PHASE_STATE.json` foi atualizado.

## 12. Parada Por Stop Condition

Parar imediatamente quando qualquer item de `automation/STOP_CONDITIONS.md` ocorrer.

Ao parar por bloqueio, produzir:

1. resumo do bloqueio;
2. fase afetada;
3. retries ja consumidos;
4. acao humana recomendada.

O Cursor nao deve continuar o roadmap enquanto `blocked = true`.

## 13. Gates Por Fase

### Fase 20

Fase: `memory-adapter-expansion-and-storage-parity`

Spec:

`specs/580-phase-20-memory-adapter-expansion-and-storage-parity.md`

Diagnosis spec:

`specs/581-phase-20-memory-adapter-expansion-and-storage-parity-diagnosis.md`

Report:

`docs/reports/phase-20-memory-adapter-expansion-and-storage-parity-report.md`

Gate:

1. `classification = PASS`
2. `decision = READY FOR PHASE 21`

### Fase 21

Fase: `agent-invocation-and-conversation-contract-parity`

Spec:

`specs/600-phase-21-agent-invocation-and-conversation-contract-parity.md`

Diagnosis spec:

`specs/601-phase-21-agent-invocation-and-conversation-contract-parity-diagnosis.md`

Report:

`docs/reports/phase-21-agent-invocation-and-conversation-contract-parity-report.md`

Gate:

1. `classification = PASS`
2. `decision = READY FOR PHASE 22`

### Fase 22

Fase: `workflow-state-persistence-and-resume-parity`

Spec:

`specs/620-phase-22-workflow-state-persistence-and-resume-parity.md`

Diagnosis spec:

`specs/621-phase-22-workflow-state-persistence-and-resume-parity-diagnosis.md`

Report:

`docs/reports/phase-22-workflow-state-persistence-and-resume-parity-report.md`

Gate:

1. `classification = PASS`
2. `decision = READY FOR PHASE 23`

### Fase 23

Fase: `resumable-streams-and-playground-runtime-parity`

Spec:

`specs/640-phase-23-resumable-streams-and-playground-runtime-parity.md`

Diagnosis spec:

`specs/641-phase-23-resumable-streams-and-playground-runtime-parity-diagnosis.md`

Report:

`docs/reports/phase-23-resumable-streams-and-playground-runtime-parity-report.md`

Gate:

1. `classification = PASS`
2. `decision = READY FOR PHASE 24`

### Fase 24

Fase: `hosting-integration-and-{{UPSTREAM_NAME_LOWER}}-app-migration-readiness`

Spec:

`specs/660-phase-24-hosting-integration-and-{{UPSTREAM_NAME_LOWER}}-app-migration-readiness.md`

Diagnosis spec:

`specs/661-phase-24-hosting-integration-and-{{UPSTREAM_NAME_LOWER}}-app-migration-readiness-diagnosis.md`

Report:

`docs/reports/phase-24-hosting-integration-and-{{UPSTREAM_NAME_LOWER}}-app-migration-readiness-report.md`

Gate:

1. `classification = PASS`
2. `decision = READY FOR MIGRATION EXECUTION`

## 14. Resposta Final Obrigatoria Do Cursor

Ao encerrar cada iteracao, responder com:

1. objetivo;
2. specs lidas;
3. skills aplicadas;
4. arquivos alterados;
5. comandos executados;
6. testes executados;
7. report lido;
8. classificacao;
9. decisao;
10. estado atualizado;
11. gaps restantes;
12. se continuou, proxima fase;
13. se bloqueou, motivo e acao humana recomendada.

## 15. Prompt Mestre Para Iniciar O Autopilot

Use este prompt para iniciar a execucao automatizada do roadmap:

```text
Execute o autopilot do roadmap do projeto {{PROJECT_SLUG}}.

Regras obrigatorias:

1. Leia antes de qualquer acao:
   - AGENTS.md
   - skills/00-skill-index/SKILL.md
   - automation/AUTOPILOT.md
   - automation/ROADMAP.json
   - automation/PHASE_STATE.json
   - automation/STOP_CONDITIONS.md

2. Comece exatamente pela fase indicada em automation/PHASE_STATE.json.

3. Use automation/ROADMAP.json como lista fechada de fases. Nao invente fases, nao reordene fases, nao pule fases e nao simplifique o roadmap.

4. Execute somente a fase atual por vez.

5. Para cada fase, siga este ciclo completo:
   - criar ou validar a spec da fase;
   - implementar a spec da fase;
   - criar ou validar a spec de diagnostico;
   - executar o diagnostico;
   - localizar o report no caminho exato definido no roadmap;
   - ler o report completo;
   - extrair classification e decision;
   - atualizar automation/PHASE_STATE.json;
   - avancar somente se o gate da fase for satisfeito exatamente.

6. O report da fase e a fonte de verdade para avanco. Nunca avance por narrativa, intencao, resumo otimista, testes verdes isolados ou ausencia aparente de erro.

7. Avance somente quando:
   - classification == PASS;
   - decision for exatamente a decisao esperada da fase atual;
   - go test ./... passar;
   - go test -race ./... passar;
   - OpenAPI estiver valida quando aplicavel;
   - nao houver divergencia material entre spec e implementacao;
   - nenhuma stop condition estiver ativa.

8. Pare imediatamente se houver qualquer stop condition em automation/STOP_CONDITIONS.md, incluindo:
   - FAIL;
   - BLOCKED;
   - PARTIAL;
   - NOT READY;
   - retries excedidos;
   - go test ./... falhou;
   - go test -race ./... falhou;
   - OpenAPI invalida ou divergente;
   - report ausente;
   - report sem decisao final explicita;
   - breaking change sem aprovacao explicita;
   - divergencia material entre spec e codigo;
   - necessidade de decisao humana de arquitetura;
   - rename/naming inconsistente afetando contrato publico;
   - roadmap inconsistente com automation/PHASE_STATE.json.

9. Use no maximo 2 retries corretivos por fase. Depois disso, pare e registre bloqueio em automation/PHASE_STATE.json.

10. Ao bloquear, produza:
   - resumo do bloqueio;
   - fase afetada;
   - retries ja consumidos;
   - stop condition acionada;
   - evidencia observada;
   - acao humana recomendada.

11. Ao concluir uma fase com gate aprovado:
   - atualize last_completed_phase;
   - atualize last_report;
   - atualize last_decision;
   - zere retry_count;
   - marque blocked como false;
   - avance current_phase para a proxima fase;
   - deixe status como pending para a proxima fase.

12. Ao concluir a Fase 24 com:
   - classification = PASS;
   - decision = READY FOR MIGRATION EXECUTION;
   atualize automation/PHASE_STATE.json com status = finished e last_completed_phase = 24.

13. Continue automaticamente fase por fase ate concluir todas as fases ou atingir uma stop condition.

Formato obrigatorio da resposta ao final de cada iteracao:
- objetivo
- specs lidas
- skills aplicadas
- arquivos alterados
- comandos executados
- testes executados
- report lido
- classification
- decision
- estado atualizado
- gaps restantes
- proxima fase ou motivo de bloqueio
```

## 16. Prompt De Continuacao Normal

Use este prompt quando o Cursor ja estiver no meio do processo e precisar continuar de onde parou:

```text
Continue de onde parou no autopilot do {{PROJECT_SLUG}}.

Antes de qualquer acao, leia:
- AGENTS.md
- skills/00-skill-index/SKILL.md
- automation/AUTOPILOT.md
- automation/ROADMAP.json
- automation/PHASE_STATE.json
- automation/STOP_CONDITIONS.md

Retome exatamente da fase indicada em automation/PHASE_STATE.json.

Nao reinicie fases ja concluidas.
Nao pule gates.
Nao avance por narrativa.
Nao reordene o roadmap.
Nao altere os nomes das fases, specs, diagnosis specs ou reports.
Nao execute fase futura antes da fase atual passar no gate.

Se automation/PHASE_STATE.json indicar blocked = true, pare a continuacao normal e use o fluxo de retomada apos bloqueio.

Para a fase atual:
1. leia a spec de construcao definida em automation/ROADMAP.json;
2. crie a spec se ela ainda nao existir;
3. implemente ou continue a implementacao somente dentro do escopo da fase atual;
4. leia a spec de diagnostico definida em automation/ROADMAP.json;
5. crie a spec de diagnostico se ela ainda nao existir;
6. execute o diagnostico;
7. rode go test ./...;
8. rode go test -race ./...;
9. valide OpenAPI quando aplicavel;
10. gere ou localize o report no caminho exato definido no roadmap;
11. leia o report completo;
12. extraia classification e decision;
13. atualize automation/PHASE_STATE.json;
14. avance somente se classification == PASS e decision for exatamente a decisao esperada da fase atual.

Pare imediatamente se qualquer stop condition em automation/STOP_CONDITIONS.md ocorrer.

Use no maximo 2 retries corretivos por fase. Se os retries forem excedidos, registre bloqueio em automation/PHASE_STATE.json e pare.

Formato obrigatorio da resposta:
- objetivo
- specs lidas
- skills aplicadas
- arquivos alterados
- comandos executados
- testes executados
- report lido
- classification
- decision
- estado atualizado
- gaps restantes
- proxima fase ou motivo de bloqueio
```

## 17. Prompt De Retomada Apos Bloqueio

Use este prompt quando houver `blocked = true`, retries excedidos ou parada por stop condition:

```text
Retome o autopilot do {{PROJECT_SLUG}} apos bloqueio de forma segura.

Antes de qualquer acao, leia:
- AGENTS.md
- skills/00-skill-index/SKILL.md
- automation/AUTOPILOT.md
- automation/ROADMAP.json
- automation/PHASE_STATE.json
- automation/STOP_CONDITIONS.md

Inspecione primeiro o bloqueio registrado em automation/PHASE_STATE.json.

Nao resete o roadmap.
Nao apague historico de bloqueio.
Nao zere retry_count sem resolver o bloqueio.
Nao avance para fase futura enquanto blocked = true.
Nao reinicie fases ja concluidas.
Nao pule gates.
Nao altere ROADMAP.json para contornar bloqueio.

Procedimento obrigatorio:

1. Identifique:
   - current_phase;
   - status;
   - retry_count;
   - last_completed_phase;
   - last_report;
   - last_decision;
   - blocked;
   - block_reason.

2. Localize a fase atual em automation/ROADMAP.json.

3. Leia:
   - spec de construcao da fase atual;
   - spec de diagnostico da fase atual;
   - report da fase atual, se existir;
   - stop condition relacionada ao bloqueio.

4. Atue primeiro no desbloqueio da fase atual.

5. Corrija somente o que estiver dentro do escopo da fase atual.

6. Se o bloqueio exigir decisao humana de arquitetura, breaking change sem aprovacao, rename publico sensivel, alteracao de gate, reordenacao de roadmap ou capability fora da fase atual, pare e informe a acao humana recomendada.

7. Se o bloqueio for corrigivel:
   - incremente retry_count se esta for uma tentativa corretiva;
   - aplique a correcao;
   - rode go test ./...;
   - rode go test -race ./...;
   - valide OpenAPI quando aplicavel;
   - execute novamente o diagnostico da fase;
   - gere ou atualize o report no caminho exato definido no roadmap;
   - leia o report completo;
   - extraia classification e decision.

8. So desbloqueie automation/PHASE_STATE.json se:
   - classification == PASS;
   - decision for exatamente a decisao esperada da fase atual;
   - go test ./... passou;
   - go test -race ./... passou;
   - OpenAPI passou quando aplicavel;
   - a stop condition original foi resolvida;
   - nenhuma nova stop condition esta ativa.

9. Se o gate for satisfeito:
   - defina blocked = false;
   - limpe block_reason;
   - zere retry_count;
   - atualize last_completed_phase;
   - atualize last_report;
   - atualize last_decision;
   - avance current_phase para a proxima fase, ou marque status = finished se a fase concluida for 24.

10. Se o gate nao for satisfeito:
   - mantenha blocked = true;
   - mantenha current_phase na fase atual;
   - atualize block_reason;
   - preserve o historico relevante em last_report e last_decision;
   - pare.

11. Se retry_count exceder 2:
   - mantenha blocked = true;
   - defina status = blocked;
   - registre block_reason como retries corretivos excedidos;
   - pare.

Formato obrigatorio da resposta:
- objetivo
- specs lidas
- skills aplicadas
- arquivos alterados
- comandos executados
- testes executados
- bloqueio original
- acao de desbloqueio executada
- report lido
- classification
- decision
- estado atualizado
- gaps restantes
- proxima fase ou motivo de bloqueio persistente
```
</file>

<file path="automation/STOP_CONDITIONS.md">
# {{PROJECT_SLUG}} Autopilot Stop Conditions

## 1. Regra Geral

O autopilot deve parar imediatamente quando qualquer condicao deste documento for encontrada.

Ao parar, o autopilot deve:

1. manter `current_phase` na fase afetada;
2. definir `status = "blocked"` em `automation/PHASE_STATE.json`;
3. definir `blocked = true`;
4. registrar `block_reason` com motivo objetivo;
5. preservar `retry_count`;
6. preservar `last_report` e `last_decision` quando existirem;
7. produzir resumo curto do bloqueio;
8. nao avancar para a proxima fase.

Para `interactive_sdd_autopilot`, a mesma regra geral se aplica a `automation/INTERACTIVE_STATE.json`: manter a etapa afetada, definir `status = "blocked"`, preservar report/classificacao/decisao quando existirem, registrar `block_reason` e nao avancar para a proxima etapa.

## 2. Stop Conditions De Report

Parar imediatamente se:

1. o report da fase atual estiver ausente;
2. o report estiver em caminho diferente do definido em `automation/ROADMAP.json`;
3. o report for inconsistente com a fase atual;
4. o report nao tiver classificacao explicita;
5. o report nao tiver decisao final explicita;
6. o report trouxer `classification = FAIL`;
7. o report trouxer `classification = BLOCKED`;
8. o report trouxer `classification = PARTIAL`;
9. o report trouxer decisao contendo `NOT READY`;
10. o report trouxer decisao diferente da esperada para a fase atual;
11. o report depender de evidencia externa nao versionada;
12. o report declarar testes obrigatorios nao executados;
13. o report declarar diagnostico incompleto;
14. o report declarar divergencia material entre spec e codigo;
15. o report declarar decisao humana pendente.

## 3. Stop Conditions De Retry

Parar imediatamente se:

1. `retry_count` exceder `2`;
2. a mesma falha persistir apos `2` retries corretivos;
3. o retry exigiria alterar fase futura;
4. o retry exigiria remover criterio de aceite;
5. o retry exigiria reordenar roadmap;
6. o retry exigiria criar capability fora da fase atual;
7. o retry exigiria breaking change sem aprovacao explicita.

## 4. Stop Conditions De Testes

Parar imediatamente se:

1. `go test ./...` falhar;
2. `go test -race ./...` falhar;
3. teste de conformidade exigido pela fase falhar;
4. teste de contrato publico exigido pela fase falhar;
5. teste de integracao exigido pela fase falhar;
6. teste que comprova criterio de aceite estiver ausente;
7. teste estiver pulado sem justificativa aceita pela spec de diagnostico;
8. suite falhar por race condition;
9. suite falhar por dependencia real externa nao prevista;
10. suite exigir segredo, credencial ou servico hospedado fora do escopo.

## 5. Stop Conditions De OpenAPI

Parar imediatamente se:

1. OpenAPI estiver invalida;
2. OpenAPI estiver divergente da implementacao;
3. OpenAPI estiver divergente da spec da fase;
4. contrato HTTP publico mudar sem atualizacao de spec;
5. endpoint publico for removido sem aprovacao explicita;
6. schema publico for renomeado sem aprovacao explicita;
7. campo publico obrigatorio for removido sem aprovacao explicita;
8. comportamento documentado por OpenAPI nao tiver teste correspondente quando a fase tocar API.

## 6. Stop Conditions De Spec

Parar imediatamente se:

1. a spec de construcao da fase atual estiver ausente apos a etapa de criacao ou validacao;
2. a spec de diagnostico da fase atual estiver ausente apos a etapa de criacao ou validacao;
3. a spec de construcao nao definir objetivo, escopo, fora de escopo, contratos, testes e criterios de aceite;
4. a spec de diagnostico nao definir sinais observaveis, modos de falha, comandos de verificacao, classificacao e decisao final;
5. houver divergencia material entre spec e codigo;
6. houver comportamento implementado fora da spec;
7. houver capability declarada apenas em documentacao narrativa;
8. houver conflito entre specs sem decisao explicita;
9. houver necessidade de decisao humana de arquitetura;
10. a fase depender de spec futura ainda nao executada.

## 7. Stop Conditions De Contrato Publico

Parar imediatamente se:

1. houver breaking change sem aprovacao explicita;
2. API publica for redefinida silenciosamente;
3. naming publico mudar sem justificativa explicita;
4. rename afetar contrato publico sem spec correspondente;
5. tipo de `internal/*` vazar para assinatura exportada;
6. pacote `pkg/*` depender de `internal/*`, exceto pela ponte publica permitida em `pkg/app`;
7. erro publico observavel mudar sem documentacao e teste;
8. comportamento de cancelamento mudar sem spec e teste;
9. ordem de execucao observavel mudar sem spec e teste;
10. compatibilidade conceitual com {{UPSTREAM_NAME}} divergir sem justificativa registrada.

## 8. Stop Conditions De Roadmap E Estado

Parar imediatamente se:

1. `automation/ROADMAP.json` estiver ausente;
2. `automation/PHASE_STATE.json` estiver ausente;
3. `automation/PHASE_STATE.json` nao for JSON valido;
4. `automation/ROADMAP.json` nao for JSON valido;
5. `current_phase` nao existir em `automation/ROADMAP.json`;
6. `last_completed_phase` for maior ou igual a `current_phase`;
7. `blocked = true` e o autopilot tentar avancar sem desbloqueio;
8. roadmap estiver inconsistente com `PHASE_STATE.json`;
9. fase for pulada;
10. fase for executada fora de ordem;
11. fase nova for adicionada sem pedido explicito;
12. fase existente for removida;
13. gate esperado for alterado;
14. report esperado for alterado.

## 9. Stop Conditions De Escopo

Parar imediatamente se:

1. a correcao exigir abrir capability fora da fase atual;
2. a correcao exigir implementar fase futura;
3. a correcao exigir reordenar fases;
4. a correcao exigir simplificar o roadmap;
5. a correcao exigir remover diagnostico;
6. a correcao exigir mascarar gap com texto narrativo;
7. a correcao exigir aceitar evidencia insuficiente;
8. a correcao exigir depender de hosted service fora do escopo;
9. a correcao exigir {{UPSTREAM_OPS_NAME}};
10. a correcao exigir decisao humana de arquitetura.

## 10. Stop Conditions Do Modo Interativo

Parar imediatamente no `interactive_sdd_autopilot` se:

1. a solicitacao inicial estiver materialmente ambigua e ainda depender de resposta do usuario;
2. a trilha exigir comportamento novo sem spec de construcao;
3. a trilha exigir comportamento novo sem spec de diagnostico;
4. a decisao entre amendar spec existente e criar nova trilha dual-spec estiver ausente;
5. `automation/INTERACTIVE_STATE.json` estiver ausente ou invalido;
6. `current_request.current_step` estiver inconsistente com a execucao;
7. o report da etapa atual estiver ausente;
8. o report nao trouxer classificacao explicita;
9. o report nao trouxer decisao final explicita;
10. o report trouxer `classification = PARTIAL`, `classification = FAIL` ou `classification = BLOCKED`;
11. a decisao final nao autorizar avanco, retry ou conclusao;
12. a implementacao exigiria alterar `automation/ROADMAP.json` ou `automation/PHASE_STATE.json` para contornar gate;
13. a implementacao exigiria capability publica do framework fora de spec;
14. a implementacao exigiria {{UPSTREAM_OPS_NAME}}, hosted service, dashboard, control plane ou deploy tooling;
15. a trilha tentar concluir sem evidencia versionada;
16. houver decisao humana de arquitetura pendente.

Ao parar no modo interativo, registrar o bloqueio em `automation/INTERACTIVE_STATE.json`.

## 11. Relatorio Obrigatorio Ao Parar

Quando parar por qualquer stop condition, o autopilot deve produzir:

1. resumo do bloqueio;
2. fase afetada;
3. nome da fase;
4. spec de construcao;
5. spec de diagnostico;
6. report esperado;
7. report encontrado, se houver;
8. classificacao encontrada, se houver;
9. decisao encontrada, se houver;
10. retries ja consumidos;
11. stop condition acionada;
12. evidencia observada;
13. acao humana recomendada.
</file>

<file path="doc.go">
// Package {{PROJECT_SLUG}} is the root documentation package for {{PROJECT_TITLE}}.
package {{PROJECT_SLUG}}
</file>

<file path="docs/ai/ai-contribution-contract.md">
# AI Contribution Contract

Este contrato define a baseline operacional expandida para contribuicoes assistidas por AI na `{{PROJECT_SLUG}}`.
Ele complementa `AGENTS.md` e materializa a `Spec 350`, secoes 5.1, 5.2, 5.5 e 7.1-7.5, preservando a baseline recuperada pela `Spec 348`.

## Regras normativas

1. Antes de editar codigo, ler `AGENTS.md`, a spec governante, a spec de diagnostico correspondente quando existir, `skills/00-skill-index/SKILL.md` e as skills aplicaveis.
2. Toda tarefa deve chegar com `objetivo`, `specs lidas` com a spec governante identificada, `arquivos em escopo`, `entregaveis minimos`, `validacoes obrigatorias`, `regras` e `formato de saida`.
3. Nenhuma feature pode ser implementada sem spec. Se a spec estiver ausente, incompleta ou ambigua, a execucao deve parar e o gap precisa ser declarado.
4. Mudanca de comportamento exige teste no mesmo patch, ou excecao formal explicitamente registrada em `docs/ai/compliance-exceptions.md`.
5. Mudanca de API publica exige atualizacao da spec correspondente, dos testes de contrato afetados e declaracao do impacto em docs ou superficie publica.
6. `pkg/*` deve continuar livre de dependencias em `internal/*`, exceto pela ponte publica permitida em `pkg/app`.
7. Docs, prompts, briefs, scripts e CI nao podem afirmar capacidade que o repositorio nao entrega de fato.
8. Toda entrega deve declarar, no minimo: `objetivo`, `specs lidas`, `skills aplicadas`, `arquivos alterados`, `comandos executados`, `testes executados` e `gaps restantes`.
9. Quando houver impacto em docs operacionais ou superficie publica, esse impacto deve ser declarado de forma explicita na tarefa, no PR e na documentacao alterada.
10. O checker nao pode depender de excecao silenciosa. Toda excecao ativa de compliance, teste ou enforcement deve ser auditavel no artefato formal do repositorio.

## Uso esperado

- Use `docs/ai/task-input-format.md` como formato padrao de entrada.
- Use `docs/ai/prompts/` como base de prompt operacional, nao como substituto da leitura das specs.
- Use `docs/ai/briefs/` para briefs reais e rastreaveis por spec.
- Use `docs/ai/compliance-exceptions.md` quando alguma exigencia obrigatoria precisar de excecao formal temporaria.

## Nao aceitavel

- implementar feature fora de spec
- deixar marcador vazio em assets operacionais
- tratar gap conhecido como capacidade implementada
- omitir spec governante, testes executados ou gaps restantes quando a mudanca tocar comportamento observavel
- usar excecao informal em comentario, script ou narrativa de PR sem registro auditavel
</file>

<file path="docs/ai/briefs/ai-diagnosis-brief.md">
# Brief real - AI diagnosis

## Objetivo

Diagnosticar o estado real de uma trilha, patch ou artefato operacional contra a spec governante, produzindo classificacao auditavel e sem drift aspiracional.

## Specs governantes

- `AGENTS.md`
- spec de diagnostico da trilha
- spec de construcao correspondente
- `docs/ai/ai-contribution-contract.md`

## Escopo desta entrega

- classificar cada dimensao relevante como PASS, PARTIAL, FAIL ou BLOCKED
- sustentar conclusoes com evidencia por arquivo e, quando necessario, por comando
- separar gap historico, gap remanescente e regressao atual
- registrar justificativa arquitetural quando algum item nao se aplicar diretamente

## Fora do escopo

- reescrever a spec durante o diagnostico
- tratar ausencia de artefato obrigatorio como detalhe menor
- inferir maturidade nao sustentada por script, teste ou artefato versionado

## Evidencias esperadas

- resumo executivo
- avaliacao por dimensao
- evidencias por arquivo
- gaps remanescentes
- decisao final rastreavel
</file>

<file path="docs/ai/briefs/ai-implementation-brief.md">
# Brief real - AI implementation

## Objetivo

Executar uma mudanca implementativa assistida por AI com rastreabilidade por spec, validacao minima real e declaracao explicita de gaps restantes.

## Specs governantes

- `AGENTS.md`
- spec governante da tarefa
- spec de diagnostico correspondente, quando existir
- `docs/ai/ai-contribution-contract.md`
- `docs/ai/task-input-format.md`

## Escopo desta entrega

- ler primeiro a spec governante e os arquivos diretamente afetados
- implementar a menor mudanca correta dentro do escopo pedido
- atualizar testes no mesmo patch quando houver mudanca de comportamento
- sincronizar docs operacionais, prompts, briefs, checker, Makefile ou CI quando o fluxo operacional for parte da mudanca

## Fora do escopo

- introduzir feature fora de spec
- reabrir arquitetura sem necessidade observavel
- declarar capacidade operacional sem artefato correspondente

## Evidencias esperadas

- arquivos alterados listados de forma objetiva
- comandos e testes executados declarados de forma rastreavel
- impacto em docs ou superficie publica declarado quando existir
- gaps restantes registrados de forma explicita
</file>

<file path="docs/ai/briefs/ai-remediation-bugfix-brief.md">
# Brief real - AI remediation ou bugfix

## Objetivo

Corrigir um bug ou gap operacional ja diagnosticado, restaurando aderencia a spec com a menor mudanca correta e com evidencia real de fechamento.

## Specs governantes

- `AGENTS.md`
- spec governante da correcao
- spec de diagnostico que evidenciou o problema, quando existir
- `docs/ai/ai-contribution-contract.md`
- `docs/ai/compliance-exceptions.md`, quando houver excecao ativa relacionada

## Escopo desta entrega

- reproduzir ou citar a evidencia do problema antes da correcao
- limitar a mudanca ao bug ou gap diagnosticado
- atualizar testes, checker, docs ou CI quando eles fizerem parte da causa ou da confirmacao do fechamento
- declarar o que foi resolvido e o que continua pendente

## Fora do escopo

- usar remediation para abrir feature nova
- remover exigencia de compliance sem justificativa arquitetural
- esconder pendencia atras de excecao nao auditavel

## Evidencias esperadas

- reproducao ou referencia concreta do problema
- arquivos alterados e comandos executados
- testes que confirmam a correcao
- gaps restantes ou excecoes formais ainda ativas
</file>

<file path="docs/ai/briefs/spec-348-baseline-recovery-brief.md">
# Brief real - Spec 348 baseline recovery

## Objetivo

Recuperar a baseline minima de maturidade operacional da trilha de AI development da `{{PROJECT_SLUG}}`, trocando marcadores vazios e gaps abertos por mecanismos versionados, executaveis e auditaveis.

## Specs governantes

- `specs/348-ai-operational-maturity-baseline-recovery.md`
- `specs/349-ai-operational-maturity-baseline-recovery-diagnosis.md`
- `AGENTS.md`
- `specs/020-repository-architecture.md`
- `specs/060-guardrails.md`

## Escopo desta entrega

- criar assets operacionais minimos em `docs/ai/`
- versionar regra do Cursor, PR template e config de `pre-commit`
- adicionar checker simples e auditavel
- substituir targets vazios no `Makefile`
- alinhar `.github/workflows/ci.yml` aos targets reais
- recuperar testes basicos em `pkg/types` e `pkg/guardrail`
- atualizar apenas os docs centrais estritamente necessarios

## Fora do escopo

- suite completa de safety regression
- policy engine avancado
- expansao ampla de prompts e briefs
- reabertura de toda a trilha historica 300-347

## Evidencias esperadas

- `make lint`, `make vet` e `make check-compliance` executaveis
- `.pre-commit-config.yaml` usando hooks locais simples
- prompts e briefs reutilizaveis, sem texto de enchimento
- `_test.go` presentes em `pkg/types` e `pkg/guardrail`
- docs centrais alinhados ao novo estado real do repo
</file>

<file path="docs/ai/compliance-exceptions.md">
# AI Compliance Exceptions

Este arquivo registra excecoes formais, pequenas e auditaveis para compliance, testes e enforcement da trilha de AI development.
Ele existe para materializar a `Spec 350`, secao 5.5, sem criar bypass silencioso no checker.

## Estado atual

Ha uma excecao ativa para PRs Dependabot de atualizacao de dependencias.

## Formato de registro

Cada excecao ativa deve listar:

- `id`: identificador curto e estavel
- `escopo`: arquivo, check, teste ou fluxo afetado
- `justificativa`: motivo tecnico observavel
- `owner`: responsavel logico pela revisao
- `criterio de revisao`: prazo, trigger ou condicao objetiva de encerramento

## Excecoes ativas

- `id`: `dependabot-pr-body-template`
  - `escopo`: `scripts/check-pr-body.sh` em eventos `pull_request` gerados pelo Dependabot.
  - `justificativa`: Dependabot gera corpo de PR proprio com changelog e comandos operacionais, sem usar `.github/PULL_REQUEST_TEMPLATE.md`. Exigir headings humanos bloqueia a automacao leve de atualizacao de dependencias aprovada pela `Spec 726`, sem aumentar seguranca ou rastreabilidade do PR automatizado.
  - `owner`: mantenedores do repositorio `{{PROJECT_SLUG}}`.
  - `criterio de revisao`: revisar quando a politica de PR template mudar, quando Dependabot permitir template customizado compativel ou quando a automacao de dependencia for substituida.
</file>

<file path="docs/ai/ops-guide.md">
# Guia de operacao AI - {{PROJECT_SLUG}}

Base normativa atual:

- `AGENTS.md`
- `specs/010-feature-matrix.md`
- `specs/348-ai-operational-maturity-baseline-recovery.md`
- `specs/349-ai-operational-maturity-baseline-recovery-diagnosis.md`
- `specs/350-ai-operational-maturity-expansion.md`
- `specs/351-ai-operational-maturity-expansion-diagnosis.md`
- `specs/352-ai-historical-spec-convergence.md`
- `specs/353-ai-historical-spec-convergence-diagnosis.md`
- `specs/342-ai-development-remediation-closure.md`
- `specs/343-ai-development-remediation-closure-diagnosis.md`
- `specs/344-ai-development-doc-coherence-cleanup.md`
- `specs/345-ai-development-doc-coherence-cleanup-diagnosis.md`
- `specs/346-ai-development-doc-truthfulness-closure.md`
- `specs/347-ai-development-doc-truthfulness-closure-diagnosis.md`
- `docs/ai/spec-lineage.md`

Este guia descreve apenas os artefatos de AI development realmente versionados no repositorio em 2026-04-09.

## 1. Iniciar tarefa com Codex ou Cursor

1. Ler `AGENTS.md` e a spec governante da mudanca.
2. Ler `specs/010-feature-matrix.md` quando a tarefa tocar status, cobertura ou paridade.
3. Ler `skills/00-skill-index/SKILL.md` e as skills aplicaveis ao escopo.
4. Usar `docs/ai/task-input-format.md` como formato padrao de entrada.
5. Reutilizar os prompts oficiais de `docs/ai/prompts/` e os briefs versionados em `docs/ai/briefs/` quando fizer sentido.
6. Exigir que a resposta do agente liste objetivo, specs lidas, skills aplicadas, arquivos alterados, comandos executados, testes executados e gaps restantes.

Estado atual de assets operacionais:

- existe `docs/ai/ai-contribution-contract.md`
- existe `docs/ai/task-input-format.md`
- existe `docs/ai/compliance-exceptions.md` com caminho formal para excecoes auditaveis
- existe `docs/ai/prompts/` com prompts reais para implementacao, review, diagnostico e remediation/bugfix
- existe `docs/ai/briefs/` com briefs reais para implementacao, diagnostico e remediation/bugfix
- existe `.cursor/rules/{{PROJECT_SLUG}}.mdc`
- existe `test/security/` com baseline minima real baseada em guardrails publicos
- existe `make test-security` e o CI o executa explicitamente

Esta baseline continua enxuta e repo-native. Ela nao substitui specs, mas evita remontar o fluxo operacional do zero a cada tarefa e materializa enforcement semantico leve.

## 2. Validar a entrega

Campos minimos esperados na resposta do agente:

- [ ] objetivo
- [ ] specs lidas
- [ ] skills aplicadas
- [ ] arquivos alterados
- [ ] comandos executados
- [ ] testes executados
- [ ] gaps restantes

Comandos realmente disponiveis hoje:

- `make fmt-check`
- `make lint`
- `make vet`
- `make check-compliance`
- `make test-security`
- `make test`
- `make race`
- `make coverage`
- `pre-commit run --all-files`

Gaps operacionais atuais:

- a baseline de AI continua minima; nao existe suite ampla de safety regression
- specs `310/311` e `330/331` foram formalmente superadas por `348/349` e `350/351` respectivamente (decisao sob Spec 352; ver `docs/ai/spec-lineage.md`)

## 3. Estado documental da trilha AI

| Item | Estado atual | Artefatos reais |
| --- | --- | --- |
| Enforcement baseline | baseline minima recuperada e expandida com checker semantico leve, caminho formal de excecoes, `test-security` e CI real; specs `310/311` formalmente superadas por `348/349` (Spec 352) | `specs/348-ai-operational-maturity-baseline-recovery.md`, `specs/349-ai-operational-maturity-baseline-recovery-diagnosis.md`, `specs/350-ai-operational-maturity-expansion.md`, `specs/351-ai-operational-maturity-expansion-diagnosis.md`, `docs/reports/ai-enforcement-baseline-report.md`, `scripts/check-compliance.sh`, `docs/ai/spec-lineage.md` |
| Remediation closure | trilha historica `342/343` continua nao consolidada; gap normativo de `310/311` resolvido pela convergencia (Spec 352), mas D5 (safety) permanece FAIL | `specs/342-ai-development-remediation-closure.md`, `specs/343-ai-development-remediation-closure-diagnosis.md`, `docs/reports/ai-remediation-closure-report.md` |
| Doc coherence cleanup | specs existem | `specs/344-ai-development-doc-coherence-cleanup.md`, `specs/345-ai-development-doc-coherence-cleanup-diagnosis.md` |
| Truthfulness closure | specs existem; esta e a trilha usada para corrigir os docs centrais | `specs/346-ai-development-doc-truthfulness-closure.md`, `specs/347-ai-development-doc-truthfulness-closure-diagnosis.md` |
| Governance baseline | baseline expandida agora versionada em docs, prompts, briefs, excecoes formais e rule set do Cursor | `docs/ai/ai-contribution-contract.md`, `docs/ai/task-input-format.md`, `docs/ai/compliance-exceptions.md`, `docs/ai/prompts/`, `docs/ai/briefs/`, `.cursor/rules/{{PROJECT_SLUG}}.mdc` |
| Operabilidade Codex/Cursor | baseline operacional minima expandida disponivel | `docs/ai/task-input-format.md`, `docs/ai/prompts/`, `docs/ai/briefs/`, `.cursor/rules/{{PROJECT_SLUG}}.mdc` |
| Safety regression baseline | baseline minima real presente; specs `330/331` formalmente superadas por `350/351` (Spec 352); suite dedicada ampla continua ausente | `specs/350-ai-operational-maturity-expansion.md`, `specs/351-ai-operational-maturity-expansion-diagnosis.md`, `test/security/`, `docs/ai/spec-lineage.md` |
| Learning loop | gap documental | specs `340/341` e relatorio dedicado nao encontrados |

## 4. Onde encontrar o que

| Artefato real | Caminho |
| --- | --- |
| Workflow normativo | `AGENTS.md` |
| Indice de skills | `skills/00-skill-index/SKILL.md` |
| Feature matrix | `specs/010-feature-matrix.md` |
| Specs do repositorio | `specs/` |
| Contrato de contribuicao AI | `docs/ai/ai-contribution-contract.md` |
| Formato padrao de entrada | `docs/ai/task-input-format.md` |
| Excecoes formais de compliance | `docs/ai/compliance-exceptions.md` |
| Prompts operacionais | `docs/ai/prompts/` |
| Briefs versionados | `docs/ai/briefs/` |
| Guia operacional AI atual | `docs/ai/ops-guide.md` |
| Relatorio de enforcement AI | `docs/reports/ai-enforcement-baseline-report.md` |
| Relatorio de remediation closure AI | `docs/reports/ai-remediation-closure-report.md` |
| Linhagem normativa da trilha AI | `docs/ai/spec-lineage.md` |
| Workflow de CI existente | `.github/workflows/ci.yml` |
| Comandos versionados | `Makefile` |
| Hooks locais | `.pre-commit-config.yaml` |
| Checker de compliance | `scripts/check-compliance.sh` |
| Baseline minima de safety | `test/security/` |
| Fluxo de contribuicao | `CONTRIBUTING.md` |

## 5. Gaps explicitos desta trilha

Os itens abaixo ja foram citados historicamente em docs da trilha AI, mas nao existem no estado atual do repositorio:

- specs `300`, `301`, `320`, `321`, `340` e `341`
- `docs/ai/sensitive-changes-checklist.md`
- `docs/ai/threat-model.md`
- `docs/ai/gap-tracker.md`
- `docs/ai/anti-patterns.md`
- `docs/ai/postmortem-template.md`

Esses itens devem ser tratados como gaps documentais ou operacionais, nunca como capacidade implementada.

Specs formalmente superadas (nao sao mais gaps; ver `docs/ai/spec-lineage.md`):

- `310/311` — superadas por `348/349` (decisao sob Spec 352)
- `330/331` — superadas por `350/351` (decisao sob Spec 352)
</file>

<file path="docs/ai/prompts/audit-or-diagnose.md">
# Prompt Operacional - Auditoria ou diagnostico

Use este prompt quando a tarefa principal for auditar implementacao, docs, enforcement ou aderencia a spec.

```text
Leia primeiro:
- AGENTS.md
- a spec de diagnostico governante
- a spec de construcao correspondente
- os docs, scripts, testes e workflows citados pela spec

Objetivo:
Auditar o estado real do repositorio contra a spec, sem superestimar capacidade implementada.

Regras de execucao:
- classifique cada dimensao exigida pela spec como PASS, PARTIAL, FAIL ou BLOCKED
- cite evidencias concretas por arquivo e, quando necessario, por comando executado
- diferencie gap historico de gap introduzido pela mudanca atual
- nao trate ausencia de artefato como detalhe menor quando ela for criterio de aceite
- se a implementacao divergir da documentacao, relate a divergencia explicitamente

Formato de saida:
1. resumo executivo
2. avaliacao por dimensao
3. evidencias por arquivo
4. gaps remanescentes
5. decisao final
```
</file>

<file path="docs/ai/prompts/diagnose-task.md">
# Prompt Operacional - Diagnostico

Use este prompt quando a tarefa principal for diagnosticar estado real, auditar aderencia a spec ou produzir um relatorio de confirmacao.

```text
Leia primeiro:
- AGENTS.md
- a spec de diagnostico governante
- a spec de construcao correspondente
- os docs, scripts, testes e workflows citados pelas specs

Objetivo:
Diagnosticar o estado real do repositorio contra a spec, sem superestimar maturidade, cobertura ou enforcement.

Regras de execucao:
- classifique cada dimensao exigida pela spec como PASS, PARTIAL, FAIL ou BLOCKED
- cite evidencias concretas por arquivo e, quando necessario, por comando executado
- diferencie gap historico, gap remanescente e regressao introduzida
- se algum item nao se aplicar diretamente, registre a justificativa arquitetural
- nao trate ausencia de artefato obrigatorio como detalhe menor

Formato de saida:
1. resumo executivo
2. avaliacao por dimensao
3. evidencias por arquivo
4. gaps remanescentes
5. decisao final
```
</file>

<file path="docs/ai/prompts/implement-task.md">
# Prompt Operacional - Implementacao de tarefa

Use este prompt quando a tarefa principal for implementar ou corrigir algo no repositorio.

```text
Leia primeiro:
- AGENTS.md
- a spec governante da mudanca
- a spec de diagnostico correspondente, quando existir
- specs/010-feature-matrix.md, quando a tarefa tocar status, cobertura, risco ou paridade
- skills/00-skill-index/SKILL.md
- as skills aplicaveis ao escopo
- os arquivos diretamente afetados

Objetivo:
Implementar a tarefa pedida com a menor mudanca correta, mantendo rastreabilidade por spec.

Regras de execucao:
- identifique explicitamente qual spec e qual secao governam a mudanca antes de editar
- nao implemente feature fora de spec; se faltar especificacao, pare e declare o gap
- mantenha fronteiras entre `pkg/*` e `internal/*`
- atualize testes no mesmo patch quando o comportamento mudar
- se algum teste ou check obrigatorio nao se aplicar, registre excecao formal em `docs/ai/compliance-exceptions.md`
- declare impacto em docs ou superficie publica quando houver
- atualize docs operacionais somente quando houver artefato real sustentando a afirmacao
- nao deixe marcadores vazios em Makefile, CI, prompts, briefs, templates ou scripts

Validacao minima:
- execute os comandos mais relevantes para o escopo alterado
- declare claramente o que foi executado, o que falhou e o que ficou pendente

Formato de saida:
- objetivo
- specs lidas
- skills aplicadas
- arquivos alterados
- comandos executados
- testes executados
- gaps restantes
```
</file>

<file path="docs/ai/prompts/remediation-bugfix.md">
# Prompt Operacional - Remediation ou bugfix

Use este prompt quando a tarefa principal for corrigir um bug, fechar um gap diagnosticado ou remediar um desvio de compliance sem reabrir escopo amplo.

```text
Leia primeiro:
- AGENTS.md
- a spec governante da correcao
- a spec de diagnostico que evidenciou o gap, quando existir
- os arquivos afetados pelo bug ou desvio
- skills/00-skill-index/SKILL.md
- as skills aplicaveis ao escopo

Objetivo:
Corrigir o problema com a menor mudanca correta, restaurando aderencia a spec e deixando evidencia auditavel da remediacao.

Regras de execucao:
- reproduza o problema ou descreva a evidencia concreta antes de editar
- mantenha a correcao dentro do escopo diagnosticado; nao use remediation para introduzir feature nova
- atualize testes, checker, docs ou CI no mesmo patch quando eles fizerem parte da causa ou da evidencia da correcao
- se algum requisito obrigatorio continuar pendente, registre o gap remanescente e a justificativa
- use `docs/ai/compliance-exceptions.md` apenas quando uma excecao formal continuar necessaria apos a remediacao

Validacao minima:
- execute os checks que reproduzem ou confirmam a correcao
- declare claramente o que ficou resolvido e o que ainda permanece aberto

Formato de saida:
- objetivo
- specs lidas
- skills aplicadas
- arquivos alterados
- comandos executados
- testes executados
- gaps restantes
```
</file>

<file path="docs/ai/prompts/review-change.md">
# Prompt Operacional - Review de mudanca

Use este prompt quando a tarefa principal for revisar codigo, docs, testes, CI ou enforcement antes de merge.

```text
Leia primeiro:
- AGENTS.md
- a spec governante da mudanca revisada
- a spec de diagnostico correspondente, quando existir
- os arquivos alterados e os testes relacionados
- skills/00-skill-index/SKILL.md
- as skills aplicaveis ao escopo revisado

Objetivo:
Revisar a mudanca contra a spec, priorizando bugs, regressao comportamental, riscos operacionais, gaps de teste e drift documental.

Regras de execucao:
- apresente findings primeiro, ordenados por severidade
- cite arquivos e comportamento observavel, nao opinioes vagas
- trate ausencia de spec, teste ou doc obrigatorio como finding real quando afetar criterio de aceite
- diferencie bug real, risco potencial e gap residual
- se nao houver findings, diga isso explicitamente e registre riscos ou coberturas ainda limitadas

Formato de saida:
1. findings
2. perguntas ou premissas abertas
3. resumo curto da mudanca
```
</file>

<file path="docs/ai/spec-lineage.md">
# AI Development Spec Lineage

Spec governante: `specs/352-ai-historical-spec-convergence.md`
Data da decisao: 2026-04-09

---

## 1. Decisao de convergencia

**Estrategia adotada: substituicao formal (Estrategia A da Spec 352, secao 5).**

As specs historicamente planejadas `310/311` e `330/331` nunca foram criadas como arquivos no repositorio e nunca governaram nenhuma implementacao real. A funcionalidade prevista para essas trilhas foi absorvida por specs posteriores que de fato orientaram a implementacao.

A partir desta decisao:

- `310/311` sao formalmente superadas por `348/349`
- `330/331` sao formalmente superadas por `350/351`
- Nenhum arquivo de backfill retroativo sera criado para 310, 311, 330 ou 331
- A rastreabilidade entre trilha planejada e trilha real e garantida por este documento

Esta decisao nao inventa historico de execucao. As specs 310/311 e 330/331 nunca existiram como arquivos e nunca governaram trabalho passado.

---

## 2. Tabela de rastreabilidade

| Spec planejada (ausente) | Funcionalidade prevista | Superada por (existente) | Evidencia de absorcao |
| --- | --- | --- | --- |
| 310 - AI development enforcement baseline | checker de compliance, PR template, linter real, pre-commit, targets reais de CI | `specs/348-ai-operational-maturity-baseline-recovery.md` (gaps G1-G4) | `scripts/check-compliance.sh`, `Makefile`, `.github/workflows/ci.yml`, `.github/PULL_REQUEST_TEMPLATE.md`, `.pre-commit-config.yaml` |
| 311 - AI development enforcement diagnosis | auditoria da baseline de enforcement | `specs/349-ai-operational-maturity-baseline-recovery-diagnosis.md` (dimensoes D1-D6) | `docs/reports/ai-enforcement-baseline-report.md` |
| 330 - AI safety regression baseline | baseline minima de testes de seguranca, input malicioso, output validation, fail-closed | `specs/350-ai-operational-maturity-expansion.md` (gap G4, secao 5.4) | `test/security/guardrail_security_test.go` |
| 331 - AI safety regression diagnosis | auditoria da baseline de safety | `specs/351-ai-operational-maturity-expansion-diagnosis.md` (dimensao D4) | `docs/reports/ai-expansion-diagnosis-report.md` |

---

## 3. Classificacao das specs da trilha AI (300+)

### Specs governantes (existem e governam implementacao real)

| Spec | Nome | Par diagnostico | Status |
| --- | --- | --- | --- |
| 342 | AI development remediation closure | 343 | existente; D1 e D5 historicamente FAIL; gap normativo de D1 resolvido por esta convergencia |
| 344 | AI development doc coherence cleanup | 345 | existente |
| 346 | AI development doc truthfulness closure | 347 | existente |
| 348 | AI operational maturity baseline recovery | 349 | existente; superou 310/311 |
| 350 | AI operational maturity expansion | 351 | existente; superou 330/331 |
| 352 | AI historical spec convergence | 353 | existente; esta decisao |

### Specs formalmente superadas (nunca existiram como arquivos)

| Spec planejada | Superada por | Justificativa |
| --- | --- | --- |
| 310 | 348 | funcionalidade absorvida; nenhuma implementacao foi guiada por 310 |
| 311 | 349 | funcionalidade absorvida; nenhuma auditoria foi guiada por 311 |
| 330 | 350 | funcionalidade absorvida; nenhuma implementacao foi guiada por 330 |
| 331 | 351 | funcionalidade absorvida; nenhuma auditoria foi guiada por 331 |

### Specs historicas ausentes (sem substituicao formal nesta convergencia)

| Spec planejada | Funcionalidade prevista | Estado |
| --- | --- | --- |
| 300/301 | (referenciadas historicamente sem descricao precisa) | ausentes; sem substituicao formal; tratadas como gap documental |
| 320/321 | (referenciadas historicamente sem descricao precisa) | ausentes; sem substituicao formal; tratadas como gap documental |
| 340/341 | learning loop | ausentes; sem substituicao formal; tratadas como gap documental |

---

## 4. Impacto na Spec 342

A `Spec 342` referencia `specs/310` e `specs/311` como dependencias em sua secao 5.1 (Gap G1) e no prompt de execucao (secao 10). A `Spec 342` tambem referencia `specs/330` e `specs/331` no prompt de execucao.

A partir desta convergencia:

- Toda referencia a 310/311 na Spec 342 deve ser lida como apontando para 348/349
- Toda referencia a 330/331 na Spec 342 deve ser lida como apontando para 350/351
- O FAIL em D1 do relatorio de remediation closure (`docs/reports/ai-remediation-closure-report.md`) refletia o gap normativo de 310/311; esse gap normativo esta agora resolvido pela substituicao formal
- D5 do relatorio de remediation closure (safety gap) permanece inalterado por esta convergencia

---

## 5. Guia para contribuidores futuros

### O que ler para entender a trilha AI

1. `AGENTS.md` — workflow normativo e regras globais
2. `docs/ai/ops-guide.md` — guia operacional com estado atual de assets
3. Este documento (`docs/ai/spec-lineage.md`) — linhagem e rastreabilidade historica
4. As specs governantes listadas na secao 3 deste documento

### Como interpretar specs da trilha AI

- Se a spec existe como arquivo em `specs/`, ela e governante ou historica conforme a secao 3 deste documento
- Se a spec nao existe como arquivo e esta listada como "formalmente superada", sua funcionalidade foi absorvida pela spec indicada na tabela de rastreabilidade
- Se a spec nao existe como arquivo e esta listada como "historica ausente", ela e um gap documental reconhecido e nao governou nenhuma implementacao
- Nunca assumir que uma spec ausente governou implementacao real

### Onde esta o enforcement real

A baseline de enforcement da trilha AI esta materializada em:

- `scripts/check-compliance.sh` — checker semantico
- `.pre-commit-config.yaml` — hooks locais
- `.github/workflows/ci.yml` — pipeline de CI
- `.github/PULL_REQUEST_TEMPLATE.md` — template de PR
- `docs/ai/ai-contribution-contract.md` — contrato de contribuicao
- `docs/ai/compliance-exceptions.md` — caminho formal de excecoes
- `test/security/` — baseline minima de safety

Esses artefatos foram criados sob as specs 348 e 350, nao sob 310 ou 330.
</file>

<file path="docs/ai/task-input-format.md">
# Task Input Format

Use este formato quando abrir uma tarefa para Codex, Cursor ou outro agente.
Ele complementa `AGENTS.md`, reduz drift entre prompt, spec e resultado entregue, e materializa a `Spec 350`, secoes 5.1 e 5.2.

## Campos obrigatorios

- `Objetivo`: resultado concreto esperado.
- `Specs lidas`: lista das specs e docs obrigatorios de leitura, com a spec governante identificada explicitamente.
- `Arquivos em escopo`: arquivos ou diretorios que o agente deve inspecionar primeiro.
- `Entregaveis minimos`: artefatos que precisam existir ao final.
- `Impacto esperado em docs ou superficie publica`: declare `none` quando nao houver impacto relevante.
- `Validacoes obrigatorias`: comandos, testes ou checks a executar.
- `Regras`: restricoes de escopo, arquitetura ou operacao.
- `Formato de saida`: campos que a resposta final deve listar.

## Template

```md
Leia primeiro:
- AGENTS.md
- specs/<spec-governante>.md
- specs/<spec-diagnostico-correspondente>.md # quando a trilha tiver spec de diagnostico
- specs/010-feature-matrix.md # quando tocar status, cobertura ou paridade
- arquivos e docs adicionais realmente necessarios

Objetivo:
Descreva o resultado observavel esperado.

Specs lidas:
- AGENTS.md
- specs/<spec-governante>.md
- specs/<spec-diagnostico-correspondente>.md # quando existir
- outros docs realmente necessarios

Arquivos em escopo:
- caminhos que precisam ser lidos antes da implementacao

Entregaveis minimos:
1. artefato ou mudanca 1
2. artefato ou mudanca 2

Impacto esperado em docs ou superficie publica:
- none

Validacoes obrigatorias:
- comando ou teste 1
- comando ou teste 2

Regras:
- limite de escopo
- restricao arquitetural
- criterio de auditabilidade
- se alguma validacao obrigatoria nao se aplicar, registrar excecao formal em `docs/ai/compliance-exceptions.md`

Formato de saida:
- objetivo
- specs lidas
- skills aplicadas
- arquivos alterados
- comandos executados
- testes executados
- gaps restantes
```

## Quando detalhar mais

- Se houver mudanca de API publica, cite a spec do modulo, os testes de contrato afetados e o impacto em docs ou superficie publica.
- Se houver excecao formal em vez de teste ou check, declare a justificativa no input e vincule `docs/ai/compliance-exceptions.md`.
- Se o escopo tocar CI, docs operacionais, prompts ou briefs, liste os artefatos reais que devem ser sincronizados no mesmo patch.
</file>

<file path="docs/feature-request-lifecycle.md">
# Feature Request Lifecycle

## Objetivo

Este documento descreve o ciclo humano de uma solicitacao interativa no `{{PROJECT_SLUG}}`.

Ele complementa `docs/interactive-sdd-autopilot.md`, `automation/INTERACTIVE_AUTOPILOT.md`, `automation/INTERACTIVE_RUNBOOK.md`, `automation/STOP_CONDITIONS.md` e `skills/23-{{PROJECT_SLUG}}-sdd-autopilot/SKILL.md`.

O ciclo abaixo descreve como uma pessoa deve iniciar, acompanhar, revisar, bloquear, retomar e aceitar uma trilha de `interactive_sdd_autopilot`.

## 1. Pedido Inicial

O usuario descreve uma feature, evolucao, bug, refatoracao, investigacao ou mudanca documental.

O pedido inicial nao e uma spec aprovada. Ele e a entrada para intake.

Pedido recomendado:

```text
Execute o interactive_sdd_autopilot do {{PROJECT_SLUG}} para esta solicitacao:
<descreva a mudanca observavel>

Escopo conhecido:
- <arquivos, modulo ou comportamento esperado>

Fora do escopo:
- <limites que nao devem ser cruzados>
```

Se o usuario nao informar escopo ou criterios suficientes, o Cursor deve normalizar o que existe e perguntar pelo menor conjunto util de informacoes.

## 2. Intake

Durante o intake, o Cursor deve registrar:

1. texto original;
2. tipo do pedido;
3. objetivo;
4. escopo conhecido;
5. fora do escopo quando identificavel;
6. incertezas materiais;
7. specs candidatas;
8. skills candidatas;
9. decisao inicial.

Exemplo de intake versionado: `docs/reports/phase-26-interactive-sdd-autopilot-smoke-evidence.md`.

No smoke da Fase 26, o pedido foi classificado como `docs`, com objetivo de consolidar exemplos humanos de uso do modo interativo. O escopo ficou restrito a docs, reports e `automation/INTERACTIVE_STATE.json`.

## 3. Refinamento

O Cursor deve perguntar antes de editar quando faltarem informacoes materiais sobre:

1. comportamento observavel;
2. escopo;
3. fora de escopo;
4. contrato publico;
5. impacto arquitetural;
6. estrategia de testes;
7. criterio diagnostico;
8. criterio de aceite;
9. decisao humana de arquitetura.

As perguntas devem ser objetivas e poucas por vez.

Pedido que deve perguntar ou bloquear:

```text
Execute o interactive_sdd_autopilot do {{PROJECT_SLUG}} para esta solicitacao:
Melhore o autopilot para ser mais seguro.
```

Esse pedido nao define objetivo observavel, escopo, diagnostico, testes nem criterio de aceite. O Cursor nao deve implementar baseado nessa frase.

## 4. Decisao De Spec

O Cursor deve decidir entre `amend` e nova trilha dual-spec.

Use `amend` quando a solicitacao pertence a um dominio ja governado e a spec existente consegue receber a regra, criterio ou exemplo sem misturar responsabilidades. O smoke da Fase 26 usou esse caminho porque a mudanca era onboarding humano exigido por `Spec 700` e diagnosticado por `Spec 701`.

Abra nova trilha dual-spec quando a solicitacao cria novo bounded context, novo contrato principal, diagnostico separado ou impacto transversal que merece governanca propria.

Uma nova trilha so pode seguir quando tiver:

1. spec de construcao;
2. spec de diagnostico.

Se a decisao de spec estiver ausente, a trilha deve bloquear antes de implementacao.

## 5. Specs Governantes

Antes de implementar, uma pessoa deve conseguir apontar quais specs governam a mudanca.

Para uma trilha `amend`, confira:

1. specs existentes lidas;
2. secoes governantes;
3. motivo para nao abrir nova dual-spec;
4. lacunas registradas, se houver.

Para uma nova trilha dual-spec, confira:

1. spec de construcao com objetivo, escopo, fora de escopo, comportamento observavel, testes e criterios de aceite;
2. spec de diagnostico com sinais, modos de falha, comandos, evidencias, classificacao, decisao e report esperado;
3. aderencia entre as duas specs.

## 6. Implementacao Governada

A implementacao so comeca quando as specs governantes cobrem a mudanca.

Durante a implementacao, o Cursor deve:

1. editar apenas arquivos no escopo;
2. preservar `pkg/` e `internal/` quando a trilha for somente governanca, docs ou onboarding;
3. nao alterar `automation/ROADMAP.json` nem `automation/PHASE_STATE.json` para representar trilha interativa;
4. adicionar ou atualizar testes quando comportamento mudar;
5. atualizar docs quando comportamento operacional ou publico mudar;
6. registrar gaps que nao puderem ser fechados.

No smoke da Fase 26, a implementacao governada foi documental: docs e reports foram atualizados; nenhuma capability de produto foi criada.

## 7. Diagnostico

O diagnostico deve seguir a spec diagnostica da trilha.

Ele deve registrar:

1. comandos executados;
2. testes executados;
3. evidencias coletadas;
4. skips justificados;
5. gaps restantes;
6. classificacao;
7. decisao final.

Para mudanca documental, o diagnostico pode incluir validacao JSON, `go test ./...`, `go test -race ./...`, lints dos arquivos alterados e revisao de escopo proibido. `make fmt` e OpenAPI podem ser justificados como nao aplicaveis quando nenhum arquivo Go ou HTTP/OpenAPI mudar.

Quando a solicitacao usar `@kodus/agent-readiness`, aplique `skills/24-agent-readiness-governance/SKILL.md`. O report deve ficar em `docs/reports/agent-readiness/`, cada achado deve ser classificado para o contexto do `{{PROJECT_SLUG}}`, e o resultado deve ser tratado como evidencia auxiliar. O score bruto nao autoriza avanco nem bloqueio sem o report diagnostico da trilha.

## 8. Report E Gate

O report e a fonte de verdade para avanco.

Classificacoes permitidas:

1. `PASS`;
2. `PARTIAL`;
3. `FAIL`;
4. `BLOCKED`.

O Cursor so pode avancar quando o report declarar `PASS` e a decisao final autorizar a proxima etapa ou a conclusao.

Exemplo de report de trilha: `docs/reports/phase-26-interactive-sdd-autopilot-smoke-report.md`.

Uma pessoa deve conferir:

1. se o report corresponde ao `current_request.report_file`;
2. se a etapa do report corresponde a `current_request.current_step`;
3. se `classification` e `decision` existem;
4. se a decisao autoriza avanco, retry ou conclusao;
5. se os testes obrigatorios foram executados ou justificados;
6. se nenhuma stop condition permanece ativa.

## 9. Estado Interativo

`automation/INTERACTIVE_STATE.json` e o estado da trilha interativa.

Ele deve registrar:

1. `mode = "interactive_sdd_autopilot"`;
2. `status`;
3. `blocked` e `block_reason`;
4. request id;
5. tipo do pedido;
6. pedido original;
7. etapa atual;
8. specs governantes;
9. spec de construcao;
10. spec de diagnostico;
11. report atual;
12. retry count;
13. classificacao e decisao mais recentes;
14. criterios de conclusao;
15. historico relevante quando houver smoke, bloqueio ou retomada.

Durante revisao humana, `INTERACTIVE_STATE.json` deve bater com o report. Se o estado disser `completed`, o report deve declarar `PASS` e decisao de conclusao. Se o estado disser `blocked`, deve haver motivo objetivo e acao humana recomendada.

## 10. Bloqueio

A trilha deve bloquear quando:

1. faltar spec obrigatoria;
2. faltar resposta humana para requisito material;
3. a decisao entre `amend` e nova trilha dual-spec estiver ausente;
4. teste obrigatorio falhar;
5. diagnostico obrigatorio nao tiver sido executado;
6. report estiver ausente ou ambiguo;
7. classificacao for `PARTIAL`, `FAIL` ou `BLOCKED`;
8. a correcao exigir escopo proibido;
9. a correcao exigiria alterar o roadmap fixo para contornar gate;
10. houver decisao humana de arquitetura pendente.

Ao bloquear, o Cursor deve registrar motivo objetivo e acao humana recomendada em `automation/INTERACTIVE_STATE.json`, no report ou na evidencia da trilha, e na resposta final.

No smoke da Fase 26, o bloqueio controlado aconteceu em `report_review` porque o report da etapa ainda estava ausente ou sem `classification` e `decision`.

## 11. Retomada

A retomada deve partir de artefatos existentes, nao de narrativa da conversa:

1. ler `automation/INTERACTIVE_STATE.json`;
2. ler a spec de construcao;
3. ler a spec de diagnostico;
4. ler report ou evidencia do bloqueio;
5. corrigir somente a causa do bloqueio;
6. preservar historico e retry count;
7. reexecutar diagnostico;
8. gerar ou atualizar report;
9. avancar apenas se `classification = PASS` e `decision` autorizar.

O smoke da Fase 26 demonstra esse fluxo registrando bloqueio por report ausente, retomada apos criacao do report da trilha e conclusao com retry count preservado em `1`.

## 12. Conclusao

Uma solicitacao esta concluida quando:

1. todas as etapas necessarias foram executadas;
2. specs e docs exigidas foram atualizadas;
3. diagnosticos obrigatorios foram executados;
4. report final declara `PASS`;
5. decisao final declara conclusao;
6. nenhuma stop condition esta ativa;
7. `automation/INTERACTIVE_STATE.json` registra estado coerente com a conclusao;
8. gaps restantes estao ausentes ou explicitamente registrados.

No smoke da Fase 26, a conclusao ficou registrada em `automation/INTERACTIVE_STATE.json` com `status = "completed"` e no report de smoke com decisao `SMOKE COMPLETE - READY FOR PHASE 26 AUDIT`.

## 13. Fluxo Completo Bem-Sucedido

Fluxo real usado como referencia:

1. Pedido: adicionar exemplos humanos de intake, ambiguidade, retomada e checklist.
2. Intake: tipo `docs`, escopo documental, sem mudanca em `pkg/*` ou `internal/*`.
3. Decisao: `amend` em `Spec 700` e `Spec 701`.
4. Implementacao: docs, reports e `automation/INTERACTIVE_STATE.json`.
5. Diagnostico: JSON valido, `go test ./...`, `go test -race ./...` e lints.
6. Report: `docs/reports/phase-26-interactive-sdd-autopilot-smoke-report.md`.
7. Gate: `classification = PASS`.
8. Decisao: `SMOKE COMPLETE - READY FOR PHASE 26 AUDIT`.
9. Estado: `status = "completed"`.

## 14. Fluxo Bloqueado E Retomado

Fluxo real usado como referencia:

1. A trilha chegou a `report_review`.
2. O report da etapa ainda estava ausente ou sem `classification` e `decision`.
3. O gate bloqueou a trilha.
4. `automation/INTERACTIVE_STATE.json` preservou `retry_count = 1`, classificacao `BLOCKED`, decisao de stop e acao humana recomendada.
5. A retomada leu estado, specs e evidencia do bloqueio.
6. A causa foi resolvida com a criacao do report de smoke.
7. O diagnostico foi reexecutado.
8. O report passou com `classification = PASS`.
9. A trilha foi concluida sem alterar `ROADMAP.json`, `PHASE_STATE.json`, gate ou criterio de aceite.

## 15. Checklist Humano

Antes de aceitar a conclusao de uma trilha, revise:

1. pedido inicial e intake normalizado;
2. decisao entre `amend` e nova trilha dual-spec;
3. spec de construcao e spec de diagnostico;
4. arquivos alterados;
5. comandos e testes executados;
6. report com `classification` e `decision`;
7. `automation/INTERACTIVE_STATE.json`;
8. ausencia de mudancas indevidas em `automation/ROADMAP.json` e `automation/PHASE_STATE.json`;
9. ausencia de capability de produto fora de spec;
10. bloqueios e retomadas preservados quando existirem;
11. gaps restantes explicitamente registrados.
</file>

<file path="docs/interactive-sdd-autopilot.md">
# Interactive SDD Autopilot

## Objetivo

O `interactive_sdd_autopilot` e o modo de trabalho para transformar uma solicitacao inicial em uma trilha Spec Driven Development completa dentro do `{{PROJECT_SLUG}}`.

Ele existe para orientar o desenvolvimento do repositorio. Ele nao cria API publica, runtime feature, dashboard, runner hospedado, {{UPSTREAM_OPS_NAME}} ou capability de produto do framework.

Use este guia como onboarding humano. As fontes normativas continuam sendo `AGENTS.md`, `specs/680-phase-25-interactive-sdd-autopilot-foundation.md`, `specs/681-phase-25-interactive-sdd-autopilot-foundation-diagnosis.md`, `specs/700-phase-26-interactive-sdd-autopilot-verification-and-human-onboarding.md`, `specs/701-phase-26-interactive-sdd-autopilot-verification-and-human-onboarding-diagnosis.md`, `automation/INTERACTIVE_AUTOPILOT.md`, `automation/INTERACTIVE_RUNBOOK.md`, `automation/STOP_CONDITIONS.md` e `skills/23-{{PROJECT_SLUG}}-sdd-autopilot/SKILL.md`.

## Quando Usar

Use o `interactive_sdd_autopilot` quando uma pessoa trouxer uma solicitacao inicial que precisa virar entrega governada por specs, por exemplo:

1. feature nova;
2. evolucao de comportamento existente;
3. bug com impacto em contrato ou diagnostico;
4. refatoracao que muda fluxo operacional ou risco arquitetural;
5. documentacao operacional ou onboarding de desenvolvimento;
6. investigacao que pode terminar em implementacao, bloqueio ou decisao de spec.

O modo e adequado quando a solicitacao precisa de intake, decisao entre `amend` e nova trilha dual-spec, implementacao governada, diagnostico, report, gate e atualizacao de `automation/INTERACTIVE_STATE.json`.

## Quando Nao Usar

Nao use o modo interativo para executar o roadmap fixo. O `phase_autopilot` continua governado por `automation/ROADMAP.json`, `automation/PHASE_STATE.json`, `automation/AUTOPILOT.md` e `automation/RUNBOOK.md`.

Tambem nao use o modo interativo para:

1. registrar feature request em `automation/ROADMAP.json`;
2. atualizar `automation/PHASE_STATE.json` para representar uma trilha interativa;
3. contornar gate, stop condition, report ausente ou teste obrigatorio falho;
4. criar capability publica sem spec de produto propria;
5. depender de dashboard, hosted service, control plane, runner gerenciado, deploy tooling ou {{UPSTREAM_OPS_NAME}};
6. tratar conversa, plano, prompt, TODO ou ausencia aparente de erro como evidencia de conclusao;
7. implementar quando objetivo, escopo, contrato, diagnostico, teste ou criterio de aceite estiver materialmente ambiguo.

Se a tarefa for apenas uma pergunta de leitura, uma explicacao curta ou uma verificacao sem mudanca de artefato, responda diretamente e nao abra trilha interativa.

## Como Iniciar Uma Solicitacao

Use um pedido explicito e inclua o comportamento esperado, o escopo conhecido e qualquer limite importante:

```text
Execute o interactive_sdd_autopilot do {{PROJECT_SLUG}} para esta solicitacao:
Adicionar uma pagina de documentacao explicando como revisar reports de trilhas interativas.

Escopo conhecido:
- documentacao humana;
- exemplos de revisao;
- nenhum codigo Go.

Fora do escopo:
- alterar gates;
- alterar ROADMAP.json ou PHASE_STATE.json;
- criar capability de produto.
```

Se o pedido for curto, o Cursor deve normalizar o que estiver claro e perguntar quando faltar informacao material. O pedido inicial e intake, nao spec aprovada.

## Como O Intake Funciona

Durante intake, o Cursor deve registrar:

1. texto original do pedido;
2. `request_type`, como `feature`, `evolution`, `bug`, `refactor`, `docs` ou `investigation`;
3. objetivo;
4. escopo conhecido;
5. fora do escopo quando ja for identificavel;
6. incertezas materiais;
7. specs candidatas;
8. skills candidatas;
9. decisao inicial: perguntar, amendar spec existente ou preparar nova trilha dual-spec.

Exemplo de resposta esperada durante intake:

```text
Objetivo: documentar como revisar reports de trilhas interativas.
Tipo: docs.
Escopo conhecido: documentacao humana, exemplos e checklist de revisao.
Fora do escopo: alterar gates, roadmap fixo, phase state, pkg/*, internal/* ou capability de produto.
Incertezas: confirmar se a mudanca apenas documenta o fluxo ou se tambem altera regras de gate.
Specs candidatas: Spec 700, Spec 701, Spec 680, Spec 681.
Skills candidatas: 00, 05, 21, 23.
Decisao inicial: tratar como intake; perguntar se houver mudanca de regra, ou seguir como amend documental se o escopo permanecer apenas em docs.
```

Pedido ambiguo que deve perguntar ou bloquear:

```text
Execute o interactive_sdd_autopilot do {{PROJECT_SLUG}} para esta solicitacao:
Melhore o autopilot para ser mais seguro.
```

Resultado esperado: o Cursor deve perguntar ou bloquear por ambiguidade material, porque objetivo, escopo, comportamento observavel, diagnostico, testes e criterios de aceite nao estao definidos.

## Como A Decisao De Spec E Tomada

A decisao segue `skills/05-{{PROJECT_SLUG}}-spec-architect/SKILL.md` e `skills/05-{{PROJECT_SLUG}}-spec-architect/references/decision-rules.md`.

Use `amend` quando a mudanca pertence claramente a uma spec ou dominio ja governado. O smoke da Fase 26 fez isso: a solicitacao `phase-26-smoke-onboarding-docs` era documental, pertencia ao dominio da Fase 26 e foi governada por `Spec 700` e `Spec 701`, sem abrir nova dual-spec.

Abra nova trilha dual-spec quando a mudanca introduzir novo bounded context, novo contrato principal, impacto transversal proprio ou diagnostico separado. Nesse caso, a implementacao so pode seguir depois de existirem:

1. spec de construcao com objetivo, escopo, fora de escopo, comportamento observavel, testes e criterios de aceite;
2. spec de diagnostico com sinais observaveis, modos de falha, comandos, evidencias, classificacao, decisao e caminho de report.

Se a decisao entre `amend` e nova trilha dual-spec estiver ausente, a trilha deve bloquear.

## Maquina De Estados Interativa

`automation/INTERACTIVE_AUTOPILOT.md` define os estados permitidos:

1. `idle`;
2. `intake`;
3. `requirements`;
4. `build_spec`;
5. `diagnosis_spec`;
6. `implementation`;
7. `diagnosis`;
8. `report_review`;
9. `next_step_decision`;
10. `blocked`;
11. `completed`.

Fluxo normal:

```text
idle -> intake -> requirements -> build_spec -> diagnosis_spec -> implementation -> diagnosis -> report_review -> next_step_decision -> completed
```

Transicoes tambem podem voltar de `next_step_decision` para `requirements` ou `implementation` quando o report autorizar nova etapa. Qualquer etapa pode ir para `blocked` quando uma stop condition aparecer.

Uma transicao fora da lista de `automation/INTERACTIVE_AUTOPILOT.md` deve ser tratada como inconsistencia de estado.

## Como Acompanhar `INTERACTIVE_STATE.json`

Leia `automation/INTERACTIVE_STATE.json` para saber onde a trilha esta. Os campos humanos mais importantes sao:

1. `mode`: deve ser `interactive_sdd_autopilot`;
2. `status`: mostra a etapa operacional atual ou final;
3. `max_retries_per_step`: limite de retry corretivo, atualmente `2`;
4. `blocked` e `block_reason`: indicam bloqueio ativo;
5. `current_request.request_id`: identifica a solicitacao;
6. `current_request.request_type`: classifica o pedido;
7. `current_request.original_request`: preserva o pedido inicial;
8. `current_request.current_step`: mostra a etapa da trilha;
9. `current_request.governing_specs`: lista specs lidas e governantes;
10. `current_request.build_spec_file` e `diagnosis_spec_file`: apontam specs da trilha;
11. `current_request.report_file`: aponta o report usado como gate;
12. `current_request.retry_count`: mostra retries consumidos;
13. `current_request.last_classification` e `last_decision`: registram o ultimo gate;
14. `current_request.completion_criteria`: mostra os criterios usados para concluir;
15. `smoke_history`: preserva eventos do smoke da Fase 26, incluindo bloqueio e retomada.

No smoke da Fase 26, o estado final registra `last_completed_request = "phase-26-smoke-onboarding-docs"`, `status = "completed"`, `last_classification = "PASS"` e `last_decision = "SMOKE COMPLETE - READY FOR PHASE 26 AUDIT"`.

## Diagnostico E Gate

Ao terminar uma etapa implementavel, o Cursor deve executar o diagnostico definido pela spec diagnostica. Para mudancas documentais como o smoke da Fase 26, o diagnostico aceito incluiu:

1. validacao JSON de `automation/INTERACTIVE_STATE.json`, `automation/ROADMAP.json` e `automation/PHASE_STATE.json`;
2. `go test ./...`;
3. `go test -race ./...`;
4. `ReadLints` nos arquivos alterados;
5. revisao de escopo proibido;
6. justificativa de `make fmt` como nao aplicavel quando nenhum arquivo Go muda;
7. justificativa de OpenAPI como nao aplicavel quando nenhuma API HTTP muda.

O gate so permite avanco quando:

1. specs governantes foram lidas;
2. spec de construcao cobre a etapa;
3. spec de diagnostico define verificacao;
4. diagnostico foi executado ou justificado;
5. report foi gerado ou localizado;
6. o report corresponde a etapa atual;
7. `classification = PASS`;
8. `decision` autoriza avanco ou conclusao;
9. nenhuma stop condition esta ativa;
10. `automation/INTERACTIVE_STATE.json` esta coerente.

Testes verdes sem report nao autorizam avanco. Narrativa, plano futuro, TODO e ausencia aparente de erro tambem nao autorizam avanco.

## Como Um Bloqueio Aparece

Um bloqueio deve aparecer em tres lugares:

1. `automation/INTERACTIVE_STATE.json`, com `status = "blocked"`, `blocked = true`, `block_reason`, etapa afetada e retry preservado;
2. report ou evidencia da trilha, com stop condition acionada e acao humana recomendada;
3. resposta final do Cursor, com motivo de bloqueio e proxima acao humana.

Stop conditions comuns no modo interativo:

1. solicitacao materialmente ambigua sem resposta humana;
2. spec de construcao ausente para comportamento novo;
3. spec de diagnostico ausente para comportamento novo;
4. decisao `amend` versus nova dual-spec ausente;
5. report ausente;
6. report sem classificacao ou decisao final;
7. `classification = PARTIAL`, `FAIL` ou `BLOCKED`;
8. teste obrigatorio falho;
9. tentativa de alterar `ROADMAP.json` ou `PHASE_STATE.json` para contornar gate;
10. necessidade de capability fora de spec, {{UPSTREAM_OPS_NAME}}, hosted service ou decisao humana de arquitetura.

## Como Retomar Apos Bloqueio

A retomada parte de artefatos, nao da memoria da conversa:

1. leia `automation/INTERACTIVE_STATE.json`;
2. leia a spec de construcao;
3. leia a spec de diagnostico;
4. leia o report ou a evidencia do bloqueio;
5. identifique a stop condition exata;
6. corrija apenas a causa do bloqueio dentro do escopo aprovado;
7. preserve `retry_count` e historico;
8. execute novamente o diagnostico;
9. gere ou atualize o report;
10. avance somente se `classification = PASS` e `decision` autorizar.

Se houver decisao humana pendente, nao continue automaticamente.

Prompt de retomada:

```text
Retome o interactive_sdd_autopilot do {{PROJECT_SLUG}} apos bloqueio.
Leia automation/INTERACTIVE_STATE.json, a spec de construcao, a spec de diagnostico e o report atual.
Corrija somente a causa do bloqueio e avance apenas se classification = PASS e decision autorizar.
```

## Exemplo Completo Bem-Sucedido

Este exemplo usa o smoke operacional versionado da Fase 26 como fonte de verdade:

1. Pedido inicial: adicionar ao onboarding humano exemplos de intake, pedido ambiguo, retomada apos bloqueio e checklist de revisao.
2. Intake: `request_id = "phase-26-smoke-onboarding-docs"`, `request_type = "docs"`, objetivo documental, escopo restrito a docs, reports e `automation/INTERACTIVE_STATE.json`.
3. Decisao de spec: `amend` em `Spec 700` e `Spec 701`, porque o pedido pertence ao dominio da Fase 26 e nao cria novo bounded context.
4. Alteracao governada: docs e reports foram atualizados; `pkg/*`, `internal/*`, `ROADMAP.json` e `PHASE_STATE.json` ficaram fora do escopo.
5. Diagnostico: validacao JSON, `go test ./...`, `go test -race ./...` e lints dos arquivos alterados passaram.
6. Report: `docs/reports/phase-26-interactive-sdd-autopilot-smoke-report.md` declarou `classification = PASS`.
7. Decisao: `SMOKE COMPLETE - READY FOR PHASE 26 AUDIT`.
8. Estado: `automation/INTERACTIVE_STATE.json` registrou `status = "completed"` e preservou criterios de conclusao.

Artefatos para conferir:

1. `docs/reports/phase-26-interactive-sdd-autopilot-smoke-evidence.md`;
2. `docs/reports/phase-26-interactive-sdd-autopilot-smoke-report.md`;
3. `automation/INTERACTIVE_STATE.json`;
4. `docs/reports/phase-26-interactive-sdd-autopilot-verification-and-human-onboarding-report.md`.

## Exemplo Completo Bloqueado E Retomado

Este exemplo tambem vem do smoke da Fase 26.

Fluxo bloqueado:

1. A trilha chegou a `report_review` depois de intake, decisao de spec e alteracao documental planejada.
2. Antes da criacao do report de smoke, o gate nao podia avancar porque nao existia report com `classification` e `decision`.
3. A stop condition acionada foi report ausente ou sem classificacao/decisao explicita, conforme `automation/STOP_CONDITIONS.md`.
4. O estado esperado ficou com `status = "blocked"`, `current_step = "report_review"`, `retry_count = 1`, `blocked = true` e acao humana recomendada para gerar report diagnostico e reler o gate.
5. Nenhum gate, criterio de aceite, `ROADMAP.json` ou `PHASE_STATE.json` foi alterado para contornar o bloqueio.

Retomada:

1. A retomada leu `automation/INTERACTIVE_STATE.json`, `Spec 700`, `Spec 701` e `docs/reports/phase-26-interactive-sdd-autopilot-smoke-evidence.md`.
2. A causa do bloqueio foi resolvida com a criacao de `docs/reports/phase-26-interactive-sdd-autopilot-smoke-report.md`.
3. O diagnostico foi executado novamente.
4. O report declarou `classification = PASS`.
5. A decisao final autorizou conclusao do smoke.
6. `retry_count` permaneceu `1`, preservando a tentativa corretiva.
7. `automation/INTERACTIVE_STATE.json` preservou `smoke_history` com `blocked`, `resume` e `completed`.

## Checklist Humano Antes De Aceitar Conclusao

Antes de aceitar uma trilha interativa como concluida, confirme:

1. o pedido inicial foi tratado como intake, nao como spec aprovada;
2. a decisao entre `amend` e nova trilha dual-spec esta registrada;
3. a spec de construcao foi lida e cobre a mudanca;
4. a spec de diagnostico foi lida e define o diagnostico;
5. o report corresponde a etapa atual;
6. `classification = PASS`;
7. `decision` autoriza avanco ou conclusao;
8. `automation/INTERACTIVE_STATE.json` registra request, specs, report, classificacao, decisao, retries e bloqueio quando houver;
9. `automation/ROADMAP.json` e `automation/PHASE_STATE.json` nao foram usados como estado interativo;
10. stop conditions foram respeitadas;
11. nenhuma capability de produto, {{UPSTREAM_OPS_NAME}}, hosted service, dashboard ou deploy tooling foi aberta;
12. gaps restantes estao registrados no report ou na resposta final.

## Artefatos Para Revisar Em PR

Revise, quando aplicavel:

1. specs de construcao e diagnostico;
2. `automation/INTERACTIVE_STATE.json`;
3. report da trilha;
4. comandos e testes executados;
5. docs alteradas;
6. cursor rules;
7. `skills/23-{{PROJECT_SLUG}}-sdd-autopilot/SKILL.md`;
8. report `agent-readiness` em `docs/reports/agent-readiness/` quando a trilha usar `skills/24-agent-readiness-governance/SKILL.md`;
9. ausencia de mudancas indevidas em `automation/ROADMAP.json` e `automation/PHASE_STATE.json`;
10. ausencia de mudancas em `pkg/*` e `internal/*` quando a trilha for apenas governanca, docs ou onboarding.

## Agent Readiness Como Evidencia Auxiliar

Use `@kodus/agent-readiness` somente como insumo auxiliar. O score bruto nao substitui spec, report da trilha, stop condition ou decisao de gate.

Quando uma trilha interativa incluir readiness:

1. leia `skills/24-agent-readiness-governance/SKILL.md`;
2. salve ou referencie o report em `docs/reports/agent-readiness/`;
3. classifique cada achado como `worth_it_for_{{PROJECT_SLUG}}`, `optional_for_{{PROJECT_SLUG}}` ou `out_of_scope_for_{{PROJECT_SLUG}}`;
4. trate achados de app, SaaS, deploy, dashboard, hosted service ou {{UPSTREAM_OPS_NAME}} como fora de escopo salvo spec futura explicita;
5. referencie o report de readiness como evidencia de apoio no report diagnostico da trilha.
</file>

<file path="docs/releases/.gitkeep">

</file>

<file path="docs/reports/agent-readiness/.gitkeep">

</file>

<file path="examples/.gitkeep">

</file>

<file path="go.mod.tmpl">
module {{MODULE_PATH}}

go 1.26.3
</file>

<file path="internal/.gitkeep">

</file>

<file path="pkg/.gitkeep">

</file>

<file path="scripts/check-compliance.sh">
#!/usr/bin/env bash

set -euo pipefail

ROOT_DIR="$(cd "$(dirname "${BASH_SOURCE[0]}")/.." && pwd)"
cd "$ROOT_DIR"

log() {
	printf 'check-compliance: %s\n' "$*"
}

fail() {
	printf 'check-compliance: FAIL: %s\n' "$*" >&2
	exit 1
}

assert_file() {
	local file="$1"
	[[ -f "$file" ]] || fail "missing required file: $file"
}

assert_dir() {
	local dir="$1"
	[[ -d "$dir" ]] || fail "missing required directory: $dir"
}

assert_grep() {
	local pattern="$1"
	local file="$2"
	grep -Eq "$pattern" "$file" || fail "expected pattern [$pattern] in $file"
}

required_files=(
	"docs/ai/ai-contribution-contract.md"
	"docs/ai/task-input-format.md"
	"docs/ai/compliance-exceptions.md"
	".cursor/rules/{{PROJECT_SLUG}}.mdc"
	".github/PULL_REQUEST_TEMPLATE.md"
	".pre-commit-config.yaml"
	".github/workflows/ci.yml"
	".github/dependabot.yml"
	"SECURITY.md"
	"Makefile"
	"scripts/check-compliance.sh"
	"scripts/check-secrets.sh"
	"scripts/check-pr-body.sh"
	"scripts/create-pr-from-template.sh"
	"docs/ai/prompts/implement-task.md"
	"docs/ai/prompts/review-change.md"
	"docs/ai/prompts/diagnose-task.md"
	"docs/ai/prompts/remediation-bugfix.md"
	"docs/ai/briefs/ai-implementation-brief.md"
	"docs/ai/briefs/ai-diagnosis-brief.md"
	"docs/ai/briefs/ai-remediation-bugfix-brief.md"
)

for file in "${required_files[@]}"; do
	assert_file "$file"
done

[[ -x "scripts/check-compliance.sh" ]] || fail "scripts/check-compliance.sh must be executable"
[[ -x "scripts/check-secrets.sh" ]] || fail "scripts/check-secrets.sh must be executable"
[[ -x "scripts/check-pr-body.sh" ]] || fail "scripts/check-pr-body.sh must be executable"
[[ -x "scripts/create-pr-from-template.sh" ]] || fail "scripts/create-pr-from-template.sh must be executable"
assert_dir "test/security"

shopt -s nullglob

prompt_files=(docs/ai/prompts/*.md)
brief_files=(docs/ai/briefs/*.md)
(( ${#prompt_files[@]} > 0 )) || fail "docs/ai/prompts must contain at least one markdown file"
(( ${#brief_files[@]} > 0 )) || fail "docs/ai/briefs must contain at least one markdown file"

security_tests=($(find test/security -type f -name '*_test.go'))
(( ${#security_tests[@]} > 0 )) || fail "test/security must contain at least one *_test.go file"

for dir in $(find pkg -type f -name '*.go' ! -name '*_test.go' -exec dirname {} \; | sort -u); do
	assert_file "$dir/doc.go"
	tests=("$dir"/*_test.go)
	(( ${#tests[@]} > 0 )) || fail "missing *_test.go in $dir"
done

for target in lint vet check-compliance check-pr-body test-security; do
	assert_grep "^${target}:" "Makefile"
	assert_grep "run: make ${target}" ".github/workflows/ci.yml"
done

assert_grep "^secrets-check:" "Makefile"
assert_grep "make secrets-check" ".pre-commit-config.yaml"
assert_grep "package-ecosystem: gomod" ".github/dependabot.yml"
assert_grep "package-ecosystem: github-actions" ".github/dependabot.yml"
assert_grep "make test-security" "SECURITY.md"

placeholder_patterns=(
	'\\bplaceholder\\b'
	'\\bTODO\\b'
	'\\bTBD\\b'
)

operational_files=(
	"Makefile"
	".github/workflows/ci.yml"
	".github/PULL_REQUEST_TEMPLATE.md"
	".pre-commit-config.yaml"
	".cursor/rules/{{PROJECT_SLUG}}.mdc"
	"docs/ai/ai-contribution-contract.md"
	"docs/ai/task-input-format.md"
	"docs/ai/compliance-exceptions.md"
	"${prompt_files[@]}"
	"${brief_files[@]}"
)

for file in "${operational_files[@]}"; do
	for pattern in "${placeholder_patterns[@]}"; do
		if grep -Eiq "$pattern" "$file"; then
			fail "forbidden placeholder marker found in $file"
		fi
	done
done

assert_grep "^## Objetivo$" ".github/PULL_REQUEST_TEMPLATE.md"
assert_grep "^## Specs lidas$" ".github/PULL_REQUEST_TEMPLATE.md"
assert_grep "^## Arquivos alterados$" ".github/PULL_REQUEST_TEMPLATE.md"
assert_grep "^## Impacto em docs e superficie publica$" ".github/PULL_REQUEST_TEMPLATE.md"
assert_grep "^## Comandos executados$" ".github/PULL_REQUEST_TEMPLATE.md"
assert_grep "^## Testes executados$" ".github/PULL_REQUEST_TEMPLATE.md"
assert_grep "^## Excecoes formais$" ".github/PULL_REQUEST_TEMPLATE.md"
assert_grep "^## Gaps restantes$" ".github/PULL_REQUEST_TEMPLATE.md"
assert_grep "^## Checklist$" ".github/PULL_REQUEST_TEMPLATE.md"
assert_grep "PULL_REQUEST_TEMPLATE.md" "AGENTS.md"
assert_grep "scripts/create-pr-from-template.sh" "AGENTS.md"
assert_grep "check-pr-body" ".github/workflows/ci.yml"
assert_grep "^# AI Contribution Contract$" "docs/ai/ai-contribution-contract.md"
assert_grep "^# Task Input Format$" "docs/ai/task-input-format.md"
assert_grep "spec governante" "docs/ai/ai-contribution-contract.md"
assert_grep "docs/ai/compliance-exceptions.md" "docs/ai/ai-contribution-contract.md"
assert_grep "arquivos alterados" "docs/ai/ai-contribution-contract.md"
assert_grep "testes executados" "docs/ai/ai-contribution-contract.md"
assert_grep "gaps restantes" "docs/ai/ai-contribution-contract.md"
assert_grep "^## Campos obrigatorios$" "docs/ai/task-input-format.md"
assert_grep '`Specs lidas`' "docs/ai/task-input-format.md"
assert_grep '`Impacto esperado em docs ou superficie publica`' "docs/ai/task-input-format.md"
assert_grep "^Specs lidas:$" "docs/ai/task-input-format.md"
assert_grep "^Impacto esperado em docs ou superficie publica:$" "docs/ai/task-input-format.md"
assert_grep "docs/ai/compliance-exceptions.md" "docs/ai/task-input-format.md"
assert_grep "^- arquivos alterados$" "docs/ai/task-input-format.md"
assert_grep "^- testes executados$" "docs/ai/task-input-format.md"
assert_grep "^- gaps restantes$" "docs/ai/task-input-format.md"
assert_grep "^# AI Compliance Exceptions$" "docs/ai/compliance-exceptions.md"
assert_grep "^## Estado atual$" "docs/ai/compliance-exceptions.md"
assert_grep "^## Formato de registro$" "docs/ai/compliance-exceptions.md"
assert_grep "^## Excecoes ativas$" "docs/ai/compliance-exceptions.md"

log "ok"
</file>

<file path="scripts/check-file-size.sh">
#!/usr/bin/env bash
set -euo pipefail

WARN_THRESHOLD=${WARN_THRESHOLD:-700}
ERROR_THRESHOLD=${ERROR_THRESHOLD:-1000}
EXCEPTIONS_FILE="${EXCEPTIONS_FILE:-.file-size-exceptions}"

exceptions=()
if [[ -f "$EXCEPTIONS_FILE" ]]; then
  while IFS= read -r line; do
    line="${line%%#*}"
    line="$(echo "$line" | xargs)"
    [[ -n "$line" ]] && exceptions+=("$line")
  done < "$EXCEPTIONS_FILE"
fi

is_excepted() {
  local file="$1"
  for exc in "${exceptions[@]+"${exceptions[@]}"}"; do
    [[ "$file" == "$exc" ]] && return 0
  done
  return 1
}

warnings=0
errors=0

while IFS= read -r file; do
  lines=$(wc -l < "$file" | tr -d ' ')

  if is_excepted "$file"; then
    continue
  fi

  if (( lines > ERROR_THRESHOLD )); then
    echo "ERROR: $file has $lines lines (threshold: $ERROR_THRESHOLD)"
    errors=$((errors + 1))
  elif (( lines > WARN_THRESHOLD )); then
    echo "WARNING: $file has $lines lines (threshold: $WARN_THRESHOLD)"
    warnings=$((warnings + 1))
  fi
done < <(find . -name '*.go' -not -name '*_test.go' -not -path './.git/*' -not -path './vendor/*' | sort)

if (( errors > 0 )); then
  echo ""
  echo "FAIL: $errors file(s) exceed $ERROR_THRESHOLD lines."
  echo "Add justified exceptions to $EXCEPTIONS_FILE or refactor per specs/020 section 8."
  exit 1
fi

if (( warnings > 0 )); then
  echo ""
  echo "INFO: $warnings file(s) exceed $WARN_THRESHOLD lines (review recommended)."
fi

exit 0
</file>

<file path="scripts/check-pr-body.sh">
#!/usr/bin/env bash

set -euo pipefail

ROOT_DIR="$(cd "$(dirname "${BASH_SOURCE[0]}")/.." && pwd)"
cd "$ROOT_DIR"

fail() {
	printf 'check-pr-body: FAIL: %s\n' "$*" >&2
	exit 1
}

usage() {
	cat <<'EOF'
Usage:
  scripts/check-pr-body.sh --body-file <path>
  scripts/check-pr-body.sh < body.md

When running in GitHub Actions on pull_request events, the script also reads the
PR body automatically from GITHUB_EVENT_PATH.
EOF
}

is_dependabot_pull_request_event() {
	local event_path="$1"

	python3 - "$event_path" <<'PY'
import json
import sys

with open(sys.argv[1], "r", encoding="utf-8") as fh:
    payload = json.load(fh)

pull_request = payload.get("pull_request") or {}
user_login = (pull_request.get("user") or {}).get("login", "")
sender_login = (payload.get("sender") or {}).get("login", "")
head_ref = (pull_request.get("head") or {}).get("ref", "")

if (
    user_login in {"dependabot[bot]", "app/dependabot"}
    or sender_login == "dependabot[bot]"
    or head_ref.startswith("dependabot/")
):
    sys.exit(0)

sys.exit(1)
PY
}

body_file=""
temp_file=""

cleanup() {
	if [[ -n "$temp_file" && -f "$temp_file" ]]; then
		rm -f "$temp_file"
	fi
}
trap cleanup EXIT

while [[ $# -gt 0 ]]; do
	case "$1" in
	--body-file)
		shift
		[[ $# -gt 0 ]] || fail "--body-file requires a path"
		body_file="$1"
		;;
	-h|--help)
		usage
		exit 0
		;;
	*)
		fail "unknown argument: $1"
		;;
	esac
	shift
done

if [[ -n "$body_file" ]]; then
	[[ -f "$body_file" ]] || fail "body file not found: $body_file"
elif [[ -n "${GITHUB_EVENT_PATH:-}" && -f "${GITHUB_EVENT_PATH:-}" ]]; then
	if is_dependabot_pull_request_event "$GITHUB_EVENT_PATH"; then
		printf 'check-pr-body: skipped for Dependabot generated dependency PR\n'
		exit 0
	fi

	temp_file="$(mktemp)"
	python3 - "$GITHUB_EVENT_PATH" "$temp_file" <<'PY'
import json
import pathlib
import sys

event_path = pathlib.Path(sys.argv[1])
out_path = pathlib.Path(sys.argv[2])
with event_path.open("r", encoding="utf-8") as fh:
    payload = json.load(fh)
body = payload.get("pull_request", {}).get("body", "")
out_path.write_text(body, encoding="utf-8")
PY
	body_file="$temp_file"
elif [[ ! -t 0 ]]; then
	temp_file="$(mktemp)"
	cat >"$temp_file"
	body_file="$temp_file"
else
	usage
	fail "no PR body source provided"
fi

[[ -s "$body_file" ]] || fail "PR body is empty"

required_headings=(
	"## Objetivo"
	"## Specs lidas"
	"## Skills aplicadas"
	"## Arquivos alterados"
	"## Impacto em docs e superficie publica"
	"## Comandos executados"
	"## Testes executados"
	"## Excecoes formais"
	"## Gaps restantes"
	"## Checklist"
)

for heading in "${required_headings[@]}"; do
	grep -Fqx "$heading" "$body_file" || fail "missing required heading: $heading"
done

forbidden_literals=(
	"Descreva o resultado observavel desta mudanca."
	"- [ ] \`AGENTS.md\`"
	"- [ ] spec(s) governante(s) citadas abaixo"
	"- \`specs/...\`"
	"- \`skills/...\`"
	"- liste os arquivos principais alterados neste PR"
	"- \`none\`, ou descreva o impacto relevante quando existir"
	"- liste comandos executados alem dos checks padrao, quando houver"
	"- descreva gaps, tradeoffs ou itens fora do escopo"
)

for literal in "${forbidden_literals[@]}"; do
	if grep -Fqx -- "$literal" "$body_file"; then
		fail "template placeholder not replaced: $literal"
	fi
done

printf 'check-pr-body: ok\n'
</file>

<file path="scripts/check-secrets.sh">
#!/usr/bin/env bash

set -euo pipefail

ROOT_DIR="$(cd "$(dirname "${BASH_SOURCE[0]}")/.." && pwd)"
cd "$ROOT_DIR"

fail() {
	printf 'check-secrets: FAIL: %s\n' "$*" >&2
	exit 1
}

command -v git >/dev/null 2>&1 || fail "git not found"

tmp_file="$(mktemp)"
cleanup() {
	rm -f "$tmp_file"
}
trap cleanup EXIT

patterns=(
	'AKIA[0-9A-Z]{16}'
	'gh[pousr]_[A-Za-z0-9_]{36,}'
	'github_pat_[A-Za-z0-9_]{22,}'
	'([Aa][Pp][Ii][_-]?[Kk][Ee][Yy]|[Aa][Cc][Cc][Ee][Ss][Ss][_-]?[Tt][Oo][Kk][Ee][Nn]|[Aa][Uu][Tt][Hh][_-]?[Tt][Oo][Kk][Ee][Nn]|[Ss][Ee][Cc][Rr][Ee][Tt]|[Pp][Aa][Ss][Ss][Ww][Oo][Rr][Dd]|[Pp][Aa][Ss][Ss][Ww][Dd])[[:space:]]*[:=][[:space:]]*["'\'']?[A-Za-z0-9_./+=-]{16,}'
	'-----BEGIN (RSA |DSA |EC |OPENSSH |PGP )?PRIVATE KEY-----'
)

skip_file() {
	case "$1" in
	docs/reports/agent-readiness/*.raw.json)
		return 0
		;;
	*)
		return 1
		;;
	esac
}

findings=0

while IFS= read -r -d '' file; do
	[[ -f "$file" ]] || continue
	if skip_file "$file"; then
		continue
	fi

	for pattern in "${patterns[@]}"; do
		if LC_ALL=C grep -IEn -- "$pattern" "$file" >"$tmp_file"; then
			while IFS= read -r line; do
				printf 'check-secrets: potential secret in %s:%s\n' "$file" "$line" >&2
				findings=$((findings + 1))
			done <"$tmp_file"
		fi
	done
done < <(git ls-files -z --cached --others --exclude-standard)

if (( findings > 0 )); then
	fail "potential secrets found; remove the value, rotate it if real, and document only redacted evidence"
fi

printf 'check-secrets: ok\n'
</file>

<file path="scripts/create-pr-from-template.sh">
#!/usr/bin/env bash

set -euo pipefail

ROOT_DIR="$(cd "$(dirname "${BASH_SOURCE[0]}")/.." && pwd)"
cd "$ROOT_DIR"

TEMPLATE_FILE=".github/PULL_REQUEST_TEMPLATE.md"

usage() {
	cat <<'EOF'
Usage:
  scripts/create-pr-from-template.sh [gh pr create args...]

Examples:
  scripts/create-pr-from-template.sh --base main --title "docs: update report"
  scripts/create-pr-from-template.sh --base main --title "feat: x" --body-file pr-body.md

Behavior:
  - If --body-file is provided, the file is validated before calling gh pr create.
  - If --body-file is not provided, the script copies the repository PR template
    to a temporary file, opens it in $EDITOR when interactive, validates it, and
    then uses it as the PR body.
  - Passing --body is rejected to keep the workflow anchored to the repository
    template.
EOF
}

require_command() {
	command -v "$1" >/dev/null 2>&1 || {
		printf 'create-pr-from-template: FAIL: missing required command: %s\n' "$1" >&2
		exit 1
	}
}

require_command gh
require_command python3

body_file=""
gh_args=()

while [[ $# -gt 0 ]]; do
	case "$1" in
	--body)
		printf 'create-pr-from-template: FAIL: --body is not allowed; use --body-file or the repository template flow\n' >&2
		exit 1
		;;
	--body-file)
		gh_args+=("$1")
		shift
		[[ $# -gt 0 ]] || {
			printf 'create-pr-from-template: FAIL: --body-file requires a path\n' >&2
			exit 1
		}
		body_file="$1"
		gh_args+=("$1")
		;;
	-h|--help)
		usage
		exit 0
		;;
	*)
		gh_args+=("$1")
		;;
	esac
	shift
done

temp_file=""
cleanup() {
	if [[ -n "$temp_file" && -f "$temp_file" ]]; then
		rm -f "$temp_file"
	fi
}
trap cleanup EXIT

if [[ -z "$body_file" ]]; then
	temp_file="$(mktemp)"
	cp "$TEMPLATE_FILE" "$temp_file"
	body_file="$temp_file"

	if [[ -t 1 ]]; then
		editor="${EDITOR:-}"
		if [[ -z "$editor" ]]; then
			printf 'create-pr-from-template: FAIL: EDITOR is not set; pass --body-file with a filled template\n' >&2
			exit 1
		fi
		"$editor" "$body_file"
	else
		printf 'create-pr-from-template: FAIL: non-interactive use requires --body-file with a filled template\n' >&2
		exit 1
	fi

	gh_args+=("--body-file" "$body_file")
fi

./scripts/check-pr-body.sh --body-file "$body_file"
gh pr create "${gh_args[@]}"
</file>

<file path="skills/00-skill-index/SKILL.md">
---
name: {{PROJECT_SLUG}}-{{PROJECT_SLUG}}-go-skill-index
description: Entry point for Go agents working on {{PROJECT_SLUG}}. Use whenever you need to know which specialized skill governs a change or to verify global guardrails before starting or finishing any task.
---

# Codex Go Skill Index
Goal: ensure every change applies the right specialized skill and respects global guardrails for {{PROJECT_SLUG}}.

## When to Use
- Always load this index before starting work to route yourself to the correct specialized skill.
- Re-check before code review or PR merge to confirm no relevant skill was missed.

## Non-negotiables
1. **Spec Driven Development**: every feature traces to a spec in `specs/`; never implement behavior without normative coverage.
2. **Architecture Boundaries**: `pkg/` contains stable public contracts; `internal/` contains runtime details; `pkg/app` is the only bridge to `internal/runtime`.
3. **Context Propagation**: propagate `context.Context` as the standard boundary for execution, cancellation and lifecycle.
4. **Effective Go**: follow the 13-go skill for context, errors, interfaces, concurrency, and tests.
5. **Observability**: emit events, logs and hooks through public contracts; never leak internal types into public signatures.

## Do / Don't
- **Do**: Identify the skill ID(s) below that match the asset you modify and load them fully.
- **Do**: Record in PR description which skills were followed.
- **Don't**: Mix patterns from multiple skills without reconciling terminology.
- **Don't**: Introduce new libs or infra outside these skills without design review.

## Interfaces / Contracts
- Use this table to map work areas to skills:
  - `01-hexagonal-architecture`: layered architecture, ports, adapters, composition root wiring.
  - `02-tenant-auth-quotas`: multi-tenant context, auth, rate limits (inherited from infra context -- see note in skill).
  - `03-async-command-processing`: async patterns, command queues, workers (inherited from infra context).
  - `04-streaming-sse-ws`: streaming, SSE, WebSocket replay (inherited from infra context).
  - `05-{{PROJECT_SLUG}}-spec-architect`: spec refinement, dual-spec workflow, prompt generation for {{PROJECT_SLUG}}.
  - `06-multi-provider-model-routing`: model routing, provider config (inherited from infra context).
  - `07-postgres-pgx-migrate`: pgx usage, migrations (inherited from infra context).
  - `08-redis-cache-streams`: caching, locks, streams (inherited from infra context).
  - `09-sqs-fifo-idempotency`: FIFO queues, dedup, DLQ (inherited from infra context).
  - `10-http-gin-openapi`: REST handlers, middleware, OpenAPI (inherited from infra context).
  - `11-grpc-interceptors-health`: gRPC services, interceptors (inherited from infra context).
  - `12-observability-zap-otel-prom`: logging, tracing, metrics (inherited from infra context).
  - `13-go-idiomatic-effective-go`: Go idioms, context, errors, concurrency, tests, file size guidelines.
  - `14-config-viper-flags`: config loading, feature flags (inherited from infra context).
  - `15-devx-ci-precommit-changelog`: tooling, Makefile, CI, pre-commit, changelog.
  - `16-object-calisthenics`: reduce nesting, use strong types, keep functions small in domain and application layers.
  - `17-solid-go-ports`: SOLID principles with ports/adapters and strict dependency direction.
  - `18-security-owasp-api`: API security, input validation, secrets (inherited from infra context).
  - `19-prompt-injection-llm-safety`: LLM safety, prompt injection defense (inherited from infra context).
  - `20-testing-strategy-regression-load-containers`: testing pyramid, regression, conformance, contracts.
  - `21-documentation-open-source`: godoc, README, CONTRIBUTING, CHANGELOG, examples.
  - `22-release-versioning-governance`: SemVer, release readiness, changelog, release notes and distribution path.
  - `23-{{PROJECT_SLUG}}-sdd-autopilot`: interactive SDD autopilot for feature requests, dual-spec trails, diagnosis, reports and gates.
  - `24-agent-readiness-governance`: `@kodus/agent-readiness` analysis filtered for {{PROJECT_SLUG}}'s Go library/framework scope.
  - `25-{{PROJECT_SLUG}}-ultra-rigid-sdd`: ultra-rigid evidence-first SDD for phase planning, autopilot prompt generation, paired implementation and diagnosis specs, reconciliation, readiness/release gates, and framework-vs-application boundaries.

## Checklists
**Before starting**
- [ ] Identify change scope and map to skill IDs above.
- [ ] Load each skill and required resource files.
- [ ] Confirm terminology alignment with `AGENTS.md` and relevant specs.

**During work**
- [ ] Keep architecture boundaries intact (`pkg/` vs `internal/` per `specs/020`).
- [ ] Apply Effective Go and observability patterns consistently.
- [ ] Respect file size guidelines from skill 13 and `specs/020` section 8.

**Before PR**
- [ ] Run relevant tests/linters per skill instructions.
- [ ] Update docs/diagrams referenced by skills if behavior changed.
- [ ] Capture skill checklist status in PR template.
- [ ] If touching core behavior, apply skill `20-testing-strategy-regression-load-containers`.

## Definition of Done
- Referenced skills applied with evidence (code/comments/tests) for every touched subsystem.
- No contradictions with non-negotiables.
- PR checklist attached and green.
- Documentation updated (or explicitly not needed).

## Minimal Examples
- Example PR note: `Skills: 01-hexagonal-architecture, 13-go-idiomatic-effective-go, 20-testing-strategy`. Include checklist results inline.
- Example routing decision: modifying `pkg/agent` + `internal/runtime` => load skills 13 + 17 + 20.
- Example testing routing: modifying conformance suites + memory backend => load skills 20 + 21.
- Example interactive request: user asks for a new feature from free text => load skills 05 + 23, then any domain-specific skills.
- Example readiness review: user asks to run or apply `@kodus/agent-readiness` => load skills 20 + 21 + 24, and skill 23 when it is part of an interactive trail.
- Example ultra-rigid phase planning: user asks for {{PROJECT_SLUG}} phase prompts, paired SDD specs, diagnosis gates, roadmap closure, or readiness confirmation => load skills 05 + 25, and skill 23 when it is part of an interactive trail.
</file>

<file path="skills/01-hexagonal-architecture/SKILL.md">
---
name: hexagonal-architecture
description: Maintain the hexagonal (ports-and-adapters) structure for {{PROJECT_SLUG}}. Enforce layered separation between public contracts (pkg/*), runtime internals (internal/*), and composition root (pkg/app).
---

# Hexagonal Architecture Playbook
Goal: keep the codebase layered so features remain testable, extensible, and free from internal leakage.

## When to Use
- Adding or modifying public contracts, runtime logic, adapters, or command binaries.
- Touching wiring/DI for composition root or server entry points.
- Reviewing PRs that blend responsibilities across layers.

## Non-negotiables
1. `pkg/*` contains stable public contracts; `internal/*` contains runtime details.
2. `pkg/app` is the only public package authorized to import `internal/runtime`.
3. Adapters live behind ports; each adapter implements one interface.
4. CMD packages only assemble dependencies, parse config, and start servers.
5. Prefer small interfaces and composition over large abstractions.

## Do / Don't
- **Do** keep structs immutable via constructors or value copies; return domain errors.
- **Do** define ports near consumers (e.g., `agent.Engine`, `memory.Store`).
- **Do** create DTOs in adapters when external schemas differ from domain models.
- **Don't** import adapters into public contract packages.
- **Don't** expose `internal/*` types in exported signatures, errors, or examples.
- **Don't** leak infrastructure details into public API surface.

## Interfaces / Contracts
- Ports should be consumer-driven and tiny (see Effective Go skill). Example:
  ```go
  type Store interface {
      Load(ctx context.Context, sessionID string) (Snapshot, error)
      Save(ctx context.Context, sessionID string, delta Delta) error
  }
  ```
- Wiring contract: `pkg/app` builds the composition root; tests supply fakes via public interfaces.

## Checklists
**Before coding**
- [ ] Identify which layer changes; confirm dependencies only flow inward.
- [ ] Define/adjust port interfaces before touching adapters.
- [ ] Plan DTOs for any new external schemas.

**During**
- [ ] Keep constructors in adapters, injecting via interfaces.
- [ ] Add unit tests per layer (public contract tests, runtime detail tests, wiring smoke tests).

**After**
- [ ] Verify `go test ./...` in touched packages.
- [ ] Update diagrams/docs if new adapter exists.
- [ ] Mention port/interface changes in PR description.

## Definition of Done
- Layer boundaries intact; no forbidden imports.
- Ports documented and implemented with fakes in tests.
- Composition root compiles and respects config skill.
- Public contract tests cover new logic; adapters have coverage or manual verification noted.

## Minimal Examples
- Adding a new memory backend: create adapter implementing `memory.Store` port; public contract unaware of backend specifics.
- New tool: implement `tool.Tool` interface in adapter; register via `pkg/tool.Registry`.
</file>

<file path="skills/01-hexagonal-architecture/resources/repo_layout.md">
# Repository Layout (Hexagonal)

```
/cmd/
  api/
    main.go              # wires Gin HTTP server, DI container, config loading
  worker/
    main.go              # wires worker loop, SQS consumer, Redis, pgx
/internal/
  domain/
    conversation/
      entity.go          # pure domain models (no infra imports)
      service.go         # business rules (ports only)
  application/
    usecase/
      send_message.go    # orchestrates domain + ports, transactional boundary
    port/
      repo.go            # interfaces consumed by use cases
      notifier.go
  adapter/
    postgres/
      conversation_repo.go  # pgx implementation of repo port
    redis/
      rate_limit_store.go
    aws/
      sqs_producer.go
    crm/
      {{EXTERNAL_SYSTEM_LOWER}}_client.go    # only place where CRM SDK appears
/pkg/
  config/
  logger/
  telemetry/

```

**Guidelines**
- Each adapter implements exactly one port interface; fan-in happens at application layer.
- Domain structs keep JSON tags off; adapters add DTOs for IO formats.
- Command binaries import only `/internal/{application,adapter,domain}` through a wiring package (e.g., `/internal/app`).
- Unit tests live next to code; integration/system tests live under `/test/`.
</file>

<file path="skills/02-tenant-auth-quotas/SKILL.md">
---
name: tenant-auth-quotas
description: Apply multi-tenant ({{DOMAIN_ACTOR}}) context, JWT/API key auth, CRM-agnostic adapters, and per-{{DOMAIN_ACTOR}} quotas/rate limits.
---

> **Nota de contexto:** esta skill foi herdada de outro projeto (sac-agents) e contém
> terminologia específica daquele domínio ({{DOMAIN_ACTOR}}_id, CRM, SQS, etc.).
> Os patterns técnicos são aplicáveis quando a {{PROJECT_SLUG}} adotar infraestrutura equivalente.
> Até lá, tratar como referência conceitual, não como regra operacional direta.

# Tenant Auth & Quotas
Goal: guarantee every request is authenticated, scoped, and throttled per {{DOMAIN_ACTOR}} with CRM adapters isolated.

## When to Use
- Building or modifying auth middleware, tenant context propagation, CRM adapters, or per-{{DOMAIN_ACTOR}} quotas.
- Touching Redis rate limit keys or tenant-aware config.
- Handling {{DOMAIN_ACTOR}} onboarding/offboarding logic.

## Non-negotiables
1. Every inbound request yields a `TenantContext` (see resource file) stored in `context.Context`.
2. Auth supports both JWT and API key; pick via `Authorization` scheme.
3. Rate limiting enforced before calling application layer; tokens stored per {{DOMAIN_ACTOR}}.
4. CRM-specific code stays in adapters and receives sanitized tenant context copies.
5. Logs, metrics, traces include `{{DOMAIN_ACTOR}}_id`, `auth_method`, `rate_limit_tier`.

## Do / Don't
- **Do** short-circuit unauthorized traffic with 401/403 before hex layers.
- **Do** expose derived flags (features, tiers) via context to use cases.
- **Do** cache JWKS per {{DOMAIN_ACTOR}} with TTL + background refresh.
- **Don't** hardcode {{DOMAIN_ACTOR}}-specific behavior; use config skill for overrides.
- **Don't** bypass rate limits in tests without explicit helper toggles.
- **Don't** log raw API keys or JWTs.

## Interfaces / Contracts
- Middleware signature: `func InjectTenantContext(ctx context.Context, headers http.Header) (context.Context, error)`.
- `TenantContext` schema lives in [tenant_context_contract.md](resources/tenant_context_contract.md).
- Rate limit store port example:
  ```go
  type RateLimiter interface {
      Allow(ctx context.Context, {{DOMAIN_ACTOR}}ID, key string, limit int, window time.Duration) (allowed bool, remaining int, reset time.Time, err error)
  }
  ```

## Checklists
**Before coding**
- [ ] Identify auth methods affected (JWT, API key, both).
- [ ] Determine required scopes/tiers for feature.
- [ ] Confirm rate limit thresholds per {{DOMAIN_ACTOR}} (default + overrides).

**During**
- [ ] Validate token signatures and expiration.
- [ ] Populate context with canonical IDs + correlation metadata.
- [ ] Apply Redis quota mutation atomically with scripts or Lua.

**After**
- [ ] Add unit tests covering allowed, throttled, and unauthorized cases.
- [ ] Update OpenAPI/GRPC docs to reflect auth scheme.
- [ ] Verify logs redact secrets.

## Definition of Done
- Tenant context available to downstream handlers/use cases.
- Rate limit metrics exported per {{DOMAIN_ACTOR}} tier.
- CRM adapters receive only sanitized, necessary attributes.
- Error responses include traceable IDs and skill-aligned messages.

## Minimal Examples
- HTTP middleware verifying JWT, storing `TenantContext`, passing to Gin handlers via `ctx.Request.Context()`.
- CLI script creating API key row with hashed secret + tier flag, then invalidating Redis caches.
</file>

<file path="skills/02-tenant-auth-quotas/resources/tenant_context_contract.md">
# Tenant Context Contract

## Headers / Metadata
- `X-{{DOMAIN_ACTOR_TITLE}}-ID` (required) — canonical tenant identifier.
- `X-Request-ID` (required) — upstream-provided, fallback to generated UUID v4.
- `X-Conversation-ID` (optional) — for conversational flows; generate when absent.
- `Authorization` — `Bearer <JWT>` for server-to-server auth or `ApiKey <token>` for key auth.

## Context Struct
```go
type TenantContext struct {
    {{DOMAIN_ACTOR_TITLE}}ID       string
    AuthMethod     AuthMethod // JWT or APIKey
    Subject        string      // user_id or integration name
    Scopes         []string
    ConversationID string
    RateLimitTier  string      // standard, burst, premium
    RequestID      string
}
```

## JWT Expectations
- Validate with shared JWKS cache per {{DOMAIN_ACTOR}}.
- Required claims: `iss` ({{DOMAIN_ACTOR}} auth domain), `sub`, `exp`, `{{DOMAIN_ACTOR}}_id`.
- Optional: `scopes`, `roles`, `tier`.

## API Key Expectations
- Stored hashed in Postgres `{{DOMAIN_ACTOR}}_api_keys` table.
- Keys scoped to features; rate limit tier derived from row.
- Rotate every 90 days; log usage with request metadata.

## Rate Limit Inputs
- Redis bucket key: `rate:{{DOMAIN_ACTOR}}:{{{DOMAIN_ACTOR}}_id}:window:{minute}`.
- Limits expressed as requests per rolling minute, override via {{DOMAIN_ACTOR}} flags.
- Burst handling: allow +20% headroom with TTL-limited counter.

## CRM Adapter Boundary
- Tenant context flows into CRM adapter through interface `CRMContextProvider` returning sanitized fields so CRM-specific code never mutates shared state.
</file>

<file path="skills/03-async-command-processing/SKILL.md">
---
name: async-command-processing
description: Enforce HTTP 202 ACK pattern, SQS FIFO command queueing, and idempotent worker loops for asynchronous commands.
---

> **Nota de contexto:** esta skill foi herdada de outro projeto (sac-agents) e contém
> terminologia específica daquele domínio ({{DOMAIN_ACTOR}}_id, CRM, SQS, etc.).
> Os patterns técnicos são aplicáveis quando a {{PROJECT_SLUG}} adotar infraestrutura equivalente.
> Até lá, tratar como referência conceitual, não como regra operacional direta.

# Async Command Processing
Goal: keep API surfaces responsive (ACK 202) while guaranteeing exactly-once semantics through SQS FIFO and idempotent workers.

## When to Use
- Implementing or updating command handlers, queues, or worker services.
- Changing HTTP endpoints that enqueue work.
- Adjusting DLQ handling or replay tooling.

## Non-negotiables
1. HTTP handlers enqueue command then return `202 Accepted` + `request_id` + `command_id`.
2. Queue is AWS SQS FIFO with `MessageGroupId = {{DOMAIN_ACTOR}}:conversation` and content-based dedup.
3. Worker uses inbox table (Postgres) for idempotency — insert-first, skip if exists.
4. Visibility timeout >= 2x max processing time; extend when doing long calls.
5. DLQ reprocessing documented; no silent drops.

## Do / Don't
- **Do** log enqueue success/failure with metrics `commands_enqueued_total`.
- **Do** batch deletes when finishing messages to reduce API calls.
- **Do** publish completion events to Redis Streams for streaming consumers.
- **Don't** block HTTP handlers waiting for worker completion.
- **Don't** disable dedup; tune TTL if collisions happen.
- **Don't** forget to set `command_type` attribute for observability.

## Interfaces / Contracts
- Command envelope lives in [command_contract.md](resources/command_contract.md).
- Application port example:
  ```go
  type CommandQueue interface {
      Enqueue(ctx context.Context, cmd CommandEnvelope) (commandID string, err error)
  }
  ```
- Worker signature: `func (w *Worker) Handle(ctx context.Context, msg sqs.Message) error` returning error to retry.

## Checklists
**Before coding**
- [ ] Confirm command schema changes across producers/consumers.
- [ ] Determine ack payload (ids, estimated SLA).
- [ ] Size visibility timeout and DLQ config.

**During**
- [ ] Wrap enqueue in context-aware retries with backoff.
- [ ] Write inbox UPSERT logic to enforce idempotency.
- [ ] Emit metrics (enqueue/processing/dlq) + structured logs.

**After**
- [ ] Add integration test using LocalStack for enqueue/dequeue.
- [ ] Document runbook for DLQ replay.
- [ ] Update OpenAPI to state 202 behavior.

## Definition of Done
- HTTP + worker flows compile and pass tests.
- Commands survive retries without duplication.
- DLQ alarm configured (CloudWatch or similar).
- Observability dashboards updated (latency, backlog, DLQ size).

## Minimal Examples
- Gin handler: validate payload, build `CommandEnvelope`, call `queue.Enqueue`, respond `202` with JSON `{ "command_id": "...", "request_id": "..." }`.
- Worker: fetch message, insert inbox row, execute use case, delete message when success; on failure, log and return error.
</file>

<file path="skills/03-async-command-processing/resources/command_contract.md">
# Command Envelope Contract

```json
{
  "version": 1,
  "{{DOMAIN_ACTOR}}_id": "{{DOMAIN_ACTOR}}-123",
  "conversation_id": "conv-789",
  "command_id": "uuid",
  "command_type": "send_agent_reply",
  "payload": {
    "message_id": "uuid",
    "body": "...",
    "metadata": {
      "language": "en",
      "priority": "default"
    }
  },
  "dedup_key": "{{DOMAIN_ACTOR}}-123:conv-789:send_agent_reply:uuid",
  "created_at": "2026-03-05T12:00:00Z"
}
```

- **SQS FIFO** queue name: `cmd-{{DOMAIN_ACTOR}}-conversation.fifo`.
- **MessageGroupId**: `{{DOMAIN_ACTOR}}:{{{DOMAIN_ACTOR}}_id}:conversation:{conversation_id}`.
- **MessageDeduplicationId**: SHA256 of `dedup_key`.
- **DLQ**: `cmd-{{DOMAIN_ACTOR}}-conversation-dlq` with `maxReceiveCount=3`.
- **Attributes**: `request_id`, `trace_id`, `command_type`.

## Worker Expectations
- Poll via long polling (20s) with visibility timeout >= worker SLA.
- Use idempotent inbox table (Postgres) storing `command_id`, `status`, `processed_at`.
- Confirm processing before deleting message; ack errors by leaving message to retry.
- Publish completion events to Redis Streams when needed for streaming skill.
</file>

<file path="skills/04-streaming-sse-ws/SKILL.md">
---
name: streaming-sse-ws
description: Deliver agent updates through Redis Streams with SSE/WebSocket fan-out, cursored replay, and heartbeat guarantees.
---

> **Nota de contexto:** esta skill foi herdada de outro projeto (sac-agents) e contém
> terminologia específica daquele domínio ({{DOMAIN_ACTOR}}_id, CRM, SQS, etc.).
> Os patterns técnicos são aplicáveis quando a {{PROJECT_SLUG}} adotar infraestrutura equivalente.
> Até lá, tratar como referência conceitual, não como regra operacional direta.

# Streaming SSE & WebSocket
Goal: provide consistent, replayable streaming of agent events for UI/webhooks using Redis Streams and SSE/WS transports.

## When to Use
- Adding/modifying streaming endpoints, consumers, or event schemas.
- Working on Redis Streams retention/replay logic.
- Handling UI subscription bugs or performance tuning.

## Non-negotiables
1. Redis Stream per conversation: `stream:conversation:{{{DOMAIN_ACTOR}}_id}:{conversation_id}`.
2. Every event increments `seq` monotonically; use as cursor within payload.
3. SSE default, WebSocket optional but uses same event envelope.
4. Reconnection/resume supported through cursor tokens (Redis IDs).
5. API responses never block on downstream consumers; they publish events asynchronously.

## Do / Don't
- **Do** push events from workers/use cases after durable state write (Postgres).
- **Do** store structured payload JSON; avoid string concatenation.
- **Do** emit `heartbeat` events for idle streams.
- **Don't** keep per-client state in Redis; rely on cursors instead.
- **Don't** drop events silently; log + metric `stream_dropped_events_total`.
- **Don't** re-use SSE endpoint for non-stream payloads.

## Interfaces / Contracts
- Event envelope spec lives in [event_envelope.md](resources/event_envelope.md).
- Publisher port:
  ```go
  type EventStream interface {
      Publish(ctx context.Context, {{DOMAIN_ACTOR}}ID, conversationID string, event Event) (cursor string, err error)
      Read(ctx context.Context, {{DOMAIN_ACTOR}}ID, conversationID, cursor string, count int64) ([]Event, string, error)
  }
  ```
- SSE handler contract: use Gin middleware to upgrade connection and flush per event.

## Checklists
**Before coding**
- [ ] Define event types + payload schemas (delta, completed, error, heartbeat).
- [ ] Decide retention + backpressure thresholds.
- [ ] Plan auth via tenant skill ({{DOMAIN_ACTOR}}_id on context).

**During**
- [ ] Wrap Redis calls with context + timeout; handle `XREAD` cancellation.
- [ ] Write SSE/WS tests for reconnect/resume cases.
- [ ] Ensure gzip disabled (per SSE spec) but TLS enforced.

**After**
- [ ] Update OpenAPI/WebSocket docs with cursor semantics.
- [ ] Add Grafana panels: backlog length, fan-out latency.
- [ ] Run load test for at least 100 concurrent clients.

## Definition of Done
- SSE + WS endpoints resume from cursor and heartbeat on idle.
- Redis retention + metrics configured.
- Event schema documented and versioned.
- Observability includes stream publish/read counters.

## Minimal Examples
- Worker publishes ADK delta: `Event{Type: "agent.delta", Payload: DeltaJSON, Seq: nextSeq}` -> Redis stream -> UI SSE.
- Client reconnect: sends last `cursor`, server `XREAD` from `cursor+1` and replays.
</file>

<file path="skills/04-streaming-sse-ws/resources/event_envelope.md">
# Event Envelope (Redis Streams + SSE/WS)

```
Stream: stream:conversation:{{{DOMAIN_ACTOR}}_id}:{conversation_id}
Entry ID: millisecond seq from Redis (e.g., 1719580000000-0)
Fields:
  seq            -> incrementing int maintained per conversation
  cursor         -> `${timestamp}-${seq}` used for replay tokens
  {{DOMAIN_ACTOR}}_id      -> canonical {{DOMAIN_ACTOR}}
  conversation_id-> canonical conversation
  event_type     -> e.g., agent.delta, agent.completed, system.error
  payload        -> JSON string describing body
  request_id     -> originating HTTP request id
```

## SSE Contract
- Endpoint: `/conversations/:id/events`.
- Query params: `cursor` (optional). Default = latest delivered.
- Response headers: `Cache-Control: no-store`, `Content-Type: text/event-stream`.
- Event format: `event: <event_type>`, `id: <cursor>`, `data: <payload JSON>`.

## WebSocket Contract
- Path: `/ws/conversations/:id/stream`.
- Protocol message:
  ```json
  {
    "type": "subscribe",
    "conversation_id": "conv-1",
    "cursor": "1719580000000-0"
  }
  ```
- Server pushes JSON with fields `cursor`, `event_type`, `payload`.
- Heartbeat every 15s with `event_type=heartbeat`.

## Replay Rules
- Clients send `cursor`; server reads stream via `XREAD` from `cursor+1`.
- Retention: keep last 24h or 500 events, whichever longer.
- When cursor missing/expired, send `410 Gone` SSE or WS error advising restart without cursor.
</file>

<file path="skills/05-{{PROJECT_SLUG}}-spec-architect/SKILL.md">
---
name: {{PROJECT_SLUG}}-spec-architect
description: refine vague {{PROJECT_SLUG}} change requests into a dual-spec package with one build spec and one diagnostic spec, plus cursor/codex execution briefs and a pr checklist. use when the user provides an issue, ticket, prd, bug report, or free text for {{PROJECT_SLUG}} and needs architectural clarification, mandatory questioning, spec drafting or amendment, and enforcement of the rule that every new track must define how it is built and how it is diagnosed.
---

# {{PROJECT_SLUG}} Spec Architect

Conduza pedidos vagos ate virarem um pacote dual-spec executavel no padrao do `{{PROJECT_SLUG}}`. Atue como assistente de arquitetura e gate de especificacao: nao deixe implementacao seguir enquanto houver lacunas materiais na spec de construcao ou na spec de diagnostico.

## Fluxo obrigatorio

1. Leia `references/{{PROJECT_SLUG}}-rules.md` antes de estruturar a resposta.
2. Identifique a natureza da entrada: `issue`, `ticket`, `bug report`, `prd` ou `texto livre`.
3. Localize a spec canonica do dominio em `references/spec-routing-and-examples.md`.
4. Entre em modo de refinamento obrigatorio.
5. Faca perguntas objetivas ate fechar **as duas specs**. Nao assuma comportamento material faltante.
6. Decida entre **amendar spec existente** ou **criar nova trilha dual-spec ligada a anterior** usando `references/decision-rules.md`.
7. Redija duas specs no formato de `references/spec-template.md` e no estilo normativo do repo:
   - uma **spec de construcao**
   - uma **spec de diagnostico**
8. Gere dois prompts operacionais: um para Cursor e outro para Codex.
9. Gere checklist de PR alinhado ao `AGENTS.md`.

## Regras inegociaveis

1. Trate `specs/` como source of truth.
2. Exija rastreabilidade entre pedido, specs, arquivos provaveis, testes e criterios de aceitacao.
3. Nao permita implementacao se objetivo, escopo, fora de escopo, contratos, impacto arquitetural, riscos, trade-offs, testes ou criterios de aceitacao estiverem materialmente incompletos.
4. Nao permita considerar uma trilha nova suficientemente spec driven se existir apenas spec de construcao sem spec de diagnostico.
5. Prefira **estender** as specs existentes do modulo.
6. Crie nova trilha dual-spec apenas quando houver novo bounded context, novo contrato principal ou separacao arquitetural relevante.
7. Use linguagem normativa e observavel. Evite termos vagos como “melhorar”, “otimizar” ou “ajustar” sem comportamento verificavel.
8. Quando o pedido conflitar com specs existentes, aponte o conflito e refine ate haver decisao explicita.
9. Quando faltar informacao critica, continue perguntando. Nao preencha lacunas criticas com premissas.
10. So gere prompt de implementacao depois que as duas specs estiverem fechadas.
11. Sempre inclua bloco arquitetural com impacto em modulos, riscos e trade-offs.
12. Sempre inclua bloco diagnostico com sinais observaveis, modos de falha, logs, traces, metricas, health checks e procedimento de troubleshooting quando aplicavel.

## Modo de refinamento obrigatorio

Durante o refinamento, feche no minimo estes pontos para **construcao** e **diagnostico**:

- problema observavel
- objetivo
- escopo
- fora de escopo
- spec existente relacionada
- secoes ou contratos afetados
- impacto em modulos (`pkg/*`, `internal/*`, `test/*`, `examples/*` quando aplicavel)
- riscos
- trade-offs
- estrategia de testes
- criterios de aceitacao
- sinais esperados em runtime
- modos de falha relevantes
- logs, metricas, traces ou consultas de troubleshooting
- criterio de confirmacao diagnostica

Use `references/refinement-checklist.md` como roteiro. Faca poucas perguntas por vez, mas continue ate fechar as duas specs.

## Regra de decisao: estender vs nova trilha dual-spec

Use `references/decision-rules.md`.

Resumo operacional:
- **Amendar specs existentes** quando a mudanca claramente pertence ao mesmo modulo ou contrato ja governado.
- **Criar nova trilha dual-spec** quando a mudanca introduz novo dominio, novo contrato principal ou deixaria a spec atual confusa e sobrecarregada.
- Nunca duplique regras normativas que ja vivem em outra spec; referencie-as.

## Estrutura de saida obrigatoria

Entregue sempre nesta ordem:

1. **Leitura arquitetural**
   - tipo de entrada normalizada
   - spec(s) relacionada(s)
   - recomendacao: amendar ou criar nova trilha dual-spec
   - lacunas ainda abertas, se houver
2. **Perguntas de refinamento**
   - somente enquanto as specs nao estiverem fechadas
3. **Spec de construcao**
   - no formato do repo
   - com secoes normativas e criterio de aceitacao observavel
4. **Spec de diagnostico**
   - no formato do repo
   - com sinais observaveis, troubleshooting e criterio de confirmacao
5. **Prompt para Cursor**
   - com escopo, arquivos provaveis, limites, testes e formato de resposta
6. **Prompt para Codex**
   - com o mesmo contrato, mas orientado a execucao mais autonoma e verificacao
7. **Checklist de PR**
   - alinhado a `AGENTS.md`

## Prompt para Cursor

Ao gerar o prompt para Cursor:
- instrua a ler `AGENTS.md` e as specs citadas antes de editar
- aponte quais secoes governam a tarefa
- limite a implementacao ao escopo definido
- exija declaracao de arquivos alterados, testes executados e gaps restantes
- exija que divergencias em relacao ao {{UPSTREAM_NAME}} sejam justificadas pela spec ou feature matrix
- exija verificacao de aderencia entre implementacao, spec de construcao e spec de diagnostico

## Prompt para Codex

Ao gerar o prompt para Codex:
- inclua todas as exigencias do prompt do Cursor
- enfatize execucao autonoma com validacao local
- exija proposta de plano breve antes das alteracoes
- exija atualizacao de testes se houver mudanca de comportamento
- exija validacao de logs, traces ou metricas previstos na spec de diagnostico quando aplicavel
- exija resposta final com: specs lidas, secoes atendidas, arquivos alterados, testes executados, riscos, trade-offs, sinais diagnosticados e gaps

## Checklist de PR

Monte checklist objetivo contendo:
- specs lidas e secoes aplicadas
- motivo para amendar ou criar nova trilha dual-spec
- arquivos alterados
- testes adicionados/atualizados/executados
- criterios de aceitacao cobertos
- cobertura diagnostica prevista ou implementada
- gaps restantes
- divergencias intencionais de paridade
- documentacao atualizada

## Exemplos obrigatorios

Use `references/spec-routing-and-examples.md` para modelar exemplos concretos em:
- `030-agent`
- `050-memory`
- `070-workflows`

## Estilo de escrita

1. Escreva em portugues do Brasil.
2. Preserve os nomes tecnicos do repo exatamente como existem em codigo e specs.
3. Prefira frases curtas, normativas e verificaveis.
4. Quando redigir spec, use titulos numerados e listas normativas.
5. Quando ainda faltarem respostas, pare antes da implementacao e faca novas perguntas.
</file>

<file path="skills/05-{{PROJECT_SLUG}}-spec-architect/agents/openai.yaml">
interface:
  display_name: "{{PROJECT_SLUG}} Spec Architect"
  short_description: "Refina pedidos vagos ate virarem specs executaveis do {{PROJECT_SLUG}}, com prompts para Cursor/Codex e checklist de PR."
  category: "developer"
</file>

<file path="skills/05-{{PROJECT_SLUG}}-spec-architect/references/decision-rules.md">
# Decision rules

## Preferir amendar specs existentes
Escolha este caminho quando:
- a mudanca pertence claramente ao mesmo modulo ja governado
- altera comportamento ja previsto parcialmente
- adiciona regra, contrato ou criterio de aceite dentro do mesmo dominio
- a trilha diagnostica continua acoplada ao mesmo runtime ou contrato existente
- manter no mesmo documento melhora a rastreabilidade

## Criar nova trilha dual-spec ligada a anterior
Escolha este caminho quando:
- a mudanca introduz novo bounded context
- a mudanca cria novo contrato principal
- a extensao deixaria a spec atual longa, confusa ou com responsabilidade misturada
- ha impacto transversal que merece governanca propria
- o diagnostico passa a exigir sinais, fluxos operacionais ou troubleshooting de um dominio separado

## Regra pratica
1. Procure a spec canonica do modulo.
2. Se houver aderencia forte de escopo, amende as specs existentes.
3. Se a aderencia for fraca ou houver novo dominio, crie nova trilha dual-spec com referencia cruzada.
4. Nunca duplique regras normativas ja existentes; referencie-as.
</file>

<file path="skills/05-{{PROJECT_SLUG}}-spec-architect/references/refinement-checklist.md">
# Refinement checklist

Use esta checklist para transformar pedido vago em pacote dual-spec fechado.

## Campos obrigatorios de construcao
- problema observavel
- objetivo
- escopo
- fora de escopo
- spec ou dominio relacionado
- contratos afetados
- impacto em modulos
- riscos
- trade-offs
- estrategia de testes
- criterios de aceitacao

## Campos obrigatorios de diagnostico
- sinais observaveis esperados
- sintomas de falha
- hipoteses principais
- logs relevantes
- metricas relevantes
- traces relevantes
- health checks ou consultas operacionais
- procedimento de troubleshooting
- criterio de confirmacao diagnostica

## Perguntas base
1. Qual comportamento observavel precisa mudar?
2. Qual spec ou modulo atual parece governar isso?
3. O que entra e o que fica explicitamente fora?
4. Ha mudanca de API publica, contrato interno ou apenas runtime?
5. Quais arquivos ou areas sao provavelmente afetados?
6. Quais falhas, cancelamentos, limites ou invariantes precisam ser preservados?
7. Como validar por teste que a mudanca foi feita corretamente?
8. Quais criterios de aceitacao tornam a mudanca verificavel?
9. Quais riscos e trade-offs a mudanca introduz?
10. A mudanca cabe em spec existente ou cria nova trilha dual-spec?
11. Quais sinais em runtime indicam comportamento saudavel?
12. Quais logs, metricas ou traces ajudam a diferenciar sucesso, regressao e falha?
13. Como um operador confirma o diagnostico e reduz o espaco de busca?

## Gate de fechamento
O pacote de specs so esta fechado quando todos os campos obrigatorios de construcao e diagnostico estiverem materialmente respondidos.
</file>

<file path="skills/05-{{PROJECT_SLUG}}-spec-architect/references/spec-routing-and-examples.md">
# Routing and examples for {{PROJECT_SLUG}}

## Dominios canonicos
- `030-agent`: API publica do agent, execucao sincrona/streaming, tools, memory, guardrails, hooks.
- `050-memory`: semantica de memoria conversacional, `SessionID`, `Store`, persistencia, isolamento por sessao.
- `070-workflows`: builder, workflow runnable, state, retry, branching, hooks, history.

## Heuristicas de roteamento
- Pedidos sobre execucao do agent, output streaming, hooks ou tool loop tendem a `030-agent`.
- Pedidos sobre continuidade de conversa, persistencia, `SessionID`, adapters de storage ou snapshots tendem a `050-memory`.
- Pedidos sobre cadeia de steps, branching, retry, suspend/resume ou execution history tendem a `070-workflows`.

## Exemplo 1: pedido vago sobre agent
### Pedido inicial
"quero melhorar o streaming do agent"

### Refinamento esperado
- O que muda no comportamento observavel do stream?
- Afeta ordem de eventos, cancelamento, flush final ou tratamento de erro?
- O contrato e do `pkg/agent` ou apenas runtime interno?
- Quais criterios de aceitacao validam a melhoria?
- Quais sinais, logs ou traces distinguem stream saudavel de stream regressivo?

### Direcao recomendada
- Comecar por `030-agent`
- Amendar specs existentes se o contrato ja existir parcialmente
- Exigir spec de construcao e spec de diagnostico antes de gerar prompt de execucao

## Exemplo 2: pedido vago sobre memory
### Pedido inicial
"quero salvar melhor a memoria"

### Refinamento esperado
- O problema e semantica de `SessionID`, estrategia de persistencia ou isolamento?
- O store atual falha em `Load`, `Save` ou concorrencia?
- A mudanca altera contrato publico de `Store`?
- Quais cenarios de sucesso, ausencia de sessao e falha devem ser cobertos?
- Como diagnosticar corrupcao, ausencia ou isolamento incorreto de memoria?

### Direcao recomendada
- Comecar por `050-memory`
- Amendar specs existentes se continuar sendo o mesmo contrato de memoria
- Exigir sinais diagnosticos para confirmar isolamento, persistencia e falhas de store

## Exemplo 3: pedido vago sobre workflows
### Pedido inicial
"quero workflows mais flexiveis"

### Refinamento esperado
- Flexiveis em branching, retry, suspend/resume ou history?
- A mudanca afeta `Builder`, `Workflow`, `State`, `StepResult` ou hooks?
- O comportamento e novo dominio ou extensao do contrato atual?
- Como validar por testes a nova flexibilidade?
- Como diagnosticar loops incorretos, retries excessivos ou history inconsistente?

### Direcao recomendada
- Comecar por `070-workflows`
- Criar nova trilha dual-spec apenas se surgir um dominio separado do contrato atual
- Exigir spec de diagnostico para sinais de runtime, retries e consistencia de history

## Bundle de saida esperado
Depois das specs fechadas, entregue:
1. leitura arquitetural
2. spec de construcao
3. spec de diagnostico
4. prompt Cursor
5. prompt Codex
6. checklist de PR
</file>

<file path="skills/05-{{PROJECT_SLUG}}-spec-architect/references/spec-template.md">
# Templates de spec {{PROJECT_SLUG}}

Use estes formatos como esqueletos minimos, adaptando a numeracao e o titulo ao modulo.

## Template de spec de construcao

```md
# Spec NNN: <titulo do modulo> - construcao

## 1. Objetivo

Este documento define <contrato/modulo/comportamento> em `{{PROJECT_SLUG}}`, com foco em:

- <objetivo 1>
- <objetivo 2>
- <objetivo 3>

Esta spec complementa a `Spec 000`, a `Spec 010` e a `Spec 020` quando aplicavel.

Ficam fora do escopo deste documento:

- <fora de escopo 1>
- <fora de escopo 2>

## 2. Contexto e encaixe arquitetural

- Modulos afetados: `pkg/...`, `internal/...`, `test/...`, `examples/...`
- Contratos relacionados: `Spec 0xx`, `Spec 0yy`
- Impacto em paridade: <nenhum | baixo | medio | alto>

## 3. Regras normativas

1. <regra observavel 1>
2. <regra observavel 2>
3. <regra observavel 3>

## 4. Contratos e comportamento observavel

- <entrada>
- <saida>
- <erros>
- <cancelamento>
- <efeitos colaterais controlados>

## 5. Riscos e trade-offs

- Risco: <...>
- Trade-off: <...>

## 6. Estrategia de testes

- <teste de sucesso>
- <teste de falha>
- <teste de cancelamento/limite quando aplicavel>

## 7. Criterios de aceitacao

1. <criterio verificavel 1>
2. <criterio verificavel 2>
3. <criterio verificavel 3>
```

## Template de spec de diagnostico

```md
# Spec NNN: <titulo do modulo> - diagnostico

## 1. Objetivo

Este documento define como observar, depurar e confirmar o comportamento de <contrato/modulo/comportamento> em `{{PROJECT_SLUG}}`, com foco em:

- <sinais de saude>
- <sinais de regressao>
- <confirmacao diagnostica>

Ficam fora do escopo deste documento:

- <investigacoes nao cobertas>
- <ferramentas externas fora do runtime, se aplicavel>

## 2. Contexto e encaixe arquitetural

- Modulos afetados: `pkg/...`, `internal/...`, `test/...`, `examples/...`
- Contratos relacionados: `Spec 0xx`, `Spec 0yy`
- Dependencias operacionais: <logs | metrics | traces | health checks | consultas>

## 3. Sinais observaveis e modos de falha

1. <sinal saudavel 1>
2. <modo de falha 1>
3. <modo de falha 2>

## 4. Fontes diagnosticas

- Logs: <campos, eventos, correlacao>
- Metricas: <nome, dimensao, interpretacao>
- Traces: <spans, atributos, correlacao>
- Health checks ou consultas: <comandos, endpoints, invariantes>

## 5. Procedimento de troubleshooting

1. <passo 1>
2. <passo 2>
3. <passo 3>

## 6. Riscos e trade-offs

- Risco: <...>
- Trade-off: <...>

## 7. Criterios de confirmacao diagnostica

1. <criterio verificavel 1>
2. <criterio verificavel 2>
3. <criterio verificavel 3>
```

## Regras de estilo
- usar linguagem normativa e observavel
- explicitar o que muda e o que nao muda
- evitar palavras vagas sem criterio verificavel
- manter rastreabilidade com specs relacionadas
- tratar construcao e diagnostico como specs irmas da mesma trilha
</file>

<file path="skills/05-{{PROJECT_SLUG}}-spec-architect/references/{{PROJECT_SLUG}}-rules.md">
# {{PROJECT_SLUG}} operating rules

Use este arquivo como resumo operacional do `AGENTS.md`.

## Source of truth
- `specs/` governa requisitos por modulo.
- `specs/000-compatibility-target.md` define compatibilidade e regras globais.
- `specs/001-non-goals.md` define exclusoes de escopo.
- `specs/010-feature-matrix.md` funciona como checklist de cobertura, prioridade, risco e paridade.
- `specs/020-repository-architecture.md` governa fronteiras entre `pkg/`, `internal/`, `test/` e `examples/`.
- Specs especificas, como `030-agent`, `050-memory` e `070-workflows`, governam contratos de modulo.
- Se implementacao e spec divergirem, a spec prevalece.

## Workflow obrigatorio
1. Ler specs relevantes antes de propor implementacao.
2. Identificar spec e secao antes de alterar codigo.
3. Nao implementar feature fora da spec.
4. Se a spec estiver ausente, incompleta ou ambigua, parar e registrar gap.
5. Toda mudanca de comportamento exige teste.
6. Toda entrega deve dizer qual spec foi atendida e quais criterios de aceitacao foram cobertos.
7. Toda trilha nova deve possuir duas specs complementares:
   - uma spec de construcao, definindo objetivo, escopo, contratos, regras, arquitetura, testes e criterios de aceite;
   - uma spec de diagnostico, definindo sinais observaveis, modos de falha, hipoteses, metricas, logs, traces, troubleshooting e criterios de confirmacao.
8. Nao considerar uma trilha nova suficientemente spec driven se existir apenas spec de construcao sem spec de diagnostico.
9. Se faltar qualquer uma das duas specs, tratar como gap de especificacao antes da implementacao.

## Regras arquiteturais
- `pkg/` contem contratos publicos estaveis.
- `internal/` contem runtime e detalhes operacionais nao exportados.
- Nao expor `internal/*` em API publica.
- Preferir interfaces pequenas e composicao.
- Evitar dependencias externas desnecessarias.
- `pkg/app` e a ponte publica para `internal/runtime`.

## Regras de saida para tarefas de codigo
Sempre exigir que a implementacao final informe:
- specs lidas e secoes usadas
- arquivos alterados
- decisoes e trade-offs
- testes adicionados, atualizados ou executados
- gaps restantes
- divergencias intencionais em relacao ao {{UPSTREAM_NAME}}
- aderencia entre spec de construcao e spec de diagnostico
</file>

<file path="skills/06-multi-provider-model-routing/SKILL.md">
---
name: multi-provider-model-routing
description: Configure and read per-{{DOMAIN_ACTOR}} provider/model routing with Postgres source of truth, Redis cache, and JSONB params validation.
---

> **Nota de contexto:** esta skill foi herdada de outro projeto (sac-agents) e contém
> terminologia específica daquele domínio ({{DOMAIN_ACTOR}}_id, CRM, SQS, etc.).
> Os patterns técnicos são aplicáveis quando a {{PROJECT_SLUG}} adotar infraestrutura equivalente.
> Até lá, tratar como referência conceitual, não como regra operacional direta.

# Multi-provider Model Routing
Goal: allow each {{DOMAIN_ACTOR}} to select providers/models/params safely while keeping Postgres authoritative and Redis fast.

## When to Use
- Updating {{DOMAIN_ACTOR}} model selection, routing logic, or cache behavior.
- Implementing new provider integrations or validation rules.
- Building admin APIs for {{DOMAIN_ACTOR}} model configuration.

## Non-negotiables
1. Postgres table `control.{{DOMAIN_ACTOR}}_model_config` is the single source of truth.
2. Redis cache TTL <=5m; always write-through/invalidate after DB updates.
3. Provider/model combos validated against allowlist prior to persisting.
4. JSONB `params` validated in application layer; store as-is after validation.
5. Config retrieval returns defaults when {{DOMAIN_ACTOR}}-specific row missing.

## Do / Don't
- **Do** wrap DB writes in transactions when multiple tables touched (e.g., audit log).
- **Do** log config loads with `{{DOMAIN_ACTOR}}_id`, `provider`, `model` (no secrets).
- **Do** guard concurrency by using `SELECT ... FOR UPDATE` when editing via admin API.
- **Don't** let workers fetch provider config per message without caching.
- **Don't** hardcode provider parameters in code; keep in JSONB per {{DOMAIN_ACTOR}}.
- **Don't** skip validation just because params are flexible.

## Interfaces / Contracts
- Schema + cache described in [{{DOMAIN_ACTOR}}_model_config_schema.md](resources/{{DOMAIN_ACTOR}}_model_config_schema.md).
- Port example:
  ```go
  type {{DOMAIN_ACTOR_TITLE}}ModelConfigStore interface {
      Get(ctx context.Context, {{DOMAIN_ACTOR}}ID string) ({{DOMAIN_ACTOR_TITLE}}ModelConfig, error)
      Update(ctx context.Context, cfg {{DOMAIN_ACTOR_TITLE}}ModelConfig) error
  }
  ```
- Validation contract: `ValidateProviderParams(provider string, params map[string]any) error` per provider module.

## Checklists
**Before coding**
- [ ] List providers/models affected; confirm allowlist updates.
- [ ] Decide default fallback behavior.
- [ ] Plan cache invalidation strategy.

**During**
- [ ] Fetch config via Redis first, then DB.
- [ ] Validate JSONB structure; reject unknown keys when necessary.
- [ ] Merge feature flags (skill 14) for dynamic overrides.

**After**
- [ ] Add unit/integration tests covering cache hit/miss and invalid configs.
- [ ] Document migrations if schema changed.
- [ ] Update Grafana panels for config load latency/cache hit ratio.

## Definition of Done
- Config API/use case returns correct provider/model per {{DOMAIN_ACTOR}}.
- Cache coherence proven via tests or TTL instrumentation.
- Validation errors clear and localized.
- Downstream consumers (ADK core, LLM adapters) read typed config.

## Minimal Examples
- Cache miss flow: `cfg, err := store.Get(ctx, {{DOMAIN_ACTOR}}ID)` -> miss -> query Postgres -> store in Redis -> return typed struct.
- Admin update: begin tx, update row, insert audit entry, delete Redis key, commit.
</file>

<file path="skills/06-multi-provider-model-routing/resources/seller_model_config_schema.md">
# {{DOMAIN_ACTOR}}_model_config Schema

Table: `control.{{DOMAIN_ACTOR}}_model_config`

| Column              | Type        | Notes |
|---------------------|-------------|-------|
| {{DOMAIN_ACTOR}}_id           | text PK     | matches tenant context |
| provider            | text        | e.g., `openai`, `anthropic`, `vertex` |
| model               | text        | provider-specific model id |
| params              | jsonb       | arbitrary provider config (temperature, top_p, tool hints) |
| max_output_tokens   | int         | guardrail |
| tools               | jsonb       | list of enabled tools |
| updated_at          | timestamptz | audit |
| updated_by          | text        | human or system |

## Redis Cache
- Key: `{{DOMAIN_ACTOR}}:model-config:{{{DOMAIN_ACTOR}}_id}`
- TTL: 5 minutes (refresh early on misses or critical updates).
- Value: JSON copy of table row.

## API Validation Rules
- Validate provider/model against allowlist.
- Validate `params` schema per provider at application layer (not DB constraint).
- Default fallback row stored under {{DOMAIN_ACTOR}} `default` per environment.

## Multi-provider Routing
1. Fetch config from Redis; on miss, query Postgres and warm cache.
2. Merge feature flags (skill 14) before returning to caller.
3. Provide typed struct:
```go
type {{DOMAIN_ACTOR_TITLE}}ModelConfig struct {
    {{DOMAIN_ACTOR_TITLE}}ID string
    Provider string
    Model    string
    Params   map[string]any
    MaxOutputTokens int
    Tools   []string
}
```
</file>

<file path="skills/07-postgres-pgx-migrate/SKILL.md">
---
name: postgres-pgx-migrate
description: Work with Postgres via pgx (pool, transactions, jsonb) and manage multi-schema migrations using golang-migrate.
---

> **Nota de contexto:** esta skill foi herdada de outro projeto (sac-agents) e contém
> terminologia específica daquele domínio ({{DOMAIN_ACTOR}}_id, CRM, SQS, etc.).
> Os patterns técnicos são aplicáveis quando a {{PROJECT_SLUG}} adotar infraestrutura equivalente.
> Até lá, tratar como referência conceitual, não como regra operacional direta.

# Postgres + pgx + Migrations
Goal: keep Postgres the reliable source with safe migrations, pgx best practices, and multi-schema isolation.

## When to Use
- Writing repositories, transactions, or raw SQL via pgx.
- Adding/modifying migrations, schemas, or indexes.
- Handling jsonb fields or multi-schema setups.

## Non-negotiables
1. Use `pgxpool.Pool` injected via ports; no global connections.
2. Context passed to every query with deadlines.
3. Migrations organized per module/schema, executed via Make target.
4. Multi-tenant tables always include `{{DOMAIN_ACTOR}}_id` and relevant indexes.
5. jsonb used only when schema flexibility needed; accompany with constraints if critical.

## Do / Don't
- **Do** use `pgx.NamedArgs` or struct scanning for clarity.
- **Do** wrap multi-step writes in transactions (`BeginTx` + `defer tx.Rollback`).
- **Do** log slow queries (>200ms) with SQL + plan hints.
- **Don't** use `database/sql` directly; stay on pgx.
- **Don't** rely on implicit transactions in migrations.
- **Don't** cast jsonb to text for filtering when indexes can be used.

## Interfaces / Contracts
- Repository port example:
  ```go
  type ConversationRepository interface {
      WithTx(ctx context.Context, fn func(context.Context, pgx.Tx) error) error
      Save(ctx context.Context, tx pgx.Tx, conv Conversation) error
  }
  ```
- Migration standards documented in [migration_rules.md](resources/migration_rules.md).

## Checklists
**Before coding**
- [ ] Determine schema (control/conversation/telemetry) and access pattern.
- [ ] Decide if transaction boundaries need saga/support.
- [ ] Plan indexes for new queries.

**During**
- [ ] Use prepared statements/batch when iterating.
- [ ] Handle pgx errors with `errors.Is(err, pgx.ErrNoRows)` etc.
- [ ] Structure migrations with reversible down scripts.

**After**
- [ ] Run `make migrate MODULE=<module>` locally.
- [ ] Add integration tests hitting new queries.
- [ ] Update ERD/reference docs when schema changes.

## Definition of Done
- Queries context-aware, tested, and instrumented.
- Migrations applied locally without conflicts.
- Rollback path documented.
- Schema changes communicated to dependent skills (config, routing, async).

## Minimal Examples
- Transaction: `repo.WithTx(ctx, func(ctx context.Context, tx pgx.Tx) error { ... })` ensuring commit/rollback.
- Migration snippet adding jsonb column with index:
  ```sql
  ALTER TABLE conversation.messages ADD COLUMN metadata jsonb DEFAULT '{}'::jsonb NOT NULL;
  CREATE INDEX CONCURRENTLY idx_messages_metadata ON conversation.messages USING GIN (metadata);
  ```
</file>

<file path="skills/07-postgres-pgx-migrate/resources/migration_rules.md">
# Migration Rules

## Tooling
- Use `golang-migrate/migrate` with `.sql` files stored under `db/migrations/<module>`.
- Naming: `<timestamp>_<module>_<action>.up.sql` / `.down.sql`.
- Run via Make target `make migrate MODULE=<module>`.

## Multi-schema Layout
- `control` — tenant + config tables.
- `conversation` — chat state, inbox, transcripts.
- `telemetry` — durable metrics/log aggregates.

## Conventions
- Always include `{{DOMAIN_ACTOR}}_id` first column on multi-tenant tables.
- Use `jsonb` for flexible payloads; add GIN indexes when querying nested fields.
- Default `updated_at`/`created_at` via triggers or DEFAULT.

## Transaction Rules
- Wrap migrations altering multiple tables in a transaction.
- Avoid DDL in long-running transactions during peak hours.
- For large data backfills, use `ALTER TABLE ... ADD COLUMN` with defaults `NULL`, then backfill in batches.

## Testing
- Use `pgxpool` against `docker compose up postgres` for integration tests.
- Provide helper SQL under `db/seeds/` for fixture loading.
</file>

<file path="skills/08-redis-cache-streams/SKILL.md">
---
name: redis-cache-streams
description: Use Redis for caches, rate limits, locks, result stores, and streams with consistent key naming and TTLs.
---

> **Nota de contexto:** esta skill foi herdada de outro projeto (sac-agents) e contém
> terminologia específica daquele domínio ({{DOMAIN_ACTOR}}_id, CRM, SQS, etc.).
> Os patterns técnicos são aplicáveis quando a {{PROJECT_SLUG}} adotar infraestrutura equivalente.
> Até lá, tratar como referência conceitual, não como regra operacional direta.

# Redis Cache & Streams
Goal: operate Redis safely for caching, throttling, locking, and streaming without violating {{DOMAIN_ACTOR}} isolation.

## When to Use
- Touching Redis-based caches, rate limits, locks, or stream publishers/consumers.
- Designing result stores for async operations.
- Tuning TTLs or eviction policies.

## Non-negotiables
1. Keys always namespaced with `{{DOMAIN_ACTOR}}_id`; include `conversation_id` when relevant.
2. TTLs defined upfront; no immortal cache entries.
3. Use `go-redis/v9` with context deadlines/timeouts.
4. Lua scripts or atomic commands for rate limits/locks to avoid race conditions.
5. Observability counters for hits/misses, rate limit events, lock contention.

## Do / Don't
- **Do** centralize key patterns referencing [redis_keyspace.md](resources/redis_keyspace.md).
- **Do** instrument `redis_client_duration_seconds` histogram.
- **Do** prefer `XADD`/`XREAD` for streaming requirements.
- **Don't** store large payloads (>512KB) without compression.
- **Don't** rely on `KEYS` in production; use SCAN.
- **Don't** swallow Redis errors; propagate up for retries/fallback.

## Interfaces / Contracts
- Cache port:
  ```go
  type Cache interface {
      Get(ctx context.Context, key string, dest any) (bool, error)
      Set(ctx context.Context, key string, value any, ttl time.Duration) error
      Delete(ctx context.Context, key string) error
  }
  ```
- Rate limiter uses script returning remaining tokens; see tenant skill.
- Stream publisher interface defined in streaming skill.

## Checklists
**Before coding**
- [ ] Decide TTL, eviction strategy, serialization format (JSON, MsgPack).
- [ ] Confirm key pattern and ownership.
- [ ] Evaluate failure modes (cache miss path, fallback).

**During**
- [ ] Use pipelining/batching for multiple operations.
- [ ] Wrap Lua scripts in go helpers for reuse.
- [ ] Add structured logs for lock acquisition/release.

**After**
- [ ] Add tests with miniredis or redis-server via docker compose.
- [ ] Update Grafana dashboards for new keys/metrics.
- [ ] Document runbook for clearing caches or replaying streams.

## Definition of Done
- Keys follow canonical naming and TTL.
- Error handling + fallbacks tested.
- Metrics/logs show hit ratio, rate limit counts, lock contention.
- Stream consumers/producers pass load tests.

## Minimal Examples
- Cache: `cache.Set(ctx, fmt.Sprintf("{{DOMAIN_ACTOR}}:model-config:%s", {{DOMAIN_ACTOR}}ID), cfg, 5*time.Minute)`.
- Lock: `SET lock:tool:{{DOMAIN_ACTOR}}-123:conv-1 token NX PX 15000` followed by Lua compare/delete on release.
</file>

<file path="skills/08-redis-cache-streams/resources/redis_keyspace.md">
# Redis Keyspace Map

| Purpose         | Key Pattern                                           | TTL          |
|-----------------|-------------------------------------------------------|--------------|
| {{DOMAIN_ACTOR_TITLE}} config   | `{{DOMAIN_ACTOR}}:model-config:{{{DOMAIN_ACTOR}}_id}`                     | 5m           |
| Rate limit      | `rate:{{DOMAIN_ACTOR}}:{{{DOMAIN_ACTOR}}_id}:window:{minute}`             | 60s          |
| Locks           | `lock:{namespace}:{{{DOMAIN_ACTOR}}_id}:{resource}`             | 15s default  |
| Conversation KV | `conv:state:{{{DOMAIN_ACTOR}}_id}:{conversation_id}`            | 24h          |
| Result store    | `result:{command_id}`                                 | 15m          |
| Redis Stream    | `stream:conversation:{{{DOMAIN_ACTOR}}_id}:{conversation_id}`   | retention 24h|
| Idempotency     | `inbox:{{{DOMAIN_ACTOR}}_id}:{command_id}`                      | 24h          |

## Scripts
- Lua script for rate limit: atomic increment + expiry when first set.
- Lock acquisition: `SET lock ... NX PX 15000` + unique token, release via Lua compare/delete.

## Observability Keys
- `metrics:rate_limit:{{{DOMAIN_ACTOR}}_id}` for aggregated counters.
- Use Redis MONITOR only locally (heavy!).
</file>

<file path="skills/09-sqs-fifo-idempotency/SKILL.md">
---
name: sqs-fifo-idempotency
description: Operate the SQS FIFO queues with {{DOMAIN_ACTOR}}:conversation grouping, content dedup, DLQ policies, and inbox-based idempotency.
---

> **Nota de contexto:** esta skill foi herdada de outro projeto (sac-agents) e contém
> terminologia específica daquele domínio ({{DOMAIN_ACTOR}}_id, CRM, SQS, etc.).
> Os patterns técnicos são aplicáveis quando a {{PROJECT_SLUG}} adotar infraestrutura equivalente.
> Até lá, tratar como referência conceitual, não como regra operacional direta.

# SQS FIFO & Idempotency
Goal: ensure command processing is exactly-once using SQS FIFO ordering + Postgres inbox ledger and solid DLQ handling.

## When to Use
- Modifying queue producers/consumers, dedup keys, or DLQ logic.
- Building replay tooling or monitoring.
- Changing inbox schema or processing semantics.

## Non-negotiables
1. `MessageGroupId = {{DOMAIN_ACTOR}}:{{{DOMAIN_ACTOR}}_id}:conversation:{conversation_id}`.
2. `MessageDeduplicationId` derived from deterministic key (hash of {{DOMAIN_ACTOR}} + conversation + command).
3. Inbox table guarded by unique constraint on `command_id` (or dedup key) + `{{DOMAIN_ACTOR}}_id`.
4. No manual deletes from queue without persisting state; always let worker ack.
5. DLQ policy matches [dlq_policy.md](resources/dlq_policy.md).

## Do / Don't
- **Do** include tracing attributes (`request_id`, `trace_id`).
- **Do** extend visibility timeout if downstream call > default.
- **Do** store result/outcome in Postgres for auditing.
- **Don't** rely solely on SQS dedup for idempotency; inbox must exist.
- **Don't** log full payloads with PII; use hashes/ids.
- **Don't** perform destructive replays without backup.

## Interfaces / Contracts
- Inbox repository sample:
  ```go
  type InboxStore interface {
      Insert(ctx context.Context, cmd CommandEnvelope) (inserted bool, err error)
      Complete(ctx context.Context, commandID string, status string) error
  }
  ```
- Replay command script references DLQ policy resource.

## Checklists
**Before coding**
- [ ] Confirm dedup key format and retention (5 minutes default for FIFO).
- [ ] Ensure inbox table indexes support new queries.
- [ ] Validate IAM permissions for queue access.

**During**
- [ ] Wrap SQS send/receive with retries/backoff.
- [ ] On worker start, warm up metrics for queue depth/backlog.
- [ ] Test failure paths (throwing error vs panic) to ensure retries happen.

**After**
- [ ] Update runbooks for DLQ replay if semantics changed.
- [ ] Add integration tests via LocalStack covering dedup collisions.
- [ ] Validate CloudWatch alarms triggered appropriately.

## Definition of Done
- Commands processed exactly once per {{DOMAIN_ACTOR}} conversation even under retries.
- DLQ/backlog metrics visible and alarms configured.
- Replay tooling documented and tested.
- Inbox table consistent with queue schema.

## Minimal Examples
- Dedup key builder: `fmt.Sprintf("%s:%s:%s:%s", {{DOMAIN_ACTOR}}ID, conversationID, commandType, payload.MessageID)` hashed with SHA256.
- Inbox upsert pattern: `INSERT ... ON CONFLICT DO NOTHING` returning bool to skip duplicate work.
</file>

<file path="skills/09-sqs-fifo-idempotency/resources/dlq_policy.md">
# DLQ Policy

- Primary queue: `cmd-{{DOMAIN_ACTOR}}-conversation.fifo`
- DLQ: `cmd-{{DOMAIN_ACTOR}}-conversation-dlq`
- Redrive policy: `maxReceiveCount=3`

## Alerting
- CloudWatch alarm when DLQ inflight > 0 for >5m.
- PagerDuty auto-trigger with {{DOMAIN_ACTOR}}_id tags.

## Replay Procedure
1. Investigate root cause using logs (filter by `command_id`).
2. Fix bug/config if needed.
3. Use script `cmd/replay_dlq`:
   - Pull messages from DLQ
   - Recompute dedup id if necessary
   - Push back to main queue with same attributes
4. Document incident with command/{{DOMAIN_ACTOR}} details.

## Retention
- DLQ message retention 14 days.
- Encrypt with KMS (same key as main queue).
</file>

<file path="skills/10-http-gin-openapi/SKILL.md">
---
name: http-gin-openapi
description: Build REST endpoints with Gin + standard middleware, 202 ACK flows, and synchronized OpenAPI specs.
---

> **Nota de contexto:** esta skill foi herdada de outro projeto (sac-agents) e contém
> terminologia específica daquele domínio ({{DOMAIN_ACTOR}}_id, CRM, SQS, etc.).
> Os patterns técnicos são aplicáveis quando a {{PROJECT_SLUG}} adotar infraestrutura equivalente.
> Até lá, tratar como referência conceitual, não como regra operacional direta.

# HTTP (Gin + OpenAPI)
Goal: ship REST handlers that honor tenant middleware chain, async ACK pattern, and stay documented via OpenAPI.

## When to Use
- Creating/updating HTTP handlers, middleware, routing, or OpenAPI specs.
- Working on REST-specific validation or response envelopes.
- Adding endpoints exposing streaming cursors or command IDs.

## Non-negotiables
1. Middleware order matches [http_contracts.md](resources/http_contracts.md).
2. All handlers accept/propagate `context.Context` from `*gin.Context`.
3. Responses follow standard envelopes (data/error) and include `request_id`.
4. Async endpoints return `202 Accepted` with command info.
5. OpenAPI kept in sync and linted.

## Do / Don't
- **Do** validate payloads with binding + struct tags; add explicit enums.
- **Do** convert domain errors to HTTP codes using centralized mapper.
- **Do** attach OTel span attributes ({{DOMAIN_ACTOR}}_id, route, status).
- **Don't** perform long-running work inside handler; enqueue commands.
- **Don't** leak internal error messages; map to sanitized codes.
- **Don't** forget CORS/CSRF requirements when exposing to browsers.

## Interfaces / Contracts
- Handler skeleton:
  ```go
  func (h *Handler) CreateCommand(c *gin.Context) {
      ctx := c.Request.Context()
      tenant := tenantctx.From(ctx)
      var req CreateCommandRequest
      if err := c.ShouldBindJSON(&req); err != nil {
          h.respondError(c, ErrInvalidPayload)
          return
      }
      cmdID, err := h.usecase.Enqueue(ctx, tenant, req)
      if err != nil {
          h.respondError(c, err)
          return
      }
      c.JSON(http.StatusAccepted, gin.H{"request_id": tenant.RequestID, "command_id": cmdID})
  }
  ```
- Contracts + envelopes detailed in [http_contracts.md](resources/http_contracts.md).

## Checklists
**Before coding**
- [ ] Confirm OpenAPI paths/methods needed.
- [ ] Determine auth scopes + rate limit tiers.
- [ ] Define request/response schemas + examples.

**During**
- [ ] Implement handler + tests (table-driven) covering success/error.
- [ ] Update OpenAPI + run lint.
- [ ] Wire route in `cmd/api/main.go` respecting middleware order.

**After**
- [ ] Verify `go test ./internal/adapter/http/...`.
- [ ] Hit endpoint via `curl` or Postman to confirm 202 + headers.
- [ ] Update docs/changelog referencing endpoint.

## Definition of Done
- Handler integrated with middleware, tests, and OpenAPI spec.
- Responses consistent and observability emitting route metrics.
- Async flows enqueued correctly.
- Docs include sample request/response.

## Minimal Examples
- `GET /healthz` returns 200 soon with build info (no tenant context required) but still logs/traces.
- `POST /conversations/:id/messages` binds JSON, enqueues command, returns 202 with command_id.
</file>

<file path="skills/10-http-gin-openapi/resources/http_contracts.md">
# HTTP Contracts

## Middleware Chain (Gin)
1. `RequestID` — ensures `X-Request-ID` header; generates when absent.
2. `TenantAuth` — maps headers to `TenantContext` (skill 02).
3. `RateLimit` — rejects with 429 when quota exceeded.
4. `Recovery` — zap logging + sanitized 500 payload.
5. `Tracing` — OTel HTTP instrumentation.
6. `Metrics` — Prometheus histogram for latency/status.

## Standard Responses
- Success envelope:
  ```json
  {
    "request_id": "...",
    "data": {...}
  }
  ```
- Error envelope:
  ```json
  {
    "request_id": "...",
    "error": {
      "code": "unauthorized",
      "message": "token expired"
    }
  }
  ```
- Async command POST returns `202` with `{ "request_id", "command_id" }`.

## OpenAPI Workflow
1. Update `api/openapi.yaml` after handler changes.
2. Run `make openapi-lint`.
3. Regenerate client/server stubs if needed.
4. Commit spec + generated files in same PR.
</file>

<file path="skills/11-grpc-interceptors-health/SKILL.md">
---
name: grpc-interceptors-health
description: Build gRPC services with standardized interceptor chains, auth, validation, and health checks.
---

> **Nota de contexto:** esta skill foi herdada de outro projeto (sac-agents) e contém
> terminologia específica daquele domínio ({{DOMAIN_ACTOR}}_id, CRM, SQS, etc.).
> Os patterns técnicos são aplicáveis quando a {{PROJECT_SLUG}} adotar infraestrutura equivalente.
> Até lá, tratar como referência conceitual, não como regra operacional direta.

# gRPC Interceptors & Health
Goal: keep our gRPC services consistent with HTTP behavior, including auth/tenant context, validation, and observability.

## When to Use
- Creating/modifying gRPC services, protos, or server wiring.
- Adjusting interceptors, metadata handling, or health endpoints.
- Adding validation via protovalidate/PGV.

## Non-negotiables
1. Interceptor order follows [grpc_contracts.md](resources/grpc_contracts.md) for both unary and streaming.
2. Metadata carries {{DOMAIN_ACTOR}} context identical to HTTP headers.
3. Validation errors map to `codes.InvalidArgument` with details.
4. Health service (`grpc.health.v1.Health`) always registered and ready state tied to dependencies.
5. TLS + mTLS enforced per environment configuration.

## Do / Don't
- **Do** generate protos via buf; keep versioned.
- **Do** instrument metrics/traces per method.
- **Do** convert domain errors to gRPC status via central mapper.
- **Don't** bypass interceptors to run raw handlers.
- **Don't** expose experimental methods without feature flags.
- **Don't** break backwards compatibility without version bump.

## Interfaces / Contracts
- Server builder signature:
  ```go
  func NewServer(cfg Config, deps Deps) *grpc.Server {
      unary := grpc.ChainUnaryInterceptor(requestID, tenantAuth, logging, metrics, tracing, recovery, authz)
      stream := grpc.ChainStreamInterceptor(...)
      return grpc.NewServer(grpc.Creds(cfg.Creds), grpc.UnaryInterceptor(unary), grpc.StreamInterceptor(stream))
  }
  ```
- Contracts for metadata, validation, interceptors in [grpc_contracts.md](resources/grpc_contracts.md).

## Checklists
**Before coding**
- [ ] Update proto definitions + buf.yaml when needed.
- [ ] Determine auth scopes + metadata requirements.
- [ ] Plan validation rules (protovalidate) and mapping to errors.

**During**
- [ ] Regenerate Go code (`buf generate`).
- [ ] Implement service logic inside application/use case layers (not adapter) calling ports.
- [ ] Add unit tests using bufconn or grpc-go testing harness.

**After**
- [ ] Run `buf lint` + `buf breaking` against main.
- [ ] Verify health endpoint responds `SERVING` locally.
- [ ] Update docs referencing new RPCs.

## Definition of Done
- Service builds with interceptors, validation, and health wiring.
- Tests cover success, auth failure, validation failure.
- Observability dashboards updated for RPC metrics.
- Buf artifacts committed.

## Minimal Examples
- New RPC `RunAgent`: validate request via protovalidate, call use case, return streaming updates via server-side streaming with same interceptors.
- Health check toggled to `NOT_SERVING` when Redis unavailable by hooking dependency watcher.
</file>

<file path="skills/11-grpc-interceptors-health/resources/grpc_contracts.md">
# gRPC Contracts

## Interceptor Stack (Unary & Stream)
1. `RequestID` — propagate metadata `x-request-id`.
2. `TenantAuth` — builds TenantContext from metadata (mirrors HTTP headers).
3. `Logging` — zap structured logs.
4. `Metrics` — Prometheus histogram/counters per method.
5. `Tracing` — OTel span injection/extraction.
6. `Recovery` — convert panics to gRPC errors.
7. `Authz` — optional per-method scope checks.

## Health Service
- Use `grpc/health/grpc_health_v1`. Register as `health.Server`.
- Endpoint `/grpc.health.v1.Health/Check` must require no auth.

## Validation
- Use `protovalidate` (bufbuild) or `envoyproxy/protoc-gen-validate`.
- Validation runs in interceptor before hitting handlers.
- Errors mapped to `codes.InvalidArgument` with field paths.

## Metadata Expectations
- `x-{{DOMAIN_ACTOR}}-id`
- `authorization`
- `x-request-id`
- `x-conversation-id` (optional)

## OpenAPI parity
- For shared functionality, keep proto comments + buf registry docs in sync with HTTP spec.
</file>

<file path="skills/12-observability-zap-otel-prom/SKILL.md">
---
name: observability-zap-otel-prom
description: Instrument services with zap JSON logging, OpenTelemetry traces, Prometheus metrics, and Grafana dashboards including correlation IDs.
---

> **Nota de contexto:** esta skill foi herdada de outro projeto (sac-agents) e contém
> terminologia específica daquele domínio ({{DOMAIN_ACTOR}}_id, CRM, SQS, etc.).
> Os patterns técnicos são aplicáveis quando a {{PROJECT_SLUG}} adotar infraestrutura equivalente.
> Até lá, tratar como referência conceitual, não como regra operacional direta.

# Observability (zap + OTel + Prom)
Goal: provide consistent insight across logs, traces, and metrics with correlation IDs for every {{DOMAIN_ACTOR}} interaction.

## When to Use
- Adding/changing logging, tracing, metrics, or Grafana dashboards.
- Introducing new components (HTTP handler, worker, stream) needing instrumentation.
- Investigating incidents needing better observability.

## Non-negotiables
1. Logs emitted via zap in JSON format with required fields (see resource).
2. Traces instrumented using OpenTelemetry SDK exporting to collector defined in config.
3. Metrics exported via Prometheus `/metrics` endpoint with cataloged names.
4. Every log/metric includes `request_id` or `command_id` + `{{DOMAIN_ACTOR}}_id` when available.
5. Grafana dashboards updated when new metrics added.

## Do / Don't
- **Do** wrap contexts with `otel.Tracer` spans at boundaries (HTTP/gRPC/worker).
- **Do** log at `info` for successful checkpoints, `warn` for recoverable errors, `error` for failures.
- **Do** add exemplars to histograms when supported.
- **Don't** log sensitive payloads (PII, secrets); mask before logging.
- **Don't** create ad-hoc metric names; follow catalog.
- **Don't** emit logs inside hot loops unless at debug level (and guard with flag).

## Interfaces / Contracts
- Required log fields in [log_fields.md](resources/log_fields.md).
- Metrics list in [metrics_catalog.md](resources/metrics_catalog.md).
- Standard middleware wiring:
  ```go
  logger := zap.L()
  tracer := otel.Tracer("api")
  meter := global.Meter("api")
  ```

## Checklists
**Before coding**
- [ ] Decide which spans/metrics/logs are necessary.
- [ ] Confirm labels/tags align with catalog.
- [ ] Ensure config skill exposes endpoints/keys.

**During**
- [ ] Add context-aware logging (`logger.With(zap.String("{{DOMAIN_ACTOR}}_id", {{DOMAIN_ACTOR}}ID))`).
- [ ] Wrap operations with spans/histograms.
- [ ] Emit Prom metrics with `promauto` or manual registration.

**After**
- [ ] Run `make telemetry-check` (if available) or `go test` for instrumentation packages.
- [ ] Verify data showing up locally (OTel collector + Prom scrape).
- [ ] Update Grafana dashboards or snapshots.

## Definition of Done
- Logs/traces/metrics confirm new path and share IDs for correlation.
- Dashboards updated with panels for new metrics.
- Alerts configured/adjusted if thresholds changed.
- Documentation/Runbooks mention instrumentation.

## Minimal Examples
- Add HTTP span: `ctx, span := tracer.Start(ctx, "CreateCommand")` -> `defer span.End()` -> record attributes.
- Metric use: `commandsEnqueued.WithLabelValues({{DOMAIN_ACTOR}}ID, commandType).Inc()` after enqueue.
</file>

<file path="skills/12-observability-zap-otel-prom/resources/log_fields.md">
# Required Log Fields (zap JSON)

| Field            | Type   | Notes |
|------------------|--------|-------|
| `timestamp`      | ISO8601| auto
| `level`          | string | info/warn/error
| `message`        | string | short, no punctuation at end
| `request_id`     | string | always present
| `trace_id`       | string | OTel trace span context
| `{{DOMAIN_ACTOR}}_id`      | string | optional for public endpoints, required elsewhere
| `conversation_id`| string | when available
| `command_id`     | string | for async flows
| `component`      | string | e.g., http.api, worker.processor
| `error`          | string | err message when level>=error
| `duration_ms`    | number | when logging operation durations

Add structured payloads via nested objects (zap fields) instead of string concatenation.
</file>

<file path="skills/12-observability-zap-otel-prom/resources/metrics_catalog.md">
# Metrics Catalog

| Metric | Type | Labels | Description |
|--------|------|--------|-------------|
| `http_requests_total` | counter | `route`, `method`, `status`, `{{DOMAIN_ACTOR}}_id` | Count of HTTP requests |
| `http_request_duration_seconds` | histogram | `route`, `{{DOMAIN_ACTOR}}_id` | Latency |
| `commands_enqueued_total` | counter | `command_type`, `{{DOMAIN_ACTOR}}_id` | Number of commands queued |
| `command_processing_duration_seconds` | histogram | `command_type`, `{{DOMAIN_ACTOR}}_id` | Worker processing time |
| `redis_operations_total` | counter | `operation`, `result` | Redis ops |
| `redis_operation_duration_seconds` | histogram | `operation` | Redis latency |
| `rate_limit_blocks_total` | counter | `{{DOMAIN_ACTOR}}_id`, `tier` | Requests throttled |
| `adk_step_total` | counter | `{{DOMAIN_ACTOR}}_id`, `step`, `status` | ADK tool usage |
| `adk_run_duration_seconds` | histogram | `{{DOMAIN_ACTOR}}_id` | Agent run durations |
| `grpc_requests_total` | counter | `service`, `method`, `code` | RPC calls |
| `otel_traces_exported_total` | counter | `backend` | Exporter health |
</file>

<file path="skills/13-go-idiomatic-effective-go/SKILL.md">
---
name: go-idiomatic-effective-go
description: Apply Effective Go patterns: context propagation, error handling, small interfaces, safe concurrency, table-driven tests.
---

# Go Idiomatic Guide
Goal: ensure every Go change follows Effective Go practices for maintainability and safety.

## When to Use
- All Go changes (default skill) especially domain/application logic, concurrency, and testing.
- Reviewing PRs for idiomatic pitfalls.

## Non-negotiables
1. Every IO/public function accepts `context.Context` as first parameter.
2. Errors wrapped with `%w`, no panics for control flow.
3. Interfaces live close to consumers and remain minimal.
4. Concurrency uses context-aware `errgroup`/channels; no leaks.
5. Tests are table-driven with deterministic seeds/fakes.

## Do / Don't
- **Do** return `(value, error)` instead of `(bool)` flags.
- **Do** prefer composition over inheritance-like embedding when clarity suffers.
- **Do** implement `String()`/`MarshalJSON` on domain types when helpful.
- **Don't** expose struct fields publicly without need.
- **Don't** swallow errors; wrap and propagate.
- **Don't** spin goroutines without cancellation plan.

## File Size Guidelines

Go source files should remain reasonably small and focused on a single responsibility.

### Recommended File Size

| Lines of Code | Guideline                                         |
| ------------- | ------------------------------------------------- |
| 150–400       | Ideal size                                        |
| 400–700       | Acceptable if the file has a clear responsibility |
| 700+          | Reevaluate the file structure                     |
| 1000+         | Strong code smell – consider splitting the file   |

### Design Principles

1. **Prefer cohesion over arbitrary size limits.**
   A file should group types and functions that belong to the same concept.

2. **Avoid "god files".**
   Files that contain models, business logic, persistence logic, and HTTP handlers should be split.

3. **Prefer multiple small files over one large file.**
   The Go compiler compiles packages, not individual files, so there is no meaningful performance penalty.

4. **Name files according to responsibility.**

Example:

```
order/
    order.go
    service.go
    repository.go
```

Instead of:

```
models.go
services.go
utils.go
```

### When to Split a File

Consider splitting when:

* The file contains multiple unrelated types.
* Navigation becomes difficult.
* Different architectural layers appear in the same file.
* The file exceeds ~700 lines.

### Rule of Thumb

> Organize Go files by **responsibility**, not by arbitrary grouping or size.

## Interfaces / Contracts
- Error handling reference: [error_handling.md](resources/error_handling.md).
- Concurrency patterns: [concurrency_patterns.md](resources/concurrency_patterns.md).
- Interface template:
  ```go
  type CommandQueue interface {
      Enqueue(ctx context.Context, cmd CommandEnvelope) (string, error)
  }
  ```

## Checklists
**Before coding**
- [ ] Decide ownership of context + cancellation.
- [ ] Identify domain boundaries requiring interfaces.
- [ ] Plan tests (unit vs integration) + fakes.

**During**
- [ ] Keep functions short; extract helpers when logic grows.
- [ ] Use `errors.Is/As` when branching on errors.
- [ ] Manage goroutines with errgroup or explicit Stop channels.

**After**
- [ ] Run `go test ./...` and linters (go vet, staticcheck if available).
- [ ] Review diff for wrapped errors and context usage.
- [ ] Update tests to cover regressions + edge cases.

## Definition of Done
- Code builds, tests pass, lints clean.
- No data races (run `go test -race` when concurrency touched).
- Interfaces documented and fakes implemented for tests.
- Error messages short, actionable, no punctuation.

## Minimal Examples
- Table-driven test snippet:
  ```go
  func TestRateLimiter(t *testing.T) {
      cases := []struct {
          name string
          limit int
          wantAllow bool
      }{
          {"under", 10, true},
          {"over", 0, false},
      }
      for _, tc := range cases {
          t.Run(tc.name, func(t *testing.T) {
              got := limiter.Allow(tc.limit)
              if got != tc.wantAllow {
                  t.Fatalf("Allow()=%v want %v", got, tc.wantAllow)
              }
          })
      }
  }
  ```
- Context usage: `func (s *Service) Handle(ctx context.Context, req Request) error { ctx, span := tracer.Start(ctx, "Service.Handle"); defer span.End(); ... }`.
</file>

<file path="skills/13-go-idiomatic-effective-go/resources/concurrency_patterns.md">
# Concurrency Patterns

## Context + ErrGroup
```go
ctx, cancel := context.WithTimeout(parent, 30*time.Second)
defer cancel()
eg, ctx := errgroup.WithContext(ctx)
for _, task := range tasks {
    task := task
    eg.Go(func() error {
        return doWork(ctx, task)
    })
}
if err := eg.Wait(); err != nil {
    return fmt.Errorf("work failed: %w", err)
}
```

- Always respect ctx.Done() inside goroutines.
- No goroutine leaks: ensure go routines exit on cancel.

## Channels
- Prefer typed channels with clear ownership; close only by sender.
- Buffer small amounts when bridging to IO (e.g., streaming deltas) but flush on shutdown.

## Worker Pools
- Use struct encapsulating pool; avoid global channels.
- Provide `Stop()` that closes input and waits for workers.
</file>

<file path="skills/13-go-idiomatic-effective-go/resources/error_handling.md">
# Error Handling Patterns

## Wrapping
- Use `fmt.Errorf("context: %w", err)` for propagation.
- Define sentinel errors in package and compare with `errors.Is`.

## Classification
- Domain errors (`ErrNotFound`, `ErrUnauthorized`) stay in application layer for mapping to transport codes.
- Infrastructure errors bubble up with contextual message; add retries in caller.

## Logging
- Log errors once per request at boundary (transport or worker) with request metadata.

## Retries
- Use `backoff` (e.g., `cenkalti/backoff/v4`) or custom linear/exponential for IO operations.
- Cancel retries when context done.

## Testing
- Provide fake implementations returning errors to test handling.
</file>

<file path="skills/14-config-viper-flags/SKILL.md">
---
name: config-viper-flags
description: Manage configuration via Viper, load env-aware settings, and operate feature flags per {{DOMAIN_ACTOR}}.
---

> **Nota de contexto:** esta skill foi herdada de outro projeto (sac-agents) e contém
> terminologia específica daquele domínio ({{DOMAIN_ACTOR}}_id, CRM, SQS, etc.).
> Os patterns técnicos são aplicáveis quando a {{PROJECT_SLUG}} adotar infraestrutura equivalente.
> Até lá, tratar como referência conceitual, não como regra operacional direta.

# Config & Feature Flags
Goal: centralize configuration (Viper) and {{DOMAIN_ACTOR}}-specific feature flags without hardcoding per-environment differences.

## When to Use
- Changing config structure, adding options, or wiring Viper usage.
- Implementing feature flag evaluation or overrides.
- Handling secrets, env var parsing, or config-driven behavior.

## Non-negotiables
1. Viper loads defaults, config file, then env vars (highest priority).
2. Config objects passed via dependency injection; no global `viper.Get` outside setup.
3. Feature flags stored in Postgres + cached in Redis (60s TTL) with per-{{DOMAIN_ACTOR}} overrides.
4. {{DOMAIN_ACTOR_TITLE}}-level toggles evaluated in application layer before calling providers.
5. Secrets fetched from secure store (SSM/Doppler) not committed.

## Do / Don't
- **Do** create typed config structs covering all modules; avoid map[string]any.
- **Do** expose config via `pkg/config` package returning immutable copies.
- **Do** document new fields in [config_shape.md](resources/config_shape.md).
- **Don't** read env vars deep inside business logic.
- **Don't** use feature flags for sensitive security controls (deploy gating instead).
- **Don't** rely on default values silently; log config summary on startup.

## Interfaces / Contracts
- Config struct example:
  ```go
  type AppConfig struct {
      Env   string
      HTTP  HTTPConfig
      Redis RedisConfig
      Flags FlagConfig
  }
  ```
- Feature flag store: `flagStore.Enabled(ctx, {{DOMAIN_ACTOR}}ID, "streaming_default")`.
- Config shape documented in resource file.

## Checklists
**Before coding**
- [ ] Decide which module(s) need new config knobs.
- [ ] Define defaults + env overrides.
- [ ] Plan migration path for existing deployments.

**During**
- [ ] Update config structs + Viper bindings.
- [ ] Add validation logic (panic on missing required fields) during startup.
- [ ] Implement flag evaluation paths with caching + metrics.

**After**
- [ ] Update docs/resources + sample env files.
- [ ] Add tests ensuring config parsing works (use `t.Setenv`).
- [ ] Bump config version in release/changelog.

## Definition of Done
- Config struct + Viper wiring compile and pass tests.
- Feature flags accessible per {{DOMAIN_ACTOR}} with redis cache + fallback.
- Secrets handled securely.
- Documented shape + usage instructions.

## Minimal Examples
- Startup: `cfg := config.Load()` -> `server := app.NewServer(cfg)`.
- Flag check: `if !flagStore.Enabled(ctx, {{DOMAIN_ACTOR}}ID, "streaming_default") { disableStreaming() }`.
</file>

<file path="skills/14-config-viper-flags/resources/config_shape.md">
# Config Shape

Use Viper loading order: env vars > config file > defaults.

```yaml
app:
  env: dev|staging|prod
  http_port: 8080
  grpc_port: 9090
  log_level: info
  feature_flags:
    global:
      streaming_default: true
    {{DOMAIN_ACTOR}}_overrides:
      {{DOMAIN_ACTOR}}-123:
        streaming_default: false
        provider_override: "openai:gpt-4o"
redis:
  url: redis://localhost:6379
postgres:
  url: postgres://postgres:postgres@localhost:5432/app?sslmode=disable
sqs:
  queue_url: http://localstack:4566/000000000000/cmd-{{DOMAIN_ACTOR}}-conversation.fifo
telemetry:
  otel_collector: http://otel-collector:4317
  metrics_port: 2112
```

## Feature Flag Access
```go
type FlagStore interface {
    Enabled(ctx context.Context, {{DOMAIN_ACTOR}}ID, flag string) bool
}
```
- Flags stored in Postgres `control.feature_flags` table, cached in Redis for 60s.

## Secrets
- Pull from AWS SSM or Doppler; never commit to repo.
</file>

<file path="skills/15-devx-ci-precommit-changelog/SKILL.md">
---
name: devx-ci-precommit-changelog
description: Maintain developer experience: docker-compose stack, Makefile targets, GitHub Actions CI, pre-commit hooks, conventional commits, changelog automation.
---

# DevX, CI & Release Hygiene
Goal: keep local dev setup, CI, and release process smooth so every engineer/agent can ship confidently.

## When to Use
- Editing Makefile, docker-compose, CI workflows, or pre-commit configs.
- Adding dependencies to local stack (postgres/redis/localstack/otel/prom/grafana).
- Working on release automation, changelog, or commit conventions.

## Non-negotiables
1. Make targets documented and kept in sync with automation (see resource).
2. Local stack uses docker-compose with services: Postgres, Redis, LocalStack, OTel Collector, Prometheus, Grafana.
3. Pre-commit hooks installed and passing before pushing.
4. GitHub Actions workflows kept green; updates include workflow docs.
5. Conventional commits enforced plus changelog updates via git-cliff.

## Do / Don't
- **Do** use `.env` files for local credentials; exclude from git.
- **Do** provide troubleshooting steps in PRs when dev-env changes.
- **Do** ensure CI secrets referenced via GitHub env vars, not plaintext.
- **Don't** break backward compatibility for make targets; if renaming, provide shim.
- **Don't** skip updating docs when docker services change.
- **Don't** push directly to main; rely on PR with CI status.

## Interfaces / Contracts
- Make targets listed in [make_targets.md](resources/make_targets.md).
- CI pipeline & release process described in [ci_pipeline.md](resources/ci_pipeline.md).
- Pre-commit config expected under `.pre-commit-config.yaml` referencing hooks above.

## Checklists
**Before coding**
- [ ] Identify which developer workflow component changes.
- [ ] Confirm dependencies/services impacted.
- [ ] Notify teammates if downtime required.

**During**
- [ ] Update Makefile/docker-compose/CI config atomically.
- [ ] Test `make dev-up` + `make dev-down` locally.
- [ ] Run `pre-commit run --all-files` to ensure hook validity.

**After**
- [ ] Ensure GitHub Actions succeed on branch (use workflow_dispatch if needed).
- [ ] Update `CHANGELOG.md` via `make release` (dry run acceptable) for notable changes.
- [ ] Document new steps in README or skill resources as needed.

## Definition of Done
- Make + CI + pre-commit reflect latest commands.
- Local stack instructions verified and reproducible.
- Conventional commit message used; changelog updated if release-worthy.
- Automation scripts (git-cliff, commitlint) still functional.

## Minimal Examples
- Adding new service: update `docker-compose.yml`, extend `make dev-up` to depend on service, document env vars.
- Changing CI step: modify `.github/workflows/ci.yml`, run `act` or remote run, mention in PR.
</file>

<file path="skills/15-devx-ci-precommit-changelog/resources/ci_pipeline.md">
# CI/CD Pipeline

## Pre-commit
- Hooks: `gofmt`, `golines`, `golangci-lint`, `go test ./...`, `markdownlint`, `commitlint` (conventional commits).
- Run `pre-commit install` after cloning.

## GitHub Actions
1. `ci.yml`
   - Steps: checkout -> setup Go -> cache -> `make lint` -> `make test` -> upload coverage.
2. `integration.yml`
   - Spins docker-compose (postgres, redis, localstack) -> runs integration tests -> publishes artifacts.
3. `release.yml`
   - Triggered on tags `v*` -> runs `git-cliff` to generate changelog -> builds docker images -> pushes to registry.

## Conventional Commits
- Format: `type(scope): subject`
- Types: feat, fix, chore, docs, refactor, test, ci, perf, build.
- Subject in imperative mood; max ~50 chars.

## Changelog
- `git-cliff` config under `.github/changelog.toml`.
- Run `make release` to update `CHANGELOG.md` automatically.
</file>

<file path="skills/15-devx-ci-precommit-changelog/resources/make_targets.md">
# Make Targets

| Target | Description |
|--------|-------------|
| `make bootstrap` | install go tools, pre-commit hooks |
| `make lint` | run golangci-lint, staticcheck |
| `make test` | run `go test ./...` with race + coverage |
| `make migrate MODULE=<module>` | run golang-migrate against local DB |
| `make dev-up` | docker compose up (postgres, redis, localstack, otel, prom, grafana) |
| `make dev-down` | stop compose stack |
| `make openapi-lint` | validate OpenAPI spec |
| `make buf` | run buf lint/generate |
| `make telemetry-check` | verify metrics endpoint + otel collector |
| `make release` | run changelog + tag (git-cliff + semver) |
</file>

<file path="skills/16-object-calisthenics/SKILL.md">
---
name: object-calisthenics-go-core
description: Apply Object Calisthenics adapted to Go in pkg/* and internal/* to reduce complexity, improve readability, and preserve hexagonal boundaries.
---

# Object Calisthenics for Go Core
Goal: reduce accidental complexity in public contracts and runtime internals while preserving architectural boundaries.

## When to Use
- Changing or reviewing `pkg/*` and `internal/*`.
- Refactoring long functions, nested conditionals, or bulky structs.
- Creating input/output types that carry structured identifiers.

## Non-negotiables
1. Max 1 indentation level per function; use guard clauses and early returns.
2. Avoid `else`; return early and keep the happy path flat.
3. Keep functions small (target <= 20 lines of logic; refactor when > 30).
4. No primitive obsession for identifiers; use strong types when clarity benefits.
5. Prefer max 3 function arguments; use input structs for larger inputs.
6. Avoid god structs; split by responsibility and compose explicitly.
7. Encapsulate collections with domain rules (example: `Messages`).
8. `pkg/*` must not import `internal/*` except through the authorized bridge (`pkg/app`).

## Do / Don't
- **Do** extract helpers named by intent (`validateTenantScope`, `buildCommandEnvelope`).
- **Do** centralize invariant checks near domain types.
- **Do** keep side effects at the edges (ports/adapters).
- **Don't** pass raw `string` for tenant identifiers across layers.
- **Don't** mix validation, orchestration, and infrastructure in one method.
- **Don't** add conditional trees when polymorphism or map-based dispatch is clearer.

## Interfaces / Contracts
- Canonical typed identifiers:
  ```go
  type {{DOMAIN_ACTOR_TITLE}}ID string
  type ConversationID string
  type RequestID string
  type Seq int64
  ```
- Input object for use cases:
  ```go
  type SendMessageInput struct {
      {{DOMAIN_ACTOR_TITLE}}ID       {{DOMAIN_ACTOR_TITLE}}ID
      ConversationID ConversationID
      RequestID      RequestID
      Body           string
  }
  ```
- Collection wrapper with invariant:
  ```go
  type Messages struct {
      items []Message
  }
  ```

## Checklists
**Before**
- [ ] Identify functions with nested control flow or too many args.
- [ ] Identify primitive IDs that should become strong types.
- [ ] Confirm invariant ownership (domain type vs use case).

**During**
- [ ] Flatten control flow with guard clauses.
- [ ] Extract helpers by intention, not by technical detail.
- [ ] Keep tenant scope (`{{DOMAIN_ACTOR}}_id`) explicit in all core inputs.

**After**
- [ ] Re-read function top-to-bottom: happy path is obvious.
- [ ] Ensure signatures are minimal and typed.
- [ ] Confirm no adapter imports leaked into `domain`/`application`.

## Definition of Done
- Core code is flatter, smaller, and easier to scan.
- Inputs/IDs are typed and tenant-safe.
- Domain invariants are explicit and tested.
- No hexagonal boundary violations.

## Minimal Examples
- Replace:
  ```go
  if req.{{DOMAIN_ACTOR_TITLE}}ID == "" {
      return ErrInvalid{{DOMAIN_ACTOR_TITLE}}ID
  } else {
      // continue
  }
  ```
  With:
  ```go
  if req.{{DOMAIN_ACTOR_TITLE}}ID == "" {
      return ErrInvalid{{DOMAIN_ACTOR_TITLE}}ID
  }
  // continue
  ```
- Prefer `SendMessageInput` over `SendMessage(ctx, {{DOMAIN_ACTOR}}ID, conversationID, requestID, body, source, locale string)`.
- Full before/after snippets: [examples_go.md](resources/examples_go.md).
</file>

<file path="skills/16-object-calisthenics/resources/examples_go.md">
# Object Calisthenics in Go: Before/After

## 1) One indentation level + no else

Before:
```go
func (u *SendMessageUseCase) Execute(ctx context.Context, in SendMessageInput) error {
    if in.{{DOMAIN_ACTOR_TITLE}}ID != "" {
        if in.ConversationID != "" {
            if in.Body != "" {
                return u.repo.Save(ctx, in)
            } else {
                return ErrEmptyBody
            }
        } else {
            return ErrInvalidConversationID
        }
    } else {
        return ErrInvalid{{DOMAIN_ACTOR_TITLE}}ID
    }
}
```

After:
```go
func (u *SendMessageUseCase) Execute(ctx context.Context, in SendMessageInput) error {
    if in.{{DOMAIN_ACTOR_TITLE}}ID == "" {
        return ErrInvalid{{DOMAIN_ACTOR_TITLE}}ID
    }
    if in.ConversationID == "" {
        return ErrInvalidConversationID
    }
    if in.Body == "" {
        return ErrEmptyBody
    }
    return u.repo.Save(ctx, in)
}
```

## 2) Strong types over primitives

Before:
```go
func BuildStreamKey({{DOMAIN_ACTOR}}ID, conversationID string) string {
    return {{DOMAIN_ACTOR}}ID + ":" + conversationID
}
```

After:
```go
type {{DOMAIN_ACTOR_TITLE}}ID string
type ConversationID string

func BuildStreamKey({{DOMAIN_ACTOR}}ID {{DOMAIN_ACTOR_TITLE}}ID, conversationID ConversationID) string {
    return string({{DOMAIN_ACTOR}}ID) + ":" + string(conversationID)
}
```

## 3) Limit arguments with input struct

Before:
```go
func NewCommand({{DOMAIN_ACTOR}}ID, requestID, conversationID, body, channel, locale string) Command
```

After:
```go
type NewCommandInput struct {
    {{DOMAIN_ACTOR_TITLE}}ID       {{DOMAIN_ACTOR_TITLE}}ID
    RequestID      RequestID
    ConversationID ConversationID
    Body           string
    Channel        string
    Locale         string
}

func NewCommand(in NewCommandInput) Command
```

## 4) Encapsulate collections with invariants

Before:
```go
type Conversation struct {
    Messages []Message
}
```

After:
```go
type Messages struct {
    items []Message
}

func (m *Messages) Append(msg Message) error {
    if msg.Seq <= 0 {
        return ErrInvalidSeq
    }
    m.items = append(m.items, msg)
    return nil
}

type Conversation struct {
    Messages Messages
}
```
</file>

<file path="skills/17-solid-go-ports/SKILL.md">
---
name: solid-go-hexagonal-ports
description: Apply SOLID in Go with hexagonal architecture using small ports, explicit adapters, and strict dependency direction for {{PROJECT_SLUG}}.
---

# SOLID for Go + Ports/Adapters
Goal: enforce SOLID with Go idioms so public contracts (`pkg/*`) and runtime internals (`internal/*`) stay stable, testable, and independent from infrastructure.

## When to Use
- Creating or changing public interfaces in `pkg/*`.
- Implementing new adapters (memory backends, tool adapters, server transports).
- Reviewing contracts for responsibility leaks across `pkg/` and `internal/` boundaries.

## Non-negotiables
1. DIP: public contracts depend on interfaces only; runtime wiring belongs to `pkg/app` and `internal/runtime`.
2. ISP: ports are small (1-3 methods) and defined by consumer needs.
3. SRP: public types handle contracts; runtime handles orchestration; adapters integrate external systems.
4. OCP: new backends arrive as new adapters; core logic stays unchanged.
5. LSP: fakes/mocks preserve invariants, errors, and expected behavior contracts.
6. `pkg/*` never imports `internal/*` except through `pkg/app`.
7. Contracts keep context explicit via `context.Context` as the standard boundary.

## Do / Don't
- **Do** split broad interfaces into focused ports.
- **Do** return classified errors from public contracts, not provider-specific details.
- **Do** keep adapter mapping isolated at boundaries.
- **Don't** put orchestration logic in transport handlers.
- **Don't** make a shared mega-port for unrelated capabilities.
- **Don't** branch core logic by provider name; use adapter polymorphism.

## Interfaces / Contracts
- Consumer-owned port:
  ```go
  type Store interface {
      Load(ctx context.Context, sessionID string) (Snapshot, error)
      Save(ctx context.Context, sessionID string, delta Delta) error
  }
  ```
- Engine boundary:
  ```go
  type Engine interface {
      Run(ctx context.Context, agent *Agent, req Request) (Response, error)
      Stream(ctx context.Context, agent *Agent, req Request) (Stream, error)
  }
  ```
- Error contract: fakes must return `ErrNotFound`, `ErrInvalidConfig`, or wrapped errors equivalent to real adapters.

## Checklists
**Before**
- [ ] Decide which layer owns the new behavior (public contract / runtime / adapter).
- [ ] Define minimal ports from consumer needs.
- [ ] Identify invariant/error contract to preserve in fakes.

**During**
- [ ] Keep interfaces in consumer package, not provider package.
- [ ] Ensure adapter only translates protocol/SDK concerns.

**After**
- [ ] Add unit tests with fakes for contract behavior.
- [ ] Verify new adapter works by wiring changes only in `pkg/app` or `cmd/*`.
- [ ] Confirm no public package imports internal packages directly.

## Definition of Done
- Ports are minimal, cohesive, and consumer-driven.
- New backends can be added without core edits.
- Fakes behave like production adapters for success and failures.
- Dependency direction remains hexagonal and compile-time safe.

## Minimal Examples
- Adding a memory backend: implement `memory.Store` in adapter package; public contract unaware of backend specifics.
- Split interface:
  ```go
  type Engine interface { Run(ctx context.Context, agent *Agent, req Request) (Response, error) }
  type HistorySink interface { Append(ctx context.Context, entry HistoryEntry) error }
  ```
- Review checklist: [ports_checklist.md](resources/ports_checklist.md).
</file>

<file path="skills/17-solid-go-ports/resources/ports_checklist.md">
# Ports Review Checklist (Go + Hexagonal)

## Architecture
- [ ] Port is defined on consumer side (`application` use case package).
- [ ] `domain`/`application` import no adapter package.
- [ ] Wiring for concrete adapter exists only in `cmd/*`.

## Interface Shape (ISP + DIP)
- [ ] Interface has 1-3 methods.
- [ ] Method names express business intent, not transport/SDK terms.
- [ ] Method arguments carry required scope (`{{DOMAIN_ACTOR}}_id`, `conversation_id`, `request_id`).
- [ ] No interface combines unrelated responsibilities.

## Responsibilities (SRP)
- [ ] Handler only parses/validates I/O and returns HTTP/gRPC response.
- [ ] Use case coordinates policy, idempotency, and orchestration.
- [ ] Adapter maps between core DTOs and external APIs/DB/queue.

## Extension Safety (OCP)
- [ ] Adding provider/CRM requires only new adapter + wiring.
- [ ] No core switch/case by provider in `domain`/`application`.
- [ ] Shared contracts remain stable after extension.

## Substitutability (LSP)
- [ ] Fakes return same semantic errors as real adapter.
- [ ] Fakes preserve constraints (timeouts, not-found, conflicts).
- [ ] Tests validate behavior, not implementation details.

## Observability and Tenant Safety
- [ ] Logs include `{{DOMAIN_ACTOR}}_id`, `request_id`, `conversation_id` when relevant.
- [ ] Metrics/traces preserve same correlation fields.
- [ ] No cross-tenant read/write path in adapter logic.
</file>

<file path="skills/18-security-owasp-api/SKILL.md">
---
name: security-owasp-api-baseline
description: Enforce OWASP API Security baseline for auth, tenant isolation, validation, headers, secrets handling, and CI security checks.
---

> **Nota de contexto:** esta skill foi herdada de outro projeto (sac-agents) e contém
> terminologia específica daquele domínio ({{DOMAIN_ACTOR}}_id, CRM, SQS, etc.).
> Os patterns técnicos são aplicáveis quando a {{PROJECT_SLUG}} adotar infraestrutura equivalente.
> Até lá, tratar como referência conceitual, não como regra operacional direta.

# OWASP API Security Baseline
Goal: keep APIs tenant-safe and production-hardened with consistent controls from request entry to CI gates.

## When to Use
- Any change in auth, authorization, tenant resolution, middleware, or endpoints.
- Touching request parsing, validation, headers, rate limiting, secrets, or logging.
- Updating CI, dependencies, linters, or release automation affecting security posture.

## Non-negotiables
1. AuthN/AuthZ mandatory and scoped by `{{DOMAIN_ACTOR}}_id`; deny by default on missing/invalid scope.
2. Rate limiting required per `{{DOMAIN_ACTOR}}_id` and per endpoint.
3. Validate and bound all inputs (body/query/header/path) with explicit size limits.
4. Use secure defaults: CORS policy explicit, TLS in production, security headers enabled.
5. Return standardized errors; never leak stack traces, internals, credentials, or provider payloads.
6. Access secrets only via `SecretsProvider`; never log tokens/keys/secrets.
7. Logs must sanitize PII and sensitive content while preserving `{{DOMAIN_ACTOR}}_id`, `request_id`, `conversation_id`, `seq`.
8. CI must run `govulncheck` and `gosec` (directly or via `golangci-lint`) and keep dependency versions pinned/reviewed.

## Do / Don't
- **Do** enforce tenant scope in middleware and re-check in use case for critical actions.
- **Do** reject oversized payloads early (413/400).
- **Do** keep audit-safe logs and metrics for auth/rate-limit denials.
- **Don't** trust CRM/webhook payloads without validation/sanitization.
- **Don't** accept wildcard CORS in production.
- **Don't** expose raw error values from drivers/SDKs directly to clients.

## Interfaces / Contracts
- Auth context contract:
  ```go
  type AuthContext struct {
      {{DOMAIN_ACTOR_TITLE}}ID       string
      Subject        string
      Scopes         []string
      RequestID      string
      ConversationID string
  }
  ```
- Error envelope contract:
  ```json
  {"error":{"code":"forbidden","message":"access denied","request_id":"req-123"}}
  ```
- References:
  - [security_headers.md](resources/security_headers.md)
  - [ci_security_checks.md](resources/ci_security_checks.md)

## Checklists
**Before**
- [ ] Identify attack surface (endpoint, middleware, adapter, worker input).
- [ ] Define authz scope and rate-limit strategy by `{{DOMAIN_ACTOR}}_id`.
- [ ] Define max input sizes and accepted content types.

**During**
- [ ] Add/keep validation for path/query/header/body.
- [ ] Enforce secure headers and CORS policy.
- [ ] Sanitize logs and map internal errors to safe public codes.

**After**
- [ ] Run security checks (`govulncheck`, `gosec`/`golangci-lint`).
- [ ] Verify no secrets or PII leakage in logs and errors.
- [ ] Confirm observability includes denied/auth/rate-limit metrics.

## Definition of Done
- Endpoint/auth flow is tenant-scoped and deny-by-default.
- Input size/shape constraints are explicit and tested.
- Security headers and error envelope are consistent.
- CI security gates are active and passing.

## Minimal Examples
- Reject oversized body:
  ```go
  c.Request.Body = http.MaxBytesReader(c.Writer, c.Request.Body, 1<<20) // 1 MiB
  ```
- Safe error mapping:
  ```go
  if errors.Is(err, ErrForbidden) {
      writeError(c, http.StatusForbidden, "forbidden", "access denied")
      return
  }
  ```
</file>

<file path="skills/18-security-owasp-api/resources/ci_security_checks.md">
# CI Security Checks (Makefile + GitHub Actions)

## Mandatory Gates
- `govulncheck` on module code.
- `gosec` scan (direct or through `golangci-lint`).
- Dependency hygiene: pinned versions and automated update review (Dependabot/Renovate if configured).

## Makefile Example
```makefile
.PHONY: security
security:
	GOCACHE=/tmp/go-build GOPATH=/tmp/go go run golang.org/x/vuln/cmd/govulncheck@latest ./...
	gosec ./...

.PHONY: ci-checks
ci-checks: test lint security
```

If `gosec` runs through `golangci-lint`, keep:
```yaml
# .golangci.yml
linters:
  enable:
    - gosec
```

## GitHub Actions Example
```yaml
name: ci
on:
  pull_request:
  push:
    branches: [main]

jobs:
  checks:
    runs-on: ubuntu-latest
    steps:
      - uses: actions/checkout@v4
      - uses: actions/setup-go@v5
        with:
          go-version-file: go.mod
      - run: make test
      - run: make lint
      - run: make security
```

## Dependency Hygiene
- Keep direct dependencies explicit in `go.mod`.
- Avoid floating tool versions in CI where reproducibility matters.
- If Dependabot is enabled, require PR review for security updates before merge.
</file>

<file path="skills/18-security-owasp-api/resources/security_headers.md">
# Recommended Security Headers

Apply at HTTP edge (Gin middleware/reverse proxy). Keep consistent across API and docs endpoints where applicable.

## Required
- `X-Content-Type-Options: nosniff`
  - Prevent MIME sniffing.
- `X-Frame-Options: DENY`
  - Prevent clickjacking for API surfaces.
- `Referrer-Policy: no-referrer`
  - Minimize referrer leakage.
- `Permissions-Policy: geolocation=(), microphone=(), camera=()`
  - Disable unused browser capabilities.
- `Content-Security-Policy: default-src 'none'; frame-ancestors 'none'; base-uri 'none'`
  - For API responses/docs where feasible; tune per route if serving assets.
- `Cache-Control: no-store` (for auth/session/sensitive endpoints)
  - Avoid caching sensitive data.

## Conditional
- `Strict-Transport-Security: max-age=31536000; includeSubDomains`
  - Enable only in HTTPS production environments.
- `Access-Control-Allow-Origin`
  - Never `*` for authenticated endpoints; use explicit allowlist.
- `Access-Control-Allow-Headers`
  - Keep minimal (`Authorization`, `Content-Type`, correlation headers as needed).

## Correlation and Safety
- Always propagate `request_id` in response headers (for support/audit).
- Never include secrets, raw tokens, or provider stack traces in headers.
</file>

<file path="skills/19-prompt-injection-llm-safety/SKILL.md">
---
name: prompt-injection-llm-safety
description: Protect agent orchestration and tools against prompt injection, tool injection, data exfiltration, and cross-tenant leakage.
---

> **Nota de contexto:** esta skill foi herdada de outro projeto (sac-agents) e contém
> terminologia específica daquele domínio ({{DOMAIN_ACTOR}}_id, CRM, SQS, etc.).
> Os patterns técnicos são aplicáveis quando a {{PROJECT_SLUG}} adotar infraestrutura equivalente.
> Até lá, tratar como referência conceitual, não como regra operacional direta.

# Prompt Injection and LLM Safety
Goal: enforce policy gates and least-privilege boundaries so agent/tool flows remain safe, tenant-scoped, and observable.

## When to Use
- Changing ADK orchestration, tool registration, tool execution, or handoff actions.
- Ingesting untrusted CRM/email/HTML content into prompts.
- Implementing agent features that read/write tenant data or call external systems.

## Non-negotiables
1. Separate trusted `system/policy` instructions from untrusted `user/tool` content.
2. Enforce tool allowlist by agent profile and `{{DOMAIN_ACTOR}}_id` feature flag.
3. Apply least privilege: every tool call scoped to `{{DOMAIN_ACTOR}}_id` + `conversation_id` + `request_id`.
4. Execute mandatory `PolicyGate` before applying `OrchestratorActions` (handoff/transfer/escalate/tool-call).
5. Sanitize/normalize inbound content (strip active HTML/script, truncate, normalize encoding).
6. Emit logs + metrics for all blocked actions and policy denials.
7. Never allow cross-tenant context retrieval or tool execution.

## Do / Don't
- **Do** classify content source (`trusted`, `untrusted`, `external`) before prompting.
- **Do** include deny reasons and policy IDs in audit logs.
- **Do** fail closed on validator/policy timeouts.
- **Don't** let model output directly execute tools without validation.
- **Don't** pass raw HTML/email bodies as trusted instructions.
- **Don't** return hidden prompts, secrets, or other-tenant data in responses.

## Interfaces / Contracts
- Suggested ports:
  ```go
  type PolicyGate interface {
      Evaluate(ctx context.Context, action OrchestratorAction, context PolicyContext) (Decision, error)
  }

  type ToolInputValidator interface {
      Validate(ctx context.Context, tool string, input map[string]any) (Decision, error)
  }

  type ContentSanitizer interface {
      Sanitize(ctx context.Context, content ContentInput) (SanitizedContent, error)
  }
  ```
- Decision contract:
  ```go
  type Decision struct {
      Allow  bool
      Reason string
      RuleID string
  }
  ```
- References:
  - [threat_model.md](resources/threat_model.md)
  - [policy_gate_examples.md](resources/policy_gate_examples.md)

## Checklists
**Before**
- [ ] Define trust boundaries and untrusted sources.
- [ ] Define per-{{DOMAIN_ACTOR}} tool allowlist and feature flags.
- [ ] Define deny behavior and observability fields.

**During**
- [ ] Sanitize inbound CRM/tool text before prompt assembly.
- [ ] Evaluate policy gate before any privileged action.
- [ ] Validate tool inputs and enforce {{DOMAIN_ACTOR}}/conversation scope.

**After**
- [ ] Add tests for allow/deny paths, including timeout/fail-closed cases.
- [ ] Verify metrics/logs for blocked prompt/tool injection attempts.
- [ ] Confirm no cross-tenant leakage in traces, logs, or stream events.

## Definition of Done
- Prompt/tool execution path has explicit policy checkpoints.
- Unsafe actions are denied with auditable reasons.
- Tool access is least-privilege and {{DOMAIN_ACTOR}}-scoped.
- Injection and exfiltration scenarios are covered by tests.

## Minimal Examples
- Fail-closed gate:
  ```go
  decision, err := gate.Evaluate(ctx, action, policyCtx)
  if err != nil || !decision.Allow {
      return ErrPolicyDenied
  }
  ```
- Scope check before tool run:
  ```go
  if input.{{DOMAIN_ACTOR_TITLE}}ID != ctx{{DOMAIN_ACTOR_TITLE}}ID || input.ConversationID != ctxConversationID {
      return ErrScopeViolation
  }
  ```
</file>

<file path="skills/19-prompt-injection-llm-safety/resources/policy_gate_examples.md">
# Policy Gate Examples

## 1) Deny cross-tenant tool call

Action:
```json
{
  "type": "tool_call",
  "tool": "crm_get_ticket",
  "input": {
    "{{DOMAIN_ACTOR}}_id": "{{DOMAIN_ACTOR}}-b",
    "conversation_id": "conv-10",
    "ticket_id": "123"
  }
}
```

Context:
```json
{
  "{{DOMAIN_ACTOR}}_id": "{{DOMAIN_ACTOR}}-a",
  "conversation_id": "conv-10",
  "request_id": "req-1"
}
```

Decision:
```json
{
  "allow": false,
  "reason": "{{DOMAIN_ACTOR}} scope mismatch",
  "rule_id": "scope.{{DOMAIN_ACTOR}}.match"
}
```

## 2) Deny disallowed tool by profile

Action:
```json
{
  "type": "tool_call",
  "tool": "secrets_dump",
  "input": {}
}
```

Decision:
```json
{
  "allow": false,
  "reason": "tool not allowlisted for {{DOMAIN_ACTOR}} profile",
  "rule_id": "tool.allowlist.profile"
}
```

## 3) Allow sanitized CRM summary

Input content:
```json
{
  "source": "crm_email_html",
  "text": "<html><body>Pedido 123<script>alert(1)</script></body></html>"
}
```

Sanitized:
```json
{
  "text": "Pedido 123",
  "truncated": false
}
```

Decision:
```json
{
  "allow": true,
  "reason": "content sanitized and scope valid",
  "rule_id": "content.sanitized.ok"
}
```

## 4) Fail-closed on validator timeout

Decision:
```json
{
  "allow": false,
  "reason": "tool input validator timeout",
  "rule_id": "validator.timeout.fail_closed"
}
```
</file>

<file path="skills/19-prompt-injection-llm-safety/resources/threat_model.md">
# LLM Safety Threat Model

| Threat | Vector | Impact | Required Mitigations | Detection Signals |
|---|---|---|---|---|
| Direct prompt injection | User message tries to override system policy | Unsafe tool calls, policy bypass | Strict system/user separation, PolicyGate before actions, fail-closed | Spike in policy denials by rule ID |
| Indirect prompt injection | CRM/email/HTML includes hidden instructions | Agent follows attacker text from external content | ContentSanitizer, source tagging (`untrusted`), tool allowlist | Sanitizer strips script/tags; blocked tool attempts |
| Tool injection | Model fabricates tool name/params outside allowlist | Unauthorized actions or data writes | ToolInputValidator + allowlist by {{DOMAIN_ACTOR}} profile | Denied tool names/params in logs |
| Data exfiltration | Prompt asks for secrets/internal prompts | Secret leakage, privacy incident | Secret redaction, deny sensitive intents, output filters | Redaction counters, denied exfil intents |
| Cross-tenant leakage | Tool query omits or tampers `{{DOMAIN_ACTOR}}_id` filter | Tenant data breach | Mandatory {{DOMAIN_ACTOR}} scope in all tool contracts, query guards, tests | Scope violation metrics/logs |
| Context window poisoning | Large hostile content dominates prompt | Policy dilution and wrong actions | Truncation + priority ordering + policy restatement | High truncation count + repeated deny |
| Action hijack via handoff | Malicious content requests transfer/escalation | Unauthorized workflow transitions | PolicyGate on `handoff`/`transfer` actions | Denied orchestrator action metrics |

## Invariant Fields
- Always carry `{{DOMAIN_ACTOR}}_id`, `request_id`, `conversation_id`, and `seq` in decision logs/metrics.
- Block decisions must include `reason` and `rule_id`.
</file>

<file path="skills/20-testing-strategy-regression-load-containers/SKILL.md">
---
name: testing-strategy-regression-load-containers
description: Apply when changing async flows, workers, streams, database, migrations, model routing, or core business logic to keep feature behavior authoritative while enforcing regression, load, and testcontainers-based reliability checks.
---

# Testing Strategy: Regression, Load & Containers
Goal: preserve business behavior first, then prove it with fast unit tests, reliable container-based integration tests, critical-path E2E checks, and repeatable load runs.

## When to Use
- Implementing or changing business rules in `domain` or `application`.
- Modifying async flow pieces: HTTP `202`, workers, SQS FIFO, Redis Streams, SSE/WS replay.
- Changing Postgres schemas, migrations, `{{DOMAIN_ACTOR}}_model_config`, inbox/idempotency, Redis TTLs, or queue contracts.
- Adding CI gates, `make` targets, or load/performance coverage.
- Reviewing a failing test where the correct fix is unclear and behavior must be checked against contract/ADR.

## Non-negotiables
1. Features and approved contracts win; tests do not invent behavior on their own.
2. If a failing test implies a behavior change, call that out explicitly and require contract/ADR approval before changing production logic.
3. When a test fails, the order is fixed: validate contract/ADR -> adjust the test to match the approved contract -> only then change behavior if the contract/ADR is updated in the same PR.
4. Test pyramid is mandatory:
   - Unit: `domain` + `application/usecases` with fakes only.
   - Integration/component: repos, migrations, Redis Streams, model routing, queues using `testcontainers-go`.
   - E2E: few critical flows covering `POST /messages` -> queue -> worker -> Redis Streams -> SSE/WS.
5. Local regression must have one entrypoint: `make test-regression` = unit + integration + E2E.
6. CI guardrail: if `internal/domain/**` or `internal/application/**` changes, the PR must also change either contract artifacts (`docs/**` or `api/**`) or tests (`internal/**/*_test.go` or `test/**`).
7. CI guardrail: if core behavior files and tests change together, the PR must also change contract artifacts (`docs/**` or `api/**`); tests cannot redefine product behavior alone.
8. Load tests live under `test/load/` and cover ACK latency, stream latency, error rate, and end-to-end completion time.
9. Container-based tests use `testcontainers-go` for Postgres, Redis, and LocalStack; prefer suite-level reuse and always cleanup with context timeouts.
10. Run flaky-prone integration/E2E suites with `-count=1`; avoid hidden shared state.
11. Tests must preserve project invariants: `{{DOMAIN_ACTOR}}_id`, `request_id`, `conversation_id`, `seq`, async `202`, SQS FIFO grouping/dedup/DLQ, Redis Streams replay, no secret/PII logging.
12. Folder conventions are fixed:
   - Unit: near source as `*_test.go` in `internal/**`
   - Integration: `test/integration/**`
   - E2E: `test/e2e/**`
   - Load: `test/load/**`

## Do / Don't
- **Do** add or update the smallest test layer that proves the intended contract.
- **Do** state plainly when a test change is only reflecting existing business rules versus proposing new behavior.
- **Do** validate {{DOMAIN_ACTOR}} isolation in tests that touch persistence, queues, caches, or streams.
- **Do** keep container fixtures deterministic: fixed schemas, migrations, seed data, explicit cleanup.
- **Do** prefer fake agent engines in E2E unless the scenario is explicitly about ADK/provider integration.
- **Don't** rewrite business logic just to satisfy an outdated or overly strict test.
- **Don't** turn every change into E2E; keep most assertions in unit or integration layers.
- **Don't** use real cloud dependencies in CI for routine integration coverage; use `testcontainers-go` + LocalStack.
- **Don't** assert on sensitive logs, prompts, or secrets.
- **Don't** make POST handlers wait for worker completion; test `202` ACK, not synchronous completion.

## Interfaces / Contracts
- Read [test_pyramid.md](resources/test_pyramid.md) when deciding which layer to add.
- Read [guardrails_feature_over_tests.md](resources/guardrails_feature_over_tests.md) when a red test suggests changing behavior.
- Read [testcontainers_patterns.md](resources/testcontainers_patterns.md) for container fixture rules and infra scopes.
- Read [load_testing_patterns.md](resources/load_testing_patterns.md) for `k6`/`ghz` scenarios and SLO-style thresholds.
- Read [make_targets.md](resources/make_targets.md) and [ci_jobs.md](resources/ci_jobs.md) when touching automation.
- CI enforcement lives in `scripts/ci/guardrails.sh`; keep the script, workflow, docs, and this skill aligned.
- Canonical contracts under test:
  - HTTP: `POST /v1/conversations/{conversation_id}/messages` always returns `202 Accepted` + `request_id`.
  - Queue: `MessageGroupId = {{DOMAIN_ACTOR}}_id:conversation_id`, content-based dedup, DLQ `maxReceiveCount=3`.
  - Stream: Redis Streams carry monotonic `seq`; SSE/WS replay with `after` cursor.
  - Tenant safety: data, cache, stream, and queue assertions must scope by `{{DOMAIN_ACTOR}}_id`.

## Checklists
**Before**
- [ ] Confirm whether the change is behavior-preserving or behavior-changing.
- [ ] If core behavior will change, identify which contract artifact (`docs/**` or `api/**`) must change in the same PR.
- [ ] Map the change to the minimum required test layer using the pyramid.
- [ ] If infra is involved, decide which `testcontainers-go` fixtures are needed.
- [ ] If load-sensitive, define which latency/error metrics matter before coding.

**During**
- [ ] Keep unit tests near the business code and avoid infra there.
- [ ] Do not let tests become the only artifact describing a behavior change in `domain` or `application`.
- [ ] Use `context.WithTimeout` in container tests and explicit teardown paths.
- [ ] Run migrations inside Postgres integration tests; do not fake schema behavior.
- [ ] Validate `{{DOMAIN_ACTOR}}_id`, `request_id`, `conversation_id`, and `seq` where the flow crosses boundaries.
- [ ] Mark any behavior-changing test update in the PR/commit notes.

**After**
- [ ] Run `make test-regression`.
- [ ] Run `make ci-guardrails` or confirm the CI guardrail step will evaluate the same diff.
- [ ] Run the relevant load target when latency, throughput, or stream fan-out changed.
- [ ] Ensure CI jobs mirror the local regression entrypoint.
- [ ] Confirm no tests require logging secrets, prompts, or tenant data to pass.
- [ ] Update resources/automation if new test targets or directories were introduced.

## Definition of Done
- Business behavior matches the approved contract/ADR; no silent semantic drift introduced by tests.
- Any core behavior change is accompanied by updated contract artifacts or explicit confirmation that tests reflect the existing contract.
- `make test-regression` covers unit + integration + E2E and is green locally or in CI.
- Container tests validate the touched infra boundaries with `testcontainers-go`.
- Critical async path still proves `202` ACK, FIFO queueing, worker processing, Redis Streams replay, and SSE/WS delivery where applicable.
- Load scripts exist or are updated for changed hot paths, with recorded thresholds for ACK/error/end-to-end behavior.
- Tests remain tenant-safe, deterministic, and free of sensitive-data assertions.

## Minimal Examples
Table-driven unit test:
```go
func TestClassifier_Classify(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name string
		in   Ticket
		want Priority
	}{
		{name: "vip escalates", in: Ticket{{{DOMAIN_ACTOR_TITLE}}ID: "s1", VIP: true}, want: PriorityHigh},
		{name: "default normal", in: Ticket{{{DOMAIN_ACTOR_TITLE}}ID: "s1"}, want: PriorityNormal},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := Classifier{}.Classify(tt.in)
			if got != tt.want {
				t.Fatalf("got %v want %v", got, tt.want)
			}
		})
	}
}
```

Integration with `testcontainers-go` + Postgres + migrations:
```go
ctx, cancel := context.WithTimeout(context.Background(), 2*time.Minute)
defer cancel()

pg := postgres.RunContainer(ctx, testcontainers.WithImage("postgres:16-alpine"))
t.Cleanup(func() { _ = testcontainers.TerminateContainer(pg) })

dsn := pg.ConnectionString(ctx, "sslmode=disable")
require.NoError(t, migrateUp(ctx, dsn))
repo := postgresrepo.New(dsn)
```

Redis Streams append + read-after:
```go
cursor, err := stream.Publish(ctx, "{{DOMAIN_ACTOR}}-1", "conv-9", Event{Seq: 7, Type: "agent.delta"})
require.NoError(t, err)

events, next, err := stream.Read(ctx, "{{DOMAIN_ACTOR}}-1", "conv-9", cursor, 10)
require.NoError(t, err)
require.Len(t, events, 1)
require.Equal(t, int64(7), events[0].Seq)
_ = next
```

SQS FIFO in LocalStack with receive + DLQ path:
```go
stack := localstack.Run(ctx)
client := sqsClientFor(stack)

sendFIFO(t, client, queueURL, Message{
	{{DOMAIN_ACTOR_TITLE}}ID: "{{DOMAIN_ACTOR}}-1", ConversationID: "conv-9", RequestID: "req-1",
})

msg := receiveOne(t, client, queueURL)
require.Equal(t, "{{DOMAIN_ACTOR}}-1:conv-9", aws.ToString(msg.Attributes["MessageGroupId"]))

forceRetriesToDLQ(t, client, queueURL, dlqURL, msg, 3)
assertEventuallyInDLQ(t, client, dlqURL)
```

E2E happy path with fake engine:
```go
fakeEngine.Emit(
	Event{Type: "agent.delta", Seq: 1},
	Event{Type: "agent.completed", Seq: 2},
)

resp := postMessage(t, apiURL, "{{DOMAIN_ACTOR}}-1", "conv-9")
require.Equal(t, http.StatusAccepted, resp.StatusCode)

events := readSSE(t, streamURL, "{{DOMAIN_ACTOR}}-1", "conv-9", "")
require.Equal(t, "agent.delta", events[0].Type)
require.Equal(t, "agent.completed", events[1].Type)
```
</file>

<file path="skills/20-testing-strategy-regression-load-containers/resources/ci_jobs.md">
# CI Jobs

## Required Pre-merge Jobs
1. `guardrails`
   - Runs `bash scripts/ci/guardrails.sh`
   - Fails when core behavior changes are not accompanied by tests and/or contract artifacts according to the repo rules.
1. `unit`
   - Runs `make test-unit`
   - Fails fast on business-rule regressions.
2. `integration`
   - Runs `make test-integration`
   - Requires Docker and validates Postgres, Redis, LocalStack paths via `testcontainers-go`.
3. `e2e`
   - Runs `make test-e2e`
   - Small suite proving `202`, queue handoff, worker output, stream replay.
4. `regression`
   - Runs `make test-regression`
   - Optional aggregator job, required before merge if separate jobs are not mandatory checks.

## Recommended Performance Jobs
- `load-smoke`
  - Trigger on PR label or nightly.
  - Runs short `make load-api` and `make load-e2e`.
- `load-nightly`
  - Runs broader `make load-stream-sse`, `make load-stream-ws`, `make load-e2e`.
  - Publishes trend artifacts for ACK p95/p99, error rate, end-to-end duration.

## CI Rules
- Do not require external shared cloud infra for routine regression jobs.
- Run the guardrail script before tests so invalid PR shapes fail early.
- Keep E2E job small; prefer deterministic fake agent engine.
- Persist artifacts when jobs fail: logs, `k6` summaries, container startup output, flaky retry evidence.
- If business behavior changes, CI must fail until tests and contract docs are updated together.
</file>

<file path="skills/20-testing-strategy-regression-load-containers/resources/guardrails_feature_over_tests.md">
# Guardrails: Feature Over Tests

## Decision Rule
- Contract, ADR, and approved business behavior come first.
- Tests verify behavior; they do not create new behavior by themselves.

## When a Test Fails
1. Ask: is the test asserting documented behavior or inventing a new rule?
2. Validate the contract artifact first: `docs/**`, `api/**`, event schema, OpenAPI/proto, ADR.
3. If documented behavior is correct, fix code, fixture, or the test to reflect that contract and keep the contract intact.
4. If the test proposes new behavior, stop and call it out explicitly in the PR/issue before changing production logic.
5. Update code, tests, and contract docs together when the business rule truly changes.

## CI Enforcement
- Guardrail CI #1:
  - If `internal/domain/**` or `internal/application/**` changed, require at least one change in `docs/**`, `api/**`, `internal/**/*_test.go`, or `test/**`.
- Guardrail CI #2:
  - If core files and tests changed together, also require a change in `docs/**` or `api/**`.
- Implementation entrypoint:
  - `bash scripts/ci/guardrails.sh`

## Required Agent Language
- Use wording like:
  - `This test failure indicates a behavior mismatch with the approved contract.`
  - `This test update would change business behavior; contract/ADR approval is required before modifying production code.`
  - `I am updating the test to reflect existing behavior, not introducing a new rule.`

## Anti-patterns
- Changing queue/stream/tenant semantics only because a brittle test assumed the wrong behavior.
- Making handlers synchronous so an E2E test can assert completion in one HTTP response.
- Relaxing constraints or idempotency rules to avoid fixing fixtures or migrations.
</file>

<file path="skills/20-testing-strategy-regression-load-containers/resources/load_testing_patterns.md">
# Load Testing Patterns

## Tools
- Preferred: `k6` for HTTP, SSE, WebSocket, and full pipeline scenarios.
- Optional: `ghz` for gRPC-specific endpoints if the repo adds critical gRPC hot paths.

## Required Scenarios
- `make load-api`
  - Focus: `POST /messages` ACK only.
  - Validate API returns `202` quickly without waiting for worker completion.
- `make load-stream-sse`
  - Focus: SSE subscribe, reconnect, replay via `after`, event fan-out.
- `make load-stream-ws`
  - Focus: WebSocket subscribe, steady-state delivery, reconnect.
- `make load-e2e`
  - Focus: POST -> queue -> worker -> Redis Streams -> client receives `agent.completed`.

## Minimum Metrics
- ACK latency: p95 and p99.
- Stream delivery latency: first delta and completion.
- End-to-end completion time.
- Error rate.
- Throughput: requests/sec or conversations/min.

## Suggested Thresholds
- ACK `POST /messages`
  - p95 <= 300 ms
  - p99 <= 750 ms
  - error rate < 1%
- SSE / WS first event
  - p95 <= 1 s after worker emits
- End-to-end completion
  - define per use case; start with p95 <= 10 s for fake-engine local runs

## Guardrails
- Use {{DOMAIN_ACTOR}}-scoped ids in load data.
- Do not disable auth/tenant checks just to reach higher throughput.
- Keep fake agent engines deterministic unless the goal is provider benchmarking.
- Store summaries as CI artifacts so regressions are comparable over time.
</file>

<file path="skills/20-testing-strategy-regression-load-containers/resources/make_targets.md">
# Make Targets

| Target | Purpose | Expected command shape |
|---|---|---|
| `make test` | fast default developer loop | `go test ./...` or repo equivalent |
| `make test-unit` | unit-only feedback | `go test ./internal/...` |
| `make test-integration` | Postgres/Redis/LocalStack via `testcontainers-go` | `go test -count=1 ./test/integration/...` |
| `make test-e2e` | critical async paths only | `go test -count=1 ./test/e2e/...` |
| `make test-regression` | one-liner pre-merge gate | `make test-unit && make test-integration && make test-e2e` |
| `make load-api` | ACK `202` latency and error rate | run `k6` script under `test/load/api` |
| `make load-stream-sse` | SSE fan-out/replay load | run `k6` script under `test/load/stream_sse` |
| `make load-stream-ws` | WebSocket stream load | run `k6` script under `test/load/stream_ws` |
| `make load-e2e` | full pipeline throughput | run `k6` scenario covering POST -> worker -> stream completion |

## Notes
- Keep `make test-regression` stable; extend internals, do not rename casually.
- Integration and E2E targets should set `-count=1` by default.
- Load targets should accept overridable env vars like `BASE_URL`, `SELLER_ID`, `VUS`, `DURATION`.
- If a folder does not exist yet, create it before adding the target:
  - `test/integration/`
  - `test/e2e/`
  - `test/load/`
</file>

<file path="skills/20-testing-strategy-regression-load-containers/resources/test_pyramid.md">
# Test Pyramid

## Layer Rules

| Layer | Scope | Infra | Typical assertions | Location |
|---|---|---|---|---|
| Unit | `domain`, `application`, validators, pure mappers | none; use fakes/stubs | business rules, errors, branching, {{DOMAIN_ACTOR}} isolation logic | `internal/**/**/*_test.go` |
| Integration / Component | repos, migrations, Redis Streams, {{DOMAIN_ACTOR}} model routing, queue publishers/consumers | `testcontainers-go` for Postgres, Redis, LocalStack | schema constraints, `jsonb`, inbox idempotency, TTL, stream replay, FIFO attributes | `test/integration/**` |
| E2E | critical async journeys only | local app stack + fake agent engine + real infra containers | `202` ACK, worker handoff, emitted events, replay/resume, correlation ids | `test/e2e/**` |
| Load | hot-path performance and availability | local stack or ephemeral env | p95/p99 ACK, error %, end-to-end completion, stream fan-out latency | `test/load/**` |

## Selection Heuristics
- Prefer unit tests when only one use case or domain rule changes.
- Prefer integration when correctness depends on SQL, migrations, Redis behavior, queue semantics, or serialization.
- Prefer E2E only for cross-boundary contracts the business depends on: ACK, worker pipeline, stream delivery, replay.
- Prefer load tests when touching request throughput, stream fan-out, worker concurrency, or storage latency.

## Required Coverage by Subsystem
- `domain` / `application`: unit first; add integration only if persistence or queue contracts changed.
- Postgres repositories: integration mandatory for migrations, constraints, `jsonb`, inbox, `{{DOMAIN_ACTOR}}_model_config`.
- Redis streams / result store / rate limit: integration mandatory for replay, TTL, and isolation by `{{DOMAIN_ACTOR}}_id`.
- SQS FIFO / worker loops: integration mandatory with LocalStack; E2E for the business-critical path only.
- HTTP `POST /messages`: E2E mandatory to prove `202` and decoupling from worker completion.

## Anti-patterns
- Replacing missing integration coverage with mocks around SQL, Redis streams, or SQS semantics.
- Adding broad E2E suites for simple use-case branching.
- Keeping flaky cross-test shared state instead of resetting fixtures or using unique ids.
</file>

<file path="skills/20-testing-strategy-regression-load-containers/resources/testcontainers_patterns.md">
# Testcontainers Patterns

## Containers Required by Concern
- Postgres
  - Validate migrations, `jsonb` columns, constraints, inbox idempotency, `{{DOMAIN_ACTOR}}_model_config`.
- Redis
  - Validate Streams replay, result-store TTL, rate-limit counters, {{DOMAIN_ACTOR}}-scoped keys.
- LocalStack
  - Validate SQS FIFO ordering/grouping, content-based dedup, retry/DLQ with `maxReceiveCount=3`.

## Fixture Rules
- Create `context.WithTimeout` for every suite and every container startup.
- Prefer one container set per suite/package when tests can isolate by schema/key prefixes/ids.
- Use `t.Cleanup` or suite teardown to terminate containers.
- Run migrations against the actual test Postgres container; never bypass them with hand-created tables.
- Use unique `{{DOMAIN_ACTOR}}_id`, `conversation_id`, `request_id` values per test to prevent state bleed.
- Set `go test -count=1` for integration/E2E targets.

## Suggested Patterns
- Postgres
  - Start container -> build DSN -> run migrations -> construct repo -> seed fixtures.
- Redis
  - Start container -> create isolated client/db -> publish stream events -> read with `after` cursor.
- LocalStack
  - Start container -> create FIFO queue + DLQ + redrive policy -> publish -> receive -> simulate retry path.

## Failure Diagnostics
- On timeout, dump container logs and connection settings, not secrets.
- Assert on canonical fields only: `{{DOMAIN_ACTOR}}_id`, `request_id`, `conversation_id`, `seq`, queue attributes, Redis ids.
- Avoid sleeps when possible; use `require.Eventually` with bounded waits.
</file>

<file path="skills/21-documentation-open-source/SKILL.md">
---
name: documentation-open-source
description: Godoc for pkg/* packages, README maintenance, CONTRIBUTING guide, CHANGELOG updates, example READMEs, and spec-to-doc traceability. Use when creating or changing public packages, examples, or any user-facing documentation artifact.
---

# Open Source Documentation
Goal: ensure every public surface of {{PROJECT_SLUG}} has traceable, idiomatic, and up-to-date documentation.

## When to Use
- Creating or modifying a package under `pkg/*` (godoc required).
- Changing public API signatures, types, or observable behavior (README, examples, and spec must follow).
- Adding a new example under `examples/` (example README and smoke test required).
- Preparing a release or recording a behavior change (CHANGELOG entry required).
- Onboarding contributors or accepting first external PR (CONTRIBUTING.md must be current).

## Non-negotiables
1. Every package in `pkg/*` must have a `doc.go` with valid godoc following Go conventions.
2. Godoc opening sentence must start with the declared name (`Package X ...`, `Type Y ...`, `Func Z ...`).
3. Root `README.md` must contain: project description, installation, quick-start snippet, link to pkg.go.dev, and license.
4. Every example under `examples/` must have its own `README.md` with purpose, how to run, and expected output.
5. Public API changes require documentation updates in the same PR as the code change.
6. `CONTRIBUTING.md` must exist and cover: local setup, commit conventions, PR workflow, and how to run tests.
7. `CHANGELOG.md` must be updated for release-worthy changes (integrated with git-cliff per skill 15).
8. Every documentation artifact must be traceable to at least one spec.

## Do / Don't
- **Do** update docs in the same PR as the code change.
- **Do** use a dedicated `doc.go` file for packages with long documentation.
- **Do** validate markdown links and formatting before merge.
- **Do** keep examples compilable and runnable (`go run` or `go test`).
- **Do** reference `pkg.go.dev` from README instead of duplicating API details.
- **Don't** duplicate content between README and godoc; README links, godoc details the API.
- **Don't** document `internal/*` types or behavior in public-facing docs.
- **Don't** create documentation that is not traceable to a spec.
- **Don't** leave an example without a README or without execution instructions.

## Interfaces / Contracts
- `doc.go` template and conventions: [doc_go_template.md](resources/doc_go_template.md).
- Root and sub-module README checklist: [readme_checklist.md](resources/readme_checklist.md).
- Example README template: [example_readme_template.md](resources/example_readme_template.md).
- Expected file locations: `pkg/*/doc.go`, `README.md`, `CONTRIBUTING.md`, `CHANGELOG.md`, `examples/*/README.md`.

## Checklists
**Before**
- [ ] Identify which documentation artifacts are affected by the change.
- [ ] Check if the affected package already has a `doc.go`.
- [ ] Confirm which spec governs the change for traceability.

**During**
- [ ] Write or update `doc.go` for every touched `pkg/*` package.
- [ ] Update root `README.md` if public API surface changed.
- [ ] Write or update example `README.md` for affected examples.
- [ ] Add `CHANGELOG.md` entry if the change is release-worthy.

**After**
- [ ] Validate markdown links (no broken references).
- [ ] Confirm all affected examples compile and run.
- [ ] Verify spec traceability is explicit in PR description.
- [ ] Run `go doc` locally to confirm godoc renders correctly.

## Definition of Done
- All touched `pkg/*` packages have valid `doc.go` with proper opening sentences.
- Root `README.md` reflects current public API if it changed.
- Affected examples have `README.md` and compile successfully.
- `CHANGELOG.md` updated for release-worthy changes.
- `CONTRIBUTING.md` exists and is current.
- No public documentation references `internal/*` types directly.

## Minimal Examples
- Well-structured `doc.go` (reference model: `pkg/memory/doc.go`):
  ```go
  // Package tool defines the contract for executable tools and the
  // helpers used to compose toolkits in {{PROJECT_SLUG}}.
  //
  // # Defining a tool
  //
  // Implement the [Tool] interface and register it with [Registry.Register].
  //
  // # Composing toolkits
  //
  // Use [NewToolkit] to group related tools under a shared namespace.
  package tool
  ```
- Example README (`examples/basic-agent/README.md`):
  ```markdown
  # Basic Agent

  Demonstrates creating and running an agent with a single tool.

  ## How to run

      go run ./examples/basic-agent

  ## Expected output

      Agent replied: Hello from {{PROJECT_SLUG}}!

  ## Concepts

  - Agent creation via `pkg/agent`
  - Tool registration via `pkg/tool`
  ```
- CHANGELOG entry:
  ```markdown
  ## [Unreleased]
  ### Added
  - `pkg/memory`: FileStore with atomic writes and restart persistence.
  ```
</file>

<file path="skills/21-documentation-open-source/resources/doc_go_template.md">
# doc.go Template and Conventions

## Purpose

Every package under `pkg/*` must have a `doc.go` file that provides godoc-compatible documentation. This file contains only a package-level comment and the `package` clause.

## Structure

```go
// Package <name> <one-sentence summary starting with a verb>.
//
// <Expanded description: what the package provides, its role in {{PROJECT_SLUG}},
// and key abstractions.>
//
// # <Section heading>
//
// <Details about a specific topic: choosing implementations, configuring
// behavior, extending via interfaces, etc.>
//
//   - [TypeOrFunc] short description of the first option.
//
//   - [TypeOrFunc] short description of the second option.
//
// # <Another section heading>
//
// <Guidance for users who want to extend or customize.>
package <name>
```

## Rules

1. **Opening sentence**: must start with `Package <name>` followed by a verb phrase that summarizes the package purpose. This sentence appears in pkg.go.dev package listings.
2. **No blank line** between the comment block and the `package` clause.
3. **Section headings** use `# Heading` (single `#`) inside the comment to create godoc sections (Go 1.19+).
4. **Symbol references** use `[SymbolName]` to create hyperlinks in pkg.go.dev.
5. **Lists** use `//   - ` (three spaces + dash) for indented bullet points rendered by godoc.
6. **Code examples** inside doc comments use indentation (tab or four spaces after `//`).
7. **Keep it focused**: `doc.go` documents the package contract, not implementation details from `internal/*`.
8. **Traceability**: mention the governing spec when the package implements a specific spec requirement (e.g., "as defined in spec 030").

## Reference Model

The canonical example in this repository is `pkg/memory/doc.go`, which demonstrates:

- Opening sentence naming the package
- Contextual description separating conversational memory from working memory
- `# Choosing a backend` section with two listed implementations
- `# Implementing a custom adapter` section for extensibility guidance
- Symbol references via `[Store]`, `[InMemoryStore]`, `[FileStore]`, `[NewFileStore]`

## When NOT to Use doc.go

- Packages with only one or two exported symbols may document directly above those declarations instead of in a separate `doc.go`.
- `internal/*` packages do not require `doc.go` (but inline comments on exported-within-internal symbols are still encouraged).
</file>

<file path="skills/21-documentation-open-source/resources/example_readme_template.md">
# Example README Template

Use this template when creating a `README.md` for any example under `examples/`.

---

## Template

```markdown
# <Example Name>

<One sentence explaining what this example demonstrates.>

## Prerequisites

- Go 1.22+ installed
- <Any other prerequisites, or remove this section if none>

## How to run

    go run ./examples/<example-name>

## Expected output

    <Paste the exact terminal output the user should see.>

## Concepts

- <Concept 1>: brief explanation, link to `pkg/<package>` or spec.
- <Concept 2>: brief explanation, link to `pkg/<package>` or spec.
```

---

## Rules

1. **Title** must match the directory name in human-readable form (e.g., `basic-agent` becomes `Basic Agent`).
2. **Purpose sentence** must be concrete: "Demonstrates X using Y", not "Shows how things work".
3. **How to run** must use the exact `go run` command from the repository root.
4. **Expected output** must reflect what a clean run produces. If output varies (e.g., timestamps), use placeholders like `<timestamp>`.
5. **Concepts** must link to the relevant `pkg/*` package or spec so the reader can dive deeper.
6. Do not include implementation details or internal package references.
7. Keep the README short; the example code is the primary documentation.

## Reference

See `examples/demo-app/README.md` for an existing example in the repository.
</file>

<file path="skills/21-documentation-open-source/resources/readme_checklist.md">
# README Checklist

Use this checklist when creating or reviewing the root `README.md` or any sub-module README.

## Root README.md

### Required sections

- [ ] **Project name and description**: one paragraph explaining what {{PROJECT_SLUG}} is and its relationship to {{UPSTREAM_NAME}}.
- [ ] **Installation**: `go get` command with the module path.
- [ ] **Quick start**: minimal compilable code snippet showing the most common use case (e.g., creating and running an agent).
- [ ] **Documentation link**: link to pkg.go.dev for full API reference.
- [ ] **Contributing**: link to `CONTRIBUTING.md`.
- [ ] **License**: license type and link to `LICENSE` file.

### Recommended sections

- [ ] **Badges**: CI status, Go version, license badge, Go Report Card.
- [ ] **Roadmap**: link to specs or high-level milestones.
- [ ] **Non-scope**: brief mention of what {{PROJECT_SLUG}} does NOT cover ({{UPSTREAM_OPS_NAME}}) with link to `specs/001-non-goals.md`.
- [ ] **Examples**: list of examples under `examples/` with one-line descriptions.

### Rules

1. Quick-start code must compile without importing `internal/*`.
2. Do not duplicate godoc content; link to pkg.go.dev instead.
3. Keep the README concise; detailed API documentation belongs in `doc.go` files.
4. Update the README in the same PR when public API changes affect user-facing instructions.
5. Badge URLs must point to stable endpoints (GitHub Actions, pkg.go.dev, goreportcard).

## Example README (examples/*/README.md)

See [example_readme_template.md](example_readme_template.md) for the full template.

### Minimum content

- [ ] **Title**: name of the example.
- [ ] **Purpose**: one sentence explaining what the example demonstrates.
- [ ] **How to run**: exact command (`go run ./examples/<name>`).
- [ ] **Expected output**: what the user should see in the terminal.
- [ ] **Concepts**: list of {{PROJECT_SLUG}} concepts demonstrated, with links to relevant `pkg/*` godoc or specs.
</file>

<file path="skills/22-release-versioning-governance/SKILL.md">
---
name: release-versioning-governance
description: Govern versioning and release decisions for {{PROJECT_SLUG}}. Use when an agent needs to decide whether a change warrants a new release, classify patch/minor/major impact, detect breaking changes in the public surface, prepare release checklist, changelog and release notes, and recommend the correct distribution path such as no release, prerelease, stable tag or hotfix.
---

# Release Versioning Governance
Goal: make release, versioning and distribution decisions for `{{PROJECT_SLUG}}` consistently and transparently, with explicit SemVer reasoning and hard gates before any stable release.

## When to Use
- Evaluating whether a merged or proposed change should produce a new release.
- Classifying a change as `patch`, `minor`, `major`, `prerelease`, `hotfix`, or `no release`.
- Preparing release checklist, changelog, release notes, distribution recommendation, or compatibility assessment.
- Auditing whether a candidate release is safe to publish without compromising stable consumers.

## Non-negotiables
1. Use SemVer as the public release contract.
2. Treat `pkg/*`, documented HTTP contracts, and published OpenAPI as public surface.
3. Treat `internal/*` as non-public unless surfaced through `pkg/*`, HTTP, OpenAPI, starter flow, or documented behavior.
4. Never recommend mutating an already published release; publish a new version instead.
5. Stable release requires passing release gates; if evidence is incomplete, recommend `NOT READY` or `prerelease`.
6. Do not classify by intent alone; classify by observable impact on consumers.
7. For `{{PROJECT_SLUG}}`, phase reports and parity diagnostics are release evidence, not optional context.

## Public Surface Rules
Use these rules before classifying impact:
- **Public API:** exported contracts in `pkg/*`, OpenAPI-described endpoints, starter-facing documented setup flow, documented config and CLI/demo usage that the repo treats as supported.
- **Non-public API:** `internal/*`, test helpers, temporary build seams, comments-only changes, reports, and internal file moves with no public behavior impact.
- **Potentially public behavior:** example/starter flows, `/docs`, `/openapi.yaml`, documented response envelopes, SSE event shapes, trace/read surfaces if documented as part of the framework.

## Release Classification Rules
Read [release_policy.md](resources/release_policy.md) and [version_decision_tree.md](resources/version_decision_tree.md) before deciding.

### `no release`
Choose when:
- only docs/reports/spec reconciliation changed
- comments or internal-only refactors changed
- tests changed without public behavior change
- examples changed without affecting documented supported flow

### `patch`
Choose when:
- bug fix preserves public API and documented behavior
- validation, error mapping, docs surface, or runtime defect is corrected compatibly
- OpenAPI/doc bundle is corrected to match already-existing behavior

### `minor`
Choose when:
- new backward-compatible capability is added
- new optional endpoint, config, trace surface, starter path, or documented public contract is added compatibly
- existing behavior is extended without breaking prior consumers

### `major`
Choose when:
- exported symbol removal or signature change in `pkg/*`
- incompatible HTTP or OpenAPI contract change
- renamed/removed documented endpoint
- incompatible behavior change in starter or canonical setup flow
- changed semantics that force consumer code changes even if symbols remain

### `prerelease`
Choose when:
- feature is real but baseline is not yet stable enough for general stable release
- diagnostics say capability is present but not yet fully closed
- API exists but you want consumer trial without stability promise
- recommended tags: `vX.Y.Z-alpha.N`, `vX.Y.Z-beta.N`, `vX.Y.Z-rc.N`
- `alpha.N`: early technical feedback or unstable shape
- `beta.N`: mostly complete candidate for controlled integration with known gaps documented
- `rc.N`: stable-release candidate where only localized fixes are expected before final tag

Prerelease stages describe candidate maturity only. They do not hide breaking changes, missing gates, or known incompatibilities.

### `hotfix`
Choose when:
- urgent fix targets latest stable line
- change is minimal and clearly `patch`
- risk of waiting for next normal release is unacceptable

## Required Inputs
When applying this skill, gather at least:
- change summary or diff summary
- touched files/modules
- whether public surface changed
- current latest released version or intended baseline
- relevant phase/spec/report evidence if the repo uses phase gates for this change

If any of the above is missing, say what is missing and classify provisionally.

## Required Outputs
Always produce all of the following sections:
1. **Release decision**
2. **Version bump recommendation**
3. **Why this is not lower/higher impact**
4. **Breaking change assessment**
5. **Release checklist status**
6. **Changelog draft**
7. **Release notes draft**
8. **Recommended distribution path**
9. **Residual risks / blockers**

Use the templates in [output_templates.md](resources/output_templates.md).

## Release Gates
Apply [release_checklist.md](resources/release_checklist.md) strictly.

At minimum, before recommending a stable release, verify:
- `go test ./...`
- `go test -race ./...`
- relevant smoke/conformance checks for changed surface
- OpenAPI/docs alignment if HTTP surface changed
- feature matrix/reports not contradicting the proposed release story
- no open blocker-level gap for the affected capability

If any gate is unknown, explicitly mark as `UNVERIFIED`.

## Distribution Recommendation Rules
Recommend exactly one path:
- `NO_RELEASE`
- `PRERELEASE_TAG`
- `STABLE_TAG`
- `HOTFIX_TAG`
- `HOLD_RELEASE`

Map them as follows:
- docs/spec/report only => usually `NO_RELEASE`
- compatible completed feature with green gates => `STABLE_TAG`
- incomplete but useful capability => `PRERELEASE_TAG`
- urgent patch on stable line => `HOTFIX_TAG`
- blockers or unverified gates => `HOLD_RELEASE`

## {{PROJECT_SLUG}} Specific Rules
1. Use phase diagnostics as evidence for release maturity.
2. If a phase/report says `NOT READY`, do not recommend stable release for that track.
3. If a phase closes a capability but known S2/S3 gaps remain, reflect them in risk and release notes.
4. If change only reconciles specs/docs without new consumer behavior, prefer `NO_RELEASE`.
5. If OpenAPI, `/docs`, starter, or public traces change materially, treat them as public-facing changes.
6. For `v0.x.y`, breaking changes are still breaking changes; do not hide them just because major is zero.
7. Recommend `v1.0.0` only when public API/starter/docs/release policy are intentionally declared stable.

## Checklists
**Before deciding**
- [ ] Identify touched public surface.
- [ ] Identify current and target version line.
- [ ] Identify whether change is behavior, docs-only, or compatibility-affecting.
- [ ] Load release policy, decision tree, checklist, and templates.

**Before stable recommendation**
- [ ] Core tests green.
- [ ] Race tests green.
- [ ] Changed capability diagnostics not blocked.
- [ ] Changelog and release notes drafted.
- [ ] Distribution path chosen.

## Definition of Done
A release recommendation is complete only when:
- SemVer classification is explicit.
- Public-surface impact is explained.
- Checklist status is explicit (`PASS`, `FAIL`, `UNVERIFIED`).
- Changelog and release notes are drafted.
- Distribution path is explicit.
- Residual blockers are named honestly.

## Minimal Examples
- Docs/spec reconciliation only -> `NO_RELEASE`, maybe roll into next minor release notes.
- New backward-compatible HTTP endpoint -> `minor`, usually `STABLE_TAG` if gates are green.
- Renamed exported `pkg/*` symbol -> `major` or defer to next major line.
- Urgent bug fix in stable persistence path -> `patch`, `HOTFIX_TAG` if urgent.
</file>

<file path="skills/22-release-versioning-governance/resources/output_templates.md">
# Output Templates

## 1. Release Decision
- Decision: `NO_RELEASE | PRERELEASE_TAG | STABLE_TAG | HOTFIX_TAG | HOLD_RELEASE`
- Recommended version: `vX.Y.Z` or `vX.Y.Z-alpha.N | vX.Y.Z-beta.N | vX.Y.Z-rc.N`
- Classification: `patch | minor | major | no-release`
- Prerelease stage: `N/A | alpha.N | beta.N | rc.N`
- Prerelease stage rationale:

## 2. Breaking Change Assessment
- Public surface touched:
- Breaking?: `yes | no | possible`
- Why:

## 3. Checklist Status
| Check | Status | Notes |
| --- | --- | --- |
| go test ./... | PASS/FAIL/UNVERIFIED | |
| go test -race ./... | PASS/FAIL/UNVERIFIED | |
| Relevant smoke/conformance | PASS/FAIL/UNVERIFIED | |
| OpenAPI/docs alignment | PASS/FAIL/UNVERIFIED | |
| Prerelease stage justified | PASS/FAIL/N/A | |

## 4. Changelog Draft
### Added
- ...
### Changed
- ...
### Fixed
- ...
### Deprecated
- ...
### Removed
- ...

## 5. Release Notes Draft
### Summary
...

### Prerelease Stage
...

### Consumer Impact
...

### Upgrade Notes
...

### Risks
...

## 6. Distribution Recommendation
- Path:
- Reason:
- Suggested tag/branch action:
</file>

<file path="skills/22-release-versioning-governance/resources/release_checklist.md">
# Release Checklist

## Required checks
- [ ] Change summary is explicit.
- [ ] Touched files/modules are listed.
- [ ] Public surface impact assessed.
- [ ] Breaking-change assessment completed.
- [ ] Recommended version bump selected.
- [ ] Distribution path selected.
- [ ] Prerelease stage selected and justified when using `PRERELEASE_TAG`.
- [ ] Prerelease stage does not hide breaking change, missing gates, or known incompatibility.
- [ ] `go test ./...` passed or explicitly marked `UNVERIFIED`.
- [ ] `go test -race ./...` passed or explicitly marked `UNVERIFIED`.
- [ ] Relevant smoke/conformance checks passed or explicitly marked `UNVERIFIED`.
- [ ] OpenAPI/docs alignment checked if HTTP/public docs changed.
- [ ] Release notes drafted.
- [ ] Changelog drafted.

## Optional checks
- [ ] Prerelease justification names `alpha.N`, `beta.N`, or `rc.N`.
- [ ] Hotfix urgency justified.
- [ ] Known residual risks listed.
- [ ] Rollback/hold recommendation documented if needed.
</file>

<file path="skills/22-release-versioning-governance/resources/release_policy.md">
# Release Policy for {{PROJECT_SLUG}}

## 1. Versioning Model
Use SemVer:
- `MAJOR`: incompatible public change
- `MINOR`: backward-compatible feature
- `PATCH`: backward-compatible fix

For `v0.x.y`, still classify breakage honestly. The repo may remain in `v0`, but the release recommendation must still flag breaking changes.

## 2. Public Contract Boundaries
Treat as public:
- exported contracts in `pkg/*`
- documented HTTP endpoints
- OpenAPI schemas and envelopes
- starter / getting-started flow when documented as canonical
- documented config and trace surfaces

Treat as non-public unless surfaced indirectly:
- `internal/*`
- smoke/test helpers
- report-only artifacts
- implementation-only seams

## 3. Stable Release Conditions
Recommend stable release only when:
- no blocker gap is open for the changed capability
- validation commands passed or are explicitly confirmed
- docs and OpenAPI match real behavior if relevant
- release notes/changelog are consistent with actual change scope

## 4. Prerelease Conditions
Prefer prerelease when:
- capability is newly added but not yet hardened
- diagnostics are partial or recent
- external feedback is desired before stability promise

Use the stage to describe candidate maturity:
- `alpha.N`: early technical feedback, validation, or unstable shape. Expect API, payload, docs, or behavior refinements before stability.
- `beta.N`: mostly complete candidate for controlled integration. Core behavior is present, contracts are documented, and known gaps are listed.
- `rc.N`: stable-release candidate. No broad functional changes are expected before the final tag; only localized fixes or final documentation adjustments should remain.

Prerelease stages do not replace SemVer classification. A breaking change is still `MAJOR`, and missing gates or known incompatibilities must remain explicit in the release decision.

## 5. Hotfix Conditions
Use hotfix only when all are true:
- targets a released stable line
- urgent consumer-facing defect
- patch-scoped fix
- low blast radius

## 6. Release Immutability Rule
Never recommend changing a published version in place.
Always publish a new version.
</file>

<file path="skills/22-release-versioning-governance/resources/version_decision_tree.md">
# Version Decision Tree

## Step 1: Did public behavior change?
- No -> likely `NO_RELEASE` or at most roll into next scheduled release notes.
- Yes -> continue.

## Step 2: Is any consumer action required to keep working?
- Yes -> breaking change -> `major`.
- No -> continue.

## Step 3: Is this new capability or only a fix?
- New compatible capability -> `minor`.
- Compatible correction/fix -> `patch`.

## Step 4: Is release maturity incomplete?
- Yes -> convert recommendation to `PRERELEASE_TAG` and choose a stage:
  - `alpha.N` for early technical feedback or unstable shape.
  - `beta.N` for controlled integration with core behavior present and known gaps documented.
  - `rc.N` for a stable-release candidate where only localized fixes are expected before final tag.
- No -> continue.

## Step 5: Is urgency exceptional on stable line?
- Yes -> `HOTFIX_TAG`.
- No -> `STABLE_TAG`.

## Additional checks
- If only spec/report/README/feature-matrix changed without consumer-visible behavior: `NO_RELEASE`.
- If OpenAPI changed because actual runtime changed, classify from runtime impact, not from file type alone.
- If starter/getting-started path changed materially for new users, treat as public-surface impact.
- Prerelease stage never lowers SemVer impact; breaking changes remain breaking.
</file>

<file path="skills/23-{{PROJECT_SLUG}}-sdd-autopilot/SKILL.md">
---
name: {{PROJECT_SLUG}}-sdd-autopilot
description: Operate {{PROJECT_SLUG}}'s interactive SDD autopilot for feature requests, evolutions, bugs, refactors and operational documentation. Use when a user asks Cursor to take an initial request through requirements, dual-spec, implementation, diagnosis, report review, gates, state updates and completion without breaking the fixed roadmap autopilot.
---

# {{PROJECT_SLUG}} SDD Autopilot

Goal: conduct an interactive feature request through Spec Driven Development until completion or a stop condition, while preserving the existing `phase_autopilot`.

## When to Use

Use this skill when the user asks to:

- implement a feature or evolution from an initial request
- turn a vague request into specs and implementation
- continue an interactive SDD trail
- diagnose or unblock an interactive SDD trail
- run the `interactive_sdd_autopilot`

Do not use this skill for the fixed roadmap execution. Use `automation/RUNBOOK.md` for `phase_autopilot`.

## Required Reading

Before editing, read:

1. `AGENTS.md`
2. `skills/00-skill-index/SKILL.md`
3. `skills/05-{{PROJECT_SLUG}}-spec-architect/SKILL.md`
4. `specs/680-phase-25-interactive-sdd-autopilot-foundation.md`
5. `specs/681-phase-25-interactive-sdd-autopilot-foundation-diagnosis.md`
6. `automation/INTERACTIVE_AUTOPILOT.md`
7. `automation/INTERACTIVE_RUNBOOK.md`
8. `automation/INTERACTIVE_STATE.json`
9. `automation/STOP_CONDITIONS.md`

Then read the specs, skills and docs that govern the requested domain.

## Non-Negotiables

1. Treat the user request as intake, not as a spec.
2. Do not implement behavior outside approved specs.
3. Every new trail must have one build spec and one diagnosis spec.
4. If a material requirement is ambiguous, ask before editing.
5. The report is the source of truth for advancement.
6. Advance only when `classification = PASS` and the decision allows the next step.
7. Stop immediately on any stop condition.
8. Do not mutate `automation/ROADMAP.json` or `automation/PHASE_STATE.json` for interactive trails.
9. Do not create product capability, `pkg/*` API, runtime feature, {{UPSTREAM_OPS_NAME}} or hosted dependency unless governed by a separate approved product spec.

## Interactive Workflow

### 1. Intake

Normalize the request:

- original request
- request type: `feature`, `evolution`, `bug`, `refactor`, `docs` or `investigation`
- intended outcome
- known scope
- ambiguous points
- candidate specs
- candidate skills

If the request is materially incomplete, ask the smallest useful set of questions and stop until answered.

### 2. Spec Decision

Use `skills/05-{{PROJECT_SLUG}}-spec-architect/SKILL.md`.

Decide:

- amend existing specs when the request belongs to an existing module or contract
- create a new dual-spec trail when the request introduces a new bounded context, primary contract or diagnostic domain

Record the decision and the governing sections.

### 3. Dual-Spec Gate

Before implementation, ensure:

- build spec defines objective, scope, non-goals, contracts or observable behavior, tests and acceptance criteria
- diagnosis spec defines signals, failure modes, commands, evidence, classification, decision and report path

If either spec is missing or incomplete, create or refine it first.

### 4. Implementation

Implement only the scoped behavior covered by the specs.

Rules:

- preserve `pkg/` and `internal/` boundaries from `Spec 020`
- update tests when behavior changes
- update docs when public or operational behavior changes
- do not change fixed roadmap files to make the interactive trail pass
- keep evidence traceable from request to specs, code, tests, docs and report

### 5. Diagnosis

Run the diagnosis spec.

Minimum checks when applicable:

- targeted tests for changed packages
- `go test ./...`
- `go test -race ./...`
- OpenAPI check when HTTP/OpenAPI changes
- documentation and rules checks when governance changes
- verification that no forbidden scope entered

If a required check cannot run, record the technical reason in the report.

### 6. Report Review

Read the generated report completely.

Extract:

- `classification`
- `decision`
- tests executed
- skipped checks
- gaps
- stop conditions

Never advance based only on tests, narrative or confidence.

### 7. Advance, Retry, Block or Complete

Advance only when:

- report exists and matches the current trail step
- `classification = PASS`
- decision explicitly allows advancement
- required checks passed or were explicitly not applicable
- no stop condition is active
- `automation/INTERACTIVE_STATE.json` is coherent

Block when the report is `PARTIAL`, `FAIL`, `BLOCKED`, missing, ambiguous or depends on human decision.

Use retries only within the configured retry limit.

## State Rules

Use `automation/INTERACTIVE_STATE.json` for interactive trails.

Never use interactive state to represent the fixed roadmap. Never use `automation/PHASE_STATE.json` to represent an interactive feature request.

When updating state, preserve:

- request id
- request type
- current step
- governing specs
- build spec
- diagnosis spec
- report path
- retry count
- latest classification
- latest decision
- blocked status and reason
- completion criteria

## Stop Conditions

Stop immediately when:

- requirements remain materially ambiguous
- build spec is missing for a new trail
- diagnosis spec is missing for a new trail
- report is missing or inconsistent
- classification is not `PASS`
- decision does not allow advancement
- tests required by the diagnosis spec fail
- implementation would exceed the approved spec
- a change would mutate fixed roadmap state to bypass a gate
- a human architecture decision is required

## Final Response Format

Always answer with:

- objetivo
- specs lidas
- skills aplicadas
- arquivos alterados
- comandos executados
- testes executados
- report lido
- classificacao
- decisao
- estado atualizado
- gaps restantes
- proxima etapa ou motivo de bloqueio

## Minimal Prompts

Start an interactive trail:

```text
Execute o interactive_sdd_autopilot para esta solicitacao.
Comece por intake, levante requisitos, decida entre amendar spec existente ou abrir nova trilha dual-spec, implemente apenas depois de specs suficientes, execute diagnostico, leia o report e avance somente quando o gate passar.
```

Continue an interactive trail:

```text
Continue o interactive_sdd_autopilot a partir de automation/INTERACTIVE_STATE.json.
Leia o estado, specs e report atuais, preserve o phase_autopilot e avance somente se o report permitir.
```
</file>

<file path="skills/24-agent-readiness-governance/SKILL.md">
---
name: agent-readiness-governance
description: Govern agent-readiness analysis for {{PROJECT_SLUG}} feature, refactor and PR workflows. Use when running, reading or applying @kodus/agent-readiness results, deciding which findings matter for a Go library/framework, creating readiness reports, attaching readiness evidence to PRs, updating docs, or deciding whether readiness regressions should block a change.
---

# Agent Readiness Governance

Goal: use `@kodus/agent-readiness` as an evidence input for `{{PROJECT_SLUG}}` without turning generic app/SaaS checks into inappropriate PR blockers for a Go library/framework.

## Required reading

Before applying this skill, read:

1. `AGENTS.md`
2. `skills/00-skill-index/SKILL.md`
3. `skills/21-documentation-open-source/SKILL.md`
4. `skills/20-testing-strategy-regression-load-containers/SKILL.md`
5. `skills/23-{{PROJECT_SLUG}}-sdd-autopilot/SKILL.md` when the work is part of an interactive autopilot trail
6. `docs/interactive-sdd-autopilot.md` when the work is part of an interactive autopilot trail
7. the current `agent-readiness` report if one already exists
8. `specs/720-agent-readiness-governance-assessment.md` and `specs/721-agent-readiness-governance-assessment-diagnosis.md` when running the baseline {{PROJECT_SLUG}} readiness assessment

If release, CI, security or OpenAPI files are changed, also route through the relevant domain skill from `skills/00-skill-index/SKILL.md`.

## Non-negotiables

1. Treat `agent-readiness` as an advisory and evidence source, not a universal gate.
2. Filter every finding through the `{{PROJECT_SLUG}}` context: Go library/framework, public API in `pkg/*`, docs/specs/reports, automation and agent workflows.
3. Do not block PRs on checks that are irrelevant to a library/framework.
4. Do block or request remediation when a change degrades agent operability in a relevant area.
5. Always classify findings as `worth_it_for_{{PROJECT_SLUG}}`, `optional_for_{{PROJECT_SLUG}}` or `out_of_scope_for_{{PROJECT_SLUG}}`.
6. Always produce a readiness report artifact for feature/refactor PRs when this skill is invoked.
7. Do not let a high raw score hide a relevant regression.
8. Do not let a low raw score force out-of-scope work.
9. When a finding is rejected as not relevant, record the reason.
10. When this skill is part of an interactive autopilot trail, the readiness report is supporting evidence; the phase/diagnosis report remains the gate source of truth.

## Classification rules

### `worth_it_for_{{PROJECT_SLUG}}`

Classify a finding as `worth_it_for_{{PROJECT_SLUG}}` when it improves one or more of:

- agent ability to understand repository purpose, architecture or constraints
- SDD/dual-spec workflow clarity
- Cursor/agent rules and automation reliability
- PR review quality and reproducibility
- testability, smoke coverage, conformance or regression safety
- CI checks for library quality
- security hygiene relevant to a Go library/framework
- open-source contributor onboarding
- public docs, examples and API discoverability
- explicit evidence for generated changes

These findings may block a PR when the PR introduces or worsens the relevant deficiency.

### `optional_for_{{PROJECT_SLUG}}`

Classify a finding as `optional_for_{{PROJECT_SLUG}}` when it can help but should not block normal PRs unless the PR directly changes that area. Examples:

- additional devcontainer quality when local setup already works
- broader CI matrix beyond the current supported baseline
- dependency update automation beyond minimal safety
- extra templates or governance docs that improve contributor experience but are not required for the current change
- extra examples when existing examples already prove the modified capability

### `out_of_scope_for_{{PROJECT_SLUG}}`

Classify a finding as `out_of_scope_for_{{PROJECT_SLUG}}` when it is primarily for an application, SaaS product or deployed service rather than a Go library/framework. Do not block PRs on these unless the project explicitly opens a spec making them relevant.

Common non-blocking/out-of-scope examples:

- production deployment pipeline requirements
- hosting platform readiness
- app runtime monitoring dashboards outside local observability
- bundle-size or frontend asset analysis when no frontend product is changed
- product analytics instrumentation
- Kubernetes/infra hardening unrelated to library usage
- cloud deployment checks
- SaaS incident response processes
- Docker/devcontainer requirements when the repo already has a documented Go-only local flow and no containerized workflow is in scope

## Required workflow

### 1. Collect inputs

Gather:

- change summary or user request
- touched files
- current branch/PR context when available
- existing `agent-readiness` report, or command output if a new run was performed
- relevant specs, reports, docs and automation files

If no report exists and the task explicitly requires readiness assessment, run or request a report using the repository-approved method. Prefer JSON output if available, but do not depend on JSON if text is the only available artifact.

### 2. Normalize the report

Produce a normalized list of findings with:

- id or stable name
- original finding text
- source category/pillar if available
- severity/score impact if available
- affected files or repository areas
- initial classification
- final `{{PROJECT_SLUG}}` classification
- blocker status
- rationale

Use `resources/classification-rules.md`.

### 3. Decide PR impact

A finding may block a PR only when all are true:

- classification is `worth_it_for_{{PROJECT_SLUG}}`
- the PR introduces, worsens or leaves unaddressed a relevant deficiency in files it touches
- remediation is feasible inside the current scope or the PR must explicitly defer it via spec/report
- the finding is not contradicted by stronger repository evidence

A finding must not block a PR when:

- classification is `optional_for_{{PROJECT_SLUG}}` or `out_of_scope_for_{{PROJECT_SLUG}}`
- it is unrelated to changed surfaces
- it requires product/SaaS/deployment work outside repository scope
- it conflicts with existing non-goals or phase scope

### 4. Generate report artifacts

For every readiness assessment, create or update a report under one of:

- `docs/reports/agent-readiness/<YYYY-MM-DD>-<short-scope>.md`
- `docs/reports/agent-readiness/<pr-number-or-branch>.md`

If the repo does not yet have the directory, create it.

The report must use `resources/report-template.md`.

### 5. Update PR and documentation

When operating in a PR workflow:

- include a PR section using `resources/pr-template-snippet.md`
- attach or link the readiness report path
- list only relevant blockers
- list ignored findings with rationale

Update project docs only when:

- a recurring relevant gap is discovered
- a user-facing contributor workflow changes
- agent/autopilot rules change
- README/AGENTS/docs are contradicted by the readiness result

### 6. Feed autopilot decisions carefully

For interactive SDD autopilot:

- readiness findings can create a new improvement trail only if classified `worth_it_for_{{PROJECT_SLUG}}`
- readiness findings cannot override a phase report gate
- readiness findings cannot force work outside the active feature/refactor scope unless the user approves a new trail
- readiness report should be referenced as supporting evidence in the feature/refactor diagnosis report

## Required output

When this skill is used, respond with these sections:

1. **Readiness source**
2. **Scope interpreted for {{PROJECT_SLUG}}**
3. **Classified findings**
4. **PR blockers**
5. **Non-blocking recommendations**
6. **Out-of-scope findings ignored**
7. **Report artifact path**
8. **Docs/PR updates required**
9. **Autopilot impact**

## References

- Use `resources/classification-rules.md` for filtering rules.
- Use `resources/report-template.md` for the generated report shape.
- Use `resources/pr-template-snippet.md` for PR text.
</file>

<file path="skills/24-agent-readiness-governance/resources/classification-rules.md">
# {{PROJECT_TITLE}} Agent Readiness Classification Rules

## Purpose

Filter generic `agent-readiness` findings through the scope of `{{PROJECT_SLUG}}`: a Go library/framework with SDD, automation, specs, reports, examples and public API contracts.

## Classification decision tree

### Step 1: Is the finding about agent ability to modify this repository safely?

If yes, classify as `worth_it_for_{{PROJECT_SLUG}}` unless it is already fully covered by existing evidence.

Typical examples:

- missing or stale agent instructions
- unclear architecture boundaries
- missing spec/report convention
- missing PR checklist for generated changes
- missing tests for changed public behavior
- missing CI gate for tests
- missing docs for a public or operational surface

### Step 2: Is it helpful but not required for the current change?

If yes, classify as `optional_for_{{PROJECT_SLUG}}`.

Typical examples:

- additional examples beyond minimum coverage
- optional devcontainer when normal Go setup is documented
- extra documentation polish not required for the touched surface
- expanded CI matrix when base CI is green

### Step 3: Is it app/SaaS/deployment specific?

If yes, classify as `out_of_scope_for_{{PROJECT_SLUG}}` unless the repository has a spec that explicitly brings it into scope.

Typical examples:

- production deployment platform
- cloud infrastructure monitoring
- product analytics
- frontend bundle checks
- SaaS incident-management workflows
- hosted environment operations

## Blocker policy

### Block PR

Block or request remediation when a `worth_it_for_{{PROJECT_SLUG}}` finding applies directly to changed surfaces and the PR would leave agents less able to reason about, test or safely modify the project.

Examples:

- PR changes public API but docs/OpenAPI/specs are not updated.
- PR changes automation but `AGENTS.md` or runbooks are now misleading.
- PR changes behavior with no test, smoke or diagnosis evidence.
- PR changes release, security or PR governance and leaves contradictions.

### Do not block PR

Do not block when:

- the finding is `optional_for_{{PROJECT_SLUG}}`
- the finding is `out_of_scope_for_{{PROJECT_SLUG}}`
- the finding is unrelated to the files changed
- the finding is historical and does not affect the current PR
- remediation would require a separate feature/spec not approved in the current trail

## Required rationale phrases

Use direct labels in reports:

- `BLOCKS_PR`: must be fixed or explicitly waived by spec/report.
- `DOES_NOT_BLOCK_PR`: advisory only.
- `OUT_OF_SCOPE`: should not be worked in this PR.
- `SEPARATE_TRAIL_REQUIRED`: relevant but too large for current PR.

## {{PROJECT_TITLE}}-specific relevant areas

Readiness findings are especially relevant when touching:

- `AGENTS.md`
- `.cursor/rules/*`
- `automation/*`
- `skills/*`
- `specs/*`
- `docs/*`
- `docs/reports/*`
- `.github/workflows/*`
- `README.md`
- `pkg/*` public API
- `api/openapi/*` and `api/openapi.yaml`
- `test/*`
- `examples/*`

## {{PROJECT_TITLE}}-specific non-goals

Do not introduce these merely to satisfy readiness score:

- hosted control plane
- {{UPSTREAM_OPS_NAME}}-style managed services
- SaaS dashboards
- production deployment platform
- Kubernetes deployment automation
- mandatory container workflow
- application analytics
- frontend performance tooling unless the changed area is a frontend app
</file>

<file path="skills/24-agent-readiness-governance/resources/pr-template-snippet.md">
## Agent Readiness

- Report: `<docs/reports/agent-readiness/...>`
- Source: `<agent-readiness command/report>`
- {{PROJECT_TITLE}}-filtered result: `<READY | READY_WITH_WARNINGS | NOT_READY>`
- PR impact: `<BLOCKS_PR | DOES_NOT_BLOCK_PR>`

### Relevant blockers

- [ ] None
- [ ] `<blocker id>` - `<required fix>`

### Optional recommendations

- `<recommendation>`

### Out-of-scope findings ignored

- `<finding>` - ignored because `<reason>`

### Notes

This readiness check filters generic `agent-readiness` output through the `{{PROJECT_SLUG}}` library/framework scope. Out-of-scope app/SaaS/deployment findings do not block this PR.
</file>

<file path="skills/24-agent-readiness-governance/resources/report-template.md">
# Agent Readiness Report - <scope>

Date: <YYYY-MM-DD>
Scope: <feature/refactor/PR/branch>
Source: <agent-readiness command/report path>

## 1. Executive summary

- Raw readiness score: `<score or unknown>`
- {{PROJECT_TITLE}}-filtered result: `<READY | READY_WITH_WARNINGS | NOT_READY>`
- PR impact: `<BLOCKS_PR | DOES_NOT_BLOCK_PR>`
- Relevant blockers: `<count>`
- Optional recommendations: `<count>`
- Out-of-scope findings ignored: `<count>`

## 2. Scope interpreted for {{PROJECT_SLUG}}

Describe why this assessment matters for a Go library/framework and which areas were considered in scope.

## 3. Source evidence

List:

- report path or command used
- timestamp
- format (`text`, `json`, or `manual summary`)
- relevant commit/branch/PR if known

## 4. Classified findings

| Finding | Original category | {{PROJECT_TITLE}} classification | PR impact | Rationale |
| --- | --- | --- | --- | --- |
| `<id/name>` | `<category>` | `worth_it_for_{{PROJECT_SLUG}}` | `BLOCKS_PR` | `<why>` |
| `<id/name>` | `<category>` | `optional_for_{{PROJECT_SLUG}}` | `DOES_NOT_BLOCK_PR` | `<why>` |
| `<id/name>` | `<category>` | `out_of_scope_for_{{PROJECT_SLUG}}` | `OUT_OF_SCOPE` | `<why>` |

## 5. PR blockers

List only findings that block the current PR.

For each blocker:

- finding
- affected files
- required fix
- acceptance evidence

## 6. Non-blocking recommendations

List findings worth considering later, grouped by theme.

## 7. Out-of-scope findings ignored

List findings intentionally ignored for this PR with rationale.

## 8. Documentation updates

State whether updates are required for:

- `AGENTS.md`
- `README.md`
- `docs/*`
- `automation/*`
- `.cursor/rules/*`
- `skills/*`
- PR description

## 9. Autopilot impact

State whether this report:

- creates no new autopilot work
- adds tasks to the current trail
- requires a separate SDD trail
- blocks advancement

## 10. Final decision

Use exactly one:

- `READY`
- `READY_WITH_WARNINGS`
- `NOT_READY`

Then explain in one paragraph.
</file>

<file path="skills/25-{{PROJECT_SLUG}}-ultra-rigid-sdd/SKILL.md">
---
name: {{PROJECT_SLUG}}-ultra-rigid-sdd
description: use for {{PROJECT_SLUG}} spec-driven development, autopilot prompt generation, phase planning, implementation specs, diagnosis specs, reconciliation phases, release/readiness gates, and any request that must follow the user's ultra-rigid evidence-first project standard. applies to {{PROJECT_SLUG}} framework parity, documentation, conformance, roadmap closure, and future sdd phases.
---

# {{PROJECT_TITLE}} Ultra-Rigid SDD

## Core rule

Always apply the ultra-rigid model when the user asks for {{PROJECT_SLUG}} phases, specs, diagnosis specs, Cursor/Codex/autopilot prompts, roadmap validation, parity work, documentation reconciliation, readiness confirmation, release/versioning governance, or any continuation of the {{PROJECT_SLUG}} SDD process.

The default posture is: **scope locked, evidence-first, no accidental capability expansion, paired feature spec + diagnosis spec, and explicit gates.**

## Project boundary

Treat `{{PROJECT_SLUG}}` as a reusable, versionable Go framework/library for agents, sub-agents, tools, memory, workflows, HTTP exposure, observability, provider adapters, conformance, and examples.

Do not treat `{{PROJECT_SLUG}}` as the product application.

Do not put product-domain responsibilities into the framework core:
- no {{DOMAIN_ACTOR_TITLE}} Portal or {{EXTERNAL_SYSTEM}}-specific logic
- no queues or business workers
- no auth/RBAC of the final product
- no domain repositories such as {{DOMAIN_ENTITY_SET}}
- no hosted {{UPSTREAM_OPS_NAME}}/control plane work unless explicitly authorized
- no application payload formats as framework contracts unless intentionally generalized

The consumer application may use {{PROJECT_SLUG}}, but must remain responsible for external systems, queues, product auth, domain rules, and production deployment composition.

## Mandatory invocation workflow

When this skill is triggered:

1. **Restate the locked scope in one short paragraph.**
2. **Classify the request type:** feature phase, diagnosis phase, paired SDD phase, reconciliation, readiness confirmation, release/versioning, prompt generation, or audit.
3. **Check for prior completion:** if the user says a roadmap or phase is complete, do not reopen it unless evidence shows a real blocker.
4. **Generate only the smallest sufficient scope.**
5. **Always use the ultra-rigid output contract.**
6. **Include docs/demo/report updates whenever the task affects public behavior or discovery.**
7. **End with a binary gate or next-step decision.**

For detailed templates, read `references/ultra-rigid-output-contract.md` when generating prompts/specs, and read `references/sdd-phase-template.md` when creating a full SDD phase.

## Ultra-rigid defaults

### Always enforce

- Do not open capability nova without explicit user authorization.
- Do not reopen completed roadmaps by narrative speculation.
- Do not redesign runtime for reconciliation, diagnosis, polish, or readiness tasks.
- Do not mix framework readiness with application readiness.
- Do not treat examples/helpers as substitutes for public framework contracts.
- Do not treat smoke tests as proof of full parity unless the diagnosis explicitly says so.
- Do not persist a full runtime operation context automatically.
- Do not bypass docs, feature matrix, reports, or examples when affected.
- Do not skip the diagnosis spec.
- Do not skip explicit `PASS / PARTIAL / FAIL / BLOCKED` classification.

### Evidence order

Use this precedence when reasoning:

1. Exported contracts in `pkg/*`, real HTTP/OpenAPI surface, real behavior.
2. Automated tests, conformance, smoke checks, race checks, benchmarks when applicable.
3. Normative specs and diagnosis specs.
4. Reports, docs, feature matrix, README, examples.
5. Historical notes and chat context.

If lower-precedence artifacts disagree with code/tests, mark them as drift and propose reconciliation, not capability expansion.

## Required SDD pairing

For any implementation phase, produce both:

1. Feature/implementation spec.
2. Diagnosis spec.

Also produce prompts in this order:

1. Prompt to create the feature spec.
2. Prompt to implement the feature spec.
3. Prompt to create the diagnosis spec.
4. Prompt to execute the diagnosis.

Each prompt must include:
- task
- mandatory file path
- objective
- locked scope
- prohibited scope
- required artifacts
- validation gates
- exact output format
- rejection criteria

## Required phase sections

Every feature spec must include:

1. objetivo
2. motivacao
3. pergunta principal
4. escopo
5. requisitos de design
6. blocos obrigatorios da fase
7. requisitos funcionais
8. decisoes obrigatorias de modelagem
9. criterios de aceitacao
10. evidencia obrigatoria
11. fora do escopo
12. perguntas secundarias

Every diagnosis spec must include:

1. objetivo
2. pergunta principal
3. escopo
4. checklist obrigatorio
5. criterios de classificacao
6. saida obrigatoria
7. criterios de aceitacao

## Mandatory gate language

Use explicit binary decisions. Examples:

- `READY FOR PHASE N`
- `NOT READY FOR PHASE N`
- `FRAMEWORK READY FOR GO APP REWRITE`
- `FRAMEWORK NOT READY FOR GO APP REWRITE`
- `READY FOR STABLE RELEASE GOVERNANCE`
- `NOT READY FOR STABLE RELEASE GOVERNANCE`
- `ROADMAP COMPLETE`
- `ROADMAP NOT COMPLETE`

Never end with only "recommended" or "looks good".

## Validation commands

Default validation gates:

```bash
go test ./...
go test -race ./...
go vet ./...
```

Add these when applicable:

```bash
go run ./cmd/openapi-bundle -check
```

For examples/demos:

```bash
go test ./test/smoke/...
go test ./examples/...
go build ./examples/<example-name>/
```

For storage changes, require restart-proof tests.
For streaming changes, require replay/resume tests.
For public API changes, require semver impact notes.
For concurrency changes, require race checks and at least one concurrent scenario.

If a command cannot be run, require the implementing agent to state why and classify the gate as `BLOCKED` or `PARTIAL`, not `PASS`.

## Documentation and demo rules

Whenever public behavior changes, require updates to affected artifacts, typically:

- `README.md`
- `specs/010-feature-matrix.md`
- `docs/*.md`
- `docs/reports/*.md`
- `examples/*/README.md`
- reference consumer or starter example when applicable
- OpenAPI bundle when HTTP surface changes
- `CHANGELOG.md` or release docs when API surface changes

Do not allow implementation-only completion when discovery artifacts become stale.

## Rejection criteria

A generated response is invalid if it:

- opens capability nova outside the requested scope
- reopens a completed roadmap without hard evidence
- omits paired feature and diagnosis specs for an SDD phase
- omits exact file paths
- omits validation commands
- omits docs/demo/report updates when applicable
- confuses {{PROJECT_SLUG}} framework readiness with application readiness
- treats product-domain concerns as {{PROJECT_SLUG}} core
- gives a generic plan without gate language
- lacks `PASS / PARTIAL / FAIL / BLOCKED` criteria

## Standard concise response when acknowledging a completed roadmap

Use this stance:

> Treat the completed roadmap as closed. Do not reopen the main parity track. If residual issues exist, propose only a small reconciliation or readiness-confirmation phase with locked scope, evidence-first checks, and no new capability.

## Standard concise response when generating a future phase

Use this stance:

> Generate one minimal SDD phase, with feature spec + diagnosis spec + four autopilot prompts. Lock the scope, declare prohibited scope, list required artifacts, require validation gates, require docs/demo updates, and include rejection criteria.
</file>

<file path="skills/25-{{PROJECT_SLUG}}-ultra-rigid-sdd/agents/openai.yaml">
interface:
  display_name: "{{PROJECT_TITLE}} Ultra-Rigid SDD"
  short_description: "Generate {{PROJECT_SLUG}} SDD specs, diagnoses, and autopilot prompts with strict evidence-first gates."
</file>

<file path="skills/25-{{PROJECT_SLUG}}-ultra-rigid-sdd/references/sdd-phase-template.md">
# SDD Phase Template

Use this reference when creating a full paired SDD phase for {{PROJECT_SLUG}}.

## Feature spec template

```md
# Spec <number>: Phase <n> - <title>

## 1. Objetivo

## 2. Motivacao

## 3. Pergunta principal

## 4. Escopo

### Dentro do escopo

### Fora do escopo

## 5. Requisitos de design

## 6. Blocos obrigatorios desta fase

## 7. Requisitos funcionais

## 8. Decisoes obrigatorias de modelagem

## 9. Criterios de aceitacao

## 10. Evidencia obrigatoria

## 11. Fora do escopo tecnico explicito

## 12. Perguntas secundarias
```

## Diagnosis spec template

```md
# Spec <number>: Phase <n> - <title> Diagnosis

## 1. Objetivo

## 2. Pergunta principal

## 3. Escopo

## 4. Checklist obrigatorio

## 5. Classificacao

- PASS
- PARTIAL
- FAIL
- BLOCKED

## 6. Saida obrigatoria

`docs/reports/phase-<n>-<slug>-report.md`

## 7. Criterios de aceitacao
```

## File naming pattern

Prefer:

- `specs/<number>-phase-<n>-<slug>.md`
- `specs/<number>-phase-<n>-<slug>-diagnosis.md`
- `docs/reports/phase-<n>-<slug>-report.md`

If the repo already uses a different numbering range, follow the repo's current numbering.

## Prompt generation pattern

For each phase, generate exactly four prompts:

1. Create feature spec.
2. Implement feature spec.
3. Create diagnosis spec.
4. Execute diagnosis spec.

## Boundary language to reuse

Use these phrases when relevant:

- `nao abrir capability nova`
- `nao reabrir o roadmap principal`
- `tratar {{PROJECT_SLUG}} como lib reutilizavel e versionavel`
- `avaliar evidencia real, nao intencao`
- `diferenciar readiness de framework de readiness da aplicacao`
- `gerar report canonico da fase`
- `concluir com decisao binaria`
</file>

<file path="skills/25-{{PROJECT_SLUG}}-ultra-rigid-sdd/references/ultra-rigid-output-contract.md">
# Ultra-Rigid Output Contract

Use this reference whenever generating Cursor/Codex/autopilot prompts, specs, diagnosis specs, or audit reports for {{PROJECT_SLUG}}.

## Required response shape for an autopilot prompt set

For each phase, output exactly:

1. `fase <n> - <title>`
2. objective
3. locked scope
4. prohibited scope
5. implementation spec file path
6. implementation prompt
7. implementation execution prompt
8. diagnosis spec file path
9. diagnosis spec prompt
10. diagnosis execution prompt
11. gate language
12. rejection criteria

## Required prompt sections

Every prompt must include:

- `TAREFA`
- `ARQUIVO OBRIGATORIO` or `SPEC DE REFERENCIA`
- `OBJETIVO`
- `CONTEXTO OBRIGATORIO`
- `ESCOPO PERMITIDO`
- `ESCOPO PROIBIDO`
- `ARTEFATOS A REVISAR` or `ARTEFATOS A ALTERAR`
- `TESTES/GATES OBRIGATORIOS`
- `FORMATO DE SAIDA OBRIGATORIO`
- `CRITERIOS DE REJEICAO`

## Required implementation response shape

For implementation prompts, require this exact final structure:

```markdown
## 1. Escopo executado
- o que foi alterado
- o que foi deliberadamente nao alterado

## 2. Arquivos alterados
- lista completa

## 3. Implementacao realizada
- problema
- arquivo(s)
- solucao aplicada

## 4. Validacao executada
- comandos rodados
- resultado de cada comando

## 5. Confirmacao de escopo
- por que nao abriu capability fora do escopo

## 6. Limitacoes remanescentes
- gaps reais, se houver
```

## Required diagnosis response shape

For diagnosis execution prompts, require this exact final structure:

```markdown
## 1. Classificacao final
- PASS, PARTIAL, FAIL ou BLOCKED

## 2. Decisao binaria
- READY/NOT READY language appropriate to the phase

## 3. Principais evidencias
- lista objetiva

## 4. Gaps encontrados
- gaps reais somente

## 5. Confirmacao de boundary
- framework vs application boundary, if relevant
```

## Classification rules

Use these meanings consistently:

- `PASS`: all mandatory axes satisfied with evidence and no blocker.
- `PARTIAL`: meaningful progress, but one or more required axes remain incomplete or ambiguous.
- `FAIL`: required work was not completed or scope was violated.
- `BLOCKED`: evidence could not be collected, commands could not run, or the repo state cannot be verified.

Never classify `PASS` when required commands were not run unless the phase explicitly allows a documentation-only audit without commands.

## Capability expansion guard

If a task is reconciliation, polish, confirmation, diagnosis, versioning governance, or documentation alignment, require the agent to state:

> Esta fase nao abre capability nova.

If a task truly needs a new capability, require explicit user authorization or a new feature phase.
</file>

<file path="skills/25-{{PROJECT_SLUG}}-ultra-rigid-sdd/references/{{PROJECT_SLUG}}-boundaries.md">
# {{PROJECT_TITLE}} Boundaries

Use this reference when a request risks mixing framework responsibilities with application responsibilities.

## {{PROJECT_TITLE}} core may include

- agent runtime and invocation contracts
- tools and tool middleware
- memory and conversation framework contracts
- workflow contracts
- supervisor/sub-agent contracts
- provider adapters as reusable framework integrations
- HTTP/API exposure as an optional framework server
- OpenAPI artifacts
- local observability hooks/helpers
- conformance suites
- examples/reference consumer
- release/versioning governance

## Consumer application should own

- domain rules
- product-specific auth/RBAC
- external queues
- third-party business integrations
- data repositories for product entities
- channel-specific payload formats
- deployment topology
- hosted platform concerns
- business observability dashboards

## Decision rule

If a requested change is useful for many {{PROJECT_SLUG}} consumers and can be expressed as a reusable contract, it may belong in {{PROJECT_SLUG}}.

If it encodes one product's domain, one channel's payload, one tenant's policy, or one external vendor workflow, keep it out of {{PROJECT_SLUG}} core and expose extension points instead.
</file>

<file path="specs/000-project-mission.md">
# Spec 000: Project Mission

## 1. Objetivo

`{{PROJECT_SLUG}}` existe para entregar `{{PROJECT_DESCRIPTION}}` como um projeto Go versionável, testável e operável por humanos e agentes.

## 2. Regras globais

1. Nenhuma feature entra sem especificação aprovada.
2. Toda capacidade suportada precisa de critério de aceitação observável.
3. Toda implementação deve ter teste ou diagnóstico correspondente.
4. O design público deve ser idiomático em Go.
</file>

<file path="specs/001-non-goals.md">
# Spec 001: Non-goals

## 1. Objetivo

Este documento define limites explícitos para evitar expansão acidental de escopo em `{{PROJECT_SLUG}}`.

## 2. Fora do escopo inicial

1. Serviços hospedados obrigatórios para o runtime local.
2. Dashboards, control planes ou automação de deploy como dependência de desenvolvimento.
3. Regras de domínio de produto que não pertençam ao projeto.
</file>

<file path="specs/010-feature-matrix.md">
# Spec 010: Feature Matrix

## 1. Objetivo

Esta matriz lista capabilities planejadas, status, prioridade, risco e evidência para `{{PROJECT_SLUG}}`.

## 2. Matriz

| Feature | Descrição | Obrigatória? | Prioridade | Módulo provável | Estado | Verificação | Risco | Observações |
| --- | --- | --- | --- | --- | --- | --- | --- | --- |
| Bootstrap governance | Baseline AI-first, specs, skills, automation, docs/ai e CI. | Sim | P0 | repo | Implementado | PASS | Médio | Criado pela Phase 0 bootstrap. |
</file>

<file path="specs/020-repository-architecture.md">
# Spec 020: Repository Architecture

## 1. Objetivo

Este documento define a arquitetura de repositorio da `{{PROJECT_SLUG}}` para:

- separar contratos publicos estaveis do runtime interno
- reduzir acoplamento entre capacidades centrais
- preparar a base para testes de conformidade, examples executaveis e extensibilidade futura
- tornar explicitas as regras de evolucao e compatibilidade do modulo

Esta spec complementa a `Spec 000` e a `Spec 010`. Ela nao define comportamento de uma feature isolada; ela define a forma como o repositorio deve ser organizado para que as features sejam implementadas sem erosao arquitetural.

## 2. Principios arquiteturais

Os principios abaixo sao normativos para a organizacao do repositorio:

1. API publica pequena e intencional. Apenas contratos, tipos e pontos de extensao que precisem ser importados por usuarios devem viver em `pkg/`.
2. Runtime escondido. Logica de orquestracao, execucao, pipeline interno e detalhes de estado devem viver em `internal/`.
3. Dependencias aciclicas. A arvore de pacotes deve formar camadas claras e sem ciclos.
4. Contratos antes de implementacoes. Interfaces, structs de configuracao, requests, responses e erros observaveis devem ser definidos nas camadas publicas antes da concretizacao do runtime.
5. `context.Context` como fronteira padrao. Cancelamento, timeout e lifecycle devem atravessar a API publica de forma idiomatica.
6. Observabilidade sem vazamento. Eventos internos podem existir para runtime, historico e instrumentacao, mas nao devem aparecer em assinaturas exportadas sem promocao explicita para uma API publica estavel.
7. Extensibilidade por composicao. Tools, memory, guardrails e workflows devem expor contratos estaveis para extensao sem exigir importacao de pacotes internos.
8. Testabilidade por camada. Cada camada deve poder ser validada com testes coerentes com sua responsabilidade.
9. Root package minimo. O pacote raiz `{{PROJECT_SLUG}}` deve continuar pequeno, reservado para documentacao do modulo e, no maximo, metadados globais estaveis.
10. Compatibilidade observavel acima de detalhe estrutural. O que importa para estabilidade e o contrato importavel e o comportamento observavel, nao o desenho interno do runtime.

## 3. Organizacao de pacotes

### 3.1 Arvore sugerida

```text
.
├── doc.go
├── go.mod
├── pkg
│   ├── agent
│   ├── app
│   ├── guardrail
│   ├── logger
│   ├── memory
│   ├── server
│   ├── tool
│   ├── types
│   └── workflow
├── internal
│   ├── events
│   ├── runtime
│   │   ├── agent
│   │   ├── lifecycle
│   │   ├── memory
│   │   ├── registry
│   │   ├── tool
│   │   └── workflow
│   └── validation
├── test
│   ├── conformance
│   │   ├── agent
│   │   ├── app
│   │   ├── guardrail
│   │   ├── memory
│   │   ├── tool
│   │   └── workflow
│   └── fixtures
│       ├── cases
│       └── golden
├── examples
│   ├── basic-agent
│   ├── custom-memory
│   ├── http-server
│   └── workflow-chain
└── specs
```

### 3.2 Papel de cada pacote

| Caminho | Papel | Natureza |
| --- | --- | --- |
| `pkg/types` | Tipos compartilhados de baixo nivel: ids, metadata, envelopes, erros observaveis e contratos pequenos usados por mais de um pacote publico. Nao deve virar deposito generico. | Publico estavel |
| `pkg/logger` | Fachada de logging e adaptadores sobre `log/slog`, incluindo noop/test logger quando necessario. | Publico estavel |
| `pkg/tool` | Contratos de tool, schema de input/output, helpers de registro e composicao de toolkits. | Publico estavel |
| `pkg/memory` | Contratos de memoria, working memory e persistencia conversacional plugavel. | Publico estavel |
| `pkg/guardrail` | Contratos de validacao e decisao antes/depois da execucao, com resultados observaveis de permitir, bloquear, transformar ou falhar. | Publico estavel |
| `pkg/workflow` | Definicao publica de workflows, steps, branching, retry, hooks e historico observavel. | Publico estavel |
| `pkg/agent` | Definicao publica de agentes, requests/responses, configuracao e registries relacionados ao dominio de agente. | Publico estavel |
| `pkg/app` | Composition root publica: builder da aplicacao, wiring de registries, lifecycle, startup e integracao controlada com o runtime interno. | Publico estavel |
| `pkg/server` | Adaptadores publicos para expor uma `app` por transporte sem contaminar os pacotes centrais com detalhes HTTP ou streaming. | Publico estavel, opcional |
| `internal/runtime` | Implementacao concreta do motor de execucao, roteamento, pipelines, registries internos, state transitions e lifecycle operacional. | Interno |
| `internal/events` | Modelo interno de eventos e envelopes usados pelo runtime para historico, hooks e observabilidade local. | Interno |
| `internal/validation` | Validacoes, normalizacao de configuracao e garantias de consistencia antes da execucao. | Interno |
| `test/conformance` | Suites que validam paridade comportamental via API publica usando fixtures e cenarios de referencia. | Teste |
| `test/fixtures` | Dados de entrada, saida esperada, cenarios e golden files usados por suites de conformidade. | Teste |
| `examples` | Programas minimos e executaveis que demonstram o uso da API publica e servem como smoke tests/documentacao viva. | Suporte |
| `specs` | Documentacao normativa de arquitetura, escopo e comportamento esperado. | Documentacao |

### 3.3 Papel das subcamadas internas

- `internal/runtime/agent`: execucao concreta de agentes e handoffs internos.
- `internal/runtime/tool`: pipeline de invocacao de tools, coercao interna e tratamento de falhas.
- `internal/runtime/workflow`: motor de steps, branching, retries e historico.
- `internal/runtime/memory`: adaptacao entre contratos publicos de memoria e armazenamentos concretos.
- `internal/runtime/registry`: registries internos, lookup otimizado e wiring de componentes.
- `internal/runtime/lifecycle`: startup, shutdown, cleanup e coordenacao com `context.Context`.

## 4. Regras de dependencia entre pacotes

### 4.1 Camadas

As camadas do repositorio devem seguir a direcao abaixo:

1. Base publica: `pkg/types`, `pkg/logger`
2. Contratos publicos de extensao: `pkg/tool`, `pkg/memory`, `pkg/guardrail`
3. Contratos publicos de orquestracao: `pkg/workflow`, `pkg/agent`
4. Composicao publica: `pkg/app`
5. Adaptadores publicos: `pkg/server`
6. Implementacao interna: `internal/validation`, `internal/events`, `internal/runtime`
7. Verificacao e suporte: `test/*`, `examples`, `specs`

### 4.2 Regras normativas

As regras abaixo devem ser seguidas ao criar ou revisar imports:

1. `pkg/types` e `pkg/logger` so podem depender de stdlib ou dependencias externas estritamente fundacionais e estaveis.
2. `pkg/tool`, `pkg/memory` e `pkg/guardrail` podem depender de `pkg/types` e `pkg/logger`, mas nao de `pkg/app`, `pkg/server` ou `internal/*`.
3. `pkg/workflow` pode depender de `pkg/types`, `pkg/logger`, `pkg/tool`, `pkg/memory` e `pkg/guardrail`.
4. `pkg/agent` pode depender de `pkg/types`, `pkg/logger`, `pkg/tool`, `pkg/memory` e `pkg/guardrail`; deve evitar dependencia direta de `pkg/workflow` para nao misturar dominios centrais.
5. `pkg/app` e o unico pacote publico autorizado a importar `internal/runtime` e `internal/validation`.
6. `pkg/server` deve depender de `pkg/app` e outros contratos publicos, nunca de `internal/*`.
7. `internal/runtime` pode depender de qualquer `pkg/*`, alem de `internal/events` e `internal/validation`.
8. `internal/events` nao deve depender de `pkg/app` ou `pkg/server`; se um evento precisar ser observavel publicamente, ele deve ser promovido para um contrato estavel em `pkg/*`.
9. `internal/validation` nao deve depender de `internal/runtime`; validacao precisa permanecer reutilizavel e sem acoplamento ao motor.
10. `examples` so podem depender de `pkg/*`.
11. `test/conformance` deve exercitar a biblioteca via `pkg/*`, nunca via `internal/*`.
12. `test/fixtures` nao pode ser importado por `pkg/*` nem por `internal/*`.
13. `specs` e `examples` nunca sao dependencias de `pkg/*`.
14. Nenhum pacote em `pkg/*` pode depender de `pkg/server`.
15. Nenhum pacote em `pkg/*` pode depender de `examples`, `test/*` ou `specs`.
16. Qualquer necessidade de importar `internal/*` fora de `pkg/app` indica erro de fronteira arquitetural e deve ser corrigida.

### 4.3 Matriz resumida de importacao

| Origem | Pode importar |
| --- | --- |
| `pkg/types` | stdlib |
| `pkg/logger` | stdlib, dependencias de logging aprovadas |
| `pkg/tool` | `pkg/types`, `pkg/logger` |
| `pkg/memory` | `pkg/types`, `pkg/logger` |
| `pkg/guardrail` | `pkg/types`, `pkg/logger` |
| `pkg/workflow` | `pkg/types`, `pkg/logger`, `pkg/tool`, `pkg/memory`, `pkg/guardrail` |
| `pkg/agent` | `pkg/types`, `pkg/logger`, `pkg/tool`, `pkg/memory`, `pkg/guardrail` |
| `pkg/app` | qualquer `pkg/*`, `internal/runtime`, `internal/validation` |
| `pkg/server` | `pkg/app`, `pkg/types`, `pkg/logger` |
| `internal/events` | `pkg/types`, `pkg/logger`, stdlib |
| `internal/validation` | `pkg/*`, stdlib |
| `internal/runtime` | qualquer `pkg/*`, `internal/events`, `internal/validation` |
| `examples` | qualquer `pkg/*` |
| `test/conformance` | qualquer `pkg/*`, `test/fixtures` |

## 5. O que pode ser publico vs interno

### 5.1 Deve ser publico

Devem viver em `pkg/`:

- interfaces e contratos que o usuario implementa
- tipos de request/response que atravessam a API
- erros e resultados observaveis
- builders, registries e configuracoes necessarias para compor a biblioteca
- hooks e pontos de extensao deliberadamente suportados
- adaptadores que o usuario precisa importar para expor a biblioteca em outro contexto

### 5.2 Deve ser interno

Devem viver em `internal/`:

- detalhes do motor de execucao
- state machines, steps intermediarios e pipeline de runtime
- envelopes de eventos usados apenas para coordenacao interna
- caches, validacoes, defaults e normalizacao sem contrato publico
- helpers de wiring e registries que nao precisem ser estendidos pelo usuario
- estruturas criadas apenas para performance, concorrencia ou simplificacao interna

### 5.3 Regras de fronteira

1. `internal/*` nao pode aparecer em exemplos de uso da biblioteca.
2. Tipos de `internal/*` nao podem ser expostos em assinaturas exportadas, campos exportados ou erros retornados por `pkg/*`.
3. Se um conceito interno precisar ser customizado por usuarios, ele deve ser promovido para um contrato publico em `pkg/*` antes de qualquer extensao adicional.
4. Tipos compartilhados entre mais de um pacote publico devem ir para `pkg/types` apenas quando forem realmente transversais e estaveis.
5. Tipos usados por apenas um pacote publico devem permanecer no proprio pacote, nao em `pkg/types`.
6. O pacote raiz `{{PROJECT_SLUG}}` nao deve virar umbrella API; imports publicos devem continuar explicitos em `pkg/...`.

## 6. Estrategia de versionamento e compatibilidade

### 6.1 Unidade de versionamento

O repositorio deve seguir versionamento semantico no nivel do modulo Go.

- `MAJOR`: quebra de compatibilidade em `pkg/*` ou mudanca observavel nao retrocompativel em comportamento documentado.
- `MINOR`: adicao retrocompativel de capacidades, tipos, funcoes, pacotes publicos ou pontos de extensao.
- `PATCH`: correcao de bugs e ajustes internos sem quebra de contrato publico.

### 6.2 Escopo da garantia de compatibilidade

A garantia de compatibilidade cobre:

- identificadores exportados em `pkg/*`
- comportamento observavel documentado em specs aprovadas
- formatos e semantica de erros publicos
- ordem de execucao relevante quando ela fizer parte do contrato observavel
- examples oficiais e suites de conformidade que representem cenarios suportados

A garantia de compatibilidade nao cobre:

- qualquer pacote em `internal/*`
- organizacao interna do runtime
- utilitarios de `test/*`
- fixtures auxiliares que nao representem contrato publico
- documentos em `specs` que ainda nao tenham sido convertidos em comportamento implementado

### 6.3 Politica para fase pre-v1

Antes de `v1.0.0`, mudancas breaking em `pkg/*` ainda podem ocorrer, mas somente quando:

1. a spec correspondente for atualizada
2. a mudanca vier acompanhada de ajuste nas suites de conformidade e examples afetados
3. a divergencia estiver claramente registrada no changelog ou release note

### 6.4 Promocao de contratos

Qualquer artefato que hoje exista em `internal/*` e precise de garantia de estabilidade deve seguir este fluxo:

1. promover o conceito para `pkg/*`
2. documentar o novo contrato em spec
3. adicionar teste de conformidade ou contrato publico correspondente
4. so depois simplificar ou trocar a implementacao interna

## 7. Estrategia de testes por camada

### 7.1 `pkg/*`

Objetivo: validar contratos publicos, ergonomia e comportamento local de cada pacote.

- testes unitarios por pacote em `pkg/<nome>/*_test.go`
- foco em construtores, validacao de configuracao, erros publicos e comportamento documentado
- exemplos de package doc e exemplos executaveis devem compilar sem importar `internal/*`

### 7.2 `internal/*`

Objetivo: validar corretude do motor e detalhes de execucao.

- testes unitarios deterministas em `internal/.../*_test.go`
- foco em concorrencia, ordering, retries, lifecycle, pipelines e normalizacao
- mocks ou fakes devem consumir contratos de `pkg/*`, nao recriar contratos paralelos

### 7.3 `test/conformance`

Objetivo: validar paridade conceitual e comportamental a partir da superficie publica.

- testes devem usar `pkg/app` e demais `pkg/*`
- fixtures em `test/fixtures` devem descrever entradas, eventos observaveis, erros e saidas esperadas
- cenarios devem cobrir sucesso, falha, cancelamento e ordem relevante de operacoes
- suites de conformidade sao o principal gate de compatibilidade antes de release

### 7.4 `examples`

Objetivo: servir como documentacao viva e smoke test de integracao.

- cada exemplo deve demonstrar um caminho de uso real da API publica
- exemplos devem ser pequenos, executaveis e sem infraestrutura externa obrigatoria
- examples oficiais devem ser cobertos por `go test` ou por smoke test equivalente

### 7.5 Estrategia por camada funcional

| Camada | Tipo principal de teste | Resultado esperado |
| --- | --- | --- |
| `pkg/types`, `pkg/logger` | unitario | contratos basicos e estabilidade de tipos |
| `pkg/tool`, `pkg/memory`, `pkg/guardrail` | unitario + contrato | extensibilidade e erros observaveis |
| `pkg/workflow`, `pkg/agent` | unitario + contrato | definicao publica coerente e previsivel |
| `pkg/app` | integracao local | wiring correto com runtime interno |
| `pkg/server` | integracao local | adaptacao de transporte sem contaminar o core |
| `internal/runtime` | unitario determinista | execucao correta e acoplamento interno sob controle |
| `test/conformance` | conformidade | paridade comportamental do modulo |
| `examples` | smoke test | usabilidade minima e compilacao real da API |

## 8. Maintainability e coesao de arquivo

### 8.1 Faixas de tamanho

Arquivos Go de producao devem seguir as faixas abaixo, derivadas de `skills/13-go-idiomatic-effective-go/SKILL.md`:

| Linhas | Diretriz |
| --- | --- |
| <= 400 | Ideal |
| 401–700 | Aceitavel se houver responsabilidade unica clara |
| 701–1000 | Revisao obrigatoria de coesao, imports e agrupamento |
| > 1000 | Code smell forte — exige proposta explicita de split por responsabilidade |

### 8.2 Criterios de coesao

Quando um arquivo exceder 700 linhas, avaliar:

1. numero de responsabilidades visiveis
2. mistura de camadas arquiteturais no mesmo arquivo
3. numero de funcoes grandes ou helpers de natureza diferente
4. acoplamento com detalhes internos que poderiam ser isolados
5. dificuldade de teste isolado

### 8.3 Regras normativas

1. Preferir multiplos arquivos pequenos a um unico arquivo grande. O compilador Go compila pacotes, nao arquivos individuais.
2. Nomear arquivos por responsabilidade, nao por tipo generico (evitar `utils.go`, `helpers.go`, `models.go`).
3. Quando um split for necessario, preservar integralmente o contrato publico e o comportamento observavel.
4. Toda refatoracao de split deve ser acompanhada de execucao dos testes do pacote afetado e das suites de conformidade relacionadas.
5. Nao introduzir novos arquivos acima de 700 linhas sem justificativa tecnica explicita.

## 9. Regras curtas para enforcement

As regras abaixo devem ser usadas como checklist rapido em PRs:

- `pkg` nunca depende de `examples`, `test` ou `specs`
- `pkg/server` nunca sobe dependencia para o core
- `pkg/app` e a unica ponte publica para `internal/runtime`
- `internal` nao vaza para API publica
- `test/conformance` valida comportamento via API publica, nao via runtime interno
- `examples` mostram apenas imports publicos
- `pkg/types` nao vira pacote deposito
- contratos publicos sobem antes de implementacoes internas
- mudanca breaking em `pkg/*` exige revisao de spec e teste de conformidade
- arquivos acima de 700 linhas exigem revisao de coesao conforme secao 8

## 10. Decisao

Fica adotada a organizacao com `pkg/` como superficie publica estavel e `internal/` como implementacao privada do repositorio, com `pkg/app` como unica ponte oficial entre contratos publicos e runtime interno. Esta arquitetura deve ser usada como base para as proximas specs e para a implementacao incremental da `{{PROJECT_SLUG}}`.
</file>

<file path="specs/680-phase-0-bootstrap-foundation.md">
# Spec 680: Phase 0 - Bootstrap Foundation

## 1. Objetivo

Definir a fundação AI-first inicial de `{{PROJECT_SLUG}}`, incluindo governança para agentes, specs, skills, automation dual, docs operacionais, CI e scripts de compliance.

## 2. Motivação

O projeto deve nascer com padrões explícitos para que agentes consigam operar com baixa ambiguidade e alta rastreabilidade.

## 3. Pergunta principal

`{{PROJECT_SLUG}}` possui uma baseline versionada, verificável e executável para desenvolvimento Go assistido por agentes?

## 4. Escopo

### Dentro do escopo

1. Criar `AGENTS.md`.
2. Criar `skills/`.
3. Criar `.cursor/rules/` e hooks.
4. Criar `automation/`.
5. Criar `docs/ai/`.
6. Criar CI, scripts e PR template.

### Fora do escopo

1. Criar feature de domínio do produto.
2. Introduzir serviço hospedado obrigatório.
3. Abrir API pública sem spec específica.

## 5. Requisitos de design

1. A baseline deve ser local-first.
2. A governança deve ser evidence-first.
3. Trilhas novas devem exigir dual-spec.
4. Reports devem usar `PASS`, `PARTIAL`, `FAIL` ou `BLOCKED`.

## 6. Blocos obrigatórios desta fase

Governança para agentes, automação dual, specs base, CI/scripts e documentação humana.

## 7. Requisitos funcionais

Agentes devem saber o que ler, quando perguntar, como diagnosticar e como reportar evidência.

## 8. Decisões obrigatórias de modelagem

`ROADMAP.json` governa roadmap fixo; `INTERACTIVE_STATE.json` governa trilhas interativas; exceções formais vivem em `docs/ai/compliance-exceptions.md`.

## 9. Critérios de aceitação

1. `make check-compliance` passa.
2. `make test` passa.
3. `automation/*.json` é JSON válido.

## 10. Evidência obrigatória

Arquivos versionados, saída de comandos e report de bootstrap.

## 11. Fora do escopo técnico explícito

Não criar domínio de produto, hosted service, dashboard ou deploy tooling obrigatório.

## 12. Perguntas secundárias

O modo interativo inicia? O modo fixo opera? O CI detecta drift?
</file>

<file path="specs/681-phase-0-bootstrap-foundation-diagnosis.md">
# Spec 681: Phase 0 - Bootstrap Foundation Diagnosis

## 1. Objetivo

Definir a auditoria da Phase 0 de `{{PROJECT_SLUG}}`, confirmando que a baseline AI-first foi criada com evidência real.

## 2. Pergunta principal

A baseline AI-first está pronta para operar agentes de desenvolvimento Go com specs, skills, automation, reports e CI?

## 3. Escopo

Verificar arquivos obrigatórios, JSON de automation, scripts executáveis, docs/ai e prompts.

## 4. Checklist obrigatório

- [ ] `AGENTS.md` existe.
- [ ] `skills/00-skill-index/SKILL.md` existe.
- [ ] `automation/AUTOPILOT.md` existe.
- [ ] `automation/INTERACTIVE_AUTOPILOT.md` existe.
- [ ] `.cursor/rules/` existe.
- [ ] `.github/PULL_REQUEST_TEMPLATE.md` existe.
- [ ] `scripts/check-compliance.sh` existe e é executável.

## 5. Classificação

- `PASS`: todos os itens obrigatórios atendidos e checks verdes.
- `PARTIAL`: baseline existe mas algum eixo material está incompleto.
- `FAIL`: baseline não opera.
- `BLOCKED`: falta decisão humana ou dependência externa.

## 6. Saída obrigatória

`docs/reports/phase-0-bootstrap-foundation-report.md`

## 7. Critérios de aceitação

O report deve concluir `READY FOR FIRST FEATURE TRAIL` ou `NOT READY FOR FIRST FEATURE TRAIL`.
</file>

<file path="test/conformance/.gitkeep">

</file>

<file path="test/fixtures/.gitkeep">

</file>

<file path="test/security/.gitkeep">

</file>

<file path="test/security/security_baseline_test.go">
package security_test

import "testing"

func TestSecurityBaselineExists(t *testing.T) {
	t.Parallel()
}
</file>

