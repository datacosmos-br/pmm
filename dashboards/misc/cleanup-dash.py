#!/usr/bin/env python3

import argparse
import copy
import json
import pathlib
import sys


def set_dashboard_id_to_null(dashboard):
    """To remove any dashboard id. New one is set by grafana."""
    for element in enumerate(dashboard.copy()):
        if "id" in element:
            dashboard["id"] = None

    return dashboard


def set_editable(dashboard):
    """Set Editable Dashboard."""
    if "editable" not in dashboard:
        return dashboard

    dashboard["editable"] = False
    return dashboard


def set_refresh(dashboard):
    """Set Dashboard refresh."""
    if "refresh" not in dashboard:
        return dashboard

    dashboard["refresh"] = False
    return dashboard


def set_timezone(dashboard):
    """Set Dashboard Time zone."""
    dashboard["timezone"] = ""
    return dashboard


def set_time(dashboard):
    """Set Dashboard Time Range."""
    dashboard["time"]["from"] = "now-12h"
    dashboard["time"]["to"] = "now"
    return dashboard


def main() -> None:
    parser = argparse.ArgumentParser(description="Dashboard cleaner")
    parser.add_argument("dashboard_file", type=str, help="dashboard file to cleanup")
    parser.add_argument("--check-only", action="store_true", help="check only mode")
    args = parser.parse_args()

    with pathlib.Path(args.dashboard_file).open(encoding="utf-8") as dashboard_file:
        dashboard = json.loads(dashboard_file.read())
        raw_dashboard = copy.deepcopy(dashboard)

    CLEANUPERS = [
        set_editable,
        set_time,
        set_timezone,
        set_refresh,
        set_dashboard_id_to_null,
    ]

    for func in CLEANUPERS:
        dashboard = func(dashboard)

    dashboard_json = json.dumps(
        dashboard,
        sort_keys=True,
        indent=4,
        separators=(",", ": "),
        ensure_ascii=False,
    )

    if args.check_only:
        if raw_dashboard == dashboard:
            sys.exit(0)
        else:

            def jv(v):
                return json.dumps(v)

            issues = []
            if raw_dashboard.get("editable") != dashboard.get("editable"):
                issues.append(
                    f"  editable: {jv(raw_dashboard.get('editable'))} -> {jv(dashboard.get('editable'))}",
                )
            if raw_dashboard.get("refresh") != dashboard.get("refresh"):
                issues.append(
                    f"  refresh: {jv(raw_dashboard.get('refresh'))} -> {jv(dashboard.get('refresh'))}",
                )
            if raw_dashboard.get("timezone") != dashboard.get("timezone"):
                issues.append(
                    f"  timezone: {jv(raw_dashboard.get('timezone'))} -> {jv(dashboard.get('timezone'))}",
                )
            if raw_dashboard.get("time", {}).get("from") != dashboard.get(
                "time",
                {},
            ).get("from"):
                issues.append(
                    f"  time.from: {jv(raw_dashboard.get('time', {}).get('from'))} -> {jv(dashboard.get('time', {}).get('from'))}",
                )
            if raw_dashboard.get("time", {}).get("to") != dashboard.get("time", {}).get(
                "to",
            ):
                issues.append(
                    f"  time.to: {jv(raw_dashboard.get('time', {}).get('to'))} -> {jv(dashboard.get('time', {}).get('to'))}",
                )
            if raw_dashboard.get("id") != dashboard.get("id"):
                issues.append(
                    f"  id: {jv(raw_dashboard.get('id'))} -> {jv(dashboard.get('id'))}",
                )
            for _issue in issues:
                pass
            sys.exit(1)

    with pathlib.Path(args.dashboard_file).open(
        "w",
        encoding="utf-8",
    ) as dashboard_file:
        dashboard_file.write(dashboard_json)
        dashboard_file.write("\n")


if __name__ == "__main__":
    main()
