# AI Development Spec Lineage

Spec governante: `specs/352-ai-historical-spec-convergence.md`
Data da decisao: 2026-04-09

---

## 1. Decisao de convergencia

**Estrategia adotada: substituicao formal (Estrategia A da Spec 352, secao 5).**

As specs historicamente planejadas `310/311` e `330/331` nunca foram criadas como arquivos no repositorio e nunca governaram nenhuma implementacao real. A funcionalidade prevista para essas trilhas foi absorvida por specs posteriores que de fato orientaram a implementacao.

A partir desta decisao:

- `310/311` sao formalmente superadas por `348/349`
- `330/331` sao formalmente superadas por `350/351`
- Nenhum arquivo de backfill retroativo sera criado para 310, 311, 330 ou 331
- A rastreabilidade entre trilha planejada e trilha real e garantida por este documento

Esta decisao nao inventa historico de execucao. As specs 310/311 e 330/331 nunca existiram como arquivos e nunca governaram trabalho passado.

---

## 2. Tabela de rastreabilidade

| Spec planejada (ausente) | Funcionalidade prevista | Superada por (existente) | Evidencia de absorcao |
| --- | --- | --- | --- |
| 310 - AI development enforcement baseline | checker de compliance, PR template, linter real, pre-commit, targets reais de CI | `specs/348-ai-operational-maturity-baseline-recovery.md` (gaps G1-G4) | `scripts/check-compliance.sh`, `Makefile`, `.github/workflows/ci.yml`, `.github/PULL_REQUEST_TEMPLATE.md`, `.pre-commit-config.yaml` |
| 311 - AI development enforcement diagnosis | auditoria da baseline de enforcement | `specs/349-ai-operational-maturity-baseline-recovery-diagnosis.md` (dimensoes D1-D6) | `docs/reports/ai-enforcement-baseline-report.md` |
| 330 - AI safety regression baseline | baseline minima de testes de seguranca, input malicioso, output validation, fail-closed | `specs/350-ai-operational-maturity-expansion.md` (gap G4, secao 5.4) | `test/security/guardrail_security_test.go` |
| 331 - AI safety regression diagnosis | auditoria da baseline de safety | `specs/351-ai-operational-maturity-expansion-diagnosis.md` (dimensao D4) | `docs/reports/ai-expansion-diagnosis-report.md` |

---

## 3. Classificacao das specs da trilha AI (300+)

### Specs governantes (existem e governam implementacao real)

| Spec | Nome | Par diagnostico | Status |
| --- | --- | --- | --- |
| 342 | AI development remediation closure | 343 | existente; D1 e D5 historicamente FAIL; gap normativo de D1 resolvido por esta convergencia |
| 344 | AI development doc coherence cleanup | 345 | existente |
| 346 | AI development doc truthfulness closure | 347 | existente |
| 348 | AI operational maturity baseline recovery | 349 | existente; superou 310/311 |
| 350 | AI operational maturity expansion | 351 | existente; superou 330/331 |
| 352 | AI historical spec convergence | 353 | existente; esta decisao |

### Specs formalmente superadas (nunca existiram como arquivos)

| Spec planejada | Superada por | Justificativa |
| --- | --- | --- |
| 310 | 348 | funcionalidade absorvida; nenhuma implementacao foi guiada por 310 |
| 311 | 349 | funcionalidade absorvida; nenhuma auditoria foi guiada por 311 |
| 330 | 350 | funcionalidade absorvida; nenhuma implementacao foi guiada por 330 |
| 331 | 351 | funcionalidade absorvida; nenhuma auditoria foi guiada por 331 |

### Specs historicas ausentes (sem substituicao formal nesta convergencia)

| Spec planejada | Funcionalidade prevista | Estado |
| --- | --- | --- |
| 300/301 | (referenciadas historicamente sem descricao precisa) | ausentes; sem substituicao formal; tratadas como gap documental |
| 320/321 | (referenciadas historicamente sem descricao precisa) | ausentes; sem substituicao formal; tratadas como gap documental |
| 340/341 | learning loop | ausentes; sem substituicao formal; tratadas como gap documental |

---

## 4. Impacto na Spec 342

A `Spec 342` referencia `specs/310` e `specs/311` como dependencias em sua secao 5.1 (Gap G1) e no prompt de execucao (secao 10). A `Spec 342` tambem referencia `specs/330` e `specs/331` no prompt de execucao.

A partir desta convergencia:

- Toda referencia a 310/311 na Spec 342 deve ser lida como apontando para 348/349
- Toda referencia a 330/331 na Spec 342 deve ser lida como apontando para 350/351
- O FAIL em D1 do relatorio de remediation closure (`docs/reports/ai-remediation-closure-report.md`) refletia o gap normativo de 310/311; esse gap normativo esta agora resolvido pela substituicao formal
- D5 do relatorio de remediation closure (safety gap) permanece inalterado por esta convergencia

---

## 5. Guia para contribuidores futuros

### O que ler para entender a trilha AI

1. `AGENTS.md` — workflow normativo e regras globais
2. `docs/ai/ops-guide.md` — guia operacional com estado atual de assets
3. Este documento (`docs/ai/spec-lineage.md`) — linhagem e rastreabilidade historica
4. As specs governantes listadas na secao 3 deste documento

### Como interpretar specs da trilha AI

- Se a spec existe como arquivo em `specs/`, ela e governante ou historica conforme a secao 3 deste documento
- Se a spec nao existe como arquivo e esta listada como "formalmente superada", sua funcionalidade foi absorvida pela spec indicada na tabela de rastreabilidade
- Se a spec nao existe como arquivo e esta listada como "historica ausente", ela e um gap documental reconhecido e nao governou nenhuma implementacao
- Nunca assumir que uma spec ausente governou implementacao real

### Onde esta o enforcement real

A baseline de enforcement da trilha AI esta materializada em:

- `scripts/check-compliance.sh` — checker semantico
- `.pre-commit-config.yaml` — hooks locais
- `.github/workflows/ci.yml` — pipeline de CI
- `.github/PULL_REQUEST_TEMPLATE.md` — template de PR
- `docs/ai/ai-contribution-contract.md` — contrato de contribuicao
- `docs/ai/compliance-exceptions.md` — caminho formal de excecoes
- `test/security/` — baseline minima de safety

Esses artefatos foram criados sob as specs 348 e 350, nao sob 310 ou 330.
