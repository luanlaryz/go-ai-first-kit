# {{PROJECT_SLUG}} Autopilot Runbook

## 1. Objetivo

Este runbook instrui o Cursor a executar o roadmap automatizado do `{{PROJECT_SLUG}}` como maquina de estados, uma fase por vez, usando reports como fonte de verdade e gates explicitos para avanco.

O Cursor deve executar somente a fase atual indicada por `automation/PHASE_STATE.json`.

Este runbook governa apenas o modo `phase_autopilot`.

Para solicitacoes interativas de feature, evolucao, bug, refatoracao ou documentacao, use `automation/INTERACTIVE_RUNBOOK.md`, `automation/INTERACTIVE_AUTOPILOT.md` e `automation/INTERACTIVE_STATE.json`.

O modo interativo nao deve alterar `automation/ROADMAP.json` nem `automation/PHASE_STATE.json` para representar uma trilha de feature request.

## 2. Leitura Obrigatoria Inicial

Antes de qualquer implementacao, diagnostico ou atualizacao de estado, o Cursor deve ler:

1. `AGENTS.md`;
2. `skills/00-skill-index/SKILL.md`;
3. `automation/AUTOPILOT.md`;
4. `automation/ROADMAP.json`;
5. `automation/PHASE_STATE.json`;
6. `automation/STOP_CONDITIONS.md`.

Depois deve identificar:

1. `current_phase`;
2. fase correspondente em `automation/ROADMAP.json`;
3. spec de construcao da fase atual;
4. spec de diagnostico da fase atual;
5. report esperado da fase atual;
6. gate esperado da fase atual;
7. `retry_count`;
8. estado de `blocked`.

Se `blocked = true`, seguir a rotina de retomada apos bloqueio antes de executar qualquer nova fase.

## 3. Execucao Somente Da Fase Atual

O Cursor deve executar somente a fase indicada por `current_phase`.

O Cursor nao deve:

1. executar fase futura;
2. reexecutar fase ja concluida;
3. pular fase;
4. reordenar roadmap;
5. alterar gates;
6. trocar nomes de arquivos;
7. substituir report esperado;
8. alterar `last_completed_phase` sem report aprovado.

## 4. Procedimento Da Fase Atual

Para a fase atual, executar obrigatoriamente:

1. criar ou validar spec;
2. implementar a spec;
3. criar ou validar diagnosis spec;
4. executar diagnostico;
5. localizar e ler o report;
6. atualizar `automation/PHASE_STATE.json`;
7. avancar somente quando o gate for satisfeito;
8. parar imediatamente quando houver stop condition.

## 5. Criar Ou Validar Spec

Localizar `spec_file` da fase atual em `automation/ROADMAP.json`.

Se existir:

1. ler a spec completa;
2. registrar secoes governantes;
3. confirmar objetivo, escopo, fora de escopo, contratos, testes e criterios de aceite;
4. identificar arquivos provaveis de implementacao;
5. identificar docs e testes exigidos.

Se nao existir:

1. criar a spec no caminho exato definido por `spec_file`;
2. manter o nome da fase exatamente como definido em `automation/ROADMAP.json`;
3. escrever objetivo, motivacao, pergunta principal, escopo, fora de escopo, contratos, arquitetura, testes, criterios de aceite e limites;
4. nao implementar antes de a spec existir.

## 6. Implementar

Durante a implementacao:

1. ler codigo existente antes de editar;
2. alterar somente arquivos necessarios para a fase atual;
3. preservar arquitetura `pkg/` e `internal/`;
4. preservar contratos publicos salvo autorizacao explicita da spec;
5. adicionar ou atualizar testes quando houver mudanca de comportamento;
6. atualizar docs quando houver comportamento observavel, API publica ou fluxo operacional afetado;
7. manter feature matrix coerente quando maturidade, cobertura, prioridade ou paridade mudarem;
8. registrar gaps se a spec exigir algo que nao pode ser implementado com seguranca.

## 7. Criar Ou Validar Diagnosis Spec

Localizar `diagnosis_spec_file` da fase atual em `automation/ROADMAP.json`.

Se existir:

1. ler a spec completa;
2. identificar checklist de diagnostico;
3. identificar comandos obrigatorios;
4. identificar sinais observaveis;
5. identificar modos de falha;
6. identificar criterio de classificacao;
7. identificar formato da decisao final.

Se nao existir:

1. criar a spec no caminho exato definido por `diagnosis_spec_file`;
2. vincular a diagnostico a fase atual;
3. definir sinais observaveis, modos de falha, comandos, evidencias, troubleshooting, classificacao e decisao final;
4. nao executar diagnostico antes de a spec existir.

## 8. Executar Diagnostico

Executar o diagnostico conforme a spec de diagnostico da fase atual.

Validacoes obrigatorias:

1. aderencia entre spec de construcao e implementacao;
2. aderencia entre spec de diagnostico e evidencias;
3. `go test ./...`;
4. `go test -race ./...`;
5. validacao OpenAPI quando aplicavel;
6. testes especificos da fase;
7. docs e feature matrix quando aplicavel;
8. ausencia de breaking change nao aprovado;
9. ausencia de divergencia material entre spec e codigo.

O diagnostico deve gerar o report no caminho exato definido por `report_file`.

## 9. Localizar E Ler Report

Depois do diagnostico, localizar `report_file`.

O Cursor deve:

1. confirmar que o arquivo existe;
2. confirmar que o report corresponde a fase atual;
3. ler o report completo;
4. extrair `classification`;
5. extrair `decision`;
6. comparar com `advance_only_if.classification`;
7. comparar com `advance_only_if.decision`.

Nao avancar se qualquer campo estiver ausente, ambiguo ou divergente.

## 10. Atualizar `PHASE_STATE.json`

### 10.1 Ao Iniciar Fase

Definir:

1. `status = "running"`;
2. `blocked = false`;
3. `block_reason = ""`.

### 10.2 Ao Entrar Em Diagnostico

Definir:

1. `status = "diagnosing"`.

### 10.3 Ao Passar No Gate

Se `classification` e `decision` forem exatamente os esperados:

1. definir `last_completed_phase` como a fase atual;
2. definir `last_report` como o report lido;
3. definir `last_decision` como a decisao lida;
4. definir `retry_count = 0`;
5. definir `blocked = false`;
6. definir `block_reason = ""`;
7. se houver proxima fase, definir `current_phase` como a proxima fase;
8. se houver proxima fase, definir `status = "pending"`;
9. se a fase atual for 24, definir `status = "finished"` e manter `current_phase = 24`.

### 10.4 Ao Bloquear

Se houver stop condition:

1. manter `current_phase` na fase atual;
2. definir `status = "blocked"`;
3. definir `blocked = true`;
4. definir `block_reason` com motivo objetivo;
5. preservar ou atualizar `retry_count`;
6. definir `last_report` se houver report lido;
7. definir `last_decision` se houver decisao lida;
8. parar.

## 11. Politica De Avanco

Avancar somente quando todas as condicoes forem verdadeiras:

1. fase atual foi implementada ou validada;
2. spec de construcao existe;
3. spec de diagnostico existe;
4. diagnostico foi executado;
5. report foi gerado no caminho correto;
6. report foi lido;
7. `classification` e exatamente a esperada;
8. `decision` e exatamente a esperada;
9. `go test ./...` passou;
10. `go test -race ./...` passou;
11. OpenAPI passou quando aplicavel;
12. nenhuma stop condition esta ativa;
13. `automation/PHASE_STATE.json` foi atualizado.

## 12. Parada Por Stop Condition

Parar imediatamente quando qualquer item de `automation/STOP_CONDITIONS.md` ocorrer.

Ao parar por bloqueio, produzir:

1. resumo do bloqueio;
2. fase afetada;
3. retries ja consumidos;
4. acao humana recomendada.

O Cursor nao deve continuar o roadmap enquanto `blocked = true`.

## 13. Gates Por Fase

### Fase 20

Fase: `memory-adapter-expansion-and-storage-parity`

Spec:

`specs/580-phase-20-memory-adapter-expansion-and-storage-parity.md`

Diagnosis spec:

`specs/581-phase-20-memory-adapter-expansion-and-storage-parity-diagnosis.md`

Report:

`docs/reports/phase-20-memory-adapter-expansion-and-storage-parity-report.md`

Gate:

1. `classification = PASS`
2. `decision = READY FOR PHASE 21`

### Fase 21

Fase: `agent-invocation-and-conversation-contract-parity`

Spec:

`specs/600-phase-21-agent-invocation-and-conversation-contract-parity.md`

Diagnosis spec:

`specs/601-phase-21-agent-invocation-and-conversation-contract-parity-diagnosis.md`

Report:

`docs/reports/phase-21-agent-invocation-and-conversation-contract-parity-report.md`

Gate:

1. `classification = PASS`
2. `decision = READY FOR PHASE 22`

### Fase 22

Fase: `workflow-state-persistence-and-resume-parity`

Spec:

`specs/620-phase-22-workflow-state-persistence-and-resume-parity.md`

Diagnosis spec:

`specs/621-phase-22-workflow-state-persistence-and-resume-parity-diagnosis.md`

Report:

`docs/reports/phase-22-workflow-state-persistence-and-resume-parity-report.md`

Gate:

1. `classification = PASS`
2. `decision = READY FOR PHASE 23`

### Fase 23

Fase: `resumable-streams-and-playground-runtime-parity`

Spec:

`specs/640-phase-23-resumable-streams-and-playground-runtime-parity.md`

Diagnosis spec:

`specs/641-phase-23-resumable-streams-and-playground-runtime-parity-diagnosis.md`

Report:

`docs/reports/phase-23-resumable-streams-and-playground-runtime-parity-report.md`

Gate:

1. `classification = PASS`
2. `decision = READY FOR PHASE 24`

### Fase 24

Fase: `hosting-integration-and-{{UPSTREAM_NAME_LOWER}}-app-migration-readiness`

Spec:

`specs/660-phase-24-hosting-integration-and-{{UPSTREAM_NAME_LOWER}}-app-migration-readiness.md`

Diagnosis spec:

`specs/661-phase-24-hosting-integration-and-{{UPSTREAM_NAME_LOWER}}-app-migration-readiness-diagnosis.md`

Report:

`docs/reports/phase-24-hosting-integration-and-{{UPSTREAM_NAME_LOWER}}-app-migration-readiness-report.md`

Gate:

1. `classification = PASS`
2. `decision = READY FOR MIGRATION EXECUTION`

## 14. Resposta Final Obrigatoria Do Cursor

Ao encerrar cada iteracao, responder com:

1. objetivo;
2. specs lidas;
3. skills aplicadas;
4. arquivos alterados;
5. comandos executados;
6. testes executados;
7. report lido;
8. classificacao;
9. decisao;
10. estado atualizado;
11. gaps restantes;
12. se continuou, proxima fase;
13. se bloqueou, motivo e acao humana recomendada.

## 15. Prompt Mestre Para Iniciar O Autopilot

Use este prompt para iniciar a execucao automatizada do roadmap:

```text
Execute o autopilot do roadmap do projeto {{PROJECT_SLUG}}.

Regras obrigatorias:

1. Leia antes de qualquer acao:
   - AGENTS.md
   - skills/00-skill-index/SKILL.md
   - automation/AUTOPILOT.md
   - automation/ROADMAP.json
   - automation/PHASE_STATE.json
   - automation/STOP_CONDITIONS.md

2. Comece exatamente pela fase indicada em automation/PHASE_STATE.json.

3. Use automation/ROADMAP.json como lista fechada de fases. Nao invente fases, nao reordene fases, nao pule fases e nao simplifique o roadmap.

4. Execute somente a fase atual por vez.

5. Para cada fase, siga este ciclo completo:
   - criar ou validar a spec da fase;
   - implementar a spec da fase;
   - criar ou validar a spec de diagnostico;
   - executar o diagnostico;
   - localizar o report no caminho exato definido no roadmap;
   - ler o report completo;
   - extrair classification e decision;
   - atualizar automation/PHASE_STATE.json;
   - avancar somente se o gate da fase for satisfeito exatamente.

6. O report da fase e a fonte de verdade para avanco. Nunca avance por narrativa, intencao, resumo otimista, testes verdes isolados ou ausencia aparente de erro.

7. Avance somente quando:
   - classification == PASS;
   - decision for exatamente a decisao esperada da fase atual;
   - go test ./... passar;
   - go test -race ./... passar;
   - OpenAPI estiver valida quando aplicavel;
   - nao houver divergencia material entre spec e implementacao;
   - nenhuma stop condition estiver ativa.

8. Pare imediatamente se houver qualquer stop condition em automation/STOP_CONDITIONS.md, incluindo:
   - FAIL;
   - BLOCKED;
   - PARTIAL;
   - NOT READY;
   - retries excedidos;
   - go test ./... falhou;
   - go test -race ./... falhou;
   - OpenAPI invalida ou divergente;
   - report ausente;
   - report sem decisao final explicita;
   - breaking change sem aprovacao explicita;
   - divergencia material entre spec e codigo;
   - necessidade de decisao humana de arquitetura;
   - rename/naming inconsistente afetando contrato publico;
   - roadmap inconsistente com automation/PHASE_STATE.json.

9. Use no maximo 2 retries corretivos por fase. Depois disso, pare e registre bloqueio em automation/PHASE_STATE.json.

10. Ao bloquear, produza:
   - resumo do bloqueio;
   - fase afetada;
   - retries ja consumidos;
   - stop condition acionada;
   - evidencia observada;
   - acao humana recomendada.

11. Ao concluir uma fase com gate aprovado:
   - atualize last_completed_phase;
   - atualize last_report;
   - atualize last_decision;
   - zere retry_count;
   - marque blocked como false;
   - avance current_phase para a proxima fase;
   - deixe status como pending para a proxima fase.

12. Ao concluir a Fase 24 com:
   - classification = PASS;
   - decision = READY FOR MIGRATION EXECUTION;
   atualize automation/PHASE_STATE.json com status = finished e last_completed_phase = 24.

13. Continue automaticamente fase por fase ate concluir todas as fases ou atingir uma stop condition.

Formato obrigatorio da resposta ao final de cada iteracao:
- objetivo
- specs lidas
- skills aplicadas
- arquivos alterados
- comandos executados
- testes executados
- report lido
- classification
- decision
- estado atualizado
- gaps restantes
- proxima fase ou motivo de bloqueio
```

## 16. Prompt De Continuacao Normal

Use este prompt quando o Cursor ja estiver no meio do processo e precisar continuar de onde parou:

```text
Continue de onde parou no autopilot do {{PROJECT_SLUG}}.

Antes de qualquer acao, leia:
- AGENTS.md
- skills/00-skill-index/SKILL.md
- automation/AUTOPILOT.md
- automation/ROADMAP.json
- automation/PHASE_STATE.json
- automation/STOP_CONDITIONS.md

Retome exatamente da fase indicada em automation/PHASE_STATE.json.

Nao reinicie fases ja concluidas.
Nao pule gates.
Nao avance por narrativa.
Nao reordene o roadmap.
Nao altere os nomes das fases, specs, diagnosis specs ou reports.
Nao execute fase futura antes da fase atual passar no gate.

Se automation/PHASE_STATE.json indicar blocked = true, pare a continuacao normal e use o fluxo de retomada apos bloqueio.

Para a fase atual:
1. leia a spec de construcao definida em automation/ROADMAP.json;
2. crie a spec se ela ainda nao existir;
3. implemente ou continue a implementacao somente dentro do escopo da fase atual;
4. leia a spec de diagnostico definida em automation/ROADMAP.json;
5. crie a spec de diagnostico se ela ainda nao existir;
6. execute o diagnostico;
7. rode go test ./...;
8. rode go test -race ./...;
9. valide OpenAPI quando aplicavel;
10. gere ou localize o report no caminho exato definido no roadmap;
11. leia o report completo;
12. extraia classification e decision;
13. atualize automation/PHASE_STATE.json;
14. avance somente se classification == PASS e decision for exatamente a decisao esperada da fase atual.

Pare imediatamente se qualquer stop condition em automation/STOP_CONDITIONS.md ocorrer.

Use no maximo 2 retries corretivos por fase. Se os retries forem excedidos, registre bloqueio em automation/PHASE_STATE.json e pare.

Formato obrigatorio da resposta:
- objetivo
- specs lidas
- skills aplicadas
- arquivos alterados
- comandos executados
- testes executados
- report lido
- classification
- decision
- estado atualizado
- gaps restantes
- proxima fase ou motivo de bloqueio
```

## 17. Prompt De Retomada Apos Bloqueio

Use este prompt quando houver `blocked = true`, retries excedidos ou parada por stop condition:

```text
Retome o autopilot do {{PROJECT_SLUG}} apos bloqueio de forma segura.

Antes de qualquer acao, leia:
- AGENTS.md
- skills/00-skill-index/SKILL.md
- automation/AUTOPILOT.md
- automation/ROADMAP.json
- automation/PHASE_STATE.json
- automation/STOP_CONDITIONS.md

Inspecione primeiro o bloqueio registrado em automation/PHASE_STATE.json.

Nao resete o roadmap.
Nao apague historico de bloqueio.
Nao zere retry_count sem resolver o bloqueio.
Nao avance para fase futura enquanto blocked = true.
Nao reinicie fases ja concluidas.
Nao pule gates.
Nao altere ROADMAP.json para contornar bloqueio.

Procedimento obrigatorio:

1. Identifique:
   - current_phase;
   - status;
   - retry_count;
   - last_completed_phase;
   - last_report;
   - last_decision;
   - blocked;
   - block_reason.

2. Localize a fase atual em automation/ROADMAP.json.

3. Leia:
   - spec de construcao da fase atual;
   - spec de diagnostico da fase atual;
   - report da fase atual, se existir;
   - stop condition relacionada ao bloqueio.

4. Atue primeiro no desbloqueio da fase atual.

5. Corrija somente o que estiver dentro do escopo da fase atual.

6. Se o bloqueio exigir decisao humana de arquitetura, breaking change sem aprovacao, rename publico sensivel, alteracao de gate, reordenacao de roadmap ou capability fora da fase atual, pare e informe a acao humana recomendada.

7. Se o bloqueio for corrigivel:
   - incremente retry_count se esta for uma tentativa corretiva;
   - aplique a correcao;
   - rode go test ./...;
   - rode go test -race ./...;
   - valide OpenAPI quando aplicavel;
   - execute novamente o diagnostico da fase;
   - gere ou atualize o report no caminho exato definido no roadmap;
   - leia o report completo;
   - extraia classification e decision.

8. So desbloqueie automation/PHASE_STATE.json se:
   - classification == PASS;
   - decision for exatamente a decisao esperada da fase atual;
   - go test ./... passou;
   - go test -race ./... passou;
   - OpenAPI passou quando aplicavel;
   - a stop condition original foi resolvida;
   - nenhuma nova stop condition esta ativa.

9. Se o gate for satisfeito:
   - defina blocked = false;
   - limpe block_reason;
   - zere retry_count;
   - atualize last_completed_phase;
   - atualize last_report;
   - atualize last_decision;
   - avance current_phase para a proxima fase, ou marque status = finished se a fase concluida for 24.

10. Se o gate nao for satisfeito:
   - mantenha blocked = true;
   - mantenha current_phase na fase atual;
   - atualize block_reason;
   - preserve o historico relevante em last_report e last_decision;
   - pare.

11. Se retry_count exceder 2:
   - mantenha blocked = true;
   - defina status = blocked;
   - registre block_reason como retries corretivos excedidos;
   - pare.

Formato obrigatorio da resposta:
- objetivo
- specs lidas
- skills aplicadas
- arquivos alterados
- comandos executados
- testes executados
- bloqueio original
- acao de desbloqueio executada
- report lido
- classification
- decision
- estado atualizado
- gaps restantes
- proxima fase ou motivo de bloqueio persistente
```
