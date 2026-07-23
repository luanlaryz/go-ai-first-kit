# Feature Request Lifecycle

## Objetivo

Este documento descreve como uma pessoa inicia, acompanha, revisa, bloqueia, retoma e aceita uma solicitacao interativa no `{{PROJECT_SLUG}}`.

Ele e o guia do lado humano. O guia de execucao do modo interativo, incluindo a maquina de estados, o gate e os campos de `automation/INTERACTIVE_STATE.json`, esta em `docs/interactive-sdd-autopilot.md`. As fontes normativas sao `AGENTS.md`, `automation/INTERACTIVE_AUTOPILOT.md`, `automation/INTERACTIVE_RUNBOOK.md`, `automation/STOP_CONDITIONS.md` e `skills/23-{{PROJECT_SLUG}}-sdd-autopilot/SKILL.md`.

Os exemplos deste documento sao ilustrativos: eles mostram o formato esperado e nao declaram a existencia de specs, reports ou trilhas historicas neste repositorio.

## 1. Iniciar uma solicitacao

O pedido inicial e intake, nao spec aprovada. Descreva a mudanca observavel, o escopo conhecido e os limites:

```text
Execute o interactive_sdd_autopilot do {{PROJECT_SLUG}} para esta solicitacao:
<descreva a mudanca observavel>

Escopo conhecido:
- <arquivos, modulo ou comportamento esperado>

Fora do escopo:
- <limites que nao devem ser cruzados>
```

Se voce nao souber detalhar escopo ou criterios, envie o que existe: o agente deve normalizar o pedido e perguntar pelo menor conjunto util de informacoes.

Pedido que deve gerar perguntas ou bloqueio, nunca implementacao direta:

```text
Execute o interactive_sdd_autopilot do {{PROJECT_SLUG}} para esta solicitacao:
Melhore o autopilot para ser mais seguro.
```

Esse pedido nao define objetivo observavel, escopo, diagnostico, testes nem criterio de aceite.

## 2. O que conferir no intake

Depois do intake, uma pessoa deve conseguir ler na resposta do agente:

1. o texto original preservado;
2. o tipo do pedido (`feature`, `evolution`, `bug`, `refactor`, `docs` ou `investigation`);
3. objetivo e escopo conhecido;
4. incertezas materiais e perguntas abertas;
5. specs e skills candidatas;
6. a decisao inicial: perguntar, amendar spec existente ou preparar nova trilha dual-spec.

Se o agente pergunta, responda apenas as pendencias materiais. Perguntas sobre comportamento observavel, contrato publico, estrategia de testes e criterio de aceite bloqueiam a trilha ate serem respondidas.

## 3. O que conferir na decisao de spec

A decisao entre `amend` e nova trilha dual-spec segue `skills/05-{{PROJECT_SLUG}}-spec-architect/SKILL.md`.

Para `amend`, confira se o agente registrou as specs existentes lidas, as secoes governantes e o motivo para nao abrir nova trilha.

Para nova trilha, confira se existem, antes de qualquer implementacao:

1. spec de construcao com objetivo, escopo, fora de escopo, comportamento observavel, testes e criterios de aceite;
2. spec de diagnostico com sinais observaveis, modos de falha, comandos, evidencias, classificacao e decisao esperada.

Decisao de spec ausente e motivo de bloqueio, nao de tolerancia.

## 4. O que conferir durante a implementacao

1. Somente arquivos no escopo aprovado foram editados.
2. `pkg/` e `internal/` ficaram intactos quando a trilha for so de governanca, docs ou onboarding.
3. `automation/ROADMAP.json` e `automation/PHASE_STATE.json` nao foram usados para representar a trilha interativa.
4. Testes acompanharam qualquer mudanca de comportamento.
5. Gaps que nao puderam ser fechados foram registrados, nao omitidos.

## 5. O que conferir no report e no gate

O report da etapa e a fonte de verdade para avanco. Classificacoes permitidas: `PASS`, `PARTIAL`, `FAIL` e `BLOCKED`. O agente so avanca com `classification = PASS` e decisao final autorizando.

Revise:

1. o report existe no caminho registrado em `current_request.report_file`;
2. a etapa do report corresponde a `current_request.current_step`;
3. `classification` e `decision` existem e sao explicitas;
4. comandos e testes obrigatorios foram executados ou justificados pela spec de diagnostico;
5. nenhuma stop condition permanece ativa.

Para mudanca documental, o diagnostico tipico inclui validacao JSON dos estados alterados, `go test ./...`, `go test -race ./...` e lints dos arquivos alterados; `make fmt` e OpenAPI podem ser justificados como nao aplicaveis quando nenhum arquivo Go ou HTTP muda.

## 6. Bloqueio e retomada

Um bloqueio legitimo aparece em tres lugares: em `automation/INTERACTIVE_STATE.json` (`blocked = true` com `block_reason` objetivo), no report ou evidencia da trilha, e na resposta final do agente com a acao humana recomendada.

Na retomada, exija que o agente parta dos artefatos versionados — estado, specs da trilha e report do bloqueio — e corrija somente a causa registrada, preservando `retry_count` e historico. Retomada baseada na memoria da conversa nao e aceitavel. O limite e de `2` retries corretivos por etapa; alem disso, a trilha para e espera decisao humana.

Nunca aceite como "solucao" de bloqueio: alterar gate, criterio de aceite, `ROADMAP.json` ou `PHASE_STATE.json`.

## 7. Checklist humano antes de aceitar a conclusao

1. Pedido inicial tratado como intake, nao como spec aprovada.
2. Decisao entre `amend` e nova trilha dual-spec registrada.
3. Spec de construcao e spec de diagnostico lidas e cobrindo a mudanca.
4. Arquivos alterados dentro do escopo aprovado.
5. Comandos e testes executados ou justificados.
6. Report com `classification = PASS` e decisao de conclusao.
7. `automation/INTERACTIVE_STATE.json` coerente com o report (`completed` exige `PASS`).
8. Nenhuma mudanca indevida em `automation/ROADMAP.json` ou `automation/PHASE_STATE.json`.
9. Nenhuma capability de produto criada fora de spec.
10. Bloqueios e retomadas preservados no historico quando existirem.
11. Gaps restantes ausentes ou explicitamente registrados.

## 8. Artefatos para revisar em PR

Quando a trilha virar PR, revise: specs de construcao e diagnostico, `automation/INTERACTIVE_STATE.json`, o report da trilha, docs alteradas e, quando a trilha usar `skills/24-agent-readiness-governance/SKILL.md`, o report de readiness em `docs/reports/agent-readiness/` como evidencia auxiliar — nunca como substituto do gate.
