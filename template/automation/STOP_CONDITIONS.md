# {{PROJECT_SLUG}} Autopilot Stop Conditions

## 1. Regra Geral

O autopilot deve parar imediatamente quando qualquer condicao deste documento for encontrada.

Ao parar, o autopilot deve:

1. manter `current_phase` na fase afetada;
2. definir `status = "blocked"` em `automation/PHASE_STATE.json`;
3. definir `blocked = true`;
4. registrar `block_reason` com motivo objetivo;
5. preservar `retry_count`;
6. preservar `last_report` e `last_decision` quando existirem;
7. produzir resumo curto do bloqueio;
8. nao avancar para a proxima fase.

Para `interactive_sdd_autopilot`, a mesma regra geral se aplica a `automation/INTERACTIVE_STATE.json`: manter a etapa afetada, definir `status = "blocked"`, preservar report/classificacao/decisao quando existirem, registrar `block_reason` e nao avancar para a proxima etapa.

## 2. Stop Conditions De Report

Parar imediatamente se:

1. o report da fase atual estiver ausente;
2. o report estiver em caminho diferente do definido em `automation/ROADMAP.json`;
3. o report for inconsistente com a fase atual;
4. o report nao tiver classificacao explicita;
5. o report nao tiver decisao final explicita;
6. o report trouxer `classification = FAIL`;
7. o report trouxer `classification = BLOCKED`;
8. o report trouxer `classification = PARTIAL`;
9. o report trouxer decisao contendo `NOT READY`;
10. o report trouxer decisao diferente da esperada para a fase atual;
11. o report depender de evidencia externa nao versionada;
12. o report declarar testes obrigatorios nao executados;
13. o report declarar diagnostico incompleto;
14. o report declarar divergencia material entre spec e codigo;
15. o report declarar decisao humana pendente.

## 3. Stop Conditions De Retry

Parar imediatamente se:

1. `retry_count` exceder `2`;
2. a mesma falha persistir apos `2` retries corretivos;
3. o retry exigiria alterar fase futura;
4. o retry exigiria remover criterio de aceite;
5. o retry exigiria reordenar roadmap;
6. o retry exigiria criar capability fora da fase atual;
7. o retry exigiria breaking change sem aprovacao explicita.

## 4. Stop Conditions De Testes

Parar imediatamente se:

1. `go test ./...` falhar;
2. `go test -race ./...` falhar;
3. teste de conformidade exigido pela fase falhar;
4. teste de contrato publico exigido pela fase falhar;
5. teste de integracao exigido pela fase falhar;
6. teste que comprova criterio de aceite estiver ausente;
7. teste estiver pulado sem justificativa aceita pela spec de diagnostico;
8. suite falhar por race condition;
9. suite falhar por dependencia real externa nao prevista;
10. suite exigir segredo, credencial ou servico hospedado fora do escopo.

## 5. Stop Conditions De OpenAPI

Parar imediatamente se:

1. OpenAPI estiver invalida;
2. OpenAPI estiver divergente da implementacao;
3. OpenAPI estiver divergente da spec da fase;
4. contrato HTTP publico mudar sem atualizacao de spec;
5. endpoint publico for removido sem aprovacao explicita;
6. schema publico for renomeado sem aprovacao explicita;
7. campo publico obrigatorio for removido sem aprovacao explicita;
8. comportamento documentado por OpenAPI nao tiver teste correspondente quando a fase tocar API.

## 6. Stop Conditions De Spec

Parar imediatamente se:

1. a spec de construcao da fase atual estiver ausente apos a etapa de criacao ou validacao;
2. a spec de diagnostico da fase atual estiver ausente apos a etapa de criacao ou validacao;
3. a spec de construcao nao definir objetivo, escopo, fora de escopo, contratos, testes e criterios de aceite;
4. a spec de diagnostico nao definir sinais observaveis, modos de falha, comandos de verificacao, classificacao e decisao final;
5. houver divergencia material entre spec e codigo;
6. houver comportamento implementado fora da spec;
7. houver capability declarada apenas em documentacao narrativa;
8. houver conflito entre specs sem decisao explicita;
9. houver necessidade de decisao humana de arquitetura;
10. a fase depender de spec futura ainda nao executada.

## 7. Stop Conditions De Contrato Publico

Parar imediatamente se:

1. houver breaking change sem aprovacao explicita;
2. API publica for redefinida silenciosamente;
3. naming publico mudar sem justificativa explicita;
4. rename afetar contrato publico sem spec correspondente;
5. tipo de `internal/*` vazar para assinatura exportada;
6. pacote `pkg/*` depender de `internal/*`, exceto pela ponte publica permitida em `pkg/app`;
7. erro publico observavel mudar sem documentacao e teste;
8. comportamento de cancelamento mudar sem spec e teste;
9. ordem de execucao observavel mudar sem spec e teste;
10. compatibilidade conceitual com {{UPSTREAM_NAME}} divergir sem justificativa registrada.

## 8. Stop Conditions De Roadmap E Estado

Parar imediatamente se:

1. `automation/ROADMAP.json` estiver ausente;
2. `automation/PHASE_STATE.json` estiver ausente;
3. `automation/PHASE_STATE.json` nao for JSON valido;
4. `automation/ROADMAP.json` nao for JSON valido;
5. `current_phase` nao existir em `automation/ROADMAP.json`;
6. `last_completed_phase` for maior ou igual a `current_phase`;
7. `blocked = true` e o autopilot tentar avancar sem desbloqueio;
8. roadmap estiver inconsistente com `PHASE_STATE.json`;
9. fase for pulada;
10. fase for executada fora de ordem;
11. fase nova for adicionada sem pedido explicito;
12. fase existente for removida;
13. gate esperado for alterado;
14. report esperado for alterado.

## 9. Stop Conditions De Escopo

Parar imediatamente se:

1. a correcao exigir abrir capability fora da fase atual;
2. a correcao exigir implementar fase futura;
3. a correcao exigir reordenar fases;
4. a correcao exigir simplificar o roadmap;
5. a correcao exigir remover diagnostico;
6. a correcao exigir mascarar gap com texto narrativo;
7. a correcao exigir aceitar evidencia insuficiente;
8. a correcao exigir depender de hosted service fora do escopo;
9. a correcao exigir {{UPSTREAM_OPS_NAME}};
10. a correcao exigir decisao humana de arquitetura.

## 10. Stop Conditions Do Modo Interativo

Parar imediatamente no `interactive_sdd_autopilot` se:

1. a solicitacao inicial estiver materialmente ambigua e ainda depender de resposta do usuario;
2. a trilha exigir comportamento novo sem spec de construcao;
3. a trilha exigir comportamento novo sem spec de diagnostico;
4. a decisao entre amendar spec existente e criar nova trilha dual-spec estiver ausente;
5. `automation/INTERACTIVE_STATE.json` estiver ausente ou invalido;
6. `current_request.current_step` estiver inconsistente com a execucao;
7. o report da etapa atual estiver ausente;
8. o report nao trouxer classificacao explicita;
9. o report nao trouxer decisao final explicita;
10. o report trouxer `classification = PARTIAL`, `classification = FAIL` ou `classification = BLOCKED`;
11. a decisao final nao autorizar avanco, retry ou conclusao;
12. a implementacao exigiria alterar `automation/ROADMAP.json` ou `automation/PHASE_STATE.json` para contornar gate;
13. a implementacao exigiria capability publica do framework fora de spec;
14. a implementacao exigiria {{UPSTREAM_OPS_NAME}}, hosted service, dashboard, control plane ou deploy tooling;
15. a trilha tentar concluir sem evidencia versionada;
16. houver decisao humana de arquitetura pendente.

Ao parar no modo interativo, registrar o bloqueio em `automation/INTERACTIVE_STATE.json`.

## 11. Relatorio Obrigatorio Ao Parar

Quando parar por qualquer stop condition, o autopilot deve produzir:

1. resumo do bloqueio;
2. fase afetada;
3. nome da fase;
4. spec de construcao;
5. spec de diagnostico;
6. report esperado;
7. report encontrado, se houver;
8. classificacao encontrada, se houver;
9. decisao encontrada, se houver;
10. retries ja consumidos;
11. stop condition acionada;
12. evidencia observada;
13. acao humana recomendada.
