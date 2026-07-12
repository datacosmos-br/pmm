#!/usr/bin/env python3
"""Normalizes Grafana dashboard JSON files to the PMM committed form.

Applies the canonical cleanup (editable/refresh off, empty timezone,
``now-12h``..``now`` range, ``"id": null``) and rewrites the file with the
deterministic dump format. With ``--check-only`` the file is left untouched:
exit 0 when already clean, exit 1 with the offending fields on stderr
otherwise. Consumed by ``.github/workflows/dashboards.yml`` (stdlib only —
the workflow runner has no workspace dependencies).
"""

import copy
import json
import pathlib
import sys
from typing import TypeAlias

# NOTE (p/ outro agente): JsonValue/JsonMap modelam o contrato JSON externo
# (arquivos de dashboard do Grafana) — stdlib typing, sem flext (pmm nao e
# projeto cosmos-*; CI roda no GitHub Actions sem dependencias do workspace).
# Forma classica (string forward-ref): compativel com qualquer python3 do
# runner GitHub Actions; o statement `type` (PEP 695) exigiria 3.12+.
JsonValue: TypeAlias = (
    "str | int | float | bool | list[JsonValue] | dict[str, JsonValue] | None"
)
JsonMap: TypeAlias = dict[str, JsonValue]

_CLEANED_FROM = "now-12h"
_CLEANED_TO = "now"


def set_dashboard_id_to_null(dashboard: JsonMap) -> JsonMap:
    """Removes any dashboard id (a new one is set by Grafana)."""
    if "id" in dashboard:
        dashboard["id"] = None
    return dashboard


def set_editable(dashboard: JsonMap) -> JsonMap:
    """Disables dashboard editing when the flag is present."""
    if "editable" in dashboard:
        dashboard["editable"] = False
    return dashboard


def set_refresh(dashboard: JsonMap) -> JsonMap:
    """Disables dashboard auto-refresh when the flag is present."""
    if "refresh" in dashboard:
        dashboard["refresh"] = False
    return dashboard


def set_timezone(dashboard: JsonMap) -> JsonMap:
    """Clears the dashboard timezone (defaults to the browser timezone)."""
    dashboard["timezone"] = ""
    return dashboard


def set_time(dashboard: JsonMap) -> JsonMap:
    """Resets the dashboard time range to the canonical default."""
    time_range = dashboard.setdefault("time", {})
    if isinstance(time_range, dict):
        time_range["from"] = _CLEANED_FROM
        time_range["to"] = _CLEANED_TO
    return dashboard


def _clean(dashboard: JsonMap) -> JsonMap:
    """Applies every cleanup step in the canonical order."""
    for cleanup in (
        set_editable,
        set_time,
        set_timezone,
        set_refresh,
        set_dashboard_id_to_null,
    ):
        dashboard = cleanup(dashboard)
    return dashboard


def _report_issues(raw: JsonMap, cleaned: JsonMap) -> None:
    """Writes the fields that differ between the raw and cleaned dashboard."""
    fields = ("editable", "refresh", "timezone", "id")
    for field in fields:
        if raw.get(field) != cleaned.get(field):
            sys.stderr.write(
                f"  {field}: {json.dumps(raw.get(field))} -> {json.dumps(cleaned.get(field))}\n",
            )
    raw_time = raw.get("time", {})
    cleaned_time = cleaned.get("time", {})
    if isinstance(raw_time, dict) and isinstance(cleaned_time, dict):
        for bound in ("from", "to"):
            if raw_time.get(bound) != cleaned_time.get(bound):
                sys.stderr.write(
                    f"  time.{bound}: {json.dumps(raw_time.get(bound))} -> "
                    f"{json.dumps(cleaned_time.get(bound))}\n",
                )


def _parse_cli(argv: list[str]) -> tuple[pathlib.Path, bool]:
    """Parses ``[--check-only] <dashboard_file>`` (zero-dependency CLI).

    NOTE (p/ outro agente): TID251 bane argparse fora de cosmos-*; o padrao
    manual aprovado e este (precedente: scripts/legado/argocd/*).
    """
    check_only = False
    positional: list[str] = []
    for arg in argv:
        if arg == "--check-only":
            check_only = True
        else:
            positional.append(arg)
    if len(positional) != 1:
        sys.stderr.write("usage: cleanup-dash.py [--check-only] <dashboard_file>\n")
        sys.exit(2)
    return pathlib.Path(positional[0]), check_only


def main(argv: list[str] | None = None) -> None:
    """Entry point: cleans one dashboard file or checks its cleanliness."""
    path, check_only = _parse_cli(sys.argv[1:] if argv is None else argv)
    dashboard: JsonMap = json.loads(path.read_text(encoding="utf-8"))
    raw_dashboard = copy.deepcopy(dashboard)

    dashboard = _clean(dashboard)

    if check_only:
        if raw_dashboard == dashboard:
            sys.exit(0)
        _report_issues(raw_dashboard, dashboard)
        sys.exit(1)

    dashboard_json = json.dumps(
        dashboard,
        sort_keys=True,
        indent=4,
        separators=(",", ": "),
        ensure_ascii=False,
    )
    path.write_text(dashboard_json + "\n", encoding="utf-8")


if __name__ == "__main__":
    main()
