# {{PROJECT_TITLE}} Boundaries

Use this reference when a request risks mixing framework responsibilities with application responsibilities.

## {{PROJECT_TITLE}} core may include

- agent runtime and invocation contracts
- tools and tool middleware
- memory and conversation framework contracts
- workflow contracts
- supervisor/sub-agent contracts
- provider adapters as reusable framework integrations
- HTTP/API exposure as an optional framework server
- OpenAPI artifacts
- local observability hooks/helpers
- conformance suites
- examples/reference consumer
- release/versioning governance

## Consumer application should own

- domain rules
- product-specific auth/RBAC
- external queues
- third-party business integrations
- data repositories for product entities
- channel-specific payload formats
- deployment topology
- hosted platform concerns
- business observability dashboards

## Decision rule

If a requested change is useful for many {{PROJECT_SLUG}} consumers and can be expressed as a reusable contract, it may belong in {{PROJECT_SLUG}}.

If it encodes one product's domain, one channel's payload, one tenant's policy, or one external vendor workflow, keep it out of {{PROJECT_SLUG}} core and expose extension points instead.
