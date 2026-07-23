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
5. `automation/INTERACTIVE_AUTOPILOT.md`;
6. `automation/INTERACTIVE_STATE.json`;
7. `automation/STOP_CONDITIONS.md`.

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

Leia uma spec de construcao ou diagnostico da trilha somente se ela existir.
Durante intake, caminhos de specs ainda inexistentes sao candidatos a criar,
nunca arquivos obrigatorios que ja devam ser lidos. Antes de implementar
comportamento novo, complete e leia a dual-spec.

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
- automation/INTERACTIVE_AUTOPILOT.md
- automation/INTERACTIVE_STATE.json
- automation/STOP_CONDITIONS.md

Comece por intake.
Levante requisitos.
Decida entre amendar spec existente ou abrir nova trilha dual-spec.
Leia specs existentes somente quando forem relevantes e estiverem presentes.
Trate specs novas como candidatas a criar; crie ou refine a dual-spec antes de implementar.
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
