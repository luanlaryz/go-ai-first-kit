"""Shared high-risk classification for plan review.

Both `scripts/review-plan.py` (which refuses to WRITE an OPERATOR_FALLBACK
approval for a high-risk plan) and `scripts/verify-plan-review.sh` (which refuses
to ACCEPT one) classify risk from this single module. Two copies of the list would
drift, and the drift would silently reopen the fallback path the gate exists to
close.

Path markers are matched as substrings. Word markers are matched with word
boundaries and explicit alternatives, so "author" does not count as "auth" and
"secretary" does not count as "secret": inflated risk blocks legitimate work and
teaches people to distrust the gate.
"""

from __future__ import annotations

import re

PATH_MARKERS = {
    "pkg/": "public API surface",
    "api/": "published contract",
    "migrations/": "database migration",
    "docs/decisions/": "ADR",
}

WORD_MARKERS = {
    r"auth|authentication|authorization|autenticacao|autorizacao": "authentication or authorization",
    r"security|seguranca": "security",
    r"secret|secrets|credential|credentials|credencial|credenciais": "secret handling",
    r"adr|adrs": "ADR",
    r"breaking": "breaking change",
    r"concurrency|concurrent|concorrencia|goroutine|goroutines": "concurrency",
    r"queue|queues|fila|filas|dlq": "queue semantics",
    r"replay": "replay",
    r"idempotency|idempotencia|idempotente": "idempotency",
    r"ordering|ordenacao": "ordering",
}


def detect_high_risk(raw: str) -> list[str]:
    """Return the sorted list of high-risk reasons matched by the plan text."""
    text = raw.lower()
    reasons: list[str] = []
    for marker, reason in PATH_MARKERS.items():
        if marker in text and reason not in reasons:
            reasons.append(reason)
    for pattern, reason in WORD_MARKERS.items():
        if reason in reasons:
            continue
        if re.search(rf"\b(?:{pattern})\b", text):
            reasons.append(reason)
    return sorted(reasons)
