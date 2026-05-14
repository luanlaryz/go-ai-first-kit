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
