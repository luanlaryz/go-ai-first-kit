---
name: 30-third-party-service-integrations
description: Design, implement and review third-party service integrations using ports/adapters, encrypted per-tenant credentials, bounded credential cache, timeouts, retries, circuit breakers and sanitized observability. Use when adding or changing any adapter that calls an external system, or any agent tool that reaches outside the process.
---

# Third-party Service Integrations

Goal: manter integracoes externas seguras, escopadas por tenant e substituiveis, permitindo muitos
servicos sem que nenhum deles vire dependencia arquitetural do core.

## When to Use

- Adicionar ou alterar integracao com {{EXTERNAL_SYSTEM}} ou qualquer provider HTTP/gRPC externo.
- Construir tools de agente que leem ou escrevem em sistemas externos.
- Adicionar configuracao de credencial, lookup de token, cache de token, auth de adapter, retry,
  circuit breaker ou mapeamento de erro de provider.

## Required Reading

1. `AGENTS.md` e a spec governante da integracao.
2. `skills/01-hexagonal-architecture/SKILL.md` e `skills/17-solid-go-ports/SKILL.md`.
3. `skills/18-security-owasp-api/SKILL.md` e `skills/19-prompt-injection-llm-safety/SKILL.md` quando a
   chamada for disparada por agente ou tool.
4. `skills/08-redis-cache-streams/SKILL.md` e `skills/12-observability-zap-otel-prom/SKILL.md` quando o
   escopo alcancar cache ou telemetria.

## Non-negotiables

1. Todo servico tem um port de saida explicito e um adapter especifico do vendor. Nenhuma chamada
   direta de SDK ou HTTP a partir do dominio ou da aplicacao.
2. Contratos do core usam DTOs canonicos. Payload externo, tipos de SDK e erros de provider ficam
   dentro do adapter.
3. Credenciais sao escopadas por {{DOMAIN_ACTOR}} e por chave de servico, cifradas em repouso e nunca
   registradas em log, span, fixture ou artefato de review.
4. Cache de credencial tem TTL explicito e limitado. Cache sem TTL e sem invalidacao vira credencial
   revogada que continua funcionando.
5. Atualizacao de credencial invalida a chave de cache afetada **antes** de a mudanca ser considerada
   concluida.
6. Adapters usam deadline de contexto, timeout configurado, retry limitado com backoff e mapeamento
   sanitizado de erro. Retry infinito transforma indisponibilidade do provider em indisponibilidade
   propria.
7. Tool com efeito colateral passa por gate de politica antes de chamar o servico externo.
8. Observabilidade preserva correlacao e chave de servico, resultado e latencia, sem dado pessoal nem
   segredo.

## Do / Don't

- **Do** definir ports pequenos, de propriedade do consumidor, com 1 a 3 metodos.
- **Do** separar resolucao de credencial da operacao de negocio quando isso mantem as interfaces
  pequenas.
- **Do** validar chave de servico, tipo de auth, limites de TTL e formato do token antes de persistir.
- **Do** retornar erro de dominio ou aplicacao a partir do caso de uso, nunca erro cru do provider.
- **Don't** ramificar logica do core por nome de provider; a selecao acontece no wiring.
- **Don't** fazer fallback para credencial global compartilhada sem decisao registrada: fallback
  silencioso quebra o isolamento por tenant.
- **Don't** cachear token em variavel global de processo sem TTL nem invalidacao.
- **Don't** colocar loop de retry dentro da tool se isso altera a semantica observavel definida pelo
  contrato do adapter.

## Interfaces / Contracts

Port de resolucao de credencial:

```go
type CredentialResolver interface {
    Resolve(ctx context.Context, tenant TenantID, service ServiceKey) (Credential, error)
}
```

Port de servico (exemplo):

```go
type OrderReader interface {
    Get(ctx context.Context, tenant TenantID, ref OrderRef) (Order, error)
}
```

Formato de persistencia esperado (campos minimos):

```text
integration_credential(
  tenant_id,
  service_key,
  auth_type,
  token_prefix,
  token_hash,
  token_encrypted,
  ttl_seconds,
  is_active,
  last_used_at,
  audit fields
)
```

O `token_prefix` e o `token_hash` existem para permitir diagnostico e deteccao de rotacao sem
descriptografar nada.

Formato de chave de cache:

```text
tenant:{tenant_id}:integration:{service_key}:credential
```

## Checklists

**Before**
- [ ] Identificar o port canonico e os DTOs que o caso de uso precisa.
- [ ] Definir chave de servico, tipo de auth, TTL default/maximo e politica de fallback.
- [ ] Identificar gates de politica e risco de prompt injection se um agente ou tool disparar a chamada.
- [ ] Identificar endpoints de configuracao necessarios para gerenciar a credencial.

**During**
- [ ] Manter codigo de SDK e cliente HTTP apenas no adapter.
- [ ] Resolver credencial com descriptografia e cache com TTL explicito.
- [ ] Adicionar timeout, retry com backoff e circuit breaker quando aplicavel.
- [ ] Mapear falha de provider para erro de aplicacao sanitizado.
- [ ] Emitir metricas e traces sem segredo nem dado pessoal no payload.

**After**
- [ ] Teste unitario do caso de uso com fakes.
- [ ] Teste de adapter com servidor falso, cobrindo timeout, retry e falha de auth.
- [ ] Teste de integracao para cifragem de credencial, TTL de cache e invalidacao.
- [ ] Atualizar contrato publicado e docs se os endpoints de gerenciamento mudaram.
- [ ] Atualizar runbook de rotacao/revogacao de token e comportamento em indisponibilidade do provider.

## Definition of Done

- Integracao esta atras de port, adapter e wiring na composition root.
- Credenciais cifradas em repouso, cacheadas com TTL limitado e invalidadas na atualizacao.
- Logs, traces e metricas preservam correlacao e nunca expoem token nem payload sensivel.
- Testes provam sucesso, falha de auth, comportamento de timeout/retry, cache hit/miss e invalidacao.
- Spec, contrato publicado e runbooks relevantes atualizados.
