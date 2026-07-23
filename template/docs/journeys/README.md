# Jornadas do {{PROJECT_TITLE}}

Guias humanos passo a passo. Cada jornada aponta para as fontes normativas em vez de duplicá-las, e termina com a evidência esperada, as condições de bloqueio e a próxima ação humana.

| Jornada | Quando usar | Fontes normativas |
| --- | --- | --- |
| [01 — Primeira contribuição](01-primeira-contribuicao.md) | Acabou de clonar ou gerar o projeto e vai fazer a primeira mudança | [AGENTS.md](../../AGENTS.md), [skills/00-skill-index/SKILL.md](../../skills/00-skill-index/SKILL.md) |
| [02 — Fase do roadmap](02-fase-do-roadmap.md) | Vai executar ou retomar uma fase do `phase_autopilot` | [automation/AUTOPILOT.md](../../automation/AUTOPILOT.md), [automation/RUNBOOK.md](../../automation/RUNBOOK.md) |
| [03 — Backlog e trilha SDD](03-backlog-e-trilha-sdd.md) | Tem um pedido, gap ou bug para registrar e transformar em trilha governada | [skills/26-backlog-item-intake/SKILL.md](../../skills/26-backlog-item-intake/SKILL.md), [interactive-sdd-autopilot.md](../interactive-sdd-autopilot.md) |
| [04 — Release e decisões](04-release-e-decisoes.md) | Vai recomendar release, atualizar changelog ou registrar ADR | [release-versioning-policy.md](../release-versioning-policy.md), [skills/22-release-versioning-governance/SKILL.md](../../skills/22-release-versioning-governance/SKILL.md) |

Regras que valem para todas as jornadas:

- O report da etapa é a fonte de verdade para avanço; testes verdes ou narrativa não satisfazem gate.
- Pare imediatamente em qualquer stop condition de [automation/STOP_CONDITIONS.md](../../automation/STOP_CONDITIONS.md).
- Specs ainda inexistentes são candidatas a criar, nunca leitura obrigatória antecipada.
