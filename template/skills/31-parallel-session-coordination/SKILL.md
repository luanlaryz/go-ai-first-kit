---
name: 31-parallel-session-coordination
description: Coordinate parallel agent sessions working on different tasks in the same repo. Use before reserving sequential IDs (BLG-NNNN, specs, migrations), creating branches, publishing PRs, touching shared local infra, or whenever another active session is detected - to prevent ID collisions, duplicated branches/PRs, backlog merge conflicts and shared-infra interference.
---

# Parallel Session Coordination

Goal: permitir que varias sessoes de agente trabalhem simultaneamente no mesmo
repositorio, em tarefas diferentes, sem colidir em IDs sequenciais, branches, PRs,
arquivos append-heavy ou infraestrutura local compartilhada.

O desenho aceita que **nao existe exclusao mutua entre sessoes**: a estrategia e
deteccao + reserva tardia + integracao, nunca sobrescrita. Uma sessao que assume
exclusividade produz exatamente os incidentes que esta skill previne.

## When to Use

- Antes de reservar qualquer ID sequencial: `BLG-NNNN`, `specs/NNN` (par dual),
  `migrations/NNNN`, numero de skill, plan-id em `.cursor/plans/` +
  `automation/PLAN_REVIEWS/`, nome de arquivo de corpo de PR.
- Antes de criar branch ou publicar (`git push`, `gh pr create`, `gh pr edit`).
- Antes de subir, derrubar ou migrar infra local compartilhada (banco, cache,
  emuladores, portas de dev server).
- Ao detectar qualquer sinal de sessao paralela (Protocolo A).
- Ao receber push rejeitado por non-fast-forward.

## Required Reading

1. `AGENTS.md`
2. `skills/26-backlog-item-intake/SKILL.md` (identidade `BLG-NNNN` e template)
3. `skills/05-{{PROJECT_SLUG}}-spec-architect/SKILL.md` (numeracao dual-spec)
4. `.cursor/rules/pre-pr-before-push.mdc` (gate `make pre-pr`)
5. `.cursor/rules/pr-body-gate.mdc` (corpo de PR via body-file)

## Non-Negotiables

1. Reserva tardia: ID sequencial e calculado no ultimo momento antes do commit
   que o usa, nunca no planejamento. Todo recalculo usa **sort numerico**
   (lexicografico quebra na virada de casa decimal).
2. Duas fontes na reserva: base recem-fetchada **e** PRs abertos (texto,
   `headRefName` e paths de arquivos). Um ID pode estar consumido num PR ainda
   nao mergeado, invisivel na base.
3. Branch nomeada `{tipo}/blg{NNNN}-<slug>` (`fix/`, `feat/`, `docs/`, `tests/`,
   `chore/`). O ID no nome e a chave de deduplicacao entre sessoes. Renumerou
   antes de publicar? Renomeie a branch.
4. Force push e proibido para resolver colisao. Push rejeitado por
   non-fast-forward significa parar e aplicar o Protocolo D.
5. Trabalho que gera commit nasce em worktree dedicado a partir da base
   recem-fetchada; nunca commitar do workspace principal quando houver trabalho
   em andamento nao relacionado.
6. Conflito no backlog preserva os itens de **ambas** as sessoes; conflito de
   lockfile nao se resolve a mao - rebase e regenere com a toolchain.
7. Infra compartilhada nao se recria, migra ou semeia sem confirmar que nenhuma
   outra sessao a usa; container de outro projeto nunca e derrubado por conflito
   de porta.

## Protocolo A - Deteccao de sessoes paralelas

Rode no inicio da sessao e antes de decisoes de publicacao. Sinais em ordem de
forca:

```bash
git worktree list                       # worktrees ativos: sinal mais direto
gh pr list --state open --json number,title,headRefName,isDraft
git fetch --all -q && git for-each-ref --sort=-committerdate refs/remotes/origin --format='%(committerdate:short) %(refname:short)' | head -10
docker ps                               # sinal auxiliar: infra de pe, nao diz quem usa
```

Qualquer sinal positivo = assumir coordenacao ativa e aplicar os protocolos B-F
integralmente.

## Protocolo B - Reserva tardia de IDs

```bash
git fetch origin main -q

# maior item no backlog (numerico; lexicografico quebra na virada BLG-0999 -> BLG-1000):
git show origin/main:docs/backlog/Backlog.md \
  | rg -o '^### BLG-([0-9]{4})' -r '$1' | sort -n | tail -1

# maior spec e migration na base:
git ls-tree -r --name-only origin/main specs/ | rg -o 'specs/([0-9]{3})' -r '$1' | sort -n | tail -1
git ls-tree --name-only origin/main migrations/ | rg -o 'migrations/([0-9]{4})' -r '$1' | sort -n | tail -1

# IDs citados em PRs abertos (titulo, corpo e branch):
gh pr list --state open --json title,body,headRefName --jq '.[]|"\(.title) \(.body // "") \(.headRefName)"' | rg -o '(BLG-|blg)[0-9]+' | rg -o '[0-9]+' | sort -n | uniq

# specs/migrations tocados por PRs abertos (paths de arquivos):
gh pr list --state open --json number --jq '.[].number' | xargs -I{} gh pr view {} --json files --jq '.files[].path' | rg '^(specs|migrations)/[0-9]+' | sort -u
```

Gate executavel: `scripts/check-parallel-collision.sh` automatiza este protocolo -
detecta os IDs que a branch introduz e cruza as tres fontes. Roda no
`beforeShellExecution` de `git push`/`gh pr create` e em `make guardrails`/`pre-pr`.
Os comandos acima seguem validos para reserva manual e diagnostico.

Limitacao declarada: a varredura de PRs e **heuristica, nao garantia**. Um PR pode
reservar um ID sem cita-lo em titulo, corpo, branch ou paths - por isso o
cruzamento de tres fontes. A colisao residual e absorvida pelo Protocolo D quando
o merge acusar.

Em colisao: renumere, revalide toda referencia cruzada (plano, corpo do PR, specs,
backlog) e **renomeie a branch se ainda nao publicada**. Se ja publicada, registre
a discrepancia nome-da-branch vs ID no corpo do PR.

## Protocolo C - Gate pre-publicacao

Imediatamente antes de `git push` / `gh pr create`, nesta ordem:

```bash
scripts/check-parallel-collision.sh                # IDs + Protocolo C de uma vez
git fetch origin main -q
git ls-remote origin <branch>                      # ja existe? -> Protocolo D
gh pr list --state open --head <branch>            # PR aberto para a branch? (chave objetiva)
# re-executar Protocolo B (IDs ainda livres?)
make pre-pr PR_BODY=<arquivo>                      # exit 0 obrigatorio
```

Depois de criar ou editar o PR:

```bash
gh pr view <N> --json isDraft,state,mergeable      # draft silencioso custa uma rodada
```

## Protocolo D - Integracao pos-colisao

Quando a branch remota ja contem trabalho de outra sessao (push rejeitado ou
`ls-remote` positivo):

1. `git fetch origin <branch>` e compare as solucoes:
   `git diff origin/<branch> HEAD -- <paths>`.
2. Se equivalentes, integre por rebase mantendo so o que agrega:
   `git rebase origin/<branch>` e resolva conflitos preservando a base remota.
3. Prove o fast-forward antes de publicar:
   `git merge-base --is-ancestor origin/<branch> HEAD && echo fast-forward-ok`.
4. Publique sem force e, se necessario, sem mover a branch local:
   `git push origin <sha>:refs/heads/<branch>`.
5. Registre no corpo do PR que houve convergencia de sessoes.

Armadilha documentada: depois de integrar, **nao** rebasear sobre a base - isso
reescreve os commits da outra sessao ja publicados e quebra o fast-forward. O diff
do PR usa merge-base com a base, entao o rebase sobre ela e desnecessario.

## Protocolo E - Infra local compartilhada

Verificacao passiva antes de qualquer acao que altere estado:

```bash
docker ps                                   # o que esta de pe e em quais portas
# versao de schema aplicada no banco compartilhado vs ultima migration do checkout
ls migrations/ | tail -1
```

- Versao de schema no banco alem da ultima migration do checkout = outra sessao
  com migration nova aplicada; **pare e coordene** antes de migrar.
- Proibido recriar, migrar ou semear infra compartilhada sem confirmar que
  nenhuma sessao a usa.
- Porta ocupada por container de outro projeto: use o servico existente quando
  compativel, ou suba isolado com projeto/portas alternativas e banco dedicado.
  Nunca derrube o container alheio.

## Protocolo F - Backlog e arquivos append-heavy

Este repositorio mantem o backlog em **arquivo unico**: `docs/backlog/Backlog.md`.
Isso e mais simples de navegar e mais sujeito a conflito de merge que um modelo de
um arquivo por item. A mitigacao e procedimental:

- Item novo entra **sempre ao final** da secao de itens, e o ID e reservado no
  ultimo momento (Protocolo B).
- Rebase imediatamente antes do push, para que o conflito apareca localmente e
  nao no merge.
- Conflito em `docs/backlog/Backlog.md` resolve-se preservando os itens de
  **ambas** as sessoes. Descartar o item alheio e perda silenciosa de trabalho.
- Se duas sessoes reservaram o mesmo `BLG-NNNN`, renumere o seu (nao o alheio se o
  outro ja publicou), atualize as referencias cruzadas e rode
  `scripts/check-parallel-collision.sh` de novo.

**O caso mais perigoso nao e o conflito: e o merge que o git resolve sozinho.**
Duas PRs que adicionam itens em regioes diferentes do arquivo podem mesclar sem
conflito e produzir dois itens com o mesmo ID. Portanto: PR que toca o backlog
precisa estar atualizada com a base **no momento do merge**, nao apenas quando o
CI rodou. O CI valida o PR isolado; o resultado do merge nao passa por gate.

Para os demais arquivos append-heavy (indice de skills, matriz de features):

- Item novo sempre ao final da secao; rebase imediatamente antes do push;
  conflito resolve-se preservando o conteudo de **ambas** as sessoes.
- Alternativa avaliada e rejeitada: `merge=union` via `.gitattributes` -
  quebraria tabelas e contadores agregados em silencio.

## Do / Don't

- Do: recalcular IDs com sort numerico imediatamente antes do commit.
- Do: usar `gh pr list --state open --head <branch>` como chave de deduplicacao.
- Do: conferir `isDraft` apos criar/editar PR.
- Don't: force push para "destravar" push rejeitado.
- Don't: resolver conflito de backlog descartando o item da outra sessao.
- Don't: editar lockfile a mao em conflito.
- Don't: confiar apenas no texto de PRs abertos para saber IDs reservados.

## Stop Conditions

- Integracao exigiria force push ou descarte de trabalho alheio.
- PR aberto de outra sessao duplica o mesmo objetivo (decisao humana sobre qual
  segue).
- ID em disputa com PR aberto de outra sessao (coordenar antes de renumerar).
- Infra compartilhada em uso quando o trabalho exige recria-la ou migra-la.

## Definition of Done

- Nenhum ID sequencial publicado colide com a base nem com PRs abertos.
- Branch publicada segue `{tipo}/blg{NNNN}-<slug>` com o ID correto.
- Push realizado sem force; colisoes integradas com fast-forward provado.
- `scripts/check-parallel-collision.sh` PASS.
- PR publicado fora de draft, com `make pre-pr` exit 0.

## Final Response Format

Ao aplicar esta skill, registre na entrega: sinais de sessao paralela detectados,
IDs recalculados (valor e fontes), colisoes encontradas e protocolo aplicado,
comandos executados e gaps restantes.
