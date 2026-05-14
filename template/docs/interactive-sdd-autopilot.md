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
