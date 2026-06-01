# Backlog Update Procedure

Use este procedimento sempre que criar ou atualizar `docs/backlog/Backlog.md`.

## 1. Preparar contexto

Leia:

1. `AGENTS.md`
2. `skills/00-skill-index/SKILL.md`
3. `skills/26-backlog-item-intake/SKILL.md`
4. `docs/backlog/Backlog.md`
5. fontes citadas pela entrada
6. specs governantes ou candidatas

## 2. Preservar fontes históricas

Não edite nem migre automaticamente arquivos auxiliares em `docs/backlog/**` (planos importados, backlogs herdados, relatórios de pesquisa ou validação).

Esses arquivos podem ser citados em:

- `## Backlogs auxiliares e fontes históricas`
- `## Itens candidatos a triagem`

## 3. Deduplicar

Antes de criar item:

```bash
rg "BLG-[0-9]{4}" docs/backlog/Backlog.md
rg "<palavra-chave>|<spec>|<fonte>" docs/backlog/Backlog.md
```

Se existir item equivalente:

- não crie novo ID;
- atualize evidência, histórico ou status do item existente;
- registre a fonte adicional em `Histórico`.

## 4. Calcular próximo ID

1. Liste todos os IDs `BLG-NNNN`.
2. Use o maior número + 1.
3. Se não houver item real, o primeiro ID é `BLG-0001`.
4. Não reserve ID para candidato histórico não triado.

## 5. Inserir ou atualizar item

Use `references/backlog-item-template.md`.

Campos obrigatórios:

- status;
- tipo;
- prioridade;
- severidade;
- valor;
- complexidade;
- risco;
- decisão de escopo;
- spec governante;
- spec de diagnóstico;
- fonte;
- evidência;
- impacto em API pública;
- risco de breaking change;
- problema;
- resultado esperado;
- escopo;
- fora do escopo;
- plano;
- testes;
- stop conditions;
- prompts quando aplicável;
- histórico.

## 6. Atualizar índice de priorização

Em `## Índice de priorização`, liste apenas itens `BLG-*` reais.

Ordenação recomendada:

1. `P0`
2. `P1`
3. `P2`
4. `P3`
5. bloqueados por decisão humana
6. rejeitados
7. concluídos

Dentro da mesma prioridade, prefira maior risco e maior valor.

## 7. Atualizar candidatos a triagem

Use `## Itens candidatos a triagem` para fontes ainda não promovidas.

Cada candidato deve conter:

- fonte;
- recorte;
- decisão atual;
- motivo para não criar `BLG-*` ainda;
- próxima ação de triagem.

## 8. Validar

Execute checks documentais:

```bash
test -f docs/backlog/Backlog.md
rg "# {{PROJECT_TITLE}} Backlog|## Índice de priorização|## Itens" docs/backlog/Backlog.md
rg "## Backlogs auxiliares e fontes históricas|## Itens candidatos a triagem" docs/backlog/Backlog.md
rg "BLG-[0-9]{4}" docs/backlog/Backlog.md
```

Se o backlog ainda não tiver item real, o último comando pode não encontrar resultado. Registre isso como esperado, não como falha.

## 9. Responder

Informe:

- item criado ou atualizado;
- classificação;
- decisão final;
- arquivos alterados;
- comandos executados;
- gaps restantes;
- próxima etapa.
