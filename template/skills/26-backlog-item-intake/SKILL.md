---
name: backlog-item-intake
description: Govern {{PROJECT_SLUG}} backlog intake. Use when turning requests, gaps, bugs, diagnostics, recommendations or ideas into canonical backlog items in docs/backlog/Backlog.md with SDD classification, evidence, implementation planning and paired autopilot prompts.
---

# Backlog Item Intake

Goal: transformar entradas brutas em itens governados de backlog do `{{PROJECT_SLUG}}`, sem abrir capability fora de spec e sem tratar conversa, plano ou diagnóstico histórico como autorização de implementação.

## When to Use

Use esta skill quando o usuário pedir para:

- registrar pedido, gap, bug, recomendação, diagnóstico ou ideia no backlog;
- transformar achados de report em itens priorizáveis;
- triagear candidatos antes de abrir trilha SDD;
- criar prompts autopilot para implementação e diagnóstico de um item;
- atualizar `docs/backlog/Backlog.md`.

Não use para implementar o item. Esta skill termina no backlog governado e nos prompts/planos necessários para uma trilha posterior.

## Required Reading

Antes de editar backlog, leia:

1. `AGENTS.md`
2. `skills/00-skill-index/SKILL.md`
3. `skills/05-{{PROJECT_SLUG}}-spec-architect/SKILL.md`
4. `skills/23-{{PROJECT_SLUG}}-sdd-autopilot/SKILL.md`
5. `skills/25-{{PROJECT_SLUG}}-ultra-rigid-sdd/SKILL.md`
6. `docs/backlog/Backlog.md`
7. `references/backlog-item-template.md`
8. `references/backlog-classification-rules.md`
9. `references/autopilot-prompt-template.md`
10. `references/backlog-update-procedure.md`

Leia também specs, reports e docs citados pela entrada. Se a entrada vier de fonte histórica, leia a fonte antes de promover qualquer candidato.

## Non-Negotiables

1. `docs/backlog/Backlog.md` é o backlog canônico.
2. Não invente item real quando só houver estrutura inicial.
3. Não migre lotes históricos automaticamente para `BLG-*`.
4. Não use backlog como substituto de spec, diagnóstico, report ou aprovação humana.
5. Todo item deve registrar evidência; quando faltar, escreva `evidência não fornecida`.
6. Todo item implementável deve declarar spec governante existente ou `SPEC_GAP_BLOCKED`.
7. Toda nova trilha futura deve exigir spec de construção e spec de diagnóstico.
8. Não altere `automation/ROADMAP.json` nem `automation/PHASE_STATE.json` para representar backlog.
9. Não crie capability de produto, {{UPSTREAM_OPS_NAME}}, hosted service, Admin UI, billing, quotas, RBAC ou control plane.
10. Não altere `pkg/*`, `internal/*` ou API pública durante intake de backlog.

## Intake Workflow

### 1. Normalize a entrada

Registre:

- texto original ou fonte;
- tipo: `feature`, `bug`, `docs`, `refactor`, `diagnosis_gap`, `parity_gap`, `research` ou `governance`;
- objetivo observável;
- evidência disponível;
- specs candidatas;
- skills candidatas;
- incertezas materiais.

Se objetivo, escopo, evidência ou aceite estiverem materialmente ausentes, classifique como candidato ou bloqueado; não promova para implementação.

### 2. Deduplicate

Antes de criar novo item:

- busque IDs `BLG-[0-9]{4}` em `docs/backlog/Backlog.md`;
- procure título, palavras-chave e specs relacionadas;
- se já existir item equivalente, atualize o item existente com nova evidência em vez de duplicar.

### 3. Classify

Use `references/backlog-classification-rules.md` para preencher:

- status;
- prioridade;
- severidade;
- valor;
- complexidade;
- risco;
- tipo de spec necessária;
- impacto de API pública;
- risco de breaking change;
- decisão de escopo.

### 4. Persist

Use `references/backlog-update-procedure.md`.

O próximo ID deve ser o maior `BLG-NNNN` existente + 1. Se nenhum item existir, comece em `BLG-0001`.

### 5. Add implementation plan

Cada item aprovado para backlog deve ter plano resumido:

- specs a ler/criar;
- arquivos prováveis;
- passos de implementação;
- testes e checks;
- docs/reports esperados;
- stop conditions.

### 6. Add paired autopilot prompts

Para itens não bloqueados por spec gap, inclua dois prompts no item:

- `Prompt autopilot SDD - implementação`
- `Prompt autopilot SDD - diagnóstico`

Use `references/autopilot-prompt-template.md`. Prompts devem ser autocontidos e rejeitar implementação fora de spec.

## Backlog Statuses

- `candidate`: entrada registrada, ainda sem triagem completa.
- `ready_for_spec`: recorte suficiente para criar/amendar specs.
- `ready_for_implementation`: specs governantes existem e diagnóstico está definido.
- `blocked`: falta decisão, evidência, spec ou escopo.
- `in_progress`: trilha SDD em execução fora desta skill.
- `done`: item concluído com report `PASS`.
- `rejected`: fora de escopo, duplicado ou incompatível com non-goals.

## Final Response Format

Sempre responda com:

- objetivo
- specs lidas
- skills aplicadas
- arquivos alterados
- comandos executados
- testes executados
- classificação
- decisão final
- gaps restantes
