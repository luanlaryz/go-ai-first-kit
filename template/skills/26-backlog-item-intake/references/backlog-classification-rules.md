# Backlog Classification Rules

Use estas regras para classificar itens antes de persistir ou atualizar `docs/backlog/Backlog.md`.

## Status

- `candidate`: há fonte ou ideia, mas falta triagem completa.
- `ready_for_spec`: o recorte é claro o bastante para criar ou amendar specs.
- `ready_for_implementation`: specs de construção e diagnóstico existem e cobrem o item.
- `blocked`: falta spec, evidência, decisão humana, acesso ou escopo material.
- `in_progress`: uma trilha SDD real está em execução.
- `done`: há report com `classification = PASS` e decisão final de conclusão.
- `rejected`: item duplicado, fora do escopo ou incompatível com non-goals.

## Tipo

- `feature`: nova capacidade de framework.
- `bug`: comportamento existente diverge de spec, teste ou contrato.
- `docs`: documentação, report, runbook, skill ou regra.
- `refactor`: reorganização sem mudança observável.
- `diagnosis_gap`: lacuna descoberta por report, auditoria ou check.
- `parity_gap`: diferença relevante contra {{UPSTREAM_NAME}} ou referência autorizada.
- `research`: investigação ainda sem decisão de implementação.
- `governance`: mudança em processo SDD, backlog, rules, prompts ou gates.

## Prioridade

- `P0`: bloqueia release, gate SDD, segurança de escopo ou capacidade central aprovada.
- `P1`: destrava trilha importante ou corrige desvio relevante de spec.
- `P2`: melhoria útil, não bloqueante, com valor claro.
- `P3`: baixa urgência, pesquisa ou limpeza oportunista.

Prioridade não deve ser herdada automaticamente de fonte histórica. Recalcule contra specs atuais e escopo do `{{PROJECT_SLUG}}`.

## Severidade

- `critical`: quebra contrato público, segurança de escopo, release gate ou integridade SDD.
- `high`: falha funcional relevante ou risco alto de drift.
- `medium`: lacuna verificável com workaround ou impacto limitado.
- `low`: melhoria pequena, documentação ou ergonomia.
- `none`: ideia, pesquisa ou item sem defeito associado.

## Valor

- `high`: reduz risco de implementação errada, destrava fase/trilha ou melhora paridade core.
- `medium`: melhora descoberta, manutenção, diagnósticos ou cobertura.
- `low`: valor incremental sem impacto imediato.

## Complexidade

- `S`: documentação ou ajuste local simples.
- `M`: toca múltiplos docs, specs ou testes, sem API pública.
- `L`: exige dual-spec, múltiplos módulos ou contrato público.
- `XL`: exige desenho faseado, migração, compatibilidade ou decisão humana.

## Decisão de escopo

- `in_scope`: cabe no core/framework Go e possui cobertura normativa suficiente.
- `spec_gap_blocked`: pode caber, mas falta spec de construção ou diagnóstico.
- `consumer_app_responsibility`: pertence à aplicação consumidora, não ao framework.
- `out_of_scope`: conflita com `specs/001-non-goals.md` ou {{UPSTREAM_OPS_NAME}}.
- `duplicate`: já existe item equivalente.

## Impacto em API pública

- `none`: não toca `pkg/*`, exemplos públicos nem contratos observáveis.
- `additive`: adiciona contrato público compatível, com spec e testes.
- `breaking`: quebra contrato ou comportamento público; exige aprovação explícita.
- `unknown`: impacto ainda não determinado; item deve ficar `blocked` ou `ready_for_spec`.

## Regras de bloqueio imediato

Classifique como `blocked` ou `rejected` quando:

- faltar evidência material;
- faltar spec governante para comportamento novo;
- faltar spec de diagnóstico para nova trilha;
- a proposta exigir {{UPSTREAM_OPS_NAME}}, hosted service, Admin UI, billing, quotas, RBAC ou control plane;
- a proposta depender de `automation/ROADMAP.json` ou `automation/PHASE_STATE.json` como estado do backlog;
- a proposta exigir alteração de API pública sem spec aprovada;
- houver conflito entre specs sem decisão humana.

## Regra para fontes históricas auxiliares

Itens oriundos de fontes auxiliares em `docs/backlog/**` (planos importados, backlogs herdados, relatórios de pesquisa) devem permanecer como candidatos até triagem individual. Não converta lotes históricos em `BLG-*` de uma vez.

Ao triagear um candidato histórico:

1. leia a fonte histórica e o report de validação aplicável;
2. filtre contra `specs/001-non-goals.md`;
3. confirme se o recorte é framework-level, local-first e sem {{UPSTREAM_OPS_NAME}};
4. registre specs necessárias;
5. só então crie ou atualize item canônico.
