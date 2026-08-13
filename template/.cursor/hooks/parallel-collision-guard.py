#!/usr/bin/env python3
"""Parallel collision guard - beforeShellExecution hook logic.

Roda scripts/check-parallel-collision.sh antes de `git push` / `gh pr create` e
pede confirmacao quando um ID sequencial desta branch (BLG-NNNN, specs/NNN,
migrations/NNNN) ja esta ocupado na base, num PR aberto ou numa branch remota.

Por que so em push/PR, e nao em commit: e quando a colisao vira publica e cara de
desfazer, e evita consulta de rede a cada commit.

Best-effort: o gate definitivo e scripts/check-parallel-collision.sh no pre-pr.
Fail-open em erro interno (emite allow); hooks.json `failClosed` cobre
crash/timeout. Le o JSON do evento em stdin.
"""
from __future__ import annotations

import json
import re
import subprocess
import sys
from pathlib import Path

TRIGGER = re.compile(r"git\s+push|gh\s+pr\s+create")
TIMEOUT_SECONDS = 45


def emit(obj: dict) -> None:
    sys.stdout.write(json.dumps(obj))
    sys.exit(0)


def main() -> None:
    try:
        raw = sys.stdin.read()
        data = json.loads(raw) if raw.strip() else {}
    except Exception:
        emit({"permission": "allow"})

    if not TRIGGER.search(data.get("command") or ""):
        emit({"permission": "allow"})

    try:
        root = subprocess.run(
            ["git", "rev-parse", "--show-toplevel"], capture_output=True, text=True
        ).stdout.strip()
    except Exception:
        root = ""
    if not root:
        emit({"permission": "allow"})

    gate = Path(root) / "scripts/check-parallel-collision.sh"
    if not gate.exists():
        emit({"permission": "allow"})

    try:
        proc = subprocess.run(
            ["bash", str(gate)],
            capture_output=True,
            text=True,
            cwd=root,
            timeout=TIMEOUT_SECONDS,
        )
    except Exception:
        # Rede lenta ou indisponivel nao pode virar bloqueio: o gate do pre-pr
        # roda de novo antes do PR.
        emit({"permission": "allow"})

    if proc.returncode == 0:
        emit({"permission": "allow"})

    detail = (proc.stderr or proc.stdout or "").strip()
    collisions = [
        line.split("FAIL:", 1)[1].strip()
        for line in detail.splitlines()
        if "FAIL:" in line
    ]
    resumo = "; ".join(collisions) if collisions else detail[:400]

    emit(
        {
            "permission": "ask",
            "user_message": (
                "Colisao de ID entre sessoes paralelas detectada antes de publicar: "
                f"{resumo}. Renumere antes do push (skills/31-parallel-session-coordination, "
                "Protocolo B)."
            ),
            "agent_message": (
                "Parallel collision guard: "
                f"{resumo}. Recalcule o ID com sort numerico contra a base, PRs abertos e "
                "branches remotas; revalide as referencias cruzadas (backlog, plano, corpo do "
                "PR, specs) e renomeie a branch se ainda nao publicada. "
                "Gate definitivo: scripts/check-parallel-collision.sh."
            ),
        }
    )


try:
    main()
except SystemExit:
    raise
except Exception:
    emit({"permission": "allow"})
