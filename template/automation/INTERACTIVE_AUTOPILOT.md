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
2. `skills/23-{{PROJECT_SLUG}}-sdd-autopilot/SKILL.md`;
3. spec de construcao aprovada da trilha atual, quando existir;
4. spec de diagnostico aprovada da trilha atual, quando existir;
5. `automation/INTERACTIVE_AUTOPILOT.md`;
6. `automation/INTERACTIVE_RUNBOOK.md`;
7. `automation/INTERACTIVE_STATE.json`;
8. `automation/STOP_CONDITIONS.md`;
9. report da etapa atual, quando tiver sido gerado.

Uma spec ainda inexistente e apenas candidata durante intake e decisao de spec.
Ela so passa a ser leitura obrigatoria depois de criada ou aprovada. Nenhuma
implementacao ou diagnostico de comportamento novo pode iniciar sem a dual-spec
existente.

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
