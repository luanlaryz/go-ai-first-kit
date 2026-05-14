# Spec 680: Phase 0 - Bootstrap Foundation

## 1. Objetivo

Definir a fundação AI-first inicial de `{{PROJECT_SLUG}}`, incluindo governança para agentes, specs, skills, automation dual, docs operacionais, CI e scripts de compliance.

## 2. Motivação

O projeto deve nascer com padrões explícitos para que agentes consigam operar com baixa ambiguidade e alta rastreabilidade.

## 3. Pergunta principal

`{{PROJECT_SLUG}}` possui uma baseline versionada, verificável e executável para desenvolvimento Go assistido por agentes?

## 4. Escopo

### Dentro do escopo

1. Criar `AGENTS.md`.
2. Criar `skills/`.
3. Criar `.cursor/rules/` e hooks.
4. Criar `automation/`.
5. Criar `docs/ai/`.
6. Criar CI, scripts e PR template.

### Fora do escopo

1. Criar feature de domínio do produto.
2. Introduzir serviço hospedado obrigatório.
3. Abrir API pública sem spec específica.

## 5. Requisitos de design

1. A baseline deve ser local-first.
2. A governança deve ser evidence-first.
3. Trilhas novas devem exigir dual-spec.
4. Reports devem usar `PASS`, `PARTIAL`, `FAIL` ou `BLOCKED`.

## 6. Blocos obrigatórios desta fase

Governança para agentes, automação dual, specs base, CI/scripts e documentação humana.

## 7. Requisitos funcionais

Agentes devem saber o que ler, quando perguntar, como diagnosticar e como reportar evidência.

## 8. Decisões obrigatórias de modelagem

`ROADMAP.json` governa roadmap fixo; `INTERACTIVE_STATE.json` governa trilhas interativas; exceções formais vivem em `docs/ai/compliance-exceptions.md`.

## 9. Critérios de aceitação

1. `make check-compliance` passa.
2. `make test` passa.
3. `automation/*.json` é JSON válido.

## 10. Evidência obrigatória

Arquivos versionados, saída de comandos e report de bootstrap.

## 11. Fora do escopo técnico explícito

Não criar domínio de produto, hosted service, dashboard ou deploy tooling obrigatório.

## 12. Perguntas secundárias

O modo interativo inicia? O modo fixo opera? O CI detecta drift?
