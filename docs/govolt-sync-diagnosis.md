# Diagnóstico de sincronização govolt → go-ai-first-kit

Fonte upstream: `github.com/inventa-shop/govolt` (default branch `main`, último push observado 2026-05-29).
Baseline do kit: extraído do govolt no commit `4085345` (`chore: bootstrap go-ai-first-kit extracted from govolt`).
Método: comparação evidence-first via árvore (`gh api repos/inventa-shop/govolt/git/trees/main?recursive=1`) e conteúdo bruto (`gh api .../contents/<path> -H "Accept: application/vnd.github.raw"`) contra `template/` e `internal/diagnose`.

Este relatório é a fonte de verdade do que foi trazido para o kit nesta rodada e do que foi deliberadamente deixado de fora.

## 1. Resumo executivo

Desde o bootstrap, a governança AI-first do govolt evoluiu em quatro frentes:

- A) Backlog governado (camada nova, ausente no kit).
- B) Releases / changelog / ADR (camada nova, ausente no kit).
- C) Quality gate (já em paridade no kit, exceto GitHub Pages).
- D) Arquitetura (já em paridade no kit; ganha o sistema de ADRs da frente B).

Decisões aplicadas: integração completa (template + AGENTS/índice + motor `gakit diagnose` + regeneração de `prompt-bootstrap` e embed); GitHub Pages excluído como específico do govolt.

## 2. Diff estrutural (governança AI-first)

### 2.1 Novo no govolt, ausente no kit (trazer)

- `skills/26-backlog-item-intake/` (`SKILL.md`, `agents/openai.yaml`, `references/backlog-item-template.md`, `references/backlog-classification-rules.md`, `references/backlog-update-procedure.md`, `references/autopilot-prompt-template.md`).
- `docs/backlog/Backlog.md` (backlog canônico com itens `BLG-NNNN`, lifecycle e classificação).
- `docs/decisions/` (sistema de ADRs: `README.md` MADR-like + ADRs `0010..0019`).
- `docs/release-versioning-policy.md`, `docs/release-notes-policy.md`, `docs/release-checklist.md`.
- `docs/plans/` e `docs/reports/` como diretórios de artefato versionado (o kit só tinha `docs/reports/agent-readiness/.gitkeep`).
- `CHANGELOG.md` Keep a Changelog 1.1.0 + SemVer (o `template/CHANGELOG.md` do kit é mínimo).

### 2.2 Novo no govolt, deliberadamente NÃO trazido (descartar)

- `.github/workflows/pages.yml` + `scripts/check-pages-site.sh` + passo `make check-pages-site`: publicação GitHub Pages ligada à fase 42 (techmeeting) do govolt. Específico do govolt; fora do escopo de um kit AI-first genérico.
- `.cursor/plans/*.plan.md`: artefatos do plan mode do Cursor (trabalho em andamento), não governança.
- Conteúdo de domínio do govolt: `docs/backlog/v2/**` (LiteLLM v2, `presentation_v2`), ADRs `0010..0019`, specs `030..831`, itens `BLG-0002..0017`, gaps `G0-G17`.

### 2.3 Compartilhado com mudança de conteúdo (atualizar)

- `AGENTS.md`: govolt adicionou 2 regras de backlog (uma em "Source of truth", uma em "Required execution workflow").
- `skills/00-skill-index/SKILL.md`: govolt adicionou a linha `26-backlog-item-intake` na tabela e um exemplo de roteamento de backlog.

## 3. Matriz de paridade — quality gate (frente C)

Comparação `template/` do kit vs govolt (idêntico salvo indicado):

- `make fmt-check` — em paridade.
- `make lint` (staticcheck) — em paridade.
- `make vet` — em paridade.
- `make check-compliance` (`scripts/check-compliance.sh`) — em paridade.
- `make check-pr-body` (`scripts/check-pr-body.sh`) — em paridade.
- `make test-security` = `security-tests` + `secrets-check` + `vulncheck` (govulncheck) — em paridade.
- `make check-file-size` (limites 700/1000 + `.file-size-exceptions`) — em paridade.
- `make test` / `make race` — em paridade (passos `test` e `race` no CI).
- `make coverage` — alvo do Makefile em paridade, mas NÃO é passo do `ci.yml` em nenhum dos dois (sinal local, não gate de CI).
- `.pre-commit-config.yaml` (hooks fmt-fix/lint/vet/check-compliance/check-file-size/secrets-check) — em paridade.
- `.github/workflows/ci.yml` — em paridade EXCETO o passo `Check Pages site` (`make check-pages-site`), excluído por decisão.
- `scripts/check-secrets.sh`, `scripts/check-pr-body.sh`, `scripts/create-pr-from-template.sh` — byte-idênticos ao govolt (confirmado por `diff`).

Drift encontrado e corrigido:

- `.editorconfig` estava simplificado no kit (faltavam regras para Go/Makefile/yaml/sh; `indent_size` divergente). Alinhado byte-a-byte ao govolt.
- `.file-size-exceptions` não existe no govolt (o kit o mantém como arquivo de exceções com header; `check-file-size.sh` lida com sua ausência). Mantido como adição benigna do kit.

Conclusão da frente C: nenhuma ferramenta de quality gate nova precisa ser criada; verificada a paridade byte-a-byte, corrigido o drift do `.editorconfig` e garantida a ausência de Pages.

## 4. Matriz de paridade — arquitetura (frente D)

- `specs/020-repository-architecture.md` (fronteiras `pkg/` vs `internal/`, ponte `pkg/app`, limites de tamanho de arquivo) — em paridade.
- `skills/01-hexagonal-architecture`, `skills/16-object-calisthenics`, `skills/17-solid-go-ports` — em paridade.
- `internal/diagnose/categories/hexagonal.go` (checker de arquitetura no `gakit diagnose`) — em paridade.
- Novo artefato de arquitetura/governança: sistema de ADRs (`docs/decisions/`), trazido pela frente B.

## 5. Tabela de decisão por item

- skill `26-backlog-item-intake` → portar-parametrizado (`govolt`→`{{PROJECT_SLUG}}`, `Voltagent`→`{{UPSTREAM_NAME}}`, `VoltOps`→`{{UPSTREAM_OPS_NAME}}`; remover regra LiteLLM-específica, generalizar fontes históricas).
- `docs/backlog/Backlog.md` → portar-parametrizado, vazio de itens reais (estrutura canônica com `# {{PROJECT_TITLE}} Backlog`).
- `docs/decisions/README.md` → portar-parametrizado, índice vazio, faixa `0001-0009` reservada; remover índice govolt `0010..0019`.
- ADR seed/template → criar genérico (`0001-record-architecture-decisions.md`), não copiar ADRs do govolt.
- `docs/release-versioning-policy.md` / `release-notes-policy.md` / `release-checklist.md` → portar-parametrizado; ancorar em skills 15/21/22 e `specs/020`; remover números de spec govolt (460/461, 240/241).
- `CHANGELOG.md` → reescrever como Keep a Changelog 1.1.0 + SemVer com footer usando `{{MODULE_PATH}}`.
- `docs/plans/` e `docs/reports/` → criar `.gitkeep` + documentar convenção `NN-*-implementation-plan.md` / `NN-*-implementation-report.md`.
- `AGENTS.md` / `skills/00-skill-index` → atualizar (backlog + ADR/release na source of truth + skill mapping).
- GitHub Pages (`pages.yml`, `check-pages-site.sh`, `make check-pages-site`) → divergência intencional (não portar).
- `.cursor/plans/*` e `docs/backlog/v2/**` → descartar (govolt-specific).

## 6. Impacto no motor `gakit diagnose`

Para os novos sinais refletirem no diagnóstico (e não só no template), sem quebrar invariantes de teste:

- `internal/diagnose/engine_test.go::TestRunAgainstTemplate` trava exatamente 7 categorias e score global `>= 65`: não criar categoria nova; dobrar sinais nas categorias existentes.
- `internal/diagnose/categories/governance_test.go::TestRequiredComplianceFilesMatchShellScript` compara `RequiredComplianceFiles` com `template/scripts/check-compliance.sh` por igualdade exata: backlog/ADR/release NÃO viram gate duro de compliance (ficam como sinal de qualidade), preservando os dois lados intactos.
- Sinais adicionados: backlog (pilar AI-first), ADR + release policy (pilar Governança), changelog Keep-a-Changelog (pilar Documentação), cada um como `qualityCheck` adicional com incremento do total correspondente.

## 7. Divergências intencionais registradas

- GitHub Pages do govolt não entra no kit.
- Conteúdo de domínio do govolt (LiteLLM, VoltAgent, ADRs 0010-0019, specs e itens BLG reais) não entra no template; apenas as estruturas/processos parametrizados entram.
