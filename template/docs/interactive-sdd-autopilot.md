# Interactive SDD Autopilot

## Objetivo

O `interactive_sdd_autopilot` e o modo de trabalho para transformar uma solicitacao inicial em uma trilha Spec Driven Development completa dentro do `{{PROJECT_SLUG}}`.

Ele existe para orientar o desenvolvimento do repositorio. Ele nao cria API publica, runtime feature, dashboard, runner hospedado, {{UPSTREAM_OPS_NAME}} ou capability de produto do framework.

Use este guia como onboarding humano. As fontes normativas sao `AGENTS.md`,
as specs existentes e aplicaveis em `specs/`, `automation/INTERACTIVE_AUTOPILOT.md`,
`automation/INTERACTIVE_RUNBOOK.md`, `automation/INTERACTIVE_STATE.json`,
`automation/STOP_CONDITIONS.md` e
`skills/23-{{PROJECT_SLUG}}-sdd-autopilot/SKILL.md`.

No template recem-renderizado, os caminhos de specs garantidamente existentes
sao `specs/000-project-mission.md`, `specs/001-non-goals.md`,
`specs/010-feature-matrix.md`, `specs/020-repository-architecture.md`,
`specs/680-phase-0-bootstrap-foundation.md` e
`specs/681-phase-0-bootstrap-foundation-diagnosis.md`. As duas ultimas
pertencem a fase atual do `phase_autopilot`; uma trilha interativa continua
separada e usa `automation/INTERACTIVE_STATE.json`.

Specs de uma trilha interativa que ainda nao existem sao candidatas a criar,
nunca leitura obrigatoria antecipada.

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
Specs candidatas: specs de baseline relevantes e, se necessario, uma nova dual-spec a criar.
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

Use `amend` quando a mudanca pertence claramente a uma spec existente e aprovada. Registre as secoes governantes antes de editar.

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

## Diagnostico E Gate

Ao terminar uma etapa implementavel, o Cursor deve executar o diagnostico definido pela spec diagnostica. Para uma mudanca documental, a spec pode exigir:

1. validacao JSON dos arquivos de estado alterados, quando aplicavel;
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

## Exemplo Ilustrativo De Conclusao

Este fluxo e apenas ilustrativo e nao declara a existencia de artefatos
historicos:

1. o pedido inicial entra como `docs`, com escopo, fora de escopo e incertezas registrados;
2. o intake decide entre amendar uma spec existente ou criar uma dual-spec;
3. se uma nova trilha for necessaria, as specs de construcao e diagnostico sao criadas e lidas antes da implementacao;
4. a mudanca e executada somente no escopo aprovado;
5. o diagnostico coleta as evidencias e gera o report no caminho definido pela trilha;
6. o report e lido por completo e precisa declarar `classification = PASS` com decisao que autorize conclusao;
7. somente entao `automation/INTERACTIVE_STATE.json` pode registrar `status = "completed"`.

## Exemplo Ilustrativo De Bloqueio E Retomada

1. se a trilha chegar a `report_review` sem report com classificacao e decisao, ela fica bloqueada;
2. o estado preserva `current_step`, `retry_count`, `blocked = true` e um `block_reason` objetivo;
3. nenhum gate, criterio de aceite, `ROADMAP.json` ou `PHASE_STATE.json` pode ser alterado para contornar o bloqueio;
4. a retomada le os artefatos existentes, corrige somente a causa dentro do escopo e reexecuta o diagnostico;
5. o estado so pode desbloquear quando o novo report satisfizer o gate e nenhuma stop condition permanecer ativa.

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
