# {{PROJECT_SLUG}} Autopilot

## 1. Missao

O autopilot do `{{PROJECT_SLUG}}` executa o roadmap fase por fase como uma maquina de estados controlada por evidencia.

A missao e:

1. executar o roadmap do `{{PROJECT_SLUG}}` fase por fase;
2. operar como maquina de estados baseada em `automation/PHASE_STATE.json`;
3. usar o report da fase como fonte de verdade para decisao de avanco;
4. nunca avancar por narrativa, intencao, resumo otimista ou ausencia de erro aparente;
5. preservar o fluxo Spec Driven Development definido em `AGENTS.md`;
6. respeitar a regra dual-spec: toda fase precisa de spec de construcao e spec de diagnostico;
7. parar imediatamente quando qualquer stop condition for encontrada.

O autopilot nao decide que uma fase esta pronta. O autopilot le o report da fase e so avanca quando o report contem exatamente o gate esperado para a fase atual.

## 1.1 Modos De Autopilot

O repositorio possui dois modos de automacao governada:

1. `phase_autopilot`: modo deste documento, fechado sobre `automation/ROADMAP.json` e `automation/PHASE_STATE.json`;
2. `interactive_sdd_autopilot`: modo interativo definido em `automation/INTERACTIVE_AUTOPILOT.md`, iniciado por solicitacao de feature ou evolucao e registrado em `automation/INTERACTIVE_STATE.json`.

O modo `interactive_sdd_autopilot` e aditivo. Ele nao pode reordenar, simplificar, remover ou reinterpretar o roadmap fixo deste documento.

Trilhas interativas nao devem ser registradas em `automation/ROADMAP.json`.

Solicitacoes interativas nao devem alterar `automation/PHASE_STATE.json`.

## 2. Fonte De Verdade

A execucao deve seguir esta hierarquia:

1. `AGENTS.md`;
2. spec de construcao da fase atual;
3. spec de diagnostico da fase atual;
4. `automation/ROADMAP.json`;
5. `automation/PHASE_STATE.json`;
6. `automation/STOP_CONDITIONS.md`;
7. report gerado para a fase atual.

Quando houver conflito entre implementacao e spec aprovada, a spec prevalece.

Quando houver conflito entre narrativa da conversa e report versionado, o report versionado prevalece para decisao de gate.

Quando o report estiver ausente, incompleto ou inconsistente com a fase atual, a fase fica bloqueada.

## 3. Roadmap Governado

O autopilot deve executar exatamente estas fases, nesta ordem:

1. Fase 20: `memory-adapter-expansion-and-storage-parity`
2. Fase 21: `agent-invocation-and-conversation-contract-parity`
3. Fase 22: `workflow-state-persistence-and-resume-parity`
4. Fase 23: `resumable-streams-and-playground-runtime-parity`
5. Fase 24: `hosting-integration-and-{{UPSTREAM_NAME_LOWER}}-app-migration-readiness`

O autopilot nao pode inventar fases, remover fases, reordenar fases, agrupar fases ou antecipar trabalho de fase futura.

## 4. Maquina De Estados

O estado persistente vive em `automation/PHASE_STATE.json`.

Campos obrigatorios:

1. `current_phase`: fase que deve ser executada agora;
2. `status`: estado operacional da fase atual;
3. `last_completed_phase`: ultima fase concluida com gate aprovado;
4. `retry_count`: numero de retries corretivos ja consumidos na fase atual;
5. `last_report`: caminho do ultimo report lido;
6. `last_decision`: decisao final extraida do ultimo report;
7. `blocked`: indica se o autopilot esta bloqueado;
8. `block_reason`: motivo objetivo do bloqueio.

Estados permitidos para `status`:

1. `pending`;
2. `running`;
3. `diagnosing`;
4. `blocked`;
5. `completed`;
6. `finished`.

Transicoes permitidas:

1. `pending` -> `running`
2. `running` -> `diagnosing`
3. `diagnosing` -> `completed`
4. `diagnosing` -> `blocked`
5. `running` -> `blocked`
6. `completed` -> `pending` da proxima fase
7. `completed` da Fase 24 -> `finished`

Qualquer transicao fora dessas deve ser tratada como inconsistencia de estado.

## 5. Ciclo Obrigatorio Por Fase

Para cada fase, o autopilot deve executar o ciclo completo abaixo, sem pular etapas.

### 5.1 Criar Ou Validar A Spec Da Fase

O autopilot deve localizar a spec de construcao definida em `automation/ROADMAP.json`.

Se a spec existir:

1. ler a spec completa;
2. identificar objetivo, escopo, fora de escopo, contratos, testes e criterios de aceite;
3. verificar se a spec cobre materialmente a implementacao esperada para a fase;
4. registrar no resultado da iteracao quais secoes governaram a fase.

Se a spec nao existir:

1. criar a spec no caminho exato definido em `automation/ROADMAP.json`;
2. respeitar `AGENTS.md`;
3. respeitar a fase e o nome definidos no roadmap;
4. nao criar fase nova;
5. nao alterar a ordem do roadmap;
6. nao implementar antes de a spec existir.

### 5.2 Implementar A Spec Da Fase

A implementacao deve ficar restrita ao escopo da fase atual.

O autopilot deve:

1. ler o codigo existente relevante antes de editar;
2. preservar contratos publicos salvo quando a spec exigir e justificar explicitamente;
3. atualizar testes no mesmo patch quando houver mudanca de comportamento;
4. atualizar documentacao operacional ou publica quando houver impacto observavel;
5. manter `pkg/*` livre de dependencia em `internal/*`, exceto pela ponte publica permitida em `pkg/app`;
6. manter rastreabilidade entre spec, codigo, testes e docs;
7. nao abrir capability nova fora da fase atual;
8. nao mascarar gap com documentacao narrativa.

### 5.3 Criar Ou Validar A Spec De Diagnostico

O autopilot deve localizar a spec de diagnostico definida em `automation/ROADMAP.json`.

Se a spec existir:

1. ler a spec completa;
2. identificar sinais observaveis, modos de falha, comandos, evidencias esperadas e criterio de confirmacao;
3. verificar que a auditoria consegue classificar a fase como `PASS`, `PARTIAL`, `FAIL` ou `BLOCKED`;
4. verificar que a auditoria exige decisao final explicita.

Se a spec nao existir:

1. criar a spec no caminho exato definido em `automation/ROADMAP.json`;
2. vincular a spec de diagnostico a fase atual;
3. incluir sinais observaveis, sintomas de falha, hipoteses, metricas, logs, traces, health checks quando aplicavel, comandos de verificacao, troubleshooting e criterio de confirmacao;
4. nao executar diagnostico antes de a spec existir.

### 5.4 Executar O Diagnostico

O diagnostico deve seguir a spec de diagnostico da fase atual.

O autopilot deve executar, no minimo:

1. verificacao de aderencia entre spec de construcao e implementacao;
2. verificacao de aderencia entre spec de diagnostico e evidencias coletadas;
3. `go test ./...`;
4. `go test -race ./...`;
5. validacao de OpenAPI quando a fase tocar API, runtime HTTP, playground, app hosting ou contrato documentado por OpenAPI;
6. testes especificos exigidos pela spec da fase;
7. leitura de docs, exemplos, feature matrix e reports relacionados quando a fase exigir;
8. classificacao objetiva da fase;
9. geracao do report no caminho exato definido em `automation/ROADMAP.json`.

Falha em qualquer validacao obrigatoria e bloqueio, exceto quando a spec de diagnostico autorizar explicitamente classificacao diferente e o report documentar a justificativa com evidencia.

### 5.5 Localizar O Report Gerado

O autopilot deve localizar o report no caminho exato da fase atual:

1. Fase 20: `docs/reports/phase-20-memory-adapter-expansion-and-storage-parity-report.md`
2. Fase 21: `docs/reports/phase-21-agent-invocation-and-conversation-contract-parity-report.md`
3. Fase 22: `docs/reports/phase-22-workflow-state-persistence-and-resume-parity-report.md`
4. Fase 23: `docs/reports/phase-23-resumable-streams-and-playground-runtime-parity-report.md`
5. Fase 24: `docs/reports/phase-24-hosting-integration-and-{{UPSTREAM_NAME_LOWER}}-app-migration-readiness-report.md`

Report ausente e stop condition.

Report em caminho diferente nao satisfaz o gate.

Report inconsistente com a fase atual e stop condition.

### 5.6 Ler A Decisao Final

O autopilot deve ler o report completo e extrair a classificacao e a decisao final.

A classificacao deve ser uma destas:

1. `PASS`
2. `PARTIAL`
3. `FAIL`
4. `BLOCKED`

A decisao deve ser exatamente a decisao esperada da fase atual.

Gates esperados:

1. Fase 20:
   - `classification = PASS`
   - `decision = READY FOR PHASE 21`

2. Fase 21:
   - `classification = PASS`
   - `decision = READY FOR PHASE 22`

3. Fase 22:
   - `classification = PASS`
   - `decision = READY FOR PHASE 23`

4. Fase 23:
   - `classification = PASS`
   - `decision = READY FOR PHASE 24`

5. Fase 24:
   - `classification = PASS`
   - `decision = READY FOR MIGRATION EXECUTION`

Decisao ausente e stop condition.

Classificacao ausente e stop condition.

Decisao ambigua e stop condition.

### 5.7 Atualizar `automation/PHASE_STATE.json`

O autopilot deve atualizar `automation/PHASE_STATE.json` ao final de cada tentativa.

Quando a fase inicia:

1. `status = "running"`
2. `blocked = false`
3. `block_reason = ""`

Quando a fase entra em diagnostico:

1. `status = "diagnosing"`

Quando a fase passa no gate:

1. `last_completed_phase` recebe a fase concluida;
2. `last_report` recebe o caminho do report lido;
3. `last_decision` recebe a decisao final do report;
4. `retry_count = 0`;
5. `blocked = false`;
6. `block_reason = ""`;
7. se houver proxima fase, `current_phase` recebe a proxima fase e `status = "pending"`;
8. se a fase concluida for a Fase 24, `current_phase = 24` e `status = "finished"`.

Quando a fase bloqueia:

1. `current_phase` permanece na fase atual;
2. `status = "blocked"`;
3. `last_report` recebe o caminho do report lido, se existir;
4. `last_decision` recebe a decisao extraida, se existir;
5. `blocked = true`;
6. `block_reason` recebe motivo objetivo;
7. `retry_count` reflete os retries corretivos consumidos.

### 5.8 Avancar Apenas Se O Gate For Satisfeito

O autopilot so pode avancar quando todas as condicoes forem verdadeiras:

1. report existe no caminho esperado;
2. report corresponde a fase atual;
3. `classification == PASS`;
4. `decision` e exatamente a decisao esperada para a fase atual;
5. `go test ./...` passou;
6. `go test -race ./...` passou;
7. validacao OpenAPI passou quando aplicavel;
8. nao ha divergencia material entre spec e implementacao;
9. nao ha stop condition ativa;
10. `automation/PHASE_STATE.json` foi atualizado de forma coerente.

## 6. Regras De Nao Avanco

O autopilot nunca deve avancar se o report trouxer qualquer uma destas condicoes:

1. `PARTIAL`;
2. `FAIL`;
3. `BLOCKED`;
4. `NOT READY`;
5. decisao ausente;
6. classificacao ausente;
7. report ausente;
8. report inconsistente com a fase atual;
9. report com decisao diferente da esperada;
10. report que declara evidencia insuficiente;
11. report que indica testes obrigatorios nao executados;
12. report que indica divergencia material entre spec e implementacao;
13. report que depende de decisao humana pendente.

O autopilot tambem nao deve avancar por inferencia baseada em testes verdes se o report nao satisfizer o gate.

## 7. Retry Policy

Cada fase permite no maximo `2` retries corretivos.

Retry corretivo significa uma tentativa de corrigir falhas identificadas durante implementacao, testes, diagnostico ou leitura do report da fase atual.

Politica:

1. ao encontrar falha corrigivel na fase atual, incrementar `retry_count`;
2. executar a correcao dentro do escopo da fase atual;
3. reexecutar as validacoes obrigatorias;
4. regenerar ou atualizar o report da fase;
5. reler o gate;
6. avancar somente se o gate for satisfeito.

Se `retry_count` exceder `2`, o autopilot deve parar.

Ao parar por retries excedidos, deve produzir um relatorio curto de bloqueio contendo:

1. fase afetada;
2. spec de construcao;
3. spec de diagnostico;
4. report esperado;
5. retries consumidos;
6. falhas persistentes;
7. acao humana recomendada.

## 8. Disciplina De Escopo

O autopilot deve manter disciplina estrita de escopo.

Regras obrigatorias:

1. nao abrir capability nova fora da fase atual;
2. nao implementar itens de fase futura;
3. nao reordenar o roadmap;
4. nao remover criterios de aceite;
5. nao mascarar gap com documentacao narrativa;
6. preferir reconciliacao documental e testes quando o codigo ja estiver correto;
7. nao redefinir API publica silenciosamente;
8. nao alterar naming publico sem justificativa explicita;
9. nao renomear contrato publico para satisfazer preferencia estetica;
10. nao introduzir dependencia externa desnecessaria;
11. nao introduzir hosted service como requisito para runtime local;
12. nao declarar capacidade operacional sem artefato real correspondente;
13. nao tratar intencao, plano, comentario ou TODO como entrega concluida.

## 9. Validacoes Obrigatorias

As validacoes abaixo sao obrigatorias em todas as fases, salvo quando a spec de diagnostico registrar explicitamente que uma validacao nao se aplica a fase atual e justificar com evidencia.

Bloqueios obrigatorios:

1. `go test ./...` falhar;
2. `go test -race ./...` falhar;
3. OpenAPI quebrar quando a fase tocar API, runtime HTTP, playground, app hosting ou contrato documentado por OpenAPI;
4. diagnostico nao gerar report;
5. report estar ausente;
6. report nao conter decisao final explicita;
7. existir divergencia material entre spec e implementacao;
8. existir breaking change sem aprovacao explicita em spec;
9. existir rename ou naming inconsistente afetando contrato publico;
10. existir necessidade de decisao humana de arquitetura;
11. `automation/PHASE_STATE.json` divergir do roadmap;
12. specs de construcao e diagnostico da fase nao existirem apos a etapa de criacao ou validacao.

## 10. Evidence-First

Toda conclusao do autopilot deve apontar para evidencia versionada ou comando executado.

Evidencias aceitas:

1. spec versionada;
2. codigo versionado;
3. teste versionado;
4. saida de comando executado;
5. documentacao versionada;
6. feature matrix;
7. report da fase;
8. diff revisavel;
9. diagnostico versionado.

Evidencias nao aceitas como conclusao de fase:

1. intencao;
2. resumo narrativo;
3. plano futuro;
4. comentario sem teste;
5. TODO;
6. ausencia de erro visual;
7. suposicao de compatibilidade;
8. report de fase diferente;
9. report sem classificacao;
10. report sem decisao final.

## 11. Encerramento Da Automacao

O autopilot termina com sucesso somente quando a Fase 24 tiver:

1. report no caminho `docs/reports/phase-24-hosting-integration-and-{{UPSTREAM_NAME_LOWER}}-app-migration-readiness-report.md`;
2. `classification = PASS`;
3. `decision = READY FOR MIGRATION EXECUTION`;
4. `go test ./...` verde;
5. `go test -race ./...` verde;
6. nenhuma stop condition ativa;
7. `automation/PHASE_STATE.json` com `status = "finished"`;
8. `last_completed_phase = 24`.

Qualquer outra parada deve ser tratada como bloqueio operacional.
