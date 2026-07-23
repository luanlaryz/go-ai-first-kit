# Jornada 03 — Do backlog à trilha SDD

Para quem tem um pedido, gap, bug, diagnóstico ou ideia e quer transformá-lo em entrega governada. O caminho tem duas etapas independentes: registrar no backlog (sem implementar) e, quando priorizado, abrir a trilha interativa de SDD.

## Parte A — Registrar no backlog

Fonte normativa: [skills/26-backlog-item-intake/SKILL.md](../../skills/26-backlog-item-intake/SKILL.md).

1. Normalize a entrada: texto original, tipo, objetivo observável, evidência disponível, specs e skills candidatas, incertezas.
2. Deduplique: procure itens `BLG-*` equivalentes em [docs/backlog/Backlog.md](../backlog/Backlog.md) antes de criar um novo.
3. Classifique com as regras da skill (status, prioridade, severidade, risco, impacto em API pública).
4. Persista usando o template de item da skill; o primeiro item do projeto recebe `BLG-0001`.
5. Para itens não bloqueados por gap de spec, inclua os dois prompts pareados de autopilot (implementação e diagnóstico).

Regras que não se negociam: o backlog não substitui spec, report nem decisão humana; nenhum item vira implementação sem spec de construção e spec de diagnóstico; intake não altera `pkg/*`, `internal/*` nem API pública.

## Parte B — Abrir a trilha interativa

Fonte normativa: [interactive-sdd-autopilot.md](../interactive-sdd-autopilot.md) (modelos de pedido) e [automation/INTERACTIVE_RUNBOOK.md](../../automation/INTERACTIVE_RUNBOOK.md).

1. Formule o pedido explícito com comportamento esperado, escopo conhecido e fora de escopo.
2. Acompanhe o intake: tipo do pedido, incertezas materiais e decisão inicial. Pedido ambíguo deve gerar perguntas ou bloqueio, não implementação.
3. Confira a decisão de spec: `amend` em spec existente ou nova trilha dual-spec (skill [05](../../skills/05-{{PROJECT_SLUG}}-spec-architect/SKILL.md) governa a decisão). Nova trilha só segue com spec de construção e spec de diagnóstico criadas.
4. Acompanhe implementação e diagnóstico pelo estado em `automation/INTERACTIVE_STATE.json`.
5. Aceite a conclusão apenas com report `classification = PASS` e decisão que autorize encerramento — o checklist humano completo está em [feature-request-lifecycle.md](../feature-request-lifecycle.md).

## Encerramento

- Evidência esperada: item `BLG-*` atualizado (Parte A); dual-spec, report `PASS` e `INTERACTIVE_STATE.json` coerente (Parte B).
- Bloqueia quando: falta decisão de spec, falta resposta humana para requisito material, teste obrigatório falha ou a correção exigiria escopo proibido.
- Próxima ação humana: em bloqueio, ler `block_reason` e o report da etapa, responder a pendência material e retomar a partir dos artefatos — nunca da memória da conversa.
