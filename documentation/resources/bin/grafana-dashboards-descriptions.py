#! /usr/bin/env python3
"""Writes md files with PMM dashboard panel titles/descriptions to the cwd.

Reads every dashboard JSON from a local clone of
https://github.com/percona/grafana-dashboards/ (expected at ``repo_src``) and
emits one ``dashboard-<name>.md`` file per dashboard, preserving the historical
output contract (titles, image links, and panel descriptions).
"""

import json
import pathlib

# Path to local git clone of https://github.com/percona/grafana-dashboards/
REPO_SRC = pathlib.Path("../../grafana-dashboards/dashboards/")

PANEL_TYPES = {"graph", "singlestat"}


def _write_dashboard(path: pathlib.Path) -> None:
    """Renders one dashboard JSON file into its markdown description file."""
    # Title and image come from the filename.
    title = path.name.replace("_", " ").replace(".json", "")
    image = "PMM_" + path.name.replace(".json", "") + ".jpg"
    titlelc = path.name.replace("_", "-").replace(".json", "").lower()

    with path.open(encoding="utf-8") as fp:
        dashboard = json.load(fp)

    md_path = pathlib.Path("dashboard-" + titlelc + ".md")
    with md_path.open("w", encoding="utf-8") as md:
        md.write("# " + title + "\n\n")
        md.write("![image](../images/" + image + ")\n\n")

        for panel in dashboard["panels"]:
            if panel["type"] != "row":
                continue
            if "title" in panel:
                md.write("## " + panel["title"] + "\n\n")
            if "description" in panel:
                md.write(panel["description"] + "\n\n")
            for sub in panel.get("panels", []):
                if (
                    sub["type"] in PANEL_TYPES
                    and "title" in sub
                    and "description" in sub
                ):
                    md.write("### " + sub["title"] + "\n\n")
                    md.write(sub["description"] + "\n\n")


def main() -> None:
    """Entry point: renders all dashboards found under REPO_SRC."""
    if not REPO_SRC.is_dir():
        return
    for path in sorted(REPO_SRC.glob("*.json")):
        _write_dashboard(path)


if __name__ == "__main__":
    main()
