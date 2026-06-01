# Autopilot Prompt Template

Use estes templates dentro de um item `BLG-*` quando o item estiver pronto para abrir trilha SDD posterior.

Não gere prompt de implementação para item `rejected`. Para item `blocked`, gere prompt apenas para desbloqueio de spec/evidência.

## Prompt autopilot SDD - implementação

```md
Execute o `interactive_sdd_autopilot` para implementar o item `<BLG-ID> - <título>` do backlog canônico `docs/backlog/Backlog.md`.

Leia primeiro:
- `AGENTS.md`
- `skills/00-skill-index/SKILL.md`
- `skills/05-{{PROJECT_SLUG}}-spec-architect/SKILL.md`
- `skills/23-{{PROJECT_SLUG}}-sdd-autopilot/SKILL.md`
- `skills/25-{{PROJECT_SLUG}}-ultra-rigid-sdd/SKILL.md`
- `skills/26-backlog-item-intake/SKILL.md`
- `<spec de construção governante>`
- `<spec de diagnóstico governante>`
- `docs/backlog/Backlog.md`

Objetivo:
Implementar somente o comportamento descrito no item `<BLG-ID>`, preservando o escopo de framework Go do `{{PROJECT_SLUG}}`.

Escopo permitido:
- `<listar recorte aprovado do item>`

Escopo proibido:
- Não alterar `automation/ROADMAP.json` nem `automation/PHASE_STATE.json`.
- Não abrir {{UPSTREAM_OPS_NAME}}, hosted service, dashboard, Admin UI, control plane, auth enterprise, RBAC, billing ou quotas.
- Não alterar `pkg/*`, `internal/*` ou API pública sem cobertura explícita da spec governante.
- Não tratar este prompt, conversa ou backlog como substituto de spec.

Entregáveis mínimos:
1. Mudanças de implementação estritamente cobertas por spec.
2. Testes atualizados ou adicionados quando houver mudança de comportamento.
3. Docs/specs atualizadas se a mudança alterar comportamento observável ou API pública.
4. Report diagnóstico da etapa no caminho definido pela spec de diagnóstico.

Validações obrigatórias:
- `<comandos específicos da spec de diagnóstico>`
- `go test ./...`
- `go vet ./...`

Regras:
- Pare se a spec de construção ou diagnóstico estiver ausente, ambígua ou incompatível.
- Pare se qualquer validação obrigatória falhar sem correção possível dentro do escopo.
- Registre gaps, skips e limitações no report.

Formato de saída:
- objetivo
- specs lidas
- skills aplicadas
- arquivos alterados
- comandos executados
- testes executados
- report lido
- classificação
- decisão
- estado atualizado
- gaps restantes
- próxima etapa ou motivo de bloqueio
```

## Prompt autopilot SDD - diagnóstico

```md
Execute o diagnóstico do item `<BLG-ID> - <título>` conforme a spec de diagnóstico governante e o backlog canônico `docs/backlog/Backlog.md`.

Leia primeiro:
- `AGENTS.md`
- `skills/00-skill-index/SKILL.md`
- `skills/23-{{PROJECT_SLUG}}-sdd-autopilot/SKILL.md`
- `skills/25-{{PROJECT_SLUG}}-ultra-rigid-sdd/SKILL.md`
- `skills/26-backlog-item-intake/SKILL.md`
- `<spec de construção governante>`
- `<spec de diagnóstico governante>`
- `docs/backlog/Backlog.md`

Objetivo:
Verificar por evidência se o item `<BLG-ID>` foi atendido sem expandir escopo proibido.

Checklist obrigatório:
1. Confirmar que os arquivos alterados correspondem ao item e às specs.
2. Confirmar que não houve alteração proibida em `automation/ROADMAP.json`, `automation/PHASE_STATE.json`, `pkg/*`, `internal/*` ou API pública sem spec.
3. Executar validações documentais ou de conteúdo aplicáveis.
4. Executar `<comandos específicos da spec de diagnóstico>`.
5. Executar `go test ./...` e `go vet ./...` se o ambiente permitir.
6. Registrar skips com causa técnica.
7. Gerar ou atualizar report diagnóstico no caminho exigido.

Critérios de classificação:
- `PASS`: todos os critérios obrigatórios estão atendidos ou justificados pela spec, sem stop condition ativa.
- `PARTIAL`: há artefato material incompleto, validação não executada sem justificativa suficiente ou gap não bloqueante.
- `FAIL`: escopo proibido foi alterado, spec foi violada ou validação obrigatória falhou.
- `BLOCKED`: falta acesso, decisão humana, spec, evidência ou ambiente para concluir.

Decisão final permitida:
- `concluir solicitação`
- `executar retry corretivo`
- `parar por stop condition`
- `abrir próxima etapa SDD`

Formato de saída:
- objetivo
- specs lidas
- skills aplicadas
- arquivos revisados
- comandos executados
- testes executados
- skips ou limitações
- gaps restantes
- classificação
- decisão final
```

## Prompt de desbloqueio para item bloqueado

```md
Desbloqueie o item `<BLG-ID> - <título>` antes de qualquer implementação.

Tarefa:
Identificar a menor ação necessária para sair de `blocked`: criar/amendar spec de construção, criar/amendar spec de diagnóstico, coletar evidência, decidir escopo ou rejeitar o item.

Saída obrigatória:
- lacuna bloqueante
- ação humana ou artefato necessário
- specs a criar ou amendar
- decisão recomendada: `ready_for_spec`, `ready_for_implementation`, `rejected` ou `blocked`
```
