#!/usr/bin/env python3
"""Governed change guard - beforeShellExecution hook logic.

Asks the user to register the governed-change trio (BLG-NNNN + dual-spec +
reviewed plan) when a `git commit`/`git push` or `gh pr create` is about to run
and the staged/working diff touches governed scope (pkg/**, internal/**, api/**,
migrations/**) without an active interactive track in
automation/INTERACTIVE_STATE.json.

Best-effort: the definitive gate is scripts/check-governed-change.sh in
pre-pr/CI. Fail-open on internal errors (emits allow); hooks.json `failClosed`
covers process crash/timeout. Reads the hook event JSON from stdin.
"""
from __future__ import annotations

import json
import re
import subprocess
import sys
from pathlib import Path

GOVERNED = re.compile(r"^(pkg/|internal/|api/|migrations/)")
TRIGGER = re.compile(r"git\s+commit|git\s+push|gh\s+pr\s+create")
ITEM_ID = re.compile(r"BLG-\d{4}")
INACTIVE_STATUS = {"idle", "track_closed", ""}


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

    def git(*args: str) -> str:
        try:
            return subprocess.run(
                ["git", "-C", root, *args], capture_output=True, text=True
            ).stdout
        except Exception:
            return ""

    files: set[str] = set()
    for spec in (["diff", "--cached", "--name-only"], ["diff", "--name-only"]):
        for f in git(*spec).splitlines():
            if f.strip():
                files.add(f.strip())
    if not any(GOVERNED.match(f) for f in files):
        emit({"permission": "allow"})

    # An active interactive track is the signal that the trio was registered.
    # The item id may live anywhere in current_request, so the whole state blob
    # is scanned rather than one fixed field.
    active = False
    item = ""
    state = Path(root) / "automation/INTERACTIVE_STATE.json"
    if state.exists():
        try:
            text = state.read_text(encoding="utf-8")
            parsed = json.loads(text)
            status = str(parsed.get("status") or "").strip()
            active = status not in INACTIVE_STATUS and parsed.get("current_request") is not None
            found = ITEM_ID.search(text)
            item = found.group(0) if found else ""
        except Exception:
            active = False

    item_ok = False
    if item:
        backlog = Path(root) / "docs" / "backlog" / "Backlog.md"
        if backlog.is_file():
            try:
                item_ok = re.search(
                    rf"^###\s+{re.escape(item)}\b", backlog.read_text(encoding="utf-8"), re.MULTILINE
                ) is not None
            except Exception:
                item_ok = False

    if active and item_ok:
        emit({"permission": "allow"})

    emit({
        "permission": "ask",
        "user_message": (
            "Mudanca em escopo governado (pkg/internal/api/migrations) sem trilha ativa registrada. "
            "Registre BLG-NNNN no backlog + par dual-spec + plano revisado antes de commit/push "
            "(skills/29-governed-change-workflow)."
        ),
        "agent_message": (
            "Governed-change guard: escopo governado sem trilha ativa em "
            "automation/INTERACTIVE_STATE.json com item BLG-NNNN presente em docs/backlog/Backlog.md. "
            "Aplique skills/29-governed-change-workflow (backlog + dual-spec + plano com review). "
            "Gate definitivo: scripts/check-governed-change.sh no pre-pr/CI."
        ),
    })


try:
    main()
except SystemExit:
    raise
except Exception:
    emit({"permission": "allow"})
