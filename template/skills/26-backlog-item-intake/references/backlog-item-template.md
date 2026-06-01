# Backlog Item Template

Use este template para cada item persistido em `docs/backlog/Backlog.md`.

## Template

```md
### BLG-0000 - Título curto e verificável

- **Status**: `candidate | ready_for_spec | ready_for_implementation | blocked | in_progress | done | rejected`
- **Tipo**: `feature | bug | docs | refactor | diagnosis_gap | parity_gap | research | governance`
- **Prioridade**: `P0 | P1 | P2 | P3`
- **Severidade**: `critical | high | medium | low | none`
- **Valor**: `high | medium | low`
- **Complexidade**: `S | M | L | XL`
- **Risco**: `high | medium | low`
- **Decisão de escopo**: `in_scope | spec_gap_blocked | consumer_app_responsibility | out_of_scope | duplicate`
- **Spec governante**: `specs/<arquivo>.md#<seção>` ou `SPEC_GAP_BLOCKED`
- **Spec de diagnóstico**: `specs/<arquivo>.md#<seção>` ou `SPEC_GAP_BLOCKED`
- **Skills aplicáveis**: `skills/00-skill-index/SKILL.md`, `skills/<id>/SKILL.md`
- **Fonte**: caminho, issue, report, conversa ou `evidência não fornecida`
- **Evidência**: evidência versionada, comando, report, trecho observado ou `evidência não fornecida`
- **Impacto em API pública**: `none | additive | breaking | unknown`
- **Risco de breaking change**: `none | low | medium | high | unknown`

#### Problema

Descreva o problema observável sem prometer solução.

#### Resultado esperado

Descreva o resultado verificável que tornaria o item concluído.

#### Escopo proposto

- Inclua apenas o recorte mínimo governado.
- Cite limites de arquitetura e comportamento observável.

#### Fora do escopo

- Liste explicitamente {{UPSTREAM_OPS_NAME}}, hosted services, Admin UI, billing, quotas, RBAC, control plane ou outros itens proibidos quando relevantes.

#### Plano de implementação

1. Ler `AGENTS.md`, `skills/00-skill-index/SKILL.md` e specs governantes.
2. Confirmar se specs existentes bastam ou se a trilha exige dual-spec.
3. Implementar apenas o comportamento coberto por spec.
4. Atualizar testes/docs/reports conforme impacto.
5. Executar diagnóstico e registrar report.

#### Testes e validações previstas

- `go test ./...`
- `go vet ./...`
- Checks documentais específicos do item.
- Outros comandos exigidos pela spec de diagnóstico.

#### Stop conditions

- Spec de construção ausente.
- Spec de diagnóstico ausente.
- Evidência material ausente.
- Decisão humana pendente.
- Escopo proibido por `specs/001-non-goals.md`.
- Mudança exigiria `pkg/*`, `internal/*` ou API pública sem spec aprovada.

#### Prompt autopilot SDD - implementação

Cole aqui o prompt de implementação gerado a partir de `references/autopilot-prompt-template.md`.

#### Prompt autopilot SDD - diagnóstico

Cole aqui o prompt de diagnóstico gerado a partir de `references/autopilot-prompt-template.md`.

#### Histórico

- YYYY-MM-DD: item criado a partir de `<fonte>`.
```

## Regras de preenchimento

- Não remova campos obrigatórios.
- Use `evidência não fornecida` quando não houver prova verificável.
- Use `SPEC_GAP_BLOCKED` quando a implementação depender de spec ausente.
- Não inclua prompt executável para item `rejected`.
- Para item `blocked`, o prompt deve começar com a ação de desbloqueio, não com implementação.
