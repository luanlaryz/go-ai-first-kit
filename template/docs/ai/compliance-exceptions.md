# AI Compliance Exceptions

Este arquivo registra excecoes formais, pequenas e auditaveis para compliance, testes e enforcement da trilha de AI development.
Ele e o unico caminho aceito para excecoes: nenhuma excecao vale sem registro auditavel aqui, e nenhum registro aqui cria bypass silencioso no checker.

## Estado atual

Ha uma excecao ativa para PRs Dependabot de atualizacao de dependencias.

## Formato de registro

Cada excecao ativa deve listar:

- `id`: identificador curto e estavel
- `escopo`: arquivo, check, teste ou fluxo afetado
- `justificativa`: motivo tecnico observavel
- `owner`: responsavel logico pela revisao
- `criterio de revisao`: prazo, trigger ou condicao objetiva de encerramento

## Excecoes ativas

- `id`: `dependabot-pr-body-template`
  - `escopo`: `scripts/check-pr-body.sh` em eventos `pull_request` gerados pelo Dependabot.
  - `justificativa`: Dependabot gera corpo de PR proprio com changelog e comandos operacionais, sem usar `.github/PULL_REQUEST_TEMPLATE.md`. Exigir headings humanos bloquearia a automacao leve de atualizacao de dependencias sem aumentar seguranca ou rastreabilidade do PR automatizado; os demais checks de CI continuam obrigatorios.
  - `owner`: mantenedores do repositorio `{{PROJECT_SLUG}}`.
  - `criterio de revisao`: revisar quando a politica de PR template mudar, quando Dependabot permitir template customizado compativel ou quando a automacao de dependencia for substituida.
