# Documentação do go-ai-first-kit

Índice por objetivo. Cada documento declara apenas o que os artefatos versionados sustentam.

## Avaliar o kit

- [capabilities.md](capabilities.md): catálogo de capacidades, tipos de verificação, estados e limites — comece aqui para saber o que é automação, o que é processo governado e o que é apenas guidance.
- [adoption-guide.md](adoption-guide.md): quando adotar, quando evitar e a sequência verificável de adoção.

## Gerar e diagnosticar um projeto

- [cli-reference.md](cli-reference.md): comandos e flags de `gakit create`, `gakit template list`, `gakit diagnose` e `gakit version`.
- [../README.md](../README.md): fluxo rápido de criação e o fallback por prompt único.
- Após gerar, o projeto contém sua própria central em `docs/README.md`, com catálogo (`docs/ai/capabilities.md`) e jornadas (`docs/journeys/`).

## Entender a arquitetura documental

- [how-it-works.md](how-it-works.md): as três partes do kit (CLI, template, prompt master) e a separação entre os dois autopilots.
- [capabilities.md](capabilities.md): seção "Capacidades entregues nos projetos gerados" resume as camadas renderizadas.

## Manter a baseline do kit

- [govolt-sync-diagnosis.md](govolt-sync-diagnosis.md): diagnóstico do diff `govolt → kit`, matriz de paridade e decisões por item.
- [govolt-sync-baseline.json](govolt-sync-baseline.json): pin do SHA upstream aplicado e exclusões da rodada.
- [extracted-from-govolt.md](extracted-from-govolt.md): proveniência e divergências intencionais.
- [sac-agents-sync-diagnosis.md](sac-agents-sync-diagnosis.md): diagnóstico do diff `sac-agents → kit` (sincronização downstream), decisões por artefato e limitações aceitas.
- [sac-agents-sync-baseline.json](sac-agents-sync-baseline.json): pin do SHA downstream aplicado e exclusões da rodada.
- [../CHANGELOG.md](../CHANGELOG.md): histórico de mudanças do kit.

Regra editorial: toda afirmação nova nestes documentos deve apontar para artefato renderizável ou comando executável. Capacidades futuras entram como gap ou trabalho a especificar, nunca como entrega concluída.
