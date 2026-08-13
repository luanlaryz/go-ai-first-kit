# Diagnóstico de sincronização sac-agents → go-ai-first-kit

Fonte downstream: `github.com/inventa-shop/sac-agents`, `origin/main` (`d5b89b19e718`, 2026-08-13).
Direção: **downstream → kit**. O `sac-agents` foi gerado a partir deste kit e evoluiu camadas de governança que o `template/` não tinha; esta rodada devolve ao kit o que é genérico.
Método: leitura de artefatos via `git show origin/main:<path>`. O checkout local do downstream estava em branch de feature atrasada em relação à `main` e **não** é fonte válida — a primeira varredura desta rodada perdeu o eixo de sessões paralelas exatamente por isso.

O pin operacional da rodada está em [`docs/sac-agents-sync-baseline.json`](sac-agents-sync-baseline.json).

## 1. Resumo executivo

O downstream evoluiu em cinco frentes desde a geração:

- A) Segurança de acesso a credencial e nuvem por hooks fail-closed (camada nova, ausente no kit).
- B) Governança de planos: nenhum plano executa sem review registrado e verificável (camada nova).
- C) Governed change: mudança em escopo governado exige backlog + dual-spec + plano revisado, com gate no PR (camada nova).
- D) Coordenação de sessões paralelas: reserva tardia de IDs e gate de colisão contra três fontes (camada nova).
- E) Gates de qualidade de design Go, observabilidade, fronteiras hexagonais e limpeza de pacote (camada nova; o kit tinha apenas tamanho de arquivo).

Decisão aplicada: integração completa das cinco frentes em `template/`, com generalização por placeholders e adaptação de identificadores (`TASK-NNN` → `BLG-NNNN`). Conteúdo de domínio de atendimento ao cliente permanece fora.

## 2. Remapeamento de skills

O downstream numera skills até 35; o kit vai até 26. As portadas entram por append, sem renumerar as existentes:

| Downstream | Kit | Skill |
|---|---|---|
| `28-dual-model-plan-review` | `27-dual-model-plan-review` | Gate de review de planos |
| `29-plan-review-autopilot` | `28-plan-review-autopilot` | Preparação do gate sem burlá-lo |
| `33-governed-change-workflow` | `29-governed-change-workflow` | Fluxo mestre de mudança governada |
| `22-third-party-service-integrations` | `30-third-party-service-integrations` | Integrações externas via ports |
| `35-parallel-session-coordination` | `31-parallel-session-coordination` | Coordenação multi-sessão |

Referências cruzadas dentro dos artefatos portados foram reescritas para a numeração do kit. Nenhum artefato portado pode citar número de skill do downstream.

## 3. Decisão por artefato

### 3.1 Portado com generalização

Hooks (`template/.cursor/hooks/`):

- `agent-cloud-access-guard.{py,sh}` + selftest: classificação por alvo (conta, profile, contexto, endpoint), não por verbo. Contas e marcadores do downstream saíram do código para `cloud-access-config.json`; sem configuração, todo alvo resolve `unknown` e o comportamento é fail-safe (mutação nega, leitura pergunta).
- `env-read-guard.{py,sh}` + selftest: exige aprovação humana antes de ler `.env`, liberando `.env.example`/`.sample`/`.template`.
- `governed-change-guard.{py,sh}`: pergunta antes de commit/push/PR quando o diff toca escopo governado sem trilha ativa.
- `parallel-collision-guard.{py,sh}`: bloqueia publicação com ID sequencial já reservado por outra sessão.

Rules (`template/.cursor/rules/`): `agent-aws-access`, `dual-model-plan-review`, `plan-review-autopilot`, `governed-change-enforcement`, `sdd-non-negotiable`, `pre-pr-before-push`, `pr-body-gate`, `parallel-session-safety`, `go-design-size`, `architecture-boundaries`, `observability`, `package-clean`, `http-security-tenant`, `third-party-service-integrations`, `docs-sync-maintenance`, `ci-no-silent-skip`, `test-suite-session-isolation`, `regression-coverage-required`.

Scripts (`template/scripts/`): `check-agent-hooks.sh`, `check-governed-change.sh`, `check-sdd-compliance.sh`, `verify-plan-review.sh`, `review-plan.py`, `resolve-open-pr-body.sh`, `resolve-pr-body-by-number.sh`, `check-parallel-collision.sh` (+ selftest), `check-architecture-boundaries.sh`, `check-observability.sh`, `check-package-clean.sh`, `check-function-size.sh`, `check-port-size.sh`, `check-ignored-errors.sh`, `check-public-route-security.sh`, `check-test-isolation.sh`, `env-keys.sh`.

Plan templates (`template/.cursor/plans/templates/`): seis modelos genéricos (`adapter-port`, `openapi-change`, `migration-change`, `runtime-change`, `worker-change`, `golden-case`).

### 3.2 Adaptado, não copiado

Estas rules citavam gates que não existem no kit; a versão portada só promete o que o kit entrega:

- `ci-no-silent-skip`: no downstream cita jobs de admin UI, testes de carga k6 e harness de paridade. No kit, cita apenas `guardrails`, `security`, `race` e a sincronia `pre-pr` ↔ CI.
- `regression-coverage-required` e `test-suite-session-isolation`: o prefixo protegido `scripts/dev-regression/` vira a árvore de testes do kit (`test/`), com `check-test-isolation.sh` parametrizado por variável.
- `pre-pr-before-push`: removidas as dependências de targets inexistentes no kit (validação de paridade e auditoria de acoplamento a vendor).
- `docs-sync-maintenance`: no downstream sincroniza dois documentos de onboarding específicos; no kit aponta para `docs/journeys/` e `docs/ai/capabilities.md`.
- Protocolo F da skill de sessões paralelas: reescrito para backlog single-file (ver seção 5).

### 3.3 Não portado

- Skill do runtime de agentes proprietário: o slot 05 do kit já é ocupado pelo spec-architect e o runtime é escolha de cada projeto.
- Skill e rule de padrões de admin frontend: acopladas ao portal interno da organização.
- Skill de migração da aplicação legada: trilha de projeto, não capacidade de kit.
- Skill e rule do eixo de resolução de intenção por domínio: lógica de atendimento ao cliente.
- Rules `sac-agents-baseline.mdc` e `sac-agents.mdc`: já cobertas por `template/.cursor/rules/{{PROJECT_SLUG}}.mdc`.
- `check-kubectl-context.sh`: a decisão de contexto explícito já é aplicada pelo `agent-cloud-access-guard` (nega `kubectl` sem `--context`); um gate separado seria redundante.
- `check-autopilot-structure.sh`: acoplado ao backlog um-arquivo-por-item e às trilhas do downstream. O kit cobre estrutura com `check-compliance.sh`.
- `check-backlog-consistency.sh`, `generate-backlog.py`, `split-backlog.py`: pertencem ao modelo um-arquivo-por-item, fora do escopo desta rodada.
- Gates de admin UI, testes de carga, publicação de site e harness de paridade: específicos do downstream.
- Plan templates `tool-port` e `flow-port`: amarrados à migração da aplicação legada.

### 3.4 Referências normativas: skills e runbooks, não specs numeradas

As rules do downstream citam specs numeradas locais (`610`/`611` para governed change, `886`/`887` para hardening de credencial, `guardrails/001` para os limites não-negociáveis).

O kit **não** ganhou equivalentes numerados. Motivo: as rules existentes do kit (`{{PROJECT_SLUG}}.mdc` e as cinco modulares de autopilot) citam `AGENTS.md`, skills e arquivos de `automation/` — nunca specs numeradas. Criar oito specs só para servir de âncora de rule iria contra essa convenção, inflaria o starter e criaria mais superfície para o teste de integridade de referências quebrar.

As rules portadas apontam então para o que o kit já sustenta: a skill correspondente, o runbook (`docs/runbooks/agent-cloud-access.md`) e o gate executável. `template/specs/` segue com as seis specs originais, e a afirmação em `AGENTS.md` sobre quais specs têm leitura exigida por caminho fixo permanece verdadeira.

Consequência aceita: um projeto que queira spec normativa para esses guardrails cria o par dual-spec no seu próprio repositório. O kit entrega o gate, não a spec do gate.

## 4. Compartilhado com mudança de conteúdo

- `template/AGENTS.md`: novas entradas no mapeamento de skills (27–31) e no workflow de execução obrigatório.
- `template/skills/00-skill-index/SKILL.md`: linhas das skills 27–31 e roteamento por gatilho.
- `template/Makefile`: alvos `pre-pr`, `guardrails`, `verify-plan-reviews`, `review-plan`, `review-plan-turn-based`, `check-agent-hooks`, `check-parallel-collision` e os checks de design.
- `template/.github/PULL_REQUEST_TEMPLATE.md`: seções que o gate de governed change valida.
- `template/.github/workflows/ci.yml`: job de governança PR-scoped e agregação dos guardrails.
- `template/scripts/check-compliance.sh`: novos arquivos obrigatórios.
- `template/CONTRIBUTING.md`: `python3` e `jq` como dependências de desenvolvimento.

## 5. Limitação conhecida: backlog single-file

No downstream, o backlog é `automation/backlog/TASK-NNN.md` (um arquivo por item) com índice gerado. Duas sessões que registram itens diferentes não conflitam, porque escrevem arquivos distintos.

O kit mantém `docs/backlog/Backlog.md` em arquivo único. Consequência: duas sessões que registram itens simultaneamente editam o mesmo arquivo e podem conflitar. Isto é **limitação aceita** desta rodada, não gap silencioso.

Mitigação implementada:

- Item novo sempre ao final da seção, e rebase imediatamente antes do push (skill 31, Protocolo F).
- Conflito preserva os itens de **ambas** as sessões; nunca se resolve descartando o item alheio.
- Colisão de `BLG-NNNN`, `specs/NNN` e `migrations/NNNN` é detectada por `scripts/check-parallel-collision.sh` contra três fontes.

Migração para um-arquivo-por-item muda a estrutura da skill 26 e do próprio `Backlog.md`; fica como trabalho a especificar, não como entrega desta rodada.

## 6. Gaps declarados

**Motor de diagnóstico.** `gakit diagnose` ganhou sinais nesta rodada: `security.go` passou a exigir os guards de credencial, o gate de wiring e o runbook, e a checar `failClosed` em `hooks.json`; `governance.go` passou a exigir o gate de mudança governada, o gate de review de plano e allowlists com contrato de expiração. Não foram adicionados sinais para os gates de design Go (tamanho de função, tamanho de port, erro descartado): eles dependem de interpretar `.guardrails/*.yaml`, o que exige contrato próprio no motor. Fica como gap.

**Suíte de testes do kit.** `go test ./...` no repositório do kit tem uma falha pré-existente e independente deste port: o pacote `template` não compila porque `template/doc.go` declara `package {{PROJECT_SLUG}}`, que só é Go válido depois da renderização. Os testes que governam este port (integridade de referências e atualidade do prompt master) estão no pacote raiz e passam.

**Gate de tamanho de arquivo.** O kit mantém o `scripts/check-file-size.sh` em bash com `.file-size-exceptions`, em vez do subcomando `filesize` do tool Go. Isso deixa dois formatos de allowlist convivendo: o histórico do kit e o `.guardrails/*.yaml` com expiração. Unificar é trabalho a especificar.

## 7. Endurecimentos aplicados após revisão

A revisão do port encontrou bypasses que existiam no artefato de origem e foram corrigidos aqui, não portados:

- **Escopo das seções do PR.** O gate de governed change do downstream faz fallback para o corpo inteiro quando a seção não contém o dado (`section(body, "Backlog") or body`). Consequência: um `BLG` citado em `## Regressao` satisfazia `## Backlog`, apesar da mensagem de erro dizer o contrário. Como as seções são obrigatórias no template de PR, o fallback foi removido e a busca é estritamente escopada.
- **Auto-aprovação por frontmatter.** `review-plan.py` recusa planner igual a reviewer, mas `verify-plan-review.sh` não checava — bastava editar o frontmatter à mão. O verificador agora recusa também.
- **`OPERATOR_FALLBACK` forjado.** Mesma assimetria: o escritor bloqueava alto risco, o verificador só exigia `operator_fallback_reason`. A classificação de risco passou a viver em `scripts/lib/plan_risk.py`, importada pelos dois, para que escritor e verificador não possam divergir.
- **Falso positivo de risco.** A detecção do downstream usa substring: `auth` casava com "author", `secret` com "secretary". Isso inflava risco e bloqueava fallback legítimo. Agora usa fronteira de palavra com alternativas explícitas.
- **Schema do allowlist governado.** O parser Python lia apenas `path` e `expires_at`, aceitando entrada sem `owner`, `justification` ou `ref` — campos que a documentação exige e que o tool Go valida nas outras allowlists. O parser passou a exigir o schema completo.
- **Cobertura de refs remotas.** O gate de colisão extraía só o item de backlog do nome da branch; número de spec e de migration em nome de branch passavam. Os três tipos agora são extraídos.
- **Regressão declarada de forma vaga.** A seção `## Regressao` exigia apenas a substring `test/`. Agora exige `test/<pacote>::<cenário>` concreto.
- **Dependência indireta como adoção.** `check-observability.sh` tratava qualquer menção em `go.mod` como adoção da stack, incluindo `// indirect`. Agora só dependência direta exige wiring.
- **`failClosed` do governed-change guard.** O gate de wiring não exigia a flag para esse hook, embora `hooks.json` a declare. Passou a exigir.

## 8. Procedimento para a próxima rodada

1. Fixar `origin/main` do downstream e confirmar que o checkout comparado é a `main`, não uma branch de feature.
2. Atualizar `docs/sac-agents-sync-baseline.json` com SHA e data.
3. Classificar cada artefato novo: portar com placeholders, adaptar (gate equivalente no kit), ou descartar (domínio).
4. Garantir que nenhuma rule portada cite script, target, spec ou número de skill inexistente no kit.
5. Regenerar `prompt-bootstrap-go-ai-first.md` e rodar `go test ./...` no repo do kit.
