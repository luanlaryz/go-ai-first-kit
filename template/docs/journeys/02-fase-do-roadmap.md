# Jornada 02 — Executar uma fase do roadmap

Para quem vai operar o `phase_autopilot`: executar a fase atual do roadmap fixo ou retomar após bloqueio. As fontes normativas são [automation/AUTOPILOT.md](../../automation/AUTOPILOT.md) e [automation/RUNBOOK.md](../../automation/RUNBOOK.md); esta jornada só organiza a operação humana.

## 1. Conferir o estado antes de iniciar

Leia `automation/PHASE_STATE.json` e localize em `automation/ROADMAP.json` a entrada cujo `id` é `current_phase`. Anote `spec_file`, `diagnosis_spec_file`, `report_file` e `advance_only_if`. Em um projeto recém-gerado, a fase atual é a fase 0 (bootstrap), governada pelo par de specs 680/681.

Se `blocked = true`, não inicie execução normal: vá direto ao passo 4.

## 2. Iniciar a execução

Cole em um agente o prompt mestre da seção 15 de [automation/RUNBOOK.md](../../automation/RUNBOOK.md). O agente deve executar somente a fase atual: criar ou validar a dual-spec, implementar, diagnosticar, gerar o report no caminho exato de `report_file` e atualizar `PHASE_STATE.json`.

O que observar enquanto acompanha:

- a implementação só começa depois que as duas specs da fase existem e foram lidas;
- o roadmap não pode ser reordenado, e gates não podem ser editados;
- trilhas interativas não entram aqui — `ROADMAP.json` e `PHASE_STATE.json` pertencem só ao modo de fases.

## 3. Conferir o gate

A fase só avança quando todas as condições valem:

1. o report existe no caminho exato de `report_file` e corresponde à fase atual;
2. `classification` e `decision` são exatamente os valores de `advance_only_if`;
3. `go test ./...` e `go test -race ./...` passaram;
4. nenhuma stop condition de [automation/STOP_CONDITIONS.md](../../automation/STOP_CONDITIONS.md) está ativa;
5. `PHASE_STATE.json` foi atualizado de forma coerente.

## 4. Retomar após bloqueio

Use o prompt de retomada da seção 17 de [automation/RUNBOOK.md](../../automation/RUNBOOK.md). Regras que você deve fiscalizar: não resetar roadmap, não zerar `retry_count` sem resolver a causa, não avançar com `blocked = true`, e no máximo `max_retries_per_phase` retries corretivos (definido em `ROADMAP.json`).

## Encerramento

- Evidência esperada: report da fase com classificação e decisão exatas do gate, e `PHASE_STATE.json` apontando a próxima fase (`pending`) ou `finished` na última.
- Bloqueia quando: report ausente ou inconsistente, teste obrigatório falho, retries excedidos ou decisão humana de arquitetura pendente.
- Próxima ação humana: em bloqueio persistente, ler o motivo em `block_reason`, decidir a correção de escopo e, se a causa for um pedido novo, encaminhá-lo pela [jornada 03](03-backlog-e-trilha-sdd.md) em vez de alterar o roadmap.
