#! /usr/bin/env python3
# Write md files containing PMM dashboard panel titles and descriptions to current dir

import glob
import json
import pathlib

# Path to local git clone of https://github.com/percona/grafana-dashboards/
repo_src = "../../grafana-dashboards/dashboards/"

if not pathlib.Path(repo_src).is_dir():
    exit

# Dict of dashboard files
dashboard_files = glob.glob(repo_src + "*.json")
# For each, open the file, read in fields
for filename in dashboard_files:
    with pathlib.Path(filename).open(encoding="utf-8") as fp:
        # Title and image come from filename
        title = pathlib.Path(filename).name.replace("_", " ").replace(".json", "")
        image = "PMM_" + pathlib.Path(filename).name.replace(".json", "") + ".jpg"
        titlelc = (
            pathlib.Path(filename).name.replace("_", "-").replace(".json", "").lower()
        )

        with pathlib.Path("dashboard-" + titlelc + ".md").open(
            "w",
            encoding="utf-8",
        ) as md:
            x = json.load(fp)
            md.write("# " + title + "\n\n")
            md.write("![image](../images/" + image + ")\n\n")

            for p in x["panels"]:
                if p["type"] == "row":
                    if "title" in p:
                        md.write("## " + p["title"] + "\n\n")

                    if "description" in p:
                        md.write(p["description"] + "\n\n")

                    if "panels" in p:
                        for p2 in p["panels"]:
                            if p2["type"] in {"graph", "singlestat"}:
                                if "title" in p2 and "description" in p2:
                                    md.write("### " + p2["title"] + "\n\n")
                                    md.write(p2["description"] + "\n\n")

        md.close
    fp.close
