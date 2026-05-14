# Ultra-Rigid Output Contract

Use this reference whenever generating Cursor/Codex/autopilot prompts, specs, diagnosis specs, or audit reports for {{PROJECT_SLUG}}.

## Required response shape for an autopilot prompt set

For each phase, output exactly:

1. `fase <n> - <title>`
2. objective
3. locked scope
4. prohibited scope
5. implementation spec file path
6. implementation prompt
7. implementation execution prompt
8. diagnosis spec file path
9. diagnosis spec prompt
10. diagnosis execution prompt
11. gate language
12. rejection criteria

## Required prompt sections

Every prompt must include:

- `TAREFA`
- `ARQUIVO OBRIGATORIO` or `SPEC DE REFERENCIA`
- `OBJETIVO`
- `CONTEXTO OBRIGATORIO`
- `ESCOPO PERMITIDO`
- `ESCOPO PROIBIDO`
- `ARTEFATOS A REVISAR` or `ARTEFATOS A ALTERAR`
- `TESTES/GATES OBRIGATORIOS`
- `FORMATO DE SAIDA OBRIGATORIO`
- `CRITERIOS DE REJEICAO`

## Required implementation response shape

For implementation prompts, require this exact final structure:

```markdown
## 1. Escopo executado
- o que foi alterado
- o que foi deliberadamente nao alterado

## 2. Arquivos alterados
- lista completa

## 3. Implementacao realizada
- problema
- arquivo(s)
- solucao aplicada

## 4. Validacao executada
- comandos rodados
- resultado de cada comando

## 5. Confirmacao de escopo
- por que nao abriu capability fora do escopo

## 6. Limitacoes remanescentes
- gaps reais, se houver
```

## Required diagnosis response shape

For diagnosis execution prompts, require this exact final structure:

```markdown
## 1. Classificacao final
- PASS, PARTIAL, FAIL ou BLOCKED

## 2. Decisao binaria
- READY/NOT READY language appropriate to the phase

## 3. Principais evidencias
- lista objetiva

## 4. Gaps encontrados
- gaps reais somente

## 5. Confirmacao de boundary
- framework vs application boundary, if relevant
```

## Classification rules

Use these meanings consistently:

- `PASS`: all mandatory axes satisfied with evidence and no blocker.
- `PARTIAL`: meaningful progress, but one or more required axes remain incomplete or ambiguous.
- `FAIL`: required work was not completed or scope was violated.
- `BLOCKED`: evidence could not be collected, commands could not run, or the repo state cannot be verified.

Never classify `PASS` when required commands were not run unless the phase explicitly allows a documentation-only audit without commands.

## Capability expansion guard

If a task is reconciliation, polish, confirmation, diagnosis, versioning governance, or documentation alignment, require the agent to state:

> Esta fase nao abre capability nova.

If a task truly needs a new capability, require explicit user authorization or a new feature phase.
